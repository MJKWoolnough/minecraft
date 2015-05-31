// Package nbt implements a full Named Binary Tag reader/writer, based on the specs at
// http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt
package nbt

import (
	"fmt"

	"github.com/MJKWoolnough/equaler"
)

// Tag Types
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
	equaler.Equaler
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
func (t Tag) Equal(e equaler.Equaler) bool {
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
	return fmt.Sprintf("%s(%q): %s", t.data.Type(), t.name, t.data)
}

type end struct{}

func (end) Copy() Data {
	return &end{}
}

func (end) Equal(e equaler.Equaler) bool {
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
type Byte int8

// Copy simply returns a copy the the data
func (b Byte) Copy() Data {
	return b
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (b Byte) Equal(e equaler.Equaler) bool {
	if m, ok := e.(Byte); ok {
		return b == m
	}
	return false
}

func (b Byte) String() string {
	return fmt.Sprintf("%d", b)
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
func (s Short) Equal(e equaler.Equaler) bool {
	if m, ok := e.(Short); ok {
		return s == m
	}
	return false
}

func (s Short) String() string {
	return fmt.Sprintf("%d", s)
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
func (i Int) Equal(e equaler.Equaler) bool {
	if m, ok := e.(Int); ok {
		return i == m
	}
	return false
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
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
func (l Long) Equal(e equaler.Equaler) bool {
	if m, ok := e.(Long); ok {
		return l == m
	}
	return false
}

func (l Long) String() string {
	return fmt.Sprintf("%d", l)
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
func (f Float) Equal(e equaler.Equaler) bool {
	if m, ok := e.(Float); ok {
		return f == m
	}
	return false
}

func (f Float) String() string {
	return fmt.Sprintf("%f", f)
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
func (d Double) Equal(e equaler.Equaler) bool {
	if m, ok := e.(Double); ok {
		return d == m
	}
	return false
}

func (d Double) String() string {
	return fmt.Sprintf("%f", d)
}

// Type returns the TagID of the data
func (Double) Type() TagID {
	return TagDouble
}

// ByteArray is an implementation of the Data interface
type ByteArray []int8

// Copy simply returns a copy the the data
func (b ByteArray) Copy() Data {
	c := ByteArray(make([]int8, len(b)))
	copy(c, b)
	return c
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (b ByteArray) Equal(e equaler.Equaler) bool {
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
	return fmt.Sprintf("[%d bytes] %v", len(b), []int8(b))
}

// Type returns the TagID of the data
func (ByteArray) Type() TagID {
	return TagByteArray
}

// String is an implementation of the Data interface
type String string

// Copy simply returns a copy the the data
func (s String) Copy() Data {
	return s
}

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality
func (s String) Equal(e equaler.Equaler) bool {
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
func (l *List) Equal(e equaler.Equaler) bool {
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
	s := fmt.Sprintf("%d entries of type %s {", len(l.data), l.tagType)
	for _, d := range l.data {
		s += fmt.Sprintf("\n	%s: %s", l.tagType, indent(d.String()))
	}
	return s + "\n}"
}

// Set sets the data at the given position. It does not append
func (l *List) Set(i int32, data Data) error {
	if i < 0 || i >= int32(len(l.data)) {
		return &BadRange{}
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
func (c Compound) Equal(e equaler.Equaler) bool {
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
	s := fmt.Sprintf("%d entries {", len(c))
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
func (i IntArray) Equal(e equaler.Equaler) bool {
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
	return fmt.Sprintf("[%d ints] %v", len(i), []int32(i))
}

// Type returns the TagID of the data
func (IntArray) Type() TagID {
	return TagIntArray
}
