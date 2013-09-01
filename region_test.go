package minecraft

import (
	"testing"
)

func TestRegion(t *testing.T) {
	m := NewMemPath()
	r := new(region)
	blocks := []struct {
		x, y, z int32
		block   *Block
	}{
		{0, 0, 0, &Block{BlockId: 12, Data: 13}},
		{0, 250, 0, &Block{BlockId: 2, Data: 1}},
		{15, 0, 15, &Block{BlockId: 45, Data: 14}},
		{15, 250, 15, &Block{BlockId: 67, Data: 4}},
		{18, 25, 1, &Block{BlockId: 143, Data: 7}},
		{1, 35, 18, &Block{BlockId: 431, Data: 6}},
		{500, 0, 10, &Block{BlockId: 32, Data: 8}},
		{432, 120, 400, &Block{BlockId: 98, Data: 13}},
		{10, 45, 43, &Block{BlockId: 179, Data: 5}},
		{20, 101, 91, &Block{BlockId: 258, Data: 11}},
		{342, 201, 40, &Block{BlockId: 891, Data: 5}},
		{87, 43, 251, &Block{BlockId: 1003, Data: 8}},
		{90, 98, 511, &Block{BlockId: 365, Data: 15}},
		{411, 100, 192, &Block{BlockId: 13, Data: 0}},
		{306, 255, 77, &Block{BlockId: 1, Data: 14}},
		{511, 255, 49, &Block{BlockId: 84, Data: 4}},
		{85, 80, 165, &Block{BlockId: 531, Data: 2}},
		{19, 231, 61, &Block{BlockId: 712, Data: 8}},
		{32, 1, 6, &Block{BlockId: 162, Data: 3}},
		{23, 8, 99, &Block{BlockId: 801, Data: 10}},
	}
	for n, block := range blocks {
		if err := r.SetBlock(m, block.x, block.y, block.z, block.block); err != nil {
			t.Error(err.Error())
		} else if thatBlock, err := r.GetBlock(m, block.x, block.y, block.z); err != nil {
			t.Error(err.Error())
		} else if !thatBlock.Equal(block.block) {
			t.Errorf("test %d failed: expecting %s, got %s", n+1, block.block.String(), thatBlock.String())
		}
	}
	r.Save(m)
	r = new(region)
	for n, block := range blocks {
		if thatBlock, err := r.GetBlock(m, block.x, block.y, block.z); err != nil {
			t.Error(err.Error())
		} else if !thatBlock.Equal(block.block) {
			t.Errorf("test %d failed: expecting %s, got %s", n+1, block.block.String(), thatBlock.String())
		}
	}
}
