package minecraft

import (
	"strconv"

	"vimagination.zapto.org/minecraft/nbt"
)

var (
	chunkRequired = []struct {
		name string
		nbt.TagID
	}{
		{"HeightMap", nbt.TagIntArray},
		{"InhabitedTime", nbt.TagLong},
		{"LastUpdate", nbt.TagLong},
		{"Sections", nbt.TagList},
		{"TerrainPopulated", nbt.TagByte},
		{"xPos", nbt.TagInt},
		{"zPos", nbt.TagInt},
	}
	chunkOther = []struct {
		name      string
		tagType   nbt.TagID
		listType  nbt.TagID
		emptyByte bool
	}{
		{"Biomes", nbt.TagByteArray, nbt.TagByte, false},
		{"Entities", nbt.TagList, nbt.TagCompound, true},
		{"TileEntities", nbt.TagList, nbt.TagCompound, true},
		{"TileTicks", nbt.TagList, nbt.TagCompound, false},
	}
)

type chunk struct {
	sections     [16]*section
	biomes       nbt.ByteArray
	data         nbt.Compound
	heightMap    nbt.IntArray
	tileEntities map[uint16]nbt.Compound
	tileTicks    map[uint16][]nbt.Compound
}

func (c *chunk) GetNBT() nbt.Tag {
	data := c.data.Copy().(nbt.Compound)
	sections := make([]nbt.Data, 0, 16)

	for i := 0; i < 16; i++ {
		if c.sections[i] != nil {
			sections = append(sections, c.sections[i].section)
		}
	}

	sectionList := nbt.NewEmptyList(nbt.TagCompound)
	tileEntities := make(nbt.ListCompound, 0, len(c.tileEntities))

	sectionList.Append(sections...)
	data.Set(nbt.NewTag("Sections", sectionList))

	for _, cmp := range c.tileEntities {
		if cmp != nil {
			tileEntities = append(tileEntities, cmp)
		}
	}

	data.Set(nbt.NewTag("TileEntities", &tileEntities))

	tileTicks := make(nbt.ListCompound, 0, len(c.tileTicks))

	for _, cmpa := range c.tileTicks {
		for _, cmp := range cmpa {
			tileTicks = append(tileTicks, cmp)
		}
	}

	if len(tileTicks) == 0 {
		data.Remove("TileTicks")
	} else {
		data.Set(nbt.NewTag("TileTicks", &tileTicks))
	}

	return nbt.NewTag("", nbt.Compound{nbt.NewTag("Level", data)})
}

