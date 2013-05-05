# nbtparser
--
    import "github.com/MJKWoolnough/minecraft/nbtparser"

Package nbtparser implements a full Named Binary Tag parser, based on the specs at
http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt

## Usage

```go
const (
	NBTTag_End = uint8(iota)
	NBTTag_Byte
	NBTTag_Short
	NBTTag_Int
	NBTTag_Long
	NBTTag_Float
	NBTTag_Double
	NBTTag_ByteArray
	NBTTag_String
	NBTTag_List
	NBTTag_Compound
	NBTTag_IntArray
)
```
Tag Types

#### func  TagName

```go
func TagName(id uint8) string
```
TagName converts a tag id into its canonical name.

#### type NBTTag

```go
type NBTTag interface {
	equaler.Equaler // Allows the tags to be compared to other tags.
	fmt.Stringer
	io.ReaderFrom
	io.WriterTo
	Type() uint8
	TypeString() string
	Name() string
	Copy() NBTTag //Produces a deep-copy of the tag
	TagEnd() *NBTTagEnd
	TagByte() *NBTTagByte
	TagShort() *NBTTagShort
	TagInt() *NBTTagInt
	TagLong() *NBTTagLong
	TagFloat() *NBTTagFloat
	TagDouble() *NBTTagDouble
	TagByteArray() *NBTTagByteArray
	TagString() *NBTTagString
	TagList() *NBTTagList
	TagCompound() *NBTTagCompound
	TagIntArray() *NBTTagIntArray
}
```

NBTTag is the main interface for all of the NBT types. All tags implement the
io.ReaderFrom and io.WriterTo interfaces.

#### type NBTTagByte

```go
type NBTTagByte struct {
}
```


#### func  NewTagByte

```go
func NewTagByte(name string, data int8) *NBTTagByte
```

#### func (*NBTTagByte) Copy

```go
func (n *NBTTagByte) Copy() NBTTag
```

#### func (*NBTTagByte) Equal

```go
func (n *NBTTagByte) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagByte) Get

```go
func (n *NBTTagByte) Get() int8
```

#### func (NBTTagByte) Name

```go
func (n NBTTagByte) Name() string
```

#### func (*NBTTagByte) ReadFrom

```go
func (n *NBTTagByte) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagByte) Set

```go
func (n *NBTTagByte) Set(data int8)
```

#### func (NBTTagByte) SetName

```go
func (n NBTTagByte) SetName(name string)
```

#### func (*NBTTagByte) String

```go
func (n *NBTTagByte) String() string
```

#### func (*NBTTagByte) TagByte

```go
func (n *NBTTagByte) TagByte() *NBTTagByte
```

#### func (NBTTagByte) TagByteArray

```go
func (n NBTTagByte) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagCompound

```go
func (n NBTTagByte) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagDouble

```go
func (n NBTTagByte) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagEnd

```go
func (n NBTTagByte) TagEnd() *NBTTagEnd
```

#### func (NBTTagByte) TagFloat

```go
func (n NBTTagByte) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagInt

```go
func (n NBTTagByte) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagIntArray

```go
func (n NBTTagByte) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagList

```go
func (n NBTTagByte) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagLong

```go
func (n NBTTagByte) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagShort

```go
func (n NBTTagByte) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) TagString

```go
func (n NBTTagByte) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByte) Type

```go
func (n NBTTagByte) Type() uint8
```

#### func (NBTTagByte) TypeString

```go
func (n NBTTagByte) TypeString() string
```

#### func (*NBTTagByte) WriteTo

```go
func (n *NBTTagByte) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagByteArray

```go
type NBTTagByteArray struct {
}
```


#### func  NewTagByteArray

```go
func NewTagByteArray(name string, data []byte) *NBTTagByteArray
```

#### func (*NBTTagByteArray) Append

```go
func (n *NBTTagByteArray) Append(data byte)
```

#### func (*NBTTagByteArray) Copy

```go
func (n *NBTTagByteArray) Copy() NBTTag
```

#### func (*NBTTagByteArray) Equal

```go
func (n *NBTTagByteArray) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagByteArray) Get

```go
func (n *NBTTagByteArray) Get(index int32) byte
```

