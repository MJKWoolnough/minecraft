// Copyright (c) 2013 - Michael Woolnough <michael.woolnough@gmail.com>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package minecraft

import (
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"github.com/MJKWoolnough/io-watcher"
	"github.com/MJKWoolnough/memio"
	"github.com/MJKWoolnough/minecraft/nbt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"time"
)

var filename *regexp.Regexp

type Path interface {
	// Returns a nil nbt.Tag when chunk does not exists
	GetChunk(int32, int32) (nbt.Tag, error)
	SetChunk(...nbt.Tag) error
	RemoveChunk(int32, int32) error
	ReadLevelDat() (nbt.Tag, error)
	WriteLevelDat(nbt.Tag) error
	GetRegions() [][2]int32
	// 	GetChunks(int32, int32) [][2]int32
}

const (
	GZip byte = 1
	Zlib byte = 2
)

type FilePath struct {
	dirname string
	lock    bool
}

// NewFilePath constructs a new directory based path to read from.
func NewFilePath(dirname string) (*FilePath, error) {
	dirname = path.Clean(dirname)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		return nil, err
	}
	p := &FilePath{
		dirname,
		false,
	}
	p.Lock()
	return p, nil
}

func (p *FilePath) GetChunk(x, z int32) (nbt.Tag, error) {
	if !p.lock {
		return nil, &NoLock{}
	}
	f, err := os.Open(path.Join(p.dirname, "region", fmt.Sprintf("r.%d.%d.mca", x>>5, z>>5)))
	if os.IsNotExist(err) {
		return nil, nil
	}
	defer f.Close()
	pos := int64((z&31)<<5 | (x & 31))
	if _, err = f.Seek(4*pos, os.SEEK_SET); err != nil {
		return nil, err
	}
	var locationSize uint32
	if err = binary.Read(f, binary.BigEndian, &locationSize); err != nil {
		return nil, err
	} else if locationSize>>8 == 0 {
		return nil, nil
	} else if _, err = f.Seek(int64(locationSize>>8<<12), os.SEEK_SET); err != nil {
		return nil, err
	}
	reader := io.LimitReader(f, int64(locationSize&255<<12))
	var (
		length      uint32
		compression byte
	)
	if err = binary.Read(reader, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	reader = io.LimitReader(reader, int64(length))
	if err = binary.Read(reader, binary.BigEndian, &compression); err != nil {
		return nil, err
	}
	switch compression {
	case GZip:
		gReader, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		defer gReader.Close()
		reader = gReader
	case Zlib:
		if reader, err = zlib.NewReader(reader); err != nil {
			return nil, err
		}
	default:
		return nil, &UnknownCompression{compression}
	}
	data, _, err := nbt.ReadNBTFrom(reader)
	return data, err
}

type rc struct {
	pos int32
	buf []byte
}

func (p *FilePath) SetChunk(data ...nbt.Tag) error {
	if !p.lock {
		return &NoLock{}
	}
	regions := make(map[uint64][]rc)
	poses := make([]uint64, 0)
	for _, d := range data {
		x, z, err := chunkCoords(d)
		if err != nil {
			return err
		}
		pos := uint64(z)<<32 | uint64(uint32(x))
		for _, p := range poses {
			if p == pos {
				return &ConflictError{x, z}
			}
		}
		poses = append(poses, pos)
		r := uint64(z>>5)<<32 | uint64(uint32(x>>5))
		reg := rc{pos: (z&31)<<5 | (x & 31)}
		zl := zlib.NewWriter(memio.Create(&reg.buf))
		if _, err := d.WriteTo(zl); err != nil {
			return err
		}
		zl.Close()
		if regions[r] == nil {
			regions[r] = []rc{reg}
		} else {
			regions[r] = append(regions[r], reg)
		}
	}
	for rId, chunks := range regions {
		if err := p.setChunks(int32(rId&0xffffffff), int32(rId>>32), chunks); err != nil {
			return err
		}
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
	if err := os.MkdirAll(path.Join(p.dirname, "region"), 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(path.Join(p.dirname, "region", fmt.Sprintf("r.%d.%d.mca", x, z)), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	var positions [1024]uint32
	if err = binary.Read(f, binary.BigEndian, positions[:]); err != nil && err != io.EOF {
		return err
	}
	var todoChunks []rc
	for _, chunk := range chunks {
		newSize := uint32(len(chunk.buf)+5) >> 12
		if uint32(len(chunk.buf))&4095 > 0 {
			newSize++
		}
		if positions[chunk.pos]&255 == newSize {
			if _, err = f.Seek(4*int64(chunk.pos)+4096, os.SEEK_SET); err != nil { // Write the time, then the data
				return err
			} else if err = binary.Write(f, binary.BigEndian, uint32(time.Now().Unix())); err != nil {
				return err
			} else if _, err = f.Seek(int64(positions[chunk.pos])>>8<<12, os.SEEK_SET); err != nil {
				return err
			} else if err = binary.Write(f, binary.BigEndian, uint32(len(chunk.buf))+1); err != nil {
				return err
			} else if err = binary.Write(f, binary.BigEndian, Zlib); err != nil {
				return err
			} else if err = binary.Write(f, binary.BigEndian, chunk.buf); err != nil {
				return err
			}
		} else {
			todoChunks = append(todoChunks, chunk)
			positions[chunk.pos] = 0
		}
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
		if _, err = f.Seek(4*int64(chunk.pos), os.SEEK_SET); err != nil {
			return err
		} else if err = binary.Write(f, binary.BigEndian, positions[0]); err != nil {
			return err
		} else if _, err = f.Seek(4*(int64(chunk.pos)+1024), os.SEEK_SET); err != nil { // Write the time, then the data
			return err
		} else if err = binary.Write(f, binary.BigEndian, uint32(time.Now().Unix())); err != nil {
			return err
		} else if _, err = f.Seek(int64(newPosition)<<12, os.SEEK_SET); err != nil {
			return err
		} else if err = binary.Write(f, binary.BigEndian, uint32(len(chunk.buf))+1); err != nil {
			return err
		} else if err = binary.Write(f, binary.BigEndian, Zlib); err != nil {
			return err
		} else if err = binary.Write(f, binary.BigEndian, chunk.buf); err != nil {
			return err
		} else if writeLastByte { // Make filesize mod 4096 (for minecraft compatibility)
			if _, err = f.Seek(int64(newPosition+newSize)<<12-1, os.SEEK_SET); err != nil {
				return err
			} else if err = binary.Write(f, binary.BigEndian, byte(0)); err != nil {
				return err
			}

		}
	}
	return nil
}

func (p *FilePath) RemoveChunk(x, z int32) error {
	chunkX := x & 31
	regionX := x >> 5
	chunkZ := z & 31
	regionZ := z >> 5
	f, err := os.OpenFile(path.Join(p.dirname, "region", fmt.Sprintf("r.%d.%d.mca", regionX, regionZ)), os.O_WRONLY, 0666)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Seek(int64(chunkZ<<5|chunkX)*4, os.SEEK_SET); err != nil {
		return err
	}
	return binary.Write(f, binary.BigEndian, int32(0))
}

func (p *FilePath) ReadLevelDat() (nbt.Tag, error) {
	if !p.lock {
		return nil, &NoLock{}
	}
	f, err := os.Open(path.Join(p.dirname, "level.dat"))
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	defer f.Close()
	g, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	data, _, err := nbt.ReadNBTFrom(g)
	return data, err
}

func (p *FilePath) WriteLevelDat(data nbt.Tag) error {
	if !p.lock {
		return &NoLock{}
	}
	f, err := os.OpenFile(path.Join(p.dirname, "level.dat"), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil
	}
	defer f.Close()
	g := gzip.NewWriter(f)
	defer g.Close()
	_, err = data.WriteTo(g)
	return err
}

// GetRegions returns a list of region x,z coords of all generated regions.
func (p *FilePath) GetRegions() [][2]int32 {
	files, _ := ioutil.ReadDir(path.Join(p.dirname, "region"))
	toRet := make([][2]int32, 0)
	var x, z int32
	for _, file := range files {
		if !file.IsDir() {
			if nums := filename.FindStringSubmatch(file.Name()); nums != nil {
				fmt.Sscan(nums[1], &x)
				fmt.Sscan(nums[2], &z)
				toRet = append(toRet, [2]int32{x, z})
			}
		}
	}
	return toRet
}

// Update tracks the lock file for updates to remove the lock.
func (p *FilePath) Update(filname string, mode uint8) {
	p.lock = false
	watcher.StopWatch(p.dirname)
}

// Lock will retake the lock file if it has been lost. May cause corruption.
func (p *FilePath) Lock() {
	if p.lock {
		return
	}
	session := path.Join(p.dirname, "session.lock")
	if f, err := os.Create(session); err != nil {
		return //??
	} else {
		fmt.Fprintf(f, "%d", timestampMS())
		f.Close()
	}
	watcher.Watch(session, p)
	p.lock = true
}

func (p *FilePath) Defrag(x, z int32) error {
	return nil
}

type memPath struct {
	level  []byte
	chunks map[uint64][]byte
}

func NewMemPath() *memPath {
	return &memPath{chunks: make(map[uint64][]byte)}
}

func (m *memPath) GetChunk(x, z int32) (nbt.Tag, error) {
	pos := uint64(z)<<32 | uint64(uint32(x))
	if m.chunks[pos] == nil {
		return nil, nil
	}
	return m.read(m.chunks[pos])
}

func (m *memPath) SetChunk(data ...nbt.Tag) error {
	for _, d := range data {
		x, z, err := chunkCoords(d)
		if err != nil {
			return err
		}
		var buf []byte
		if err = m.write(d, &buf); err != nil {
			return err
		}
		pos := uint64(z)<<32 | uint64(uint32(x))
		m.chunks[pos] = buf
	}
	return nil
}

func (m *memPath) RemoveChunk(x, z int32) error {
	pos := uint64(z)<<32 | uint64(uint32(x))
	delete(m.chunks, pos)
	return nil
}

func (m *memPath) ReadLevelDat() (nbt.Tag, error) {
	if len(m.level) == 0 {
		return nil, nil
	}
	return m.read(m.level)
}

func (m *memPath) WriteLevelDat(data nbt.Tag) error {
	return m.write(data, &m.level)
}

func (m *memPath) GetRegions() [][2]int32 {
	toRet := make([][2]int32, 0)
JP:
	for i := range m.chunks {
		x := int32(i&0xffffffff) >> 5
		z := int32(i>>32) >> 5
		for _, j := range toRet {
			if j[0] == x && j[1] == z {
				continue JP
			}
		}
		toRet = append(toRet, [2]int32{x, z})
	}
	return toRet
}

func (m *memPath) read(buf []byte) (nbt.Tag, error) {
	z, err := zlib.NewReader(memio.Open(buf))
	if err != nil {
		return nil, err
	}
	data, _, err := nbt.ReadNBTFrom(z)
	return data, err
}

func (m *memPath) write(data nbt.Tag, buf *[]byte) error {
	z := zlib.NewWriter(memio.Create(buf))
	defer z.Close()
	_, err := data.WriteTo(z)
	return err
}

func chunkCoords(data nbt.Tag) (x int32, z int32, err error) {
	if data.TagId() != nbt.Tag_Compound {
		err = &WrongTypeError{"[Chunk Base]", nbt.Tag_Compound, data.TagId()}
		return
	} else if lTag := data.Data().(*nbt.Compound).Get("Level"); lTag == nil {
		err = &MissingTagError{"[Chunk Base]->Level"}
		return
	} else if lTag.TagId() != nbt.Tag_Compound {
		err = &WrongTypeError{"[Chunk Base]->Level", nbt.Tag_Compound, lTag.TagId()}
		return
	} else {
		lCmp := lTag.Data().(*nbt.Compound)
		if xPos := lCmp.Get("xPos"); xPos == nil {
			err = &MissingTagError{"[Chunk Base]->Level->xPos"}
			return
		} else if xPos.TagId() != nbt.Tag_Int {
			err = &WrongTypeError{"[Chunk Base]->Level->xPos", nbt.Tag_Int, xPos.TagId()}
			return
		} else {
			x = int32(*xPos.Data().(*nbt.Int))
		}
		if zPos := lCmp.Get("zPos"); zPos == nil {
			err = &MissingTagError{"[Chunk Base]->Level->zPos"}
			return
		} else if zPos.TagId() != nbt.Tag_Int {
			err = &WrongTypeError{"[Chunk Base]->Level->zPos", nbt.Tag_Int, zPos.TagId()}
			return
		} else {
			z = int32(*zPos.Data().(*nbt.Int))
		}
	}
	return
}

func init() {
	filename = regexp.MustCompile(`^r.(-?[0-9]+).(-?[0-9]+).mca$`)
}
