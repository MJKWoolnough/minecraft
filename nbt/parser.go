// Package nbt implements a full Named Binary Tag reader/writer, based on the specs at
// http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt
package nbt

import (
	"fmt"
	"io"

	"github.com/MJKWoolnough/bytewrite"
	"github.com/MJKWoolnough/equaler"
)

type Data interface {
	io.ReaderFrom
	io.WriterTo
	equaler.Equaler
	CReadFrom(*Config, io.Reader) (int64, error)
	CWriteTo(*Config, io.Writer) (int64, error)
	Copy() Data
	String() string
	Type() TagId
}

type Tag struct {
	name String
	data Data
}

func ReadNBTFrom(f io.Reader) (*Tag, int64, error) {
	n := new(Tag)
	count, err := n.ReadFrom(f)
	return n, count, err

}

func NewTag(name string, d Data) (n *Tag) {
	m := Tag{
		name: String(name),
		data: d,
	}
	return &m
}

func (n *Tag) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *Tag) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var (
		c    int
		d    int64
		data [1]byte
	)
	c, err = io.ReadFull(f, data[:])
	total += int64(c)
	if err != nil {
		err = ReadError{"named TagId", err}
		return
	}
	tagType := TagId(data[0])
	if tagType == Tag_End {
		n.data = new(end)
	} else {
		if n.data, err = config.newFromTag(tagType); err != nil {
			return
		}
		d, err = n.name.CReadFrom(config, f)
		total += d
		if err != nil {
			err = ReadError{"name", err}
			return
		}
		d, err = n.data.CReadFrom(config, f)
		total += d
		if err != nil {
			if _, ok := err.(*ReadError); !ok {
				err = &ReadError{tagType.String(), err}
			}
		}
	}
	return
}

func (n *Tag) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n *Tag) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	var (
		c int
		d int64
	)
	tagType := n.data.Type()
	c, err = f.Write([]byte{byte(tagType)})
	total += int64(c)
	if err != nil {
		err = WriteError{"named TagId", err}
		return
	}
	if tagType == Tag_End {
		return
	}
	d, err = n.name.CWriteTo(config, f)
	total += d
	if err != nil {
		return
	}
	d, err = n.data.CWriteTo(config, f)
	total += d
	return
}

func (n *Tag) Copy() *Tag {
	return &Tag{
		n.name,
		n.data.Copy(),
	}
}

func (n *Tag) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Tag); ok {
		if n.data.Type() == m.data.Type() && n.name == m.name {
			return n.data.Equal(m.data)
		}
	}
	return false
}

func (n *Tag) Data() Data {
	return n.data
}

func (n *Tag) Name() string {
	return string(n.name)
}

func (n *Tag) TagId() TagId {
	return n.data.Type()
}

func (n *Tag) String() string {
	return fmt.Sprintf("%s(%q): %s", n.data.Type(), n.name, n.data)
}

type end struct{}

func (end) CReadFrom(*Config, io.Reader) (total int64, err error) {
	return
}

func (end) ReadFrom(io.Reader) (total int64, err error) {
	return
}

func (end) CWriteTo(*Config, io.Writer) (total int64, err error) {
	return
}

func (end) WriteTo(io.Writer) (total int64, err error) {
	return
}

func (end) Copy() Data {
	return &end{}
}

func (end) Equal(e equaler.Equaler) bool {
	_, ok := e.(*end)
	if !ok {
		_, ok = e.(end)
	}
	return ok
}

func (end) Type() TagId {
	return Tag_End
}

func (end) String() string {
	return ""
}

type Byte int8

func NewByte(d int8) *Byte {
	e := Byte(d)
	return &e
}

func (n *Byte) ReadFrom(f io.Reader) (total int64, err error) {
	var (
		c    int
		data [1]byte
	)
	c, err = io.ReadFull(f, data[:])
	total += int64(c)
	*n = Byte(data[0])
	return
}

func (n *Byte) CReadFrom(c *Config, f io.Reader) (total int64, err error) {
	return n.ReadFrom(f)
}

func (n Byte) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write([]byte{byte(n)})
	total += int64(c)
	return
}

func (n Byte) CWriteTo(c *Config, f io.Writer) (total int64, err error) {
	return n.WriteTo(f)
}

