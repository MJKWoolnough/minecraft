# nbt
--
    import "github.com/MJKWoolnough/minecraft/nbt"

Package nbt implements a full Named Binary Tag reader/writer, based on the specs at
http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt

## Usage

```go
var BadRange = errors.New("given index was out-of-range")
```
BadRange is an error that occurs when trying to set an item on a list which is
outside of the current limits of the list.

#### type Byte

```go
type Byte int8
```

Byte is an implementation of the Data interface

#### func (Byte) Copy

```go
func (b Byte) Copy() Data
```
Copy simply returns a copy the the data

#### func (Byte) Equal

```go
func (b Byte) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (Byte) String

```go
func (b Byte) String() string
```

#### func (Byte) Type

```go
func (Byte) Type() TagID
```
Type returns the TagID of the data

#### type ByteArray

```go
type ByteArray []int8
```

ByteArray is an implementation of the Data interface

#### func (ByteArray) Bytes

```go
func (b ByteArray) Bytes() []byte
```
Converts the ByteArray (actually int8) to an actual slice of bytes. NB: Uses
unsafe, so the underlying array is the same.

#### func (ByteArray) Copy

```go
func (b ByteArray) Copy() Data
```
Copy simply returns a copy the the data

#### func (ByteArray) Equal

```go
func (b ByteArray) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ByteArray) String

```go
func (b ByteArray) String() string
```

#### func (ByteArray) Type

```go
func (ByteArray) Type() TagID
```
Type returns the TagID of the data

#### type Compound

```go
type Compound []Tag
```

Compound is an implementation of the Data interface

#### func (Compound) Copy

```go
func (c Compound) Copy() Data
```
Copy simply returns a deep-copy the the data

#### func (Compound) Equal

```go
func (c Compound) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (Compound) Get

```go
func (c Compound) Get(name string) Tag
```
Get returns the tag for the given name

#### func (*Compound) Remove

```go
func (c *Compound) Remove(name string)
```
Remove removes the tag corresponding to the given name

#### func (*Compound) Set

```go
func (c *Compound) Set(tag Tag)
```
Set adds the given tag to the compound, or, if the tags name is already present,
overrides the old data

#### func (Compound) String

```go
func (c Compound) String() string
```

#### func (Compound) Type

```go
func (Compound) Type() TagID
```
Type returns the TagID of the data

#### type Data

```go
type Data interface {
	equaler.Equaler
	Copy() Data
	String() string
	Type() TagID
}
```

Data is an interface representing the many different types that a tag can be

#### type Decoder

```go
type Decoder struct {
}
```

Decoder is a type used to decode NBT streams

#### func  NewDecoder

```go
func NewDecoder(r io.Reader) Decoder
```
NewDecoder returns a Decoder using Big Endian

#### func  NewDecoderEndian

```go
func NewDecoderEndian(e byteio.EndianReader) Decoder
```
NewDecoderEndian allows you to specify your own Endian Reader

#### func (Decoder) DecodeByte

```go
func (d Decoder) DecodeByte() (Byte, error)
```
DecodeByte will read a single Byte Data

#### func (Decoder) DecodeByteArray

```go
func (d Decoder) DecodeByteArray() (ByteArray, error)
```
DecodeByteArray will read a ByteArray Data

#### func (Decoder) DecodeCompound

```go
func (d Decoder) DecodeCompound() (Compound, error)
```
DecodeCompound will read a Compound Data

#### func (Decoder) DecodeDouble

```go
func (d Decoder) DecodeDouble() (Double, error)
```
DecodeDouble will read a single Double Data

#### func (Decoder) DecodeFloat

```go
func (d Decoder) DecodeFloat() (Float, error)
```
DecodeFloat will read a single Float Data

#### func (Decoder) DecodeInt

```go
func (d Decoder) DecodeInt() (Int, error)
```
DecodeInt will read a single Int Data

#### func (Decoder) DecodeIntArray

```go
func (d Decoder) DecodeIntArray() (IntArray, error)
```
DecodeIntArray will read an IntArray Data

#### func (Decoder) DecodeList

```go
func (d Decoder) DecodeList() (*List, error)
```
DecodeList will read a List Data

#### func (Decoder) DecodeLong

```go
func (d Decoder) DecodeLong() (Long, error)
```
DecodeLong will read a single Long Data

#### func (Decoder) DecodeShort

```go
func (d Decoder) DecodeShort() (Short, error)
```
DecodeShort will read a single Short Data

#### func (Decoder) DecodeString

```go
func (d Decoder) DecodeString() (String, error)
```
DecodeString will read a String Data

#### func (Decoder) DecodeTag

```go
func (d Decoder) DecodeTag() (Tag, error)
```
DecodeTag will read a whole tag out of the decoding stream

#### type Double

```go
type Double float64
```

Double is an implementation of the Data interface

#### func (Double) Copy

```go
func (d Double) Copy() Data
```
Copy simply returns a copy the the data

#### func (Double) Equal

```go
func (d Double) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (Double) String

