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

// Package nbt implements a full Named Binary Tag reader/writer, based on the specs at
// http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt
package nbt

import (
	"fmt"
	"github.com/MJKWoolnough/bytewrite"
	"github.com/MJKWoolnough/equaler"
	"io"
)

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

type Data interface {
	io.ReaderFrom
	io.WriterTo
	equaler.Equaler
	Copy() Data
	String() string
}

type TagId uint8

func (t TagId) String() string {
	if int(t) < len(tagIdNames) {
		return tagIdNames[t]
	}
	return ""
}

type Tag struct {
	tagType TagId
	name    String
	d       Data
}

func ReadNBTFrom(f io.Reader) (*Tag, int64, error) {
	n := new(Tag)
	count, err := n.ReadFrom(f)
	return n, count, err

}

func NewTag(name string, d Data) (n *Tag) {
	tagType, err := idFromData(d)
	if err != nil {
		return nil
	}
	m := Tag{
		tagType,
		String(name),
		d,
	}
	return &m
}

func (n *Tag) ReadFrom(f io.Reader) (total int64, err error) {
	var (
		c    int
		d    int64
		data [1]byte
	)
	c, err = io.ReadFull(f, data[:])
	total += int64(c)
	if err != nil {
		err = &ReadError{"named TagId", err}
		return
	}
	n.tagType = TagId(data[0])
	if n.tagType == Tag_End {
		n.d = new(end)
	} else {
		if n.d, err = newFromTag(n.tagType); err != nil {
			return
		}
		d, err = n.name.ReadFrom(f)
		total += d
		if err != nil {
			err = &ReadError{"name", err}
			return
		}
		d, err = n.d.ReadFrom(f)
		total += d
		if err != nil {
			if _, ok := err.(*ReadError); !ok {
				err = &ReadError{n.tagType.String(), err}
			}
		}
	}
	return
}

func (n *Tag) WriteTo(f io.Writer) (total int64, err error) {
	var (
		c int
		d int64
	)
	c, err = f.Write([]byte{byte(n.tagType)})
	total += int64(c)
	if err != nil {
		err = &WriteError{"named TagId", err}
		return
	}
	if n.tagType == Tag_End {
		return
	}
	d, err = n.name.WriteTo(f)
	total += d
	if err != nil {
		return
	}
	d, err = n.d.WriteTo(f)
	total += d
	return
}

func (n *Tag) Copy() *Tag {
	return &Tag{
		n.tagType,
		n.name,
		n.d.Copy(),
	}
}

func (n *Tag) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Tag); ok {
		if n.tagType == m.tagType && n.name == m.name {
			return n.d.Equal(m.d)
		}
	}
	return false
}

func (n *Tag) Data() Data {
	return n.d
}

func (n *Tag) Name() string {
	return string(n.name)
}

func (n *Tag) TagId() TagId {
	return n.tagType
}

func (n *Tag) String() string {
	return fmt.Sprintf("%s(%q): %s", n.tagType, n.name, n.d)
}

type end struct{}

func (end) ReadFrom(f io.Reader) (total int64, err error) {
	return
}

func (end) WriteTo(f io.Writer) (total int64, err error) {
	return
}

func (end) Copy() Data {
	return &end{}
}

func (end) Equal(e equaler.Equaler) bool {
	_, ok := e.(*end)
	return ok
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

func (n *Byte) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write([]byte{byte(*n)})
	total += int64(c)
	return
}

func (n Byte) Copy() Data {
	return &n
}

func (n Byte) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Byte); ok {
		return n == *m
	}
	return false
}

func (n Byte) String() string {
	return fmt.Sprintf("%d", n)
}

type Short int16

func NewShort(d int16) *Short {
	e := Short(d)
	return &e
}

func (n *Short) ReadFrom(f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 2)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = Short(bytewrite.BigEndian.Uint16(data))
	return
}

func (n *Short) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(bytewrite.BigEndian.PutUint16(uint16(*n)))
	total += int64(c)
	return
}

func (n Short) Copy() Data {
	return &n
}

func (n Short) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Short); ok {
		return n == *m
	}
	return false
}

func (n Short) String() string {
	return fmt.Sprintf("%d", n)
}

type Int int32

func NewInt(d int32) *Int {
	e := Int(d)
	return &e
}

func (n *Int) ReadFrom(f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data[:])
	total += int64(c)
	*n = Int(bytewrite.BigEndian.Uint32(data))
	return
}

func (n *Int) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(bytewrite.BigEndian.PutUint32(uint32(*n)))
	total += int64(c)
	return
}

func (n Int) Copy() Data {
	return &n
}

func (n Int) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Int); ok {
		return n == *m
	}
	return false
}

func (n Int) String() string {
	return fmt.Sprintf("%d", n)
}

