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

import "github.com/MJKWoolnough/equaler"

const (
	Biome_Ocean Biome = iota
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
	Biome_JungleEdge
	Biome_DeepOcean
	Biome_StoneBeach
	Biome_ColdBeach
	Biome_BirchForest
	Biome_BirchForestHills
	Biome_RoofedForest
	Biome_ColdTaiga
	Biome_ColdTaigaHills
	Biome_MegaTaiga
	Biome_MegaTaigaHills
	Biome_ExtremeHillsPlus
	Biome_Savanna
	Biome_SavannaPlateau
	Biome_Mesa
	Biome_MesaPlateauF
	Biome_MesaPlateau
)

const (
	Biome_SunflowerPlains Biome = iota + 129
	Biome_DeserM
	Biome_ExtremeHillsM
	Biome_FlowerForest
	Biome_TaigaM
	Biome_SwamplandM
)
const (
	Biome_IcePlainsSpikes Biome = 140
	Biome_JungleM         Biome = 149
	Biome_JungleEdgeM     Biome = 151
	Biome_Auto            Biome = 255
)
const (
	Biome_BirchForestM Biome = iota + 155
	Biome_BirchForestHillsM
	Biome_RoofedForestM
	Biome_ColdTaigaM
)
const (
	Biome_MegaSpruceTaiga Biome = iota + 160
	Biome_MegaSpruceTaigaHills
	Biome_ExtremeHillsPlusM
	Biome_SavannaM
	Biome_SavannaPlateauM
	Biome_MesaBryce
	Biome_MesaPlateauFM
	Biome_MesaPlateauM
)

type Biome uint8

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
	case Biome_JungleEdge:
		return "Jungle Edge"
	case Biome_DeepOcean:
		return "Deep Ocean"
	case Biome_StoneBeach:
		return "Stone Beach"
	case Biome_ColdBeach:
		return "Cold Beach"
	case Biome_BirchForest:
		return "Birch Forest"
	case Biome_BirchForestHills:
		return "Birch Forest Hills"
	case Biome_RoofedForest:
		return "Roofed Forest"
	case Biome_ColdTaiga:
		return "Cold Taiga"
	case Biome_ColdTaigaHills:
		return "Cold Taiga Hills"
	case Biome_MegaTaiga:
		return "Mega Taiga"
	case Biome_MegaTaigaHills:
		return "Mega Taiga Hills"
	case Biome_ExtremeHillsPlus:
		return "Extreme Hills+"
	case Biome_Savanna:
		return "Savanna"
	case Biome_SavannaPlateau:
		return "Savanna Plateau"
	case Biome_Mesa:
		return "Mesa"
	case Biome_MesaPlateauF:
		return "Mesa Plateau F"
	case Biome_MesaPlateau:
		return "Mesa Plateau"
	case Biome_SunflowerPlains:
		return "Sunflower Plains"
	case Biome_DeserM:
		return "Desert M"
	case Biome_ExtremeHillsM:
		return "Extreme Hills M"
	case Biome_FlowerForest:
		return "Flower Forest"
	case Biome_TaigaM:
		return "Taiga M"
	case Biome_SwamplandM:
		return "Swampland M"
	case Biome_IcePlainsSpikes:
		return "Ice Plains Spikes"
	case Biome_JungleM:
		return "Jungle M"
	case Biome_JungleEdgeM:
		return "Jungle Edge M"
	case Biome_BirchForestM:
		return "BirchForestM"
	case Biome_BirchForestHillsM:
		return "BirchForestHillsM"
	case Biome_RoofedForestM:
		return "Roofed Forest M"
	case Biome_ColdTaigaM:
		return "Cold Taiga M"
	case Biome_MegaSpruceTaiga:
		return "Mega Spruce Taiga"
	case Biome_MegaSpruceTaigaHills:
		return "Mega Spruce Taiga Hills"
	case Biome_ExtremeHillsPlusM:
		return "Extreme Hills Plus M"
	case Biome_SavannaM:
		return "Savanna M"
	case Biome_SavannaPlateauM:
		return "Savanna Plateau M"
	case Biome_MesaBryce:
		return "Mesa Bryce"
	case Biome_MesaPlateauFM:
		return "Mesa Plateau F M"
	case Biome_MesaPlateauM:
		return "Mesa Plateau M"
	case Biome_Auto:
		return "Auto"
	}
	place := 0
	for n := b; n > 0; n /= 10 {
		place++
	}
	digits := make([]byte, place)
	for n := b; n > 0; n /= 10 {
		place--
		digits[place] = '0' + byte(n%10)
	}
	return "Unrecognised Biome ID - " + string(digits)
}
