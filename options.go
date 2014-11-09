package minecraft

import "github.com/MJKWoolnough/minecraft/nbt"

type Option func(*Level)

// GetLevelName sets the given string to the name of the minecraft level.
func GetLevelName(name *string) Option {
	return func(l *Level) {
		*name = string(*l.levelData.Get("LevelName").Data().(*nbt.String))
	}
}

// LevelName sets the name of the minecraft level.
func LevelName(name string) Option {
	return setOption("LevelName", nbt.NewString(name))
}

// Default minecraft generators
const (
	DefaultGenerator    = "default"
	FlatGenerator       = "flat"
	LargeBiomeGenerator = "largeBiomes"
	AmplifiedGenerator  = "amplified"
	CustomGenerator     = "customized"
	DebugGenerator      = "debug_all_block_states"
)

// Generator sets the generator type
func Generator(generator string) Option {
	return setOption("generatorName", nbt.NewString(generator))
}

// GeneratorOptions sets the generator options for a flat or cusom generator. The syntax is not checked.
func GeneratorOptions(options string) Option {
	return setOption("generatorOptions", nbt.NewString(options))
}

// Seed sets the random seed for the level
func Seed(seed int64) Option {
	return setOption("RandomSeed", nbt.NewLong(seed))
}

// MapFeatures enables/disables map feature generation (villages, strongholds, mineshafts, etc.)
func MapFeatures(mf bool) Option {
	return setOption("MapFeatures", boolToByte(mf))
}

// AllowsCommands enables/disables the cheat commands
func AllowCommands(a bool) Option {
	return setOption("allowCommands", boolToByte(a))
}

// Hardcore enables/disables hardcore mode
func Hardcore(h bool) Option {
	return setOption("hardcore", boolToByte(h))
}

// Game Modes Settings
const (
	Survival int32 = iota
	Creative
	Adventure
	Spectator
)

// GameMode sets the game mode type
func GameMode(gm int32) Option {
	return setOption("GameType", nbt.NewInt(gm))
}

// Difficulty Settings
const (
	Peaceful int8 = iota
	Easy
	Normal
	Hard
)

// Difficulty sets the level difficulty
func Difficulty(d int8) Option {
	return setOption("Difficulty", nbt.NewByte(d))
}

// DifficultyLocked locks the difficulty in game
func DifficultyLocked(l bool) Option {
	return setOption("DifficultyLocked", boolToByte(l))
}

// TicksExisted sets how many ticks have passed in game
func TicksExisted(t int64) Option {
	return setOption("Time", nbt.NewLong(t))
}

const (
	SunRise  = 0
	Noon     = 6000
	SunSet   = 12000
	MidNight = 18000
	Day      = 24000
)

// Time sets the in world time.
func Time(t int64) Option {
	return setOption("DayTime", nbt.NewLong(t))
}

// GetSpawn sets the given x, y, z coordinates to the current spawn point.
func GetSpawn(x, y, z *int32) Option {
	return func(l *Level) {
		xTag, yTag, zTag := l.levelData.Get("SpawnX"), l.levelData.Get("SpawnY"), l.levelData.Get("SpawnZ")
		*x = int32(*xTag.Data().(*nbt.Int))
		*y = int32(*yTag.Data().(*nbt.Int))
		*z = int32(*zTag.Data().(*nbt.Int))
	}
}

// Spawn sets the spawn point to the given coordinates.
func Spawn(x, y, z int32) Option {
	return func(l *Level) {
		l.levelData.Set(nbt.NewTag("SpawnX", nbt.NewInt(x)))
		l.levelData.Set(nbt.NewTag("SpawnY", nbt.NewInt(y)))
		l.levelData.Set(nbt.NewTag("SpawnZ", nbt.NewInt(z)))
		l.changed = true
	}
}

