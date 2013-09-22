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
	"github.com/MJKWoolnough/boolmap"
	"github.com/MJKWoolnough/minecraft/nbt"
	"math/rand"
	"time"
)

var (
	levelRequired = map[string]nbt.TagId{
		"LevelName": nbt.Tag_String,
		"SpawnX":    nbt.Tag_Int,
		"SpawnY":    nbt.Tag_Int,
		"SpawnZ":    nbt.Tag_Int,
	}
)

const (
	LIGHT_NONE uint8 = iota
	LIGHT_SIMPLE
	LIGHT_ALL
)

type Level struct {
	path      Path
	chunks    map[uint64]*chunk
	changes   boolmap.Map
	levelData *nbt.Compound
	lighting  uint8
	changed   bool
}

func NewLevel(location Path, ll uint8) (*Level, error) {
	var (
		levelDat nbt.Tag
		data     *nbt.Compound
		changed  bool
	)
	levelDat, err := location.ReadLevelDat()
	if err != nil {
		return nil, err
	} else if levelDat == nil {
		levelDat = nbt.NewTag("", nbt.NewCompound([]nbt.Tag{
			nbt.NewTag("Data", nbt.NewCompound([]nbt.Tag{
				nbt.NewTag("GameType", nbt.NewInt(1)),
				nbt.NewTag("generatorName", nbt.NewString("flat")),
				nbt.NewTag("generatorVersion", nbt.NewInt(0)),
				nbt.NewTag("generatorOptions", nbt.NewString("0")),
				nbt.NewTag("hardcore", nbt.NewByte(0)),
				nbt.NewTag("LastPlayed", nbt.NewLong(time.Now().Unix()*1000)),
				nbt.NewTag("LevelName", nbt.NewString("")),
				nbt.NewTag("MapFeatures", nbt.NewByte(0)),
				nbt.NewTag("RandomSeed", nbt.NewLong(rand.New(rand.NewSource(time.Now().Unix())).Int63())),
				nbt.NewTag("raining", nbt.NewByte(0)),
				nbt.NewTag("rainTime", nbt.NewInt(0)),
				nbt.NewTag("SizeOnDisk", nbt.NewLong(0)),
				nbt.NewTag("SpawnX", nbt.NewInt(0)),
				nbt.NewTag("SpawnY", nbt.NewInt(0)),
				nbt.NewTag("SpawnZ", nbt.NewInt(0)),
				nbt.NewTag("Time", nbt.NewLong(0)),
				nbt.NewTag("thundering", nbt.NewByte(0)),
				nbt.NewTag("thunderTime", nbt.NewInt(0)),
				nbt.NewTag("version", nbt.NewInt(19133)),
			})),
		}))
		changed = true
	}
	if levelDat.TagId() != nbt.Tag_Compound {
		return nil, WrongTypeError{"[BASE]", nbt.Tag_Compound, levelDat.TagId()}
	} else if d := levelDat.Data().(*nbt.Compound).Get("Data"); d != nil {
		if d.TagId() == nbt.Tag_Compound {
			data = d.Data().(*nbt.Compound)
		} else {
			return nil, &WrongTypeError{"Data", nbt.Tag_Compound, d.TagId()}
		}
	} else {
		return nil, &MissingTagError{"Data"}
	}
	for name, tagType := range levelRequired {
		if x := data.Get(name); x == nil {
			return nil, &MissingTagError{name}
		} else if x.TagId() != tagType {
			return nil, &WrongTypeError{name, tagType, x.TagId()}
		}
	}
	return &Level{
		location,
		make(map[uint64]*chunk),
		boolmap.NewMap(),
		levelDat.Data().(*nbt.Compound).Get("Data").Data().(*nbt.Compound),
		ll,
		changed,
	}, nil
}

func (l Level) GetSpawn() (x, y, z int32) {
	if l.levelData == nil {
		return
	}
	xTag, yTag, zTag := l.levelData.Get("SpawnX"), l.levelData.Get("SpawnY"), l.levelData.Get("SpawnZ")
	if xd, ok := xTag.Data().(*nbt.Int); !ok {
		return
	} else {
		x = int32(*xd)
	}
	if yd, ok := yTag.Data().(*nbt.Int); !ok {
		return
	} else {
		y = int32(*yd)
	}
	if zd, ok := zTag.Data().(*nbt.Int); ok {
		z = int32(*zd)
	}
	return
}

func (l *Level) SetSpawn(x, y, z int32) {
	l.levelData.Set(nbt.NewTag("SpawnX", nbt.NewInt(x)))
	l.levelData.Set(nbt.NewTag("SpawnY", nbt.NewInt(y)))
	l.levelData.Set(nbt.NewTag("SpawnZ", nbt.NewInt(z)))
	l.changed = true
}

