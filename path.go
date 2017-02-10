package minecraft

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/MJKWoolnough/byteio"
	"github.com/MJKWoolnough/memio"
	"github.com/MJKWoolnough/minecraft/nbt"
)

var filename *regexp.Regexp

// The Path interface allows the minecraft level to be created from/saved
// to different formats.
type Path interface {
	// Returns an empty nbt.Tag (TagEnd) when chunk does not exists
	GetChunk(int32, int32) (nbt.Tag, error)
	SetChunk(...nbt.Tag) error
	RemoveChunk(int32, int32) error
	ReadLevelDat() (nbt.Tag, error)
	WriteLevelDat(nbt.Tag) error
}

// Compression convenience constants
const (
	GZip byte = 1
	Zlib byte = 2
)

// FilePath implements the Path interface and provides a standard minecraft
// save format.
type FilePath struct {
	dirname   string
	lock      int64
	dimension string
}

// NewFilePath constructs a new directory based path to read from.
func NewFilePath(dirname string) (*FilePath, error) {
	dirname = path.Clean(dirname)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		return nil, err
	}
	p := &FilePath{dirname: dirname}
	return p, p.Lock()
}

type stickyEndianSeeker struct {
	byteio.StickyWriter
	io.Seeker
}

func (s *stickyEndianSeeker) Seek(offset int64, whence int) (int64, error) {
	if s.StickyWriter.Err != nil {
		return 0, s.StickyWriter.Err
	}
	var n int64
	n, s.StickyWriter.Err = s.Seeker.Seek(offset, whence)
	return n, s.StickyWriter.Err
}

// NewFilePathDimension create a new FilePath, but with the option to set the
// dimension that chunks are loaded from.
//
// Example. Dimension -1 == The Nether
//          Dimension  1 == The End
func NewFilePathDimension(dirname string, dimension int) (*FilePath, error) {
	fp, err := NewFilePath(dirname)
	if dimension != 0 {
		fp.dimension = "DIM" + strconv.Itoa(dimension)
	}
	return fp, err
}

func (p *FilePath) getRegionPath(x, z int32) string {
	return path.Join(p.dirname, p.dimension, "region", "r."+strconv.FormatInt(int64(x>>5), 10)+"."+strconv.FormatInt(int64(z>>5), 10)+".mca")
}

// GetChunk returns the chunk at chunk coords x, z.
func (p *FilePath) GetChunk(x, z int32) (nbt.Tag, error) {
	if !p.HasLock() {
		return nbt.Tag{}, ErrNoLock
	}
	f, err := os.Open(p.getRegionPath(x>>5, z>>5))
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return nbt.Tag{}, err
	}
	defer f.Close()
	pos := int64((z&31)<<5 | (x & 31))
	if _, err = f.Seek(4*pos, io.SeekStart); err != nil {
		return nbt.Tag{}, err
	}

	be := byteio.BigEndianReader{Reader: f}

	locationSize, _, err := be.ReadUint32()
	if locationSize>>8 == 0 {
		return nbt.Tag{}, nil
	} else if _, err = f.Seek(int64(locationSize>>8<<12), io.SeekStart); err != nil {
		return nbt.Tag{}, err
	}

	reader := io.LimitReader(f, int64(locationSize&255<<12))

	length, _, err := be.ReadUint32()
	if err != nil {
		return nbt.Tag{}, err
	}

	reader = io.LimitReader(reader, int64(length))
	compression, _, err := be.ReadUint8()
	if err != nil {
		return nbt.Tag{}, err
	}

	switch compression {
	case GZip:
		gReader, err := gzip.NewReader(reader)
		if err != nil {
			return nbt.Tag{}, err
		}
		defer gReader.Close()
		reader = gReader
	case Zlib:
		if reader, err = zlib.NewReader(reader); err != nil {
			return nbt.Tag{}, err
		}
	default:
		return nbt.Tag{}, UnknownCompression{compression}
	}

	return nbt.Decode(reader)
}

type rc struct {
	pos int32
	buf memio.Buffer
}

