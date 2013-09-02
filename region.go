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

type region struct {
	chunks  [1024]*chunk
	changed [1024]bool
}

func newRegion(path Path) *region {
	return new(region)
}

func (r *region) GetBlock(path Path, x, y, z int32) (*Block, error) {
	c, err := r.getChunk(path, x, z, false)
	if err != nil {
		return nil, err
	} else if c == nil {
		return &Block{}, nil
	}
	return c.GetBlock(x, y, z), nil
}

func (r *region) SetBlock(path Path, x, y, z int32, block *Block) error {
	c, err := r.getChunk(path, x, z, true)
	if err != nil {
		return err
	}
	c.SetBlock(x, y, z, block)
	return nil
}

func (r *region) GetBiome(path Path, x, z int32) (Biome, error) {
	c, err := r.getChunk(path, x, z, false)
	if err != nil {
		return -1, err
	} else if c == nil {
		return Biome(1), nil
	}
	return c.GetBiome(x, z), nil
}

func (r *region) SetBiome(path Path, x, z int32, biome Biome) error {
	c, err := r.getChunk(path, x, z, true)
	if err != nil {
		return err
	}
	c.SetBiome(x, z, biome)
	return nil
}

func (r *region) GetOpacity(path Path, x, y, z int32) (uint8, error) {
	c, err := r.getChunk(path, x, z, false)
	if err != nil {
		return 0, err
	}
	return c.GetOpacity(x, y, z), nil
}

func (r *region) GetHeight(path Path, x, z int32) (int32, error) {
	c, err := r.getChunk(path, x, z, false)
	if err != nil {
		return 0, err
	}
	return c.GetHeight(x, z), nil
}

func (r *region) GetBlockLight(path Path, x, y, z int32) (uint8, error) {
	c, err := r.getChunk(path, x, z, false)
	if err != nil {
		return 0, err
	}
	return c.GetBlockLight(x, y, z), nil
}

func (r *region) SetBlockLight(path Path, x, y, z int32, l uint8) error {
	c, err := r.getChunk(path, x, z, true)
	if err != nil {
		return err
	}
	c.SetBlockLight(x, y, z, l)
	return nil
}

func (r *region) GetSkyLight(path Path, x, y, z int32) (uint8, error) {
	c, err := r.getChunk(path, x, z, false)
	if err != nil {
		return 0, err
	}
	return c.GetSkyLight(x, y, z), nil
}

func (r *region) SetSkyLight(path Path, x, y, z int32, l uint8) error {
	c, err := r.getChunk(path, x, z, true)
	if err != nil {
		return err
	}
	c.SetSkyLight(x, y, z, l)
	return nil
}

func (r *region) Save(path Path) error {
	toSave := make([]nbt.Tag, 0)
	for i := 0; i < 1024; i++ {
		if r.changed[i] {
			r.changed[i] = false
			toSave = append(toSave, r.chunks[i].GetNBT())
		}
	}
	if len(toSave) > 0 {
		return path.SetChunk(toSave...)
	}
	return nil
}

func (r *region) getChunk(path Path, x, z int32, create bool) (*chunk, error) {
	chunkNum := (((z >> 4) & 31) << 5) | ((x >> 4) & 31)
	if r.chunks[chunkNum] == nil {
		chunkData, err := path.GetChunk(x, z)
		if err != nil {
			return nil, err
		}
		if chunkData != nil {
			chunk, err := newChunk(x, z, chunkData)
			if err != nil {
				return nil, err
			}
			r.chunks[chunkNum] = chunk
		} else if create {
			r.chunks[chunkNum], _ = newChunk(x, z, nil)
		}
	}
	if create {
		r.changed[chunkNum] = true
	}
	return r.chunks[chunkNum], nil
}
