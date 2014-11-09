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
	Ocean Biome = iota
	Plains
	Desert
	ExtremeHills
	Forest
	Taiga
	Swampland
	River
	Hell
	Sky
	FrozenOcean
	FrozenRiver
	IcePlains
	IceMountains
	MushroomIsland
	MushroomIslandShore
	Beach
	DesertHills
	ForestHills
	TaigaHills
	ExtremeHillsEdge
	Jungle
	JungleHills
	JungleEdge
	DeepOcean
	StoneBeach
	ColdBeach
	BirchForest
	BirchForestHills
	RoofedForest
	ColdTaiga
	ColdTaigaHills
	MegaTaiga
	MegaTaigaHills
	ExtremeHillsPlus
	Savanna
	SavannaPlateau
	Mesa
	MesaPlateauF
	MesaPlateau
)

const (
	SunflowerPlains Biome = iota + 129
	DeserM
	ExtremeHillsM
	FlowerForest
	TaigaM
	SwamplandM
)
const (
	IcePlainsSpikes Biome = 140
	JungleM         Biome = 149
	JungleEdgeM     Biome = 151
	AutoBiome       Biome = 255
)
const (
	BirchForestM Biome = iota + 155
	BirchForestHillsM
	RoofedForestM
	ColdTaigaM
)
const (
	MegaSpruceTaiga Biome = iota + 160
	MegaSpruceTaigaHills
	ExtremeHillsPlusM
	SavannaM
	SavannaPlateauM
	MesaBryce
	MesaPlateauFM
	MesaPlateauM
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
	case Ocean:
		return "Ocean"
	case Plains:
		return "Plains"
	case Desert:
		return "Desert"
	case ExtremeHills:
		return "Extreme Hills"
	case Forest:
		return "Forest"
	case Taiga:
		return "Taiga"
	case Swampland:
		return "Swampland"
	case River:
		return "River"
	case Hell:
		return "Hell"
	case Sky:
		return "Sky"
	case FrozenOcean:
		return "Frozen Ocean"
	case FrozenRiver:
		return "Frozen River"
	case IcePlains:
		return "Ice Plains"
	case IceMountains:
		return "Ice Mountains"
	case MushroomIsland:
		return "Mushroom Island"
	case MushroomIslandShore:
		return "Mushroom Island Shore"
	case Beach:
		return "Beach"
	case DesertHills:
		return "Desert Hills"
	case ForestHills:
		return "Forest Hills"
	case TaigaHills:
		return "Taiga Hills"
	case ExtremeHillsEdge:
		return "Extreme Hills Edge"
	case Jungle:
		return "Jungle"
	case JungleHills:
		return "Jungle Hills"
	case JungleEdge:
		return "Jungle Edge"
	case DeepOcean:
		return "Deep Ocean"
	case StoneBeach:
		return "Stone Beach"
	case ColdBeach:
		return "Cold Beach"
	case BirchForest:
		return "Birch Forest"
	case BirchForestHills:
		return "Birch Forest Hills"
	case RoofedForest:
		return "Roofed Forest"
	case ColdTaiga:
		return "Cold Taiga"
	case ColdTaigaHills:
		return "Cold Taiga Hills"
	case MegaTaiga:
		return "Mega Taiga"
	case MegaTaigaHills:
		return "Mega Taiga Hills"
	case ExtremeHillsPlus:
		return "Extreme Hills+"
	case Savanna:
		return "Savanna"
	case SavannaPlateau:
		return "Savanna Plateau"
	case Mesa:
		return "Mesa"
	case MesaPlateauF:
		return "Mesa Plateau F"
	case MesaPlateau:
		return "Mesa Plateau"
	case SunflowerPlains:
		return "Sunflower Plains"
	case DeserM:
		return "Desert M"
	case ExtremeHillsM:
		return "Extreme Hills M"
	case FlowerForest:
		return "Flower Forest"
	case TaigaM:
		return "Taiga M"
	case SwamplandM:
		return "Swampland M"
	case IcePlainsSpikes:
		return "Ice Plains Spikes"
	case JungleM:
		return "Jungle M"
	case JungleEdgeM:
		return "Jungle Edge M"
	case BirchForestM:
		return "BirchForestM"
	case BirchForestHillsM:
		return "BirchForestHillsM"
	case RoofedForestM:
		return "Roofed Forest M"
	case ColdTaigaM:
		return "Cold Taiga M"
	case MegaSpruceTaiga:
		return "Mega Spruce Taiga"
	case MegaSpruceTaigaHills:
		return "Mega Spruce Taiga Hills"
	case ExtremeHillsPlusM:
		return "Extreme Hills Plus M"
	case SavannaM:
		return "Savanna M"
	case SavannaPlateauM:
		return "Savanna Plateau M"
	case MesaBryce:
		return "Mesa Bryce"
	case MesaPlateauFM:
		return "Mesa Plateau F M"
	case MesaPlateauM:
		return "Mesa Plateau M"
	case AutoBiome:
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