// SetChunk saves multiple chunks at once, possibly returning a MultiError if
// multiple errors were encountered.
func (p *FilePath) SetChunk(data ...nbt.Tag) error {
	if !p.HasLock() {
		return ErrNoLock
	}
	regions := make(map[uint64][]rc)
	var (
		poses  []uint64
		errors []error
	)
	for _, d := range data {
		x, z, err := chunkCoords(d)
		if err != nil {
			errors = append(errors, FilePathSetError{x, z, err})
			continue
		}
		pos := uint64(z)<<32 | uint64(uint32(x))
		for _, p := range poses {
			if p == pos {
				errors = append(errors, ConflictError{x, z})
				continue
			}
		}
		poses = append(poses, pos)
		r := uint64(z>>5)<<32 | uint64(uint32(x>>5))
		reg := rc{pos: (z&31)<<5 | (x & 31)}
		zl := zlib.NewWriter(&reg.buf)
		err = nbt.Encode(zl, d)
		zl.Close()
		if err != nil {
			errors = append(errors, FilePathSetError{x, z, err})
			continue
		}
		if regions[r] == nil {
			regions[r] = []rc{reg}
		} else {
			regions[r] = append(regions[r], reg)
		}
	}
	for rID, chunks := range regions {
		x, z := int32(rID&0xffffffff), int32(rID>>32)
		if err := p.setChunks(x, z, chunks); err != nil {
			errors = append(errors, &FilePathSetError{x, z, err})
		}
	}
	if len(errors) > 0 {
		return MultiError{errors}
	}
	return nil
}

type sia []uint32

func (s sia) Len() int {
	return 1024
}

