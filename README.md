# minecraft
--
    import "github.com/MJKWoolnough/minecraft"

Package Minecraft will be a full featured minecraft level editor/viewer.

## Usage

```go
const (
	GZip byte = 1
	Zlib byte = 2
)
```

#### func  CoordsToRegion

```go
func CoordsToRegion(x, z int32) (int32, int32)
```

#### type Biome

```go
type Biome int8
```


```go
const (
	Biome_Auto Biome = iota - 1
	Biome_Ocean
	Biome_Plains
	Biome_Desert
	Biome_ExtremeHills
	Biome_Forest
	Biome_Taiga
	Biome_Swampland
	Biome_River
	Biome_Hell
	Biome_Sky
	Biome_FrozenOcean
	Biome_FrozenRiver
	Biome_IcePlains
	Biome_IceMountains
	Biome_MushroomIsland
	Biome_MushroomIslandShore
	Biome_Beach
	Biome_DesertHills
	Biome_ForestHills
	Biome_TaigaHills
	Biome_ExtremeHillsEdge
	Biome_Jungle
	Biome_JungleHills
)
```

#### func (Biome) Equal

```go
func (b Biome) Equal(e equaler.Equaler) bool
```

#### func (Biome) String

```go
func (b Biome) String() string
```

#### type Block

```go
type Block struct {
	BlockId uint16
	Data    uint8

	Tick bool
}
```


#### func (Block) Equal

```go
func (b Block) Equal(e equaler.Equaler) bool
```

#### func (Block) GetMetadata

```go
func (b Block) GetMetadata() []nbt.Tag
```

#### func (Block) HasMetadata

```go
func (b Block) HasMetadata() bool
```

#### func (Block) IsLiquid

```go
func (b Block) IsLiquid() bool
```

#### func (Block) Opacity

```go
func (b Block) Opacity() uint8
```
Opacity returns how much light is blocked by this block.

#### func (*Block) SetMetadata

```go
func (b *Block) SetMetadata(data []nbt.Tag)
```

#### func (Block) String

```go
func (b Block) String() string
```


#### type Level

```go
type Level struct {
	Path
}
```


#### func  NewLevel

```go
func NewLevel(location Path) (*Level, error)
```

#### func (*Level) Close

```go
func (l *Level) Close()
```

#### func (*Level) GetBiome

```go
func (l *Level) GetBiome(x, z int32) (Biome, error)
```

#### func (*Level) GetBlock

```go
func (l *Level) GetBlock(x, y, z int32) (*Block, error)
```

#### func (Level) GetName

```go
func (l Level) GetName() string
```

#### func (Level) GetSpawn

```go
func (l Level) GetSpawn() (x, y, z int32)
```

#### func (*Level) Save

```go
func (l *Level) Save() error
```

#### func (*Level) SetBiome

```go
func (l *Level) SetBiome(x, z int32, biome Biome) error
```

#### func (*Level) SetBlock

```go
func (l *Level) SetBlock(x, y, z int32, block *Block) error
```

#### func (*Level) SetName

```go
func (l *Level) SetName(name string)
```

#### func (*Level) SetSpawn

```go
func (l *Level) SetSpawn(x, y, z int32)
```


#### type Path

```go
type Path interface {
	// Returns a nil nbt.Tag when chunk does not exists
	GetChunk(int32, int32) (nbt.Tag, error)
	SetChunk(...nbt.Tag) error
	RemoveChunk(int32, int32) error
	ReadLevelDat() (nbt.Tag, error)
	WriteLevelDat(nbt.Tag) error
	GetRegions() [][2]int32
}
```

#### type FilePath

```go
type FilePath struct {
}
```


#### func  NewFilePath

```go
func NewFilePath(dirname string) (*FilePath, error)
```
NewFilePath constructs a new directory based path to read from.

#### func (*FilePath) Defrag

```go
func (p *FilePath) Defrag(x, z int32) error
```

#### func (*FilePath) GetChunk

```go
func (p *FilePath) GetChunk(x, z int32) (nbt.Tag, error)
```

#### func (*FilePath) GetRegions

```go
func (p *FilePath) GetRegions() [][2]int32
```
GetRegions returns a list of region x,z coords of all generated regions.

#### func (*FilePath) Lock

```go
func (p *FilePath) Lock()
```
Lock will retake the lock file if it has been lost. May cause corruption.

#### func (*FilePath) ReadLevelDat

```go
func (p *FilePath) ReadLevelDat() (nbt.Tag, error)
```

#### func (*FilePath) RemoveChunk

```go
func (p *FilePath) RemoveChunk(x, z int32) error
```

#### func (*FilePath) SetChunk

```go
func (p *FilePath) SetChunk(data ...nbt.Tag) error
```

#### func (*FilePath) Update

```go
func (p *FilePath) Update(filname string, mode uint8)
```
Update tracks the lock file for updates to remove the lock.

#### func (*FilePath) WriteLevelDat

```go
func (p *FilePath) WriteLevelDat(data nbt.Tag) error
```

#### type MemPath

```go
type MemPath struct {
}
```


#### func  NewMemPath

```go
func NewMemPath() *MemPath
```

#### func (*MemPath) GetChunk

```go
func (m *MemPath) GetChunk(x, z int32) (nbt.Tag, error)
```

#### func (*MemPath) GetRegions

```go
func (m *MemPath) GetRegions() [][2]int32
```

#### func (*MemPath) ReadLevelDat

```go
func (m *MemPath) ReadLevelDat() (nbt.Tag, error)
```

#### func (*MemPath) RemoveChunk

```go
func (m *MemPath) RemoveChunk(x, z int32) error
```

#### func (*MemPath) SetChunk

```go
func (m *MemPath) SetChunk(data ...nbt.Tag) error
```

#### func (*MemPath) WriteLevelDat

```go
func (m *MemPath) WriteLevelDat(data nbt.Tag) error
```

### Errors


#### type ConflictError

```go
type ConflictError struct {
	X, Z int32
}
```


#### func (ConflictError) Error

```go
func (c ConflictError) Error() string
```

#### type ExpectedData

```go
type ExpectedData struct{}
```


#### func (ExpectedData) Error

```go
func (e ExpectedData) Error() string
```


#### type MissingTagError

```go
type MissingTagError struct {
}
```


#### func (MissingTagError) Error

```go
func (m MissingTagError) Error() string
```

#### type NoLock

```go
type NoLock struct{}
```


#### func (NoLock) Error

```go
func (n NoLock) Error() string
```

#### type OOB

```go
type OOB struct{}
```


#### func (OOB) Error

```go
func (o OOB) Error() string
```


#### type UnexpectedValue

```go
type UnexpectedValue struct {
}
```


#### func (UnexpectedValue) Error

```go
func (u UnexpectedValue) Error() string
```

#### type UnknownCompression

```go
type UnknownCompression struct {
}
```


#### func (UnknownCompression) Error

```go
func (u UnknownCompression) Error() string
```

#### type WrongTypeError

```go
type WrongTypeError struct {
}
```


#### func (WrongTypeError) Error

```go
func (w WrongTypeError) Error() string
```
