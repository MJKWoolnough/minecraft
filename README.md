# minecraft
--
    import "github.com/MJKWoolnough/minecraft"

Package minecraft will be a full featured minecraft level editor/viewer.

## Usage

```go
const (
	DefaultGenerator    = "default"
	FlatGenerator       = "flat"
	LargeBiomeGenerator = "largeBiomes"
	AmplifiedGenerator  = "amplified"
	CustomGenerator     = "customized"
	DebugGenerator      = "debug_all_block_states"
)
```
Default minecraft generators

```go
const (
	Survival int32 = iota
	Creative
	Adventure
	Spectator
)
```
Game Modes Settings

```go
const (
	Peaceful int8 = iota
	Easy
	Normal
	Hard
)
```
Difficulty Settings

```go
const (
	SunRise  = 0
	Noon     = 6000
	SunSet   = 12000
	MidNight = 18000
	Day      = 24000
)
```
Time-of-day convenience constants

```go
const (
	GZip byte = 1
	Zlib byte = 2
)
```
Compression convenience constants

```go
var (
	// ErrOOB is an error returned when sanity checking section data
	ErrOOB = errors.New("Received Out-of-bounds error")
	// ErrNoLock is an error returns by path types to indicate that the lock on the
	// minecraft level has been locked and needs reinstating to continue
	ErrNoLock = errors.New("lost lock on files")
)
```

```go
var (
	// TransparentBlocks is a slice of the block ids that are transparent.
	// This is used in lighting calculations and is user overrideable for custom
	// blocks
	TransparentBlocks = TransparentBlockList{0, 6, 18, 20, 26, 27, 28, 29, 30, 31, 33, 34, 37, 38, 39, 40, 50, 51, 52, 54, 55, 59, 63, 64, 65, 66, 69, 70, 71, 75, 76, 77, 78, 79, 81, 83, 85, 90, 92, 93, 94, 96, 102, 106, 107, 117, 118, 119, 120, 750}
	// LightBlocks is a map of block ids to the amount of light they give off
	LightBlocks = LightBlockList{
		10:  15,
		11:  15,
		39:  1,
		50:  14,
		51:  15,
		62:  13,
		74:  13,
		76:  7,
		89:  15,
		90:  11,
		91:  15,
		94:  9,
		117: 1,
		119: 15,
		120: 1,
		122: 1,
		124: 15,
		130: 7,
		138: 15,
	}
)
```

#### type Biome

```go
type Biome uint8
```

Biome is a convenience type for biomes

```go
const (
	Ocean                Biome = 0
	Plains               Biome = 1
	Desert               Biome = 2
	ExtremeHills         Biome = 3
	Forest               Biome = 4
	Taiga                Biome = 5
	Swampland            Biome = 6
	River                Biome = 7
	Hell                 Biome = 8
	Sky                  Biome = 9
	FrozenOcean          Biome = 10
	FrozenRiver          Biome = 11
	IcePlains            Biome = 12
	IceMountains         Biome = 13
	MushroomIsland       Biome = 14
	MushroomIslandShore  Biome = 15
	Beach                Biome = 16
	DesertHills          Biome = 17
	ForestHills          Biome = 18
	TaigaHills           Biome = 19
	ExtremeHillsEdge     Biome = 20
	Jungle               Biome = 21
	JungleHills          Biome = 22
	JungleEdge           Biome = 23
	DeepOcean            Biome = 24
	StoneBeach           Biome = 25
	ColdBeach            Biome = 26
	BirchForest          Biome = 27
	BirchForestHills     Biome = 28
	RoofedForest         Biome = 29
	ColdTaiga            Biome = 30
	ColdTaigaHills       Biome = 31
	MegaTaiga            Biome = 32
	MegaTaigaHills       Biome = 33
	ExtremeHillsPlus     Biome = 34
	Savanna              Biome = 35
	SavannaPlateau       Biome = 36
	Mesa                 Biome = 37
	MesaPlateauF         Biome = 38
	MesaPlateau          Biome = 39
	SunflowerPlains      Biome = 129
	DeserM               Biome = 130
	ExtremeHillsM        Biome = 131
	FlowerForest         Biome = 132
	TaigaM               Biome = 133
	SwamplandM           Biome = 134
	IcePlainsSpikes      Biome = 140
	JungleM              Biome = 149
	JungleEdgeM          Biome = 151
	BirchForestM         Biome = 155
	BirchForestHillsM    Biome = 156
	RoofedForestM        Biome = 157
	ColdTaigaM           Biome = 158
	MegaSpruceTaiga      Biome = 160
	MegaSpruceTaigaHills Biome = 161
	ExtremeHillsPlusM    Biome = 162
	SavannaM             Biome = 163
	SavannaPlateauM      Biome = 164
	MesaBryce            Biome = 165
	MesaPlateauFM        Biome = 166
	MesaPlateauM         Biome = 167
	AutoBiome            Biome = 255
)
```
Biome constants

