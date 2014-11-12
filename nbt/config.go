package nbt

import "github.com/MJKWoolnough/bytewrite"

// Tag Types
const (
	TagEnd TagID = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
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

// TagId represents the type of nbt tag
type TagID uint8

func (t TagID) String() string {
	if int(t) < len(tagIdNames) {
		return tagIdNames[t]
	}
	return ""
}

// Config allows for specific configuration of endianness and types for the NBT
// parser
type Config struct {
	endian bytewrite.Endian
	types  map[TagID]func() Data
}

// NewConfig returns a new Config struct with the endianness set accordingly.
func NewConfig(endian bytewrite.Endian) *Config {
	return &Config{
		endian: endian,
		types:  make(map[TagID]func() Data),
	}
}

func (c Config) newFromTag(id TagID) (Data, error) {
	if nd, ok := c.types[id]; ok {
		return nd(), nil
	}
	var d Data
	switch id {
	case TagByte:
		d = new(Byte)
	case TagShort:
		d = new(Short)
	case TagInt:
		d = new(Int)
	case TagLong:
		d = new(Long)
	case TagFloat:
		d = new(Float)
	case TagDouble:
		d = new(Double)
	case TagByteArray:
		d = new(ByteArray)
	case TagString:
		d = new(String)
	case TagList:
		d = new(List)
	case TagCompound:
		d = new(Compound)
	case TagIntArray:
		d = new(IntArray)
	default:
		return nil, &UnknownTag{id}
	}
	return d, nil
}

// RegisterType allows the creating of new NBT tags or the overriding of
// existing ones.
//
// TagId is the new or existing id and the func() Data is a small function
// which creates a newly initialised data tag.
func (c Config) RegisterType(id TagID, nd func() Data) {
	c.types[id] = nd
}

var defaultConfig = &Config{endian: bytewrite.BigEndian}
