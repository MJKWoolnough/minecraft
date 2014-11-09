package nbt

import "github.com/MJKWoolnough/bytewrite"

// Tag Types
const (
	Tag_End TagId = iota
	Tag_Byte
	Tag_Short
	Tag_Int
	Tag_Long
	Tag_Float
	Tag_Double
	Tag_ByteArray
	Tag_String
	Tag_List
	Tag_Compound
	Tag_IntArray
)

var tagIdNames = [...]string{
	"End",
	"Byte",
	"Short",
	"Int",
	"Long",
	"Float",
	"Double",
	"Byte Array",
	"String",
	"List",
	"Compound",
	"Int Array",
}

type TagId uint8

func (t TagId) String() string {
	if int(t) < len(tagIdNames) {
		return tagIdNames[t]
	}
	return ""
}

type Config struct {
	endian bytewrite.Endian
	types  map[TagId]func() Data
}

func NewConfig(endian bytewrite.Endian) *Config {
	return &Config{
		endian: endian,
		types:  make(map[TagId]func() Data),
	}
}

func (c Config) newFromTag(id TagId) (Data, error) {
	if nd, ok := c.types[id]; ok {
		return nd(), nil
	}
	var d Data
	switch id {
	case Tag_Byte:
		d = new(Byte)
	case Tag_Short:
		d = new(Short)
	case Tag_Int:
		d = new(Int)
	case Tag_Long:
		d = new(Long)
	case Tag_Float:
		d = new(Float)
	case Tag_Double:
		d = new(Double)
	case Tag_ByteArray:
		d = new(ByteArray)
	case Tag_String:
		d = new(String)
	case Tag_List:
		d = new(List)
	case Tag_Compound:
		d = new(Compound)
	case Tag_IntArray:
		d = new(IntArray)
	default:
		return nil, &UnknownTag{id}
	}
	return d, nil
}

func (c Config) RegisterType(id TagId, nd func() Data) {
	c.types[id] = nd
}

var defaultConfig = &Config{endian: bytewrite.BigEndian}
