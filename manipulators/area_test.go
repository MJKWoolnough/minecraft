package manipulators

import (
	"testing"

	"github.com/MJKWoolnough/minecraft"
)

func TestAreaDimensions(t *testing.T) {
	t.Parallel()
	l, _ := minecraft.NewLevel(minecraft.NewMemPath())
	defer l.Close()
	tests := []struct {
		x1, y1, z1, x2, y2, z2, w, h, d int32
	}{
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1, 1, 1, 1},
		{-1, -1, -1, 0, 0, 0, 2, 2, 2},
		{64, 3, -5, 70, 1, -10, 7, 3, 6},
	}

	for n, test := range tests {
		a := NewArea(test.x1, test.y1, test.z1, test.x2, test.y2, test.z2, l)
		if a.Width() != test.w {
			t.Errorf("test %d: expecting width of %d, got %d", n+1, test.w, a.Width())
		}
		if a.Height() != test.h {
			t.Errorf("test %d: expecting height of %d, got %d", n+1, test.h, a.Height())
		}
		if a.Depth() != test.d {
			t.Errorf("test %d: expecting depth of %d, got %d", n+1, test.d, a.Depth())
		}
	}
}

func TestAreaGetSet(t *testing.T) {
	t.Parallel()
	l, _ := minecraft.NewLevel(minecraft.NewMemPath())
	a1 := NewArea(3, 4, 5, 6, 7, 8, l)
	a2 := NewArea(-3, 12, -5, -6, 17, -8, l)
	tests := []struct {
		x1, y1, z1, x2, y2, z2, x3, y3, z3 int32
		b                                  minecraft.Block
	}{
		{0, 0, 0, 3, 4, 5, -6, 12, -8, minecraft.Block{ID: 1}},
		{1, 1, 1, 4, 5, 6, -5, 13, -7, minecraft.Block{ID: 2}},
	}

	for n, test := range tests {
		a1.Set(test.x1, test.y1, test.z1, test.b)
		a2.Set(test.x1, test.y1, test.z1, test.b)
		b, _ := a1.Get(test.x1, test.y1, test.z1)
		if !b.EqualBlock(test.b) {
			t.Errorf("test %d-1: incorrect block gotten", n+1)
		}
		b, _ = a2.Get(test.x1, test.y1, test.z1)
		if !b.EqualBlock(test.b) {
			t.Errorf("test %d-2: incorrect block gotten", n+1)
		}
		b, _ = l.GetBlock(test.x2, test.y2, test.z2)
		if !b.EqualBlock(test.b) {
			t.Errorf("test %d-3: incorrect block gotten", n+1)
		}
		b, _ = l.GetBlock(test.x3, test.y3, test.z3)
		if !b.EqualBlock(test.b) {
			t.Errorf("test %d-4: incorrect block gotten", n+1)
		}
	}
}
