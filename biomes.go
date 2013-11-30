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

const (
	Biome_Auto Biome = iota - 1
	Biome_Ocean
	Biome_Plains
	Biome_Desert
	Biome_ExtremeHills
	Biome_Forest
	Biome_Taiga
	Biome_Swampland
	Biome_River
	Biome_Hell
	Biome_Sky
	Biome_FrozenOcean
	Biome_FrozenRiver
	Biome_IcePlains
	Biome_IceMountains
	Biome_MushroomIsland
	Biome_MushroomIslandShore
	Biome_Beach
	Biome_DesertHills
	Biome_ForestHills
	Biome_TaigaHills
	Biome_ExtremeHillsEdge
	Biome_Jungle
	Biome_JungleHills
)

type Biome int8

func (b Biome) Equal(e equaler.Equaler) bool {
	if c, ok := e.(*Biome); ok {
		return b == *c
	} else if c, ok := e.(Biome); ok {
		return b == c
	}
	return false
}

func (b Biome) String() string {
	switch b {
	case Biome_Auto:
		return "Auto"
	case Biome_Ocean:
		return "Ocean"
	case Biome_Plains:
		return "Plains"
	case Biome_Desert:
		return "Desert"
	case Biome_ExtremeHills:
		return "Extreme Hills"
	case Biome_Forest:
		return "Forest"
	case Biome_Taiga:
		return "Taiga"
	case Biome_Swampland:
		return "Swampland"
	case Biome_River:
		return "River"
	case Biome_Hell:
		return "Hell"
	case Biome_Sky:
		return "Sky"
	case Biome_FrozenOcean:
		return "Frozen Ocean"
	case Biome_FrozenRiver:
		return "Frozen River"
	case Biome_IcePlains:
		return "Ice Plains"
	case Biome_IceMountains:
		return "Ice Mountains"
	case Biome_MushroomIsland:
		return "Mushroom Island"
	case Biome_MushroomIslandShore:
		return "Mushroom Island Shore"
	case Biome_Beach:
		return "Beach"
	case Biome_DesertHills:
		return "Desert Hills"
	case Biome_ForestHills:
		return "Forest Hills"
	case Biome_TaigaHills:
		return "Taiga Hills"
	case Biome_ExtremeHillsEdge:
		return "Extreme Hills Edge"
	case Biome_Jungle:
		return "Jungle"
	case Biome_JungleHills:
		return "Jungle Hills"
	}
	return "Unrecognised Biome ID - " + strconv.Itoa(int(b))
}
