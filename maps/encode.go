package maps

import (
	"errors"
	"image"
	"io"

	"github.com/MJKWoolnough/minecraft/nbt"
)

// Encode writes an image an as uncompressed Minecraft map.
//
// As Minecraft expects the map to be gzip compressed, the Writer sohuld be the
// wrapped in gzip.NewWriter
func Encode(w io.Writer, i image.Image) error {
	e := Encoder{
		Dimension: -128,
		Scale:     3,
	}
	return e.Encode(w, i)
}

// Encoder lets you specify options for the Minecraft map
type Encoder struct {
	Scale, Dimension int8
	CenterX, CenterZ int32
}

// Encode writes an image an as uncompressed Minecraft map.
//
// As Minecraft expects the map to be gzip compressed, the Writer sohuld be the
// wrapped in gzip.NewWriter
func (e *Encoder) Encode(w io.Writer, im image.Image) error {
	width := im.Bounds().Dx()
	height := im.Bounds().Dy()
	if width > 0xffff || height > 0xffff || width < 0 || height < 0 {
		return InvalidDimensions
	}

	colours := make(nbt.ByteArray, 0, width*height)
	for j := im.Bounds().Min.Y; j < im.Bounds().Max.Y; j++ {
		for i := im.Bounds().Min.X; i < im.Bounds().Max.X; i++ {
			colours = append(colours, int8(palette.Index(im.At(i, j))))
		}
	}

	return nbt.Encode(w, nbt.NewTag("", nbt.Compound{
		nbt.NewTag("data", nbt.Compound{
			nbt.NewTag("scale", nbt.Byte(e.Scale)),
			nbt.NewTag("dimension", nbt.Byte(e.Dimension)),
			nbt.NewTag("height", nbt.Short(height)),
			nbt.NewTag("width", nbt.Short(width)),
			nbt.NewTag("xCenter", nbt.Int(e.CenterX)),
			nbt.NewTag("zCenter", nbt.Int(e.CenterZ)),
			nbt.NewTag("colors", colours),
		}),
	}))
}

// Errors

var InvalidDimensions = errors.New("cannot encode an image with the given dimensions")
