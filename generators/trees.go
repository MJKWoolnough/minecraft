package generator

import (
	"errors"
	"math/rand"
	"time"

	"github.com/MJKWoolnough/minecraft"
)

type tree struct {
	wood, leaves              minecraft.Block
	minHeight, maxHeight      int32
	trunkWidth, trunkDepth    int32
	leavesRadius, leavesStart int32
	branches, vines           bool
	vineChance                int32
	branchWidth               int32
	random                    Rand
}

type treeOption func(*tree)

func (g Generator) Tree(x, y, z int32, os ...treeOption) error {
	t := tree{
		wood:         minecraft.Block{BlockID: 17},
		leaves:       minecraft.Block{BlockID: 18},
		leavesStart:  4,
		trunkWidth:   1,
		trunkDepth:   1,
		leavesRadius: 3,
		minHeight:    3,
		maxHeight:    5,
	}
	for _, o := range os {
		o(&t)
	}
	if t.random == nil {
		t.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	height := t.minHeight + t.random.Int31n(t.maxHeight-t.minHeight)

	//check space
Loop:
	for j := y; j < y+height; j++ {
		for i := x; i < x+t.trunkWidth; i++ {
			for k := z; k < z+t.trunkWidth; k++ {
				b, err := g.GetBlock(i, j, k)
				if err != nil {
					return err
				}
				if b.BlockID != 0 { //check for all replaceable blocks?
					height = j - y - 1
					break Loop
				}
			}
		}
	}
	if height < t.minHeight {
		return ErrTreeNoRoom
	}

	// create main trunk
	for i := x; i < x+t.trunkWidth; i++ {
		for j := y; j < y+height; j++ {
			for k := z; k < z+t.trunkWidth; k++ {
				g.SetBlock(i, j, k, &t.wood)
			}
		}
	}

	// create branches
	if t.branches {

	}

	// create leaves

	// create vines
	if t.vines {

	}

	return nil
}

var ErrTreeNoRoom = errors.New("no room for tree")
