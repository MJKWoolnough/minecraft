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
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbt"
)

var (
	chunkRequired = []struct {
		name string
		nbt.TagId
	}{
		{"HeightMap", nbt.Tag_IntArray},
		{"InhabitedTime", nbt.Tag_Long},
		{"LastUpdate", nbt.Tag_Long},
		{"Sections", nbt.Tag_List},
		{"TerrainPopulated", nbt.Tag_Byte},
		{"xPos", nbt.Tag_Int},
		{"zPos", nbt.Tag_Int},
	}
	chunkOther = []struct {
		name      string
		tagType   nbt.TagId
		listType  nbt.TagId
		emptyByte bool
	}{
		{"Biomes", nbt.Tag_ByteArray, nbt.Tag_Byte, false},
		{"Entities", nbt.Tag_List, nbt.Tag_Compound, true},
		{"TileEntities", nbt.Tag_List, nbt.Tag_Compound, true},
		{"TileTicks", nbt.Tag_List, nbt.Tag_Compound, false},
	}
)

type chunk struct {
	sections     [16]*section
	biomes       *nbt.ByteArray
	data         *nbt.Compound
	heightMap    *nbt.IntArray
	tileEntities map[uint16]*nbt.Compound
	tileTicks    map[uint16][]*nbt.Compound
}

func (c *chunk) GetNBT() *nbt.Tag {
	data := c.data.Copy().(*nbt.Compound)
	sections := make([]nbt.Data, 0, 16)
	for i := 0; i < 16; i++ {
		if c.sections[i] != nil {
			sections = append(sections, c.sections[i].section)
		}
	}
	sectionList := nbt.NewEmptyList(nbt.Tag_Compound)
	sectionList.Append(sections...)
	data.Set(nbt.NewTag("Sections", sectionList))
	tileEntities := make([]nbt.Data, 0)
	for _, cmp := range c.tileEntities {
		if cmp != nil {
			tileEntities = append(tileEntities, cmp)
		}
	}
	data.Set(nbt.NewTag("TileEntities", nbt.NewList(tileEntities)))
	tileTicks := make([]nbt.Data, 0)
	for _, cmpa := range c.tileTicks {
		for _, cmp := range cmpa {
			tileTicks = append(tileTicks, cmp)
		}
	}
	if len(tileTicks) == 0 {
		data.Remove("TileTicks")
	} else {
		data.Set(nbt.NewTag("TileTicks", nbt.NewList(tileTicks)))
	}
	return nbt.NewTag("", nbt.NewCompound(nbt.Compound{nbt.NewTag("Level", data)}))
}

