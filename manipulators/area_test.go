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
	defer l.Close()
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

func TestAreaFill(t *testing.T) {
	t.Parallel()
	l, _ := minecraft.NewLevel(minecraft.NewMemPath())
	defer l.Close()
	b := minecraft.Block{ID: 1}
	a := NewArea(1, 1, 1, 4, 4, 4, l)
	a.Fill(b)
	for x := int32(0); x < 6; x++ {
		for y := int32(0); y < 6; y++ {
			for z := int32(0); z < 6; z++ {
				bl := b
				if x == 0 || x == 5 || y == 0 || y == 5 || z == 0 || z == 5 {
					bl = minecraft.Block{}
				}
				got, _ := l.GetBlock(x, y, z)
				if !got.EqualBlock(bl) {
					t.Errorf("at coords, %d, %d, %d, expecting %v, got %v", x, y, z, bl, got)
				}
			}
		}
	}
}

func TestAreaReplace(t *testing.T) {
	t.Parallel()
	l, _ := minecraft.NewLevel(minecraft.NewMemPath())
	defer l.Close()
	b := minecraft.Block{ID: 1}
	c := minecraft.Block{ID: 2}
	a := NewArea(0, 0, 0, 5, 5, 5, l)
	for x := int32(0); x < 6; x++ {
		for y := int32(0); y < 6; y++ {
			for z := int32(0); z < 6; z++ {
				if x == y || x == z || y == z {
					a.Set(x, y, z, b)
				}
			}
		}
	}
	a.Replace(b, c)
	for x := int32(0); x < 6; x++ {
		for y := int32(0); y < 6; y++ {
			for z := int32(0); z < 6; z++ {
				var bl minecraft.Block
				if x == y || x == z || y == z {
					bl = c
				}
				got, _ := a.Get(x, y, z)
				if !got.EqualBlock(bl) {
					t.Errorf("at coords, %d, %d, %d, expecting %v, got %v", x, y, z, bl, got)
				}
			}
		}
	}
}

func (a Area) EqualTo(b Area) bool {
	if a.Width() != b.Width() || a.Height() != b.Height() || a.Depth() != b.Depth() {
		return false
	}
	for x := int32(0); x < a.Width(); x++ {
		for y := int32(0); y < a.Height(); y++ {
			for z := int32(0); z < a.Depth(); z++ {
				b1, _ := a.Get(x, y, z)
				b2, _ := b.Get(x, y, z)
				if !b1.EqualBlock(b2) {
					return false
				}
			}
		}
	}
	return true
}

func TestAreaCopyTo(t *testing.T) {
	t.Parallel()
	l, _ := minecraft.NewLevel(minecraft.NewMemPath())
	defer l.Close()
	b := minecraft.Block{ID: 1}
	a := NewArea(0, 0, 0, 5, 5, 5, l)
	for x := int32(0); x < a.Width(); x++ {
		for y := int32(0); y < a.Height(); y++ {
			for z := int32(0); z < a.Depth(); z++ {
				if (x+y+z)%7 == 0 {
					a.Set(x, y, z, b)
				}
			}
		}
	}
	a2 := NewArea(10, 0, 0, 15, 5, 5, l)
	a.CopyTo(a2)
	if !a.EqualTo(a2) {
		t.Errorf("area2 not equal to area 1")
		return
	}
	a3 := NewArea(11, 1, 1, 16, 6, 6, l)
	a2.CopyTo(a3)
	if !a.EqualTo(a3) {
		t.Errorf("area3 not equal to area 1")
		return
	} else if a.EqualTo(a2) {
		t.Errorf("area2 equal to area 1, shouldn't be")
		return
	}
	a3.CopyTo(a2)
	if !a.EqualTo(a2) {
		t.Errorf("area2 not equal to area 1")
		return
	} else if a.EqualTo(a3) {
		t.Errorf("area3 equal to area 1, shouldn't be")
		return
	}
}