#### func (*NBTTagByteArray) GetArray

```go
func (n *NBTTagByteArray) GetArray() []byte
```

#### func (*NBTTagByteArray) Length

```go
func (n *NBTTagByteArray) Length() int32
```

#### func (NBTTagByteArray) Name

```go
func (n NBTTagByteArray) Name() string
```

#### func (*NBTTagByteArray) ReadFrom

```go
func (n *NBTTagByteArray) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagByteArray) Set

```go
func (n *NBTTagByteArray) Set(index int32, data byte)
```

#### func (NBTTagByteArray) SetName

```go
func (n NBTTagByteArray) SetName(name string)
```

#### func (*NBTTagByteArray) String

```go
func (n *NBTTagByteArray) String() string
```

#### func (NBTTagByteArray) TagByte

```go
func (n NBTTagByteArray) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagByteArray) TagByteArray

```go
func (n *NBTTagByteArray) TagByteArray() *NBTTagByteArray
```

#### func (NBTTagByteArray) TagCompound

```go
func (n NBTTagByteArray) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) TagDouble

```go
func (n NBTTagByteArray) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) TagEnd

```go
func (n NBTTagByteArray) TagEnd() *NBTTagEnd
```

#### func (NBTTagByteArray) TagFloat

```go
func (n NBTTagByteArray) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) TagInt

```go
func (n NBTTagByteArray) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) TagIntArray

```go
func (n NBTTagByteArray) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) TagList

```go
func (n NBTTagByteArray) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) TagLong

```go
func (n NBTTagByteArray) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) TagShort

```go
func (n NBTTagByteArray) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) TagString

```go
func (n NBTTagByteArray) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagByteArray) Type

```go
func (n NBTTagByteArray) Type() uint8
```

#### func (NBTTagByteArray) TypeString

```go
func (n NBTTagByteArray) TypeString() string
```

#### func (*NBTTagByteArray) WriteTo

```go
func (n *NBTTagByteArray) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagCompound

```go
type NBTTagCompound struct {
}
```


#### func  NewTagCompound

```go
func NewTagCompound(name string, data []NBTTag) *NBTTagCompound
```

#### func  ParseFile

```go
func ParseFile(file io.Reader) (*NBTTagCompound, int64, error)
```
ParseFile constructs an entire NBT Tree from a reader.

#### func (*NBTTagCompound) Append

```go
func (n *NBTTagCompound) Append(data NBTTag)
```

#### func (*NBTTagCompound) Copy

```go
func (n *NBTTagCompound) Copy() NBTTag
```

#### func (*NBTTagCompound) Equal

```go
func (n *NBTTagCompound) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagCompound) Get

```go
func (n *NBTTagCompound) Get(index int32) NBTTag
```

#### func (*NBTTagCompound) GetArray

```go
func (n *NBTTagCompound) GetArray() []NBTTag
```

#### func (*NBTTagCompound) GetTag

```go
func (n *NBTTagCompound) GetTag(tagName string) NBTTag
```

#### func (*NBTTagCompound) Length

```go
func (n *NBTTagCompound) Length() int32
```

#### func (NBTTagCompound) Name

```go
func (n NBTTagCompound) Name() string
```

#### func (*NBTTagCompound) ReadFrom

```go
func (n *NBTTagCompound) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagCompound) RemoveTag

```go
func (n *NBTTagCompound) RemoveTag(tagName string) NBTTag
```

#### func (*NBTTagCompound) Set

```go
func (n *NBTTagCompound) Set(index int32, data NBTTag)
```

#### func (*NBTTagCompound) SetArray

```go
func (n *NBTTagCompound) SetArray(data []NBTTag)
```

#### func (NBTTagCompound) SetName

```go
func (n NBTTagCompound) SetName(name string)
```

#### func (*NBTTagCompound) String

```go
func (n *NBTTagCompound) String() string
```

#### func (NBTTagCompound) TagByte

```go
func (n NBTTagCompound) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) TagByteArray

```go
func (n NBTTagCompound) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagCompound) TagCompound

```go
func (n *NBTTagCompound) TagCompound() *NBTTagCompound
```

#### func (NBTTagCompound) TagDouble