func (n Byte) Copy() Data {
	return &n
}

func (n *Byte) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Byte); ok {
		return *n == *m
	}
	return false
}

func (n Byte) String() string {
	return fmt.Sprintf("%d", n)
}

func (Byte) Type() TagId {
	return Tag_Byte
}

type Short int16

func NewShort(d int16) *Short {
	e := Short(d)
	return &e
}

func (n *Short) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *Short) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 2)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = Short(config.endian.Uint16(data))
	return
}

func (n Short) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n Short) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(config.endian.PutUint16(uint16(n)))
	total += int64(c)
	return
}

func (n Short) Copy() Data {
	return &n
}

func (n *Short) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Short); ok {
		return *n == *m
	}
	return false
}

func (n Short) String() string {
	return fmt.Sprintf("%d", n)
}

func (Short) Type() TagId {
	return Tag_Short
}

type Int int32

func NewInt(d int32) *Int {
	e := Int(d)
	return &e
}

func (n *Int) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *Int) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data[:])
	total += int64(c)
	*n = Int(config.endian.Uint32(data))
	return
}

func (n Int) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n Int) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(config.endian.PutUint32(uint32(n)))
	total += int64(c)
	return
}

func (n Int) Copy() Data {
	return &n
}

func (n *Int) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Int); ok {
		return *n == *m
	}
	return false
}

func (n Int) String() string {
	return fmt.Sprintf("%d", n)
}

func (Int) Type() TagId {
	return Tag_Int
}

type Long int64

func NewLong(d int64) *Long {
	e := Long(d)
	return &e
}

func (n *Long) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *Long) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 8)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = Long(config.endian.Uint64(data))
	return
}

func (n Long) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n Long) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(config.endian.PutUint64(uint64(n)))
	total += int64(c)
	return
}

func (n Long) Copy() Data {
	return &n
}

func (n *Long) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Long); ok {
		return *n == *m
	}
	return false
}

func (n Long) String() string {
	return fmt.Sprintf("%d", n)
}

func (Long) Type() TagId {
	return Tag_Long
}

type Float float32

func NewFloat(d float32) *Float {
	e := Float(d)
	return &e
}

func (n *Float) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *Float) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = Float(config.endian.Float32(data))
	return
}

func (n Float) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n Float) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(config.endian.PutFloat32(float32(n)))
	total += int64(c)
	return
}

func (n Float) Copy() Data {
	return &n
}

func (n *Float) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Float); ok {
		return *n == *m
	}
	return false
}

func (n Float) String() string {
	return fmt.Sprintf("%f", n)
}

func (Float) Type() TagId {
	return Tag_Float
}

type Double float64

func NewDouble(d float64) *Double {
	e := Double(d)
	return &e
}

func (n *Double) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *Double) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 8)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = Double(config.endian.Float64(data))
	return
}

func (n Double) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n Double) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(config.endian.PutFloat64(float64(n)))
	total += int64(c)
	return
}

func (n Double) Copy() Data {
	return &n
}

func (n *Double) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Double); ok {
		return *n == *m
	}
	return false
}

func (n Double) String() string {
	return fmt.Sprintf("%f", n)
}

func (Double) Type() TagId {
	return Tag_Double
}

type ByteArray []int8

func NewByteArray(d []int8) *ByteArray {
	e := ByteArray(d)
	return &e
}

func (n *ByteArray) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *ByteArray) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	length := config.endian.Uint32(data)
	data = make([]byte, length)
	*n = ByteArray(make([]int8, length))
	c, err = io.ReadFull(f, data)
	total += int64(c)
	for i := uint32(0); i < length; i++ {
		(*n)[i] = int8(data[i])
	}
	return
}

func (n ByteArray) WriteTo(f io.Writer) (int64, error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n ByteArray) CWriteTo(config *Config, f io.Writer) (int64, error) {
	data := make([]byte, len(n)+4)
	copy(data, config.endian.PutUint32(uint32(len(n))))
	for i, b := range n {
		data[i+4] = byte(b)
	}
	c, err := f.Write(data)
	return int64(c), err
}

func (n ByteArray) Copy() Data {
	c := ByteArray(make([]int8, len(n)))
	copy(c, n)
	return &c
}

