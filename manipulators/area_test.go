package manipulators

import (
	"testing"

	"github.com/MJKWoolnough/minecraft"
)

func TestAreaDimensions(t *testing.T) {
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
