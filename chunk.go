package minecraft

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbtparser"
	"io"
)

type Chunk interface {
	io.WriterTo
	Get(uint8, uint8, uint8) Block
	Set(uint8, uint8, uint8, Block)
	GetSkyLight(uint8, uint8, uint8) uint8
	SetSkyLight(uint8, uint8, uint8, uint8)
	Opacity(uint8, uint8, uint8) uint8
	IsEmpty() bool
	GetMetadata(uint8, uint8, uint8) []nbtparser.NBTTag
	SetMetadata(uint8, uint8, uint8, []nbtparser.NBTTag)
	GetTick(x, y, z uint8) bool
	SetTick(x, y, z uint8, data []nbtparser.NBTTag)
	HasChanged() bool
	Compress()
	// 	SkyUpdates() [256]bool
	// 	HighestBlock(uint8, uint8) uint8
}

type chunk struct {
	data            *bytes.Buffer
	root            *nbtparser.NBTTagCompound
	biomes          *nbtparser.NBTTagByteArray
	heightMap       *nbtparser.NBTTagIntArray
	sections        [16]Section
	sectionTag      *nbtparser.NBTTagList
	tileEntities    map[uint16]*nbtparser.NBTTagCompound
	tileEntitiesTag *nbtparser.NBTTagList
	//entities *nbtparser.NBTTagList
	//tileEntities *nbtparser.NBTTagList
	tileTicks map[uint16]*nbtparser.NBTTagCompound
	// 	updateSky       [256]bool
	changed      bool
	uncompressed bool
}

// func (c *chunk) HighestBlock(x, z uint8) uint8 {
// 	if !c.uncompressed {
// 		c.parseData()
// 	}
// 	return uint8(c.heightMap.Get(zx(x, z)))
// }

func (c *chunk) Compress() {
	if c.uncompressed {
		c.IsEmpty()
		c.data.Reset()
		c.WriteTo(c.data)
		c.uncompressed = false
		c.root = nil
		c.biomes = nil
		c.heightMap = nil
		for i := 0; i < 16; i++ {
			c.sections[i] = nil
		}
		c.sectionTag = nil
		c.tileEntities = nil
		c.tileEntitiesTag = nil
		c.tileTicks = nil
	}
}

func (c *chunk) Get(x, y, z uint8) Block {
	if c == nil {
		return BlockAir
	} else if !c.uncompressed {
		c.parseData()
	}
	section := c.sections[y>>4]
	if section == nil {
		return BlockAir
	}
	return section.Get(x, y&15, z)
}

func (c *chunk) Set(x, y, z uint8, block Block) {
	if block == nil || c == nil {
		return
	} else if !c.uncompressed {
		c.parseData()
	}
	if c.sections[y>>4] == nil {
		if BlockAir.Equal(block) {
			return
		}
		c.sections[y>>4], _ = NewSection(y >> 4)
	}
	sect := c.sections[y>>4]
	nY := y & 15
	oldOpac := sect.Opacity(x, nY, z)
	sect.Set(x, nY, z, block)
	if newOpac := sect.Opacity(x, nY, z); newOpac != oldOpac {
		// 		c.updateSky[zx(x, z)] = true
		if newOpac == 0 {
			if uint8(c.heightMap.Get(zx(x, z))) == y+1 {
				c.heightMap.Set(zx(x, z), 0)
				for i := y - 1; i >= 0; i++ {
					if c.Opacity(x, i, z) != 0 {
						c.heightMap.Set(zx(x, z), int32(i)+1)
						break
					}
				}
			}
		} else if uint8(c.heightMap.Get(zx(x, z))) <= y {
			c.heightMap.Set(zx(x, z), int32(y)+1)
		}
	}
}

func (c *chunk) GetSkyLight(x, y, z uint8) uint8 {
	if c == nil {
		return 15
	} else if !c.uncompressed {
		c.parseData()
	}
	section := c.sections[y>>4]
	if section == nil {
		return 15
	}
	return section.GetSkyLight(x, y&15, z)
}

func (c *chunk) SetSkyLight(x, y, z, skylight uint8) {
	if c == nil {
		return
	} else if !c.uncompressed {
		c.parseData()
	}
	section := c.sections[y>>4]
	if section == nil {
		return
	}
	c.changed = true
	section.SetSkyLight(x, y&15, z, skylight)
}

func (c *chunk) Opacity(x, y, z uint8) uint8 {
	if c == nil {
		return 0
	} else if !c.uncompressed {
		c.parseData()
	}
	section := c.sections[y>>4]
	if section == nil {
		return 0
	}
	return section.Opacity(x, y&15, z)
}

