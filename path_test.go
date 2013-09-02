package minecraft

import (
	"encoding/binary"
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func testPathChunkSetGet(t *testing.T, path Path) {
	toPlace := [][]nbt.Tag{
		[]nbt.Tag{
			addPos(0, 0, 3),  //0[4]
			addPos(1, 0, 2),  //0[4],1[3]
			addPos(2, 0, 1),  //0[4],1[3],2[2]
			addPos(3, 0, 0),  //0[4],1[3],2[2],3[1]
			addPos(20, 0, 2), //0[4],1[3],2[2],3[1],20[3]
			addPos(0, 20, 1), //0[4],1[3],2[2],3[1],20[3],640[2]
			addPos(-1, 0, 0),
			addPos(0, -1, 0),
			addPos(-1, -1, 1),
			addPos(-3, -3, 1),
		},
		[]nbt.Tag{
			addPos(0, 0, 1), //0[2],[2],1[3],2[2],3[1],20[3],640[2]
			addPos(1, 0, 2), //0[2],[2],1[3],2[2],3[1],20[3],640[2]
			addPos(3, 0, 1), //0[2],3[2],1[3],2[2],[1],20[3],640[2]
			addPos(4, 0, 0), //0[2],3[2],1[3],2[2],4[1],(!)20[3],(!)640[2] | (2 + 0) << 8 | 2, (2 + 2) << 8 | 2,(2 + 4) << 8 | 3,(2 + 7) << 8 | 2, (2 + 9) << 8 | 1
			addPos(-1, 0, 1),
			addPos(0, -2, 1),
			addPos(-1, -1, 0),
			addPos(-3, -3, 0),
		},
	}
	retest := []bool{
		false,
		false,
		true,
		false,
		true,
		true,
		false,
		true,
		false,
		false,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
	}
	for num, chunkList := range toPlace {
		if err := path.SetChunk(chunkList...); err != nil {
			t.Fatal(err.Error())
		}
		for i, chunk := range chunkList {
			if x, z, err := chunkCoords(chunk); err != nil {
				t.Fatal(err.Error())
			} else if thatChunk, err := path.GetChunk(x, z); err != nil {
				t.Fatal(err.Error())
			} else if thatChunk == nil {
				t.Fatalf("testPathChunkSetGet: 0-%d-%d: no chunk returned", num, i)
			} else if !thatChunk.Equal(chunk) {
				t.Fatalf("testPathChunkSetGet: 0-%d-%d: returned chunk not equal to set chunk, expecting: -\n%s\ngot: -\n%s", num, i, chunk.String(), thatChunk.String())
			}
		}
	}
	for num, chunkList := range toPlace {
		for i, chunk := range chunkList {
			if x, z, err := chunkCoords(chunk); err != nil {
				t.Fatal(err.Error())
			} else if thatChunk, err := path.GetChunk(x, z); err != nil {
				t.Fatal(err.Error())
			} else if thatChunk == nil {
				t.Fatalf("testPathChunkSetGet: 1-%d-%d: no chunk returned", num, i)
			} else if thatChunk.Equal(chunk) != retest[0] {
				t.Errorf("testPathChunkSetGet: 1-%d-%d: returned chunk not equal to set chunk, expecting: -\n%s\ngot: -\n%s", num, i, chunk.String(), thatChunk.String())
			}
			retest = retest[1:]
		}
	}
}

func testPathLevelSetGet(t *testing.T, path Path) {
	levelDat := nbt.NewTag("", nbt.NewCompound([]nbt.Tag{
		nbt.NewTag("Beep", nbt.NewCompound([]nbt.Tag{
			nbt.NewTag("SomeInt", nbt.NewInt(45)),
			nbt.NewTag("SomeString", nbt.NewString("hello")),
		})),
	}))
	if err := path.WriteLevelDat(levelDat); err != nil {
		t.Error(err.Error())
	} else if newLevelDat, err := path.ReadLevelDat(); err != nil {
		t.Error(err.Error())
	} else if !newLevelDat.Equal(levelDat) {
		t.Errorf("level data doesn't match original, expecting: -\n%s\ngot: -\n%s", levelDat.String(), newLevelDat.String())
	}
}

func testPathChunkRemove(t *testing.T, path Path) {
	toRemove := [][2]int32{
		{0, 20},
		{20, 0},
		{-3, -3},
	}
	for num, tR := range toRemove {
		if err := path.RemoveChunk(tR[0], tR[1]); err != nil {
			t.Error(err.Error())
		} else if tC, err := path.GetChunk(tR[0], tR[1]); err != nil {
			t.Error(err.Error())
		} else if tC != nil {
			t.Errorf("testPathChunkRemove %d: failed to remove chunk at %d,%d", num, tR[0], tR[1])
		}
	}
}

func testPathRegionsGet(t *testing.T, path Path) {
	regions := path.GetRegions()
	should := [][2]int32{
		{-1, -1},
		{0, -1},
		{-1, 0},
		{0, 0},
	}
	if len(regions) != len(should) {
		t.Error("returned regions slice does not match expected")
	} else {
	CL:
		for i := 0; i < len(regions); i++ {
			for j := 0; j < len(should); j++ {
				if regions[i][0] == should[j][0] && regions[i][1] == should[j][1] {
					should = append(should[:j], should[j+1:]...)
					continue CL
				}
			}
			t.Error("returned regions slice does not match expected")
			break
		}
	}
}

func addPos(x, z int32, chunkNum uint8) nbt.Tag {
	e := chunks[chunkNum].data
	e.Set(nbt.NewTag("xPos", nbt.NewInt(x)))
	e.Set(nbt.NewTag("zPos", nbt.NewInt(z)))
	return chunks[chunkNum].GetNBT()
}

func TestMemPath(t *testing.T) {
	f := NewMemPath()
	if a := len(f.GetRegions()); a != 0 {
		t.Errorf("should start with zero regions, have %d", a)
		return
	}
	testPathChunkSetGet(t, f)
	testPathLevelSetGet(t, f)
	testPathRegionsGet(t, f)
	testPathChunkRemove(t, f)
}

func TestFilePath(t *testing.T) {
	var (
		tempDir string
		err     error
		f       Path
	)
	if tempDir, err = ioutil.TempDir("", "minecraft-path-test"); err != nil {
		t.Error(err.Error())
	} else if f, err = NewFilePath(tempDir); err != nil {
		t.Error(err.Error())
	}
	if a := len(f.GetRegions()); a != 0 {
		t.Errorf("should start with zero regions, have %d", a)
		return
	}
	testPathChunkSetGet(t, f)
	testPathLevelSetGet(t, f)
	testPathRegionsGet(t, f)
	testPathChunkRemove(t, f)

	//Check Files
	file, err := os.Open(path.Join(tempDir, "region", "r.0.0.mca"))
	if err != nil {
		t.Error(err.Error())
	}
	var positions, should [1024]uint32
	if err = binary.Read(file, binary.BigEndian, positions[:]); err != nil {
		t.Error(err.Error())
	}
	file.Close()
	should[0] = (2+0)<<8 | 2 //pos 0 + offset(2), size 2
	should[1] = (2+4)<<8 | 3 //pos 4 + offset(2), size 3
	should[2] = (2+7)<<8 | 2 //pos 7 + offset(2), size 2
	should[3] = (2+2)<<8 | 2 //pos 2 + offset(2), size 2
	should[4] = (2+9)<<8 | 1 //pos 9 + offset(2), size 1

	for i := 0; i < 1024; i++ {
		if should[i] != positions[i] {
			t.Errorf("chunk position/size incorrect, expecting chunk %d at %d, got %d", i, should[i], positions[i])
			break
		}
	}
	regions := f.GetRegions()

	for _, region := range regions {
		if fi, err := os.Stat(path.Join(tempDir, "region", fmt.Sprintf("r.%d.%d.mca", region[0], region[1]))); err != nil {
			t.Error(err.Error())
		} else if s := fi.Size(); s%4096 != 0 {
			t.Errorf("regions %d,%d filesize not divisible by 4096, got %d", region[0], region[1], s)
		}
	}

	if err = os.RemoveAll(tempDir); err != nil {
		t.Error(err.Error())
	}
}

var chunks [4]chunk

func init() {
	a, _ := newChunk(0, 0, nil)
	b, _ := newChunk(0, 0, nil)
	c, _ := newChunk(0, 0, nil)
	d, _ := newChunk(0, 0, nil)
	chunks = [4]chunk{
		*a,
		*b,
		*c,
		*d,
	}
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			for k := 0; k < 256; k++ {
				if k < 8 {
					var tick []Tick
					if k%2 == 0 {
						tick = []Tick{{int32(i+j+k) % 4096, 1, -1}}
					}
					chunks[3].SetBlock(int32(i), int32(k), int32(j), &Block{
						uint16(i+j+k) % 4096,
						uint8(i),
						[]nbt.Tag{
							nbt.NewTag("testMD", nbt.NewInt(int32(i*j*k))),
						},
						tick,
					})
				}
				if k < 250 {
					chunks[2].SetBlock(int32(i), int32(k), int32(j), &Block{1, 0, nil, nil})
				} else {
					chunks[2].SetBlock(int32(i), int32(k), int32(j), &Block{
						1,
						0,
						[]nbt.Tag{
							nbt.NewTag("testMD", nbt.NewInt(int32(i*j*k))),
						},
						nil,
					})
				}
			}
			chunks[1].SetBlock(int32(i), int32(j)*16, int32(j), &Block{
				uint16(i*j + i + j),
				uint8(i),
				[]nbt.Tag{
					nbt.NewTag("testMD1", nbt.NewInt(int32(i))),
					nbt.NewTag("testMD2", nbt.NewInt(int32(i+1))),
					nbt.NewTag("testMD3", nbt.NewInt(int32(i+2))),
					nbt.NewTag("testMD4", nbt.NewInt(int32(i+3))),
					nbt.NewTag("testMD5", nbt.NewInt(int32(i+4))),
				},
				[]Tick{{int32(i*j+i+j) % 4096, 1, -1}},
			})
		}
		chunks[0].SetBlock(int32(i), int32(i), int32(i), &Block{uint16(i), uint8(i), nil, nil})
	}
}
