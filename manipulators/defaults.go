package manipulators

import "github.com/MJKWoolnough/minecraft"

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