func (c *chunk) GetMetadata(x, y, z uint8) []nbtparser.NBTTag {
	if !c.uncompressed {
		c.parseData()
	}
	xu := uint16(x & 15)
	zu := uint16(z & 15)
	yu := uint16(y & 255)
	if d := c.tileEntities[yu<<8|zu<<4|xu]; d != nil {
		return d.GetArray()
	}
	return nil
}

func (c *chunk) SetMetadata(x, y, z uint8, data []nbtparser.NBTTag) {
	if !c.uncompressed {
		c.parseData()
	}
	xu := uint16(x & 15)
	zu := uint16(z & 15)
	yu := uint16(y & 255)
	c.changed = true
	if data == nil {
		delete(c.tileEntities, yu<<8|zu<<4|xu)
	} else if c.tileEntities[yu<<8|zu<<4|xu] == nil {
		c.tileEntities[yu<<8|zu<<4|xu] = nbtparser.NewTagCompound("", data)
	} else {
		c.tileEntities[yu<<8|zu<<4|xu].SetArray(data)
	}
}

func (c *chunk) GetTick(x, y, z uint8) bool {
	return c.tileTicks[uint16(y&255)<<8|uint16(z&15)<<4|uint16(x&15)] != nil
}

func (c *chunk) SetTick(x, y, z uint8, data []nbtparser.NBTTag) {
	xu := uint16(x & 15)
	zu := uint16(z & 15)
	yu := uint16(y & 255)
	c.changed = true
	if data == nil {
		delete(c.tileTicks, yu<<8|zu<<4|xu)
	} else if c.tileTicks[yu<<8|zu<<4|xu] == nil {
		c.tileTicks[yu<<8|zu<<4|xu] = nbtparser.NewTagCompound("", data)
	} else {
		c.tileTicks[yu<<8|zu<<4|xu].SetArray(data)
	}
}

// func (c *chunk) SkyUpdates() [256]bool {
// 	c.basicSkyUpdate()
// 	return c.updateSky
// }
// 
// func (c *chunk) basicSkyUpdate() {
// 	if !c.uncompressed {
// 		c.parseData()
// 	}
// 	for i := uint8(0); i < 16; i++ {
// 		for j := uint8(0); j < 16; j++ {
// 			if c.updateSky[zx(i, j)] {
// 				ll := uint8(15)
// 				for k := int8(15); k >= 0; k-- {
// 					if sect := c.sections[k]; sect != nil {
// 						for l := int8(15); l >= 0; l -- {
// 							if ll > 0 {
// 								o := sect.Opacity(i, uint8(k) << 4 + uint8(l), j)
// 								if o > ll {
// 									ll = 0
// 								} else {
// 									ll -= o
// 								}
// 							}
// 							sect.SetSkyLight(i, uint8(k) << 4 + uint8(l), j, ll)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

func (c *chunk) IsEmpty() bool {
	if !c.uncompressed {
		return false
	} else if c == nil {
		return true
	}
	sections := make([]interface{}, 0)

	for i := 0; i < 16; i++ {
		if c.sections[i] != nil && !c.sections[i].IsEmpty() {
			sections = append(sections, c.sections[i].Data())
		} else {
			c.sections[i] = nil
		}
	}
	c.sectionTag.SetArray(nbtparser.NBTTag_Compound, sections)
	if len(sections) != 0 {
		tEntities := make([]interface{}, 0)
		for _, d := range c.tileEntities {
			tEntities = append(tEntities, d)
		}
		if len(tEntities) == 0 {
			c.tileEntitiesTag.SetArray(nbtparser.NBTTag_Byte, tEntities)
		} else {
			c.tileEntitiesTag.SetArray(nbtparser.NBTTag_Compound, tEntities)
		}
		tTicks := make([]interface{}, 0)
		for _, d := range c.tileTicks {
			tTicks = append(tTicks, d)
		}
		if ticksTag := c.root.GetTag("TileTicks"); ticksTag == nil {
			if len(tTicks) != 0 {
				c.root.Append(nbtparser.NewTagList("TileTicks", nbtparser.NBTTag_Compound, tTicks))
			}
		} else if len(tTicks) != 0 {
			if tt := ticksTag.TagList(); tt != nil {
				tt.SetArray(nbtparser.NBTTag_Compound, tTicks)
			}
		} else {
			c.root.RemoveTag("TileTicks")
		}
		return false
	}
	return true
}