func (n *ByteArray) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*ByteArray); ok {
		if len(*n) == len(*m) {
			for i, o := range *n {
				if o != (*m)[i] {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (n ByteArray) String() string {
	return fmt.Sprintf("[%d bytes] %v", len(n), []int8(n))
}

func (ByteArray) Type() TagId {
	return Tag_ByteArray
}

type String string

func NewString(d string) *String {
	e := String(d)
	return &e
}

func (n *String) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *String) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 2)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	data = make([]byte, config.endian.Uint16(data))
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = String(data)
	return
}

func (n String) WriteTo(f io.Writer) (int64, error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n String) CWriteTo(config *Config, f io.Writer) (int64, error) {
	c, err := f.Write(append(config.endian.PutUint16(uint16(len(n))), n...))
	return int64(c), err
}

func (n String) Copy() Data {
	return &n
}

func (n *String) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*String); ok {
		return *n == *m
	}
	return false
}

func (n String) String() string {
	return string(n)
}

func (String) Type() TagId {
	return Tag_String
}

type List struct {
	tagType TagId
	data    []Data
}

func NewList(data []Data) *List {
	if len(data) == 0 {
		return &List{Tag_Byte, data}
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

func NewEmptyList(tagType TagId) *List {
	return &List{
		tagType,
		make([]Data, 0),
	}
}

func (n *List) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *List) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var (
		c int
		d int64
	)
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data[:1])
	total += int64(c)
	if err != nil {
		err = &ReadError{"list TagId", err}
		return
	}
	n.tagType = TagId(data[0])
	c, err = io.ReadFull(f, data)
	total += int64(c)
	if err != nil {
		err = &ReadError{"list length", err}
		return
	}
	length := config.endian.Uint32(data)
	n.data = make([]Data, length)
	for i := uint32(0); i < length; i++ {
		if n.data[i], err = config.newFromTag(n.tagType); err != nil {
			return
		}
		d, err = n.data[i].CReadFrom(config, f)
		total += d
		if err != nil {
			return
		}
	}
	return
}

func (n *List) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n *List) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	var (
		c int
		d int64
	)
	c, err = f.Write([]byte{byte(n.tagType)})
	total += int64(c)
	if err != nil {
		return
	}
	c, err = f.Write(config.endian.PutUint32(uint32(len(n.data))))
	total += int64(c)
	if err != nil {
		return
	}
	if n.tagType != Tag_End {
		for _, data := range n.data {
			if tagId := data.Type(); tagId != n.tagType {
				err = &WrongTag{n.tagType, tagId}
				break
			}
			d, err = data.CWriteTo(config, f)
			total += d
			if err != nil {
				break
			}
		}
	}
	return
}

func (n *List) TagType() TagId {
	return n.tagType
}

func (n *List) Copy() Data {
	c := new(List)
	c.tagType = n.tagType
	c.data = make([]Data, len(n.data))
	for i, d := range n.data {
		c.data[i] = d.Copy()
	}
	return c
}

func (n *List) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*List); ok {
		if n.tagType == m.tagType && len(n.data) == len(m.data) {
			for i, o := range n.data {
				if !o.Equal(m.data[i]) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (n *List) String() string {
	s := fmt.Sprintf("%d entries of type %s {", len(n.data), n.tagType)
	for _, d := range n.data {
		s += fmt.Sprintf("\n	%s: %s", n.tagType, indent(d.String()))
	}
	return s + "\n}"
}

func (n *List) Set(i int32, data Data) error {
	if i < 0 || i >= int32(len(n.data)) {
		return &BadRange{}
	}
	if err := n.valid(data); err != nil {
		return err
	}
	n.data[i] = data
	return nil
}

func (n *List) Get(i int) Data {
	if i >= 0 && i < len(n.data) {
		return n.data[i]
	}
	return nil
}

func (n *List) Append(data ...Data) error {
	if err := n.valid(data...); err != nil {
		return err
	}
	n.data = append(n.data, data...)
	return nil
}

func (n *List) Insert(i int, data ...Data) error {
	if err := n.valid(data...); err != nil {
		return err
	}
	n.data = append(n.data[:i], append(data, n.data[i:]...)...)
	return nil
}

func (n *List) Remove(i int) {
	if i >= 0 && i < len(n.data) {
		copy(n.data[i:], n.data[i+1:])
		n.data[len(n.data)-1] = nil
		n.data = n.data[:len(n.data)-1]
	}
}

func (n *List) Len() int {
	return len(n.data)
}

func (n *List) valid(data ...Data) error {
	for _, d := range data {
		if t := d.Type(); t != n.tagType {
			return &WrongTag{n.tagType, t}
		}
	}
	return nil
}

func (List) Type() TagId {
	return Tag_List
}

type Compound []*Tag

func NewCompound(d Compound) *Compound {
	e := Compound(d)
	return &e
}

func (n *Compound) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *Compound) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var d int64
	*n = make(Compound, 0)
	for {
		data := new(Tag)
		d, err = data.CReadFrom(config, f)
		total += d
		if err != nil {
			return
		}
		if data.TagId() == Tag_End {
			break
		}
		*n = append(*n, data)
	}
	return
}

func (n Compound) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n Compound) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	var (
		c int
		d int64
	)
	for _, data := range n {
		d, err = data.CWriteTo(config, f)
		total += d
		if err != nil {
			return
		}
		if data.TagId() == Tag_End {
			return
		}
	}
	c, err = f.Write([]byte{byte(Tag_End)})
	total += int64(c)
	return
}

