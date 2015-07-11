package manipulators

import (
	"sync"

	"github.com/MJKWoolnough/minecraft"
)

type manipulator interface {
	Rotate90(minecraft.Block) minecraft.Block
	Rotate180(minecraft.Block) minecraft.Block
	Rotate270(minecraft.Block) minecraft.Block
	MirrorX(minecraft.Block) minecraft.Block
	MirrorZ(minecraft.Block) minecraft.Block
}

var (
	mmu          sync.RWMutex
	manipulators map[uint16]manipulator
)

func Register(blockID uint16, m manipulator) {
	mmu.Lock()
	defer mmu.Unlock()
	manipulators[blockID] = m
}

func RegisterBlock(block minecraft.Block, m manipulator) {
	Register(block.ID, m)
}

func Rotate90(b minecraft.Block) minecraft.Block {
	mmu.RLock()
	defer mmu.RUnlock()
	m, ok := manipulators[b.ID]
	if !ok {
		return b
	}
	return m.Rotate90(b)
}

func Rotate180(b minecraft.Block) minecraft.Block {
	mmu.RLock()
	defer mmu.RUnlock()
	m, ok := manipulators[b.ID]
	if !ok {
		return b
	}
	return m.Rotate180(b)
}

func Rotate270(b minecraft.Block) minecraft.Block {
	mmu.RLock()
	defer mmu.RUnlock()
	m, ok := manipulators[b.ID]
	if !ok {
		return b
	}
	return m.Rotate270(b)
}

func MirrorX(b minecraft.Block) minecraft.Block {
	mmu.RLock()
	defer mmu.RUnlock()
	m, ok := manipulators[b.ID]
	if !ok {
		return b
	}
	return m.MirrorX(b)
}

func MirrorZ(b minecraft.Block) minecraft.Block {
	mmu.RLock()
	defer mmu.RUnlock()
	m, ok := manipulators[b.ID]
	if !ok {
		return b
	}
	return m.MirrorZ(b)
}
