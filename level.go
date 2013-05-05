// Copyright (c) 2013 - Michael Woolnough <michael.woolnough@gmail.com>
// 
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met: 
// 
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer. 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution. 
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package minecraft

import (
	"compress/gzip"
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbtparser"
	"os"
	"strconv"
	"time"
)

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
	//UpdateLighting()
	//HighestBlock(int32, int32) uint8
}

type level struct {
	levelData *nbtparser.NBTTagCompound
	regions   map[int32]map[int32]Region
	path      string
}

// func (l *level) HighestBlock(x, z int32) uint8 {
// 	if l == nil {
// 		return 0
// 	}
// 	regionX := x >> 9
// 	regionZ := z >> 9
// 	if l.regions[regionX] == nil || l.regions[regionX][regionZ] == nil {
// 		if !l.loadLevel(regionX, regionZ) {
// 			return 0
// 		}
// 	}
// 	return l.regions[regionX][regionZ].HighestBlock(x, z)
// }
// 
// func (l *level) UpdateLighting() {
// 	regions := make([]struct{ r Region; x, z int32 }, 0)
// 	for i := range l.regions {
// 		for j, reg := range l.regions[i] {
// 			regions = append(regions, struct { r Region; x, z int32 } { reg, i, j })
// 		}
// 	}
// 	for _, region := range regions {
// 		for i := int32(0); i < 32; i++ {
// 			for j := int32(0); j < 32; j++ {
// 				data := region.r.SkyUpdates(i << 4, j << 4)
// 				if false {
// 				for k := uint8(0); k < 16; k++ {
// 					for m := uint8(0); m < 16; m++ {
// 						if data[zx(k, m)] {
// 							data[zx(k, m)] = false
// 							sX := region.x << 9 + i << 4 + int32(k)
// 							sZ := region.z << 9 + j << 4 + int32(m)
// 							h := region.r.HighestBlock(sX, sZ)
// 							g := l.HighestBlock(sX - 1, sZ)
// 							if f := l.HighestBlock(sX + 1, sZ); f < g {
// 								g = f
// 							}
// 							if f := l.HighestBlock(sX, sZ - 1); f < g {
// 								g = f
// 							}
// 							if f := l.HighestBlock(sX, sZ + 1); f < g {
// 								g = f
// 							}
// 							if g > h {
// 								g, h = h, g
// 							}
// 							l.checkSkyLight(sX, sZ, g, h)
// 							l.checkSkyLight(sX - 1, sZ, h, h)
// 							l.checkSkyLight(sX + 1, sZ, h, h)
// 							l.checkSkyLight(sX, sZ - 1, h, h)
// 							l.checkSkyLight(sX, sZ + 1, h, h)
// 						}
// 					}
// 				}
// 				}
// 				l.Compress()
// 			}
// 		}
// 	}
// }
// 
// func (l *level) checkSkyLight(x, z int32, minY, maxY uint8) {
// 	fmt.Println(x, z, minY, maxY)
// 	for i := int32(minY); i <= int32(maxY); i++ {
// 		list := make([]struct{ l uint8; dx, dy, dz int8 }, 0)
// 		if currLighting, newLighting := l.GetSkyLight(x, i, z), l.computeSkyValue(x, i, z); newLighting > currLighting {
// 			list = append(list, struct{ l uint8; dx, dy, dz int8 }{ 0, 0, 0, 0 })
// 		} else if newLighting < currLighting {
// 			list = append(list, struct{ l uint8; dx, dy, dz int8 }{ newLighting, 0, 0, 0 })
// 			for h := 0; h < len(list); h++ {
// 				j := list[h]
// 				if j.l == 0 {
// 					continue
// 				}
// 				pX, pY, pZ := x + int32(j.dx), i + int32(j.dy), z + int32(j.dz)
// 				if l.GetSkyLight(pX, pY, pZ) == j.l {
// 					l.SetSkyLight(pX, pY, pZ, 0)
// 					if uint16(j.dx * j.dx) + uint16(j.dy * j.dy) + uint16(j.dz * j.dz) <= 256 { //within 16 blocks
// 						for k := 0; k < 6; k++ { //each adjacent block
// 							qX, qY, qZ := j.dx, j.dy, j.dz
// 							m := int8(-1)
// 							if k & 1 == 1 {
// 								m = 1
// 							}
// 							switch k >> 1 {
// 								case 0:
// 									qX += m
// 								case 1:
// 									qY += m
// 								case 2:
// 									qZ += m
// 							}
// 							rX, rY, rZ := x + int32(qX), i + int32(qY), z + int32(qZ)
// 							nl := j.l
// 							if nl > 0 {
// 								if o := l.Opacity(rX, rY, rZ); o == 0 {
// 									nl--
// 								} else if o > nl {
// 									nl -= o
// 								} else {
// 									nl = 0
// 								}
// 							}
// 							if l.GetSkyLight(rX, rY, rZ) == nl {
// 								list = append(list, struct{ l uint8; dx, dy, dz int8 }{ nl, qX, qY, qZ })
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 		for h := 0; h < len(list); h++ {
// 			j := list[h]
// 			pX, pY, pZ := x + int32(j.dx), i + int32(j.dy), z + int32(j.dz)
// 			if currLighting, newLighting := l.GetSkyLight(pX, pY, pZ), l.computeSkyValue(pX, pY, pZ); currLighting != newLighting {
// 				l.SetSkyLight(pX, pY, pZ, newLighting)
// 				if currLighting < newLighting && uint16(j.dx * j.dx) + uint16(j.dy * j.dy) + uint16(j.dz * j.dz) <= 256 && len(list) < 32762 { //within 16 blocks
// 					if l.GetSkyLight(pX - 1, pY, pZ) < newLighting {
// 						list = append(list, struct{ l uint8; dx, dy, dz int8 }{ 0, j.dx - 1, j.dy, j.dz })
// 					}
// 					if l.GetSkyLight(pX + 1, pY, pZ) < newLighting {
// 						list = append(list, struct{ l uint8; dx, dy, dz int8 }{ 0, j.dx + 1, j.dy, j.dz })
// 					}
// 					if l.GetSkyLight(pX, pY - 1, pZ) < newLighting {
// 						list = append(list, struct{ l uint8; dx, dy, dz int8 }{ 0, j.dx, j.dy - 1, j.dz })
// 					}
// 					if l.GetSkyLight(pX, pY + 1, pZ) < newLighting {
// 						list = append(list, struct{ l uint8; dx, dy, dz int8 }{ 0, j.dx, j.dy + 1, j.dz })
// 					}
// 					if l.GetSkyLight(pX, pY, pZ - 1) < newLighting {
// 						list = append(list, struct{ l uint8; dx, dy, dz int8 }{ 0, j.dx, j.dy, j.dz - 1 })
// 					}
// 					if l.GetSkyLight(pX, pY, pZ + 1) < newLighting {
// 						list = append(list, struct{ l uint8; dx, dy, dz int8 }{ 0, j.dx, j.dy, j.dz + 1 })
// 					}
// 				}
// 			}
// 		}
// 	}
// }
// 
// func (l *level) computeSkyValue(x, y, z int32) uint8 {
// 	if uint8(y) >= l.HighestBlock(x, z) {
// 		return 15
// 	}
// 	newLighting := l.GetSkyLight(x - 1, y, z)
// 	if n := l.GetSkyLight(x + 1, y, z); n > newLighting {
// 		newLighting = n
// 	}
// 	if n := l.GetSkyLight(x, y - 1, z); n > newLighting {
// 		newLighting = n
// 	}
// 	if n := l.GetSkyLight(x, y + 1, z); n > newLighting {
// 		newLighting = n
// 	}
// 	if n := l.GetSkyLight(x, y, z - 1); n > newLighting {
// 		newLighting = n
// 	}
// 	if n := l.GetSkyLight(x, y, z + 1); n > newLighting {
// 		newLighting = n
// 	}
// 	opacity := l.Opacity(x, y, z)
// 	if opacity == 0 {
// 		opacity = 1
// 	}
// 	return newLighting - opacity
// }

func (l *level) Compress() {
	for i := range l.regions {
		for j := range l.regions[i] {
			l.regions[i][j].Compress()
		}
	}
}

func (l *level) Get(x, y, z int32) Block {
	if l == nil {
		return nil
	}
	regionX := x >> 9
	regionZ := z >> 9
	if l.regions[regionX] == nil || l.regions[regionX][regionZ] == nil {
		if !l.loadLevel(regionX, regionZ) {
			return BlockAir
		}
	}
	return l.regions[regionX][regionZ].Get(x, y, z)
}

func (l *level) Opacity(x, y, z int32) uint8 {
	regionX := x >> 9
	regionZ := z >> 9
	if l.regions[regionX] == nil || l.regions[regionX][regionZ] == nil {
		if !l.loadLevel(regionX, regionZ) {
			return 0
		}
	}
	return l.regions[regionX][regionZ].Opacity(x, y, z)
}

func (l *level) Set(x, y, z int32, block Block) {
	if l == nil || block == nil {
		return
	}
	regionX := x >> 9
	regionZ := z >> 9
	if l.regions[regionX] == nil {
		l.regions[regionX] = make(map[int32]Region)
	}
	if l.regions[regionX][regionZ] == nil {
		if !l.loadLevel(regionX, regionZ) {
			l.regions[regionX][regionZ], _ = NewRegion()
		}
	}
	l.regions[regionX][regionZ].Set(x, y, z, block)
}

func (l *level) GetSkyLight(x, y, z int32) uint8 {
	regionX := x >> 9
	regionZ := z >> 9
	if l.regions[regionX] == nil || l.regions[regionX][regionZ] == nil {
		if !l.loadLevel(regionX, regionZ) {
			return 15
		}
	}
	return l.regions[regionX][regionZ].GetSkyLight(x, y, z)
}

func (l *level) SetSkyLight(x, y, z int32, skylight uint8) {
	regionX := x >> 9
	regionZ := z >> 9
	if l.regions[regionX] == nil {
		l.regions[regionX] = make(map[int32]Region)
	}
	if l.regions[regionX][regionZ] == nil {
		if !l.loadLevel(regionX, regionZ) {
			return
		}
	}
	l.regions[regionX][regionZ].SetSkyLight(x, y, z, skylight)
}

func (l *level) loadLevel(x, z int32) bool {
	if l.path == "" {
		return false
	}
	if file, err := os.Open(l.path + "/region/r." + strconv.Itoa(int(x)) + "." + strconv.Itoa(int(z)) + ".mca"); err != nil {
		return false
	} else {
		defer file.Close()
		if r, err := LoadRegion(file); err != nil {
			return false
		} else {
			if l.regions[x] == nil {
				l.regions[x] = make(map[int32]Region)
			}
			l.regions[x][z] = r
		}
	}
	return true
}

func (l *level) GetName() string {
	return l.levelData.GetTag("Data").TagCompound().GetTag("LevelName").TagString().Get()
}

func (l *level) SetName(name string) {
	l.levelData.GetTag("Data").TagCompound().GetTag("LevelName").TagString().Set(name)
}

func (l *level) GetSpawn() (int32, int32, int32) {
	x := l.levelData.GetTag("Data").TagCompound().GetTag("SpawnX").TagInt().Get()
	y := l.levelData.GetTag("Data").TagCompound().GetTag("SpawnY").TagInt().Get()
	z := l.levelData.GetTag("Data").TagCompound().GetTag("SpawnZ").TagInt().Get()
	return x, y, z
}

func (l *level) SetSpawn(x, y, z int32) {
	l.levelData.GetTag("Data").TagCompound().GetTag("SpawnX").TagInt().Set(x)
	l.levelData.GetTag("Data").TagCompound().GetTag("SpawnY").TagInt().Set(y)
	l.levelData.GetTag("Data").TagCompound().GetTag("SpawnZ").TagInt().Set(z)
}

func (l *level) SaveLevelData(path string) error {
	if err := os.MkdirAll(path+"/", os.FileMode(7<<6|5<<3|5|os.ModeDir)); err != nil {
		return err
	}
	l.levelData.GetTag("Data").TagCompound().GetTag("LastPlayed").TagLong().Set(time.Now().Unix() * 1000)
	if file, err := os.Create(path + "/level.dat"); err != nil {
		return err
	} else {
		defer file.Close()
		zFile := gzip.NewWriter(file)
		defer zFile.Close()
		_, err := l.levelData.WriteTo(zFile)
		return err
	}
	return fmt.Errorf("Minecraft - SaveLevelData: Should never reach here!")
}

func (l *level) ExportOpenRegions(path string) error {
	// 	l.UpdateLighting()
	if err := os.MkdirAll(path+"/region/", os.FileMode(7<<6|5<<3|5|os.ModeDir)); err != nil {
		return err
	}
	l.path = path
	for i := range l.regions {
		for j := range l.regions[i] {
			if l.regions[i][j].HasChanged() {
				if file, err := os.Create(path + "/region/r." + strconv.Itoa(int(i)) + "." + strconv.Itoa(int(j)) + ".mca"); err != nil {
					return err
				} else {
					if err := l.regions[i][j].Export(file); err != nil {
						file.Close()
						return err
					}
					file.Close()
				}
			}
			delete(l.regions[i], j)
		}
		delete(l.regions, i)
	}
	return nil
}

func (l *level) CloseOpenRegions() {
	for i := range l.regions {
		for j := range l.regions[i] {
			delete(l.regions[i], j)
		}
		delete(l.regions, i)
	}
}

func (l *level) String() string {
	return "Minecraft Level - " + l.GetName()
}

func LoadLevel(path string) (Level, error) {
	levelD := new(level)
	if file, err := os.Open(path + "/level.dat"); err != nil {
		return nil, err
	} else {
		defer file.Close()
		if zFile, err := gzip.NewReader(file); err != nil {
			return nil, err
		} else {
			defer zFile.Close()
			if nbtFile, _, err := nbtparser.ParseFile(zFile); err != nil {
				return nil, err
			} else {
				if data := nbtFile.GetTag("Data"); data == nil {
					return nil, fmt.Errorf("Minecraft - Level: Missing 'Data' tag")
				} else if dataTag := data.TagCompound(); dataTag == nil {
					return nil, fmt.Errorf("Minecraft - Level: Data tag of wrong type")
				} else {
					if name := dataTag.GetTag("LevelName"); name == nil {
						return nil, fmt.Errorf("Minecraft - Level: Missing 'LevelName' tag")
					} else if name.TagString() == nil {
						return nil, fmt.Errorf("Minecraft - Level: LevelName tag of wrong type")
					}
					if name := dataTag.GetTag("SpawnX"); name == nil {
						return nil, fmt.Errorf("Minecraft - Level: Missing 'SpawnX' tag")
					} else if name.TagInt() == nil {
						return nil, fmt.Errorf("Minecraft - Level: SpawnX tag of wrong type")
					}
					if name := dataTag.GetTag("SpawnY"); name == nil {
						return nil, fmt.Errorf("Minecraft - Level: Missing 'SpawnY' tag")
					} else if name.TagInt() == nil {
						return nil, fmt.Errorf("Minecraft - Level: SpawnY tag of wrong type")
					}
					if name := dataTag.GetTag("SpawnZ"); name == nil {
						return nil, fmt.Errorf("Minecraft - Level: Missing 'SpawnZ' tag")
					} else if name.TagInt() == nil {
						return nil, fmt.Errorf("Minecraft - Level: SpawnZ tag of wrong type")
					}
				}
				levelD.levelData = nbtFile
				levelD.regions = make(map[int32]map[int32]Region)
				levelD.path = path
			}
		}
	}
	return levelD, nil
}

func NewLevel(name string) Level {
	levelD := new(level)
	levelD.regions = make(map[int32]map[int32]Region)
	levelD.levelData = nbtparser.NewTagCompound("", []nbtparser.NBTTag{nbtparser.NewTagCompound("Data", []nbtparser.NBTTag{
		nbtparser.NewTagByte("thundering", 0),
		nbtparser.NewTagLong("LastPlayed", time.Now().Unix()*1000),
		//nbtparser.NewTagCompound("Player", []nbtparser.NBTTag {
		//	nbtparser.NewTagList("Motion", nbtparser.NBTTag_Double, []interface{}{
		//		float64(0),
		//		float64(0),
		//		float64(0),
		//	}),
		//	nbtparser.NewTagFloat("foodExhaustionLevel", 0),
		//	nbtparser.NewTagInt("foodTickTimer", 0),
		//	nbtparser.NewTagInt("PersistentId", 527182454),
		//	nbtparser.NewTagInt("XpLevel", 0),
		//	nbtparser.NewTagShort("Health", 20),
		//	nbtparser.NewTagList("Inventory", nbtparser.NBTTag_Compound, []interface{}{}),
		//	nbtparser.NewTagShort("AttackTime", 0),
		//	nbtparser.NewTagByte("Sleeping", 0),
		//	nbtparser.NewTagShort("Fire", 0),
		//	nbtparser.NewTagInt("foodLevel", 20),
		//	nbtparser.NewTagInt("Score", 0),
		//	nbtparser.NewTagShort("DeathTime", 0),
		//	nbtparser.NewTagFloat("XpP", 0),
		//	nbtparser.NewTagByte("SleepTimer", 0),
		//	nbtparser.NewTagShort("HurtTime", 0),
		//	nbtparser.NewTagByte("OnGround", 1),
		//	nbtparser.NewTagInt("Dimension", 0),
		//	nbtparser.NewTagShort("Air", 0),
		//	nbtparser.NewTagList("Pos", nbtparser.NBTTag_Double, []interface{}{
		//		float64(20),
		//		float64(20),
		//		float64(20),
		//	}),
		//	nbtparser.NewTagFloat("foodSaturationLevel", 20),
		//	nbtparser.NewTagCompound("abilities", []nbtparser.NBTTag {
		//		nbtparser.NewTagFloat("walkSpeed", 0.1),
		//		nbtparser.NewTagFloat("flySpeed", 0.05),
		//		nbtparser.NewTagByte("mayfly", 1),
		//		nbtparser.NewTagByte("flying", 1),
		//		nbtparser.NewTagByte("invulnerable", 1),
		//		nbtparser.NewTagByte("maybuild", 1),
		//		nbtparser.NewTagByte("instabuild", 1),
		//	}),
		//	nbtparser.NewTagFloat("FallDistance", 0),
		//	nbtparser.NewTagInt("XpTotal", 0),
		//	nbtparser.NewTagList("Rotation", nbtparser.NBTTag_Float, []interface{}{
		//		float32(0),
		//		float32(0),
		//	}),
		//	nbtparser.NewTagByte("CanPickUpLoot", 1),
		//	nbtparser.NewTagInt("playerGameType", 1),
		//	nbtparser.NewTagList("EnderItems", nbtparser.NBTTag_Byte, []interface{}),
		//}),
		nbtparser.NewTagLong("RandomSeed", -6044196962465721491),
		nbtparser.NewTagInt("GameType", 1),
		nbtparser.NewTagByte("MapFeatures", 0),
		nbtparser.NewTagInt("version", 19133),
		nbtparser.NewTagLong("Time", 0),
		nbtparser.NewTagByte("raining", 0),
		nbtparser.NewTagInt("thunderTime", 0),
		nbtparser.NewTagInt("SpawnX", 20),
		nbtparser.NewTagByte("hardcore", 0),
		nbtparser.NewTagInt("SpawnY", 20),
		nbtparser.NewTagInt("SpawnZ", 20),
		nbtparser.NewTagString("LevelName", name),
		nbtparser.NewTagString("generatorName", "flat"),
		nbtparser.NewTagLong("SizeOnDisk", 0),
		nbtparser.NewTagInt("rainTime", 0),
		nbtparser.NewTagInt("generatorVersion", 0),
		nbtparser.NewTagString("generatorOptions", "0"),
		//nbtparser.NewTagByte("allowCommands", 1),
		//nbtparser.NewTagLong("DayTime", 0),
		//nbtparser.NewTagCompound("GameRules", []nbtparser.NBTTag {
		//	nbtparser.NewTagString("doFireTick", "true"),
		//	nbtparser.NewTagString("doMobLoot", "true"),
		//	nbtparser.NewTagString("doModSpawning", "true"),
		//	nbtparser.NewTagString("doTileDrops", "true"),
		//	nbtparser.NewTagString("keepInventory", "false"),
		//	nbtparser.NewTagString("mobGriefing", "true"),
		//}),
		//nbtparser.NewTagByte("initialized", 1),
	})})
	return levelD
}

func regionCoords(x, y, z int32) (uint16, uint16, uint16) {
	x1 := x % 512
	z1 := x % 512
	if x1 < 0 {
		x1 += 512
	}
	if z1 < 0 {
		z1 += 512
	}
	return uint16(x1), uint16(y), uint16(z1)
}
