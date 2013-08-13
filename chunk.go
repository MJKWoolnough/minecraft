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
	"bytes"
	"compress/zlib"
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

type NoCompressed struct{}

func (n NoCompressed) Error() string {
	return "no compressed data found"
}

type chunk struct {
	sections [16]*section
	biomes   *nbt.ByteArray
	data     *nbt.Compound
	// 	entities     map[uint16]*nbt.Compound
	tileEntities map[uint16]*nbt.Compound
	tileTicks    map[uint16]*nbt.Compound
	compressed   *bytes.Buffer
}

func newChunk(x, z int32) *chunk {
	buf := new(bytes.Buffer)
	biomes := make([]int8, 256)
	for i := 0; i < 256; i++ {
		biomes[i] = -1
	}
	zBuf := zlib.NewWriter(buf)
	nbt.NewTag("", nbt.NewCompound([]nbt.Tag{
		nbt.NewTag("Level", nbt.NewCompound([]nbt.Tag{
			nbt.NewTag("Biomes", nbt.NewByteArray(biomes)),
			nbt.NewTag("HeightMap", nbt.NewIntArray(make([]int32, 256))),
			nbt.NewTag("InhabitedTime", nbt.NewLong(0)),
			nbt.NewTag("LastUpdate", nbt.NewLong(0)),
			nbt.NewTag("Sections", nbt.NewEmptyList(nbt.Tag_Compound)),
			nbt.NewTag("TerrainPopulated", nbt.NewByte(1)),
			nbt.NewTag("xPos", nbt.NewInt(x)),
			nbt.NewTag("zPos", nbt.NewInt(z)),
		})),
	})).WriteTo(zBuf)
	zBuf.Close()
	return &chunk{compressed: buf}
}

func (c *chunk) compress() error {
	if c.compressed != nil {
		return nil
	}
	cx := int32(*c.data.Get("xPos").Data().(*nbt.Int)) << 4
	cz := int32(*c.data.Get("zPos").Data().(*nbt.Int)) << 4
	sections := make([]nbt.Data, 0, 16)
	for i := 0; i < 16; i++ {
		if c.sections[i] != nil {
			sections = append(sections, c.sections[i].section)
			c.sections[i] = nil
		}
	}
	sectionList := nbt.NewEmptyList(nbt.Tag_Compound)
	sectionList.Append(sections...)
	c.data.Set(nbt.NewTag("Sections", sectionList))
	tileEntities := make([]nbt.Data, 0)
	for pos, cmp := range c.tileEntities {
		if cmp != nil {
			x, y, z := int32(pos)&15, int32(pos>>8), int32(pos>>4)&15
			cmp.Set(nbt.NewTag("x", nbt.NewInt(cx+x)))
			cmp.Set(nbt.NewTag("y", nbt.NewInt(y)))
			cmp.Set(nbt.NewTag("z", nbt.NewInt(cz+z)))
			tileEntities = append(tileEntities, cmp)
			delete(c.tileEntities, pos)
		}
	}
	c.data.Set(nbt.NewTag("TileEntities", nbt.NewList(tileEntities)))
	c.tileEntities = nil
	tileTicks := make([]nbt.Data, 0)
	for pos, cmp := range c.tileTicks {
		if cmp != nil {
			x, y, z := int32(pos)&15, int32(pos>>8), int32(pos>>4)&15
			cmp.Set(nbt.NewTag("x", nbt.NewInt(cx+x)))
			cmp.Set(nbt.NewTag("y", nbt.NewInt(y)))
			cmp.Set(nbt.NewTag("z", nbt.NewInt(cz+z)))
			tileTicks = append(tileTicks, cmp)
			delete(c.tileTicks, pos)
		}
	}
	if len(tileTicks) == 0 {
		c.data.Remove("TileTicks")
	} else {
		c.data.Set(nbt.NewTag("TileTicks", nbt.NewList(tileTicks)))
	}
	c.tileTicks = nil
	c.compressed = new(bytes.Buffer)
	zBuf := zlib.NewWriter(c.compressed)
	_, err := nbt.NewTag("", nbt.NewCompound([]nbt.Tag{nbt.NewTag("Level", c.data)})).WriteTo(zBuf)
	zBuf.Close()
	c.data = nil
	return err
}

