package manipulators

import (
	"github.com/MJKWoolnough/minecraft"
	"github.com/MJKWoolnough/minecraft/nbt"
)

type Default struct{}

func (Default) Rotate90(block minecraft.Block) minecraft.Block {
	return block
}

func (Default) Rotate180(block minecraft.Block) minecraft.Block {
	return block
}

func (Default) Rotate270(block minecraft.Block) minecraft.Block {
	return block
}

func (Default) MirrorX(block minecraft.Block) minecraft.Block {
	return block
}

func (Default) MirrorZ(block minecraft.Block) minecraft.Block {
	return block
}

type Log struct {
	Default
}

func (Log) Rotate90(block minecraft.Block) minecraft.Block {
	switch block.Data & 12 {
	case 4:
		block.Data += 4
	case 8:
		block.Data -= 4
	}
	return block
}

func (l Log) Rotate270(block minecraft.Block) minecraft.Block {
	return l.Rotate90(block)
}

type Bits struct {
	Bitmask, North, East, South, West uint8
}

func NewBits(bm, north, east, south, west uint8) Bits {
	return Bits{
		bm, north, east, south, west,
	}
}

func (b Bits) Rotate90(block minecraft.Block) minecraft.Block {
	data := block.Data & b.Bitmask
	block.Data &^= b.Bitmask
	switch data {
	case b.North:
		block.Data |= b.East
	case b.East:
		block.Data |= b.South
	case b.South:
		block.Data |= b.West
	case b.West:
		block.Data |= b.North
	default:
		block.Data |= data
	}
	return block
}

func (b Bits) Rotate180(block minecraft.Block) minecraft.Block {
	data := block.Data & b.Bitmask
	block.Data &^= b.Bitmask
	switch data {
	case b.North:
		block.Data |= b.South
	case b.East:
		block.Data |= b.West
	case b.South:
		block.Data |= b.North
	case b.West:
		block.Data |= b.East
	default:
		block.Data |= data
	}
	return block
}

func (b Bits) Rotate270(block minecraft.Block) minecraft.Block {
	data := block.Data & b.Bitmask
	block.Data &^= b.Bitmask
	switch data {
	case b.North:
		block.Data |= b.West
	case b.East:
		block.Data |= b.North
	case b.South:
		block.Data |= b.East
	case b.West:
		block.Data |= b.South
	default:
		block.Data |= data
	}
	return block
}

func (b Bits) MirrorX(block minecraft.Block) minecraft.Block {
	data := block.Data & b.Bitmask
	block.Data &^= b.Bitmask
	switch data {
	case b.East:
		block.Data |= b.West
	case b.West:
		block.Data |= b.East
	default:
		block.Data |= data
	}
	return block
}

func (b Bits) MirrorZ(block minecraft.Block) minecraft.Block {
	data := block.Data & b.Bitmask
	block.Data &^= b.Bitmask
	switch data {
	case b.North:
		block.Data |= b.South
	case b.South:
		block.Data |= b.North
	default:
		block.Data |= data
	}
	return block
}

type Sign struct{}

func (Sign) Rotate90(block minecraft.Block) minecraft.Block {
	block.Data = (block.Data + 4) & 15
	return block
}

func (Sign) Rotate180(block minecraft.Block) minecraft.Block {
	block.Data = (block.Data + 8) & 15
	return block
}

func (Sign) Rotate270(block minecraft.Block) minecraft.Block {
	block.Data = (block.Data + 12) & 15
	return block
}

func (Sign) MirrorX(block minecraft.Block) minecraft.Block {
	block.Data = (16 - block.Data) & 15
	return block
}

func (Sign) MirrorZ(block minecraft.Block) minecraft.Block {
	block.Data = (8 - (block.Data & 7)) | (block.Data & 8)
	return block
}

type Door struct{}

func (Door) Rotate90(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		block.Data = (((block.Data & 3) + 1) & 3) | (block.Data & 4)
	}
	return block
}

func (Door) Rotate180(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		block.Data = (((block.Data & 3) + 2) & 3) | (block.Data & 4)
	}
	return block
}

func (Door) Rotate270(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		block.Data = (((block.Data & 3) + 3) & 3) | (block.Data & 4)
	}
	return block
}