```go
func (n NBTTagCompound) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) TagEnd

```go
func (n NBTTagCompound) TagEnd() *NBTTagEnd
```

#### func (NBTTagCompound) TagFloat

```go
func (n NBTTagCompound) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) TagInt

```go
func (n NBTTagCompound) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) TagIntArray

```go
func (n NBTTagCompound) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) TagList

```go
func (n NBTTagCompound) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) TagLong

```go
func (n NBTTagCompound) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) TagShort

```go
func (n NBTTagCompound) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) TagString

```go
func (n NBTTagCompound) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagCompound) Type

```go
func (n NBTTagCompound) Type() uint8
```

#### func (NBTTagCompound) TypeString

```go
func (n NBTTagCompound) TypeString() string
```

#### func (*NBTTagCompound) WriteTo

```go
func (n *NBTTagCompound) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagDouble

```go
type NBTTagDouble struct {
}
```


#### func  NewTagDouble

```go
func NewTagDouble(name string, data float64) *NBTTagDouble
```

#### func (*NBTTagDouble) Copy

```go
func (n *NBTTagDouble) Copy() NBTTag
```

#### func (*NBTTagDouble) Equal

```go
func (n *NBTTagDouble) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagDouble) Get

```go
func (n *NBTTagDouble) Get() float64
```

#### func (NBTTagDouble) Name

```go
func (n NBTTagDouble) Name() string
```

#### func (*NBTTagDouble) ReadFrom

```go
func (n *NBTTagDouble) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagDouble) Set

```go
func (n *NBTTagDouble) Set(data float64)
```

#### func (NBTTagDouble) SetName

```go
func (n NBTTagDouble) SetName(name string)
```

#### func (*NBTTagDouble) String

```go
func (n *NBTTagDouble) String() string
```

#### func (NBTTagDouble) TagByte

```go
func (n NBTTagDouble) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) TagByteArray

```go
func (n NBTTagDouble) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) TagCompound

```go
func (n NBTTagDouble) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagDouble) TagDouble

```go
func (n *NBTTagDouble) TagDouble() *NBTTagDouble
```

#### func (NBTTagDouble) TagEnd

```go
func (n NBTTagDouble) TagEnd() *NBTTagEnd
```

#### func (NBTTagDouble) TagFloat

```go
func (n NBTTagDouble) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) TagInt

```go
func (n NBTTagDouble) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) TagIntArray

```go
func (n NBTTagDouble) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) TagList

```go
func (n NBTTagDouble) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) TagLong

```go
func (n NBTTagDouble) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) TagShort

```go
func (n NBTTagDouble) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) TagString

```go
func (n NBTTagDouble) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagDouble) Type

```go
func (n NBTTagDouble) Type() uint8
```

#### func (NBTTagDouble) TypeString

```go
func (n NBTTagDouble) TypeString() string
```

#### func (*NBTTagDouble) WriteTo

```go
func (n *NBTTagDouble) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagEnd

```go
type NBTTagEnd struct {
}
```

NBTTagEnd marks the end of a compound tag

#### func  NewTagEnd

```go
func NewTagEnd() *NBTTagEnd
```

#### func (*NBTTagEnd) Copy

```go
func (n *NBTTagEnd) Copy() NBTTag
```

#### func (*NBTTagEnd) Equal

```go
func (n *NBTTagEnd) Equal(e equaler.Equaler) bool
```

#### func (NBTTagEnd) Name

```go
func (n NBTTagEnd) Name() string
```

#### func (*NBTTagEnd) ReadFrom

```go
func (n *NBTTagEnd) ReadFrom(file io.Reader) (int64, error)
```

#### func (NBTTagEnd) SetName

```go
func (n NBTTagEnd) SetName(name string)
```

#### func (*NBTTagEnd) String

```go
func (n *NBTTagEnd) String() string
```

#### func (NBTTagEnd) TagByte

```go
func (n NBTTagEnd) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagByteArray

```go
func (n NBTTagEnd) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagCompound

```go
func (n NBTTagEnd) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagDouble

```go
func (n NBTTagEnd) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagEnd) TagEnd

```go
func (n *NBTTagEnd) TagEnd() *NBTTagEnd
```