func (c *chunk) WriteTo(file io.Writer) (int64, error) {
	if c.uncompressed {
		z := zlib.NewWriter(file)
		defer z.Close()
		return nbtparser.NewTagCompound("", []nbtparser.NBTTag{c.root}).WriteTo(z)
	}
	return c.data.WriteTo(file)
}

func (c *chunk) HasChanged() bool {
	if c.changed {
		return true
	}
	return false
}

func (c *chunk) parseData() error {
	if c.uncompressed {
		return nil
	} else if c.data.Len() == 0 {
		return fmt.Errorf("Minecraft - Chunk: no data")
	} else if z, err := zlib.NewReader(c.data); err != nil {
		return err
	} else if nbtFile, _, err := nbtparser.ParseFile(z); err != nil {
		return err
	} else {
		z.Close()
		c.tileEntities = make(map[uint16]*nbtparser.NBTTagCompound)
		if level := nbtFile.GetTag("Level"); level == nil {
			return fmt.Errorf("Minecraft - Chunk: Missing 'Level' tag")
		} else if levelCmpd := level.TagCompound(); levelCmpd == nil {
			return fmt.Errorf("Minecraft - Chunk: Level tag of wrong type")
		} else {
			c.root = levelCmpd
		}
		if xPos := c.root.GetTag("xPos"); xPos == nil {
			return fmt.Errorf("Minecraft - Chunk: Missing 'xPos' tag")
		} else if xPos.TagInt() == nil {
			return fmt.Errorf("Minecraft - Chunk: xPos tag of wrong type")
		}
		if zPos := c.root.GetTag("zPos"); zPos == nil {
			return fmt.Errorf("Minecraft - Chunk: Missing 'zPos' tag")
		} else if zPos.TagInt() == nil {
			return fmt.Errorf("Minecraft - Chunk: zPos tag of wrong type")
		}
		if lUpdate := c.root.GetTag("LastUpdate"); lUpdate == nil {
			return fmt.Errorf("Minecraft - Chunk: Missing 'LastUpdate' tag")
		} else if lUpdate.TagLong() == nil {
			return fmt.Errorf("Minecraft - Chunk: LastUpdate tag of wrong type")
		}
		if biomes := c.root.GetTag("Biomes"); biomes != nil {
			if bios := biomes.TagByteArray(); bios == nil {
				return fmt.Errorf("Minecraft - Chunk: Biomes tag of wrong type")
			} else {
				c.biomes = bios
			}
		}
		if hMap := c.root.GetTag("HeightMap"); hMap == nil {
			return fmt.Errorf("Minecraft - Chunk: Missing 'HeightMap' tag")
		} else if hMA := hMap.TagIntArray(); hMA == nil {
			return fmt.Errorf("Minecraft - Chunk: HeightMap tag of wrong type")
		} else {
			c.heightMap = hMA
		}
		if sections := c.root.GetTag("Sections"); sections == nil {
			return fmt.Errorf("Minecraft - Chunk: Missing 'Sections' tag")
		} else if sectionList := sections.TagList(); sectionList == nil {
			return fmt.Errorf("Minecraft - Chunk: Sections tag of wrong type")
		} else if sectionList.GetType() != nbtparser.NBTTag_Compound {
			return fmt.Errorf("Minecraft - Chunk: Sections tag of wrong list type")
		} else {
			for i := int32(0); i < sectionList.Length(); i++ {
				if sectionCmp := sectionList.Get(i).(*nbtparser.NBTTagCompound); sectionCmp == nil {
					return fmt.Errorf("Minecraft - Chunk: Sections tag contains invalid list item")
				} else if section, err := LoadSection(sectionCmp); err != nil {
					return err
				} else {
					y := section.GetY()
					c.sections[y] = section
				}
			}
			c.sectionTag = sectionList
		}
		if entities := c.root.GetTag("Entities"); entities == nil {
			return fmt.Errorf("Minecraft - Chunk: Missing 'Entities' tag")
		} else if entitiesList := entities.TagList(); entitiesList == nil {
			return fmt.Errorf("Minecraft - Chunk: Entities tag of wrong type")
		} else if entitiesList.GetType() != nbtparser.NBTTag_Compound {
			if entitiesList.GetType() != nbtparser.NBTTag_Byte {
				return fmt.Errorf("Minecraft - Chunk: Entities tag of wrong list type")
			}
		}
		if tEntities := c.root.GetTag("TileEntities"); tEntities == nil {
			return fmt.Errorf("Minecraft - Chunk: Missing 'TileEntities' tag")
		} else if tEntitiesList := tEntities.TagList(); tEntitiesList == nil {
			return fmt.Errorf("Minecraft - Chunk: TileEntities tag of wrong type")
		} else if tEntitiesList.GetType() != nbtparser.NBTTag_Compound {
			if tEntitiesList.GetType() != nbtparser.NBTTag_Byte {
				return fmt.Errorf("Minecraft - Chunk: TileEntities tag of wrong list type")
			} else {
				c.tileEntitiesTag = tEntitiesList
			}
		} else {
			c.tileEntitiesTag = tEntitiesList
			for i := int32(0); i < tEntitiesList.Length(); i++ {
				if cmpd, ok := tEntitiesList.Get(i).(*nbtparser.NBTTagCompound); ok {
					if x, y, z, err := getTECoords(cmpd); err != nil {
						return err
					} else {
						xu := uint16(x & 15)
						zu := uint16(z & 15)
						yu := uint16(y & 255)
						c.tileEntities[yu<<8|zu<<4|xu] = cmpd
					}
				}
			}
		}
		c.tileTicks = make(map[uint16]*nbtparser.NBTTagCompound)
		if tileTicks := c.root.GetTag("TileTicks"); tileTicks != nil {
			if tileTicksList := tileTicks.TagList(); tileTicksList == nil {
				return fmt.Errorf("Minecraft - Chunk: TileTicks tag of wrong type")
			} else if tileTicksList.GetType() != nbtparser.NBTTag_Compound {
				return fmt.Errorf("Minecraft - Chunk: TileTicks tag of wrong list type")
			} else {
				for i := int32(0); i < tileTicksList.Length(); i++ {
					if cmpd, ok := tileTicksList.Get(i).(*nbtparser.NBTTagCompound); ok {
						if x, y, z, err := getTECoords(cmpd); err != nil {
							return err
						} else {
							xu := uint16(x & 15)
							zu := uint16(z & 15)
							yu := uint16(y & 255)
							c.tileTicks[yu<<8|zu<<4|xu] = cmpd
						}
					}
				}
			}
		}
	}
	c.uncompressed = true
	return nil
}

