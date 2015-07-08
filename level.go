package minecraft

import (
	"math/rand"
	"time"

	"github.com/MJKWoolnough/boolmap"
	"github.com/MJKWoolnough/minecraft/nbt"
)

var (
	levelRequired = map[string]nbt.TagID{
		"LevelName": nbt.TagString,
		"SpawnX":    nbt.TagInt,
		"SpawnY":    nbt.TagInt,
		"SpawnZ":    nbt.TagInt,
	}
)

// Level is the base type for minecraft data, all data for a minecraft level is
// either store in, or accessed from, this type
type Level struct {
	path      Path
	chunks    map[uint64]*chunk
	levelData nbt.Compound
	changed   bool
}

// NewLevel creates/Loads a minecraft level from the given path.
func NewLevel(location Path) (*Level, error) {
	var (
		levelDat nbt.Tag
		data     nbt.Compound
		changed  bool
	)
	levelDat, err := location.ReadLevelDat()
	if err != nil {
		return nil, err
	} else if levelDat.TagID() == 0 {
		levelDat = nbt.NewTag("", nbt.Compound{
			nbt.NewTag("Data", nbt.Compound{
				nbt.NewTag("version", nbt.Int(19133)),
				nbt.NewTag("initialized", nbt.Byte(0)),
				nbt.NewTag("LevelName", nbt.String("")),
				nbt.NewTag("generatorName", nbt.String(DefaultGenerator)),
				nbt.NewTag("generatorVersion", nbt.Int(0)),
				nbt.NewTag("generatorOptions", nbt.String("0")),
				nbt.NewTag("RandomSeed", nbt.Long(rand.New(rand.NewSource(time.Now().Unix())).Int63())),
				nbt.NewTag("MapFeatures", nbt.Byte(1)),
				nbt.NewTag("LastPlayed", nbt.Long(time.Now().Unix()*1000)),
				nbt.NewTag("SizeOnDisk", nbt.Long(0)),
				nbt.NewTag("allowCommands", nbt.Byte(0)),
				nbt.NewTag("hardcore", nbt.Byte(0)),
				nbt.NewTag("GameType", nbt.Int(Survival)),
				nbt.NewTag("Time", nbt.Long(0)),
				nbt.NewTag("DayTime", nbt.Long(0)),
				nbt.NewTag("SpawnX", nbt.Int(0)),
				nbt.NewTag("SpawnY", nbt.Int(0)),
				nbt.NewTag("SpawnZ", nbt.Int(0)),
				nbt.NewTag("BorderCenterX", nbt.Double(0)),
				nbt.NewTag("BorderCenterZ", nbt.Double(0)),
				nbt.NewTag("BorderSize", nbt.Double(60000000)),
				nbt.NewTag("BorderSafeZone", nbt.Double(5)),
				nbt.NewTag("BorderWarningTime", nbt.Double(15)),
				nbt.NewTag("BorderSizeLerpTarget", nbt.Double(60000000)),
				nbt.NewTag("BorderSizeLerpTime", nbt.Long(0)),
				nbt.NewTag("BorderDamagePerBlock", nbt.Double(0.2)),
				nbt.NewTag("raining", nbt.Byte(0)),
				nbt.NewTag("rainTime", nbt.Int(0)),
				nbt.NewTag("thundering", nbt.Byte(0)),
				nbt.NewTag("thunderTime", nbt.Int(0)),
				nbt.NewTag("clearWeatherTime", nbt.Int(0)),
				nbt.NewTag("GameRules", nbt.Compound{
					nbt.NewTag("commandBlockOutput", nbt.String("True")),
					nbt.NewTag("doDaylightCycle", nbt.String("True")),
					nbt.NewTag("doFireTick", nbt.String("True")),
					nbt.NewTag("doMobLoot", nbt.String("True")),
					nbt.NewTag("doMobSpawning", nbt.String("True")),
					nbt.NewTag("doTileDrops", nbt.String("True")),
					nbt.NewTag("keepInventory", nbt.String("False")),
					nbt.NewTag("logAdminCommands", nbt.String("True")),
					nbt.NewTag("mobGriefing", nbt.String("True")),
					nbt.NewTag("naturalRegeneration", nbt.String("True")),
					nbt.NewTag("randomTickSpeed", nbt.String("3")),
					nbt.NewTag("sendCommandFeedback", nbt.String("True")),
					nbt.NewTag("showDeathMessages", nbt.String("True")),
				}),
			}),
		})
		changed = true
	}
	if levelDat.TagID() != nbt.TagCompound {
		return nil, WrongTypeError{"[BASE]", nbt.TagCompound, levelDat.TagID()}
	} else if d := levelDat.Data().(nbt.Compound).Get("Data"); d.TagID() != 0 {
		if d.TagID() == nbt.TagCompound {
			data = d.Data().(nbt.Compound)
		} else {
			return nil, WrongTypeError{"Data", nbt.TagCompound, d.TagID()}
		}
	} else {
		return nil, MissingTagError{"Data"}
	}
	for name, tagType := range levelRequired {
		if x := data.Get(name); x.TagID() == 0 {
			return nil, MissingTagError{name}
		} else if x.TagID() != tagType {
			return nil, WrongTypeError{name, tagType, x.TagID()}
		}
	}
	return &Level{
		location,
		make(map[uint64]*chunk),
		levelDat.Data().(nbt.Compound).Get("Data").Data().(nbt.Compound),
		changed,
	}, nil
}

