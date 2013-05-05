# minecraft
--
    import "github.com/MJKWoolnough/minecraft"

A first run at a minecraft level editor.

## Usage

#### type Biome

```go
type Biome byte
```

Needs Implementation

#### func  NewBiome

```go
func NewBiome(biomeId uint8) Biome
```

#### func (*Biome) Equal

```go
func (b *Biome) Equal(e equaler.Equaler) bool
```

#### func (*Biome) String

```go
func (b *Biome) String() string
```

#### type Block

```go
type Block interface {
	BlockId() uint8
	Add() uint8
	Data() uint8
	Opacity() uint8
	IsLiquid() bool
	HasMetadata() bool
	GetMetadata() []nbtparser.NBTTag
	SetMetadata([]nbtparser.NBTTag)
	Tick(bool)
	ToTick() bool
	equaler.Equaler
}
```

Block allows access to the data of a minecraft block.

```go
var (
	BlockAir Block // BlockAir is useful as a nil block.
)
```

#### func  NewBlock

```go
func NewBlock(blockId, add, data uint8) Block
```

#### type Chunk

```go
type Chunk interface {
	io.WriterTo
	Get(uint8, uint8, uint8) Block
	Set(uint8, uint8, uint8, Block)
	GetSkyLight(uint8, uint8, uint8) uint8
	SetSkyLight(uint8, uint8, uint8, uint8)
	Opacity(uint8, uint8, uint8) uint8
	IsEmpty() bool
	GetMetadata(uint8, uint8, uint8) []nbtparser.NBTTag
	SetMetadata(uint8, uint8, uint8, []nbtparser.NBTTag)
	GetTick(x, y, z uint8) bool
	SetTick(x, y, z uint8, data []nbtparser.NBTTag)
	HasChanged() bool
	Compress()
}
```


#### func  LoadChunk

```go
func LoadChunk(data io.Reader) (Chunk, error)
```

#### func  NewChunk

```go
func NewChunk(x, z int32) Chunk
```

#### type Level

```go
type Level interface {
	Get(int32, int32, int32) Block
	Set(int32, int32, int32, Block)
	GetName() string
	SetName(string)
	GetSpawn() (int32, int32, int32)
	SetSpawn(int32, int32, int32)
	SaveLevelData(string) error
	ExportOpenRegions(string) error
	CloseOpenRegions()
	Compress()
	GetSkyLight(int32, int32, int32) uint8
	SetSkyLight(int32, int32, int32, uint8)
	Opacity(int32, int32, int32) uint8
}
```


#### func  LoadLevel

```go
func LoadLevel(path string) (Level, error)
```

#### func  NewLevel

```go
func NewLevel(name string) Level
```

#### type Region

```go
type Region interface {
	Get(int32, int32, int32) Block
	Set(int32, int32, int32, Block)
	GetSkyLight(int32, int32, int32) uint8
	SetSkyLight(int32, int32, int32, uint8)
	Opacity(int32, int32, int32) uint8
	Export(io.WriteSeeker) error
	HasChanged() bool
	Compress()
}
```


#### func  LoadRegion

```go
func LoadRegion(data io.ReadSeeker) (Region, error)
```

#### func  NewRegion

```go
func NewRegion() (Region, error)
```

#### type Section

```go
type Section interface {
	Get(uint8, uint8, uint8) Block
	Set(uint8, uint8, uint8, Block)
	GetSkyLight(uint8, uint8, uint8) uint8
	SetSkyLight(uint8, uint8, uint8, uint8)
	Opacity(uint8, uint8, uint8) uint8
	GetY() uint8
	IsEmpty() bool
	Data() *nbtparser.NBTTagCompound
}
```


#### func  LoadSection

```go
func LoadSection(sectionData *nbtparser.NBTTagCompound) (Section, error)
```

#### func  NewSection

```go
func NewSection(y byte) (Section, error)
```
