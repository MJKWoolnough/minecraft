package minecraft

import (
	"io"
	"io/ioutil"
	"path"
)

var filename *regexp.Regexp

type readSeekCloser interface {
	io.ReadSeeker
	io.Closer
}

type writeSeekCloser interface {
	io.WriteCloser
	io.Closer
}

// Path allows different types of data streams (directory, compressed file,
// net conn, in memory buffer etc.) to be read from.
type Path interface {
	readRegion(x, z int32) readSeekCloser
	writeRegion(x, z int32) writeSeekCloser
	readLevelDat() (io.ReadCloser, error)
	writeLevelDat() (io.WriteCloser, error)
	GetRegions() [][2]int32
	Lock()
}

type NoLock struct{}

func (n NoLock) Error() string {
	return "lost lock on files"
}

type path struct {
	dirname string
	lock bool
}

// NewPath constructs a new directory based path to read from.
func NewPath(dirname string) (Path, error) {
	dirname = path.Clean(dirname)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		return nil, err
	}
	p := &path {
		dirname,
		false,
	}
	p.Lock()
	return p
}

func (p *path) readRegion(x, z int32) (readSeekCloser, error) {
	if !p.lock {
		return nil, &NoLock{}
	}
	f, err := os.Open(path.Join(p.dirname, "region", fmt.Sprintf("r.%d.%d.mca", x, z)))
	if os.IsNotExist(err) {
		return nil, nil
	}
	return f, err
}

func (p *path) writeRegion(x, z int32) (writeSeekCloser, error) {
	if !p.lock {
		return nil, &NoLock{}
	}
	if err := os.MkdirAll(path.Join(p.dirname, "region"), 0755); err != nil {
		return nil, err
	}
	return os.OpenFile(path.Join(p.dirname, "region", fmt.Sprintf("r.%d.%d.mca", x, z)), os.O_WRONLY | os.O_CREATE, 0666)
}

func (p *path) readLevelDat() (io.ReadCloser, error) {
	if !p.lock {
		return nil, &NoLock{}
	}
	f, err := os.Open(path.Join(p.dirname, "level.dat"))
	if os.IsNotExist(err) {
		return nil, nil
	}
	return f, err
}

func (p *path) writeLevelDat() (io.WriteCloser, error) {
	if !p.lock {
		return nil, &NoLock{}
	}
	os.OpenFile(path.Join(p.dirname, "level.dat"), os.O_WRONLY | os.O_CREATE, 0666)
}

// Update tracks the lock file for updates to remove the lock.
func (p *path) Update(filname string, mode uint8) {
	p.lock = false
	watcher.RemoveWatcher(p.dirname)
}

// GetRegions returns a list of region x,z coords of all generated regions.
func (p path) GetRegions() [][2]int32 {
	files := ioutil.ReadDir(path.Join(p.dirname, "region"))
	toRet := make([][2]int32, 0)
	var x, z int32
	for _, file := range files {
		if !file.IsDir() {
			if nums := filename.FindStringSubmatch(file.Name()); nums != nil {
				fmt.Sscan(nums[0], &x)
				fmt.Sscan(nums[1], &z)
				toRet = append(toRet, [2]{ x, z })
			}
		}
	}
	return toRet
}


// Lock will retake the lock file if it has been lost. May cause corruption.
func (p *path) Lock() {
	if p.lock {
		return
	}
	session := path.Join(p.dirname, "session.lock")
	if f, err := os.Create(session); err != nil {
		return nil, err
	} else {
		fmt.Fprintf(f, "%d", timestampMS())
		f.Close()
	}
	watcher.Watch(session, l)
}

func init() {
	filename = regexp.MustCompile(`^r.(-?[0-9]+).(-?[0-9]+).mca$`)
}