func newChunk(x, z int32, data nbt.Tag) (*chunk, error) {
	if data.TagID() == 0 {
		biomes := make(nbt.ByteArray, 256)

		for i := 0; i < 256; i++ {
			biomes[i] = -1
		}

		data = nbt.NewTag("", nbt.Compound{
			nbt.NewTag("Level", nbt.Compound{
				nbt.NewTag("xPos", nbt.Int(x)),
				nbt.NewTag("zPos", nbt.Int(z)),
				nbt.NewTag("Biomes", biomes),
				nbt.NewTag("HeightMap", make(nbt.IntArray, 256)),
				nbt.NewTag("InhabitedTime", nbt.Long(0)),
				nbt.NewTag("LastUpdate", nbt.Long(0)),
				nbt.NewTag("Sections", nbt.NewEmptyList(nbt.TagCompound)),
				nbt.NewTag("TerrainPopulated", nbt.Byte(1)),
			}),
		})
	}

	c := new(chunk)

	if data.TagID() != nbt.TagCompound {
		return nil, WrongTypeError{"[Chunk Base]", nbt.TagCompound, data.TagID()}
	}

	tag := data.Data().(nbt.Compound).Get("Level")

	if tag.TagID() == 0 {
		return nil, MissingTagError{"[Chunk Base]->Level"}
	} else if tag.TagID() != nbt.TagCompound {
		return nil, WrongTypeError{"[Chunk Base]->Level", nbt.TagCompound, tag.TagID()}
	}

	c.data = tag.Data().(nbt.Compound)

	for _, req := range chunkRequired {
		if tag := c.data.Get(req.name); tag.TagID() == 0 {
			return nil, MissingTagError{req.name}
		} else if tagID := tag.TagID(); tagID != req.TagID {
			return nil, WrongTypeError{req.name, req.TagID, tagID}
		}
	}

	if tX := int32(c.data.Get("xPos").Data().(nbt.Int)); tX != x {
		return nil, UnexpectedValue{"[Chunk Base]->Level->xPos", strconv.FormatInt(int64(x), 10), strconv.FormatInt(int64(tX), 10)}
	}

	if tZ := int32(c.data.Get("zPos").Data().(nbt.Int)); tZ != z {
		return nil, UnexpectedValue{"[Chunk Base]->Level->zPos", strconv.FormatInt(int64(z), 10), strconv.FormatInt(int64(tZ), 10)}
	}

	for _, co := range chunkOther {
		if tag := c.data.Get(co.name); tag.TagID() == 0 {
			continue
		} else if tagID := tag.TagID(); tagID != co.tagType {
			return nil, WrongTypeError{co.name, co.tagType, tagID}
		} else if tagID == nbt.TagList {
			list := tag.Data().(nbt.List)

			if list.TagType() != co.listType {
				if co.emptyByte && list.Len() == 0 {
					if tt := list.TagType(); tt == nbt.TagByte || tt == nbt.TagEnd {
						continue
					}
				}

				return nil, WrongTypeError{co.name, co.listType, list.TagType()}
			}
		}
	}

	if biomes := c.data.Get("Biomes"); biomes.TagID() != 0 {
		c.biomes = biomes.Data().(nbt.ByteArray)
	} else {
		c.biomes = make(nbt.ByteArray, 256)

		for i := 0; i < 256; i++ {
			c.biomes[i] = -1
		}

		c.data.Set(nbt.NewTag("Biomes", c.biomes))
	}

	c.heightMap = c.data.Get("HeightMap").Data().(nbt.IntArray)
	c.tileEntities = make(map[uint16]nbt.Compound)

	if tileEntities := c.data.Get("TileEntities"); tileEntities.TagID() != 0 {
		if lTileEntities, ok := tileEntities.Data().(*nbt.ListCompound); ok {
			for _, tag := range *lTileEntities {
				if tag == nil {
					return nil, MissingTagError{"TileEntities->Child"}
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

	c.tileTicks = make(map[uint16][]nbt.Compound)

	if tileTicks := c.data.Get("TileTicks"); tileTicks.TagID() != 0 {
		if lTileTicks, ok := tileTicks.Data().(*nbt.ListCompound); ok {
			for _, tag := range *lTileTicks {
				if tag == nil {
					return nil, MissingTagError{"TileTicks->Child"}
				}

				x, y, z, err := getCoords(tag)
				if err != nil {
					return nil, err
				} else if id := tag.Get("i"); id.TagID() == 0 {
					return nil, MissingTagError{"TileTicks->Child->i"}
				} else if j := id.TagID(); j != nbt.TagInt {
					return nil, WrongTypeError{"TileTicks->Child->i", nbt.TagInt, j}
				} else if t := tag.Get("t"); t.TagID() == 0 {
					return nil, MissingTagError{"TileTicks->Child->t"}
				} else if j := t.TagID(); j != nbt.TagInt {
					return nil, WrongTypeError{"TileTicks->Child->t", nbt.TagInt, j}
				} else if p := tag.Get("p"); p.TagID() == 0 {
					return nil, MissingTagError{"TileTicks->Child->p"}
				} else if j := p.TagID(); j != nbt.TagInt {
					return nil, WrongTypeError{"TileTicks->Child->p", nbt.TagInt, j}
				}

				pos := xyz(x, y, z)
				c.tileTicks[pos] = append(c.tileTicks[pos], tag)
			}
		}
	}

	c.data.Remove("TileTicks")

	for _, section := range *(c.data.Get("Sections").Data().(*nbt.ListCompound)) {
		if yc := section.Get("Y"); yc.TagID() == 0 {
			return nil, MissingTagError{"Sections->Child->Y"}
		} else if yc.TagID() != nbt.TagByte {
			return nil, WrongTypeError{"Sections->Child->Y", nbt.TagByte, yc.TagID()}
		} else {
			y := int32(yc.Data().(nbt.Byte))

			var err error

			if c.sections[y], err = loadSection(section); err != nil {
				return nil, err
			}
		}
	}

	c.data.Remove("Sections")

	return c, nil
}

func (c *chunk) GetBlock(x, y, z int32) Block {
	ys := y >> 4

	if c.sections[ys] == nil {
		return Block{}
	}

	b := c.sections[ys].GetBlock(x, y, z)
	pos := xyz(x, y, z)

	if md, ok := c.tileEntities[pos]; ok && md != nil {
		b.SetMetadata(md)
	}

	if tt, ok := c.tileTicks[pos]; ok && tt != nil {
		b.ticks = make([]Tick, len(tt))

		for n, tick := range tt {
			b.ticks[n] = Tick{
				int32(tick.Get("i").Data().(nbt.Int)),
				int32(tick.Get("t").Data().(nbt.Int)),
				int32(tick.Get("p").Data().(nbt.Int)),
			}
		}
	}

	return b
}

func (c *chunk) SetBlock(x, y, z int32, b Block) {
	ys := y >> 4

	if c.sections[ys] == nil {
		if b.EqualBlock(Block{}) {
			return
		}

		c.sections[ys] = newSection(y)
	}

	c.sections[ys].SetBlock(x, y, z, b)

	if hmpos := x&15<<4 | z&15; b.Opacity() <= 1 { // All transparent blocks block 1 light when they are below the highest non-transparent block
		if y == c.heightMap[hmpos]-1 {
			c.heightMap[hmpos] = 0

			for i := y; i >= 0; i-- {
				if yt := i >> 4; c.sections[yt] != nil {
					if c.sections[yt].GetOpacity(x, i, z) > 1 {
						c.heightMap[hmpos] = i + 1

						break
					}
				}
			}
		}
	} else if y >= c.heightMap[hmpos] {
		c.heightMap[hmpos] = y + 1
	}

	pos := xyz(x, y, z)

	if b.metadata == nil {
		delete(c.tileEntities, pos)
	} else {
		comp := b.GetMetadata()

		comp.Set(nbt.NewTag("x", nbt.Int(x)))
		comp.Set(nbt.NewTag("y", nbt.Int(y)))
		comp.Set(nbt.NewTag("z", nbt.Int(z)))

		c.tileEntities[pos] = comp
	}
	if b.HasTicks() {
		ticks := b.GetTicks()

		c.tileTicks[pos] = make([]nbt.Compound, len(ticks))

		for n, tick := range ticks {
			c.tileTicks[pos][n] = nbt.Compound{
				nbt.NewTag("i", nbt.Int(tick.I)),
				nbt.NewTag("p", nbt.Int(tick.P)),
				nbt.NewTag("t", nbt.Int(tick.T)),
				nbt.NewTag("x", nbt.Int(x)),
				nbt.NewTag("y", nbt.Int(y)),
				nbt.NewTag("z", nbt.Int(z)),
			}
		}
	} else {
		delete(c.tileTicks, pos)
	}
}

func (c *chunk) GetBiome(x, z int32) Biome {
	return Biome(c.biomes[x&15<<4|z&15])
}

func (c *chunk) SetBiome(x, z int32, b Biome) {
	c.biomes[x&15<<4|z&15] = int8(b)
}

func (c *chunk) GetOpacity(x, y, z int32) uint8 {
	if y >= c.heightMap[x&15<<4|z&15] {
		return 1
	}

	ys := y >> 4

	if c.sections[ys] == nil {
		return 1
	}

	return c.sections[ys].GetOpacity(x, y, z)
}

func (c *chunk) GetHeight(x, z int32) int32 {
	return c.heightMap[x&15<<4|z&15]
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
	} else if y >= c.heightMap[x&15<<4|z&15] || y > 255 {
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

func getCoord(name string, data nbt.Compound) (int32, error) {
	tag := data.Get(name)

	if tag.TagID() == 0 {
		return 0, MissingTagError{name}
	} else if tag.TagID() != nbt.TagInt {
		return 0, WrongTypeError{name, nbt.TagInt, tag.TagID()}
	}

	return int32(tag.Data().(nbt.Int)), nil
}

func getCoords(data nbt.Compound) (x, y, z int32, err error) {
	if x, err = getCoord("x", data); err != nil {
		return
	} else if y, err = getCoord("y", data); err != nil {
		return
	}

	z, err = getCoord("z", data)

	return
}
