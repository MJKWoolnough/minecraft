# nbt
--
    import "vimagination.zapto.org/minecraft/nbt"

Package nbt implements a full Named Binary Tag reader/writer, based on the specs
### at
http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt

## Usage

```go
var ErrBadRange = errors.New("given index was out-of-range")
```
ErrBadRange is an error that occurs when trying to set an item on a list which
is outside of the current limits of the list.

#### func  Encode

```go
func Encode(w io.Writer, t Tag) error
```
Encode will encode a single tag to the writer using the default settings

#### type Bool

```go
type Bool bool
```

Bool is an implementation of the Data interface.

#### func (Bool) Copy

```go
func (b Bool) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Bool) Equal

```go
func (b Bool) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Bool) String

```go
func (b Bool) String() string
```

#### func (Bool) Type

```go
func (Bool) Type() TagID
```
Type returns the TagID of the data.

#### type Byte

```go
type Byte int8
```

Byte is an implementation of the Data interface NB: Despite being an unsigned
integer, it is still called a byte to match the spec.

#### func (Byte) Copy

```go
func (b Byte) Copy() Data
```
Copy simply returns a copy of the data

#### func (Byte) Equal

```go
func (b Byte) Equal(e interface{}) bool
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
Type returns the TagID of the data.

#### type ByteArray

```go
type ByteArray []int8
```

ByteArray is an implementation of the Data interface.

#### func (ByteArray) Bytes

```go
func (b ByteArray) Bytes() []byte
```
Bytes converts the ByteArray (actually int8) to an actual slice of bytes. NB:
May uss unsafe, so the underlying array may be the same.

#### func (ByteArray) Copy

```go
func (b ByteArray) Copy() Data
```
Copy simply returns a copy of the data.

#### func (ByteArray) Equal

```go
func (b ByteArray) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (ByteArray) String

```go
func (b ByteArray) String() string
```

#### func (ByteArray) Type

```go
func (ByteArray) Type() TagID
```
Type returns the TagID of the data.

#### type Complex128

```go
type Complex128 complex128
```

Complex128 is an implementation of the Data interface.

#### func (Complex128) Copy

```go
func (c Complex128) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Complex128) Equal

```go
func (c Complex128) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Complex128) String

```go
func (c Complex128) String() string
```

#### func (Complex128) Type

```go
func (Complex128) Type() TagID
```
Type returns the TagID of the data.

#### type Complex64

```go
type Complex64 complex64
```

Complex64 is an implementation of the Data interface.

#### func (Complex64) Copy

```go
func (c Complex64) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Complex64) Equal

```go
func (c Complex64) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Complex64) String

```go
func (c Complex64) String() string
```

#### func (Complex64) Type

```go
func (Complex64) Type() TagID
```
Type returns the TagID of the data.

#### type Compound

```go
type Compound []Tag
```

Compound is an implementation of the Data interface.

#### func (Compound) Copy

```go
func (c Compound) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (Compound) Equal

```go
func (c Compound) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Compound) Get

```go
func (c Compound) Get(name string) Tag
```
Get returns the tag for the given name.

#### func (*Compound) Remove

```go
func (c *Compound) Remove(name string)
```
Remove removes the tag corresponding to the given name.

#### func (*Compound) Set

```go
func (c *Compound) Set(tag Tag)
```
Set adds the given tag to the compound, or, if the tags name is already present,
overrides the old data.

#### func (Compound) String

```go
func (c Compound) String() string
```

#### func (Compound) Type

```go
func (Compound) Type() TagID
```
Type returns the TagID of the data.

#### type Data

```go
type Data interface {
	Equal(interface{}) bool
	Copy() Data
	String() string
	Type() TagID
}
```

Data is an interface representing the many different types that a tag can be.

#### type Decoder

```go
type Decoder struct {
}
```

Decoder is a type used to decode NBT streams.

#### func  NewDecoder

```go
func NewDecoder(r io.Reader) Decoder
```
NewDecoder returns a Decoder using Big Endian.

#### func  NewDecoderEndian

```go
func NewDecoderEndian(e byteio.EndianReader) Decoder
```
NewDecoderEndian allows you to specify your own Endian Reader.

#### func (Decoder) Decode

```go
func (d Decoder) Decode() (Tag, error)
```
Decode will read a whole tag out of the decoding stream.

#### type Double

```go
type Double float64
```

