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

type light uint8

func (l light) SkyLight() bool {
	return l&3 > 0
}

func (l light) SkyLightSimple() bool {
	return l&3 == 1
}

func (l light) SkyLightAll() bool {
	return l&3 == 3
}

func (l light) BlockLight() bool {
	return l&12 > 0
}

func (l light) BlockLightSimple() bool {
	return l&12 == 4
}

func (l light) BlockLightAll() bool {
	return l&12 == 12
}

const (
	LIGHT_SKY_NONE     light = 0
	LIGHT_SKY_SIMPLE   light = 1
	LIGHT_SKY_ALL      light = 3
	LIGHT_BLOCK_NONE   light = 0
	LIGHT_BLOCK_SIMPLE light = 4
	LIGHT_BLOCK_ALL    light = 12
)

type Level struct {
	path      Path
	chunks    map[uint64]*chunk
	changes   boolmap.Map
	levelData *nbt.Compound
	lighting  light
	changed   bool
}

// Create/Load a minecraft level from the given path. Also takes a light value
// to determine how it should it should process lighting.
func NewLevel(location Path, ll light) (*Level, error) {
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
		return nil, &WrongTypeError{"[BASE]", nbt.Tag_Compound, levelDat.TagId()}
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

// Returns the x, y, z coordinated for the current spawn point.
func (l *Level) GetSpawn() (x, y, z int32) {
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

// Sets the spawn point.
func (l *Level) SetSpawn(x, y, z int32) {
	l.levelData.Set(nbt.NewTag("SpawnX", nbt.NewInt(x)))
	l.levelData.Set(nbt.NewTag("SpawnY", nbt.NewInt(y)))
	l.levelData.Set(nbt.NewTag("SpawnZ", nbt.NewInt(z)))
	l.changed = true
}

// Get the block at coordinates x, y, z.
func (l *Level) GetBlock(x, y, z int32) (*Block, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return nil, err
	} else if c == nil {
		return &Block{}, nil
	}
	return c.GetBlock(x, y, z), nil
}

// Sets the block at coordinates x, y, z. Also processes any lighting updates if applicable.
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
		if l.lighting.SkyLight() {
			baseX := x &^ 15
			baseY := y &^ 15
			baseZ := z &^ 15
			for i := baseX; i < baseX+16; i++ {
				for k := baseZ; k < baseZ+16; k++ {
					j := baseY + 15
					if h := c.GetHeight(i, k); h < baseY || h == 0 {
						for ; j >= baseY; j-- {
							c.SetSkyLight(i, j, k, 15)
						}
					} else if l.lighting.SkyLightSimple() {
						for currLightLevel := c.GetSkyLight(i, j+1, k); j >= baseY; j-- {
							if currLightLevel > 0 {
								currLightLevel--
							}
							c.SetSkyLight(i, j, k, currLightLevel)
						}
					}
				}
			}
			if l.lighting.SkyLightAll() {
				for a := int32(0); a < 16; a++ {
					for b := int32(0); b < 16; b++ {
						if c.GetHeight(baseX+a, baseZ+b) > baseY {
							if err = l.genLighting(baseX+a, baseY+15, baseZ+b, true, false, 0); err != nil {
								return err
							}
							if err = l.genLighting(baseX+a, baseY, baseZ+b, true, false, 0); err != nil {
								return err
							}
						}
						if err = l.genLighting(baseX+a, baseY+15-b, baseZ, true, false, 0); err != nil {
							return err
						}
						if err = l.genLighting(baseX, baseY+15-b, baseZ+a, true, false, 0); err != nil {
							return err
						}
						if err = l.genLighting(baseX+a, baseY+15-b, baseZ+15, true, false, 0); err != nil {
							return err
						}
						if err = l.genLighting(baseX+15, baseY+15-b, baseZ+a, true, false, 0); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	var opacity uint8
	if l.lighting.SkyLight() {
		opacity = c.GetOpacity(x, y, z)
	}
	c.SetBlock(x, y, z, block)
	if l.lighting.SkyLight() && block.Opacity() != opacity {
		if l.lighting.SkyLightSimple() {
			nY := y
			for h := c.GetHeight(x, z); nY >= h; nY-- {
				c.SetSkyLight(x, nY, z, 15)
			}
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
		} else if err = l.genLighting(x, y, z, true, block.Opacity() > opacity, 0); err != nil {
			return err
		}
	}
	if bl := c.GetBlockLight(x, y, z); l.lighting.BlockLight() && (block.Light() != bl || block.Opacity() != opacity) {
		if l.lighting.BlockLightSimple() {
			c.SetBlockLight(x, y, z, block.Light())
		} else {
			if err = l.genLighting(x, y, z, false, block.Light() < bl, block.Light()); err != nil {
				return err
			}
		}
	}
	return nil
}

type lightCoords struct {
	x, y, z    int32
	lightLevel uint8
}

func (l *Level) genLighting(x, y, z int32, skyLight, darker bool, source uint8) error {
	var (
		getLight func(*chunk, int32, int32, int32) uint8
		setLight func(*chunk, int32, int32, int32, uint8)
	)
	if skyLight {
		getLight = (*chunk).GetSkyLight
		setLight = (*chunk).SetSkyLight
	} else {
		getLight = (*chunk).GetBlockLight
		setLight = (*chunk).SetBlockLight
	}
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return err
	} else if c == nil {
		return nil
	}
	list := make([]*lightCoords, 1)
	list[0] = &lightCoords{x, y, z, getLight(c, x, y, z)}
	changed := boolmap.NewMap()
	changed.Set((uint64(y)<<10)|(uint64(z&31)<<5)|uint64(x&31), true)
	if darker { // reset lighting on all blocks affected by the changed one (only applies if darker)
		setLight(c, x, y, z, 0)
		for i := 0; i < len(list); i++ {
			for _, s := range surroundingBlocks(list[i].x, list[i].y, list[i].z) {
				mx, my, mz := s[0], s[1], s[2]
				pos := (uint64(my) << 10) | (uint64(mz&31) << 5) | uint64(mx&31)
				if changed.Get(pos) {
					continue
				}
				if c, err = l.getChunk(mx, mz, false); err != nil {
					return err
				} else if c == nil {
					continue
				}
				shouldBe := list[i].lightLevel
				if opacity := c.GetOpacity(mx, my, mz); opacity > shouldBe {
					shouldBe = 0
				} else {
					shouldBe -= opacity
				}
				if thisLight := getLight(c, mx, my, mz); thisLight == shouldBe && shouldBe != 0 || (skyLight && thisLight == 15 && my < c.GetHeight(mx, mz)) {
					list = append(list, &lightCoords{mx, my, mz, thisLight})
					changed.Set(pos, true)
					if thisLight > 0 {
						setLight(c, mx, my, mz, 0)
					}
				}
			}
		}
	} // end lighting reset
	if !skyLight && source > 0 {
		c, _ = l.getChunk(x, z, false)
		c.SetBlockLight(x, y, z, source)
		list = list[1:]
	}
	for ; len(list) > 0; list = list[1:] {
		mx, my, mz := list[0].x, list[0].y, list[0].z
		newLight := uint8(0)
		c, _ = l.getChunk(mx, mz, false)
		if opacity := c.GetOpacity(mx, my, mz); skyLight && my >= c.GetHeight(mx, mz) { //Determine correct light level...
			newLight = 15
		} else if opacity == 15 {
			newLight = 0
		} else {
			var d *chunk
			for _, s := range surroundingBlocks(mx, my, mz) {
				nx, ny, nz := s[0], s[1], s[2]
				if d, err = l.getChunk(nx, nz, false); err != nil {
					return err
				} else if d == nil {
					continue
				}
				curr := getLight(d, nx, ny, nz)
				if curr < opacity {
					continue
				}
				curr -= opacity
				if curr > newLight {
					newLight = curr
				}
			}
		} // ...end determining light level
		setLight(c, mx, my, mz, newLight)
		if newLight > list[0].lightLevel {
			for _, s := range surroundingBlocks(mx, my, mz) {
				mx, my, mz = s[0], s[1], s[2]
				pos := (uint64(my) << 10) | (uint64(mz&31) << 5) | uint64(mx&31)
				if changed.Get(pos) {
					continue
				}
				if c, err = l.getChunk(mx, mz, false); err != nil {
					return err
				} else if c == nil {
					continue
				}
				if thisLight := getLight(c, mx, my, mz); thisLight < newLight {
					list = append(list, &lightCoords{mx, my, mz, thisLight})
					changed.Set(pos, true)
				}
			}
		}
	}
	return nil
}

// Returns the biome for the column x, z.
func (l *Level) GetBiome(x, z int32) (Biome, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return Biome_Auto, err
	} else if c == nil {
		return Biome_Plains, nil
	}
	return c.GetBiome(x, z), nil
}

// Sets the biome for the column x, z.
func (l *Level) SetBiome(x, z int32, biome Biome) error {
	c, err := l.getChunk(x, z, true)
	if err != nil {
		return err
	}
	c.SetBiome(x, z, biome)
	return nil
}

// Returns the name of the minecraft level.
func (l *Level) GetName() string {
	s := l.levelData.Get("LevelName").Data().(*nbt.String)
	return string(*s)
}

// Sets the name of the minecraft level.
func (l *Level) SetName(name string) {
	l.levelData.Set(nbt.NewTag("LevelName", nbt.NewString(name)))
	l.changed = true
}

func (l *Level) getChunk(x, z int32, create bool) (*chunk, error) {
	x >>= 4
	z >>= 4
	pos := uint64(z)<<28 | uint64(uint32(x))
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

// Saves all open chunks, but does not close them.
func (l *Level) Save() error {
	if l.changed {
		if err := l.path.WriteLevelDat(nbt.NewTag("", nbt.NewCompound([]nbt.Tag{nbt.NewTag("Data", l.levelData)}))); err != nil {
			return err
		}
		l.changed = false
	}
	toSave := make([]nbt.Tag, 0)
	for n, c := range l.chunks {
		if l.changes.Get(uint64(n)) {
			toSave = append(toSave, c.GetNBT())
		}
	}
	l.changes = boolmap.NewMap()
	if len(toSave) > 0 {
		return l.path.SetChunk(toSave...) //check multi-error
	}
	return nil
}

// Closes all open chunks, but does not save them.
func (l *Level) Close() {
	l.changed = false
	l.chunks = make(map[uint64]*chunk)
	l.changes = boolmap.NewMap()
}

func surroundingBlocks(x, y, z int32) [][3]int32 {
	sB := [6][3]int32{
		{x, y - 1, z},
		{x - 1, y, z},
		{x + 1, y, z},
		{x, y, z - 1},
		{x, y, z + 1},
		{x, y + 1, z},
	}
	if y == 0 {
		return sB[1:]
	}
	if y == 255 {
		return sB[:5]
	}
	return sB[:]
}
