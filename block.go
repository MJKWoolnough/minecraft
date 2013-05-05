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
	"github.com/MJKWoolnough/equaler"
	"github.com/MJKWoolnough/minecraft/nbtparser"
	"strconv"
)

// Block allows access to the data of a minecraft block.
type Block interface {
	BlockId() uint8
	Add() uint8
	Data() uint8
	Opacity() uint8
	IsLiquid() bool
	HasMetadata() bool
	GetMetadata() []nbtparser.NBTTag
	SetMetadata([]nbtparser.NBTTag)
	Tick(bool)
	ToTick() bool
	equaler.Equaler
}

type block struct {
	blockId  uint8
	add      uint8
	data     uint8
	metadata []nbtparser.NBTTag
	tick     bool
}

func (b *block) BlockId() uint8 {
	if b == nil {
		return 0
	}
	return b.blockId
}

// Add returns the extended block Id data
func (b *block) Add() uint8 {
	if b == nil {
		return 0
	}
	return b.add
}

// Data returns the additional data, also known as damage.
func (b *block) Data() uint8 {
	if b == nil {
		return 0
	}
	return b.data
}

func (b *block) Equal(e equaler.Equaler) bool {
	if c, ok := e.(*block); ok {
		if b.blockId == c.blockId && b.add == c.add && b.data == c.data && b.tick == c.tick {
			if len(b.metadata) > 0 {
				if len(c.metadata) > 0 {
					for _, v := range b.metadata {
						name := v.Name()
						found := false
						for _, w := range c.metadata {
							if w.Name() == name {
								if !v.Equal(w) {
									return false
								}
								found = true
							}
						}
						if !found {
							return false
						}
					}
					return true
				}
			} else {
				return len(c.metadata) == 0
			}
		}
	}
	return false
}

// Opacity returns how much light is blocked by this block.
func (b *block) Opacity() uint8 {
	if b == nil {
		return 0
	}
	if b.blockId == 8 || b.blockId == 9 {
		return 3
	}
	blockId := uint16(b.blockId) | (uint16(b.add) << 8)
	for i := 0; i < len(transparentBlocks); i++ {
		if transparentBlocks[i] == blockId {
			return 1
		}
	}
	return 16
}

func (b *block) IsLiquid() bool {
	return b.blockId == 8 || b.blockId == 9 || b.blockId == 10 || b.blockId == 11
}

func (b *block) HasMetadata() bool {
	if b == nil || b.metadata == nil || len(b.metadata) == 0 {
		return false
	}
	return true
}

func (b *block) GetMetadata() []nbtparser.NBTTag {
	if b == nil || b.metadata == nil {
		return nil
	}
	a := make([]nbtparser.NBTTag, len(b.metadata))
	for i, j := range b.metadata {
		a[i] = j.Copy()
	}
	return a
}

func (b *block) SetMetadata(data []nbtparser.NBTTag) {
	metadata := make([]nbtparser.NBTTag, 0)
	for i := 0; i < len(data); i++ {
		name := data[i].Name()
		if name == "x" || name == "y" || name == "z" || data[i].TagEnd() != nil {
			continue
		}
		metadata = append(metadata, data[i])
	}
	if len(metadata) > 0 {
		b.metadata = metadata
	} else {
		b.metadata = nil
	}
}

func (b *block) String() string {
	toRet := "Block ID: " + strconv.Itoa(int(b.blockId)) + "\n"
	toRet += "Add Data: " + strconv.Itoa(int(b.add)) + "\n"
	toRet += "Data: " + strconv.Itoa(int(b.data)) + "\n"
	if b.metadata != nil && len(b.metadata) != 0 {
		toRet += "Metadata:\n"
		for i := 0; i < len(b.metadata); i++ {
			toRet += "	" + b.metadata[i].String() + "\n"
		}
	}

	return toRet
}

// Tick sets the block to be updated, useful for growing planted saplings.
func (b *block) Tick(t bool) {
	b.tick = t
}

func (b *block) ToTick() bool {
	return b.tick
}

func NewBlock(blockId, add, data uint8) Block {
	return &block{blockId, add, data, make([]nbtparser.NBTTag, 0), false}
}