Double is an implementation of the Data interface.

#### func (Double) Copy

```go
func (d Double) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Double) Equal

```go
func (d Double) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Double) String

```go
func (d Double) String() string
```

#### func (Double) Type

```go
func (Double) Type() TagID
```
Type returns the TagID of the data.

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

#### func (Encoder) Encode

```go
func (e Encoder) Encode(t Tag) error
```
Encode will encode a whole tag to the encoding stream

#### type Float

```go
type Float float32
```

Float is an implementation of the Data interface.

#### func (Float) Copy

```go
func (f Float) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Float) Equal

```go
func (f Float) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Float) String

```go
func (f Float) String() string
```

#### func (Float) Type

```go
func (Float) Type() TagID
```
Type returns the TagID of the data.

#### type Int

```go
type Int int32
```

Int is an implementation of the Data interface.

#### func (Int) Copy

```go
func (i Int) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Int) Equal

```go
func (i Int) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Int) String

```go
func (i Int) String() string
```

#### func (Int) Type

```go
func (Int) Type() TagID
```
Type returns the TagID of the data.

#### type IntArray

```go
type IntArray []int32
```

IntArray is an implementation of the Data interface.

#### func (IntArray) Copy

```go
func (i IntArray) Copy() Data
```
Copy simply returns a copy of the data.

#### func (IntArray) Equal

```go
func (i IntArray) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (IntArray) String

```go
func (i IntArray) String() string
```

#### func (IntArray) Type

```go
func (IntArray) Type() TagID
```
Type returns the TagID of the data.

#### type List

```go
type List interface {
	Data
	Set(int, Data) error
	Get(int) Data
	Append(...Data) error
	Insert(int, ...Data) error
	Remove(int)
	Len() int
	TagType() TagID
}
```

List interface descibes the methods for the lists of different data types.

#### func  NewEmptyList

```go
func NewEmptyList(tagType TagID) List
```
NewEmptyList returns a new empty List Data type, set to the type given.

#### func  NewList

```go
func NewList(data []Data) List
```
NewList returns a new List Data type, or nil if the given data is not of all the
same Data type.

#### type ListBool

```go
type ListBool []Bool
```

ListBool satisfies the List interface for a list of Bools.

#### func (*ListBool) Append

```go
func (l *ListBool) Append(d ...Data) error
```
Append adds data to the list

#### func (ListBool) Copy

```go
func (l ListBool) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListBool) Equal

```go
func (l ListBool) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListBool) Get

```go
func (l ListBool) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListBool) Insert

```go
func (l *ListBool) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListBool) Len

```go
func (l ListBool) Len() int
```
Len returns the length of the list.

#### func (*ListBool) Remove

```go
func (l *ListBool) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListBool) Set

```go
func (l ListBool) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListBool) String

```go
func (l ListBool) String() string
```

#### func (ListBool) TagType

```go
func (ListBool) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListBool) Type

```go
func (ListBool) Type() TagID
```
Type returns the TagID of the data.

#### type ListByte

```go
type ListByte []Byte
```

ListByte satisfies the List interface for a list of Bytes.

#### func (*ListByte) Append

```go
func (l *ListByte) Append(d ...Data) error
```
Append adds data to the list

#### func (ListByte) Copy

```go
func (l ListByte) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListByte) Equal

```go
func (l ListByte) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListByte) Get

```go
func (l ListByte) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListByte) Insert

```go
func (l *ListByte) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListByte) Len

```go
func (l ListByte) Len() int
```
Len returns the length of the list.

#### func (*ListByte) Remove

```go
func (l *ListByte) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListByte) Set

```go
func (l ListByte) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListByte) String

```go
func (l ListByte) String() string
```

#### func (ListByte) TagType

```go
func (ListByte) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListByte) Type

```go
func (ListByte) Type() TagID
```
Type returns the TagID of the data.

#### type ListComplex128

```go
type ListComplex128 []Complex128
```

ListComplex128 satisfies the List interface for a list of Complex128s.

#### func (*ListComplex128) Append

```go
func (l *ListComplex128) Append(d ...Data) error
```
Append adds data to the list

#### func (ListComplex128) Copy

```go
func (l ListComplex128) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListComplex128) Equal

```go
func (l ListComplex128) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListComplex128) Get

```go
func (l ListComplex128) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListComplex128) Insert

