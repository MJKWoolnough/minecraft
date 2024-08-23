# maps
--
    import "vimagination.zapto.org/minecraft/maps"


## Usage

```go
var (
	ErrInvalidDimensions = errors.New("cannot encode an image with the given dimensions")
)
```
Errors.

#### func  BlockColor

```go
func BlockColor(b minecraft.Block) color.Color
```
BlockColor is the standard block-to-colour func.

#### func  Config

```go
func Config(r io.Reader) (image.Config, error)
```
Config reader the configuration for an uncompressed Minecraft map.

Minecraft maps are gzip compressed, so the reader given to this func should be
wrapped in gzip.NewReader.

#### func  Decode

```go
func Decode(r io.Reader) (image.Image, error)
```
Decode takes a reader for an uncompressed Minecraft map.

Minecraft maps are gzip compressed, so the reader given to this func should be
wrapped in gzip.NewReader.

#### func  Encode

```go
func Encode(w io.Writer, i image.Image) error
```
Encode writes an image an as uncompressed Minecraft map.

As Minecraft expects the map to be gzip compressed, the Writer should be the
wrapped in gzip.NewWriter.

#### type Encoder

```go
type Encoder struct {
	Scale, Dimension int8
	CenterX, CenterZ int32
}
```

Encoder lets you specify options for the Minecraft map.

#### func (*Encoder) Encode

```go
func (e *Encoder) Encode(w io.Writer, im image.Image) error
```
Encode writes an image an as uncompressed Minecraft map.

As Minecraft expects the map to be gzip compressed, the Writer should be the
wrapped in gzip.NewWriter.

#### type Image

```go
type Image struct {
}
```

Image represents a Minecraft Map.

#### func  NewMap

```go
func NewMap(l *minecraft.Level, bounds image.Rectangle, options ...Option) Image
```
NewMap creates an image from a Minecraft level.

#### func (Image) At

```go
func (i Image) At(x, z int) color.Color
```
At returns the colour at the specified coords.

#### func (Image) Bounds

```go
func (i Image) Bounds() image.Rectangle
```
Bounds returns the dimensions of the map.

#### func (Image) ColorModel

```go
func (Image) ColorModel() color.Model
```
ColorModel returns the palette for the map.

#### type Option

```go
type Option func(*Image)
```

Option represents a optional parameter for a map type.

#### func  ColorFunc

```go
func ColorFunc(c func(minecraft.Block) color.Color) Option
```
ColorFunc is an option for NewMap that specifies what colour blocks are painted
as.

#### func  FixedY

```go
func FixedY(y int32) Option
```
FixedY is an options to fix the Y-coord of the blocks to be read. By default the
highest, non-transparent block is used.

#### func  Scale

```go
func Scale(s uint8) Option
```
Scale sets the scale the map is to be rendered at.
