package minecraft

import (
	"github.com/MJKWoolnough/minecraft/nbt"
	"testing"
)

func TestNew(t *testing.T) {
	biomes := make([]int8, 256)
	biome := int8(-1)
	blocks := make([]int8, 4096)
	add := make([]int8, 2048)
	data := make([]int8, 2048)
	for i := 0; i < 256; i++ {
		biomes[i] = biome
		if biome++; biome >= 23 {
			biome = -1
		}
	}
	dataTag := nbt.NewTag("", nbt.NewCompound([]nbt.Tag{
		nbt.NewTag("Level", nbt.NewCompound([]nbt.Tag{
			nbt.NewTag("Biomes", nbt.NewByteArray(biomes)),
			nbt.NewTag("HeightMap", nbt.NewIntArray(make([]int32, 256))),
			nbt.NewTag("InhabitedTime", nbt.NewLong(0)),
			nbt.NewTag("LastUpdate", nbt.NewLong(0)),
			nbt.NewTag("Sections", nbt.NewList([]nbt.Data{
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("Blocks", nbt.NewByteArray(blocks)),
					nbt.NewTag("Add", nbt.NewByteArray(add)),
					nbt.NewTag("Data", nbt.NewByteArray(data)),
					nbt.NewTag("BlockLight", nbt.NewByteArray(make([]int8, 2048))),
					nbt.NewTag("SkyLight", nbt.NewByteArray(make([]int8, 2048))),
					nbt.NewTag("Y", nbt.NewByte(0)),
				}),
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("Blocks", nbt.NewByteArray(blocks)),
					nbt.NewTag("Add", nbt.NewByteArray(add)),
					nbt.NewTag("Data", nbt.NewByteArray(data)),
					nbt.NewTag("BlockLight", nbt.NewByteArray(make([]int8, 2048))),
					nbt.NewTag("SkyLight", nbt.NewByteArray(make([]int8, 2048))),
					nbt.NewTag("Y", nbt.NewByte(1)),
				}),
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("Blocks", nbt.NewByteArray(blocks)),
					nbt.NewTag("Add", nbt.NewByteArray(add)),
					nbt.NewTag("Data", nbt.NewByteArray(data)),
					nbt.NewTag("BlockLight", nbt.NewByteArray(make([]int8, 2048))),
					nbt.NewTag("SkyLight", nbt.NewByteArray(make([]int8, 2048))),
					nbt.NewTag("Y", nbt.NewByte(3)),
				}),
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("Blocks", nbt.NewByteArray(blocks)),
					nbt.NewTag("Add", nbt.NewByteArray(add)),
					nbt.NewTag("Data", nbt.NewByteArray(data)),
					nbt.NewTag("BlockLight", nbt.NewByteArray(make([]int8, 2048))),
					nbt.NewTag("SkyLight", nbt.NewByteArray(make([]int8, 2048))),
					nbt.NewTag("Y", nbt.NewByte(10)),
				}),
			})),
			nbt.NewTag("TileEntities", nbt.NewList([]nbt.Data{
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("id", nbt.NewString("test1")),
					nbt.NewTag("x", nbt.NewInt(-191)),
					nbt.NewTag("y", nbt.NewInt(13)),
					nbt.NewTag("z", nbt.NewInt(379)),
					nbt.NewTag("testTag", nbt.NewByte(1)),
				}),
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("id", nbt.NewString("test2")),
					nbt.NewTag("x", nbt.NewInt(-191)),
					nbt.NewTag("y", nbt.NewInt(17)),
					nbt.NewTag("z", nbt.NewInt(372)),
					nbt.NewTag("testTag", nbt.NewLong(8)),
				}),
			})),
			nbt.NewTag("Entities", nbt.NewList([]nbt.Data{
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("id", nbt.NewString("testEntity1")),
					nbt.NewTag("Pos", nbt.NewList([]nbt.Data{
						nbt.NewDouble(-190),
						nbt.NewDouble(13),
						nbt.NewDouble(375),
					})),
					nbt.NewTag("Motion", nbt.NewList([]nbt.Data{
						nbt.NewDouble(1),
						nbt.NewDouble(13),
						nbt.NewDouble(11),
					})),
					nbt.NewTag("Rotation", nbt.NewList([]nbt.Data{
						nbt.NewFloat(13),
						nbt.NewFloat(11),
					})),
					nbt.NewTag("FallDistance", nbt.NewFloat(0)),
					nbt.NewTag("Fire", nbt.NewShort(-1)),
					nbt.NewTag("Air", nbt.NewShort(300)),
					nbt.NewTag("OnGround", nbt.NewByte(1)),
					nbt.NewTag("Dimension", nbt.NewInt(0)),
					nbt.NewTag("Invulnerable", nbt.NewByte(0)),
					nbt.NewTag("PortalCooldown", nbt.NewInt(0)),
					nbt.NewTag("UUIDMost", nbt.NewLong(0)),
					nbt.NewTag("UUIDLease", nbt.NewLong(0)),
					nbt.NewTag("Riding", nbt.NewCompound([]nbt.Tag{})),
				}),
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("id", nbt.NewString("testEntity2")),
					nbt.NewTag("Pos", nbt.NewList([]nbt.Data{
						nbt.NewDouble(-186),
						nbt.NewDouble(2),
						nbt.NewDouble(378),
					})),
					nbt.NewTag("Motion", nbt.NewList([]nbt.Data{
						nbt.NewDouble(17.5),
						nbt.NewDouble(1000),
						nbt.NewDouble(54),
					})),
					nbt.NewTag("Rotation", nbt.NewList([]nbt.Data{
						nbt.NewFloat(11),
						nbt.NewFloat(13),
					})),
					nbt.NewTag("FallDistance", nbt.NewFloat(30)),
					nbt.NewTag("Fire", nbt.NewShort(4)),
					nbt.NewTag("Air", nbt.NewShort(30)),
					nbt.NewTag("OnGround", nbt.NewByte(0)),
					nbt.NewTag("Dimension", nbt.NewInt(0)),
					nbt.NewTag("Invulnerable", nbt.NewByte(1)),
					nbt.NewTag("PortalCooldown", nbt.NewInt(10)),
					nbt.NewTag("UUIDMost", nbt.NewLong(1450)),
					nbt.NewTag("UUIDLease", nbt.NewLong(6435)),
					nbt.NewTag("Riding", nbt.NewCompound([]nbt.Tag{})),
				}),
			})),
			nbt.NewTag("TileTicks", nbt.NewList([]nbt.Data{
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("i", nbt.NewInt(0)),
					nbt.NewTag("t", nbt.NewInt(0)),
					nbt.NewTag("p", nbt.NewInt(0)),
					nbt.NewTag("x", nbt.NewInt(-192)),
					nbt.NewTag("y", nbt.NewInt(0)),
					nbt.NewTag("z", nbt.NewInt(368)),
				}),
				nbt.NewCompound([]nbt.Tag{
					nbt.NewTag("i", nbt.NewInt(1)),
					nbt.NewTag("t", nbt.NewInt(34)),
					nbt.NewTag("p", nbt.NewInt(12)),
					nbt.NewTag("x", nbt.NewInt(-186)),
					nbt.NewTag("y", nbt.NewInt(11)),
					nbt.NewTag("z", nbt.NewInt(381)),
				}),
			})),
			nbt.NewTag("TerrainPopulated", nbt.NewByte(1)),
			nbt.NewTag("xPos", nbt.NewInt(-12)),
			nbt.NewTag("zPos", nbt.NewInt(23)),
		})),
	}))
	if _, err := newChunk(-12, 23, dataTag); err != nil {
		t.Fatalf("reveived unexpected error during testing, %q", err.Error())
	}
}

