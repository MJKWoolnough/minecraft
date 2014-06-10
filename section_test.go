package minecraft

import (
	"github.com/MJKWoolnough/minecraft/nbt"
	"testing"
)

func TestY(t *testing.T) {
	section := newSection(0)
	if y := *section.section.Get("Y").Data().(*nbt.Byte); y != 0 {
		t.Errorf("expecting %d, got %d", 0, y)
	}
	section = newSection(16)
	if y := *section.section.Get("Y").Data().(*nbt.Byte); y != 1 {
		t.Errorf("expecting %d, got %d", 1, y)
	}
	section.SetY(48)
	if y := *section.section.Get("Y").Data().(*nbt.Byte); y != 3 {
		t.Errorf("expecting %d, got %d", 3, y)
	}
	section.SetY(255)
	if y := *section.section.Get("Y").Data().(*nbt.Byte); y != 15 {
		t.Errorf("expecting %d, got %d", 15, y)
	}
}

func TestGetBlock(t *testing.T) {
	blocks := make([]int8, 4096)
	add := make([]int8, 2048)
	data := make([]int8, 2048)
	blocks[0] = 1
	blocks[10] = 2
	blocks[18] = 24
	blocks[19] = 13
	blocks[3475] = 45
	blocks[4054] = b2i(194)
	add[9] = b2i(12<<4 | 7)
	add[2027] = 5
	data[1737] = b2i(9 << 4)
	data[2027] = 8
	section, _ := loadSection(nbt.NewCompound(nbt.Compound{
		nbt.NewTag("Blocks", nbt.NewByteArray(blocks)),
		nbt.NewTag("Add", nbt.NewByteArray(add)),
		nbt.NewTag("Data", nbt.NewByteArray(data)),
		nbt.NewTag("BlockLight", nbt.NewByteArray(make([]int8, 2048))),
		nbt.NewTag("SkyLight", nbt.NewByteArray(make([]int8, 2048))),
		nbt.NewTag("Y", nbt.NewByte(0)),
	}))
	tests := []struct {
		xyz [3]int32
		Block
	}{
		{[3]int32{0, 0, 0}, Block{BlockId: 1}},
		{[3]int32{10, 0, 0}, Block{BlockId: 2}},
		{[3]int32{2, 0, 1}, Block{BlockId: 1816}},
		{[3]int32{3, 0, 1}, Block{BlockId: 3085}},
		{[3]int32{3, 13, 9}, Block{BlockId: 45, Data: 9}},
		{[3]int32{6, 15, 13}, Block{BlockId: 1474, Data: 8}},
		{[3]int32{9, 12, 11}, Block{}},
	}
	for n, test := range tests {
		if b := section.GetBlock(test.xyz[0], test.xyz[1], test.xyz[2]); !test.Block.Equal(b) {
			t.Errorf("test %d failed\nExpecting: %s\nGot: %s", n+1, test.Block.String(), b.String())
		}
	}
}

func TestSetBlock(t *testing.T) {
	section := newSection(32)
	tests := []struct {
		xyz [3]int32
		Block
		recheck bool
	}{
		{[3]int32{4, 7, 9}, Block{BlockId: 12, Data: 4}, false},
		{[3]int32{4, 7, 9}, Block{BlockId: 15}, false},
		{[3]int32{1, 12, 10}, Block{BlockId: 1345, Data: 12}, true},
		{[3]int32{2, 12, 10}, Block{BlockId: 765, Data: 5}, true},
		{[3]int32{3, 12, 10}, Block{BlockId: 451, Data: 11}, false},
		{[3]int32{4, 7, 9}, Block{BlockId: 761, Data: 5}, false},
		{[3]int32{4, 7, 9}, Block{}, false},
	}
	for n, test := range tests {
		section.SetBlock(test.xyz[0], test.xyz[1], test.xyz[2], &test.Block)
		if b := section.GetBlock(test.xyz[0], test.xyz[1], test.xyz[2]); !test.Block.Equal(b) {
			t.Errorf("test %d failed\nExpecting: %s\nGot: %s", n+1, test.Block.String(), b.String())
		}
	}
	for n, test := range tests {
		if test.recheck {
			if b := section.GetBlock(test.xyz[0], test.xyz[1], test.xyz[2]); !test.Block.Equal(b) {
				t.Errorf("retest %d failed\nExpecting: %s\nGot: %s", n+1, test.Block.String(), b.String())
			}
		}
	}
}

func b2i(b byte) int8 {
	return int8(b)
}

func i2b(b int8) byte {
	return byte(b)
}