#### func (NBTTagEnd) TagFloat

```go
func (n NBTTagEnd) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagInt

```go
func (n NBTTagEnd) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagIntArray

```go
func (n NBTTagEnd) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagList

```go
func (n NBTTagEnd) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagLong

```go
func (n NBTTagEnd) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagShort

```go
func (n NBTTagEnd) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) TagString

```go
func (n NBTTagEnd) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagEnd) Type

```go
func (n NBTTagEnd) Type() uint8
```

#### func (NBTTagEnd) TypeString

```go
func (n NBTTagEnd) TypeString() string
```

#### func (*NBTTagEnd) WriteTo

```go
func (n *NBTTagEnd) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagFloat

```go
type NBTTagFloat struct {
}
```


#### func  NewTagFloat

```go
func NewTagFloat(name string, data float32) *NBTTagFloat
```

#### func (*NBTTagFloat) Copy

```go
func (n *NBTTagFloat) Copy() NBTTag
```

#### func (*NBTTagFloat) Equal

```go
func (n *NBTTagFloat) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagFloat) Get

```go
func (n *NBTTagFloat) Get() float32
```

#### func (NBTTagFloat) Name

```go
func (n NBTTagFloat) Name() string
```

#### func (*NBTTagFloat) ReadFrom

```go
func (n *NBTTagFloat) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagFloat) Set

```go
func (n *NBTTagFloat) Set(data float32)
```

#### func (NBTTagFloat) SetName

```go
func (n NBTTagFloat) SetName(name string)
```

#### func (*NBTTagFloat) String

```go
func (n *NBTTagFloat) String() string
```

#### func (NBTTagFloat) TagByte

```go
func (n NBTTagFloat) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagByteArray

```go
func (n NBTTagFloat) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagCompound

```go
func (n NBTTagFloat) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagDouble

```go
func (n NBTTagFloat) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagEnd

```go
func (n NBTTagFloat) TagEnd() *NBTTagEnd
```

#### func (*NBTTagFloat) TagFloat

```go
func (n *NBTTagFloat) TagFloat() *NBTTagFloat
```

#### func (NBTTagFloat) TagInt

```go
func (n NBTTagFloat) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagIntArray

```go
func (n NBTTagFloat) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagList

```go
func (n NBTTagFloat) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagLong

```go
func (n NBTTagFloat) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagShort

```go
func (n NBTTagFloat) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) TagString

```go
func (n NBTTagFloat) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagFloat) Type

```go
func (n NBTTagFloat) Type() uint8
```

#### func (NBTTagFloat) TypeString

```go
func (n NBTTagFloat) TypeString() string
```

#### func (*NBTTagFloat) WriteTo

```go
func (n *NBTTagFloat) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagInt

```go
type NBTTagInt struct {
}
```


#### func  NewTagInt

```go
func NewTagInt(name string, data int32) *NBTTagInt
```

#### func (*NBTTagInt) Copy

```go
func (n *NBTTagInt) Copy() NBTTag
```

#### func (*NBTTagInt) Equal

```go
func (n *NBTTagInt) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagInt) Get

```go
func (n *NBTTagInt) Get() int32
```

#### func (NBTTagInt) Name

```go
func (n NBTTagInt) Name() string
```

#### func (*NBTTagInt) ReadFrom

```go
func (n *NBTTagInt) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagInt) Set

```go
func (n *NBTTagInt) Set(data int32)
```

#### func (NBTTagInt) SetName

```go
func (n NBTTagInt) SetName(name string)
```

#### func (*NBTTagInt) String

```go
func (n *NBTTagInt) String() string
```

#### func (NBTTagInt) TagByte

```go
func (n NBTTagInt) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) TagByteArray

```go
func (n NBTTagInt) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) TagCompound

```go
func (n NBTTagInt) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) TagDouble

```go
func (n NBTTagInt) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) TagEnd

```go
func (n NBTTagInt) TagEnd() *NBTTagEnd
```

#### func (NBTTagInt) TagFloat

```go
func (n NBTTagInt) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagInt) TagInt

```go
func (n *NBTTagInt) TagInt() *NBTTagInt
```