func TestBiomes(t *testing.T) {
	chunk, _ := newChunk(0, 0, nil)
	for b := Biome(-1); b < 23; b++ {
		biome := b
		for x := int32(0); x < 16; x++ {
			for z := int32(0); z < 16; z++ {
				chunk.SetBiome(x, z, biome)
				if newB := chunk.GetBiome(x, z); newB != biome {
					t.Errorf("error setting biome at co-ordinates, expecting %q, got %q", biome.String(), newB.String())
				} else if biome++; biome >= 23 {
					biome = -1
				}
			}
		}
	}
}

func TestBlock(t *testing.T) {
	chunk, _ := newChunk(0, 0, nil)
	testBlocks := []struct {
		Block
		x, y, z int32
		recheck bool
	}{
		//Test simple set
		{
			Block{
				BlockId: 12,
			},
			0, 0, 0,
			true,
		},
		//Test higher ids
		{
			Block{
				BlockId: 853,
			},
			1, 0, 0,
			true,
		},
		{
			Block{
				BlockId: 463,
			},
			2, 0, 0,
			true,
		},
		{
			Block{
				BlockId: 1001,
			},
			3, 0, 0,
			true,
		},
		//Test data set
		{
			Block{
				BlockId: 143,
				Data:    12,
			},
			0, 1, 0,
			true,
		},
		{
			Block{
				BlockId: 153,
				Data:    4,
			},
			1, 1, 0,
			true,
		},
		{
			Block{
				BlockId: 163,
				Data:    5,
			},
			2, 1, 0,
			true,
		},
		//Test metadata [un]set
		{
			Block{
				metadata: []nbt.Tag{
					nbt.NewTag("testInt2", nbt.NewInt(1743)),
					nbt.NewTag("testString2", nbt.NewString("world")),
				},
			},
			0, 0, 1,
			true,
		},
		{
			Block{
				metadata: []nbt.Tag{
					nbt.NewTag("testInt", nbt.NewInt(15)),
					nbt.NewTag("testString", nbt.NewString("hello")),
				},
			},
			1, 0, 1,
			false,
		},
		{
			Block{},
			1, 0, 1,
			true,
		},
		//Test tick [un]set
		{
			Block{
				ticks: []Tick{{123, 1, 4}, {123, 7, -1}},
			},
			0, 1, 1,
			true,
		},
		{
			Block{
				ticks: []Tick{{654, 4, 6}, {4, 63, 5}, {4, 5, 9}},
			},
			1, 1, 1,
			false,
		},
		{
			Block{},
			1, 1, 1,
			true,
		},
	}
	for _, tB := range testBlocks {
		chunk.SetBlock(tB.x, tB.y, tB.z, &tB.Block)
		if block := chunk.GetBlock(tB.x, tB.y, tB.z); !tB.Block.Equal(block) {
			t.Errorf("blocks do not match, expecting %s, got %s", tB.Block.String(), block.String())
		}
	}
	for _, tB := range testBlocks {
		if tB.recheck {
			if block := chunk.GetBlock(tB.x, tB.y, tB.z); !tB.Block.Equal(block) {
				t.Errorf("blocks do not match, expecting:-\n%s\ngot:-\n%s", tB.Block.String(), block.String())
			}
		}
	}
}

