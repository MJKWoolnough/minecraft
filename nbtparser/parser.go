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
	"encoding/binary"
	"fmt"
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

type Tag interface {
	io.ReaderFrom
	io.WriterTo
	equaler.Equaler
	Data() Data
	Name() String
	String() string
	Tag() TagId
	Copy() Tag
}

type TagId uint8

func (t TagId) String() string {
	if int(t) < len(tagIdNames) {
		return tagIdNames[t]
	}
	return ""
}

type namedTag struct {
	tagType TagId
	name    String
	d       Data
}

func NewTag(name string, d Data) (n Tag, err error) {
	m := namedTag{
		name: String(name),
		d:    d,
	}
	m.tagType, err = idFromData(d)
	n = &m
	return
}

func (n *namedTag) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	if err = binary.Read(c, binary.BigEndian, &n.tagType); err != nil {
		err = &readError{"named TagId", err}
		return
	}
	if n.tagType == Tag_End {
		n.d = new(end)
	} else {
		if n.d, err = newFromTag(n.tagType); err != nil {
			return
		}
		if _, err = n.name.ReadFrom(c); err != nil {
			err = &readError{"name", err}
			return
		}
		if _, err = n.d.ReadFrom(c); err != nil {
			if _, ok := err.(*readError); !ok {
				err = &readError{n.tagType.String(), err}
			}
		}
	}
	return
}

func (n *namedTag) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	if err = binary.Write(c, binary.BigEndian, n.tagType); err != nil {
		err = &writeError{"named TagId", err}
		return
	}
	if n.tagType == Tag_End {
		return
	}
	if _, err = n.name.WriteTo(c); err != nil {
		return
	}
	_, err = n.d.WriteTo(c)
	return
}

func (n namedTag) Copy() Tag {
	return &namedTag{
		n.tagType,
		n.name,
		n.d.Copy(),
	}
}

func (n namedTag) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*namedTag); ok {
		if n.tagType == m.tagType && n.name == m.name {
			return n.d.Equal(m.d)
		}
	}
	return false
}

func (n namedTag) Data() Data {
	return n.d
}

func (n namedTag) Name() String {
	return n.name
}

func (n namedTag) Tag() TagId {
	return n.tagType
}

func (n namedTag) String() string {
	return fmt.Sprintf("%s(%q): %s", n.tagType, n.name, indent(n.d.String()))
}

type end struct{}

func (n *end) ReadFrom(f io.Reader) (total int64, err error) {
	return
}

func (n *end) WriteTo(f io.Writer) (total int64, err error) {
	return
}

func (n end) Copy() Data {
	return &end{}
}

func (n end) Equal(e equaler.Equaler) bool {
	_, ok := e.(*end)
	return ok
}

func (n end) String() string {
	return ""
}

type Byte byte

func (n *Byte) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	err = binary.Read(c, binary.BigEndian, n)
	return
}

func (n *Byte) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	err = binary.Write(c, binary.BigEndian, n)
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

func (n *Short) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	err = binary.Read(c, binary.BigEndian, n)
	return
}

func (n *Short) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	err = binary.Write(c, binary.BigEndian, n)
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

func (n *Int) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	err = binary.Read(c, binary.BigEndian, n)
	return
}

func (n *Int) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	err = binary.Write(c, binary.BigEndian, n)
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

func (n *Long) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	err = binary.Read(c, binary.BigEndian, n)
	return
}

func (n *Long) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	err = binary.Write(c, binary.BigEndian, n)
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

func (n *Float) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	err = binary.Read(c, binary.BigEndian, n)
	return
}

func (n *Float) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	err = binary.Write(c, binary.BigEndian, n)
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
	return fmt.Sprintf("%d", n)
}

type Double float64

func (n *Double) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	err = binary.Read(c, binary.BigEndian, n)
	return
}

func (n *Double) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	err = binary.Write(c, binary.BigEndian, n)
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
	return fmt.Sprintf("%d", n)
}

type ByteArray []byte

func (n *ByteArray) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	var length uint8
	if err = binary.Read(c, binary.BigEndian, &length); err != nil {
		return
	}
	*n = make([]byte, length)
	err = binary.Read(c, binary.BigEndian, n)
	return
}