#### func (NBTTagInt) TagIntArray

```go
func (n NBTTagInt) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) TagList

```go
func (n NBTTagInt) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) TagLong

```go
func (n NBTTagInt) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) TagShort

```go
func (n NBTTagInt) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) TagString

```go
func (n NBTTagInt) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagInt) Type

```go
func (n NBTTagInt) Type() uint8
```

#### func (NBTTagInt) TypeString

```go
func (n NBTTagInt) TypeString() string
```

#### func (*NBTTagInt) WriteTo

```go
func (n *NBTTagInt) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagIntArray

```go
type NBTTagIntArray struct {
}
```


#### func  NewTagIntArray

```go
func NewTagIntArray(name string, data []int32) *NBTTagIntArray
```

#### func (*NBTTagIntArray) Append

```go
func (n *NBTTagIntArray) Append(data int32)
```

#### func (*NBTTagIntArray) Copy

```go
func (n *NBTTagIntArray) Copy() NBTTag
```

#### func (*NBTTagIntArray) Equal

```go
func (n *NBTTagIntArray) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagIntArray) Get

```go
func (n *NBTTagIntArray) Get(index int32) int32
```

#### func (*NBTTagIntArray) GetArray

```go
func (n *NBTTagIntArray) GetArray() []int32
```

#### func (*NBTTagIntArray) Length

```go
func (n *NBTTagIntArray) Length() int32
```

#### func (NBTTagIntArray) Name

```go
func (n NBTTagIntArray) Name() string
```

#### func (*NBTTagIntArray) ReadFrom

```go
func (n *NBTTagIntArray) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagIntArray) Set

```go
func (n *NBTTagIntArray) Set(index, data int32)
```

#### func (NBTTagIntArray) SetName

```go
func (n NBTTagIntArray) SetName(name string)
```

#### func (*NBTTagIntArray) String

```go
func (n *NBTTagIntArray) String() string
```

#### func (NBTTagIntArray) TagByte

```go
func (n NBTTagIntArray) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) TagByteArray

```go
func (n NBTTagIntArray) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) TagCompound

```go
func (n NBTTagIntArray) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) TagDouble

```go
func (n NBTTagIntArray) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) TagEnd

```go
func (n NBTTagIntArray) TagEnd() *NBTTagEnd
```

#### func (NBTTagIntArray) TagFloat

```go
func (n NBTTagIntArray) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) TagInt

```go
func (n NBTTagIntArray) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagIntArray) TagIntArray

```go
func (n *NBTTagIntArray) TagIntArray() *NBTTagIntArray
```

#### func (NBTTagIntArray) TagList

```go
func (n NBTTagIntArray) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) TagLong

```go
func (n NBTTagIntArray) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) TagShort

```go
func (n NBTTagIntArray) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) TagString

```go
func (n NBTTagIntArray) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagIntArray) Type

```go
func (n NBTTagIntArray) Type() uint8
```

#### func (NBTTagIntArray) TypeString

```go
func (n NBTTagIntArray) TypeString() string
```

#### func (*NBTTagIntArray) WriteTo

```go
func (n *NBTTagIntArray) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagList

```go
type NBTTagList struct {
}
```


#### func  NewTagList

```go
func NewTagList(name string, dType uint8, data []interface{}) *NBTTagList
```

#### func (*NBTTagList) Append

```go
func (n *NBTTagList) Append(data interface{}) bool
```

#### func (*NBTTagList) Copy

```go
func (n *NBTTagList) Copy() NBTTag
```

#### func (*NBTTagList) Equal

```go
func (n *NBTTagList) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagList) Get

```go
func (n *NBTTagList) Get(index int32) interface{}
```

#### func (*NBTTagList) GetArray

```go
func (n *NBTTagList) GetArray() []interface{}
```

#### func (*NBTTagList) GetType

```go
func (n *NBTTagList) GetType() uint8
```

#### func (*NBTTagList) Length

```go
func (n *NBTTagList) Length() int32
```

#### func (NBTTagList) Name

```go
func (n NBTTagList) Name() string
```

#### func (*NBTTagList) ReadFrom