func newChunk(x, z int32, data *nbt.Tag) (*chunk, error) {
	if data == nil {
		biomes := make([]int8, 256)
		for i := 0; i < 256; i++ {
			biomes[i] = -1
		}
		data = nbt.NewTag("", nbt.NewCompound(nbt.Compound{
			nbt.NewTag("Level", nbt.NewCompound(nbt.Compound{
				nbt.NewTag("xPos", nbt.NewInt(x)),
				nbt.NewTag("zPos", nbt.NewInt(z)),
				nbt.NewTag("Biomes", nbt.NewByteArray(biomes)),
				nbt.NewTag("HeightMap", nbt.NewIntArray(make([]int32, 256))),
				nbt.NewTag("InhabitedTime", nbt.NewLong(0)),
				nbt.NewTag("LastUpdate", nbt.NewLong(0)),
				nbt.NewTag("Sections", nbt.NewEmptyList(nbt.Tag_Compound)),
				nbt.NewTag("TerrainPopulated", nbt.NewByte(1)),
			})),
		}))
	}
	c := new(chunk)
	if data.TagId() != nbt.Tag_Compound {
		return nil, &WrongTypeError{"[Chunk Base]", nbt.Tag_Compound, data.TagId()}
	} else if tag := data.Data().(*nbt.Compound).Get("Level"); tag == nil {
		return nil, &MissingTagError{"[Chunk Base]->Level"}
	} else if tag.TagId() != nbt.Tag_Compound {
		return nil, &WrongTypeError{"[Chunk Base]->Level", nbt.Tag_Compound, tag.TagId()}
	} else {
		c.data = tag.Data().(*nbt.Compound)
	}
	for _, req := range chunkRequired {
		if tag := c.data.Get(req.name); tag == nil {
			return nil, &MissingTagError{req.name}
		} else if tagId := tag.TagId(); tagId != req.TagId {
			return nil, &WrongTypeError{req.name, req.TagId, tagId}
		}
	}

	if tX := int32(*c.data.Get("xPos").Data().(*nbt.Int)); tX != x {
		return nil, &UnexpectedValue{"[Chunk Base]->Level->xPos", fmt.Sprintf("%d", x), fmt.Sprintf("%d", tX)}
	}
	if tZ := int32(*c.data.Get("zPos").Data().(*nbt.Int)); tZ != z {
		return nil, &UnexpectedValue{"[Chunk Base]->Level->zPos", fmt.Sprintf("%d", z), fmt.Sprintf("%d", tZ)}
	}

	for _, co := range chunkOther {
		if tag := c.data.Get(co.name); tag == nil {
			continue
		} else if tagId := tag.TagId(); tagId != co.tagType {
			return nil, &WrongTypeError{co.name, co.tagType, tagId}
		} else if tagId == nbt.Tag_List {
			list := tag.Data().(*nbt.List)
			if list.TagType() != co.listType {
				if co.emptyByte && list.TagType() == nbt.Tag_Byte && list.Len() == 0 {
					continue
				}
				return nil, &WrongTypeError{co.name, co.listType, list.TagType()}
			}
		}
	}

	if biomes := c.data.Get("Biomes"); biomes != nil {
		c.biomes = biomes.Data().(*nbt.ByteArray)
	} else {
		biomes := make([]int8, 256)
		for i := 0; i < 256; i++ {
			biomes[i] = -1
		}
		c.biomes = nbt.NewByteArray(biomes)
		c.data.Set(nbt.NewTag("Biomes", c.biomes))
	}
	c.heightMap = c.data.Get("HeightMap").Data().(*nbt.IntArray)
	c.tileEntities = make(map[uint16]*nbt.Compound)
	if tileEntities := c.data.Get("TileEntities"); tileEntities != nil {
		if lTileEntities, ok := tileEntities.Data().(*nbt.List); ok {
			for i := 0; i < lTileEntities.Len(); i++ {
				tag := lTileEntities.Get(i).(*nbt.Compound)
				if tag == nil {
					return nil, &MissingTagError{"TileEntities->Child"}
				}
				x, y, z, err := getCoords(tag)
				if err != nil {
					return nil, err
				}
				c.tileEntities[xyz(x, y, z)] = tag
			}
		}
	}
	c.data.Remove("TileEntities")
	c.tileTicks = make(map[uint16][]*nbt.Compound)
	if tileTicks := c.data.Get("TileTicks"); tileTicks != nil {
		if lTileTicks, ok := tileTicks.Data().(*nbt.List); ok {
			for i := 0; i < lTileTicks.Len(); i++ {
				tag := lTileTicks.Get(i).(*nbt.Compound)
				if tag == nil {
					return nil, &MissingTagError{"TileTicks->Child"}
				}
				x, y, z, err := getCoords(tag)
				if err != nil {
					return nil, err
				}
				if id := tag.Get("i"); id == nil {
					return nil, &MissingTagError{"TileTicks->Child->i"}
				} else if j := id.TagId(); j != nbt.Tag_Int {
					return nil, &WrongTypeError{"TileTicks->Child->i", nbt.Tag_Int, j}
				}
				if t := tag.Get("t"); t == nil {
					return nil, &MissingTagError{"TileTicks->Child->t"}
				} else if j := t.TagId(); j != nbt.Tag_Int {
					return nil, &WrongTypeError{"TileTicks->Child->t", nbt.Tag_Int, j}
				}
				if p := tag.Get("p"); p == nil {
					return nil, &MissingTagError{"TileTicks->Child->p"}
				} else if j := p.TagId(); j != nbt.Tag_Int {
					return nil, &WrongTypeError{"TileTicks->Child->p", nbt.Tag_Int, j}
				}
				pos := xyz(x, y, z)
				c.tileTicks[pos] = append(c.tileTicks[pos], tag)
			}
		}
	}
	c.data.Remove("TileTicks")
	sections := c.data.Get("Sections").Data().(*nbt.List)
	for i := 0; i < sections.Len(); i++ {
		section := sections.Get(i).(*nbt.Compound)
		if yc := section.Get("Y"); yc == nil {
			return nil, &MissingTagError{"Sections->Child->Y"}
		} else if yc.TagId() != nbt.Tag_Byte {
			return nil, &WrongTypeError{"Sections->Child->Y", nbt.Tag_Byte, yc.TagId()}
		} else {
			y := int32(*yc.Data().(*nbt.Byte))
			var err error
			if c.sections[y], err = loadSection(section); err != nil {
				return nil, err
			}
		}
	}
	c.data.Remove("Sections")
	return c, nil
}

func (c *chunk) GetBlock(x, y, z int32) *Block {
	ys := y >> 4
	if c.sections[ys] == nil {
		return &Block{}
	}
	b := c.sections[ys].GetBlock(x, y, z)
	pos := xyz(x, y, z)
	if md, ok := c.tileEntities[pos]; ok && md != nil {
		b.SetMetadata(*md)
	}
	if tt, ok := c.tileTicks[pos]; ok && tt != nil {
		b.ticks = make([]Tick, len(tt))
		for n, tick := range tt {
			b.ticks[n] = Tick{
				int32(*tick.Get("i").Data().(*nbt.Int)),
				int32(*tick.Get("t").Data().(*nbt.Int)),
				int32(*tick.Get("p").Data().(*nbt.Int)),
			}
		}
	}
	return b
}