```go
func (l *ListComplex128) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListComplex128) Len

```go
func (l ListComplex128) Len() int
```
Len returns the length of the list.

#### func (*ListComplex128) Remove

```go
func (l *ListComplex128) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListComplex128) Set

```go
func (l ListComplex128) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListComplex128) String

```go
func (l ListComplex128) String() string
```

#### func (ListComplex128) TagType

```go
func (ListComplex128) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListComplex128) Type

```go
func (ListComplex128) Type() TagID
```
Type returns the TagID of the data.

#### type ListComplex64

```go
type ListComplex64 []Complex64
```

ListComplex64 satisfies the List interface for a list of Complex64s.

#### func (*ListComplex64) Append

```go
func (l *ListComplex64) Append(d ...Data) error
```
Append adds data to the list

#### func (ListComplex64) Copy

```go
func (l ListComplex64) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListComplex64) Equal

```go
func (l ListComplex64) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListComplex64) Get

```go
func (l ListComplex64) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListComplex64) Insert

```go
func (l *ListComplex64) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListComplex64) Len

```go
func (l ListComplex64) Len() int
```
Len returns the length of the list.

#### func (*ListComplex64) Remove

```go
func (l *ListComplex64) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListComplex64) Set

```go
func (l ListComplex64) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListComplex64) String

```go
func (l ListComplex64) String() string
```

#### func (ListComplex64) TagType

```go
func (ListComplex64) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListComplex64) Type

```go
func (ListComplex64) Type() TagID
```
Type returns the TagID of the data.

#### type ListCompound

```go
type ListCompound []Compound
```

ListCompound satisfies the List interface for a list of Compounds.

#### func (*ListCompound) Append

```go
func (l *ListCompound) Append(d ...Data) error
```
Append adds data to the list

#### func (ListCompound) Copy

```go
func (l ListCompound) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListCompound) Equal

```go
func (l ListCompound) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListCompound) Get

```go
func (l ListCompound) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListCompound) Insert

```go
func (l *ListCompound) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListCompound) Len

```go
func (l ListCompound) Len() int
```
Len returns the length of the list.

#### func (*ListCompound) Remove

```go
func (l *ListCompound) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListCompound) Set

```go
func (l ListCompound) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListCompound) String

```go
func (l ListCompound) String() string
```

#### func (ListCompound) TagType

```go
func (ListCompound) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListCompound) Type

```go
func (ListCompound) Type() TagID
```
Type returns the TagID of the data.

#### type ListData

```go
type ListData struct {
}
```

ListData is an implementation of the Data interface.

#### func (*ListData) Append

```go
func (l *ListData) Append(data ...Data) error
```
Append adds data to the list.

#### func (*ListData) Copy

```go
func (l *ListData) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (*ListData) Equal

```go
func (l *ListData) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (*ListData) Get

```go
func (l *ListData) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListData) Insert

```go
func (l *ListData) Insert(i int, data ...Data) error
```
Insert will add the given data at the specified position, moving other elements
up.

#### func (*ListData) Len

```go
func (l *ListData) Len() int
```
Len returns the length of the list.

#### func (ListData) ListBool

```go
func (l ListData) ListBool() ListBool
```
ListBool returns the list as a specifically typed List.

#### func (ListData) ListByte

```go
func (l ListData) ListByte() ListByte
```
ListByte returns the list as a specifically typed List.

#### func (ListData) ListComplex128

```go
func (l ListData) ListComplex128() ListComplex128
```
ListComplex128 returns the list as a specifically typed List.

#### func (ListData) ListComplex64

```go
func (l ListData) ListComplex64() ListComplex64
```
ListComplex64 returns the list as a specifically typed List.

#### func (ListData) ListCompound

```go
func (l ListData) ListCompound() ListCompound
```
ListCompound returns the list as a specifically typed List.

#### func (ListData) ListDouble

```go
func (l ListData) ListDouble() ListDouble
```
ListDouble returns the list as a specifically typed List.

#### func (ListData) ListEnd

```go
func (l ListData) ListEnd() ListEnd
```
ListEnd returns the list as a specifically typed List.

#### func (ListData) ListFloat

```go
func (l ListData) ListFloat() ListFloat
```
ListFloat returns the list as a specifically typed List.

#### func (ListData) ListInt

```go
func (l ListData) ListInt() ListInt
```
ListInt returns the list as a specifically typed List.

#### func (ListData) ListIntArray

```go
func (l ListData) ListIntArray() ListIntArray
```
ListIntArray returns the list as a specifically typed List.

#### func (ListData) ListLong

```go
func (l ListData) ListLong() ListLong
```
ListLong returns the list as a specifically typed List.

#### func (ListData) ListShort

```go
func (l ListData) ListShort() ListShort
```
ListShort returns the list as a specifically typed List.

#### func (ListData) ListUint16

```go
func (l ListData) ListUint16() ListUint16
```
ListUint16 returns the list as a specifically typed List.

#### func (ListData) ListUint32

```go
func (l ListData) ListUint32() ListUint32
```
ListUint32 returns the list as a specifically typed List.

#### func (ListData) ListUint64

```go
func (l ListData) ListUint64() ListUint64
```
ListUint64 returns the list as a specifically typed List.

#### func (ListData) ListUint8

```go
func (l ListData) ListUint8() ListUint8
```
ListUint8 returns the list as a specifically typed List.

#### func (*ListData) Remove

```go
func (l *ListData) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (*ListData) Set