```go
func (n *NBTTagList) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagList) Set

```go
func (n *NBTTagList) Set(index int32, data interface{}) bool
```

#### func (*NBTTagList) SetArray

```go
func (n *NBTTagList) SetArray(dType uint8, data []interface{}) bool
```

#### func (NBTTagList) SetName

```go
func (n NBTTagList) SetName(name string)
```

#### func (*NBTTagList) String

```go
func (n *NBTTagList) String() string
```

#### func (NBTTagList) TagByte

```go
func (n NBTTagList) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) TagByteArray

```go
func (n NBTTagList) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) TagCompound

```go
func (n NBTTagList) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) TagDouble

```go
func (n NBTTagList) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) TagEnd

```go
func (n NBTTagList) TagEnd() *NBTTagEnd
```

#### func (NBTTagList) TagFloat

```go
func (n NBTTagList) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) TagInt

```go
func (n NBTTagList) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) TagIntArray

```go
func (n NBTTagList) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagList) TagList

```go
func (n *NBTTagList) TagList() *NBTTagList
```

#### func (NBTTagList) TagLong

```go
func (n NBTTagList) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) TagShort

```go
func (n NBTTagList) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) TagString

```go
func (n NBTTagList) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagList) Type

```go
func (n NBTTagList) Type() uint8
```

#### func (NBTTagList) TypeString

```go
func (n NBTTagList) TypeString() string
```

#### func (*NBTTagList) WriteTo

```go
func (n *NBTTagList) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagLong

```go
type NBTTagLong struct {
}
```


#### func  NewTagLong

```go
func NewTagLong(name string, data int64) *NBTTagLong
```

#### func (*NBTTagLong) Copy

```go
func (n *NBTTagLong) Copy() NBTTag
```

#### func (*NBTTagLong) Equal

```go
func (n *NBTTagLong) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagLong) Get

```go
func (n *NBTTagLong) Get() int64
```

#### func (NBTTagLong) Name

```go
func (n NBTTagLong) Name() string
```

#### func (*NBTTagLong) ReadFrom

```go
func (n *NBTTagLong) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagLong) Set

```go
func (n *NBTTagLong) Set(data int64)
```

#### func (NBTTagLong) SetName

```go
func (n NBTTagLong) SetName(name string)
```

#### func (*NBTTagLong) String

```go
func (n *NBTTagLong) String() string
```

#### func (NBTTagLong) TagByte

```go
func (n NBTTagLong) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) TagByteArray

```go
func (n NBTTagLong) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) TagCompound

```go
func (n NBTTagLong) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) TagDouble

```go
func (n NBTTagLong) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) TagEnd

```go
func (n NBTTagLong) TagEnd() *NBTTagEnd
```

#### func (NBTTagLong) TagFloat

```go
func (n NBTTagLong) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) TagInt

```go
func (n NBTTagLong) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) TagIntArray

```go
func (n NBTTagLong) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) TagList

```go
func (n NBTTagLong) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagLong) TagLong

```go
func (n *NBTTagLong) TagLong() *NBTTagLong
```

#### func (NBTTagLong) TagShort

```go
func (n NBTTagLong) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) TagString

```go
func (n NBTTagLong) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagLong) Type

```go
func (n NBTTagLong) Type() uint8
```

#### func (NBTTagLong) TypeString

```go
func (n NBTTagLong) TypeString() string
```

#### func (*NBTTagLong) WriteTo

```go
func (n *NBTTagLong) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagShort

```go
type NBTTagShort struct {
}
```


#### func  NewTagShort

```go
func NewTagShort(name string, data int16) *NBTTagShort
```

#### func (*NBTTagShort) Copy

```go
func (n *NBTTagShort) Copy() NBTTag
```

#### func (*NBTTagShort) Equal

```go
func (n *NBTTagShort) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagShort) Get

```go
func (n *NBTTagShort) Get() int16
```

#### func (NBTTagShort) Name

```go
func (n NBTTagShort) Name() string
```

#### func (*NBTTagShort) ReadFrom

```go
func (n *NBTTagShort) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagShort) Set

```go
func (n *NBTTagShort) Set(data int16)
```

#### func (NBTTagShort) SetName

```go
func (n NBTTagShort) SetName(name string)
```

#### func (*NBTTagShort) String

```go
func (n *NBTTagShort) String() string
```

#### func (NBTTagShort) TagByte

```go
func (n NBTTagShort) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) TagByteArray

