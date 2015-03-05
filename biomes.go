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

// Biome constants
const (
	Ocean                Biome = 0
	Plains               Biome = 1
	Desert               Biome = 2
	ExtremeHills         Biome = 3
	Forest               Biome = 4
	Taiga                Biome = 5
	Swampland            Biome = 6
	River                Biome = 7
	Hell                 Biome = 8
	Sky                  Biome = 9
	FrozenOcean          Biome = 10
	FrozenRiver          Biome = 11
	IcePlains            Biome = 12
	IceMountains         Biome = 13
	MushroomIsland       Biome = 14
	MushroomIslandShore  Biome = 15
	Beach                Biome = 16
	DesertHills          Biome = 17
	ForestHills          Biome = 18
	TaigaHills           Biome = 19
	ExtremeHillsEdge     Biome = 20
	Jungle               Biome = 21
	JungleHills          Biome = 22
	JungleEdge           Biome = 23
	DeepOcean            Biome = 23
	StoneBeach           Biome = 24
	ColdBeach            Biome = 25
	BirchForest          Biome = 26
	BirchForestHills     Biome = 27
	RoofedForest         Biome = 28
	ColdTaiga            Biome = 29
	ColdTaigaHills       Biome = 30
	MegaTaiga            Biome = 31
	MegaTaigaHills       Biome = 32
	ExtremeHillsPlus     Biome = 33
	Savanna              Biome = 34
	SavannaPlateau       Biome = 35
	Mesa                 Biome = 36
	MesaPlateauF         Biome = 37
	MesaPlateau          Biome = 38
	SunflowerPlains      Biome = 129
	DeserM               Biome = 130
	ExtremeHillsM        Biome = 131
	FlowerForest         Biome = 132
	TaigaM               Biome = 133
	SwamplandM           Biome = 134
	IcePlainsSpikes      Biome = 140
	JungleM              Biome = 149
	JungleEdgeM          Biome = 151
	BirchForestM         Biome = 155
	BirchForestHillsM    Biome = 156
	RoofedForestM        Biome = 157
	ColdTaigaM           Biome = 158
	MegaSpruceTaiga      Biome = 160
	MegaSpruceTaigaHills Biome = 161
	ExtremeHillsPlusM    Biome = 162
	SavannaM             Biome = 163
	SavannaPlateauM      Biome = 164
	MesaBryce            Biome = 165
	MesaPlateauFM        Biome = 166
	MesaPlateauM         Biome = 167
	AutoBiome            Biome = 255
)

// Biome is a convenience type for biomes
type Biome uint8

// Equal is an implementation of the equaler.Equaler interface
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