func (s sia) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sia) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (p *FilePath) setChunks(x, z int32, chunks []rc) error {
	if err := os.MkdirAll(path.Join(p.dirname, p.dimension, "region"), 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(p.getRegionPath(x, z), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	var (
		bytes     [4096]byte
		positions [1024]uint32
	)
	pBytes := memio.Buffer(bytes[:])
	if _, err = io.ReadFull(f, pBytes); err != nil && err != io.EOF {
		return err
	}
	posr := byteio.BigEndianReader{Reader: &pBytes}
	for i := 0; i < 1024; i++ {
		positions[i], _, _ = posr.ReadUint32()
	}
	var todoChunks []rc
	bew := stickyEndianSeeker{byteio.StickyWriter{Writer: byteio.BigEndianWriter{Writer: f}}, f}
	for _, chunk := range chunks {
		newSize := uint32(len(chunk.buf)+5) >> 12
		if uint32(len(chunk.buf))&4095 > 0 {
			newSize++
		}
		if positions[chunk.pos]&255 == newSize {
			bew.Seek(4*int64(chunk.pos)+4096, io.SeekStart) // Write the time, then the data
			bew.WriteUint32(uint32(time.Now().Unix()))
			bew.Seek(int64(positions[chunk.pos])>>8<<12, io.SeekStart)
			bew.WriteUint32(uint32(len(chunk.buf)) + 1)
			bew.WriteUint8(Zlib)
			bew.Write(chunk.buf)
		} else {
			todoChunks = append(todoChunks, chunk)
			positions[chunk.pos] = 0
		}
	}
	if bew.Err != nil {
		return bew.Err
	}
	for _, chunk := range todoChunks {
		sort.Sort(sia(positions[:]))
		newPosition := (positions[1023] >> 8) + (positions[1023] & 255)
		if newPosition == 0 {
			newPosition = 2
		}
		lastPos := uint32(2)
		smallest := uint32(0xffffffff)
		writeLastByte := true
		newSize := uint32(len(chunk.buf) + 5)
		if newSize&4095 > 0 {
			newSize >>= 12
			newSize++
		} else {
			newSize >>= 12
		}
		// Find earliest, closest match in size for least fragmentation.
		for i := 0; i < 1024; i++ {
			loc := positions[i] >> 8
			if loc > 0 {
				size := positions[i] & 255
				if space := loc - lastPos; space >= newSize && space < smallest {
					smallest = space
					newPosition = lastPos
					writeLastByte = false // by definition it has data that is after it now, so no need to make up to mod 4096 bytes
				}
				lastPos = loc + size
			}
		}
		positions[0] = newPosition<<8 | newSize&255
		// Write the new position
		bew.Seek(4*int64(chunk.pos), io.SeekStart)
		bew.WriteUint32(positions[0])
		bew.Seek(4*(int64(chunk.pos)+1024), io.SeekStart) // Write the time, then the data
		bew.WriteUint32(uint32(time.Now().Unix()))
		bew.Seek(int64(newPosition)<<12, io.SeekStart)
		bew.WriteUint32(uint32(len(chunk.buf)) + 1)
		bew.WriteUint8(Zlib)
		bew.Write(chunk.buf)
		if writeLastByte { // Make filesize mod 4096 (for minecraft compatibility)
			bew.Seek(int64(newPosition+newSize)<<12-1, io.SeekStart)
			bew.WriteUint8(0)
		}
	}
	return bew.Err
}

// RemoveChunk deletes the chunk at chunk coords x, z.
func (p *FilePath) RemoveChunk(x, z int32) error {
	if !p.HasLock() {
		return ErrNoLock
	}
	chunkX := x & 31
	regionX := x >> 5
	chunkZ := z & 31
	regionZ := z >> 5
	f, err := os.OpenFile(p.getRegionPath(regionX, regionZ), os.O_WRONLY, 0666)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Seek(int64(chunkZ<<5|chunkX)*4, io.SeekStart); err != nil {
		return err
	}
	_, err = f.Write([]byte{0, 0, 0, 0})
	return err
}

// ReadLevelDat returns the level data.
func (p *FilePath) ReadLevelDat() (nbt.Tag, error) {
	if !p.HasLock() {
		return nbt.Tag{}, ErrNoLock
	}
	f, err := os.Open(path.Join(p.dirname, "level.dat"))
	if os.IsNotExist(err) {
		return nbt.Tag{}, nil
	} else if err != nil {
		return nbt.Tag{}, err
	}
	defer f.Close()
	g, err := gzip.NewReader(f)
	if err != nil {
		return nbt.Tag{}, err
	}
	data, err := nbt.Decode(g)
	return data, err
}

// WriteLevelDat Writes the level data.
func (p *FilePath) WriteLevelDat(data nbt.Tag) error {
	if !p.HasLock() {
		return ErrNoLock
	}
	f, err := os.OpenFile(path.Join(p.dirname, "level.dat"), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil
	}
	defer f.Close()
	g := gzip.NewWriter(f)
	defer g.Close()
	err = nbt.Encode(g, data)
	return err
}

// GetRegions returns a list of region x,z coords of all generated regions.
func (p *FilePath) GetRegions() [][2]int32 {
	files, _ := ioutil.ReadDir(path.Join(p.dirname, p.dimension, "region"))
	var toRet [][2]int32
	for _, file := range files {
		if !file.IsDir() {
			if nums := filename.FindStringSubmatch(file.Name()); nums != nil {
				x, _ := strconv.ParseInt(nums[1], 10, 32)
				z, _ := strconv.ParseInt(nums[2], 10, 32)
				toRet = append(toRet, [2]int32{int32(x), int32(z)})
			}
		}
	}
	return toRet
}

// GetChunks returns a list of all chunks within a region with coords x,z
func (p *FilePath) GetChunks(x, z int32) ([][2]int32, error) {
	if !p.HasLock() {
		return nil, ErrNoLock
	}
	f, err := os.Open(p.getRegionPath(x, z))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	pBytes := make(memio.Buffer, 4096)
	if _, err = io.ReadFull(f, pBytes); err != nil {
		return nil, err
	}

	baseX := x << 5
	baseZ := z << 5

	var toRet [][2]int32
	posr := byteio.BigEndianReader{Reader: &pBytes}
	for i := 0; i < 1024; i++ {
		if n, _, _ := posr.ReadUint32(); n > 0 {
			toRet = append(toRet, [2]int32{baseX + int32(i&31), baseZ + int32(i>>5)})
		}
	}
	return toRet, nil
}

// HasLock returns whether or not another program has taken the lock.
func (p *FilePath) HasLock() bool {
	r, err := os.Open(path.Join(p.dirname, "session.lock"))
	if err != nil {
		return false
	}
	defer r.Close()
	buf := make(memio.Buffer, 9)
	n, err := io.ReadFull(r, buf)
	if n != 8 || err != io.ErrUnexpectedEOF {
		return false
	}
	b, _, _ := byteio.BigEndianReader{Reader: &buf}.ReadInt64()
	return b == p.lock
}

// Lock will retake the lock file if it has been lost. May cause corruption.
func (p *FilePath) Lock() error {
	if p.HasLock() {
		return nil
	}
	p.lock = time.Now().UnixNano() / 1000000 // ms
	session := path.Join(p.dirname, "session.lock")
	f, err := os.Create(session)
	if err != nil {
		return err
	}
	_, err = byteio.BigEndianWriter{Writer: f}.WriteUint64(uint64(p.lock))
	f.Close()
	if err != nil {
		return err
	}
	return nil
}

// Defrag rewrites a region file to reduce wasted space.
func (p *FilePath) Defrag(x, z int32) error {
	if !p.HasLock() {
		return ErrNoLock
	}
	f, err := os.OpenFile(p.getRegionPath(x, z), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	locationSizes := make(memio.Buffer, 4096)
	if _, err = io.ReadFull(f, locationSizes); err != nil {
		return err
	}

	var (
		data       [1024][]byte
		locations  [4096]byte
		currSector uint32 = 2
	)
	locationr := byteio.BigEndianReader{Reader: &locationSizes}
	l := memio.Buffer(locations[:0])
	locationw := byteio.BigEndianWriter{Writer: &l}
	for i := 0; i < 1024; i++ {
		locationSize, _, _ := locationr.ReadUint32()
		if locationSize>>8 == 0 {
			continue
		} else if _, err = f.Seek(int64(locationSize>>8<<12), io.SeekStart); err != nil {
			return err
		}

		data[i] = make([]byte, locationSize&255<<12)

		if _, err := io.ReadFull(f, data[i]); err != nil {
			return err
		}

		locationw.WriteUint32(currSector<<8 | locationSize&255)

		currSector += locationSize & 255
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = f.Write(locations[:])
	if err != nil {
		return err
	}

	_, err = f.Seek(8192, io.SeekStart)
	if err != nil {
		return err // Try and recover first?
	}

	for _, d := range data {
		if len(d) > 0 {
			_, err = f.Write(d)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// MemPath is an in memory minecraft level format that implements the Path interface.
type MemPath struct {
	level  memio.Buffer
	chunks map[uint64]memio.Buffer
}

// NewMemPath creates a new MemPath implementation.
func NewMemPath() *MemPath {
	return &MemPath{chunks: make(map[uint64]memio.Buffer)}
}

// GetChunk returns the chunk at chunk coords x, z.
func (m *MemPath) GetChunk(x, z int32) (nbt.Tag, error) {
	pos := uint64(z)<<32 | uint64(uint32(x))
	c := m.chunks[pos]
	if c == nil {
		return nbt.Tag{}, nil
	}
	return m.read(c)
}

// SetChunk saves multiple chunks at once.
func (m *MemPath) SetChunk(data ...nbt.Tag) error {
	for _, d := range data {
		x, z, err := chunkCoords(d)
		if err != nil {
			return err
		}
		var buf memio.Buffer
		if err = m.write(d, &buf); err != nil {
			return err
		}
		pos := uint64(z)<<32 | uint64(uint32(x))
		m.chunks[pos] = buf
	}
	return nil
}

// RemoveChunk deletes the chunk at chunk coords x, z.
func (m *MemPath) RemoveChunk(x, z int32) error {
	pos := uint64(z)<<32 | uint64(uint32(x))
	delete(m.chunks, pos)
	return nil
}

// ReadLevelDat Returns the level data.
func (m *MemPath) ReadLevelDat() (nbt.Tag, error) {
	if len(m.level) == 0 {
		return nbt.Tag{}, nil
	}
	return m.read(m.level)
}

// WriteLevelDat Writes the level data.
func (m *MemPath) WriteLevelDat(data nbt.Tag) error {
	return m.write(data, &m.level)
}

func (m *MemPath) read(buf memio.Buffer) (nbt.Tag, error) {
	z, err := zlib.NewReader(&buf)
	if err != nil {
		return nbt.Tag{}, err
	}
	data, err := nbt.Decode(z)
	return data, err
}

func (m *MemPath) write(data nbt.Tag, buf *memio.Buffer) error {
	z := zlib.NewWriter(buf)
	defer z.Close()
	err := nbt.Encode(z, data)
	return err
}

func chunkCoords(data nbt.Tag) (int32, int32, error) {
	if data.TagID() != nbt.TagCompound {
		return 0, 0, WrongTypeError{"[Chunk Base]", nbt.TagCompound, data.TagID()}
	}
	lTag := data.Data().(nbt.Compound).Get("Level")
	if lTag.TagID() == 0 {
		return 0, 0, MissingTagError{"[Chunk Base]->Level"}
	} else if lTag.TagID() != nbt.TagCompound {
		return 0, 0, WrongTypeError{"[Chunk Base]->Level", nbt.TagCompound, lTag.TagID()}
	}

	lCmp := lTag.Data().(nbt.Compound)

	xPos := lCmp.Get("xPos")
	if xPos.TagID() == 0 {
		return 0, 0, MissingTagError{"[Chunk Base]->Level->xPos"}
	} else if xPos.TagID() != nbt.TagInt {
		return 0, 0, WrongTypeError{"[Chunk Base]->Level->xPos", nbt.TagInt, xPos.TagID()}
	}

	x := int32(xPos.Data().(nbt.Int))

	zPos := lCmp.Get("zPos")
	if zPos.TagID() == 0 {
		return 0, 0, MissingTagError{"[Chunk Base]->Level->zPos"}
	} else if zPos.TagID() != nbt.TagInt {
		return 0, 0, WrongTypeError{"[Chunk Base]->Level->zPos", nbt.TagInt, zPos.TagID()}
	}

	z := int32(zPos.Data().(nbt.Int))

	return x, z, nil
}

func init() {
	filename = regexp.MustCompile(`^r.(-?[0-9]+).(-?[0-9]+).mca$`)
}
