package minecraft

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbtparser"
	"io"
	"time"
)

type Region interface {
	Get(int32, int32, int32) Block
	Set(int32, int32, int32, Block)
	GetSkyLight(int32, int32, int32) uint8
	SetSkyLight(int32, int32, int32, uint8)
	Opacity(int32, int32, int32) uint8
	Export(io.WriteSeeker) error
	HasChanged() bool
	Compress()
	// 	SkyUpdates(int32, int32) [256]bool
	// 	HighestBlock(int32, int32) uint8
}

type region [1024]Chunk

// func (r *region) HighestBlock(x, z int32) uint8 {
// 	reg := zx1024(x, z)
// 	if reg > 1023 || r[reg] == nil {
// 		return 0
// 	}
// 	return r[reg].HighestBlock(uint8(x&15), uint8(z&15))
// }
// 
// func (r *region) SkyUpdates(x, z int32) [256]bool {
// 	reg := zx1024(x, z)
// 	if reg > 1023 || r[reg] == nil {
// 		return [256]bool {}
// 	}
// 	return r[reg].SkyUpdates()
// }

func (r *region) Compress() {
	for i := 0; i < 1024; i++ {
		if r[i] != nil {
			r[i].Compress()
		}
	}
}

func (r *region) Get(x, y, z int32) Block {
	if r == nil {
		return BlockAir
	}
	reg := zx1024(x, z)
	if reg > 1023 || r[reg] == nil {
		return BlockAir
	}
	xu, yu, zu := uint8(x&15), uint8(y), uint8(z&15)
	block := r[reg].Get(xu, yu, zu)
	block.SetMetadata(r[reg].GetMetadata(xu, yu, zu))
	block.Tick(r[reg].GetTick(xu, yu, zu))
	return block
}

func (r *region) Set(x, y, z int32, block Block) {
	if block == nil || r == nil {
		return
	}
	reg := zx1024(x, z)
	if reg > 1023 {
		return
	}
	if r[reg] == nil {
		r[reg] = NewChunk(x>>4, z>>4)
	}
	xu, yu, zu := uint8(x&15), uint8(y), uint8(z&15)
	if block.HasMetadata() {
		r[reg].SetMetadata(xu, yu, zu, append(block.GetMetadata(), nbtparser.NewTagInt("x", x), nbtparser.NewTagInt("y", y), nbtparser.NewTagInt("z", z)))
	} else {
		r[reg].SetMetadata(xu, yu, zu, nil)
	}
	if block.ToTick() {
		r[reg].SetTick(xu, yu, zu, []nbtparser.NBTTag{
			nbtparser.NewTagInt("i", int32(block.BlockId())),
			nbtparser.NewTagInt("t", -1),
			nbtparser.NewTagInt("x", x),
			nbtparser.NewTagInt("y", y),
			nbtparser.NewTagInt("z", z),
		})
	} else {
		r[reg].SetTick(xu, yu, zu, nil)
	}
	r[reg].Set(xu, yu, zu, block)
}

func (r *region) GetSkyLight(x, y, z int32) uint8 {
	if r == nil {
		return 15
	}
	reg := zx1024(x, z)
	if reg > 1023 || r[reg] == nil {
		return 15
	}
	return r[reg].GetSkyLight(uint8(x&15), uint8(y), uint8(z&15))
}

func (r *region) SetSkyLight(x, y, z int32, skylight uint8) {
	if r == nil {
		return
	}
	reg := zx1024(x, z)
	if reg > 1023 || r[reg] == nil {
		return
	}
	r[reg].SetSkyLight(uint8(x&15), uint8(y), uint8(z&15), skylight)
}

func (r *region) Opacity(x, y, z int32) uint8 {
	if r == nil {
		return 0
	}
	reg := zx1024(x, z)
	if reg > 1023 || r[reg] == nil {
		return 0
	}
	return r[reg].Opacity(uint8(x&15), uint8(y), uint8(z&15))
}