```go
func (l *ListData) Set(i int, data Data) error
```
Set sets the data at the given position. It does not append.

#### func (*ListData) String

```go
func (l *ListData) String() string
```

#### func (*ListData) TagType

```go
func (l *ListData) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListData) Type

```go
func (ListData) Type() TagID
```
Type returns the TagID of the data.

#### type ListDouble

```go
type ListDouble []Double
```

ListDouble satisfies the List interface for a list of Doubles.

#### func (*ListDouble) Append

```go
func (l *ListDouble) Append(d ...Data) error
```
Append adds data to the list

#### func (ListDouble) Copy

```go
func (l ListDouble) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListDouble) Equal

```go
func (l ListDouble) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListDouble) Get

```go
func (l ListDouble) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListDouble) Insert

```go
func (l *ListDouble) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListDouble) Len

```go
func (l ListDouble) Len() int
```
Len returns the length of the list.

#### func (*ListDouble) Remove

```go
func (l *ListDouble) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListDouble) Set

```go
func (l ListDouble) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListDouble) String

```go
func (l ListDouble) String() string
```

#### func (ListDouble) TagType

```go
func (ListDouble) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListDouble) Type

```go
func (ListDouble) Type() TagID
```
Type returns the TagID of the data.

#### type ListEnd

```go
type ListEnd uint32
```

ListEnd satisfies the List interface for a list of Ends.

#### func (*ListEnd) Append

```go
func (l *ListEnd) Append(d ...Data) error
```
Append adds to the list.

#### func (ListEnd) Copy

```go
func (l ListEnd) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListEnd) Equal

```go
func (l ListEnd) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (ListEnd) Get

```go
func (ListEnd) Get(_ int) Data
```
Get returns an end{}.

#### func (*ListEnd) Insert

```go
func (l *ListEnd) Insert(_ int, d ...Data) error
```
Insert adds to the list.

#### func (ListEnd) Len

```go
func (l ListEnd) Len() int
```
Len returns the length of the List.

#### func (*ListEnd) Remove

```go
func (l *ListEnd) Remove(i int)
```
Remove removes from the list.

#### func (ListEnd) Set

```go
func (l ListEnd) Set(_ int, d Data) error
```
Set does nothing as it's not applicable for ListEnd.

#### func (ListEnd) String

```go
func (l ListEnd) String() string
```

#### func (ListEnd) TagType

```go
func (ListEnd) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListEnd) Type

```go
func (ListEnd) Type() TagID
```
Type returns the TagID of the data.

#### type ListFloat

```go
type ListFloat []Float
```

ListFloat satisfies the List interface for a list of Floats.

#### func (*ListFloat) Append

```go
func (l *ListFloat) Append(d ...Data) error
```
Append adds data to the list

#### func (ListFloat) Copy

```go
func (l ListFloat) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListFloat) Equal

```go
func (l ListFloat) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListFloat) Get

```go
func (l ListFloat) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListFloat) Insert

```go
func (l *ListFloat) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListFloat) Len

```go
func (l ListFloat) Len() int
```
Len returns the length of the list.

#### func (*ListFloat) Remove

```go
func (l *ListFloat) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListFloat) Set

