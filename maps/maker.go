package maps

import (
	"image"
	"image/color"

	"github.com/MJKWoolnough/minecraft"
)

type Image struct {
	level         *minecraft.Level
	bounds        image.Rectangle
	scale         uint8
	y             int32
	width, height int
	colour        func(minecraft.Block) color.Color
}

func (Image) ColourModel() color.Model {
	return palette
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{},
		Max: image.Point{i.width, i.height},
	}
}

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

func BlockColor(b minecraft.Block) color.Color {
	id := uint32(b.ID)<<16 | uint32(b.Data)
	if c, ok := colours[id]; ok {
		return palette[c]
	}
	if c, ok := colours[id>>16]; ok {
		return palette[c]
	}
	return color.Transparent
}

type option func(*Image)

func FixedY(y int32) option {
	return func(i *Image) {
		i.y = y
	}
}

func Scale(s uint8) option {
	return func(i *Image) {
		i.scale = s
	}
}

func ColorFunc(c func(minecraft.Block) color.Color) option {
	return func(i *Image) {
		i.colour = c
	}
}

func NewMap(l *minecraft.Level, bounds image.Rectangle, options ...option) Image {
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
