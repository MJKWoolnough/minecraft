package maps // import "vimagination.zapto.org/minecraft/maps"

import (
	"image"
	"io"

	"vimagination.zapto.org/minecraft/nbt"
)

func init() {
	image.RegisterFormat("minecraftmap", "\x0a\x00\x00\x0a\x00\x04data", Decode, Config)
}

func readData(r io.Reader) (nbt.Compound, error) {
	t, err := nbt.Decode(r)
	if err != nil {
		return nil, err
	}

	if t.TagID() != nbt.TagCompound {
		return nil, nbt.WrongTag{
			Expecting: nbt.TagCompound,
			Got:       t.TagID(),
		}
	}

	d := t.Data().(nbt.Compound).Get("data")
	if d.TagID() != nbt.TagCompound {
		return nil, nbt.WrongTag{
			Expecting: nbt.TagCompound,
			Got:       d.TagID(),
		}
	}

	return d.Data().(nbt.Compound), nil
}

func getDimensions(d nbt.Compound) image.Rectangle {
	var width, height int

	if w := d.Get("width"); w.TagID() == nbt.TagShort {
		width = int(w.Data().(nbt.Short))
	}

	if h := d.Get("height"); h.TagID() == nbt.TagShort {
		height = int(h.Data().(nbt.Short))
	}

	return image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	}
}

// Decode takes a reader for an uncompressed Minecraft map.
//
// Minecraft maps are gzip compressed, so the reader given to this func should
// be wrapped in gzip.NewReader.
func Decode(r io.Reader) (image.Image, error) {
	d, err := readData(r)
	if err != nil {
		return nil, err
	}

	c := d.Get("colors")
	if c.TagID() != nbt.TagByteArray {
		return nil, nbt.WrongTag{
			Expecting: nbt.TagByteArray,
			Got:       c.TagID(),
		}
	}

	rect := getDimensions(d)

	return &image.Paletted{
		Pix:     c.Data().(nbt.ByteArray).Bytes(),
		Stride:  rect.Max.X,
		Rect:    rect,
		Palette: palette,
	}, nil
}

// Config reader the configuration for an uncompressed Minecraft map.
//
// Minecraft maps are gzip compressed, so the reader given to this func should
// be wrapped in gzip.NewReader.
func Config(r io.Reader) (image.Config, error) {
	d, err := readData(r)
	if err != nil {
		return image.Config{}, err
	}

	rect := getDimensions(d)

	return image.Config{
		ColorModel: palette,
		Width:      int(rect.Max.X),
		Height:     int(rect.Max.Y),
	}, nil
}