#### func (Biome) Equal

```go
func (b Biome) Equal(e equaler.Equaler) bool
```
Equal is an implementation of the equaler.Equaler interface

#### func (Biome) String

```go
func (b Biome) String() string
```

#### type Block

```go
type Block struct {
	ID   uint16
	Data uint8
}
```

Block is a type that represents the full information for a block, id, data,
metadata and scheduled tick data

#### func (*Block) AddTicks

```go
func (b *Block) AddTicks(t ...Tick)
```
AddTicks adds one or more scheduled ticks to the block

#### func (Block) Equal

```go
func (b Block) Equal(e equaler.Equaler) bool
```
Equal is an implementation of the equaler.Equaler interface

#### func (Block) EqualBlock

```go
func (b Block) EqualBlock(c Block) bool
```
EqualBlock checks for equality between the two blocks

#### func (Block) GetMetadata

```go
func (b Block) GetMetadata() nbt.Compound
```
GetMetadata returns a copy of the metadata for this block, or nil is it has none

#### func (Block) GetTicks

```go
func (b Block) GetTicks() []Tick
```
GetTicks returns all of the scheduled ticks for a block

#### func (Block) HasMetadata

```go
func (b Block) HasMetadata() bool
```
HasMetadata returns true the the block contains extended metadata

#### func (Block) HasTicks

```go
func (b Block) HasTicks() bool
```
HasTicks returns true if the block has any scheduled ticks

#### func (Block) IsLiquid

```go
func (b Block) IsLiquid() bool
```
IsLiquid returns true if the block id matches a liquid

#### func (Block) Light

```go
func (b Block) Light() uint8
```
Light returns how much light is generated by this block.

#### func (Block) Opacity

```go
func (b Block) Opacity() uint8
```
Opacity returns how much light is blocked by this block.

#### func (*Block) SetMetadata

```go
func (b *Block) SetMetadata(data nbt.Compound)
```
SetMetadata sets the blocks metadata to a copy of the given metadata

#### func (*Block) SetTicks

```go
func (b *Block) SetTicks(t []Tick)
```
SetTicks sets the scheduled ticks for the block, replacing any existing ones

#### func (Block) String

```go
func (b Block) String() string
```

#### type ConflictError

```go
type ConflictError struct {
	X, Z int32
}
```

ConflictError is an error return by SetChunk when trying to save a single chunk
multiple times during the same save operation

#### func (ConflictError) Error

```go
func (c ConflictError) Error() string
```

#### type FilePath

```go
type FilePath struct {
}
```

FilePath implements the Path interface and provides a standard minecraft save
format.

#### func  NewFilePath

```go
func NewFilePath(dirname string) (*FilePath, error)
```
NewFilePath constructs a new directory based path to read from.

#### func  NewFilePathDimension

```go
func NewFilePathDimension(dirname string, dimension int) (*FilePath, error)
```
NewFilePathDimension create a new FilePath, but with the option to set the
dimension that chunks are loaded from.

Example. Dimension -1 == The Nether

    Dimension  1 == The End

#### func (*FilePath) Defrag

```go
func (p *FilePath) Defrag(x, z int32) error
```
Defrag rewrites a region file to reduce wasted space.

#### func (*FilePath) GetChunk

