package minecraft

import (
	"testing"

	"vimagination.zapto.org/minecraft/nbt"
)

func TestNew(t *testing.T) {
	biomes := make(nbt.ByteArray, 256)
	biome := int8(-1)
	blocks := make(nbt.ByteArray, 4096)
	add := make(nbt.ByteArray, 2048)
	data := make(nbt.ByteArray, 2048)
	for i := 0; i < 256; i++ {
		biomes[i] = biome
		//if biome++; biome >= 23 {
		//	biome = -1
		//}
	}
	dataTag := nbt.NewTag("", nbt.Compound{
		nbt.NewTag("Level", nbt.Compound{
			nbt.NewTag("Biomes", biomes),
			nbt.NewTag("HeightMap", make(nbt.IntArray, 256)),
			nbt.NewTag("InhabitedTime", nbt.Long(0)),
			nbt.NewTag("LastUpdate", nbt.Long(0)),
			nbt.NewTag("Sections", &nbt.ListCompound{
				nbt.Compound{
					nbt.NewTag("Blocks", blocks),
					nbt.NewTag("Add", add),
					nbt.NewTag("Data", data),
					nbt.NewTag("BlockLight", make(nbt.ByteArray, 2048)),
					nbt.NewTag("SkyLight", make(nbt.ByteArray, 2048)),
					nbt.NewTag("Y", nbt.Byte(0)),
				},
				nbt.Compound{
					nbt.NewTag("Blocks", blocks),
					nbt.NewTag("Add", add),
					nbt.NewTag("Data", data),
					nbt.NewTag("BlockLight", make(nbt.ByteArray, 2048)),
					nbt.NewTag("SkyLight", make(nbt.ByteArray, 2048)),
					nbt.NewTag("Y", nbt.Byte(1)),
				},
				nbt.Compound{
					nbt.NewTag("Blocks", blocks),
					nbt.NewTag("Add", add),
					nbt.NewTag("Data", data),
					nbt.NewTag("BlockLight", make(nbt.ByteArray, 2048)),
					nbt.NewTag("SkyLight", make(nbt.ByteArray, 2048)),
					nbt.NewTag("Y", nbt.Byte(3)),
				},
				nbt.Compound{
					nbt.NewTag("Blocks", blocks),
					nbt.NewTag("Add", add),
					nbt.NewTag("Data", data),
					nbt.NewTag("BlockLight", make(nbt.ByteArray, 2048)),
					nbt.NewTag("SkyLight", make(nbt.ByteArray, 2048)),
					nbt.NewTag("Y", nbt.Byte(10)),
				},
			}),
			nbt.NewTag("TileEntities", &nbt.ListCompound{
				nbt.Compound{
					nbt.NewTag("id", nbt.String("test1")),
					nbt.NewTag("x", nbt.Int(-191)),
					nbt.NewTag("y", nbt.Int(13)),
					nbt.NewTag("z", nbt.Int(379)),
					nbt.NewTag("testTag", nbt.Byte(1)),
				},
				nbt.Compound{
					nbt.NewTag("id", nbt.String("test2")),
					nbt.NewTag("x", nbt.Int(-191)),
					nbt.NewTag("y", nbt.Int(17)),
					nbt.NewTag("z", nbt.Int(372)),
					nbt.NewTag("testTag", nbt.Long(8)),
				},
			}),
			nbt.NewTag("Entities", &nbt.ListCompound{
				nbt.Compound{
					nbt.NewTag("id", nbt.String("testEntity1")),
					nbt.NewTag("Pos", &nbt.ListDouble{
						nbt.Double(-190),
						nbt.Double(13),
						nbt.Double(375),
					}),
					nbt.NewTag("Motion", &nbt.ListDouble{
						nbt.Double(1),
						nbt.Double(13),
						nbt.Double(11),
					}),
					nbt.NewTag("Rotation", &nbt.ListFloat{
						nbt.Float(13),
						nbt.Float(11),
					}),
					nbt.NewTag("FallDistance", nbt.Float(0)),
					nbt.NewTag("Fire", nbt.Short(-1)),
					nbt.NewTag("Air", nbt.Short(300)),
					nbt.NewTag("OnGround", nbt.Byte(1)),
					nbt.NewTag("Dimension", nbt.Int(0)),
					nbt.NewTag("Invulnerable", nbt.Byte(0)),
					nbt.NewTag("PortalCooldown", nbt.Int(0)),
					nbt.NewTag("UUIDMost", nbt.Long(0)),
					nbt.NewTag("UUIDLease", nbt.Long(0)),
					nbt.NewTag("Riding", nbt.Compound{}),
				},
				nbt.Compound{
					nbt.NewTag("id", nbt.String("testEntity2")),
					nbt.NewTag("Pos", &nbt.ListDouble{
						nbt.Double(-186),
						nbt.Double(2),
						nbt.Double(378),
					}),
					nbt.NewTag("Motion", &nbt.ListDouble{
						nbt.Double(17.5),
						nbt.Double(1000),
						nbt.Double(54),
					}),
					nbt.NewTag("Rotation", &nbt.ListFloat{
						nbt.Float(11),
						nbt.Float(13),
					}),
					nbt.NewTag("FallDistance", nbt.Float(30)),
					nbt.NewTag("Fire", nbt.Short(4)),
					nbt.NewTag("Air", nbt.Short(30)),
					nbt.NewTag("OnGround", nbt.Byte(0)),
					nbt.NewTag("Dimension", nbt.Int(0)),
					nbt.NewTag("Invulnerable", nbt.Byte(1)),
					nbt.NewTag("PortalCooldown", nbt.Int(10)),
					nbt.NewTag("UUIDMost", nbt.Long(1450)),
					nbt.NewTag("UUIDLease", nbt.Long(6435)),
					nbt.NewTag("Riding", nbt.Compound{}),
				},
			}),
			nbt.NewTag("TileTicks", &nbt.ListCompound{
				nbt.Compound{
					nbt.NewTag("i", nbt.Int(0)),
					nbt.NewTag("t", nbt.Int(0)),
					nbt.NewTag("p", nbt.Int(0)),
					nbt.NewTag("x", nbt.Int(-192)),
					nbt.NewTag("y", nbt.Int(0)),
					nbt.NewTag("z", nbt.Int(368)),
				},
				nbt.Compound{
					nbt.NewTag("i", nbt.Int(1)),
					nbt.NewTag("t", nbt.Int(34)),
					nbt.NewTag("p", nbt.Int(12)),
					nbt.NewTag("x", nbt.Int(-186)),
					nbt.NewTag("y", nbt.Int(11)),
					nbt.NewTag("z", nbt.Int(381)),
				},
			}),
			nbt.NewTag("TerrainPopulated", nbt.Byte(1)),
			nbt.NewTag("xPos", nbt.Int(-12)),
			nbt.NewTag("zPos", nbt.Int(23)),
		}),
	})
	if _, err := newChunk(-12, 23, dataTag); err != nil {
		t.Fatalf("reveived unexpected error during testing, %q", err.Error())
	}
}