func (l *Level) GetBlock(x, y, z int32) (*Block, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return nil, err
	} else if c == nil {
		return &Block{}, nil
	}
	return c.GetBlock(x, y, z), nil
}

func (l *Level) SetBlock(x, y, z int32, block *Block) error {
	c, err := l.getChunk(x, z, true)
	if err != nil {
		return err
	}
	if ys := y >> 4; c.sections[ys] == nil { //crossing the object boundary for lighting
		if block.Equal(&Block{}) {
			return nil
		}
		c.sections[ys] = newSection(y)
		if l.lighting != LIGHT_NONE {
			baseX := x >> 4 << 4
			baseY := y >> 4 << 4
			baseZ := z >> 4 << 4
			for i := baseX; i < baseX+16; i++ {
				for k := baseZ; k < baseZ+16; k++ {
					j := baseY + 15
					if h := c.GetHeight(i, k); h < baseY {
						for ; j >= baseY; j-- {
							c.SetSkyLight(i, j, k, 15)
						}
					} else {
						switch l.lighting {
						case LIGHT_SIMPLE:
							for currLightLevel := c.GetSkyLight(i, j+1, k); j >= baseY; j-- {
								if currLightLevel > 0 {
									currLightLevel--
								}
								c.SetSkyLight(i, j, k, currLightLevel)
							}
						case LIGHT_ALL:
						}
					}
				}
			}
		}
	}
	var opacity uint8
	if l.lighting != LIGHT_NONE {
		opacity = c.GetOpacity(x, y, z)
	}
	c.SetBlock(x, y, z, block)
	if l.lighting != LIGHT_NONE {
		if block.Opacity() != opacity {
			nY := y
			for h := c.GetHeight(x, z); nY >= h; nY-- {
				c.SetSkyLight(x, nY, z, 15)
			}
			switch l.lighting {
			case LIGHT_SIMPLE:
				for currLightLevel := c.GetSkyLight(x, nY+1, z); nY >= 0; nY-- {
					if currLightLevel > 0 {
						if o := c.GetOpacity(x, nY, z); o < currLightLevel {
							currLightLevel -= o
						} else {
							currLightLevel = 0
						}
					}
					if c.GetSkyLight(x, nY, z) == currLightLevel {
						break
					}
					c.SetSkyLight(x, nY, z, currLightLevel)
				}
			case LIGHT_ALL:

			}
		}
		if block.Light() != c.GetBlockLight(x, y, z) {
			switch l.lighting {
			case LIGHT_SIMPLE:
				c.SetBlockLight(x, y, z, block.Light())
			case LIGHT_ALL:

			}
		}
	}
	return nil
}

func (l *Level) GetBiome(x, z int32) (Biome, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return Biome_Auto, err
	} else if c == nil {
		return Biome_Plains, nil
	}
	return c.GetBiome(x, z), nil
}

func (l *Level) SetBiome(x, z int32, biome Biome) error {
	c, err := l.getChunk(x, z, true)
	if err != nil {
		return err
	}
	c.SetBiome(x, z, biome)
	return nil
}

func (l Level) GetName() string {
	s := l.levelData.Get("LevelName").Data().(*nbt.String)
	return string(*s)
}

func (l *Level) SetName(name string) {
	l.levelData.Set(nbt.NewTag("LevelName", nbt.NewString(name)))
	l.changed = true
}

func (l *Level) getChunk(x, z int32, create bool) (*chunk, error) {
	x >>= 4
	z >>= 4
	pos := uint64(z)<<32 | uint64(uint32(x))
	if l.chunks[pos] == nil {
		chunkData, err := l.path.GetChunk(x, z)
		if err != nil {
			return nil, err
		}
		if chunkData != nil {
			chunk, err := newChunk(x, z, chunkData)
			if err != nil {
				return nil, err
			}
			l.chunks[pos] = chunk
		} else if create {
			l.chunks[pos], _ = newChunk(x, z, nil)
		}
	}
	if create {
		l.changes.Set(uint(pos), true)
	}
	return l.chunks[pos], nil
}

func (l *Level) Save() error {
	if l.changed {
		if err := l.path.WriteLevelDat(nbt.NewTag("", nbt.NewCompound([]nbt.Tag{nbt.NewTag("Data", l.levelData)}))); err != nil {
			return err
		}
		l.changed = false
	}
	toSave := make([]nbt.Tag, 0)
	for n, c := range l.chunks {
		if l.changes.Get(uint(n)) {
			toSave = append(toSave, c.GetNBT())
		}
	}
	l.changes = boolmap.NewMap()
	if len(toSave) > 0 {
		return l.path.SetChunk(toSave...)
	}
	return nil
}

func (l *Level) Close() {
	l.changed = false
	l.chunks = make(map[uint64]*chunk)
	l.changes = boolmap.NewMap()
}
