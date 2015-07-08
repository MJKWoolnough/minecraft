package manipulators

import (
	"errors"

	"github.com/MJKWoolnough/minecraft"
)

type Area struct {
	x1, y1, z1, x2, y2, z2 int32
	level                  *minecraft.Level
}

func NewArea(x1, y1, z1, x2, y2, z2 int32, level *minecraft.Level) Area {
	return Area{min(x1, x2), min(y1, y2), min(z1, z2), max(x1, x2), max(y1, y2), max(z1, z2), level}
}

func (a Area) Width() int32 {
	return a.x2 - a.x1 + 1 //no zero width, everything contains at least one block
}

func (a Area) Height() int32 {
	return a.y2 - a.y1 + 1
}

func (a Area) Depth() int32 {
	return a.z2 - a.z1 + 1
}

func (a Area) Get(x, y, z int32) (*minecraft.Block, error) {
	if x < 0 || y < 0 || z < 0 {
		return nil, ErrOOB
	}
	x += a.x1
	y += a.y1
	z += a.z1
	if x > a.x2 || y > a.y2 || z > a.z2 {
		return nil, ErrOOB
	}
	return a.level.GetBlock(x, y, z)
}

func (a Area) Set(x, y, z int32, b *minecraft.Block) error {
	if x < 0 || y < 0 || z < 0 {
		return ErrOOB
	}
	x += a.x1
	y += a.y1
	z += a.z1
	if x > a.x2 || y > a.y2 || z > a.z2 {
		return ErrOOB
	}
	return a.level.SetBlock(x, y, z, b)
}

func (a Area) Fill(b *minecraft.Block) error {
	for x := a.x1; x <= a.x2; x++ {
		for y := a.y1; y <= a.y2; y++ {
			for z := a.z1; z <= a.z2; z++ {
				if err := a.level.SetBlock(x, y, z, b); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a Area) Replace(replace, with *minecraft.Block) error {
	for x := a.x1; x <= a.x2; x++ {
		for y := a.y1; y <= a.y2; y++ {
			for z := a.z1; z <= a.z2; z++ {
				b, err := a.level.GetBlock(x, y, z)
				if err != nil {
					return err
				}
				if b.EqualBlock(replace) {
					err = a.level.SetBlock(x, y, z, with)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (a Area) CopyTo(b Area) error {
	if a.Width() != b.Width() || a.Height() != b.Height() || a.Depth() != b.Depth() {
		return ErrMismatch
	}

	var xStart, yStart, zStart, xEnd, yEnd, zEnd, xStep, yStep, zStep int32

	if a.x1 < b.x1 {
		xStart = a.Width() - 1
		xEnd = -1
		xStep = -1
	} else {
		xStart = 0
		xEnd = a.Width()
		xStep = 1
	}

	if a.y1 < b.y1 {
		yStart = a.Height() - 1
		yEnd = -1
		yStep = -1
	} else {
		yStart = 0
		yEnd = a.Height()
		yStep = 1
	}

	if a.z1 < b.z1 {
		zStart = a.Depth() - 1
		zEnd = -1
		zStep = -1
	} else {
		zStart = 0
		zEnd = a.Depth()
		zStep = 1
	}

	for x := xStart; x != xEnd; x += xStep {
		for y := yStart; y != yEnd; y += yStep {
			for z := zStart; z != zEnd; z += zStep {
				block, err := a.Get(x, y, z)
				if err != nil {
					return err
				}
				err = b.Set(x, y, z, block)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

var (
	Up    = Direction{0, 1, 0}
	Down  = Direction{0, -1, 0}
	North = Direction{0, 0, -1}
	South = Direction{0, 0, 1}
	East  = Direction{1, 0, 0}
	West  = Direction{-1, 0, 0}
)

type Direction struct {
	x, y, z int32
}

func (a Area) CopyInDirection(dir Direction, times uint) error {
	w := a.Width()
	h := a.Height()
	d := a.Depth()
	for x := a.x1; x <= a.x2; x++ {
		for y := a.y1; y <= a.y2; y++ {
			for z := a.z1; z <= a.z2; z++ {
				block, err := a.level.GetBlock(x, y, z)
				if err != nil {
					return err
				}
				i, j, k := x, y, z
				for t := times; t > 0; t-- {
					i += w * dir.x
					j += h * dir.y
					k += d * dir.z
					err = a.level.SetBlock(x, y, z, block)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil

}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

//Errors
var (
	ErrOOB      = errors.New("out of bounds")
	ErrMismatch = errors.New("areas not equal sizes")
)