func (n Compound) Copy() Data {
	c := make(Compound, len(n))
	for i, d := range n {
		c[i] = d.Copy()
	}
	return &c
}

func (n *Compound) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Compound); ok {
		if len(*n) == len(*m) {
			for _, o := range *n {
				if p := m.Get(o.Name()); p == nil || !p.Equal(o) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (n Compound) String() string {
	s := fmt.Sprintf("%d entries {", len(n))
	for _, d := range n {
		s += "\n	" + indent(d.String())
	}
	return s + "\n}"
}

func (n Compound) Get(name string) *Tag {
	for _, t := range n {
		if t.Name() == name {
			return t
		}
	}
	return nil
}

func (n *Compound) Remove(name string) {
	for i, t := range *n {
		if t.Name() == name {
			copy((*n)[i:], (*n)[i+1:])
			(*n)[len((*n))-1] = nil
			(*n) = (*n)[:len((*n))-1]
			return
		}
	}
}

func (n *Compound) Set(tag *Tag) {
	if tag.TagId() == Tag_End {
		return
	}
	name := tag.Name()
	for i, t := range *n {
		if t.Name() == name {
			(*n)[i] = tag
			return
		}
	}
	*n = append(*n, tag)
}

func (Compound) Type() TagId {
	return Tag_Compound
}

type IntArray []int32

func NewIntArray(d []int32) *IntArray {
	e := IntArray(d)
	return &e
}

func (n *IntArray) ReadFrom(f io.Reader) (total int64, err error) {
	return n.CReadFrom(defaultConfig, f)
}

func (n *IntArray) CReadFrom(config *Config, f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	if err != nil {
		return
	}
	length := config.endian.Uint32(data)
	*n = make([]int32, length)
	ints := make([]byte, 4*length)
	c, err = io.ReadFull(f, ints)
	total += int64(c)
	for i := uint32(0); i < length; i++ {
		(*n)[i] = int32(config.endian.Uint32(ints))
		ints = ints[4:]
	}
	return
}

func (n IntArray) WriteTo(f io.Writer) (total int64, err error) {
	return n.CWriteTo(defaultConfig, f)
}

func (n IntArray) CWriteTo(config *Config, f io.Writer) (total int64, err error) {
	ints := make([]byte, 4, 4*len(n)+4)
	copy(ints, config.endian.PutUint32(uint32(len(n))))
	for i := 0; i < len(n); i++ {
		ints = append(ints, bytewrite.BigEndian.PutUint32(uint32(n[i]))...)
	}
	c, err := f.Write(ints)
	return int64(c), err
}

func (n IntArray) Copy() Data {
	c := IntArray(make([]int32, len(n)))
	copy(c, n)
	return &c
}

func (n *IntArray) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*IntArray); ok {
		if len(*n) == len(*m) {
			for i, o := range *n {
				if o != (*m)[i] {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (n IntArray) String() string {
	return fmt.Sprintf("[%d ints] %v", len(n), []int32(n))
}

func (IntArray) Type() TagId {
	return Tag_IntArray
}