func (c *chunk) decompress() error {
	if c.compressed == nil {
		return nil
	}
	if c.compressed.Len() == 0 {
		return &NoCompressed{}
	}
	if zr, err := zlib.NewReader(c.compressed); err != nil {
		return err
	} else if tag, _, err := nbt.ReadNBTFrom(zr); err != nil {
		return err
	} else if tag.TagId() != nbt.Tag_Compound {
		return &WrongTypeError{"[Chunk Base]", nbt.Tag_Compound, tag.TagId()}
	} else if tagL := tag.Data().(*nbt.Compound).Get("Level"); tagL == nil {
		return &MissingTagError{"[Chunk Base]->Level"}
	} else if tagL.TagId() != nbt.Tag_Compound {
		return &WrongTypeError{"[Chunk Base]->Level", nbt.Tag_Compound, tagL.TagId()}
	} else {
		c.data = tagL.Data().(*nbt.Compound)
	}
	for _, req := range chunkRequired {
		if tag := c.data.Get(req.name); tag == nil {
			return &MissingTagError{req.name}
		} else if tagId := tag.TagId(); tagId != req.TagId {
			return &WrongTypeError{req.name, req.TagId, tagId}
		}
	}
	for _, co := range chunkOther {
		if tag := c.data.Get(co.name); tag == nil {
			continue
		} else if tagId := tag.TagId(); tagId != co.tagType {
			return &WrongTypeError{co.name, co.tagType, tagId}
		} else if tagId == nbt.Tag_List {
			list := tag.Data().(*nbt.List)
			if list.TagType() != co.listType {
				if co.emptyByte && list.TagType() == nbt.Tag_Byte && list.Len() == 0 {
					continue
				}
				return &WrongTypeError{co.name, co.listType, list.TagType()}
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
	c.tileEntities = make(map[uint16]*nbt.Compound)
	if tileEntities := c.data.Get("TileEntities"); tileEntities != nil {
		if lTileEntities, ok := tileEntities.Data().(*nbt.List); ok {
			for i := 0; i < lTileEntities.Len(); i++ {
				tag := lTileEntities.Get(i).(*nbt.Compound)
				if tag == nil {
					return &MissingTagError{"TileEntities->Child"}
				}
				x, y, z, err := getCoords(tag)
				if err != nil {
					return err
				}
				tag.Remove("x")
				tag.Remove("y")
				tag.Remove("z")
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
					return &MissingTagError{"TileTicks->Child"}
				}
				x, y, z, err := getCoords(tag)
				if err != nil {
					return err
				}
				tag.Remove("x")
				tag.Remove("y")
				tag.Remove("z")
				c.tileTicks[xyz(x, y, z)] = tag
			}
		}
	}
	c.data.Remove("TileTicks")
	sections := c.data.Get("Sections").Data().(*nbt.List)
	for i := 0; i < sections.Len(); i++ {
		section := sections.Get(i).(*nbt.Compound)
		if yc := section.Get("Y"); yc == nil {
			return &MissingTagError{"Sections->Child->Y"}
		} else if yc.TagId() != nbt.Tag_Byte {
			return &WrongTypeError{"Sections->Child->Y", nbt.Tag_Byte, yc.TagId()}
		} else {
			y := int32(*yc.Data().(*nbt.Byte))
			var err error
			if c.sections[y], err = LoadSection(section); err != nil {
				return err
			}
		}
	}
	c.data.Remove("Sections")
	c.compressed = nil
	return nil
}

func (c *chunk) GetBlock(x, y, z int32) (b *Block, err error) {
	if err := c.decompress(); err != nil {
		return nil, err
	}
	ys := y >> 4
	if c.sections[ys] == nil {
		b = &Block{}
	} else {
		if b, err = c.sections[ys].GetBlock(x, y, z); err != nil {
			return
		}
		if md, ok := c.tileEntities[xyz(x, y, z)]; ok && md != nil {
			b.SetMetadata([]nbt.Tag(*md))
		}
		if tt, ok := c.tileTicks[xyz(x, y, z)]; ok && tt != nil {
			b.Tick = true
		}
	}
	return
}

func (c *chunk) SetBlock(x, y, z int32, b *Block) error {
	if err := c.decompress(); err != nil {
		return err
	}
	ys := y >> 4
	if c.sections[ys] == nil {
		if b.Equal(&Block{}) {
			return nil
		}
		c.sections[ys] = NewSection(y)
	}
	if err := c.sections[ys].SetBlock(x, y, z, b); err != nil {
		return err
	}
	if b.metadata == nil {
		delete(c.tileEntities, xyz(x, y, z))
	} else {
		c.tileEntities[xyz(x, y, z)] = nbt.NewCompound(b.metadata)
	}
	if b.Tick {
		c.tileTicks[xyz(x, y, z)] = nbt.NewCompound([]nbt.Tag{
			nbt.NewTag("i", nbt.NewInt(int32(b.BlockId))),
			nbt.NewTag("p", nbt.NewInt(1)),
			nbt.NewTag("t", nbt.NewInt(-1)),
		})
	} else {
		delete(c.tileTicks, xyz(x, y, z))
	}
	return nil
}

func (c *chunk) GetBiome(x, z int32) (Biome, error) {
	if err := c.decompress(); err != nil {
		return -1, err
	}
	return Biome((*c.biomes)[z&15<<4|x&15]), nil
}

func (c *chunk) SetBiome(x, z int32, b Biome) error {
	if err := c.decompress(); err != nil {
		return err
	}
	(*c.biomes)[z&15<<4|x&15] = int8(b)
	return nil
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
