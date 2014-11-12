package minecraft

import "testing"

func TestNewLevel(t *testing.T) {
	m := NewMemPath()
	l, err := NewLevel(m)
	if err != nil {
		t.Error(err.Error())
		return
	}
	x, y, z := int32(1534545), int32(23), int32(-56456)
	l.Options(Spawn(x, y, z))
	l.Options(GetSpawn(&x, &y, &z))
	if x != 1534545 || y != 23 || z != -56456 {
		t.Errorf("[SG]etSpawn test failed, expecting 1534545, 23, -56456, got %d, %d, %d", x, y, z)
	}
	biomes := []struct {
		x, z int32
		Biome
	}{
		{0, 0, 1},
		{45323, 5, 6},
		{56454, 868, 4},
		{45645, 23498, 22},
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
		{0, 0, 0, &Block{BlockID: 12, Data: 13}},
		{0, 250, 0, &Block{BlockID: 2, Data: 1}},
		{185, 0, 10115, &Block{BlockID: 45, Data: 14}},
		{4564, 250, 4645, &Block{BlockID: 67, Data: 4}},
		{4232, 25, -4234234, &Block{BlockID: 143, Data: 7}},
		{-2427824, 35, 23214, &Block{BlockID: 431, Data: 6}},
		{-23478621, 0, -12341234, &Block{BlockID: 32, Data: 8}},
		{4438, 120, -3123, &Block{BlockID: 98, Data: 13}},
		{9762, 45, 3873, &Block{BlockID: 179, Data: 5}},
		{39234, 101, 37482, &Block{BlockID: 258, Data: 11}},
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
	if l, err = NewLevel(m); err != nil {
		t.Error(err.Error())
		return
	}
	l.Options(GetSpawn(&x, &y, &z))
	if x != 1534545 || y != 23 || z != -56456 {
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

func TestSkyLight(t *testing.T) {
	l, _ := NewLevel(NewMemPath())
	type posBlock struct {
		x, y, z int32
		*Block
	}
	tests := []struct {
		blocks []posBlock
		light  [][4]int32 //x, y, z, skyLight
	}{
		{[]posBlock{
			{0, 20, 0, &Block{BlockID: 39}},
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
		{[]posBlock{
			{0, 20, 0, &Block{BlockID: 1}},
		}, [][4]int32{
			{0, 21, 0, 15},
			{0, 20, 0, 0},
			{0, 19, 0, 14},
			{0, 18, 0, 14},
			{0, 16, 0, 14},
		}},
		{[]posBlock{
			{-1, 20, -1, &Block{BlockID: 1}},
			{-1, 20, 0, &Block{BlockID: 1}},
			{-1, 20, 1, &Block{BlockID: 1}},
			{0, 20, -1, &Block{BlockID: 1}},
			{0, 20, 1, &Block{BlockID: 1}},
			{1, 20, -1, &Block{BlockID: 1}},
			{1, 20, 0, &Block{BlockID: 1}},
			{1, 20, 1, &Block{BlockID: 1}},
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
		{[]posBlock{
			{0, 20, 0, &Block{}},
		}, [][4]int32{
			{0, 20, 0, 15},
			{0, 19, 0, 15},
			{0, 18, 0, 15},
			{0, 16, 0, 15},
		}},
		{[]posBlock{
			{0, 20, 0, &Block{BlockID: 1}},
		}, [][4]int32{
			{0, 20, 0, 0},
			{0, 19, 0, 13},
			{0, 19, 1, 14},
			{1, 19, 0, 14},
			{0, 18, 0, 13},
			{0, 16, 0, 13},
		}},
		{[]posBlock{
			{10, 20, 10, &Block{BlockID: 1}},
			{11, 20, 10, &Block{BlockID: 1}},
			{12, 20, 10, &Block{BlockID: 1}},
			{13, 20, 10, &Block{BlockID: 1}},
			{14, 20, 10, &Block{BlockID: 1}},
			{10, 20, 11, &Block{BlockID: 1}},
			{11, 20, 11, &Block{BlockID: 1}},
			{12, 20, 11, &Block{BlockID: 1}},
			{13, 20, 11, &Block{BlockID: 1}},
			{14, 20, 11, &Block{BlockID: 1}},
			{10, 20, 12, &Block{BlockID: 1}},
			{11, 20, 12, &Block{BlockID: 1}},
			{12, 20, 12, &Block{BlockID: 1}},
			{13, 20, 12, &Block{BlockID: 1}},
			{14, 20, 12, &Block{BlockID: 1}},
			{10, 20, 13, &Block{BlockID: 1}},
			{11, 20, 13, &Block{BlockID: 1}},
			{12, 20, 13, &Block{BlockID: 1}},
			{13, 20, 13, &Block{BlockID: 1}},
			{14, 20, 13, &Block{BlockID: 1}},
			{10, 20, 14, &Block{BlockID: 1}},
			{11, 20, 14, &Block{BlockID: 1}},
			{12, 20, 14, &Block{BlockID: 1}},
			{13, 20, 14, &Block{BlockID: 1}},
			{14, 20, 14, &Block{BlockID: 1}},

			{10, 19, 10, &Block{BlockID: 1}},
			{11, 19, 10, &Block{BlockID: 1}},
			{12, 19, 10, &Block{BlockID: 1}},
			{13, 19, 10, &Block{BlockID: 1}},
			{14, 19, 10, &Block{BlockID: 1}},
			{10, 19, 11, &Block{BlockID: 1}},
			{14, 19, 11, &Block{BlockID: 1}},
			{10, 19, 12, &Block{BlockID: 1}},
			{14, 19, 12, &Block{BlockID: 1}},
			{10, 19, 13, &Block{BlockID: 1}},
			{14, 19, 13, &Block{BlockID: 1}},
			{10, 19, 14, &Block{BlockID: 1}},
			{11, 19, 14, &Block{BlockID: 1}},
			{12, 19, 14, &Block{BlockID: 1}},
			{13, 19, 14, &Block{BlockID: 1}},
			{14, 19, 14, &Block{BlockID: 1}},

			{10, 18, 10, &Block{BlockID: 1}},
			{11, 18, 10, &Block{BlockID: 1}},
			{12, 18, 10, &Block{BlockID: 1}},
			{13, 18, 10, &Block{BlockID: 1}},
			{14, 18, 10, &Block{BlockID: 1}},
			{10, 18, 11, &Block{BlockID: 1}},
			{14, 18, 11, &Block{BlockID: 1}},
			{10, 18, 12, &Block{BlockID: 1}},
			{14, 18, 12, &Block{BlockID: 1}},
			{10, 18, 13, &Block{BlockID: 1}},
			{14, 18, 13, &Block{BlockID: 1}},
			{10, 18, 14, &Block{BlockID: 1}},
			{11, 18, 14, &Block{BlockID: 1}},
			{12, 18, 14, &Block{BlockID: 1}},
			{13, 18, 14, &Block{BlockID: 1}},
			{14, 18, 14, &Block{BlockID: 1}},

			{10, 17, 10, &Block{BlockID: 1}},
			{11, 17, 10, &Block{BlockID: 1}},
			{12, 17, 10, &Block{BlockID: 1}},
			{13, 17, 10, &Block{BlockID: 1}},
			{14, 17, 10, &Block{BlockID: 1}},
			{10, 17, 11, &Block{BlockID: 1}},
			{14, 17, 11, &Block{BlockID: 1}},
			{10, 17, 12, &Block{BlockID: 1}},
			{14, 17, 12, &Block{BlockID: 1}},
			{10, 17, 13, &Block{BlockID: 1}},
			{14, 17, 13, &Block{BlockID: 1}},
			{10, 17, 14, &Block{BlockID: 1}},
			{11, 17, 14, &Block{BlockID: 1}},
			{12, 17, 14, &Block{BlockID: 1}},
			{13, 17, 14, &Block{BlockID: 1}},
			{14, 17, 14, &Block{BlockID: 1}},

			{10, 16, 10, &Block{BlockID: 1}},
			{11, 16, 10, &Block{BlockID: 1}},
			{12, 16, 10, &Block{BlockID: 1}},
			{13, 16, 10, &Block{BlockID: 1}},
			{14, 16, 10, &Block{BlockID: 1}},
			{10, 16, 11, &Block{BlockID: 1}},
			{11, 16, 11, &Block{BlockID: 1}},
			{12, 16, 11, &Block{BlockID: 1}},
			{13, 16, 11, &Block{BlockID: 1}},
			{14, 16, 11, &Block{BlockID: 1}},
			{10, 16, 12, &Block{BlockID: 1}},
			{11, 16, 12, &Block{BlockID: 1}},
			{12, 16, 12, &Block{BlockID: 1}},
			{13, 16, 12, &Block{BlockID: 1}},
			{14, 16, 12, &Block{BlockID: 1}},
			{10, 16, 13, &Block{BlockID: 1}},
			{11, 16, 13, &Block{BlockID: 1}},
			{12, 16, 13, &Block{BlockID: 1}},
			{13, 16, 13, &Block{BlockID: 1}},
			{14, 16, 13, &Block{BlockID: 1}},
			{10, 16, 14, &Block{BlockID: 1}},
			{11, 16, 14, &Block{BlockID: 1}},
			{12, 16, 14, &Block{BlockID: 1}},
			{13, 16, 14, &Block{BlockID: 1}},
			{14, 16, 14, &Block{BlockID: 1}},
		}, [][4]int32{
			{10, 20, 10, 0},
			{14, 20, 14, 0},
			{12, 19, 12, 0},
			{13, 18, 13, 0},
			{11, 17, 11, 0},
		}},
		{[]posBlock{
			{100, 10, 100, &Block{BlockID: 8}},
		}, [][4]int32{
			{100, 10, 100, 12},
		}},
		{[]posBlock{
			{99, 12, 100, &Block{BlockID: 1}},
			{99, 11, 100, &Block{BlockID: 1}},
			{99, 10, 100, &Block{BlockID: 1}},
			{99, 9, 100, &Block{BlockID: 1}},
			{99, 8, 100, &Block{BlockID: 1}},
			{99, 7, 100, &Block{BlockID: 1}},

			{101, 12, 100, &Block{BlockID: 1}},
			{101, 11, 100, &Block{BlockID: 1}},
			{101, 10, 100, &Block{BlockID: 1}},
			{101, 9, 100, &Block{BlockID: 1}},
			{101, 8, 100, &Block{BlockID: 1}},
			{101, 7, 100, &Block{BlockID: 1}},

			{100, 12, 99, &Block{BlockID: 1}},
			{100, 11, 99, &Block{BlockID: 1}},
			{100, 10, 99, &Block{BlockID: 1}},
			{100, 9, 99, &Block{BlockID: 1}},
			{100, 8, 99, &Block{BlockID: 1}},
			{100, 7, 99, &Block{BlockID: 1}},

			{100, 12, 101, &Block{BlockID: 1}},
			{100, 11, 101, &Block{BlockID: 1}},
			{100, 10, 101, &Block{BlockID: 1}},
			{100, 9, 101, &Block{BlockID: 1}},
			{100, 8, 101, &Block{BlockID: 1}},
			{100, 7, 101, &Block{BlockID: 1}},

			{100, 8, 100, &Block{BlockID: 8}},
			{100, 7, 100, &Block{BlockID: 1}},

			{100, 11, 100, &Block{BlockID: 8}},
			{100, 9, 100, &Block{BlockID: 8}},
		}, [][4]int32{
			{100, 8, 100, 3},
			{100, 9, 100, 6},
			{100, 10, 100, 9},
			{100, 11, 100, 12},
		}},
		{[]posBlock{
			{100, 12, 100, &Block{BlockID: 8}},
			{100, 7, 100, &Block{BlockID: 8}},
			{100, 6, 100, &Block{BlockID: 1}},
		}, [][4]int32{
			{100, 6, 100, 0},
			{100, 7, 100, 0},
			{100, 8, 100, 0},
			{100, 9, 100, 3},
			{100, 10, 100, 6},
			{100, 11, 100, 9},
			{100, 12, 100, 12},
		}},
		{[]posBlock{
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

func TestBlockLight(t *testing.T) {
	l, _ := NewLevel(NewMemPath())
	type posBlock struct {
		x, y, z int32
		*Block
	}
	tests := []struct {
		blocks []posBlock
		light  [][4]int32 //x, y, z, skyLight
	}{
		{[]posBlock{ //Test 1
			{0, 10, 0, &Block{BlockID: 89}},
		}, [][4]int32{
			{0, 10, 0, 15},
			{0, 9, 0, 14},
			{0, 8, 0, 13},
			{0, 7, 0, 12},
			{0, 6, 0, 11},
			{0, 5, 0, 10},
			{0, 4, 0, 9},
			{0, 3, 0, 8},
			{0, 2, 0, 7},
			{0, 1, 0, 6},
			{0, 0, 0, 5},
			{0, 11, 0, 14},
			{0, 12, 0, 13},
			{0, 13, 0, 12},
			{0, 14, 0, 11},
			{0, 15, 0, 10},
			{1, 10, 0, 14},
			{0, 10, 1, 14},
			{1, 10, 1, 13},
		}},
		{[]posBlock{ //Test 2
			{-16, 15, 15, &Block{BlockID: 20}},
		}, [][4]int32{
			{-1, 10, 0, 14},
			{-2, 10, 0, 13},
			{-3, 10, 0, 12},
			{-1, 10, 1, 13},
			{-1, 10, 2, 12},
			{-1, 10, 3, 11},
			{-1, 11, 0, 13},
			{-2, 11, 0, 12},
		}},
		{[]posBlock{ //Test 3
			{-16, 15, 15, &Block{}},
			{-1, 10, 0, &Block{BlockID: 20}},
		}, [][4]int32{
			{-1, 10, 0, 14},
			{-2, 10, 0, 13},
			{-3, 10, 0, 12},
			{-1, 10, 1, 13},
			{-1, 10, 2, 12},
			{-1, 10, 3, 11},
			{-1, 11, 0, 13},
			{-2, 11, 0, 12},
		}},
		{[]posBlock{ //Test 4
			{1, 10, 0, &Block{BlockID: 1}},
		}, [][4]int32{
			{1, 10, 0, 0},
			{2, 10, 0, 11},
			{3, 10, 0, 10},
			{2, 10, 1, 12},
			{3, 10, 1, 11},
			{1, 9, 0, 13},
			{1, 11, 0, 13},
			{1, 9, 1, 12},
		}},
	}
	for n, test := range tests {
		for _, b := range test.blocks {
			l.SetBlock(b.x, b.y, b.z, b.Block)
		}
		for o, light := range test.light {
			if m, _ := l.getBlockLight(light[0], light[1], light[2]); int32(m) != light[3] {
				t.Errorf("test %d-%d: block light level at [%d, %d, %d] does not match expected, got %d, expecting %d", n+1, o+1, light[0], light[1], light[2], m, light[3])
			}
		}
	}
}

func BenchmarkSkyLight(b *testing.B) {
	l, _ := NewLevel(NewMemPath())
	block := &Block{BlockID: 1}
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