```go
func (l ListFloat) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListFloat) String

```go
func (l ListFloat) String() string
```

#### func (ListFloat) TagType

```go
func (ListFloat) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListFloat) Type

```go
func (ListFloat) Type() TagID
```
Type returns the TagID of the data.

#### type ListInt

```go
type ListInt []Int
```

ListInt satisfies the List interface for a list of Ints.

#### func (*ListInt) Append

```go
func (l *ListInt) Append(d ...Data) error
```
Append adds data to the list

#### func (ListInt) Copy

```go
func (l ListInt) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListInt) Equal

```go
func (l ListInt) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListInt) Get

```go
func (l ListInt) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListInt) Insert

```go
func (l *ListInt) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListInt) Len

```go
func (l ListInt) Len() int
```
Len returns the length of the list.

#### func (*ListInt) Remove

```go
func (l *ListInt) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListInt) Set

```go
func (l ListInt) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListInt) String

```go
func (l ListInt) String() string
```

#### func (ListInt) TagType

```go
func (ListInt) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListInt) Type

```go
func (ListInt) Type() TagID
```
Type returns the TagID of the data.

#### type ListIntArray

```go
type ListIntArray []IntArray
```

ListIntArray satisfies the List interface for a list of IntArrays.

#### func (*ListIntArray) Append

```go
func (l *ListIntArray) Append(d ...Data) error
```
Append adds data to the list

#### func (ListIntArray) Copy

```go
func (l ListIntArray) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListIntArray) Equal

```go
func (l ListIntArray) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListIntArray) Get

```go
func (l ListIntArray) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListIntArray) Insert

```go
func (l *ListIntArray) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListIntArray) Len

```go
func (l ListIntArray) Len() int
```
Len returns the length of the list.

#### func (*ListIntArray) Remove

```go
func (l *ListIntArray) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListIntArray) Set

```go
func (l ListIntArray) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListIntArray) String

```go
func (l ListIntArray) String() string
```

#### func (ListIntArray) TagType

```go
func (ListIntArray) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListIntArray) Type

```go
func (ListIntArray) Type() TagID
```
Type returns the TagID of the data.

#### type ListLong

```go
type ListLong []Long
```

ListLong satisfies the List interface for a list of Longs.

#### func (*ListLong) Append

```go
func (l *ListLong) Append(d ...Data) error
```
Append adds data to the list

#### func (ListLong) Copy

```go
func (l ListLong) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListLong) Equal

```go
func (l ListLong) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListLong) Get

```go
func (l ListLong) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListLong) Insert

```go
func (l *ListLong) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListLong) Len

```go
func (l ListLong) Len() int
```
Len returns the length of the list.

#### func (*ListLong) Remove

```go
func (l *ListLong) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListLong) Set

```go
func (l ListLong) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListLong) String

```go
func (l ListLong) String() string
```

#### func (ListLong) TagType

```go
func (ListLong) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListLong) Type

```go
func (ListLong) Type() TagID
```
Type returns the TagID of the data.

#### type ListShort

```go
type ListShort []Short
```

ListShort satisfies the List interface for a list of Shorts.

#### func (*ListShort) Append

```go
func (l *ListShort) Append(d ...Data) error
```
Append adds data to the list

#### func (ListShort) Copy

```go
func (l ListShort) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListShort) Equal

```go
func (l ListShort) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListShort) Get

```go
func (l ListShort) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListShort) Insert

```go
func (l *ListShort) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListShort) Len

```go
func (l ListShort) Len() int
```
Len returns the length of the list.

#### func (*ListShort) Remove

```go
func (l *ListShort) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListShort) Set

```go
func (l ListShort) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListShort) String

```go
func (l ListShort) String() string
```

#### func (ListShort) TagType

```go
func (ListShort) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListShort) Type

```go
func (ListShort) Type() TagID
```
Type returns the TagID of the data.

#### type ListUint16

```go
type ListUint16 []Uint16
```

ListUint16 satisfies the List interface for a list of Uint16s.

#### func (*ListUint16) Append

```go
func (l *ListUint16) Append(d ...Data) error
```
Append adds data to the list

#### func (ListUint16) Copy

```go
func (l ListUint16) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListUint16) Equal

```go
func (l ListUint16) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListUint16) Get

```go
func (l ListUint16) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListUint16) Insert

```go
func (l *ListUint16) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListUint16) Len

