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
	tileTicks    map[uint16]*nbt.Compound
}

func (c *chunk) GetNBT() nbt.Tag {
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
	for _, cmp := range c.tileTicks {
		if cmp != nil {
			tileTicks = append(tileTicks, cmp)
		}
	}
	if len(tileTicks) == 0 {
		data.Remove("TileTicks")
	} else {
		data.Set(nbt.NewTag("TileTicks", nbt.NewList(tileTicks)))
	}
	return nbt.NewTag("", nbt.NewCompound([]nbt.Tag{nbt.NewTag("Level", data)}))
}

func newChunk(x, z int32, data nbt.Tag) (*chunk, error) {
	if data == nil {
		biomes := make([]int8, 256)
		for i := 0; i < 256; i++ {
			biomes[i] = -1
		}
		data = nbt.NewTag("", nbt.NewCompound([]nbt.Tag{
			nbt.NewTag("Level", nbt.NewCompound([]nbt.Tag{
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
	c.tileTicks = make(map[uint16]*nbt.Compound)
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
				c.tileTicks[xyz(x, y, z)] = tag
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

func (c *chunk) GetBlock(x, y, z int32) (b *Block, err error) {
	ys := y >> 4
	if c.sections[ys] == nil {
		b = &Block{}
	} else {
		if b, err = c.sections[ys].GetBlock(x, y, z); err != nil {
			return
		}
		pos := xyz(x, y, z)
		if md, ok := c.tileEntities[pos]; ok && md != nil {
			b.SetMetadata([]nbt.Tag(*md))
		}
		if tt, ok := c.tileTicks[pos]; ok && tt != nil {
			b.Tick = true
		}
	}
	return
}

func (c *chunk) SetBlock(x, y, z int32, b *Block) error {
	ys := y >> 4
	if c.sections[ys] == nil {
		if b.Equal(&Block{}) {
			return nil
		}
		c.sections[ys] = newSection(y)
	}
	if err := c.sections[ys].SetBlock(x, y, z, b); err != nil {
		return err
	}
	if hmpos := x&15<<4 | z&15; b.Opacity() <= 1 { //All transparent blocks block 1 light when they are below the highest non-transparent block
		if y == (*c.heightMap)[hmpos] {
			(*c.heightMap)[hmpos] = 0
			for i := y; i >= 0; i-- {
				if c.sections[ys] != nil {
					if tB, _ := c.sections[ys].GetBlock(x, y, z); tB.Opacity() <= 1 {
						(*c.heightMap)[hmpos] = i
						break
					}
				}
			}
		}
	} else if y > (*c.heightMap)[hmpos] {
		(*c.heightMap)[hmpos] = y
	}
	pos := xyz(x, y, z)
	if b.metadata == nil {
		delete(c.tileEntities, pos)
	} else {
		c.tileEntities[pos] = nbt.NewCompound(b.GetMetadata())
		c.tileEntities[pos].Set(nbt.NewTag("x", nbt.NewInt(x)))
		c.tileEntities[pos].Set(nbt.NewTag("y", nbt.NewInt(y)))
		c.tileEntities[pos].Set(nbt.NewTag("z", nbt.NewInt(z)))
	}
	if b.Tick {
		c.tileTicks[pos] = nbt.NewCompound([]nbt.Tag{
			nbt.NewTag("i", nbt.NewInt(int32(b.BlockId))),
			nbt.NewTag("p", nbt.NewInt(1)),
			nbt.NewTag("t", nbt.NewInt(-1)),
			nbt.NewTag("x", nbt.NewInt(x)),
			nbt.NewTag("y", nbt.NewInt(y)),
			nbt.NewTag("z", nbt.NewInt(z)),
		})
	} else {
		delete(c.tileTicks, pos)
	}
	return nil
}

func (c *chunk) GetBiome(x, z int32) Biome {
	return Biome((*c.biomes)[x&15<<4|z&15])
}

func (c *chunk) SetBiome(x, z int32, b Biome) error {
	(*c.biomes)[x&15<<4|z&15] = int8(b)
	return nil
}

func (c *chunk) IsEmpty() bool {
	for i := 0; i < 16; i++ {
		if c.sections[i] != nil {
			if !c.sections[i].IsEmpty() {
				return false
			}
		}
	}
	return true
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