func TestHeightMap(t *testing.T) {
	tests := []struct {
		x, y, z int32
		*Block
		height int32
	}{
		{0, 0, 0, &Block{}, 0},
		{1, 0, 0, &Block{BlockId: 1}, 1},
		{1, 1, 0, &Block{BlockId: 1}, 2},
		{1, 0, 0, &Block{}, 2},
		{1, 1, 0, &Block{}, 0},
		{2, 10, 0, &Block{BlockId: 1}, 11},
		{2, 12, 0, &Block{BlockId: 1}, 13},
		{2, 12, 0, &Block{}, 11},
		{2, 10, 0, &Block{}, 0},
		{3, 15, 0, &Block{BlockId: 1}, 16},
		{3, 16, 0, &Block{BlockId: 1}, 17},
		{3, 16, 0, &Block{}, 16},
		{3, 15, 0, &Block{}, 0},
		{4, 31, 0, &Block{BlockId: 1}, 32},
		{4, 32, 0, &Block{BlockId: 1}, 33},
		{4, 32, 0, &Block{}, 32},
		{4, 31, 0, &Block{}, 0},
		{5, 16, 0, &Block{BlockId: 1}, 17},
		{5, 32, 0, &Block{BlockId: 1}, 33},
		{5, 32, 0, &Block{}, 17},
		{5, 16, 0, &Block{}, 0},
	}
	chunk, _ := newChunk(0, 0, nil)
	for n, test := range tests {
		chunk.SetBlock(test.x, test.y, test.z, test.Block)
		if h := chunk.GetHeight(test.x, test.z); h != test.height {
			t.Errorf("test %d: expecting height %d, got %d", n+1, test.height, h)
		}
	}
}