func (Door) MirrorX(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		if block.Data == 0 || block.Data == 2 {
			block.Data = ((block.Data + 2) & 3) | (block.Data & 4)
		}
	} else {
		block.Data ^= 1
	}
	return block
}

func (Door) MirrorZ(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		if block.Data == 1 || block.Data == 3 {
			block.Data = ((block.Data + 2) & 3) | (block.Data & 4)
		}
	} else {
		block.Data ^= 1
	}
	return block
}

// Constants for rail data
const (
	RailNorthSouth = 0
	RailEastWest   = 1
	RailEast       = 2
	RailWest       = 3
	RailNorth      = 4
	RailSouth      = 5
	RailNorthWest  = 6
	RailNorthEast  = 7
	RailSouthEast  = 8
	RailSouthWest  = 9
)

type Rail struct{}

func (Rail) Rotate90(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case RailNorthSouth:
		block.Data = RailEastWest
	case RailEastWest:
		block.Data = RailNorthSouth
	case RailEast:
		block.Data = RailSouth
	case RailWest:
		block.Data = RailNorth
	case RailNorth:
		block.Data = RailEast
	case RailSouth:
		block.Data = RailWest
	case RailNorthWest:
		block.Data = RailNorthEast
	case RailNorthEast:
		block.Data = RailSouthEast
	case RailSouthEast:
		block.Data = RailSouthWest
	case RailSouthWest:
		block.Data = RailNorthWest
	}
	return block
}

func (Rail) Rotate180(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case RailEast:
		block.Data = RailWest
	case RailWest:
		block.Data = RailEast
	case RailNorth:
		block.Data = RailSouth
	case RailSouth:
		block.Data = RailNorth
	case RailNorthWest:
		block.Data = RailSouthEast
	case RailNorthEast:
		block.Data = RailSouthWest
	case RailSouthEast:
		block.Data = RailNorthWest
	case RailSouthWest:
		block.Data = RailNorthEast
	}
	return block
}

func (Rail) Rotate270(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case RailNorthSouth:
		block.Data = RailEastWest
	case RailEastWest:
		block.Data = RailNorthSouth
	case RailEast:
		block.Data = RailNorth
	case RailWest:
		block.Data = RailSouth
	case RailNorth:
		block.Data = RailWest
	case RailSouth:
		block.Data = RailEast
	case RailNorthWest:
		block.Data = RailSouthWest
	case RailNorthEast:
		block.Data = RailNorthWest
	case RailSouthEast:
		block.Data = RailNorthEast
	case RailSouthWest:
		block.Data = RailSouthEast
	}
	return block
}

func (Rail) MirrorX(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case RailEast:
		block.Data = RailWest
	case RailWest:
		block.Data = RailEast
	case RailNorthWest:
		block.Data = RailNorthEast
	case RailNorthEast:
		block.Data = RailNorthWest
	case RailSouthEast:
		block.Data = RailSouthWest
	case RailSouthWest:
		block.Data = RailSouthEast
	}
	return block
}

func (Rail) MirrorZ(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case RailNorth:
		block.Data = RailSouth
	case RailSouth:
		block.Data = RailNorth
	case RailNorthWest:
		block.Data = RailSouthWest
	case RailNorthEast:
		block.Data = RailSouthEast
	case RailSouthEast:
		block.Data = RailNorthEast
	case RailSouthWest:
		block.Data = RailNorthWest
	}
	return block
}

type PowerableRail struct {
	Rail
}

func (p PowerableRail) Rotate90(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rail.Rotate90(block)
	block.Data |= powered
	return block
}

func (p PowerableRail) Rotate180(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rail.Rotate180(block)
	block.Data |= powered
	return block
}

func (p PowerableRail) Rotate270(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rail.Rotate270(block)
	block.Data |= powered
	return block
}

func (p PowerableRail) MirrorX(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rail.MirrorX(block)
	block.Data |= powered
	return block
}

func (p PowerableRail) MirrorZ(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rail.MirrorZ(block)
	block.Data |= powered
	return block
}

// Constants for rail data
const (
	LeverNorth             = 4
	LeverEast              = 1
	LeverSouth             = 3
	LeverWest              = 2
	LeverGroundNorthSouth  = 5
	LeverGroundEastWest    = 6
	LeverCeilingNorthSouth = 7
	LeverCeilingEastWest   = 0
)