type Long int64

func NewLong(d int64) *Long {
	e := Long(d)
	return &e
}

func (n *Long) ReadFrom(f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 8)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = Long(bytewrite.BigEndian.Uint64(data))
	return
}

func (n *Long) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(bytewrite.BigEndian.PutUint64(uint64(*n)))
	total += int64(c)
	return
}

func (n Long) Copy() Data {
	return &n
}

func (n Long) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Long); ok {
		return n == *m
	}
	return false
}

func (n Long) String() string {
	return fmt.Sprintf("%d", n)
}

type Float float32

func NewFloat(d float32) *Float {
	e := Float(d)
	return &e
}

func (n *Float) ReadFrom(f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = Float(bytewrite.BigEndian.Float32(data))
	return
}

func (n *Float) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(bytewrite.BigEndian.PutFloat32(float32(*n)))
	total += int64(c)
	return
}

func (n Float) Copy() Data {
	return &n
}

func (n Float) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Float); ok {
		return n == *m
	}
	return false
}

func (n Float) String() string {
	return fmt.Sprintf("%f", n)
}

type Double float64

func NewDouble(d float64) *Double {
	e := Double(d)
	return &e
}

func (n *Double) ReadFrom(f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 8)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	*n = Double(bytewrite.BigEndian.Float64(data))
	return
}

func (n *Double) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(bytewrite.BigEndian.PutFloat64(float64(*n)))
	total += int64(c)
	return
}

func (n Double) Copy() Data {
	return &n
}

func (n Double) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Double); ok {
		return n == *m
	}
	return false
}

func (n Double) String() string {
	return fmt.Sprintf("%f", n)
}

type ByteArray []int8

func NewByteArray(d []int8) *ByteArray {
	e := ByteArray(d)
	return &e
}

func (n *ByteArray) ReadFrom(f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	length := bytewrite.BigEndian.Uint32(data)
	bData := make([]byte, length)
	iData := ByteArray(make([]int8, length))
	c, err = io.ReadFull(f, bData)
	total += int64(c)
	for i := uint32(0); i < length; i++ {
		iData[i] = int8(bData[i])
	}
	*n = iData
	return
}

func (n *ByteArray) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(bytewrite.BigEndian.PutUint32(uint32(len(*n))))
	total += int64(c)
	if err != nil {
		return
	}
	data := make([]byte, len(*n))
	for i := 0; i < len(data); i++ {
		data[i] = byte((*n)[i])
	}
	c, err = f.Write(data)
	total += int64(c)
	return
}

func (n ByteArray) Copy() Data {
	c := ByteArray(make([]int8, len(n)))
	copy(c, n)
	return &c
}

func (n ByteArray) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*ByteArray); ok {
		if len(n) == len(*m) {
			for i, o := range n {
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

type String string

func NewString(d string) *String {
	e := String(d)
	return &e
}

func (n *String) ReadFrom(f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 2)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	bData := make([]byte, bytewrite.BigEndian.Uint16(data))
	c, err = io.ReadFull(f, bData)
	total += int64(c)
	*n = String(bData)
	return
}

func (n *String) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(bytewrite.BigEndian.PutUint16(uint16(len(*n))))
	total += int64(c)
	if err != nil {
		return
	}
	c, err = f.Write([]byte(*n))
	total += int64(c)
	return
}

func (n String) Copy() Data {
	return &n
}

func (n String) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*String); ok {
		return n == *m
	}
	return false
}

func (n String) String() string {
	return string(n)
}

type List struct {
	tagType TagId
	d       []Data
}

func NewList(d []Data) *List {
	if len(d) == 0 {
		return &List{Tag_Byte, d}
	}
	tagType, err := idFromData(d[0])
	if err != nil {
		return nil
	}
	for i := 1; i < len(d); i++ {
		if id, _ := idFromData(d[i]); id != tagType {
			return nil
		}
	}
	return &List{
		tagType,
		d,
	}
}

func NewEmptyList(tagType TagId) *List {
	return &List{
		tagType,
		make([]Data, 0),
	}
}

func (n *List) ReadFrom(f io.Reader) (total int64, err error) {
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
	length := bytewrite.BigEndian.Uint32(data)
	n.d = make([]Data, length)
	for i := uint32(0); i < length; i++ {
		if n.d[i], err = newFromTag(n.tagType); err != nil {
			return
		}
		d, err = n.d[i].ReadFrom(f)
		total += d
		if err != nil {
			return
		}
	}
	return
}

