package maps

import (
	"image"
	"io"

	"github.com/MJKWoolnough/minecraft/nbt"
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
		return nil, nbt.WrongTag{nbt.TagCompound, t.TagID()}
	}

	d := t.Data().(nbt.Compound).Get("data")
	if d.TagID() != nbt.TagCompound {
		return nil, nbt.WrongTag{nbt.TagCompound, d.TagID()}
	}
	return d.Data().(nbt.Compound), nil
}

func getDimensions(d nbt.Compound) image.Rectangle {
	var width, height int
	w := d.Get("width")
	if w.TagID() == nbt.TagShort {
		width = int(w.Data().(nbt.Short))
	}
	h := d.Get("height")
	if h.TagID() == nbt.TagShort {
		height = int(h.Data().(nbt.Short))
	}
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	}
}

func Decode(r io.Reader) (image.Image, error) {
	d, err := readData(r)
	if err != nil {
		return nil, err
	}

	c := d.Get("colors")
	if c.TagID() != nbt.TagByteArray {
		return nil, nbt.WrongTag{nbt.TagByteArray, c.TagID()}
	}

	rect := getDimensions(d)

	return &image.Paletted{
		Pix:     c.Data().(nbt.ByteArray).Bytes(),
		Stride:  rect.Max.X,
		Rect:    rect,
		Palette: palette,
	}, nil
}

type config struct {
	Data struct {
		Width  int16 `nbt:"width"`
		Height int16 `nbt:"height"`
	} `nbt:"data"`
}

func Config(r io.Reader) (image.Config, error) {
	var c config
	_, err := nbt.RDecode(r, &c)
	if err != nil {
		return image.Config{}, err
	}
	return image.Config{
		ColorModel: palette,
		Width:      int(c.Data.Width),
		Height:     int(c.Data.Height),
	}, nil
}
