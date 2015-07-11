package vanilla16

import (
	"github.com/MJKWoolnough/minecraft"
	"github.com/MJKWoolnough/minecraft/manipulators"
	"github.com/MJKWoolnough/minecraft/nbt"
)

type Log struct {
	manipulators.Default
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

type Doors struct{}

func (Doors) Rotate90(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		block.Data = (((block.Data & 3) + 1) & 3) | (block.Data & 4)
	}
	return block
}

func (Doors) Rotate180(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		block.Data = (((block.Data & 3) + 2) & 3) | (block.Data & 4)
	}
	return block
}

func (Doors) Rotate270(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		block.Data = (((block.Data & 3) + 3) & 3) | (block.Data & 4)
	}
	return block
}

func (Doors) MirrorX(block minecraft.Block) minecraft.Block {
	if block.Data&8 == 0 {
		if block.Data == 0 || block.Data == 2 {
			block.Data = ((block.Data + 2) & 3) | (block.Data & 4)
		}
	} else {
		block.Data ^= 1
	}
	return block
}

func (Doors) MirrorZ(block minecraft.Block) minecraft.Block {
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

type Rails struct{}

func (Rails) Rotate90(block minecraft.Block) minecraft.Block {
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

func (Rails) Rotate180(block minecraft.Block) minecraft.Block {
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

func (Rails) Rotate270(block minecraft.Block) minecraft.Block {
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

func (Rails) MirrorX(block minecraft.Block) minecraft.Block {
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

func (Rails) MirrorZ(block minecraft.Block) minecraft.Block {
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
	Rails
}

func (p PowerableRail) Rotate90(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rails.Rotate90(block)
	block.Data |= powered
	return block
}

func (p PowerableRail) Rotate180(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rails.Rotate180(block)
	block.Data |= powered
	return block
}

func (p PowerableRail) Rotate270(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rails.Rotate270(block)
	block.Data |= powered
	return block
}

func (p PowerableRail) MirrorX(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rails.MirrorX(block)
	block.Data |= powered
	return block
}

func (p PowerableRail) MirrorZ(block minecraft.Block) minecraft.Block {
	powered := block.Data & 8
	block.Data &= 7
	p.Rails.MirrorZ(block)
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

type Levers struct{}

func (Levers) Rotate90(block minecraft.Block) minecraft.Block {
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

func (Levers) Rotate180(block minecraft.Block) minecraft.Block {
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

func (Levers) Rotate270(block minecraft.Block) minecraft.Block {
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

func (Levers) MirrorX(block minecraft.Block) minecraft.Block {
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

func (Levers) MirrorZ(block minecraft.Block) minecraft.Block {
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

type Skulls struct {
	b manipulators.Bits
}

func NewSkulls() Skulls {
	return Skulls{manipulators.Bits{3, 2, 4, 3, 5}}
}

func (s Skulls) Rotate90(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, (r+4)&15))
	block.SetMetadata(m)
	return block
}

func (s Skulls) Rotate180(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, (r+8)&15))
	block.SetMetadata(m)
	return block
}

func (s Skulls) Rotate270(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, (r+12)&15))
	block.SetMetadata(m)
	return block
}

func (s Skulls) MirrorX(block minecraft.Block) minecraft.Block {
	if block.Data != 1 {
		return s.b.Rotate90(block)
	}
	m := block.GetMetadata()
	r, _ := m.Get(SkullRotationString).Data().(nbt.Byte)
	m.Set(nbt.NewTag(SkullRotationString, ((16 - (r & 3)) & 15)))
	block.SetMetadata(m)
	return block
}

func (s Skulls) MirrorZ(block minecraft.Block) minecraft.Block {
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
	manipulators.RegisterBlock(Wood, Log{})

	uewsn := manipulators.NewBits(7, 4, 1, 3, 2)
	manipulators.RegisterBlock(Torch, &uewsn)
	manipulators.RegisterBlock(RedstoneTorch, &uewsn)
	manipulators.RegisterBlock(RedstoneTorchActive, &uewsn)
	manipulators.RegisterBlock(StoneButton, &uewsn)
	manipulators.RegisterBlock(WoodenButton, &uewsn)

	nesw := manipulators.NewBits(3, 0, 1, 2, 3)
	manipulators.RegisterBlock(Bed, &nesw)
	manipulators.RegisterBlock(RedstoneRepeater, &nesw)
	manipulators.RegisterBlock(RedstoneRepeaterActive, &nesw)
	manipulators.RegisterBlock(RedstoneComparator, &nesw)
	manipulators.RegisterBlock(RedstoneComparatorActive, &nesw)
	manipulators.RegisterBlock(Cocoa, &nesw)
	manipulators.RegisterBlock(Anvil, &nesw)

	uunswe := manipulators.NewBits(7, 2, 5, 3, 4)
	manipulators.RegisterBlock(Piston, &uunswe)
	manipulators.RegisterBlock(PistonExtension, &uunswe)
	manipulators.RegisterBlock(PistonMoving, &uunswe)
	manipulators.RegisterBlock(StickyPiston, &uunswe)
	manipulators.RegisterBlock(Ladder, &uunswe)
	manipulators.RegisterBlock(SignWall, &uunswe)
	manipulators.RegisterBlock(Furnace, &uunswe)
	manipulators.RegisterBlock(FurnaceActive, &uunswe)
	manipulators.RegisterBlock(Chest, &uunswe)
	manipulators.RegisterBlock(Dispenser, &uunswe)
	manipulators.RegisterBlock(Dropper, &uunswe)
	manipulators.RegisterBlock(Hopper, &uunswe)
	manipulators.RegisterBlock(EnderChest, &uunswe)
	manipulators.RegisterBlock(TrappedChest, &uunswe)
	manipulators.RegisterBlock(LockedChest, &uunswe)

	ewsn := manipulators.NewBits(3, 3, 0, 2, 1)
	manipulators.RegisterBlock(OakStairs, &ewsn)
	manipulators.RegisterBlock(CobblestoneStairs, &ewsn)
	manipulators.RegisterBlock(BrickStairs, &ewsn)
	manipulators.RegisterBlock(StoneBrickStairs, &ewsn)
	manipulators.RegisterBlock(NetherBrickStairs, &ewsn)
	manipulators.RegisterBlock(SandStoneStairs, &ewsn)
	manipulators.RegisterBlock(SpruceStairs, &ewsn)
	manipulators.RegisterBlock(BirchStairs, &ewsn)
	manipulators.RegisterBlock(JungleStairs, &ewsn)
	manipulators.RegisterBlock(NetherQuartzStairs, &ewsn)

	manipulators.RegisterBlock(SignPost, Sign{})

	var doors Doors
	manipulators.RegisterBlock(Door, &doors)
	manipulators.RegisterBlock(IronDoor, &doors)

	manipulators.RegisterBlock(Rail, Rails{})

	var rail PowerableRail
	manipulators.RegisterBlock(Rail, &rail)
	manipulators.RegisterBlock(DetectorRail, &rail)
	manipulators.RegisterBlock(RailActivator, &rail)

	manipulators.RegisterBlock(Lever, Levers{})

	swne := manipulators.NewBits(3, 2, 3, 0, 1)
	manipulators.RegisterBlock(Pumpkin, &swne)
	manipulators.RegisterBlock(JackOLantern, &swne)
	manipulators.RegisterBlock(FenceGate, &swne)
	manipulators.RegisterBlock(EndPortalFrame, &swne)
	manipulators.RegisterBlock(TripwireHook, &swne)

	manipulators.RegisterBlock(Trapdoor, manipulators.NewBits(3, 1, 2, 0, 3))

	var mushroom Mushroom
	manipulators.RegisterBlock(BrownMushroom, &mushroom)
	manipulators.RegisterBlock(RedMushroom, &mushroom)

	manipulators.RegisterBlock(Vine, Vines{})

	manipulators.RegisterBlock(Skull, Skulls{})
}