```go
func (p *FilePath) GetChunk(x, z int32) (nbt.Tag, error)
```
GetChunk returns the chunk at chunk coords x, z.

#### func (*FilePath) GetChunks

```go
func (p *FilePath) GetChunks(x, z int32) ([][2]int32, error)
```
GetChunks returns a list of all chunks within a region with coords x,z

#### func (*FilePath) GetRegions

```go
func (p *FilePath) GetRegions() [][2]int32
```
GetRegions returns a list of region x,z coords of all generated regions.

#### func (*FilePath) HasLock

```go
func (p *FilePath) HasLock() bool
```
HasLock returns whether or not another program has taken the lock.

#### func (*FilePath) Lock

```go
func (p *FilePath) Lock() error
```
Lock will retake the lock file if it has been lost. May cause corruption.

#### func (*FilePath) ReadLevelDat

```go
func (p *FilePath) ReadLevelDat() (nbt.Tag, error)
```
ReadLevelDat returns the level data.

#### func (*FilePath) RemoveChunk

```go
func (p *FilePath) RemoveChunk(x, z int32) error
```
RemoveChunk deletes the chunk at chunk coords x, z.

#### func (*FilePath) SetChunk

```go
func (p *FilePath) SetChunk(data ...nbt.Tag) error
```
SetChunk saves multiple chunks at once, possibly returning a MultiError if
multiple errors were encountered.

#### func (*FilePath) WriteLevelDat

```go
func (p *FilePath) WriteLevelDat(data nbt.Tag) error
```
WriteLevelDat Writes the level data.

#### type FilePathSetError

```go
type FilePathSetError struct {
	X, Z int32
	Err  error
}
```

FilePathSetError is an error returned from SetChunk when some error is returned
either from converting the nbt or saving it

#### func (FilePathSetError) Error

```go
func (f FilePathSetError) Error() string
```

#### type Level

```go
type Level struct {
}
```

Level is the base type for minecraft data, all data for a minecraft level is
either store in, or accessed from, this type

#### func  NewLevel

```go
func NewLevel(location Path) (*Level, error)
```
NewLevel creates/Loads a minecraft level from the given path.

#### func (*Level) AllowCommands

```go
func (l *Level) AllowCommands(a bool)
```
AllowCommands enables/disables the cheat commands

#### func (*Level) BorderCenter

```go
func (l *Level) BorderCenter(x, z float64)
```
BorderCenter sets the position of the center of the World Border

#### func (*Level) BorderSize

```go
func (l *Level) BorderSize(w float64)
```
BorderSize sets the width of the border

#### func (*Level) Close

```go
func (l *Level) Close()
```
Close closes all open chunks, but does not save them.

#### func (*Level) CommandBlockOutput

```go
func (l *Level) CommandBlockOutput(d bool)
```
CommandBlockOutput enables/disables chat echo for command blocks

#### func (*Level) CommandFeedback

```go
func (l *Level) CommandFeedback(d bool)
```
CommandFeedback enables/disables the echo for player commands in the chat

#### func (*Level) DayLightCycle

```go
func (l *Level) DayLightCycle(d bool)
```
DayLightCycle enables/disables the day/night cycle

#### func (*Level) DeathMessages

```go
func (l *Level) DeathMessages(d bool)
```
DeathMessages enables/disables the logging of player deaths to the chat

#### func (*Level) Difficulty

```go
func (l *Level) Difficulty(d int8)
```
Difficulty sets the level difficulty

#### func (*Level) DifficultyLocked

```go
func (l *Level) DifficultyLocked(dl bool)
```
DifficultyLocked locks the difficulty in game

#### func (*Level) FireTick

```go
func (l *Level) FireTick(d bool)
```
FireTick enables/disables fire updates, such as spreading and extinguishing

#### func (*Level) GameMode

```go
func (l *Level) GameMode(gm int32)
```
GameMode sets the game mode type

#### func (*Level) Generator

```go
func (l *Level) Generator(generator string)
```
Generator sets the generator type

#### func (*Level) GeneratorOptions

```go
func (l *Level) GeneratorOptions(options string)
```
GeneratorOptions sets the generator options for a flat or custom generator. The
syntax is not checked.