func LoadChunk(data io.Reader) (Chunk, error) {
	c := new(chunk)
	c.data = new(bytes.Buffer)
	_, err := c.data.ReadFrom(data)
	return c, err
}

func NewChunk(x, z int32) Chunk {
	biomes := make([]byte, 256)
	height := make([]int32, 256)
	for i := 0; i < 256; i++ {
		biomes[i] = 1
		height[i] = 0
	}
	c := new(chunk)
	c.data = new(bytes.Buffer)
	zFile := zlib.NewWriter(c.data)
	nbtparser.NewTagCompound("", []nbtparser.NBTTag{nbtparser.NewTagCompound("Level", []nbtparser.NBTTag{
		nbtparser.NewTagList("Entities", nbtparser.NBTTag_Byte, make([]interface{}, 0)),
		nbtparser.NewTagByteArray("Biomes", biomes),
		nbtparser.NewTagLong("LastUpdate", 0),
		nbtparser.NewTagInt("xPos", x),
		nbtparser.NewTagInt("zPos", z),
		nbtparser.NewTagList("TileEntities", nbtparser.NBTTag_Byte, make([]interface{}, 0)),
		nbtparser.NewTagByte("TerrainPopulated", 1),
		nbtparser.NewTagIntArray("HeightMap", height),
		nbtparser.NewTagList("Sections", nbtparser.NBTTag_Compound, make([]interface{}, 0)),
	})}).WriteTo(zFile)
	zFile.Close()
	c.parseData()
	return c
}

func zx(x, z uint8) int32 {
	return int32(((z & 15) << 4) | (x & 15))
}

func getTECoord(coord string, cmpd *nbtparser.NBTTagCompound) (int32, error) {
	if c := cmpd.GetTag(coord); c != nil {
		if d := c.TagInt(); d != nil {
			return d.Get(), nil
		}
		return 0, fmt.Errorf("Minecraft - Chunk - TileEntities: " + coord + " tag of wrong type")
	}
	return 0, fmt.Errorf("Minecraft - Chunk - TileEntities: Missing '" + coord + "' tag")
}

func getTECoords(cmpd *nbtparser.NBTTagCompound) (int32, int32, int32, error) {
	if x, err := getTECoord("x", cmpd); err != nil {
		return 0, 0, 0, err
	} else if y, err := getTECoord("y", cmpd); err != nil {
		return x, 0, 0, err
	} else if z, err := getTECoord("z", cmpd); err != nil {
		return x, y, 0, err
	} else {
		return x, y, z, nil
	}
	return 0, 0, 0, nil
}
