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
	"fmt"
	"github.com/MJKWoolnough/equaler"
	"github.com/MJKWoolnough/minecraft/nbt"
)

type Block struct {
	BlockId  uint16
	Data     uint8
	metadata []nbt.Tag
	Tick     bool
}

func (b Block) Equal(e equaler.Equaler) bool {
	c, ok := e.(*Block)
	if !ok {
		if d, ok := e.(Block); ok {
			c = &d
		}
	}
	if c != nil {
		if b.BlockId == c.BlockId && b.Data == c.Data && b.Tick == c.Tick {
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
func (b Block) Opacity() uint8 {
	if b.BlockId == 8 || b.BlockId == 9 {
		return 3
	}
	for i := 0; i < len(transparentBlocks); i++ {
		if transparentBlocks[i] == b.BlockId {
			return 1
		}
	}
	return 16
}

func (b Block) IsLiquid() bool {
	return b.BlockId == 8 || b.BlockId == 9 || b.BlockId == 10 || b.BlockId == 11
}

func (b Block) HasMetadata() bool {
	if b.metadata == nil || len(b.metadata) == 0 {
		return false
	}
	return true
}

func (b Block) GetMetadata() []nbt.Tag {
	if b.metadata == nil {
		return nil
	}
	a := make([]nbt.Tag, len(b.metadata))
	for i, j := range b.metadata {
		a[i] = j.Copy()
	}
	return a
}

func (b *Block) SetMetadata(data []nbt.Tag) {
	metadata := make([]nbt.Tag, 0)
	for i := 0; i < len(data); i++ {
		name := data[i].Name()
		if name == "x" || name == "y" || name == "z" {
			continue
		} else if data[i].TagId() == nbt.Tag_End {
			break
		}
		metadata = append(metadata, data[i].Copy())
	}
	if len(metadata) > 0 {
		b.metadata = metadata
	} else {
		b.metadata = nil
	}
}

func (b Block) String() string {
	toRet := fmt.Sprintf("Block ID: %d\nData: %d\n", b.BlockId, b.Data)
	if b.metadata != nil && len(b.metadata) != 0 {
		toRet += "Metadata:\n"
		for i := 0; i < len(b.metadata); i++ {
			toRet += "	" + b.metadata[i].String() + "\n"
		}
	}
	if b.Tick {
		toRet += "Tick: on"
	} else {
		toRet += "Tick: off"
	}
	return toRet
}
