// Copyright (c) 2013 - Michael Woolnough <michael.woolnough@gmail.com>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
		data &= 15
	} else {
		data >>= 4
	}
	return data
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
	return
}

type section struct {
	section    *nbt.Compound
	blocks     *nbt.ByteArray
	add        *nbt.ByteArray
	data       *nbt.ByteArray
	blockLight *nbt.ByteArray
	skyLight   *nbt.ByteArray
}

func newSection(y int32) *section {
	s := new(section)
	s.blocks = nbt.NewByteArray(make([]int8, 4096))
	s.add = nbt.NewByteArray(make([]int8, 2048))
	s.data = nbt.NewByteArray(make([]int8, 2048))
	s.blockLight = nbt.NewByteArray(make([]int8, 2048))
	s.skyLight = nbt.NewByteArray(make([]int8, 2048))
	s.section = nbt.NewCompound([]nbt.Tag{
		nbt.NewTag("Blocks", s.blocks),
		nbt.NewTag("Add", s.add),
		nbt.NewTag("Data", s.data),
		nbt.NewTag("BlockLight", s.blockLight),
		nbt.NewTag("SkyLight", s.skyLight),
		nbt.NewTag("Y", nbt.NewByte(int8(y>>4))),
	})
	return s
}

func loadSection(c *nbt.Compound) (*section, error) {
	s := new(section)
	s.section = c
	blocks := c.Get("Blocks")
	if blocks == nil {
		return nil, &MissingTagError{"[SECTION]->Blocks"}
	} else if blocks.TagId() != nbt.Tag_ByteArray {
		return nil, &WrongTypeError{"Blocks", nbt.Tag_ByteArray, blocks.TagId()}
	}
	s.blocks = blocks.Data().(*nbt.ByteArray)
	if len(*s.blocks) != 4096 {
		return nil, &OOB{}
	}
	add := c.Get("Add")
	if add != nil {
		if add.TagId() != nbt.Tag_ByteArray {
			return nil, &WrongTypeError{"Add", nbt.Tag_ByteArray, add.TagId()}
		}
		s.add = add.Data().(*nbt.ByteArray)
	} else {
		s.add = nbt.NewByteArray(make([]int8, 2048))
		c.Set(nbt.NewTag("Add", s.add))
	}
	if len(*s.add) != 2048 {
		return nil, &OOB{}
	}
	data := c.Get("Data")
	if data == nil {
		return nil, &MissingTagError{"[SECTION]->Data"}
	} else if data.TagId() != nbt.Tag_ByteArray {
		return nil, &WrongTypeError{"Data", nbt.Tag_ByteArray, data.TagId()}
	}
	s.data = data.Data().(*nbt.ByteArray)
	if len(*s.data) != 2048 {
		return nil, &OOB{}
	}
	blockLight := c.Get("BlockLight")
	if blockLight == nil {
		return nil, &MissingTagError{"[SECTION]->BlockLight"}
	} else if blockLight.TagId() != nbt.Tag_ByteArray {
		return nil, &WrongTypeError{"BlockLight", nbt.Tag_ByteArray, blockLight.TagId()}
	}
	s.blockLight = blockLight.Data().(*nbt.ByteArray)
	if len(*s.blockLight) != 2048 {
		return nil, &OOB{}
	}
	skyLight := c.Get("SkyLight")
	if skyLight == nil {
		return nil, &MissingTagError{"[SECTION]->SkyLight"}
	} else if skyLight.TagId() != nbt.Tag_ByteArray {
		return nil, &WrongTypeError{"SkyLight", nbt.Tag_ByteArray, skyLight.TagId()}
	}
	s.skyLight = skyLight.Data().(*nbt.ByteArray)
	if len(*s.skyLight) != 2048 {
		return nil, &OOB{}
	}
	y := c.Get("Y")
	if blockLight == nil {
		return nil, &MissingTagError{"[SECTION]->Y"}
	} else if y.TagId() != nbt.Tag_Byte {
		return nil, &WrongTypeError{"Y", nbt.Tag_Byte, y.TagId()}
	}
	return s, nil
}

func (s *section) GetBlock(x, y, z int32) *Block {
	return &Block{
		BlockId: uint16(getNibble(*s.add, x, y, z))<<8 | uint16(byte((*s.blocks)[yzx(x, y, z)])),
		Data:    getNibble(*s.data, x, y, z),
	}
}

func (s *section) SetBlock(x, y, z int32, b *Block) {
	(*s.blocks)[yzx(x, y, z)] = int8(b.BlockId & 255)
	setNibble(*s.add, x, y, z, byte(b.BlockId>>8))
	setNibble(*s.data, x, y, z, byte(b.Data))
}

func (s *section) GetOpacity(x, y, z int32) uint8 {
	return s.GetBlock(x, y, z).Opacity()
}

func (s *section) GetBlockLight(x, y, z int32) uint8 {
	return getNibble(*s.blockLight, x, y, z)
}

func (s *section) SetBlockLight(x, y, z int32, l uint8) {
	setNibble(*s.blockLight, x, y, z, l)
}

func (s *section) GetSkyLight(x, y, z int32) uint8 {
	return getNibble(*s.skyLight, x, y, z)
}

func (s *section) SetSkyLight(x, y, z int32, l uint8) {
	setNibble(*s.skyLight, x, y, z, l)
}

func (s *section) SetY(y int32) {
	s.section.Set(nbt.NewTag("Y", nbt.NewByte(int8(y>>4))))
}

func (s *section) IsEmpty() bool {
	for _, b := range *s.blocks {
		if b != 0 {
			return false
		}
	}
	return true
}