```go
func (d Double) String() string
```

#### func (Double) Type

```go
func (Double) Type() TagID
```
Type returns the TagID of the data

#### type Encoder

```go
type Encoder struct {
}
```

Encoder is a type used to encode NBT streams

#### func  NewEncoder

```go
func NewEncoder(w io.Writer) Encoder
```
NewEncoder returns an Encoder using Big Endian

#### func  NewEncoderEndian

```go
func NewEncoderEndian(e byteio.EndianWriter) Encoder
```
NewEncoderEndian allows you to specify your own Endian Writer

#### func (Encoder) EncodeByte

```go
func (e Encoder) EncodeByte(b Byte) error
```
EncodeByte will write a single Byte Data

#### func (Encoder) EncodeByteArray

```go
func (e Encoder) EncodeByteArray(ba ByteArray) error
```
EncodeByteArray will write a ByteArray Data

#### func (Encoder) EncodeCompound

```go
func (e Encoder) EncodeCompound(c Compound) error
```
EncodeCompound will write a Compound Data

#### func (Encoder) EncodeDouble

```go
func (e Encoder) EncodeDouble(do Double) error
```
EncodeDouble will write a single Double Data

#### func (Encoder) EncodeFloat

```go
func (e Encoder) EncodeFloat(f Float) error
```
EncodeFloat will write a single Float Data

#### func (Encoder) EncodeInt

```go
func (e Encoder) EncodeInt(i Int) error
```
EncodeInt will write a single Int Data

#### func (Encoder) EncodeIntArray

```go
func (e Encoder) EncodeIntArray(ints IntArray) error
```
EncodeIntArray will write a IntArray Data

#### func (Encoder) EncodeList

```go
func (e Encoder) EncodeList(l *List) error
```
EncodeList will write a List Data

#### func (Encoder) EncodeLong

```go
func (e Encoder) EncodeLong(l Long) error
```
EncodeLong will write a single Long Data

#### func (Encoder) EncodeShort

```go
func (e Encoder) EncodeShort(s Short) error
```
EncodeShort will write a single Short Data

#### func (Encoder) EncodeString

```go
func (e Encoder) EncodeString(s String) error
```
EncodeString will write a String Data

#### func (Encoder) EncodeTag

```go
func (e Encoder) EncodeTag(t Tag) error
```
EncodeTag will encode a whole tag to the encoding stream

#### type Float

```go
type Float float32
```

Float is an implementation of the Data interface

#### func (Float) Copy

```go
func (f Float) Copy() Data
```
Copy simply returns a copy the the data

#### func (Float) Equal

```go
func (f Float) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (Float) String

```go
func (f Float) String() string
```

#### func (Float) Type

```go
func (Float) Type() TagID
```
Type returns the TagID of the data

#### type Int

```go
type Int int32
```

Int is an implementation of the Data interface

#### func (Int) Copy

```go
func (i Int) Copy() Data
```
Copy simply returns a copy the the data

#### func (Int) Equal

```go
func (i Int) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (Int) String

```go
func (i Int) String() string
```

#### func (Int) Type

```go
func (Int) Type() TagID
```
Type returns the TagID of the data

#### type IntArray

```go
type IntArray []int32
```

IntArray is an implementation of the Data interface

#### func (IntArray) Copy

```go
func (i IntArray) Copy() Data
```
Copy simply returns a copy the the data

#### func (IntArray) Equal

```go
func (i IntArray) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (IntArray) String

```go
func (i IntArray) String() string
```

#### func (IntArray) Type

```go
func (IntArray) Type() TagID
```
Type returns the TagID of the data

#### type List

```go
type List struct {
}
```

List is an implementation of the Data interface

#### func  NewEmptyList

```go
func NewEmptyList(tagType TagID) *List
```
NewEmptyList returns a new empty List Data type, set to the type given

#### func  NewList

```go
func NewList(data []Data) *List
```
NewList returns a new List Data type, or nil if the given data is not of all the
same Data type

#### func (*List) Append

```go
func (l *List) Append(data ...Data) error
```
Append adds data to the list

#### func (*List) Copy

```go
func (l *List) Copy() Data
```
Copy simply returns a deep-copy the the data

#### func (*List) Equal

```go
func (l *List) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (*List) Get

