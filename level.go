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
	"fmt"
	"github.com/MJKWoolnough/io-watcher"
	"github.com/MJKWoolnough/minecraft/nbt"
	"math/rand"
	"os"
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

type MissingTagError struct {
	tagName string
}

func (m MissingTagError) Error() {
	return fmt.Sprintf("minecraft - level: missing %q tag", m.tagName)
}

type WrongTypeError struct {
	tagName        string
	expecting, got nbt.TagId
}

func (m MissingTagError) Error() {
	return fmt.Sprintf("minecraft - level: tag %q is of incorrect type, expecting %q, got %q", m.tagName)
}

type Level struct {
	location  Path
	regions   map[int32]map[int32]Region
	levelData nbt.Tag
	changed   bool
}

func NewLevel(location Path) (*Level, error) {
	var (
		t    nbt.Tag
		data *nbt.Compound
	)
	levelDat, err = location.readLevelDat()
	if err != nil {
		return nil, err
	} else if levelDat == nil {
		t = newLevelDat()
	} else {
		defer levelDat.Close()
		file, err := gzip.NewReader(levelDat)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		t, _, err = nbt.ReadNBTFrom(file)
		if err != nil {
			return nil, err
		}
	}
	if d := t.Get("Data"); d != nil {
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
		make(map[int32]map[int32]Region),
		t,
		false,
	}, nil
}

func (l Level) GetSpawn() (x, y, z int32) {
	if l.levelData == nil {
		return
	}
	data := l.levelData.Get("Data").Data().(*nbt.Compound)
	xTag, yTag, zTag := data.Get("SpawnX"), data.Get("SpawnY"), data.Get("SpawnZ")
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
	data := l.levelData.Get("Data").Data().(*nbt.Compound)
	data.Set("SpawnX", nbt.NewInt(x))
	data.Set("SpawnY", nbt.NewInt(y))
	data.Set("SpawnZ", nbt.NewInt(z))
	l.changed = true
}

func (l *Level) GetBlock(x, y, z int32) (*Block, error) {
	r, err := l.GetRegion(CoordsToRegion(x, z))
	if err != nil {
		return nil, err
	}
	if r == nil {
		return BlockAir
	}
	return r.GetBlock(x, y, z)
}

func (l *Level) SetBlock(x, y, z int32, block *Block) error {
	r, err := l.GetRegion(CoordsToRegion(x, z))
	if err != nil {
		return err
	}
	if r != nil {
		r.SetBlock(x, y, z, block)
	}
}

func (l *Level) GetBiome(x, z int32) (Biome, error) {
	r, err := l.GetRegion(CoordsToRegion(x, z))
	if r == nil {
		return Biome_Plains, err
	}
	r.GetBiome(x, z)
}

func (l *Level) SetBiome(x, z int32, biome Biome) error {
	r, err := l.GetRegion(CoordsToRegion(x, z))
	if err != nil {
		return err
	}
	if r != nil {
		r.SetBiome(x, z, biome)
	}
}

func (l Level) GetName() string {
	s := l.levelData.Get("Data").Data().(*nbt.Compound).Get("LevelName").(*nbt.String)
	return string(*s)
}

func (l *Level) SetName(name string) {
	l.levelData.Get("Data").Data().(*nbt.Compound).Set("LevelName", nbt.NewString(name))
	l.changed = true
}

func (l *Level) GetRegion(x, z) (*Region, error) {
	if ra, ok := l.regions[x]; ok {
		if r, ok := ra[z]; ok {
			return r, nil
		}
	}
	if _, ok := l.regions[x]; !ok {
		l.regions[x] = make(map[int32]*Region)
	}
	r, err := newRegion(x, z)
	l.regions[x][z] = r
	return r, err
}

func (l *Level) Save() error {
	if !l.haveLock {
		return &NoLock{}
	}
	if l.changed {
		l, err = l.path.writeLevelDat()
		if err != nil {
			return err
		}
		defer l.Close()
		file := gzip.NewWriter(file)
		defer file.Close()
		if _, err = l.levelData.WriteTo(file); err != nil {
			return err
		}
		l.change = false
	}
	// 	for _, x := range l.regions {
	// 		for _, z := range x {
	//
	// 		}
	// 	}
	return nil
}

func (l *Level) Close() {
	l.change = false
	// 	for _, x := range l.regions {
	// 		for _, z := range x {
	//
	// 		}
	// 	}
}

func newLevelDat() nbt.Tag {
	return nbt.NewTag("", nbt.NewCompound([]Tag{
		nbt.NewTag("GameType", nbt.NewInt(1)),
		nbt.NewTag("generatorName", nbt.NewString("flat")),
		nbt.NewTag("generatorVersion", nbt.NewInt(0)),
		nbt.NewTag("generatorOptions", nbt.NewString("0")),
		nbt.NewTag("hardcore", nbt.NewByte(0)),
		nbt.NewTag("LastPlayed", nbt.NewLong(timestampMS())),
		nbt.NewTag("LevelName", nbt.NewString("")),
		nbt.NewTag("MapFeatures", nbt.NewByte(0)),
		nbt.NewTag("RandomSeed", nbt.NewLong(rand.New(rand.NewSource(time.Unix())).Int63())),
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
	}))
}

func CoordsToRegion(x, z int32) (int32, int32) {
	return x >> 9, z >> 9
}

func timestampMS() int64 {
	return time.Now().Unix() * 1000
}
