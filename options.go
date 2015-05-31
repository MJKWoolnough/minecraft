package minecraft

import "github.com/MJKWoolnough/minecraft/nbt"

// Option is a function used to set an option for a minecraft level struct
type Option func(*Level)

// GetLevelName sets the given string to the name of the minecraft level.
func (l *Level) GetLevelName() string {
	return string(l.levelData.Get("LevelName").Data().(nbt.String))
}

// LevelName sets the name of the minecraft level.
func (l *Level) LevelName(name string) {
	l.setOption("LevelName", nbt.String(name))
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
func (l *Level) Generator(generator string) {
	l.setOption("generatorName", nbt.String(generator))
}

// GeneratorOptions sets the generator options for a flat or cusom generator. The syntax is not checked.
func (l *Level) GeneratorOptions(options string) {
	l.setOption("generatorOptions", nbt.String(options))
}

// Seed sets the random seed for the level
func (l *Level) Seed(seed int64) {
	l.setOption("RandomSeed", nbt.Long(seed))
}

// MapFeatures enables/disables map feature generation (villages, strongholds, mineshafts, etc.)
func (l *Level) MapFeatures(mf bool) {
	l.setOption("MapFeatures", boolToByte(mf))
}

// AllowCommands enables/disables the cheat commands
func (l *Level) AllowCommands(a bool) {
	l.setOption("allowCommands", boolToByte(a))
}

// Hardcore enables/disables hardcore mode
func (l *Level) Hardcore(h bool) {
	l.setOption("hardcore", boolToByte(h))
}

// Game Modes Settings
const (
	Survival int32 = iota
	Creative
	Adventure
	Spectator
)

// GameMode sets the game mode type
func (l *Level) GameMode(gm int32) {
	l.setOption("GameType", nbt.Int(gm))
}

// Difficulty Settings
const (
	Peaceful int8 = iota
	Easy
	Normal
	Hard
)

// Difficulty sets the level difficulty
func (l *Level) Difficulty(d int8) {
	l.setOption("Difficulty", nbt.Byte(d))
}

// DifficultyLocked locks the difficulty in game
func (l *Level) DifficultyLocked(dl bool) {
	l.setOption("DifficultyLocked", boolToByte(dl))
}

// TicksExisted sets how many ticks have passed in game
func (l *Level) TicksExisted(t int64) {
	l.setOption("Time", nbt.Long(t))
}

// Time-of-day convenience constants
const (
	SunRise  = 0
	Noon     = 6000
	SunSet   = 12000
	MidNight = 18000
	Day      = 24000
)

// Time sets the in world time.
func (l *Level) Time(t int64) {
	l.setOption("DayTime", nbt.Long(t))
}

// GetSpawn sets the given x, y, z coordinates to the current spawn point.
func (l *Level) GetSpawn() (x int32, y int32, z int32) {
	xTag, yTag, zTag := l.levelData.Get("SpawnX"), l.levelData.Get("SpawnY"), l.levelData.Get("SpawnZ")
	x = int32(xTag.Data().(nbt.Int))
	y = int32(yTag.Data().(nbt.Int))
	z = int32(zTag.Data().(nbt.Int))
	return x, y, z
}

// Spawn sets the spawn point to the given coordinates.
func (l *Level) Spawn(x, y, z int32) {
	l.levelData.Set(nbt.NewTag("SpawnX", nbt.Int(x)))
	l.levelData.Set(nbt.NewTag("SpawnY", nbt.Int(y)))
	l.levelData.Set(nbt.NewTag("SpawnZ", nbt.Int(z)))
	l.changed = true
}

//BorderCenter sets the position of the center of the World Border
func (l *Level) BorderCenter(x, z float64) {
	l.levelData.Set(nbt.NewTag("BorderCenterX", nbt.Double(x)))
	l.levelData.Set(nbt.NewTag("BorderCenterZ", nbt.Double(z)))
	l.changed = true
}

//BorderSize sets the width of the border
func (l *Level) BorderSize(w float64) {
	l.setOption("BorderSize", nbt.Double(w))
}

// Raining sets the rain on or off
func (l *Level) Raining(raining bool) {
	l.setOption("raining", boolToByte(raining))
}

// RainTime sets the time until the rain state changes
func (l *Level) RainTime(time int32) {
	l.setOption("rainTime", nbt.Int(time))
}

// Thundering sets the lightning/thunder on or off
func (l *Level) Thundering(thundering bool) {
	l.setOption("thundering", boolToByte(thundering))
}

// ThunderTime sets the tune until the thunder state changes
func (l *Level) ThunderTime(time int32) {
	l.setOption("thunderTime", nbt.Int(time))
}

// CommandBlockOutput enables/disables chat echo for command blocks
func (l *Level) CommandBlockOutput(d bool) {
	l.setGameRule("commandBlockOutput", d)
}

// DayLightCycle enables/disables the day/night cycle
func (l *Level) DayLightCycle(d bool) {
	l.setGameRule("doDaylightCycle", d)
}

// FireTick enables/disables fire updates, such as spreading and extinguishing
func (l *Level) FireTick(d bool) {
	l.setGameRule("doFireTick", d)
}

// MobLoot enables/disable mob loot drops
func (l *Level) MobLoot(d bool) {
	l.setGameRule("doMobLoot", d)
}

// MobSpawning enables/disables mob spawning
func (l *Level) MobSpawning(d bool) {
	l.setGameRule("doMobSpawning", d)
}

// TileDrops enables/disables the dropping of items upon block breakage
func (l *Level) TileDrops(d bool) {
	l.setGameRule("doTileDrops", d)
}

// KeepInventory enables/disables the keeping of a players inventory upon death
func (l *Level) KeepInventory(d bool) {
	l.setGameRule("keepInventory", d)
}

// LogAdminCommands enables/disables the logging of admin commmands to the log
func (l *Level) LogAdminCommands(d bool) {
	l.setGameRule("logAdminCommands", d)
}

// MobGriefing enables/disables the abilty of mobs to destroy blocks
func (l *Level) MobGriefing(d bool) {
	l.setGameRule("mobGriefing", d)
}

// HealthRegeneration enables/disables the regeneration of the players health
// when their hunger is high enough
func (l *Level) HealthRegeneration(d bool) {
	l.setGameRule("naturalRegeneration", d)
}

// CommandFeedback enables/disables the echo for player commands in the chat
func (l *Level) CommandFeedback(d bool) {
	l.setGameRule("sendCommandFeedback", d)
}

// DeathMessages enables/disables the logging of player deaths to the chat
func (l *Level) DeathMessages(d bool) {
	l.setGameRule("showDeathMessages", d)
}

func (l *Level) setOption(name string, data nbt.Data) {
	l.levelData.Set(nbt.NewTag(name, data))
	l.changed = true
}

func (l *Level) setGameRule(name string, data bool) {
	var grc nbt.Compound
	gr := l.levelData.Get("GameRules")
	if gr.TagID() != 0 {
		grc = gr.Data().(nbt.Compound)
	}
	if grc.Type() == 0 {
		l.levelData.Set(nbt.NewTag("GameRules", grc))
	}
	var d nbt.String
	if data {
		d = nbt.String("True")
	} else {
		d = nbt.String("False")
	}
	grc.Set(nbt.NewTag(name, d))
	l.changed = true
}

func boolToByte(b bool) nbt.Byte {
	if b {
		return nbt.Byte(1)
	}
	return nbt.Byte(0)
}