```go
func (l *List) Get(i int) Data
```
Get returns the data at the given positon

#### func (*List) Insert

```go
func (l *List) Insert(i int, data ...Data) error
```
Insert will add the given data at the specified position, moving other elements
up.

#### func (*List) Len

```go
func (l *List) Len() int
```
Len returns the length of the list

#### func (*List) Remove

```go
func (l *List) Remove(i int)
```
Remove deletes the specified position and shifts remaing data down

#### func (*List) Set

```go
func (l *List) Set(i int32, data Data) error
```
Set sets the data at the given position. It does not append

#### func (*List) String

```go
func (l *List) String() string
```

#### func (*List) TagType

```go
func (l *List) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains

#### func (List) Type

```go
func (List) Type() TagID
```
Type returns the TagID of the data

#### type Long

```go
type Long int64
```

Long is an implementation of the Data interface

#### func (Long) Copy

```go
func (l Long) Copy() Data
```
Copy simply returns a copy the the data

#### func (Long) Equal

```go
func (l Long) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (Long) String

```go
func (l Long) String() string
```

#### func (Long) Type

```go
func (Long) Type() TagID
```
Type returns the TagID of the data

#### type ReadError

```go
type ReadError struct {
	Where string
	Err   error
}
```

ReadError is an error returned when a read error occurs

#### func (ReadError) Error

```go
func (r ReadError) Error() string
```

#### type Short

```go
type Short int16
```

Short is an implementation of the Data interface

#### func (Short) Copy

```go
func (s Short) Copy() Data
```
Copy simply returns a copy the the data

#### func (Short) Equal

```go
func (s Short) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (Short) String

```go
func (s Short) String() string
```

#### func (Short) Type

```go
func (Short) Type() TagID
```
Type returns the TagID of the data

#### type String

```go
type String string
```

String is an implementation of the Data interface

#### func (String) Copy

```go
func (s String) Copy() Data
```
Copy simply returns a copy the the data

#### func (String) Equal

```go
func (s String) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (String) String

```go
func (s String) String() string
```

#### func (String) Type

```go
func (String) Type() TagID
```
Type returns the TagID of the data

#### type Tag

```go
type Tag struct {
}
```

Tag is the main NBT type, a combination of a name and a Data type

#### func  NewTag

```go
func NewTag(name string, d Data) Tag
```
NewTag constructs a new tag with the given name and data.

#### func (Tag) Copy

```go
func (t Tag) Copy() Tag
```
Copy simply returns a deep-copy the the tag

#### func (Tag) Data

```go
func (t Tag) Data() Data
```
Data returns the tags data type

#### func (Tag) Equal

```go
func (t Tag) Equal(e equaler.Equaler) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (Tag) Name

```go
func (t Tag) Name() string
```
Name returns the tags' name

#### func (Tag) String

```go
func (t Tag) String() string
```
String returns a textual representation of the tag

#### func (Tag) TagID

```go
func (t Tag) TagID() TagID
```
TagID returns the type of the data

#### type TagID

```go
type TagID uint8
```

TagID represents the type of nbt tag

```go
const (
	TagEnd       TagID = 0
	TagByte      TagID = 1
	TagShort     TagID = 2
	TagInt       TagID = 3
	TagLong      TagID = 4
	TagFloat     TagID = 5
	TagDouble    TagID = 6
	TagByteArray TagID = 7
	TagString    TagID = 8
	TagList      TagID = 9
	TagCompound  TagID = 10
	TagIntArray  TagID = 11
)
```
Tag Types

#### func (TagID) String

```go
func (t TagID) String() string
```

#### type UnknownTag

```go
type UnknownTag struct {
	TagID
}
```

UnknownTag is an error that occurs when an unknown tag id is discovered. This
could also indicate corrupted or non-compliant data.

#### func (UnknownTag) Error

```go
func (u UnknownTag) Error() string
```

#### type WriteError

```go
type WriteError struct {
	Where string
	Err   error
}
```

WriteError is an error returned when a write error occurs

#### func (WriteError) Error

```go
func (w WriteError) Error() string
```

#### type WrongTag

```go
type WrongTag struct {
	Expecting, Got TagID
}
```

WrongTag is an error returned when a tag of the incorrect type was intended to
be added to a list.

#### func (WrongTag) Error

```go
func (w WrongTag) Error() string
```
