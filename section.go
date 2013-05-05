package minecraft

import (
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbtparser"
)

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

type section struct {
	section       *nbtparser.NBTTagCompound
	blocks        *nbtparser.NBTTagByteArray
	addTag        *nbtparser.NBTTagByteArray
	dataTag       *nbtparser.NBTTagByteArray
	blockLightTag *nbtparser.NBTTagByteArray
	skyLightTag   *nbtparser.NBTTagByteArray
}

func (s *section) Get(x, y, z uint8) Block {
	if s == nil {
		return BlockAir
	}
	return NewBlock(uint8(s.blocks.Get(int32(yzx(x, y, z)))), getNibble(s.addTag, x, y, z), getNibble(s.dataTag, x, y, z))
}

func (s *section) Set(x, y, z uint8, block Block) {
	if block == nil || s == nil {
		return
	}
	s.blocks.Set(int32(yzx(x, y, z)), block.BlockId())
	setNibble(s.addTag, x, y, z, block.Add())
	setNibble(s.dataTag, x, y, z, block.Data())
	if l, ok := lightBlocks[block.BlockId()]; ok {
		setNibble(s.blockLightTag, x, y, z, l)
	} else {
		setNibble(s.blockLightTag, x, y, z, 0)
	}
}

func (s *section) GetSkyLight(x, y, z uint8) uint8 {
	if s == nil {
		return 0
	}
	return getNibble(s.skyLightTag, x, y, z)
}

func (s *section) SetSkyLight(x, y, z, skyLight uint8) {
	if s != nil {
		setNibble(s.skyLightTag, x, y, z, skyLight)
	}
}

func (s *section) Opacity(x, y, z uint8) uint8 {
	if s == nil {
		return 0
	}
	blockId := uint16(s.blocks.Get(int32(yzx(x, y, z)))) | (uint16(getNibble(s.addTag, x, y, z)) << 8)
	if blockId == 8 || blockId == 9 {
		return 3
	}
	for i := 0; i < len(transparentBlocks); i++ {
		if transparentBlocks[i] == blockId {
			return 0
		}
	}
	return 16
}

func (s *section) GetY() uint8 {
	return uint8(s.section.GetTag("Y").TagByte().Get())
}

func (s *section) IsEmpty() bool {
	if s == nil {
		return true
	}
	for i := int32(0); i < 4096; i++ {
		if s.blocks.Get(i) != 0 {
			return false
		}
	}
	return true
}

func (s *section) Data() *nbtparser.NBTTagCompound {
	return s.section
}

func LoadSection(sectionData *nbtparser.NBTTagCompound) (Section, error) {
	if sectionData == nil {
		return nil, fmt.Errorf("Minecraft - Section: nil received")
	}
	if yTag := sectionData.GetTag("Y"); yTag == nil {
		return nil, fmt.Errorf("Minecraft - Section: Missing 'Y' tag")
	} else if yTag.TagByte() == nil {
		return nil, fmt.Errorf("Minecraft - Section: Y tag of wrong type")
	}
	newSection := new(section)
	newSection.section = sectionData
	if blocks := sectionData.GetTag("Blocks"); blocks == nil {
		return nil, fmt.Errorf("Minecraft - Section: Missing 'Blocks' tag")
	} else if blocksArray := blocks.TagByteArray(); blocksArray == nil {
		return nil, fmt.Errorf("Minecraft - Section: Blocks tag of wrong type")
	} else if blocksArray.Length() != 4096 {
		return nil, fmt.Errorf("Minecraft - Section: Incorrect number of blocks in array")
	} else {
		newSection.blocks = blocksArray
	}
	if adds := sectionData.GetTag("Add"); adds == nil {
		addArray := nbtparser.NewTagByteArray("Add", make([]byte, 2048))
		sectionData.Append(addArray)
		newSection.addTag = addArray
	} else if addArray := adds.TagByteArray(); addArray == nil {
		return nil, fmt.Errorf("Minecraft - Section: Add tag of wrong type")
	} else if addArray.Length() != 2048 {
		return nil, fmt.Errorf("Minecraft - Section: Incorrect number of add bytes in array")
	} else {
		newSection.addTag = addArray
	}
	if data := sectionData.GetTag("Data"); data == nil {
		return nil, fmt.Errorf("Minecraft - Section: Missing 'Data' tag")
	} else if dataArray := data.TagByteArray(); dataArray == nil {
		return nil, fmt.Errorf("Minecraft - Section: Data tag of wrong type")
	} else if dataArray.Length() != 2048 {
		return nil, fmt.Errorf("Minecraft - Section: Incorrect number of data bytes in array")
	} else {
		newSection.dataTag = dataArray
	}
	if blockLight := sectionData.GetTag("BlockLight"); blockLight == nil {
		return nil, fmt.Errorf("Minecraft - Section: Missing 'BlockLight' tag")
	} else if blockLightArray := blockLight.TagByteArray(); blockLightArray == nil {
		return nil, fmt.Errorf("Minecraft - Section: BlockLight tag of wrong type")
	} else if blockLightArray.Length() != 2048 {
		return nil, fmt.Errorf("Minecraft - Section: Incorrect number of BlockLight bytes in array")
	} else {
		newSection.blockLightTag = blockLightArray
	}
	if skyLight := sectionData.GetTag("SkyLight"); skyLight == nil {
		return nil, fmt.Errorf("Minecraft - Section: Missing 'skyLight' tag")
	} else if skyLightArray := skyLight.TagByteArray(); skyLightArray == nil {
		return nil, fmt.Errorf("Minecraft - Section: SkyLight tag of wrong type")
	} else if skyLightArray.Length() != 2048 {
		return nil, fmt.Errorf("Minecraft - Section: Incorrect number of SkyLight bytes in array")
	} else {
		newSection.skyLightTag = skyLightArray
	}
	return newSection, nil
}

func NewSection(y byte) (Section, error) {
	if y > 15 {
		return nil, fmt.Errorf("Minecraft - Section: Y value is too high for current limit")
	}
	skyLight := make([]byte, 2048)
	for i := 0; i < 2048; i++ {
		skyLight[i] = 0xFF
	}
	return LoadSection(nbtparser.NewTagCompound("", []nbtparser.NBTTag{
		nbtparser.NewTagByteArray("Data", make([]byte, 2048)),
		nbtparser.NewTagByteArray("SkyLight", skyLight),
		nbtparser.NewTagByteArray("BlockLight", make([]byte, 2048)),
		nbtparser.NewTagByte("Y", int8(y)),
		nbtparser.NewTagByteArray("Blocks", make([]byte, 4096)),
		nbtparser.NewTagByteArray("Add", make([]byte, 2048)),
	}))
}

func yzx(x, y, z uint8) uint16 {
	return ((uint16(y) & 15) << 8) | ((uint16(z) & 15) << 4) | (uint16(x) & 15)
}

func getNibble(arr *nbtparser.NBTTagByteArray, x, y, z uint8) uint8 {
	coord := yzx(x, y, z)
	if coord&1 == 0 {
		return uint8(arr.Get(int32(coord>>1)) & 15)
	}
	return uint8(arr.Get(int32(coord>>1)) >> 4)
}

func setNibble(arr *nbtparser.NBTTagByteArray, x, y, z uint8, data uint8) {
	coord := yzx(x, y, z)
	oldData := arr.Get(int32(coord >> 1))
	if coord&1 == 0 {
		arr.Set(int32(coord>>1), (oldData&240)|(data&15))
	} else {
		arr.Set(int32(coord>>1), (data<<4)|(oldData&15))
	}
}
