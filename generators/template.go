package generators

import (
	"github.com/MJKWoolnough/minecraft"
	"github.com/MJKWoolnough/minecraft/manipulators"
	"github.com/MJKWoolnough/minecraft/nbt"
)

type Section struct {
}

type Template struct {
	Version                         uint8
	SectionsX, SectionsY, SectionsZ uint8
	Sections                        []Section
	Blocks                          []uint16
	MetaData                        map[uint16]nbt.Compound
}

type options struct {
	Rotation, Mirror, BlockSwapper func(minecraft.Block) minecraft.Block
}

func (o *options) Rotate90() {
	o.Rotation = manipulators.Rotate90
}

func (o *options) Rotate180() {
	o.Rotation = manipulators.Rotate180
}

func (o *options) Rotate270() {
	o.Rotation = manipulators.Rotate270
}

func (o *options) MirrorX() {
	o.Mirror = manipulators.MirrorX
}

func (o *options) MirrorZ() {
	o.Mirror = manipulators.MirrorZ
}

type Modifier func(*options)

func Rotate90() Modifier {
	return options.Rotate90
}

func Rotate180() Modifier {
	return options.Rotate180
}

func Rotate270() Modifier {
	return options.Rotate270
}

func MirrorX() Modifier {
	return options.MirrorX
}

func MirrorZ() Modifier {
	return options.MirrorZ
}

func BlockSwapper(bs func(minecraft.Block) minecraft.Block) Modifier {
	return func(o *options) {
		o.BlockSwapper = bs
	}
}

func noop(b minecraft.Block) minecraft.Block {
	return b
}

func (t Template) Generate(a manipulators.Area, opts ...Modifier) error {
	o := options{
		Rotation:     noop,
		Mirror:       noop,
		BlockSwapper: noop,
	}
	for _, opt := range opts {
		opt(o)
	}
	return nil
}