func TestBiomes(t *testing.T) {
	chunk, _ := newChunk(0, 0, nbt.Tag{})
	for b := Biome(0); b < 23; b++ {
		biome := b
		for x := int32(0); x < 16; x++ {
			for z := int32(0); z < 16; z++ {
				chunk.SetBiome(x, z, biome)
				if newB := chunk.GetBiome(x, z); newB != biome {
					t.Errorf("error setting biome at co-ordinates, expecting %q, got %q", biome.String(), newB.String())
				}
			}
		}
	}
}

func TestBlock(t *testing.T) {
	chunk, _ := newChunk(0, 0, nbt.Tag{})
	testBlocks := []struct {
		Block
		x, y, z int32
		recheck bool
	}{
		//Test simple set
		{
			Block{
				ID: 12,
			},
			0, 0, 0,
			true,
		},
		//Test higher ids
		{
			Block{
				ID: 853,
			},
			1, 0, 0,
			true,
		},
		{
			Block{
				ID: 463,
			},
			2, 0, 0,
			true,
		},
		{
			Block{
				ID: 1001,
			},
			3, 0, 0,
			true,
		},
		//Test data set
		{
			Block{
				ID:   143,
				Data: 12,
			},
			0, 1, 0,
			true,
		},
		{
			Block{
				ID:   153,
				Data: 4,
			},
			1, 1, 0,
			true,
		},
		{
			Block{
				ID:   163,
				Data: 5,
			},
			2, 1, 0,
			true,
		},
		//Test metadata [un]set
		{
			Block{
				metadata: nbt.Compound{
					nbt.NewTag("testInt2", nbt.Int(1743)),
					nbt.NewTag("testString2", nbt.String("world")),
				},
			},
			0, 0, 1,
			true,
		},
		{
			Block{
				metadata: nbt.Compound{
					nbt.NewTag("testInt", nbt.Int(15)),
					nbt.NewTag("testString", nbt.String("hello")),
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
		chunk.SetBlock(tB.x, tB.y, tB.z, tB.Block)
		if block := chunk.GetBlock(tB.x, tB.y, tB.z); !tB.Block.EqualBlock(block) {
			t.Errorf("blocks do not match, expecting %s, got %s", tB.Block.String(), block.String())
		}
	}
	for _, tB := range testBlocks {
		if tB.recheck {
			if block := chunk.GetBlock(tB.x, tB.y, tB.z); !tB.Block.EqualBlock(block) {
				t.Errorf("blocks do not match, expecting:-\n%s\ngot:-\n%s", tB.Block.String(), block.String())
			}
		}
	}
}

func TestHeightMap(t *testing.T) {
	tests := []struct {
		x, y, z int32
		Block
		height int32
	}{
		{0, 0, 0, Block{}, 0},
		{1, 0, 0, Block{ID: 1}, 1},
		{1, 1, 0, Block{ID: 1}, 2},
		{1, 0, 0, Block{}, 2},
		{1, 1, 0, Block{}, 0},
		{2, 10, 0, Block{ID: 1}, 11},
		{2, 12, 0, Block{ID: 1}, 13},
		{2, 12, 0, Block{}, 11},
		{2, 10, 0, Block{}, 0},
		{3, 15, 0, Block{ID: 1}, 16},
		{3, 16, 0, Block{ID: 1}, 17},
		{3, 16, 0, Block{}, 16},
		{3, 15, 0, Block{}, 0},
		{4, 31, 0, Block{ID: 1}, 32},
		{4, 32, 0, Block{ID: 1}, 33},
		{4, 32, 0, Block{}, 32},
		{4, 31, 0, Block{}, 0},
		{5, 16, 0, Block{ID: 1}, 17},
		{5, 32, 0, Block{ID: 1}, 33},
		{5, 32, 0, Block{}, 17},
		{5, 16, 0, Block{}, 0},
	}
	chunk, _ := newChunk(0, 0, nbt.Tag{})
	for n, test := range tests {
		chunk.SetBlock(test.x, test.y, test.z, test.Block)
		if h := chunk.GetHeight(test.x, test.z); h != test.height {
			t.Errorf("test %d: expecting height %d, got %d", n+1, test.height, h)
		}
	}
}