type Lever struct{}

func (Lever) Rotate90(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	switch block.Data & 7 {
	case LeverNorth:
		block.Data = LeverEast
	case LeverEast:
		block.Data = LeverSouth
	case LeverSouth:
		block.Data = LeverWest
	case LeverWest:
		block.Data = LeverNorth
	case LeverGroundNorthSouth:
		block.Data = LeverGroundEastWest
	case LeverGroundEastWest:
		block.Data = LeverGroundNorthSouth
	case LeverCeilingNorthSouth:
		block.Data = LeverCeilingEastWest
	case LeverCeilingEastWest:
		block.Data = LeverCeilingNorthSouth
	}
	block.Data |= powered
	return block
}

func (Lever) Rotate180(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	switch block.Data & 7 {
	case LeverNorth:
		block.Data = LeverSouth
	case LeverEast:
		block.Data = LeverWest
	case LeverSouth:
		block.Data = LeverNorth
	case LeverWest:
		block.Data = LeverEast
	}
	block.Data |= powered
	return block
}

func (Lever) Rotate270(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	switch block.Data & 7 {
	case LeverNorth:
		block.Data = LeverWest
	case LeverEast:
		block.Data = LeverNorth
	case LeverSouth:
		block.Data = LeverEast
	case LeverWest:
		block.Data = LeverSouth
	case LeverGroundNorthSouth:
		block.Data = LeverGroundEastWest
	case LeverGroundEastWest:
		block.Data = LeverGroundNorthSouth
	case LeverCeilingNorthSouth:
		block.Data = LeverCeilingEastWest
	case LeverCeilingEastWest:
		block.Data = LeverCeilingNorthSouth
	}
	block.Data |= powered
	return block
}

func (Lever) MirrorX(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	switch block.Data & 7 {
	case LeverEast:
		block.Data = LeverWest
	case LeverWest:
		block.Data = LeverEast
	}
	block.Data |= powered
	return block
}

func (Lever) MirrorZ(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	switch block.Data & 7 {
	case LeverNorth:
		block.Data = LeverSouth
	case LeverSouth:
		block.Data = LeverNorth
	}
	block.Data |= powered
	return block
}

// Constants for rail data
const (
	MushroomNorth     = 2
	MushroomNorthEast = 3
	MushroomEast      = 6
	MushroomSouthEast = 9
	MushroomSouth     = 8
	MushroomSouthWest = 7
	MushroomWest      = 4
	MushroomNorthWest = 1
)

type Mushroom struct{}

func (Mushroom) Rotate90(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case MushroomNorth:
		block.Data = MushroomEast
	case MushroomNorthEast:
		block.Data = MushroomSouthEast
	case MushroomEast:
		block.Data = MushroomSouth
	case MushroomSouthEast:
		block.Data = MushroomSouthWest
	case MushroomSouth:
		block.Data = MushroomWest
	case MushroomSouthWest:
		block.Data = MushroomNorthWest
	case MushroomWest:
		block.Data = MushroomNorth
	case MushroomNorthWest:
		block.Data = MushroomNorthEast
	}
	return block
}

func (Mushroom) Rotate180(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case MushroomNorth:
		block.Data = MushroomSouth
	case MushroomNorthEast:
		block.Data = MushroomSouthWest
	case MushroomEast:
		block.Data = MushroomWest
	case MushroomSouthEast:
		block.Data = MushroomNorthWest
	case MushroomSouth:
		block.Data = MushroomNorth
	case MushroomSouthWest:
		block.Data = MushroomNorthEast
	case MushroomWest:
		block.Data = MushroomEast
	case MushroomNorthWest:
		block.Data = MushroomSouthEast
	}
	return block
}

func (Mushroom) Rotate270(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case MushroomNorth:
		block.Data = MushroomWest
	case MushroomNorthEast:
		block.Data = MushroomNorthWest
	case MushroomEast:
		block.Data = MushroomNorth
	case MushroomSouthEast:
		block.Data = MushroomNorthEast
	case MushroomSouth:
		block.Data = MushroomEast
	case MushroomSouthWest:
		block.Data = MushroomSouthEast
	case MushroomWest:
		block.Data = MushroomSouth
	case MushroomNorthWest:
		block.Data = MushroomSouthWest
	}
	return block
}

