package minecraft

import "testing"

func TestNewLevel(t *testing.T) {
	m := NewMemPath()
	l, err := NewLevel(m, 0)
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
	if l, err = NewLevel(m, 0); err != nil {
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

func TestLightingSimpleSkyLight(t *testing.T) {
	l, _ := NewLevel(NewMemPath(), LIGHT_SKY_SIMPLE)
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
				t.Errorf("test %d-%d: sky light level at [%d, %d, %d] does not match expected, got %d, expecting %d", n+1, o+1, light[0], light[1], light[2], m, light[3])
			}
		}
	}
}

func BenchmarkLightingSimpleSkyLight(b *testing.B) {
	l, _ := NewLevel(NewMemPath(), LIGHT_SKY_SIMPLE)
	block := &Block{BlockId: 1}
	for n := 0; n < b.N; n++ {
		m := int32(n)
		for i := int32(0); i < 5; i++ {
			for j := int32(0); j < 5; j++ {
				l.SetBlock(16*m+i, 20, j, block)
			}
		}
	}
}

func TestLightingAllSkyLight(t *testing.T) {
	l, _ := NewLevel(NewMemPath(), LIGHT_SKY_ALL)
	type posBlock struct {
		x, y, z int32
		*Block
	}
	tests := []struct {
		blocks []posBlock
		light  [][4]int32 //x, y, z, skyLight
	}{
		{[]posBlock{ //Test 1
			{0, 20, 0, &Block{BlockId: 39}},
		}, [][4]int32{
			{0, 20, 0, 15},
			{0, 21, 0, 15},
			{0, 19, 0, 15},
			{0, 26, 0, 15},
			{0, 31, 0, 15},
			{15, 31, 15, 15},
			{15, 15, 15, 15},
			{15, 31, 0, 15},
			{0, 15, 15, 15},
			{0, 31, 0, 15},
			{23, 23, 23, 15},
		}},
		{[]posBlock{ //Test 2
			{0, 20, 0, &Block{BlockId: 1}},
		}, [][4]int32{
			{0, 21, 0, 15},
			{0, 20, 0, 0},
			{0, 19, 0, 14},
			{0, 18, 0, 14},
			{0, 16, 0, 14},
		}},
		{[]posBlock{ //Test 3
			{-1, 20, -1, &Block{BlockId: 1}},
			{-1, 20, 0, &Block{BlockId: 1}},
			{-1, 20, 1, &Block{BlockId: 1}},
			{0, 20, -1, &Block{BlockId: 1}},
			{0, 20, 1, &Block{BlockId: 1}},
			{1, 20, -1, &Block{BlockId: 1}},
			{1, 20, 0, &Block{BlockId: 1}},
			{1, 20, 1, &Block{BlockId: 1}},
		}, [][4]int32{
			{0, 19, 0, 13},
			{-1, 19, -1, 14},
			{1, 19, 1, 14},
			{-1, 19, 1, 14},
			{1, 19, -1, 14},
			{0, 18, 0, 13},
			{-1, 18, -1, 14},
			{1, 18, 1, 14},
			{-1, 18, 1, 14},
			{1, 18, -1, 14},
		}},
		{[]posBlock{ //Test 4
			{0, 20, 0, &Block{}},
		}, [][4]int32{
			{0, 20, 0, 15},
			{0, 19, 0, 15},
			{0, 18, 0, 15},
			{0, 16, 0, 15},
		}},
		{[]posBlock{ //Test 5
			{0, 20, 0, &Block{BlockId: 1}},
		}, [][4]int32{
			{0, 20, 0, 0},
			{0, 19, 0, 13},
			{0, 19, 1, 14},
			{1, 19, 0, 14},
			{0, 18, 0, 13},
			{0, 16, 0, 13},
		}},
		{[]posBlock{ //Test 6
			{10, 20, 10, &Block{BlockId: 1}},
			{11, 20, 10, &Block{BlockId: 1}},
			{12, 20, 10, &Block{BlockId: 1}},
			{13, 20, 10, &Block{BlockId: 1}},
			{14, 20, 10, &Block{BlockId: 1}},
			{10, 20, 11, &Block{BlockId: 1}},
			{11, 20, 11, &Block{BlockId: 1}},
			{12, 20, 11, &Block{BlockId: 1}},
			{13, 20, 11, &Block{BlockId: 1}},
			{14, 20, 11, &Block{BlockId: 1}},
			{10, 20, 12, &Block{BlockId: 1}},
			{11, 20, 12, &Block{BlockId: 1}},
			{12, 20, 12, &Block{BlockId: 1}},
			{13, 20, 12, &Block{BlockId: 1}},
			{14, 20, 12, &Block{BlockId: 1}},
			{10, 20, 13, &Block{BlockId: 1}},
			{11, 20, 13, &Block{BlockId: 1}},
			{12, 20, 13, &Block{BlockId: 1}},
			{13, 20, 13, &Block{BlockId: 1}},
			{14, 20, 13, &Block{BlockId: 1}},
			{10, 20, 14, &Block{BlockId: 1}},
			{11, 20, 14, &Block{BlockId: 1}},
			{12, 20, 14, &Block{BlockId: 1}},
			{13, 20, 14, &Block{BlockId: 1}},
			{14, 20, 14, &Block{BlockId: 1}},

			{10, 19, 10, &Block{BlockId: 1}},
			{11, 19, 10, &Block{BlockId: 1}},
			{12, 19, 10, &Block{BlockId: 1}},
			{13, 19, 10, &Block{BlockId: 1}},
			{14, 19, 10, &Block{BlockId: 1}},
			{10, 19, 11, &Block{BlockId: 1}},
			{14, 19, 11, &Block{BlockId: 1}},
			{10, 19, 12, &Block{BlockId: 1}},
			{14, 19, 12, &Block{BlockId: 1}},
			{10, 19, 13, &Block{BlockId: 1}},
			{14, 19, 13, &Block{BlockId: 1}},
			{10, 19, 14, &Block{BlockId: 1}},
			{11, 19, 14, &Block{BlockId: 1}},
			{12, 19, 14, &Block{BlockId: 1}},
			{13, 19, 14, &Block{BlockId: 1}},
			{14, 19, 14, &Block{BlockId: 1}},

			{10, 18, 10, &Block{BlockId: 1}},
			{11, 18, 10, &Block{BlockId: 1}},
			{12, 18, 10, &Block{BlockId: 1}},
			{13, 18, 10, &Block{BlockId: 1}},
			{14, 18, 10, &Block{BlockId: 1}},
			{10, 18, 11, &Block{BlockId: 1}},
			{14, 18, 11, &Block{BlockId: 1}},
			{10, 18, 12, &Block{BlockId: 1}},
			{14, 18, 12, &Block{BlockId: 1}},
			{10, 18, 13, &Block{BlockId: 1}},
			{14, 18, 13, &Block{BlockId: 1}},
			{10, 18, 14, &Block{BlockId: 1}},
			{11, 18, 14, &Block{BlockId: 1}},
			{12, 18, 14, &Block{BlockId: 1}},
			{13, 18, 14, &Block{BlockId: 1}},
			{14, 18, 14, &Block{BlockId: 1}},

			{10, 17, 10, &Block{BlockId: 1}},
			{11, 17, 10, &Block{BlockId: 1}},
			{12, 17, 10, &Block{BlockId: 1}},
			{13, 17, 10, &Block{BlockId: 1}},
			{14, 17, 10, &Block{BlockId: 1}},
			{10, 17, 11, &Block{BlockId: 1}},
			{14, 17, 11, &Block{BlockId: 1}},
			{10, 17, 12, &Block{BlockId: 1}},
			{14, 17, 12, &Block{BlockId: 1}},
			{10, 17, 13, &Block{BlockId: 1}},
			{14, 17, 13, &Block{BlockId: 1}},
			{10, 17, 14, &Block{BlockId: 1}},
			{11, 17, 14, &Block{BlockId: 1}},
			{12, 17, 14, &Block{BlockId: 1}},
			{13, 17, 14, &Block{BlockId: 1}},
			{14, 17, 14, &Block{BlockId: 1}},

			{10, 16, 10, &Block{BlockId: 1}},
			{11, 16, 10, &Block{BlockId: 1}},
			{12, 16, 10, &Block{BlockId: 1}},
			{13, 16, 10, &Block{BlockId: 1}},
			{14, 16, 10, &Block{BlockId: 1}},
			{10, 16, 11, &Block{BlockId: 1}},
			{11, 16, 11, &Block{BlockId: 1}},
			{12, 16, 11, &Block{BlockId: 1}},
			{13, 16, 11, &Block{BlockId: 1}},
			{14, 16, 11, &Block{BlockId: 1}},
			{10, 16, 12, &Block{BlockId: 1}},
			{11, 16, 12, &Block{BlockId: 1}},
			{12, 16, 12, &Block{BlockId: 1}},
			{13, 16, 12, &Block{BlockId: 1}},
			{14, 16, 12, &Block{BlockId: 1}},
			{10, 16, 13, &Block{BlockId: 1}},
			{11, 16, 13, &Block{BlockId: 1}},
			{12, 16, 13, &Block{BlockId: 1}},
			{13, 16, 13, &Block{BlockId: 1}},
			{14, 16, 13, &Block{BlockId: 1}},
			{10, 16, 14, &Block{BlockId: 1}},
			{11, 16, 14, &Block{BlockId: 1}},
			{12, 16, 14, &Block{BlockId: 1}},
			{13, 16, 14, &Block{BlockId: 1}},
			{14, 16, 14, &Block{BlockId: 1}},
		}, [][4]int32{
			{10, 20, 10, 0},
			{14, 20, 14, 0},
			{12, 19, 12, 0},
			{13, 18, 13, 0},
			{11, 17, 11, 0},
		}},
		{[]posBlock{ //Test 7
			{100, 10, 100, &Block{BlockId: 8}},
		}, [][4]int32{
			{100, 10, 100, 12},
		}},
		{[]posBlock{ //Test 8
			{99, 12, 100, &Block{BlockId: 1}},
			{99, 11, 100, &Block{BlockId: 1}},
			{99, 10, 100, &Block{BlockId: 1}},
			{99, 9, 100, &Block{BlockId: 1}},
			{99, 8, 100, &Block{BlockId: 1}},
			{99, 7, 100, &Block{BlockId: 1}},

			{101, 12, 100, &Block{BlockId: 1}},
			{101, 11, 100, &Block{BlockId: 1}},
			{101, 10, 100, &Block{BlockId: 1}},
			{101, 9, 100, &Block{BlockId: 1}},
			{101, 8, 100, &Block{BlockId: 1}},
			{101, 7, 100, &Block{BlockId: 1}},

			{100, 12, 99, &Block{BlockId: 1}},
			{100, 11, 99, &Block{BlockId: 1}},
			{100, 10, 99, &Block{BlockId: 1}},
			{100, 9, 99, &Block{BlockId: 1}},
			{100, 8, 99, &Block{BlockId: 1}},
			{100, 7, 99, &Block{BlockId: 1}},

			{100, 12, 101, &Block{BlockId: 1}},
			{100, 11, 101, &Block{BlockId: 1}},
			{100, 10, 101, &Block{BlockId: 1}},
			{100, 9, 101, &Block{BlockId: 1}},
			{100, 8, 101, &Block{BlockId: 1}},
			{100, 7, 101, &Block{BlockId: 1}},

			{100, 8, 100, &Block{BlockId: 8}},
			{100, 7, 100, &Block{BlockId: 1}},

			{100, 11, 100, &Block{BlockId: 8}},
			{100, 9, 100, &Block{BlockId: 8}},
		}, [][4]int32{
			{100, 8, 100, 3},
			{100, 9, 100, 6},
			{100, 10, 100, 9},
			{100, 11, 100, 12},
		}},
		{[]posBlock{ //Test 9
			{100, 12, 100, &Block{BlockId: 8}},
			{100, 7, 100, &Block{BlockId: 8}},
			{100, 6, 100, &Block{BlockId: 1}},
		}, [][4]int32{
			{100, 6, 100, 0},
			{100, 7, 100, 0},
			{100, 8, 100, 0},
			{100, 9, 100, 3},
			{100, 10, 100, 6},
			{100, 11, 100, 9},
			{100, 12, 100, 12},
		}},
		{[]posBlock{ //Test 9
			{100, 6, 100, &Block{}},
		}, [][4]int32{
			{100, 6, 100, 13},
			{100, 7, 100, 10},
			{100, 8, 100, 7},
			{100, 9, 100, 4},
			{100, 10, 100, 6},
			{100, 11, 100, 9},
			{100, 12, 100, 12},
		}},
	}
	for n, test := range tests {
		for _, b := range test.blocks {
			l.SetBlock(b.x, b.y, b.z, b.Block)
		}
		for o, light := range test.light {
			if m, _ := l.getSkyLight(light[0], light[1], light[2]); int32(m) != light[3] {
				t.Errorf("test %d-%d: sky light level at [%d, %d, %d] does not match expected, got %d, expecting %d", n+1, o+1, light[0], light[1], light[2], m, light[3])
			}
		}
	}
}

func BenchmarkLightingAllSkyLight(b *testing.B) {
	l, _ := NewLevel(NewMemPath(), LIGHT_SKY_ALL)
	block := &Block{BlockId: 1}
	for n := 0; n < b.N; n++ {
		m := int32(n)
		for i := int32(0); i < 5; i++ {
			for j := int32(0); j < 5; j++ {
				l.SetBlock(16*m+i, 20, j, block)
			}
		}
	}
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
		return 15, nil
	}
	return c.GetSkyLight(x, y, z), nil
}