```go
func (l ListUint16) Len() int
```
Len returns the length of the list.

#### func (*ListUint16) Remove

```go
func (l *ListUint16) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListUint16) Set

```go
func (l ListUint16) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListUint16) String

```go
func (l ListUint16) String() string
```

#### func (ListUint16) TagType

```go
func (ListUint16) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListUint16) Type

```go
func (ListUint16) Type() TagID
```
Type returns the TagID of the data.

#### type ListUint32

```go
type ListUint32 []Uint32
```

ListUint32 satisfies the List interface for a list of Uint32s.

#### func (*ListUint32) Append

```go
func (l *ListUint32) Append(d ...Data) error
```
Append adds data to the list

#### func (ListUint32) Copy

```go
func (l ListUint32) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListUint32) Equal

```go
func (l ListUint32) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListUint32) Get

```go
func (l ListUint32) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListUint32) Insert

```go
func (l *ListUint32) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListUint32) Len

```go
func (l ListUint32) Len() int
```
Len returns the length of the list.

#### func (*ListUint32) Remove

```go
func (l *ListUint32) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListUint32) Set

```go
func (l ListUint32) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListUint32) String

```go
func (l ListUint32) String() string
```

#### func (ListUint32) TagType

```go
func (ListUint32) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListUint32) Type

```go
func (ListUint32) Type() TagID
```
Type returns the TagID of the data.

#### type ListUint64

```go
type ListUint64 []Uint64
```

ListUint64 satisfies the List interface for a list of Uint64s.

#### func (*ListUint64) Append

```go
func (l *ListUint64) Append(d ...Data) error
```
Append adds data to the list

#### func (ListUint64) Copy

```go
func (l ListUint64) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListUint64) Equal

```go
func (l ListUint64) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListUint64) Get

```go
func (l ListUint64) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListUint64) Insert

```go
func (l *ListUint64) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListUint64) Len

```go
func (l ListUint64) Len() int
```
Len returns the length of the list.

#### func (*ListUint64) Remove

```go
func (l *ListUint64) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListUint64) Set

```go
func (l ListUint64) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListUint64) String

```go
func (l ListUint64) String() string
```

#### func (ListUint64) TagType

```go
func (ListUint64) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListUint64) Type

```go
func (ListUint64) Type() TagID
```
Type returns the TagID of the data.

#### type ListUint8

```go
type ListUint8 []Uint8
```

ListUint8 satisfies the List interface for a list of Uint8s.

#### func (*ListUint8) Append

```go
func (l *ListUint8) Append(d ...Data) error
```
Append adds data to the list

#### func (ListUint8) Copy

```go
func (l ListUint8) Copy() Data
```
Copy simply returns a deep-copy of the data.

#### func (ListUint8) Equal

```go
func (l ListUint8) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality

#### func (ListUint8) Get

```go
func (l ListUint8) Get(i int) Data
```
Get returns the data at the given position.

#### func (*ListUint8) Insert

```go
func (l *ListUint8) Insert(i int, d ...Data) error
```
Insert will add the given data at the specified position, moving other up.

#### func (ListUint8) Len

```go
func (l ListUint8) Len() int
```
Len returns the length of the list.

#### func (*ListUint8) Remove

```go
func (l *ListUint8) Remove(i int)
```
Remove deletes the specified position and shifts remaining data down.

#### func (ListUint8) Set

```go
func (l ListUint8) Set(i int, d Data) error
```
Set sets the data at the given position. It does not append.

#### func (ListUint8) String

```go
func (l ListUint8) String() string
```

#### func (ListUint8) TagType

```go
func (ListUint8) TagType() TagID
```
TagType returns the TagID of the type of tag this list contains.

#### func (ListUint8) Type

```go
func (ListUint8) Type() TagID
```
Type returns the TagID of the data.

#### type Long

```go
type Long int64
```

Long is an implementation of the Data interface.

#### func (Long) Copy

```go
func (l Long) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Long) Equal

```go
func (l Long) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Long) String

```go
func (l Long) String() string
```

#### func (Long) Type

```go
func (Long) Type() TagID
```
Type returns the TagID of the data.

#### type ReadError

```go
type ReadError struct {
	Where string
	Err   error
}
```

ReadError is an error returned when a read error occurs.

#### func (ReadError) Error