```go
func (n NBTTagShort) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) TagCompound

```go
func (n NBTTagShort) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) TagDouble

```go
func (n NBTTagShort) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) TagEnd

```go
func (n NBTTagShort) TagEnd() *NBTTagEnd
```

#### func (NBTTagShort) TagFloat

```go
func (n NBTTagShort) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) TagInt

```go
func (n NBTTagShort) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) TagIntArray

```go
func (n NBTTagShort) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) TagList

```go
func (n NBTTagShort) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) TagLong

```go
func (n NBTTagShort) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagShort) TagShort

```go
func (n *NBTTagShort) TagShort() *NBTTagShort
```

#### func (NBTTagShort) TagString

```go
func (n NBTTagShort) TagString() *NBTTagString
```
TagString returns the tag if appropriate, nil if otherwise.

#### func (NBTTagShort) Type

```go
func (n NBTTagShort) Type() uint8
```

#### func (NBTTagShort) TypeString

```go
func (n NBTTagShort) TypeString() string
```

#### func (*NBTTagShort) WriteTo

```go
func (n *NBTTagShort) WriteTo(file io.Writer) (int64, error)
```

#### type NBTTagString

```go
type NBTTagString struct {
}
```


#### func  NewTagString

```go
func NewTagString(name string, data string) *NBTTagString
```

#### func (*NBTTagString) Copy

```go
func (n *NBTTagString) Copy() NBTTag
```

#### func (*NBTTagString) Equal

```go
func (n *NBTTagString) Equal(e equaler.Equaler) bool
```

#### func (*NBTTagString) Get

```go
func (n *NBTTagString) Get() string
```

#### func (*NBTTagString) Length

```go
func (n *NBTTagString) Length() int32
```

#### func (NBTTagString) Name

```go
func (n NBTTagString) Name() string
```

#### func (*NBTTagString) ReadFrom

```go
func (n *NBTTagString) ReadFrom(file io.Reader) (int64, error)
```

#### func (*NBTTagString) Set

```go
func (n *NBTTagString) Set(data string)
```

#### func (NBTTagString) SetName

```go
func (n NBTTagString) SetName(name string)
```

#### func (*NBTTagString) String

```go
func (n *NBTTagString) String() string
```

#### func (NBTTagString) TagByte

```go
func (n NBTTagString) TagByte() *NBTTagByte
```
TagByte returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagByteArray

```go
func (n NBTTagString) TagByteArray() *NBTTagByteArray
```
TagByteArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagCompound

```go
func (n NBTTagString) TagCompound() *NBTTagCompound
```
TagCompound returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagDouble

```go
func (n NBTTagString) TagDouble() *NBTTagDouble
```
TagDouble returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagEnd

```go
func (n NBTTagString) TagEnd() *NBTTagEnd
```

#### func (NBTTagString) TagFloat

```go
func (n NBTTagString) TagFloat() *NBTTagFloat
```
TagFloat returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagInt

```go
func (n NBTTagString) TagInt() *NBTTagInt
```
TagInt returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagIntArray

```go
func (n NBTTagString) TagIntArray() *NBTTagIntArray
```
TagIntArray returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagList

```go
func (n NBTTagString) TagList() *NBTTagList
```
TagList returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagLong

```go
func (n NBTTagString) TagLong() *NBTTagLong
```
TagLong returns the tag if appropriate, nil if otherwise.

#### func (NBTTagString) TagShort

```go
func (n NBTTagString) TagShort() *NBTTagShort
```
TagShort returns the tag if appropriate, nil if otherwise.

#### func (*NBTTagString) TagString

```go
func (n *NBTTagString) TagString() *NBTTagString
```

#### func (NBTTagString) Type

```go
func (n NBTTagString) Type() uint8
```

#### func (NBTTagString) TypeString

```go
func (n NBTTagString) TypeString() string
```

#### func (*NBTTagString) WriteTo

```go
func (n *NBTTagString) WriteTo(file io.Writer) (int64, error)
```
