package minecraft

import (
	"testing"
)

func TestNewLevel(t *testing.T) {
	m := NewMemPath()
	l, err := NewLevel(m, LIGHT_NONE)
	if err != nil {
		t.Error(err.Error())
		return
	}
	l.SetSpawn(1534545, 23, -56456)
	if x, y, z := l.GetSpawn(); x != 1534545 || y != 23 || z != -56456 {
		t.Errorf("[SG]etSpawn test failed, expecting 1534545, 23, -56456, got %d, %d, %d", x, y, z)
	}
	biomes := []struct {
		x, z int32
		Biome
	}{
		{0, 0, 1},
		{45323, 5, 6},
		{56454, 868, 4},
		{45645, 23498, -1},
		{-42536, 0, 5},
		{-23465, 5426, 9},
		{7843, -3265, 12},
		{45, -2465783, 4},
		{2553, -26582, 4},
		{-2358, -4564, 8},
		{456, 45646, 5},
	}
	blocks := []struct {
		x, y, z int32
		*Block
	}{
		{0, 0, 0, &Block{BlockId: 12, Data: 13}},
		{0, 250, 0, &Block{BlockId: 2, Data: 1}},
		{185, 0, 10115, &Block{BlockId: 45, Data: 14}},
		{4564, 250, 4645, &Block{BlockId: 67, Data: 4}},
		{4232, 25, -4234234, &Block{BlockId: 143, Data: 7}},
		{-2427824, 35, 23214, &Block{BlockId: 431, Data: 6}},
		{-23478621, 0, -12341234, &Block{BlockId: 32, Data: 8}},
		{4438, 120, -3123, &Block{BlockId: 98, Data: 13}},
		{9762, 45, 3873, &Block{BlockId: 179, Data: 5}},
		{39234, 101, 37482, &Block{BlockId: 258, Data: 11}},
	}

	for n, biome := range biomes {
		if err := l.SetBiome(biome.x, biome.z, biome.Biome); err != nil {
			t.Error(err.Error())
		} else if tBiome, err := l.GetBiome(biome.x, biome.z); err != nil {
			t.Error(err.Error())
		} else if tBiome != biome.Biome {
			t.Errorf("biome test %d: expecting %s, got %s", n, biome, tBiome)
		}
	}

	for n, block := range blocks {
		if err := l.SetBlock(block.x, block.y, block.z, block.Block); err != nil {
			t.Error(err.Error())
		} else if tBlock, err := l.GetBlock(block.x, block.y, block.z); err != nil {
			t.Error(err.Error())
		} else if !block.Equal(tBlock) {
			t.Errorf("biome test %d: expecting %s, got %s", n, block, tBlock)
		}
	}
	l.Save()
	if l, err = NewLevel(m, LIGHT_NONE); err != nil {
		t.Error(err.Error())
		return
	}
	if x, y, z := l.GetSpawn(); x != 1534545 || y != 23 || z != -56456 {
		t.Errorf("[SG]etSpawn test failed, expecting 1534545, 23, -56456, got %d, %d, %d", x, y, z)
	}
	for n, biome := range biomes {
		if tBiome, err := l.GetBiome(biome.x, biome.z); err != nil {
			t.Error(err.Error())
		} else if tBiome != biome.Biome {
			t.Errorf("biome test %d: expecting %s, got %s", n, biome.Biome, tBiome)
		}
	}
	for n, block := range blocks {
		if tBlock, err := l.GetBlock(block.x, block.y, block.z); err != nil {
			t.Error(err.Error())
		} else if !block.Equal(tBlock) {
			t.Errorf("biome test %d: expecting %s, got %s", n, block, tBlock)
		}
	}
}

// var (
// 	stone = &Block{BlockId: 1}
// 	water = &Block{BlockId: 8}
// 	glass = &Block{BlockId: 20}
// }

func TestLightingSimpleSkyLight(t *testing.T) {
	l, _ := NewLevel(NewMemPath(), LIGHT_SIMPLE)
	tests := []struct {
		x, y, z int32
		*Block
		light [][4]int32 //x, y, z, skyLight
	}{
		{0, 20, 0, &Block{BlockId: 1}, [][4]int32{{0, 20, 0, 0}, {0, 19, 0, 0}, {0, 0, 0, 0}}},
		{0, 19, 0, &Block{BlockId: 1}, [][4]int32{{0, 20, 0, 0}, {0, 19, 0, 0}, {0, 0, 0, 0}}},
		{0, 20, 0, &Block{}, [][4]int32{{1, 20, 0, 15}, {0, 19, 0, 0}, {0, 0, 0, 0}}},
		{1, 20, 0, &Block{BlockId: 8}, [][4]int32{{1, 20, 0, 12}, {1, 19, 0, 11}, {1, 0, 0, 0}}},
		{1, 19, 0, &Block{BlockId: 8}, [][4]int32{{1, 20, 0, 12}, {1, 19, 0, 9}, {1, 0, 0, 0}}},
		{1, 20, 0, &Block{}, [][4]int32{{1, 20, 0, 15}, {1, 19, 0, 12}, {1, 0, 0, 0}}},
		{0, 16, 1, &Block{BlockId: 1}, [][4]int32{{0, 16, 1, 0}, {0, 15, 1, 0}, {0, 0, 1, 0}}},
		{0, 15, 1, &Block{BlockId: 1}, [][4]int32{{0, 16, 1, 0}, {0, 15, 1, 0}, {0, 0, 1, 0}}},
		{0, 16, 1, &Block{}, [][4]int32{{0, 16, 1, 15}, {0, 15, 1, 0}, {0, 0, 1, 0}}},
		{1, 16, 1, &Block{BlockId: 8}, [][4]int32{{1, 16, 1, 12}, {1, 15, 1, 11}, {1, 0, 1, 0}}},
		{1, 15, 1, &Block{BlockId: 8}, [][4]int32{{1, 16, 1, 12}, {1, 15, 1, 9}, {1, 0, 1, 0}}},
		{1, 16, 1, &Block{}, [][4]int32{{1, 16, 1, 15}, {1, 15, 1, 12}, {1, 0, 1, 0}}},
	}
	for n, test := range tests {
		l.SetBlock(test.x, test.y, test.z, test.Block)
		for o, light := range test.light {
			if m, _ := l.getSkyLight(light[0], light[1], light[2]); int32(m) != light[3] {
				t.Errorf("SimpleSkyLight test %d-%d: sky light level at [%d, %d, %d] does not match expected, got %d, expecting %d", n+1, o+1, light[0], light[1], light[2], m, light[3])
			}
		}
	}
}

func TestLightingAllSkyLight(t *testing.T) {

}

func (l *Level) getBlockLight(x, y, z int32) (uint8, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return 0, err
	} else if c == nil {
		return 0, nil
	}
	return c.GetBlockLight(x, y, z), nil
}

func (l *Level) getSkyLight(x, y, z int32) (uint8, error) {
	c, err := l.getChunk(x, z, false)
	if err != nil {
		return 0, err
	} else if c == nil {
		return 0, nil
	}
	return c.GetSkyLight(x, y, z), nil
}