func (n *ByteArray) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	if err = binary.Write(c, binary.BigEndian, uint8(len(*n))); err != nil {
		return
	}
	err = binary.Write(c, binary.BigEndian, n)
	return
}

func (n ByteArray) Copy() Data {
	c := ByteArray(make([]byte, len(n)))
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
	return fmt.Sprintf("[%d bytes] %v", len(n), n)
}

type String string

func (n *String) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	var (
		length uint16
		Data   []byte
	)
	if err = binary.Read(c, binary.BigEndian, &length); err != nil {
		return
	}
	Data = make([]byte, length)
	if err = binary.Read(c, binary.BigEndian, &Data); err != nil {
		return
	}
	*n = String(Data)
	return
}

func (n *String) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	Data := []byte(*n)
	length := uint16(len(Data))
	if err = binary.Write(c, binary.BigEndian, length); err != nil {
		return
	}
	err = binary.Write(c, binary.BigEndian, Data)
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

func (n *List) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	if err = binary.Read(c, binary.BigEndian, &n.tagType); err != nil {
		return
	}
	var (
		length int32
		d      Data
	)
	if err = binary.Read(c, binary.BigEndian, &length); err != nil {
		return
	}
	n.d = make([]Data, 0, length)
	for i := int32(0); i < length; i++ {
		if d, err = newFromTag(n.tagType); err != nil {
			return
		}
		if _, err = d.ReadFrom(c); err != nil {
			return
		}
		n.d = append(n.d, d)
	}
	return
}

func (n *List) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	if err = binary.Write(c, binary.BigEndian, n.tagType); err != nil {
		return
	}
	if err = binary.Write(c, binary.BigEndian, int32(len(n.d))); err != nil {
		return
	}
	for _, d := range n.d {
		if _, err = d.WriteTo(c); err != nil {
			return
		}
	}
	return
}

func (n List) Copy() Data {
	c := new(List)
	c.tagType = n.tagType
	c.d = make([]Data, len(n.d))
	for i, d := range n.d {
		c.d[i] = d.Copy()
	}
	return c
}

func (n List) Equal(e equaler.Equaler) bool {
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

func (n List) String() string {
	s := fmt.Sprintf("%d entries of type %s {", len(n.d), n.tagType)
	for _, d := range n.d {
		s += fmt.Sprintf("\n	%s: %s", n.tagType, indent(d.String()))
	}
	return s + "\n}"
}

type Compound []Tag

func (n *Compound) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	*n = Compound(make([]Tag, 0))
	for {
		d := new(namedTag)
		if _, err = d.ReadFrom(f); err != nil {
			return
		}
		if d.tagType == Tag_End {
			break
		}
		*n = append(*n, d)
	}
	return
}

func (n *Compound) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	for _, d := range *n {
		if _, err = d.WriteTo(c); err != nil {
			return
		}
		if d.Tag() == Tag_End {
			return
		}
	}
	err = binary.Write(c, binary.BigEndian, Tag_End)
	return
}

func (n Compound) Copy() Data {
	c := Compound(make([]Tag, len(n)))
	for i, d := range n {
		c[i] = d.Copy()
	}
	return &c
}

func (n Compound) Equal(e equaler.Equaler) bool {
	if m, ok := e.(*Compound); ok {
		if len(n) == len(*m) {
			for i, o := range n {
				if !o.Equal((*m)[i]) {
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

type IntArray []int32

func (n *IntArray) ReadFrom(f io.Reader) (total int64, err error) {
	c := &countReader{Reader: f}
	defer c.BytesRead(&total)
	var length int32
	if err = binary.Read(c, binary.BigEndian, &length); err != nil {
		return
	}
	*n = make([]int32, length)
	err = binary.Read(c, binary.BigEndian, n)
	return
}

func (n *IntArray) WriteTo(f io.Writer) (total int64, err error) {
	c := &countWriter{Writer: f}
	defer c.BytesWritten(&total)
	if err = binary.Write(c, binary.BigEndian, int32(len(*n))); err != nil {
		return
	}
	err = binary.Write(c, binary.BigEndian, n)
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
	return fmt.Sprintf("[%d ints] %v", len(n), n)
}