func (r *region) Export(file io.WriteSeeker) error {
	if r == nil {
		return nil
	}
	totalBlocks := uint32(2)
	compression := uint8(2)
	var locations, timestamps [1024]uint32
	written := false
	for i := 0; i < 1024; i++ {
		if r[i] != nil && !r[i].IsEmpty() {
			startBlock := int64(totalBlocks) << 12
			timestamps[i] = uint32(time.Now().Unix())
			if _, err := file.Seek(startBlock+5, 0); err != nil {
				return err
			} else if _, err := r[i].WriteTo(file); err != nil {
				return err
			} else if position, err := file.Seek(0, 1); err != nil {
				return err
			} else {
				written = true
				length := uint32(position - startBlock)
				lengthBlocks := length >> 12
				if length&4095 != 0 {
					lengthBlocks++
					if _, err := file.Seek(int64(totalBlocks+lengthBlocks)<<12-1, 0); err != nil {
						return err
					}
					if err := binary.Write(file, binary.BigEndian, byte(0)); err != nil {
						return err
					}
				}
				if _, err := file.Seek(startBlock, 0); err != nil {
					return err
				}
				locations[i] = totalBlocks<<8 | (lengthBlocks & 255)
				totalBlocks += lengthBlocks
				length -= 4 //4 because length includes compression field
				if err := binary.Write(file, binary.BigEndian, &length); err != nil {
					return err
				}
				if err := binary.Write(file, binary.BigEndian, &compression); err != nil {
					return err
				}
			}
		}
	}
	if written {
		if _, err := file.Seek(0, 0); err != nil {
			return err
		} else if err := binary.Write(file, binary.BigEndian, locations); err != nil {
			return err
		} else if err := binary.Write(file, binary.BigEndian, timestamps); err != nil {
			return err
		} else if _, err := file.Seek(int64(totalBlocks)<<12, 0); err != nil {
			return err
		}
	}
	return nil

}

func (r *region) HasChanged() bool {
	for i := 0; i < 1024; i++ {
		if r[i] != nil {
			if r[i].HasChanged() {
				return true
			}
		}
	}
	return false
}

func zx1024(x, z int32) uint16 {
	z1 := (z >> 4) & 31
	x1 := (x >> 4) & 31
	return uint16((z1 << 5) | x1)
}

func LoadRegion(data io.ReadSeeker) (Region, error) {
	r := new(region)
	var locations, timestamps [1024]uint32
	if err := binary.Read(data, binary.BigEndian, &locations); err != nil {
		return nil, err
	}
	if err := binary.Read(data, binary.BigEndian, &timestamps); err != nil {
		return nil, err
	}
	for i := 0; i < 1024; i++ {
		if locations[i] > 1 {
			location := int64((locations[i] >> 8) << 12)
			if _, err := data.Seek(location, 0); err != nil {
				return nil, err
			}
			var (
				length      uint32
				compression uint8
			)
			if err := binary.Read(data, binary.BigEndian, &length); err != nil {
				return nil, err
			}
			if err := binary.Read(data, binary.BigEndian, &compression); err != nil {
				return nil, err
			}
			limited := io.LimitReader(data, int64(length-1))
			if compression == 1 {
				buf := new(bytes.Buffer)
				if zFile, err := gzip.NewReader(limited); err != nil {
					return nil, err
				} else {
					z := zlib.NewWriter(buf)
					if _, err := io.Copy(z, zFile); err != nil {
						return nil, err
					}
					z.Close()
					zFile.Close()
					limited = buf
				}
			} else if compression != 2 {
				return nil, fmt.Errorf("Unknown compression scheme: %d", compression)
			}
			if chunk, err := LoadChunk(limited); err != nil {
				return nil, err
			} else {
				r[i] = chunk
			}
		}
	}
	return r, nil
}

func NewRegion() (Region, error) {
	return LoadRegion(bytes.NewReader(make([]byte, 8192)))
}
