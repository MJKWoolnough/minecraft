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
