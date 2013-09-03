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

func NewLevel(location Path) (*Level, error) {
	var (
		levelDat nbt.Tag
		data     *nbt.Compound
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
				nbt.NewTag("LastPlayed", nbt.NewLong(timestampMS())),
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
	}
	if levelDat.TagId() != nbt.Tag_Compound {
		return nil, WrongTypeError{"[BASE]", nbt.Tag_Compound, levelDat.TagId()}
	} else if d := levelDat.Data().(*nbt.Compound).Get("Data"); d != nil {
		if d.TagId() == nbt.Tag_Compound {
			data = d.Data().(*nbt.Compound)
		} else {
			return nil, WrongTypeError{"Data", nbt.Tag_Compound, d.TagId()}
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
		LIGHT_NONE,
		false,
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
	var opacity, light uint8
	if l.lighting != LIGHT_NONE {
		opacity = c.GetOpacity(x, y, z)
		if l.lighting == LIGHT_ALL {
			light = c.GetBlockLight(x, y, z)
		}
	}
	c.SetBlock(x, y, z, block)
	if l.lighting != LIGHT_NONE {
		if block.Opacity() != opacity {

		}
		if block.Light() != light {

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
		l.changes.Set(pos, true)
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
		if l.changes.Get(n) {
			toSave = append(toSave, c.GetNBT())
			l.changes.Set(n, false)
		}
		delete(l.chunks, n)
	}
	if len(toSave) > 0 {
		return l.path.SetChunk(toSave...)
	}
	return nil
}

func (l *Level) Close() {
	l.changed = false
	for n := range l.chunks {
		delete(l.chunks, n)
	}
	l.changes = boolmap.NewMap()
}

func CoordsToRegion(x, z int32) (int32, int32) {
	return x >> 9, z >> 9
}

func timestampMS() int64 {
	return time.Now().Unix() * 1000
}