func (c *chunk) SetBlock(x, y, z int32, b *Block) {
	ys := y >> 4
	if c.sections[ys] == nil {
		if b.Equal(&Block{}) {
			return
		}
		c.sections[ys] = newSection(y)
	}
	c.sections[ys].SetBlock(x, y, z, b)
	if hmpos := x&15<<4 | z&15; b.Opacity() <= 1 { //All transparent blocks block 1 light when they are below the highest non-transparent block
		if y == (*c.heightMap)[hmpos]-1 {
			(*c.heightMap)[hmpos] = 0
			for i := y; i >= 0; i-- {
				if yt := i >> 4; c.sections[yt] != nil {
					if c.sections[yt].GetOpacity(x, i, z) > 1 {
						(*c.heightMap)[hmpos] = i + 1
						break
					}
				}
			}
		}
	} else if y >= (*c.heightMap)[hmpos] {
		(*c.heightMap)[hmpos] = y + 1
	}
	pos := xyz(x, y, z)
	if b.metadata == nil {
		delete(c.tileEntities, pos)
	} else {
		comp := b.GetMetadata()
		c.tileEntities[pos] = &comp
		c.tileEntities[pos].Set(nbt.NewTag("x", nbt.NewInt(x)))
		c.tileEntities[pos].Set(nbt.NewTag("y", nbt.NewInt(y)))
		c.tileEntities[pos].Set(nbt.NewTag("z", nbt.NewInt(z)))
	}
	if b.HasTicks() {
		ticks := b.GetTicks()
		c.tileTicks[pos] = make([]*nbt.Compound, len(ticks))
		for n, tick := range ticks {
			c.tileTicks[pos][n] = nbt.NewCompound(nbt.Compound{
				nbt.NewTag("i", nbt.NewInt(tick.I)),
				nbt.NewTag("p", nbt.NewInt(tick.P)),
				nbt.NewTag("t", nbt.NewInt(tick.T)),
				nbt.NewTag("x", nbt.NewInt(x)),
				nbt.NewTag("y", nbt.NewInt(y)),
				nbt.NewTag("z", nbt.NewInt(z)),
			})
		}
	} else {
		delete(c.tileTicks, pos)
	}
}

func (c *chunk) GetBiome(x, z int32) Biome {
	return Biome((*c.biomes)[x&15<<4|z&15])
}

func (c *chunk) SetBiome(x, z int32, b Biome) {
	(*c.biomes)[x&15<<4|z&15] = int8(b)
}

func (c *chunk) GetOpacity(x, y, z int32) uint8 {
	if y >= (*c.heightMap)[x&15<<4|z&15] {
		return 1
	}
	ys := y >> 4
	if c.sections[ys] == nil {
		return 1
	}
	return c.sections[ys].GetOpacity(x, y, z)
}

func (c *chunk) GetHeight(x, z int32) int32 {
	return (*c.heightMap)[x&15<<4|z&15]
}

func (c *chunk) GetBlockLight(x, y, z int32) uint8 {
	ys := y >> 4
	if ys < 16 && c.sections[ys] == nil {
		return 0
	} else if y > 255 {
		return 15
	} else if y < 0 {
		return 0
	}
	return c.sections[ys].GetBlockLight(x, y, z)
}

func (c *chunk) SetBlockLight(x, y, z int32, l uint8) {
	ys := y >> 4
	if ys < 16 && c.sections[ys] != nil {
		c.sections[ys].SetBlockLight(x, y, z, l)
	}
}

func (c *chunk) GetSkyLight(x, y, z int32) uint8 {
	if ys := y >> 4; ys >= 0 && ys < 16 && c.sections[ys] != nil {
		return c.sections[ys].GetSkyLight(x, y, z)
	} else if y >= (*c.heightMap)[x&15<<4|z&15] || y > 255 {
		return 15
	} else if y < 0 {
		return 0
	} else if ys < 15 && c.sections[ys+1] != nil {
		sl := c.sections[ys+1].GetSkyLight(x, 0, z)
		if d := uint8((ys+1)<<4 - y); d < sl {
			sl -= d
		} else {
			sl = 0
		}
		return sl
	}
	return 0
}

func (c *chunk) SetSkyLight(x, y, z int32, l uint8) {
	ys := y >> 4
	if ys < 16 && c.sections[ys] != nil {
		c.sections[ys].SetSkyLight(x, y, z, l)
	}
}

func (c *chunk) createSection(y int32) bool {
	if ys := y >> 4; ys >= 0 && ys < 16 && c.sections[ys] == nil {
		c.sections[ys] = newSection(y)
		return true
	}
	return false
}

func xyz(x, y, z int32) uint16 {
	return (uint16(y) << 8) | (uint16(z&15) << 4) | uint16(x&15)
}

func getCoord(name string, data *nbt.Compound) (int32, error) {
	tag := data.Get(name)
	if tag == nil {
		return 0, &MissingTagError{name}
	} else if tag.TagId() != nbt.Tag_Int {
		return 0, &WrongTypeError{name, nbt.Tag_Int, tag.TagId()}
	}
	return int32(*tag.Data().(*nbt.Int)), nil
}

func getCoords(data *nbt.Compound) (x, y, z int32, err error) {
	if x, err = getCoord("x", data); err != nil {
		return
	} else if y, err = getCoord("y", data); err != nil {
		return
	}
	z, err = getCoord("z", data)
	return
}