//BorderCenter sets the position of the center of the World Border
func BorderCenter(x, z float64) Option {
	return func(l *Level) {
		l.levelData.Set(nbt.NewTag("BorderCenterX", nbt.NewDouble(x)))
		l.levelData.Set(nbt.NewTag("BorderCenterZ", nbt.NewDouble(z)))
		l.changed = true
	}
}

//BorderSize sets the width of the border
func BorderSize(w float64) Option {
	return setOption("BorderSize", nbt.NewDouble(w))
}

// Raining sets the rain on or off
func Raining(raining bool) Option {
	return setOption("raining", boolToByte(raining))
}

// RainTime sets the time until the rain state changes
func RainTime(time int32) Option {
	return setOption("rainTime", nbt.NewInt(time))
}

// Thundering sets the lightning/thunder on or off
func Thundering(thundering bool) Option {
	return setOption("thundering", boolToByte(thundering))
}

// ThunderTime sets the tune until the tunder state changes
func ThunderTime(time int32) Option {
	return setOption("thunderTime", nbt.NewInt(time))
}

// CommandBlockOutput enables/disables chat echo for command blocks
func CommandBlockOutput(d bool) Option {
	return setGameRule("commandBlockOutput", d)
}

// DayLightCycle enables/disables the day/night cycle
func DayLightCycle(d bool) Option {
	return setGameRule("doDaylightCycle", d)
}

// FireTick enables/disables fire updates, such as spreading and extinguishing
func FireTick(d bool) Option {
	return setGameRule("doFireTick", d)
}

// MobLoot enables/disable mob loot drops
func MobLoot(d bool) Option {
	return setGameRule("doMobLoot", d)
}

// MobSpawning enables/disables mob spawning
func MobSpawning(d bool) Option {
	return setGameRule("doMobSpawning", d)
}

// TileDrops enables/disables the dropping of items upon block breakage
func TileDrops(d bool) Option {
	return setGameRule("doTileDrops", d)
}

// KeepInventory enables/disables the keeping of a players inventory upon death
func KeepInventory(d bool) Option {
	return setGameRule("keepInventory", d)
}

// LogAdminCommands enables/disables the logging of admin commmands to the log
func LogAdminCommands(d bool) Option {
	return setGameRule("logAdminCommands", d)
}

// ModGriefing enables/disables the abilty of mobs to destroy blocks
func MobGriefing(d bool) Option {
	return setGameRule("mobGriefing", d)
}

// HealthRegeneration enables/disables the regeneration of the players health
// when their hunger is high enough
func HealthRegeneration(d bool) Option {
	return setGameRule("naturalRegeneration", d)
}

// CommandFeedback enables/disables the echo for player commands in the chat
func CommandFeedback(d bool) Option {
	return setGameRule("sendCommandFeedback", d)
}

// DeathMessages enables/disables the logging of player deaths to the chat
func DeathMessages(d bool) Option {
	return setGameRule("showDeathMessages", d)
}

// Options allows for the setting (or getting) of multiple options for the
// minecraft level
func (l *Level) Options(options ...Option) {
	for _, option := range options {
		option(l)
	}
}

func setOption(name string, data nbt.Data) Option {
	return func(l *Level) {
		l.levelData.Set(nbt.NewTag(name, data))
		l.changed = true
	}
}

func setGameRule(name string, data bool) Option {
	return func(l *Level) {
		var grc *nbt.Compound
		gr := l.levelData.Get("GameRules")
		if gr != nil {
			grc = gr.Data().(*nbt.Compound)
		}
		if grc == nil {
			grc = nbt.NewCompound([]*nbt.Tag{})
			l.levelData.Set(nbt.NewTag("GameRules", grc))
		}
		var d *nbt.String
		if data {
			d = nbt.NewString("True")
		} else {
			d = nbt.NewString("False")
		}
		grc.Set(nbt.NewTag(name, d))
		l.changed = true
	}
}

func boolToByte(b bool) *nbt.Byte {
	if b {
		return nbt.NewByte(1)
	} else {
		return nbt.NewByte(0)
	}
}
