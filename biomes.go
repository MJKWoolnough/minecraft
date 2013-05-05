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
	"strconv"
)

// Needs Implementation
type Biome byte

func (b *Biome) Equal(e equaler.Equaler) bool {
	if c, ok := e.(*Biome); ok {
		return b == c
	}
	return false
}

func (b *Biome) String() string {
	switch *b {
	case 0:
		return "Ocean"
	case 1:
		return "Plains"
	case 2:
		return "Desert"
	case 3:
		return "Extreme Hills"
	case 4:
		return "Forest"
	case 5:
		return "Taiga"
	case 6:
		return "Swampland"
	case 7:
		return "River"
	case 8:
		return "Hell"
	case 9:
		return "Sky"
	case 10:
		return "Frozen Ocean"
	case 11:
		return "Frozen River"
	case 12:
		return "Ice Plains"
	case 13:
		return "Ice Mountains"
	case 14:
		return "Mushroom Island"
	case 15:
		return "Mushroom Island Shore"
	case 16:
		return "Beach"
	case 17:
		return "Desert Hills"
	case 18:
		return "Forest Hills"
	case 19:
		return "Taiga Hills"
	case 20:
		return "Extreme Hills Edge"
	case 21:
		return "Jungle"
	case 22:
		return "Jungle Hills"
	}
	return "Unrecognised Biome ID - " + strconv.Itoa(int(*b))
}

func NewBiome(biomeId uint8) Biome {
	return Biome(biomeId)
}
