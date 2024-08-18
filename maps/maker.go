package maps

import (
	"image"
	"image/color"

	"vimagination.zapto.org/minecraft"
)

// Image represents a Minecraft Map.
type Image struct {
	level         *minecraft.Level
	bounds        image.Rectangle
	scale         uint8
	y             int32
	width, height int
	colour        func(minecraft.Block) color.Color
}

// ColorModel returns the palette for the map.
func (Image) ColorModel() color.Model {
	return palette
}

// Bounds returns the dimensions of the map.
func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{},
		Max: image.Point{i.width, i.height},
	}
}

// At returns the colour at the specified coords.
func (i Image) At(x, z int) color.Color {
	var y int32

	if i.y < 0 {
		var err error

		y, err = i.level.GetHeight(int32(x), int32(z))
		if err != nil {
			return color.Transparent
		}
	} else {
		y = i.y
	}

	b, err := i.level.GetBlock(int32(x), y, int32(z))
	if err != nil {
		return color.Transparent
	}

	return i.colour(b)
}

// BlockColor is the standard block-to-colour func.
func BlockColor(b minecraft.Block) color.Color {
	if c, ok := colours[uint32(b.ID)<<16|uint32(b.Data)]; ok {
		return palette[c]
	}

	if c, ok := colours[uint32(b.ID)]; ok {
		return palette[c]
	}

	return color.Transparent
}

// Option represents a optional parameter for a map type.
type Option func(*Image)

// FixedY is an options to fix the Y-coord of the blocks to be read. By default
// the highest, non-transparent block is used.
func FixedY(y int32) Option {
	return func(i *Image) {
		i.y = y
	}
}

// Scale sets the scale the map is to be rendered at.
func Scale(s uint8) Option {
	return func(i *Image) {
		i.scale = s
	}
}

// ColorFunc is an option for NewMap that specifies what colour blocks are
// painted as.
func ColorFunc(c func(minecraft.Block) color.Color) Option {
	return func(i *Image) {
		i.colour = c
	}
}

// NewMap creates an image from a Minecraft level.
func NewMap(l *minecraft.Level, bounds image.Rectangle, options ...Option) Image {
	i := Image{
		level:  l,
		bounds: bounds,
		scale:  3,
		y:      -1,
		colour: BlockColor,
	}

	for _, o := range options {
		o(&i)
	}

	i.width = bounds.Dx() >> i.scale
	i.height = bounds.Dy() >> i.scale

	return i
}

var colours = map[uint32]uint8{}
