// Package nbt implements a full Named Binary Tag reader/writer, based on the specs at
// http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt
package nbt

import "strconv"

// Tag Types
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

var tagIDNames = [...]string{
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

// TagID represents the type of nbt tag
type TagID uint8

func (t TagID) String() string {
	if int(t) < len(tagIDNames) {
		return tagIDNames[t]
	}
	return ""
}

// Data is an interface representing the many different types that a tag can be
type Data interface {
	Equal(interface{}) bool
	Copy() Data
	String() string
	Type() TagID
}

// Tag is the main NBT type, a combination of a name and a Data type
type Tag struct {
	name string
	data Data
}

// NewTag constructs a new tag with the given name and data.
func NewTag(name string, d Data) Tag {
	return Tag{
		name: name,
		data: d,
	}
}

// Copy simply returns a deep-copy the the tag
func (t Tag) Copy() Tag {
	return Tag{
		t.name,
		t.data.Copy(),
	}
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (t Tag) Equal(e interface{}) bool {
	if m, ok := e.(Tag); ok {
		if t.data.Type() == m.data.Type() && t.name == m.name {
			return t.data.Equal(m.data)
		}
	}
	return false
}

// Data returns the tags data type
func (t Tag) Data() Data {
	if t.data == nil {
		return end{}
	}
	return t.data
}

// Name returns the tags' name
func (t Tag) Name() string {
	return t.name
}

// TagID returns the type of the data
func (t Tag) TagID() TagID {
	if t.data == nil {
		return TagEnd
	}
	return t.data.Type()
}

// String returns a textual representation of the tag
func (t Tag) String() string {
	return t.data.Type().String() + "(" + strconv.Quote(t.name) + "): " + t.data.String()
}

type end struct{}

func (end) Copy() Data {
	return &end{}
}

func (end) Equal(e interface{}) bool {
	_, ok := e.(end)
	return ok
}

func (end) Type() TagID {
	return TagEnd
}

func (end) String() string {
	return ""
}

// Byte is an implementation of the Data interface
// NB: Despite being an unsigned integer, it is still called a byte to match
// the spec.
type Byte int8

// Copy simply returns a copy the the data
func (b Byte) Copy() Data {
	return b
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (b Byte) Equal(e interface{}) bool {
	if m, ok := e.(Byte); ok {
		return b == m
	}
	return false
}

func (b Byte) String() string {
	return strconv.FormatInt(int64(b), 10)
}

// Type returns the TagID of the data
func (Byte) Type() TagID {
	return TagByte
}

// Short is an implementation of the Data interface
type Short int16

// Copy simply returns a copy the the data
func (s Short) Copy() Data {
	return s
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (s Short) Equal(e interface{}) bool {
	if m, ok := e.(Short); ok {
		return s == m
	}
	return false
}

func (s Short) String() string {
	return strconv.FormatInt(int64(s), 10)
}

// Type returns the TagID of the data
func (Short) Type() TagID {
	return TagShort
}

// Int is an implementation of the Data interface
type Int int32

// Copy simply returns a copy the the data
func (i Int) Copy() Data {
	return i
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (i Int) Equal(e interface{}) bool {
	if m, ok := e.(Int); ok {
		return i == m
	}
	return false
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Type returns the TagID of the data
func (Int) Type() TagID {
	return TagInt
}

// Long is an implementation of the Data interface
type Long int64

// Copy simply returns a copy the the data
func (l Long) Copy() Data {
	return l
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (l Long) Equal(e interface{}) bool {
	if m, ok := e.(Long); ok {
		return l == m
	}
	return false
}

func (l Long) String() string {
	return strconv.FormatInt(int64(l), 10)
}

// Type returns the TagID of the data
func (Long) Type() TagID {
	return TagLong
}

// Float is an implementation of the Data interface
type Float float32

// Copy simply returns a copy the the data
func (f Float) Copy() Data {
	return f
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (f Float) Equal(e interface{}) bool {
	if m, ok := e.(Float); ok {
		return f == m
	}
	return false
}

func (f Float) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 32)
}

// Type returns the TagID of the data
func (Float) Type() TagID {
	return TagFloat
}

// Double is an implementation of the Data interface
type Double float64

// Copy simply returns a copy the the data
func (d Double) Copy() Data {
	return d
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (d Double) Equal(e interface{}) bool {
	if m, ok := e.(Double); ok {
		return d == m
	}
	return false
}

func (d Double) String() string {
	return strconv.FormatFloat(float64(d), 'g', -1, 64)
}

// Type returns the TagID of the data
func (Double) Type() TagID {
	return TagDouble
}

// ByteArray is an implementation of the Data interface
type ByteArray []int8

// Copy simply returns a copy the the data
func (b ByteArray) Copy() Data {
	c := make(ByteArray, len(b))
	copy(c, b)
	return c
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (b ByteArray) Equal(e interface{}) bool {
	if m, ok := e.(ByteArray); ok {
		for i := 0; i < len(b); i++ {
			if b[i] != m[i] {
				return false
			}
		}
		return true
	}
	return false
}

func (b ByteArray) String() string {
	var data []byte
	for n, d := range b {
		if n > 0 {
			data = append(data, ',', ' ')
		}
		data = append(data, strconv.FormatInt(int64(d), 10)...)
	}
	return "[" + strconv.FormatInt(int64(len(b)), 10) + " bytes] [" + string(data) + "]"
}

// Type returns the TagID of the data
func (ByteArray) Type() TagID {
	return TagByteArray
}

// Bytes converts the ByteArray (actually int8) to an actual slice of bytes.
// NB: May uss unsafe, so the underlying array may be the same.
func (b ByteArray) Bytes() []byte {
	return byteArrayToByteSlice(b)
}

// String is an implementation of the Data interface
type String string

// Copy simply returns a copy the the data
func (s String) Copy() Data {
	return s
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (s String) Equal(e interface{}) bool {
	if m, ok := e.(String); ok {
		return s == m
	}
	return false
}

func (s String) String() string {
	return string(s)
}

// Type returns the TagID of the data
func (String) Type() TagID {
	return TagString
}

// List is an implementation of the Data interface
type List struct {
	tagType TagID
	data    []Data
}

// NewList returns a new List Data type, or nil if the given data is not of all
// the same Data type
func NewList(data []Data) *List {
	if len(data) == 0 {
		return &List{TagByte, data}
	}
	tagType := data[0].Type()
	for i := 1; i < len(data); i++ {
		if id := data[i].Type(); id != tagType {
			return nil
		}
	}
	return &List{
		tagType,
		data,
	}
}

// NewEmptyList returns a new empty List Data type, set to the type given
func NewEmptyList(tagType TagID) *List {
	return &List{
		tagType,
		make([]Data, 0),
	}
}

// TagType returns the TagID of the type of tag this list contains
func (l *List) TagType() TagID {
	return l.tagType
}

// Copy simply returns a deep-copy the the data
func (l *List) Copy() Data {
	c := new(List)
	c.tagType = l.tagType
	c.data = make([]Data, len(l.data))
	for i, d := range l.data {
		c.data[i] = d.Copy()
	}
	return c
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (l *List) Equal(e interface{}) bool {
	if m, ok := e.(*List); ok {
		if l.tagType == m.tagType && len(l.data) == len(m.data) {
			for i, o := range l.data {
				if !o.Equal(m.data[i]) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (l *List) String() string {
	s := strconv.FormatInt(int64(len(l.data)), 10) + " entries of type " + l.tagType.String() + " {"
	for _, d := range l.data {
		s += "\n	" + l.tagType.String() + ": " + indent(d.String())
	}
	return s + "\n}"
}

// Set sets the data at the given position. It does not append
func (l *List) Set(i int32, data Data) error {
	if i < 0 || i >= int32(len(l.data)) {
		return ErrBadRange
	}
	if err := l.valid(data); err != nil {
		return err
	}
	l.data[i] = data
	return nil
}

// Get returns the data at the given positon
func (l *List) Get(i int) Data {
	if i >= 0 && i < len(l.data) {
		return l.data[i]
	}
	return nil
}

// Append adds data to the list
func (l *List) Append(data ...Data) error {
	if err := l.valid(data...); err != nil {
		return err
	}
	l.data = append(l.data, data...)
	return nil
}

// Insert will add the given data at the specified position, moving other
// elements up.
func (l *List) Insert(i int, data ...Data) error {
	if err := l.valid(data...); err != nil {
		return err
	}
	l.data = append(l.data[:i], append(data, l.data[i:]...)...)
	return nil
}

// Remove deletes the specified position and shifts remaing data down
func (l *List) Remove(i int) {
	if i >= 0 && i < len(l.data) {
		copy(l.data[i:], l.data[i+1:])
		l.data[len(l.data)-1] = nil
		l.data = l.data[:len(l.data)-1]
	}
}

// Len returns the length of the list
func (l *List) Len() int {
	return len(l.data)
}

func (l *List) valid(data ...Data) error {
	for _, d := range data {
		if t := d.Type(); t != l.tagType {
			return &WrongTag{l.tagType, t}
		}
	}
	return nil
}

// Type returns the TagID of the data
func (List) Type() TagID {
	return TagList
}

// Compound is an implementation of the Data interface
type Compound []Tag

// Copy simply returns a deep-copy the the data
func (c Compound) Copy() Data {
	d := make(Compound, len(c))
	for i, t := range c {
		d[i] = t.Copy()
	}
	return d
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (c Compound) Equal(e interface{}) bool {
	if m, ok := e.(Compound); ok {
		if len(c) == len(m) {
			for _, o := range c {
				if p := m.Get(o.Name()); !p.Equal(o) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (c Compound) String() string {
	s := strconv.FormatInt(int64(len(c)), 10) + " entries {"
	for _, d := range c {
		s += "\n	" + indent(d.String())
	}
	return s + "\n}"
}

// Get returns the tag for the given name
func (c Compound) Get(name string) Tag {
	for _, t := range c {
		if t.Name() == name {
			return t
		}
	}
	return Tag{}
}

// Remove removes the tag corresponding to the given name
func (c *Compound) Remove(name string) {
	for i, t := range *c {
		if t.Name() == name {
			copy((*c)[i:], (*c)[i+1:])
			(*c)[len((*c))-1] = Tag{data: end{}}
			(*c) = (*c)[:len((*c))-1]
			return
		}
	}
}

// Set adds the given tag to the compound, or, if the tags name is already
// present, overrides the old data
func (c *Compound) Set(tag Tag) {
	if tag.TagID() == TagEnd {
		return
	}
	name := tag.Name()
	for i, t := range *c {
		if t.Name() == name {
			(*c)[i] = tag
			return
		}
	}
	*c = append(*c, tag)
}

// Type returns the TagID of the data
func (Compound) Type() TagID {
	return TagCompound
}

// IntArray is an implementation of the Data interface
type IntArray []int32

// Copy simply returns a copy the the data
func (i IntArray) Copy() Data {
	c := make(IntArray, len(i))
	copy(c, i)
	return c
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (i IntArray) Equal(e interface{}) bool {
	if m, ok := e.(IntArray); ok {
		if len(i) == len(m) {
			for j, o := range i {
				if o != m[j] {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (i IntArray) String() string {
	var data []byte
	for n, d := range i {
		if n > 0 {
			data = append(data, ',', ' ')
		}
		data = append(data, strconv.FormatInt(int64(d), 10)...)
	}
	return "[" + strconv.FormatInt(int64(len(i)), 10) + " bytes] [" + string(data) + "]"
}

// Type returns the TagID of the data
func (IntArray) Type() TagID {
	return TagIntArray
}

// Bool is an implementation of the Data interface
type Bool bool

// Copy simply returns a copy the the data
func (b Bool) Copy() Data {
	return b
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (b Bool) Equal(e interface{}) bool {
	if m, ok := e.(Bool); ok {
		return b == m
	}
	return false
}

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

// Type returns the TagID of the data
func (Bool) Type() TagID {
	return TagBool
}

// Uint8 is an implementation of the Data interface
type Uint8 uint8

// Copy simply returns a copy the the data
func (u Uint8) Copy() Data {
	return u
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (u Uint8) Equal(e interface{}) bool {
	if m, ok := e.(Uint8); ok {
		return u == m
	}
	return false
}

func (u Uint8) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

// Type returns the TagID of the data
func (Uint8) Type() TagID {
	return TagUint8
}

// Uint16 is an implementation of the Data interface
type Uint16 uint16

// Copy simply returns a copy the the data
func (u Uint16) Copy() Data {
	return u
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (u Uint16) Equal(e interface{}) bool {
	if m, ok := e.(Uint16); ok {
		return u == m
	}
	return false
}

func (u Uint16) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

// Type returns the TagID of the data
func (Uint16) Type() TagID {
	return TagUint16
}

// Uint32 is an implementation of the Data interface
type Uint32 uint32

// Copy simply returns a copy the the data
func (u Uint32) Copy() Data {
	return u
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (u Uint32) Equal(e interface{}) bool {
	if m, ok := e.(Uint32); ok {
		return u == m
	}
	return false
}

func (u Uint32) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

// Type returns the TagID of the data
func (Uint32) Type() TagID {
	return TagUint32
}

// Uint64 is an implementation of the Data interface
type Uint64 uint64

// Copy simply returns a copy the the data
func (u Uint64) Copy() Data {
	return u
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (u Uint64) Equal(e interface{}) bool {
	if m, ok := e.(Uint64); ok {
		return u == m
	}
	return false
}

func (u Uint64) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

// Type returns the TagID of the data
func (Uint64) Type() TagID {
	return TagUint64
}

// Complex64 is an implementation of the Data interface
type Complex64 complex64

// Copy simply returns a copy the the data
func (c Complex64) Copy() Data {
	return c
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (c Complex64) Equal(e interface{}) bool {
	if m, ok := e.(Complex64); ok {
		return c == m
	}
	return false
}

func (c Complex64) String() string {
	var str string
	if real(c) != 0 {
		str = strconv.FormatFloat(float64(real(c)), 'g', -1, 32)
	}
	if imag(c) != 0 {
		if len(str) != 0 {
			str += "+"
		}
		str += strconv.FormatFloat(float64(imag(c)), 'g', -1, 32) + "i"
	}
	return str
}

// Type returns the TagID of the data
func (Complex64) Type() TagID {
	return TagComplex64
}

// Complex128 is an implementation of the Data interface
type Complex128 complex128

// Copy simply returns a copy the the data
func (c Complex128) Copy() Data {
	return c
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (c Complex128) Equal(e interface{}) bool {
	if m, ok := e.(Complex128); ok {
		return c == m
	}
	return false
}

func (c Complex128) String() string {
	var str string
	if real(c) != 0 {
		str = strconv.FormatFloat(float64(real(c)), 'g', -1, 64)
	}
	if imag(c) != 0 {
		if len(str) != 0 {
			str += "+"
		}
		str += strconv.FormatFloat(float64(imag(c)), 'g', -1, 64) + "i"
	}
	return str
}

// Type returns the TagID of the data
func (Complex128) Type() TagID {
	return TagComplex128
}