#### func (*Level) GetBiome

```go
func (l *Level) GetBiome(x, z int32) (Biome, error)
```
GetBiome returns the biome for the column x, z.

#### func (*Level) GetBlock

```go
func (l *Level) GetBlock(x, y, z int32) (Block, error)
```
GetBlock gets the block at coordinates x, y, z.

#### func (*Level) GetHeight

```go
func (l *Level) GetHeight(x, z int32) (int32, error)
```
GetHeight returns the y coordinate for the highest non-transparent block at
column x, z.

#### func (*Level) GetLevelName

```go
func (l *Level) GetLevelName() string
```
GetLevelName sets the given string to the name of the minecraft level.

#### func (*Level) GetSpawn

```go
func (l *Level) GetSpawn() (x int32, y int32, z int32)
```
GetSpawn sets the given x, y, z coordinates to the current spawn point.

#### func (*Level) Hardcore

```go
func (l *Level) Hardcore(h bool)
```
Hardcore enables/disables hardcore mode

#### func (*Level) HealthRegeneration

```go
func (l *Level) HealthRegeneration(d bool)
```
HealthRegeneration enables/disables the regeneration of the players health when
their hunger is high enough

#### func (*Level) KeepInventory

```go
func (l *Level) KeepInventory(d bool)
```
KeepInventory enables/disables the keeping of a players inventory upon death

#### func (*Level) LevelName

```go
func (l *Level) LevelName(name string)
```
LevelName sets the name of the minecraft level.

#### func (*Level) LogAdminCommands

```go
func (l *Level) LogAdminCommands(d bool)
```
LogAdminCommands enables/disables the logging of admin commmands to the log

#### func (*Level) MapFeatures

```go
func (l *Level) MapFeatures(mf bool)
```
MapFeatures enables/disables map feature generation (villages, strongholds,
mineshafts, etc.)

#### func (*Level) MobGriefing

```go
func (l *Level) MobGriefing(d bool)
```
MobGriefing enables/disables the ability of mobs to destroy blocks

#### func (*Level) MobLoot

```go
func (l *Level) MobLoot(d bool)
```
MobLoot enables/disable mob loot drops

#### func (*Level) MobSpawning

```go
func (l *Level) MobSpawning(d bool)
```
MobSpawning enables/disables mob spawning

#### func (*Level) RainTime

```go
func (l *Level) RainTime(time int32)
```
RainTime sets the time until the rain state changes

#### func (*Level) Raining

```go
func (l *Level) Raining(raining bool)
```
Raining sets the rain on or off

#### func (*Level) Save

```go
func (l *Level) Save() error
```
Save saves all open chunks, but does not close them.

#### func (*Level) Seed

```go
func (l *Level) Seed(seed int64)
```
Seed sets the random seed for the level

#### func (*Level) SetBiome

```go
func (l *Level) SetBiome(x, z int32, biome Biome) error
```
SetBiome sets the biome for the column x, z.

#### func (*Level) SetBlock

```go
func (l *Level) SetBlock(x, y, z int32, block Block) error
```
SetBlock sets the block at coordinates x, y, z. Also processes any lighting
updates if applicable.

#### func (*Level) Spawn

```go
func (l *Level) Spawn(x, y, z int32)
```
Spawn sets the spawn point to the given coordinates.

#### func (*Level) ThunderTime

```go
func (l *Level) ThunderTime(time int32)
```
ThunderTime sets the tune until the thunder state changes

#### func (*Level) Thundering

```go
func (l *Level) Thundering(thundering bool)
```
Thundering sets the lightning/thunder on or off

#### func (*Level) TicksExisted

```go
func (l *Level) TicksExisted(t int64)
```
TicksExisted sets how many ticks have passed in game

#### func (*Level) TileDrops

```go
func (l *Level) TileDrops(d bool)
```
TileDrops enables/disables the dropping of items upon block breakage

#### func (*Level) Time

```go
func (l *Level) Time(t int64)
```
Time sets the in world time.

#### type LightBlockList

```go
type LightBlockList map[uint16]uint8
```