func (Mushroom) MirrorX(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case MushroomNorthEast:
		block.Data = MushroomNorthWest
	case MushroomEast:
		block.Data = MushroomWest
	case MushroomSouthEast:
		block.Data = MushroomSouthWest
	case MushroomSouthWest:
		block.Data = MushroomSouthEast
	case MushroomWest:
		block.Data = MushroomEast
	case MushroomNorthWest:
		block.Data = MushroomNorthEast
	}
	return block
}

func (Mushroom) MirrorZ(block minecraft.Block) minecraft.Block {
	switch block.Data {
	case MushroomNorth:
		block.Data = MushroomSouth
	case MushroomNorthEast:
		block.Data = MushroomSouthEast
	case MushroomSouthEast:
		block.Data = MushroomNorthEast
	case MushroomSouth:
		block.Data = MushroomNorth
	case MushroomSouthWest:
		block.Data = MushroomNorthWest
	case MushroomNorthWest:
		block.Data = MushroomSouthWest
	}
	return block
}

// Constants for rail data
const (
	VinesNorth = 4
	VinesEast  = 8
	VinesSouth = 1
	VinesWest  = 2
)

type Vines struct{}

func (Vines) Rotate90(block minecraft.Block) minecraft.Block {
	data := block.Data
	block.Data = 0
	if data&VinesNorth != 0 {
		block.Data = VinesEast
	}
	if data&VinesEast != 0 {
		block.Data = VinesSouth
	}
	if data&VinesSouth != 0 {
		block.Data = VinesWest
	}
	if data&VinesWest != 0 {
		block.Data = VinesNorth
	}
	return block
}

func (Vines) Rotate180(block minecraft.Block) minecraft.Block {
	data := block.Data
	block.Data = 0
	if data&VinesNorth != 0 {
		block.Data = VinesSouth
	}
	if data&VinesEast != 0 {
		block.Data = VinesEast
	}
	if data&VinesSouth != 0 {
		block.Data = VinesNorth
	}
	if data&VinesWest != 0 {
		block.Data = VinesEast
	}
	return block
}

func (Vines) Rotate270(block minecraft.Block) minecraft.Block {
	data := block.Data
	block.Data = 0
	if data&VinesNorth != 0 {
		block.Data = VinesWest
	}
	if data&VinesEast != 0 {
		block.Data = VinesNorth
	}
	if data&VinesSouth != 0 {
		block.Data = VinesEast
	}
	if data&VinesWest != 0 {
		block.Data = VinesSouth
	}
	return block
}

func (Vines) MirrorX(block minecraft.Block) minecraft.Block {
	data := block.Data
	block.Data = 0
	if data&VinesNorth != 0 {
		block.Data = VinesNorth
	}
	if data&VinesEast != 0 {
		block.Data = VinesWest
	}
	if data&VinesSouth != 0 {
		block.Data = VinesSouth
	}
	if data&VinesWest != 0 {
		block.Data = VinesEast
	}
	return block
}

func (Vines) MirrorZ(block minecraft.Block) minecraft.Block {
	data := block.Data
	block.Data = 0
	if data&VinesNorth != 0 {
		block.Data = VinesSouth
	}
	if data&VinesEast != 0 {
		block.Data = VinesEast
	}
	if data&VinesSouth != 0 {
		block.Data = VinesNorth
	}
	if data&VinesWest != 0 {
		block.Data = VinesWest
	}
	return block
}

const SkullRotationString = "Rot"

type Skull struct {
	b Bits
}

func NewSkull() Skull {
	return Skull{NewBits(3, 2, 4, 3, 5)}
}

func (s Skull) Rotate90(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, (r+4)&15))
	block.SetMetadata(m)
	return block
}

func (s Skull) Rotate180(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, (r+8)&15))
	block.SetMetadata(m)
	return block
}

func (s Skull) Rotate270(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, (r+12)&15))
	block.SetMetadata(m)
	return block
}

func (s Skull) MirrorX(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, ((16 - (r & 3)) & 15)))
	block.SetMetadata(m)
	return block
}

func (s Skull) MirrorZ(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, (8-(r&3))|(r&8)))
	block.SetMetadata(m)
	return block
}

func init() {

}