func (n *List) WriteTo(f io.Writer) (total int64, err error) {
	var (
		c int
		d int64
	)
	c, err = f.Write([]byte{byte(n.tagType)})
	total += int64(c)
	if err != nil {
		return
	}
	c, err = f.Write(bytewrite.BigEndian.PutUint32(uint32(len(n.d))))
	total += int64(c)
	if err != nil {
		return
	}
	var tagId TagId
	if n.tagType != Tag_End {
		for _, data := range n.d {
			if tagId, err = idFromData(data); err != nil {
				break
			} else if tagId != n.tagType {
				err = &WrongTag{n.tagType, tagId}
				break
			}
			d, err = data.WriteTo(f)
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
	c.d = make([]Data, len(n.d))
	for i, d := range n.d {
		c.d[i] = d.Copy()
	}
	return c
}

func (n *List) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*List); ok {
		if n.tagType == m.tagType && len(n.d) == len(m.d) {
			for i, o := range n.d {
				if !o.Equal(m.d[i]) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (n *List) String() string {
	s := fmt.Sprintf("%d entries of type %s {", len(n.d), n.tagType)
	for _, d := range n.d {
		s += fmt.Sprintf("\n	%s: %s", n.tagType, indent(d.String()))
	}
	return s + "\n}"
}

func (n *List) Set(i int32, d Data) error {
	if i < 0 || i >= int32(len(n.d)) {
		return &BadRange{}
	}
	if err := n.valid(d); err != nil {
		return err
	}
	n.d[i] = d
	return nil
}

func (n *List) Get(i int) Data {
	if i >= 0 && i < len(n.d) {
		return n.d[i]
	}
	return nil
}

func (n *List) Append(d ...Data) error {
	if err := n.valid(d...); err != nil {
		return err
	}
	n.d = append(n.d, d...)
	return nil
}

func (n *List) Insert(i int, d ...Data) error {
	if err := n.valid(d...); err != nil {
		return err
	}
	n.d = append(n.d[:i], append(d, n.d[i:]...)...)
	return nil
}

func (n *List) Remove(i int) {
	if i >= 0 && i < len(n.d) {
		copy(n.d[i:], n.d[i+1:])
		n.d[len(n.d)-1] = nil
		n.d = n.d[:len(n.d)-1]
	}
}

func (n *List) Len() int {
	return len(n.d)
}

func (n *List) valid(d ...Data) error {
	for _, e := range d {
		if t, _ := idFromData(e); t != n.tagType {
			return &WrongTag{n.tagType, t}
		}
	}
	return nil
}

type Compound []*Tag

func NewCompound(d []*Tag) *Compound {
	e := Compound(d)
	return &e
}

func (n *Compound) ReadFrom(f io.Reader) (total int64, err error) {
	var d int64
	*n = Compound(make([]*Tag, 0))
	for {
		data := new(Tag)
		d, err = data.ReadFrom(f)
		total += d
		if err != nil {
			return
		}
		if data.tagType == Tag_End {
			break
		}
		*n = append(*n, data)
	}
	return
}

func (n Compound) WriteTo(f io.Writer) (total int64, err error) {
	var (
		c int
		d int64
	)
	for _, data := range n {
		d, err = data.WriteTo(f)
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
	c := Compound(make([]*Tag, len(n)))
	for i, d := range n {
		c[i] = d.Copy()
	}
	return &c
}

func (n Compound) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Compound); ok {
		if len(n) == len(*m) {
			for _, o := range n {
				if n := m.Get(o.Name()); n == nil || !n.Equal(o) {
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

type IntArray []int32

func NewIntArray(d []int32) *IntArray {
	e := IntArray(d)
	return &e
}

func (n *IntArray) ReadFrom(f io.Reader) (total int64, err error) {
	var c int
	data := make([]byte, 4)
	c, err = io.ReadFull(f, data)
	total += int64(c)
	if err != nil {
		return
	}
	length := bytewrite.BigEndian.Uint32(data)
	*n = make([]int32, length)
	ints := make([]byte, 4*length)
	c, err = io.ReadFull(f, ints)
	total += int64(c)
	for i := uint32(0); i < length; i++ {
		(*n)[i] = int32(bytewrite.BigEndian.Uint32(ints[:4]))
		ints = ints[4:]
	}
	return
}

func (n IntArray) WriteTo(f io.Writer) (total int64, err error) {
	var c int
	c, err = f.Write(bytewrite.BigEndian.PutUint32(uint32(len(n))))
	total += int64(c)
	if err != nil {
		return
	}
	ints := make([]byte, 0, 4*len(n))
	for i := 0; i < len(n); i++ {
		ints = append(ints, bytewrite.BigEndian.PutUint32(uint32(n[i]))...)
	}
	c, err = f.Write(ints)
	total += int64(c)
	return
}

func (n IntArray) Copy() Data {
	c := IntArray(make([]int32, len(n)))
	copy(c, n)
	return &c
}

func (n IntArray) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*IntArray); ok {
		if len(n) == len(*m) {
			for i, o := range n {
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