LightBlockList is a map of block ids to the amount of light they give off

#### func (LightBlockList) Add

```go
func (l LightBlockList) Add(blockID uint16, light uint8) bool
```
Add is a convenience method for the light block list. It adds a new block id to
the list with its corresponding light level

#### func (LightBlockList) Remove

```go
func (l LightBlockList) Remove(blockID uint16) bool
```
Remove is a convenience method to remove a block id from the light block list

#### type MemPath

```go
type MemPath struct {
}
```

MemPath is an in memory minecraft level format that implements the Path
interface.

#### func  NewMemPath

```go
func NewMemPath() *MemPath
```
NewMemPath creates a new MemPath implementation.

#### func (*MemPath) GetChunk

```go
func (m *MemPath) GetChunk(x, z int32) (nbt.Tag, error)
```
GetChunk returns the chunk at chunk coords x, z.

#### func (*MemPath) ReadLevelDat

```go
func (m *MemPath) ReadLevelDat() (nbt.Tag, error)
```
ReadLevelDat Returns the level data.

#### func (*MemPath) RemoveChunk

```go
func (m *MemPath) RemoveChunk(x, z int32) error
```
RemoveChunk deletes the chunk at chunk coords x, z.

#### func (*MemPath) SetChunk

```go
func (m *MemPath) SetChunk(data ...nbt.Tag) error
```
SetChunk saves multiple chunks at once.

#### func (*MemPath) WriteLevelDat

```go
func (m *MemPath) WriteLevelDat(data nbt.Tag) error
```
WriteLevelDat Writes the level data.

#### type MissingTagError

```go
type MissingTagError struct {
	TagName string
}
```

MissingTagError is an error type returned when an expected tag is not found

#### func (MissingTagError) Error

```go
func (m MissingTagError) Error() string
```

#### type MultiError

```go
type MultiError struct {
	Errors []error
}
```

MultiError is an error type that contains multiple errors

#### func (MultiError) Error

```go
func (m MultiError) Error() string
```

#### type Option

```go
type Option func(*Level)
```

Option is a function used to set an option for a minecraft level struct

#### type Path

```go
type Path interface {
	// Returns an empty nbt.Tag (TagEnd) when chunk does not exists
	GetChunk(int32, int32) (nbt.Tag, error)
	SetChunk(...nbt.Tag) error
	RemoveChunk(int32, int32) error
	ReadLevelDat() (nbt.Tag, error)
	WriteLevelDat(nbt.Tag) error
}
```

The Path interface allows the minecraft level to be created from/saved to
different formats.

#### type Tick

```go
type Tick struct {
	I, T, P int32
}
```

Tick is a type that represents a scheduled tick

#### type TransparentBlockList

```go
type TransparentBlockList []uint16
```

TransparentBlockList is a slice of the block ids that are transparent.

#### func (*TransparentBlockList) Add

```go
func (t *TransparentBlockList) Add(blockID uint16) bool
```
Add is a convenience method for the transparent block list. It adds a new block
id to the list, making sure to not add duplicates

#### func (*TransparentBlockList) Remove

```go
func (t *TransparentBlockList) Remove(blockID uint16) bool
```
Remove is a convenience method to remove a block id from the transparent block
list

#### type UnexpectedValue

```go
type UnexpectedValue struct {
	TagName, Expecting, Got string
}
```

UnexpectedValue is an error returned from chunk loading during sanity checking

#### func (UnexpectedValue) Error

```go
func (u UnexpectedValue) Error() string
```

#### type UnknownCompression

```go
type UnknownCompression struct {
	Code byte
}
```

UnknownCompression is an error returned by path types when it encounters a
compression scheme it is not prepared to handle or an unkown compression scheme

#### func (UnknownCompression) Error

```go
func (u UnknownCompression) Error() string
```

#### type WrongTypeError

```go
type WrongTypeError struct {
	TagName        string
	Expecting, Got nbt.TagID
}
```

WrongTypeError is an error returned when a nbt tag has an unexpected type

#### func (WrongTypeError) Error

```go
func (w WrongTypeError) Error() string
```