```go
func (r ReadError) Error() string
```

#### type Short

```go
type Short int16
```

Short is an implementation of the Data interface.

#### func (Short) Copy

```go
func (s Short) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Short) Equal

```go
func (s Short) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Short) String

```go
func (s Short) String() string
```

#### func (Short) Type

```go
func (Short) Type() TagID
```
Type returns the TagID of the data.

#### type String

```go
type String string
```

String is an implementation of the Data interface.

#### func (String) Copy

```go
func (s String) Copy() Data
```
Copy simply returns a copy of the data.

#### func (String) Equal

```go
func (s String) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (String) String

```go
func (s String) String() string
```

#### func (String) Type

```go
func (String) Type() TagID
```
Type returns the TagID of the data.

#### type Tag

```go
type Tag struct {
}
```

Tag is the main NBT type, a combination of a name and a Data type.

#### func  Decode

```go
func Decode(r io.Reader) (Tag, error)
```
Decode will encode a single tag from the reader using the default settings.

#### func  NewTag

```go
func NewTag(name string, d Data) Tag
```
NewTag constructs a new tag with the given name and data.

#### func (Tag) Copy

```go
func (t Tag) Copy() Tag
```
Copy simply returns a deep-copy of the tag.

#### func (Tag) Data

```go
func (t Tag) Data() Data
```
Data returns the tags data type.

#### func (Tag) Equal

```go
func (t Tag) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Tag) Name

```go
func (t Tag) Name() string
```
Name returns the tags' name.

#### func (Tag) String

```go
func (t Tag) String() string
```
String returns a textual representation of the tag.

#### func (Tag) TagID

```go
func (t Tag) TagID() TagID
```
TagID returns the type of the data.

#### type TagID

```go
type TagID uint8
```

TagID represents the type of nbt tag.

```go
const (
	TagEnd        TagID = 0
	TagByte       TagID = 1
	TagShort      TagID = 2
	TagInt        TagID = 3
	TagLong       TagID = 4
	TagFloat      TagID = 5
	TagDouble     TagID = 6
	TagByteArray  TagID = 7
	TagString     TagID = 8
	TagList       TagID = 9
	TagCompound   TagID = 10
	TagIntArray   TagID = 11
	TagBool       TagID = 12
	TagUint8      TagID = 13
	TagUint16     TagID = 14
	TagUint32     TagID = 15
	TagUint64     TagID = 16
	TagComplex64  TagID = 17
	TagComplex128 TagID = 18
)
```
Tag Types

#### func (TagID) String

```go
func (t TagID) String() string
```

#### type Uint16

```go
type Uint16 uint16
```

Uint16 is an implementation of the Data interface.

#### func (Uint16) Copy

```go
func (u Uint16) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Uint16) Equal

```go
func (u Uint16) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Uint16) String

```go
func (u Uint16) String() string
```

#### func (Uint16) Type

```go
func (Uint16) Type() TagID
```
Type returns the TagID of the data.

#### type Uint32

```go
type Uint32 uint32
```

Uint32 is an implementation of the Data interface.

#### func (Uint32) Copy

```go
func (u Uint32) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Uint32) Equal

```go
func (u Uint32) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Uint32) String

```go
func (u Uint32) String() string
```

#### func (Uint32) Type

```go
func (Uint32) Type() TagID
```
Type returns the TagID of the data.

#### type Uint64

```go
type Uint64 uint64
```

Uint64 is an implementation of the Data interface.

#### func (Uint64) Copy

```go
func (u Uint64) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Uint64) Equal

```go
func (u Uint64) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Uint64) String

```go
func (u Uint64) String() string
```

#### func (Uint64) Type

```go
func (Uint64) Type() TagID
```
Type returns the TagID of the data.

#### type Uint8

```go
type Uint8 uint8
```

Uint8 is an implementation of the Data interface.

#### func (Uint8) Copy

```go
func (u Uint8) Copy() Data
```
Copy simply returns a copy of the data.

#### func (Uint8) Equal

```go
func (u Uint8) Equal(e interface{}) bool
```
Equal satisfies the equaler.Equaler interface, allowing for types to be checked
for equality.

#### func (Uint8) String

```go
func (u Uint8) String() string
```

#### func (Uint8) Type

```go
func (Uint8) Type() TagID
```
Type returns the TagID of the data.

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

WriteError is an error returned when a write error occurs.

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