// GetBlock gets the block at coordinates x, y, z.
func (l *Level) GetBlock(x, y, z int32) (Block, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return Block{}, err
	} else if c == nil {
		return Block{}, nil
	}
	return c.GetBlock(x, y, z), nil
}

// SetBlock sets the block at coordinates x, y, z. Also processes any lighting updates if applicable.
func (l *Level) SetBlock(x, y, z int32, block Block) error {
	var (
		c   *chunk
		err error
	)
	for mx := x - 1; mx <= x+1; mx++ {
		for mz := z - 1; mz <= z+1; mz++ {
			if c, err = l.getChunk(mx, mz, true); err != nil {
				return err
			}
			for my := y + 16; my >= 0; my -= 16 {
				if c.createSection(my) {
					break
				}
			}
		}
	}
	c, _ = l.getChunk(x, z, false)
	opacity := c.GetOpacity(x, y, z)
	c.SetBlock(x, y, z, block)
	if block.Opacity() != opacity {
		if err = l.genLighting(x, y, z, true, block.Opacity() > opacity, 0); err != nil {
			return err
		}
	}
	if bl := c.GetBlockLight(x, y, z); block.Light() != bl || block.Opacity() != opacity {
		if err = l.genLighting(x, y, z, false, block.Light() < bl, block.Light()); err != nil {
			return err
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
	changed.SetBool((uint64(y)<<10)|(uint64(z&31)<<5)|uint64(x&31), true)
	if darker { // reset lighting on all blocks affected by the changed one (only applies if darker)
		setLight(c, x, y, z, 0)
		for i := 0; i < len(list); i++ {
			for _, s := range surroundingBlocks(list[i].x, list[i].y, list[i].z) {
				mx, my, mz := s[0], s[1], s[2]
				pos := (uint64(my) << 10) | (uint64(mz&31) << 5) | uint64(mx&31)
				if changed.GetBool(pos) {
					continue
				}
				if c, err = l.getChunk(mx, mz, false); err != nil {
					return err
				} else if c == nil {
					continue
				} else if ys := my >> 4; my < 16 && c.sections[ys] == nil {
					changed.SetBool(pos, true)
					continue
				}
				shouldBe := list[i].lightLevel
				opacity := c.GetOpacity(mx, my, mz)
				if opacity > shouldBe {
					shouldBe = 0
				} else {
					shouldBe -= opacity
				}
				if thisLight := getLight(c, mx, my, mz); thisLight == shouldBe && shouldBe != 0 || (skyLight && thisLight == 15 && my < c.GetHeight(mx, mz)) {
					list = append(list, &lightCoords{mx, my, mz, thisLight})
					changed.SetBool(pos, true)
					if thisLight > 0 {
						setLight(c, mx, my, mz, 0)
					}
				}
			}
		}
	} // end lighting reset
	if source > 0 { //If this is the source of light
		c, _ = l.getChunk(x, z, false)
		c.SetBlockLight(x, y, z, source)
		list = list[1:]
		for _, s := range surroundingBlocks(x, y, z) {
			mx, my, mz := s[0], s[1], s[2]
			pos := (uint64(my) << 10) | (uint64(mz&31) << 5) | uint64(mx&31)
			if changed.GetBool(pos) {
				continue
			}
			if c, err = l.getChunk(mx, mz, false); err != nil {
				return err
			} else if c == nil {
				continue
			} else if ys := my >> 4; my < 16 && c.sections[ys] == nil {
				changed.SetBool(pos, true)
				continue
			}
			if thisLight := getLight(c, mx, my, mz); thisLight < source {
				list = append(list, &lightCoords{mx, my, mz, thisLight})
				changed.SetBool(pos, true)
			}
		}
	}
	for ; len(list) > 0; list = list[1:] {
		mx, my, mz := list[0].x, list[0].y, list[0].z
		changed.SetBool((uint64(my)<<10)|(uint64(mz&31)<<5)|uint64(mx&31), false)
		newLight := uint8(0)
		c, _ = l.getChunk(mx, mz, false)
		if skyLight && my >= c.GetHeight(mx, mz) { //Determine correct light level...
			newLight = 15
		} else if opacity := c.GetOpacity(mx, my, mz); opacity == 15 {
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
		if newLight > list[0].lightLevel || (darker && newLight == list[0].lightLevel) {
			for _, s := range surroundingBlocks(mx, my, mz) {
				mx, my, mz = s[0], s[1], s[2]
				pos := (uint64(my) << 10) | (uint64(mz&31) << 5) | uint64(mx&31)
				if changed.GetBool(pos) {
					continue
				}
				if c, err = l.getChunk(mx, mz, false); err != nil {
					return err
				} else if c == nil {
					continue
				} else if ys := my >> 4; ys < 16 && c.sections[ys] == nil {
					changed.SetBool(pos, true)
					continue
				}
				if thisLight := getLight(c, mx, my, mz); thisLight < newLight {
					list = append(list, &lightCoords{mx, my, mz, thisLight})
					changed.SetBool(pos, true)
				}
			}
		}
	}
	return nil
}

// GetBiome returns the biome for the column x, z.
func (l *Level) GetBiome(x, z int32) (Biome, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return AutoBiome, err
	} else if c == nil {
		return Plains, nil
	}
	return c.GetBiome(x, z), nil
}

// SetBiome sets the biome for the column x, z.
func (l *Level) SetBiome(x, z int32, biome Biome) error {
	c, err := l.getChunk(x, z, true)
	if err != nil {
		return err
	}
	c.SetBiome(x, z, biome)
	return nil
}

// GetHeight returns the y coordinate for the highest non-transparent block at column x, z.
func (l *Level) GetHeight(x, z int32) (int32, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil || c == nil {
		return 0, err
	}
	return c.GetHeight(x, z), nil
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
		if chunkData.TagID() != 0 {
			chunk, err := newChunk(x, z, chunkData)
			if err != nil {
				return nil, err
			}
			l.chunks[pos] = chunk
		} else if create {
			l.chunks[pos], _ = newChunk(x, z, nbt.Tag{})
		}
	}
	return l.chunks[pos], nil
}

// Save saves all open chunks, but does not close them.
func (l *Level) Save() error {
	if l.changed {
		if err := l.path.WriteLevelDat(nbt.NewTag("", nbt.Compound{nbt.NewTag("Data", l.levelData)})); err != nil {
			return err
		}
		l.changed = false
	}
	var toSave []nbt.Tag
	for _, c := range l.chunks {
		toSave = append(toSave, c.GetNBT())
	}
	if len(toSave) > 0 {
		return l.path.SetChunk(toSave...) //check multi-error
	}
	return nil
}

// Close closes all open chunks, but does not save them.
func (l *Level) Close() {
	l.changed = false
	l.chunks = make(map[uint64]*chunk)
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
