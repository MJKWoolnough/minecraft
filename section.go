package minecraft

import (
	"github.com/MJKWoolnough/minecraft/nbt"
)

func yzx(x, y, z int32) uint32 {
	return (uint32(y&15) << 8) | (uint32(z&15) << 4) | uint32(x&15)
}

func getNibble(arr nbt.ByteArray, x, y, z int32) byte {
	coord := yzx(x, y, z)
	data := byte(arr[coord>>1])
	if coord&1 == 0 {
		return data & 15
	}
	return data >> 4
}

func setNibble(arr nbt.ByteArray, x, y, z int32, data byte) {
	coord := yzx(x, y, z)
	oldData := byte(arr[coord>>1])
	if coord&1 == 0 {
		oldData = oldData&240 | data&15
	} else {
		oldData = oldData&15 | data<<4
	}
	arr[coord>>1] = int8(oldData)
}

type section struct {
	section    nbt.Compound
	blocks     nbt.ByteArray
	add        nbt.ByteArray
	data       nbt.ByteArray
	blockLight nbt.ByteArray
	skyLight   nbt.ByteArray
}

func newSection(y int32) *section {
	s := new(section)
	s.blocks = make(nbt.ByteArray, 4096)
	s.add = make(nbt.ByteArray, 2048)
	s.data = make(nbt.ByteArray, 2048)
	s.blockLight = make(nbt.ByteArray, 2048)
	s.skyLight = make(nbt.ByteArray, 2048)
	for i := 0; i < 2048; i++ {
		s.skyLight[i] = -1
	}
	s.section = nbt.Compound{
		nbt.NewTag("Blocks", s.blocks),
		nbt.NewTag("Add", s.add),
		nbt.NewTag("Data", s.data),
		nbt.NewTag("BlockLight", s.blockLight),
		nbt.NewTag("SkyLight", s.skyLight),
		nbt.NewTag("Y", nbt.Byte(y>>4)),
	}
	return s
}

func loadSection(c nbt.Compound) (*section, error) {
	s := new(section)
	s.section = c
	blocks := c.Get("Blocks")
	if blocks.TagID() == 0 {
		return nil, MissingTagError{"[SECTION]->Blocks"}
	} else if blocks.TagID() != nbt.TagByteArray {
		return nil, WrongTypeError{"Blocks", nbt.TagByteArray, blocks.TagID()}
	}
	s.blocks = blocks.Data().(nbt.ByteArray)
	if len(s.blocks) != 4096 {
		return nil, ErrOOB
	}
	add := c.Get("Add")
	if add.TagID() != 0 {
		if add.TagID() != nbt.TagByteArray {
			return nil, WrongTypeError{"Add", nbt.TagByteArray, add.TagID()}
		}
		s.add = add.Data().(nbt.ByteArray)
	} else {
		s.add = make(nbt.ByteArray, 2048)
		c.Set(nbt.NewTag("Add", s.add))
	}
	if len(s.add) != 2048 {
		return nil, ErrOOB
	}
	data := c.Get("Data")
	if data.TagID() == 0 {
		return nil, MissingTagError{"[SECTION]->Data"}
	} else if data.TagID() != nbt.TagByteArray {
		return nil, WrongTypeError{"Data", nbt.TagByteArray, data.TagID()}
	}
	s.data = data.Data().(nbt.ByteArray)
	if len(s.data) != 2048 {
		return nil, ErrOOB
	}
	blockLight := c.Get("BlockLight")
	if blockLight.TagID() == 0 {
		return nil, MissingTagError{"[SECTION]->BlockLight"}
	} else if blockLight.TagID() != nbt.TagByteArray {
		return nil, WrongTypeError{"BlockLight", nbt.TagByteArray, blockLight.TagID()}
	}
	s.blockLight = blockLight.Data().(nbt.ByteArray)
	if len(s.blockLight) != 2048 {
		return nil, ErrOOB
	}
	skyLight := c.Get("SkyLight")
	if skyLight.TagID() == 0 {
		return nil, MissingTagError{"[SECTION]->SkyLight"}
	} else if skyLight.TagID() != nbt.TagByteArray {
		return nil, WrongTypeError{"SkyLight", nbt.TagByteArray, skyLight.TagID()}
	}
	s.skyLight = skyLight.Data().(nbt.ByteArray)
	if len(s.skyLight) != 2048 {
		return nil, ErrOOB
	}
	y := c.Get("Y")
	if y.TagID() == 0 {
		return nil, MissingTagError{"[SECTION]->Y"}
	} else if y.TagID() != nbt.TagByte {
		return nil, WrongTypeError{"Y", nbt.TagByte, y.TagID()}
	}
	return s, nil
}

func (s *section) GetBlock(x, y, z int32) Block {
	return Block{
		BlockID: uint16(getNibble(s.add, x, y, z))<<8 | uint16(byte(s.blocks[yzx(x, y, z)])),
		Data:    getNibble(s.data, x, y, z),
	}
}

func (s *section) SetBlock(x, y, z int32, b Block) {
	s.blocks[yzx(x, y, z)] = int8(b.BlockID & 255)
	setNibble(s.add, x, y, z, byte(b.BlockID>>8))
	setNibble(s.data, x, y, z, byte(b.Data))
}

func (s *section) GetOpacity(x, y, z int32) uint8 {
	return s.GetBlock(x, y, z).Opacity()
}

func (s *section) GetBlockLight(x, y, z int32) uint8 {
	return getNibble(s.blockLight, x, y, z)
}

func (s *section) SetBlockLight(x, y, z int32, l uint8) {
	setNibble(s.blockLight, x, y, z, l)
}

func (s *section) GetSkyLight(x, y, z int32) uint8 {
	return getNibble(s.skyLight, x, y, z)
}

func (s *section) SetSkyLight(x, y, z int32, l uint8) {
	setNibble(s.skyLight, x, y, z, l)
}

func (s *section) SetY(y int32) {
	s.section.Set(nbt.NewTag("Y", nbt.Byte(y>>4)))
}
