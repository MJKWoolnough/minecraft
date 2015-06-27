package generator

import (
	"math/rand"
	"time"

	"github.com/MJKWoolnough/minecraft"
)

type tree struct {
	wood, leaves           minecraft.Block
	minHeight, maxHeight   int32
	trunkWidth, trunkDepth int32
	leavesRadius           int32
	branches, vines        bool
	vineChance             int32
	branchWidth            int32
	random                 Rand
}

type treeOption func(*tree)

func (g Generator) Tree(x, y, z int32, os ...treeOption) error {
	t := tree{
		wood:         minecraft.Block{BlockID: 17},
		leaves:       minecraft.Block{BlockID: 18},
		trunkWidth:   1,
		trunkDepth:   1,
		leavesRadius: 3,
		random:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	for _, o := range os {
		o(&t)
	}
	return nil
}
