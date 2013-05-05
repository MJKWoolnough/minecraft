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

// Package nbtparser implements a full Named Binary Tag parser, based on the specs at
// http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt
package nbtparser

import (
	"encoding/binary"
	"fmt"
	"github.com/MJKWoolnough/equaler"
	"io"
)

// Tag Types
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

var tagNames = map[uint8]string{
	NBTTag_End:       "End",
	NBTTag_Byte:      "Byte",
	NBTTag_Short:     "Short",
	NBTTag_Int:       "Int",
	NBTTag_Long:      "Long",
	NBTTag_Float:     "Float",
	NBTTag_Double:    "Double",
	NBTTag_ByteArray: "Byte Array",
	NBTTag_String:    "String",
	NBTTag_List:      "List",
	NBTTag_Compound:  "Compound",
	NBTTag_IntArray:  "Int Array",
}

// TagName converts a tag id into its canonical name.
func TagName(id uint8) string {
	return tagNames[id]
}

// NBTTag is the main interface for all of the NBT types.
// All tags implement the io.ReaderFrom and io.WriterTo interfaces.
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

type nbtTagLevel interface {
	NBTTag
	setLevel(uint32)
	readData(io.Reader) error
	writeData(io.Writer) error
}

type nbtString string

func (n *nbtString) ReadFrom(file io.Reader) (int64, error) {
	var (
		length uint16
		data   []byte
	)
	if err := binary.Read(file, binary.BigEndian, &length); err != nil {
		return 0, err
	}
	data = make([]byte, length)
	if err := binary.Read(file, binary.BigEndian, &data); err != nil {
		return 2, err
	}
	*n = nbtString(data)
	return int64(length) + 2, nil
}

func (n *nbtString) WriteTo(file io.Writer) (int64, error) {
	data := []byte(*n)
	length := uint16(len(data))
	if err := binary.Write(file, binary.BigEndian, &length); err != nil {
		return 0, err
	}
	if err := binary.Write(file, binary.BigEndian, []byte(data)); err != nil {
		return 2, err
	}
	return int64(length) + 2, nil
}

type nbttag struct {
	tagType uint8
	name    nbtString
}

func (n *nbttag) readTag(tagType uint8, file io.Reader) (int64, error) {
	n.tagType = tagType
	if file != nil {
		return n.name.ReadFrom(file)
	}
	return 0, nil
}

func (n *nbttag) writeTag(file io.Writer) (int64, error) {
	if err := binary.Write(file, binary.BigEndian, n.tagType); err != nil {
		return 0, err
	}
	d, err := n.name.WriteTo(file)
	return d + 1, err
}

func (n *nbttag) SetName(name string) {
	n.name = nbtString(name)
}

func (n *nbttag) TypeString() string {
	return tagNames[n.tagType]
}

func (n *nbttag) Name() string {
	return string(n.name)
}

func (n *nbttag) Type() uint8 {
	return n.tagType
}

func (n *nbttag) TagEnd() *NBTTagEnd {
	return nil
}

// TagByte returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagByte() *NBTTagByte {
	return nil
}

// TagShort returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagShort() *NBTTagShort {
	return nil
}

// TagInt returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagInt() *NBTTagInt {
	return nil
}

// TagLong returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagLong() *NBTTagLong {
	return nil
}

// TagFloat returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagFloat() *NBTTagFloat {
	return nil
}

// TagDouble returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagDouble() *NBTTagDouble {
	return nil
}

// TagByteArray returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagByteArray() *NBTTagByteArray {
	return nil
}

// TagString returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagString() *NBTTagString {
	return nil
}

// TagList returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagList() *NBTTagList {
	return nil
}

// TagCompound returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagCompound() *NBTTagCompound {
	return nil
}

// TagIntArray returns the tag if appropriate, nil if otherwise.
func (n *nbttag) TagIntArray() *NBTTagIntArray {
	return nil
}

// NBTTagEnd marks the end of a compound tag
type NBTTagEnd struct {
	*nbttag
}

func (n *NBTTagEnd) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	return n.readTag(NBTTag_End, nil)
}

func (n *NBTTagEnd) WriteTo(file io.Writer) (int64, error) {
	id := NBTTag_End
	return 1, binary.Write(file, binary.BigEndian, &id)
}

func (n *NBTTagEnd) String() string {
	return "TAG_End"
}

func (n *NBTTagEnd) TagEnd() *NBTTagEnd {
	return n
}

func (n *NBTTagEnd) Copy() NBTTag {
	return NewTagEnd()
}

func (n *NBTTagEnd) Equal(e equaler.Equaler) bool {
	_, ok := e.(*NBTTagEnd)
	return ok
}

type NBTTagByte struct {
	*nbttag
	data int8
}

func (n *NBTTagByte) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_Byte, file); err != nil {
		return 0, err
	} else {
		c = d
	}
	if err := binary.Read(file, binary.BigEndian, &n.data); err != nil {
		return 1, err
	}
	return c + 1, nil
}

func (n *NBTTagByte) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Write(file, binary.BigEndian, n.data); err != nil {
		return c, err
	}
	return c + 1, nil
}

func (n *NBTTagByte) String() string {
	return "TAG_Byte(\"" + n.Name() + "\"): " + fmt.Sprint(n.data)
}

func (n *NBTTagByte) Get() int8 {
	return n.data
}

func (n *NBTTagByte) Set(data int8) {
	n.data = data
}

func (n *NBTTagByte) TagByte() *NBTTagByte {
	return n
}

func (n *NBTTagByte) Copy() NBTTag {
	return NewTagByte(n.Name(), n.data)
}

func (n *NBTTagByte) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagByte); ok {
		if n.Name() == o.Name() && n.data == o.data {
			return true
		}
	}
	return false
}

type NBTTagShort struct {
	*nbttag
	data int16
}

func (n *NBTTagShort) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_Short, file); err != nil {
		return 0, err
	} else {
		c = d
	}
	if err := binary.Read(file, binary.BigEndian, &n.data); err != nil {
		return c, err
	}
	return c + 2, nil
}

func (n *NBTTagShort) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Write(file, binary.BigEndian, n.data); err != nil {
		return c, err
	}
	return c + 2, nil
}

func (n *NBTTagShort) String() string {
	return "TAG_Short(\"" + n.Name() + "\"): " + fmt.Sprint(n.data)
}

func (n *NBTTagShort) Get() int16 {
	return n.data
}

func (n *NBTTagShort) Set(data int16) {
	n.data = data
}

func (n *NBTTagShort) TagShort() *NBTTagShort {
	return n
}

func (n *NBTTagShort) Copy() NBTTag {
	return NewTagShort(n.Name(), n.data)
}

func (n *NBTTagShort) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagShort); ok {
		if n.Name() == o.Name() && n.data == o.data {
			return true
		}
	}
	return false
}

type NBTTagInt struct {
	*nbttag
	data int32
}

func (n *NBTTagInt) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_Int, file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Read(file, binary.BigEndian, &n.data); err != nil {
		return c, err
	}
	return c + 4, nil
}

func (n *NBTTagInt) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Write(file, binary.BigEndian, n.data); err != nil {
		return c, err
	}
	return c + 4, nil
}

func (n *NBTTagInt) String() string {
	return "TAG_Int(\"" + n.Name() + "\"): " + fmt.Sprint(n.data)
}

func (n *NBTTagInt) Get() int32 {
	return n.data
}

func (n *NBTTagInt) Set(data int32) {
	n.data = data
}

func (n *NBTTagInt) TagInt() *NBTTagInt {
	return n
}

func (n *NBTTagInt) Copy() NBTTag {
	return NewTagInt(n.Name(), n.data)
}

func (n *NBTTagInt) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagInt); ok {
		if n.Name() == o.Name() && n.data == o.data {
			return true
		}
	}
	return false
}

type NBTTagLong struct {
	*nbttag
	data int64
}

func (n *NBTTagLong) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_Long, file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Read(file, binary.BigEndian, &n.data); err != nil {
		return c, err
	}
	return c + 8, nil
}

func (n *NBTTagLong) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Write(file, binary.BigEndian, n.data); err != nil {
		return c, err
	}
	return c + 8, nil
}

func (n *NBTTagLong) String() string {
	return "TAG_Long(\"" + n.Name() + "\"): " + fmt.Sprint(n.data)
}

func (n *NBTTagLong) Get() int64 {
	return n.data
}

func (n *NBTTagLong) Set(data int64) {
	n.data = data
}

func (n *NBTTagLong) TagLong() *NBTTagLong {
	return n
}

func (n *NBTTagLong) Copy() NBTTag {
	return NewTagLong(n.Name(), n.data)
}

func (n *NBTTagLong) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagLong); ok {
		if n.Name() == o.Name() && n.data == o.data {
			return true
		}
	}
	return false
}

type NBTTagFloat struct {
	*nbttag
	data float32
}

func (n *NBTTagFloat) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_Float, file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Read(file, binary.BigEndian, &n.data); err != nil {
		return c, err
	}
	return c + 4, nil
}

func (n *NBTTagFloat) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Write(file, binary.BigEndian, n.data); err != nil {
		return c, err
	}
	return c + 4, nil
}

func (n *NBTTagFloat) String() string {
	return "TAG_Float(\"" + n.Name() + "\"): " + fmt.Sprint(n.data)
}

func (n *NBTTagFloat) Get() float32 {
	return n.data
}

func (n *NBTTagFloat) Set(data float32) {
	n.data = data
}

func (n *NBTTagFloat) TagFloat() *NBTTagFloat {
	return n
}

func (n *NBTTagFloat) Copy() NBTTag {
	return NewTagFloat(n.Name(), n.data)
}

func (n *NBTTagFloat) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagFloat); ok {
		if n.Name() == o.Name() && n.data == o.data {
			return true
		}
	}
	return false
}

type NBTTagDouble struct {
	*nbttag
	data float64
}

func (n *NBTTagDouble) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_Double, file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Read(file, binary.BigEndian, &n.data); err != nil {
		return c, err
	}
	return c + 8, nil
}

func (n *NBTTagDouble) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Write(file, binary.BigEndian, n.data); err != nil {
		return c, err
	}
	return c + 8, nil
}

func (n *NBTTagDouble) String() string {
	return "TAG_Double(\"" + n.Name() + "\"): " + fmt.Sprint(n.data)
}

func (n *NBTTagDouble) Get() float64 {
	return n.data
}

func (n *NBTTagDouble) Set(data float64) {
	n.data = data
}

func (n *NBTTagDouble) TagDouble() *NBTTagDouble {
	return n
}

func (n *NBTTagDouble) Copy() NBTTag {
	return NewTagDouble(n.Name(), n.data)
}

func (n *NBTTagDouble) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagDouble); ok {
		if n.Name() == o.Name() && n.data == o.data {
			return true
		}
	}
	return false
}

type NBTTagByteArray struct {
	*nbttag
	data []byte
}

func (n *NBTTagByteArray) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var (
		c      int64
		length int32
	)
	if d, err := n.readTag(NBTTag_ByteArray, file); err != nil {
		return d, err
	} else {
		c = d
	}
	if err := binary.Read(file, binary.BigEndian, &length); err != nil {
		return c, err
	}
	c += 4
	n.data = make([]byte, length)
	if err := binary.Read(file, binary.BigEndian, &n.data); err != nil {
		return c, err
	}
	return c + int64(length), nil
}

func (n *NBTTagByteArray) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	length := int32(len(n.data))
	if err := binary.Write(file, binary.BigEndian, length); err != nil {
		return c, err
	}
	c += 4
	if err := binary.Write(file, binary.BigEndian, n.data); err != nil {
		return c, err
	}
	return c + int64(length), nil
}

func (n *NBTTagByteArray) String() string {
	return "TAG_ByteArray(\"" + n.Name() + "\"): " + fmt.Sprint(n.data)
}

func (n *NBTTagByteArray) Length() int32 {
	return int32(len(n.data))
}

func (n *NBTTagByteArray) Get(index int32) byte {
	return n.data[index]
}

func (n *NBTTagByteArray) GetArray() []byte {
	return n.data
}

func (n *NBTTagByteArray) Set(index int32, data byte) {
	n.data[index] = data
}

func (n *NBTTagByteArray) Append(data byte) {
	n.data = append(n.data, data)
}

func (n *NBTTagByteArray) TagByteArray() *NBTTagByteArray {
	return n
}

func (n *NBTTagByteArray) Copy() NBTTag {
	b := NewTagByteArray(n.Name(), make([]byte, 0))
	for _, i := range n.data {
		b.Append(i)
	}
	return b
}

func (n *NBTTagByteArray) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagByteArray); ok {
		if n.Name() == o.Name() && len(n.data) == len(o.data) {
			for k, v := range n.data {
				if o.data[k] != v {
					return false
				}
			}
			return true
		}
	}
	return false
}

type NBTTagString struct {
	*nbttag
	data nbtString
}

func (n *NBTTagString) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_String, file); err != nil {
		return d, err
	} else {
		c = d
	}
	if d, err := n.data.ReadFrom(file); err != nil {
		return c + d, err
	} else {
		c += d
	}
	return c, nil
}

func (n *NBTTagString) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	if d, err := n.data.WriteTo(file); err != nil {
		return c + d, err
	} else {
		c += d
	}
	return c, nil
}

func (n *NBTTagString) String() string {
	return "TAG_String(\"" + n.Name() + "\"): " + n.Get()
}

func (n *NBTTagString) Length() int32 {
	return int32(len(string(n.data)))
}

func (n *NBTTagString) Get() string {
	return string(n.data)
}

func (n *NBTTagString) Set(data string) {
	n.data = nbtString(data)
}

func (n *NBTTagString) TagString() *NBTTagString {
	return n
}

func (n *NBTTagString) Copy() NBTTag {
	return NewTagString(n.Name(), n.Get())
}

func (n *NBTTagString) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagString); ok {
		if n.Name() == o.Name() && n.data == o.data {
			return true
		}
	}
	return false
}

type NBTTagList struct {
	*nbttag
	dType uint8
	data  []interface{}
	level uint32
}

func (n *NBTTagList) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_List, file); err != nil {
		return d, err
	} else {
		c = d
	}
	d, err := n.readData(file)
	return c + d, err
}

func (n *NBTTagList) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	d, err := n.writeData(file)
	return c + d, err
}

func (n *NBTTagList) readData(file io.Reader) (int64, error) {
	var (
		c      int64
		length int32
	)
	if n.nbttag == nil {
		n.nbttag = new(nbttag)
		n.tagType = NBTTag_List
	}
	if err := binary.Read(file, binary.BigEndian, &n.dType); err != nil {
		return 0, err
	} else if err := binary.Read(file, binary.BigEndian, &length); err != nil {
		return 1, err
	}
	c = 5
	n.data = make([]interface{}, length)
	for i := int32(0); i < length; i++ {
		switch n.dType {
		case NBTTag_Byte:
			var data int8
			if err := binary.Read(file, binary.BigEndian, &data); err != nil {
				return c, err
			}
			c++
			n.data[i] = data
		case NBTTag_Short:
			var data int16
			if err := binary.Read(file, binary.BigEndian, &data); err != nil {
				return c, err
			}
			c += 2
			n.data[i] = data
		case NBTTag_Int:
			var data int32
			if err := binary.Read(file, binary.BigEndian, &data); err != nil {
				return c, err
			}
			c += 4
			n.data[i] = data
		case NBTTag_Long:
			var data int64
			if err := binary.Read(file, binary.BigEndian, &data); err != nil {
				return c, err
			}
			c += 8
			n.data[i] = data
		case NBTTag_Float:
			var data float32
			if err := binary.Read(file, binary.BigEndian, &data); err != nil {
				return c, err
			}
			c += 4
			n.data[i] = data
		case NBTTag_Double:
			var data float64
			if err := binary.Read(file, binary.BigEndian, &data); err != nil {
				return c, err
			}
			c += 8
			n.data[i] = data
		case NBTTag_ByteArray:
			var length int32
			if err := binary.Read(file, binary.BigEndian, &length); err != nil {
				return c, err
			}
			c += 4
			data := make([]byte, length)
			if err := binary.Read(file, binary.BigEndian, &data); err != nil {
				return c, err
			}
			c += int64(length)
			n.data[i] = data
		case NBTTag_String:
			var data nbtString
			if d, err := data.ReadFrom(file); err != nil {
				return c + d, err
			} else {
				c += d
			}
			n.data[i] = data
		case NBTTag_List:
			data := new(NBTTagList)
			if d, err := data.readData(file); err != nil {
				return c + d, err
			} else {
				c += d
			}
			n.data[i] = data
		case NBTTag_Compound:
			data := new(NBTTagCompound)
			if d, err := data.readData(file); err != nil {
				return c + d, err
			} else {
				c += d
			}
			n.data[i] = data
		case NBTTag_IntArray:
			data := new(NBTTagIntArray)
			if d, err := data.readData(file); err != nil {
				return c + d, err
			} else {
				c += d
			}
			n.data[i] = data
		default:
			return c, fmt.Errorf("Unknown data type for NBT List: %d", n.dType)
		}
	}
	return c, nil
}

func (n *NBTTagList) writeData(file io.Writer) (int64, error) {
	if err := binary.Write(file, binary.BigEndian, &n.dType); err != nil {
		return 0, err
	}
	length := int32(len(n.data))
	if err := binary.Write(file, binary.BigEndian, &length); err != nil {
		return 1, err
	}
	c := int64(5)
	for i := 0; i < len(n.data); i++ {
		switch n.dType {
		case NBTTag_Byte:
			if data, ok := n.data[i].(int8); ok {
				if err := binary.Write(file, binary.BigEndian, &data); err != nil {
					return c, err
				}
				c++
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		case NBTTag_Short:
			if data, ok := n.data[i].(int16); ok {
				if err := binary.Write(file, binary.BigEndian, &data); err != nil {
					return c, err
				}
				c += 2
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		case NBTTag_Int:
			if data, ok := n.data[i].(int32); ok {
				if err := binary.Write(file, binary.BigEndian, &data); err != nil {
					return c, err
				}
				c += 4
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		case NBTTag_Long:
			if data, ok := n.data[i].(int64); ok {
				if err := binary.Write(file, binary.BigEndian, &data); err != nil {
					return c, err
				}
				c += 8
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		case NBTTag_Float:
			if data, ok := n.data[i].(float32); ok {
				if err := binary.Write(file, binary.BigEndian, &data); err != nil {
					return c, err
				}
				c += 4
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		case NBTTag_Double:
			if data, ok := n.data[i].(float64); ok {
				if err := binary.Write(file, binary.BigEndian, &data); err != nil {
					return c, err
				}
				c += 8
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		case NBTTag_ByteArray:
			if data, ok := n.data[i].([]byte); ok {
				length := int32(len(data))
				if err := binary.Write(file, binary.BigEndian, &length); err != nil {
					return c, err
				}
				c += 4
				if err := binary.Write(file, binary.BigEndian, &data); err != nil {
					return c, err
				}
				c += int64(length)
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		case NBTTag_String:
			if data, ok := n.data[i].(*nbtString); ok {
				if d, err := data.WriteTo(file); err != nil {
					return c + d, err
				} else {
					c += d
				}
			}
		case NBTTag_Compound:
			if tag, ok := n.data[i].(*NBTTagCompound); ok {
				if d, err := tag.writeData(file); err != nil {
					return c + d, err
				} else {
					c += d
				}
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		case NBTTag_List:
			if tag, ok := n.data[i].(*NBTTagList); ok {
				if d, err := tag.writeData(file); err != nil {
					return c + d, err
				} else {
					c += d
				}
			}
		case NBTTag_IntArray:
			if tag, ok := n.data[i].(*NBTTagIntArray); ok {
				if d, err := tag.writeData(file); err != nil {
					return c + d, err
				} else {
					c += d
				}
			} else {
				return c, fmt.Errorf("Wrong data type found, expecting: %d", n.dType)
			}
		default:
			return c, fmt.Errorf("Unknown data type found: %d", n.dType)
		}
	}
	return c, nil
}

func (n *NBTTagList) setLevel(level uint32) {
	n.level = level
	for i := 0; i < len(n.data); i++ {
		if tag, ok := n.data[i].(*NBTTagCompound); ok {
			tag.setLevel(level + 1)
			continue
		}
		if tag, ok := n.data[i].(*NBTTagList); ok {
			tag.setLevel(level + 1)
			continue
		}
	}
}

func (n *NBTTagList) String() string {
	toRet := "TAG_List(\"" + n.Name() + "\"): " + fmt.Sprint(len(n.data)) + " entries of type " + TagName(n.dType) + " {\n"
	indent := ""
	for i := uint32(0); i < n.level; i++ {
		indent += "	"
	}
	for i := 0; i < len(n.data); i++ {
		if tag, ok := n.data[i].(nbtTagLevel); ok {
			toRet += indent + "	" + tag.String() + "\n"
		} else {
			toRet += indent + "	" + TagName(n.dType) + ": " + fmt.Sprint(n.data[i]) + "\n"
		}
	}
	toRet += indent + "}"
	return toRet
}

func (n *NBTTagList) checkType(data interface{}) bool {
	switch n.dType {
	case NBTTag_Byte:
		if _, ok := data.(int8); !ok {
			return false
		}
	case NBTTag_Short:
		if _, ok := data.(int16); !ok {
			return false
		}
	case NBTTag_Int:
		if _, ok := data.(int32); !ok {
			return false
		}
	case NBTTag_Long:
		if _, ok := data.(int64); !ok {
			return false
		}
	case NBTTag_Float:
		if _, ok := data.(float32); !ok {
			return false
		}
	case NBTTag_Double:
		if _, ok := data.(float64); !ok {
			return false
		}
	case NBTTag_ByteArray:
		if _, ok := data.([]byte); !ok {
			return false
		}
	case NBTTag_List:
		if _, ok := data.(*NBTTagList); !ok {
			return false
		}
	case NBTTag_Compound:
		if _, ok := data.(*NBTTagCompound); !ok {
			return false
		}
	case NBTTag_IntArray:
		if _, ok := data.(*NBTTagIntArray); !ok {
			return false
		}
	default:
		return false
	}
	return true
}

func (n *NBTTagList) Length() int32 {
	return int32(len(n.data))
}

func (n *NBTTagList) Get(index int32) interface{} {
	return n.data[index]
}

func (n *NBTTagList) GetArray() []interface{} {
	return n.data
}

func (n *NBTTagList) GetType() uint8 {
	return n.dType
}

func (n *NBTTagList) Set(index int32, data interface{}) bool {
	if !n.checkType(data) {
		return false
	}
	if data == nil {
		t := make([]interface{}, 0)
		i := int(index)
		for k, v := range n.data {
			if k != i {
				t = append(t, v)
			}
		}
		n.data = t
	} else {
		n.data[index] = data
		n.setLevel(n.level)
	}
	return true
}

func (n *NBTTagList) SetArray(dType uint8, data []interface{}) bool {
	oldDType := n.dType
	n.dType = dType
	for i := 0; i < len(data); i++ {
		if !n.checkType(data[i]) {
			n.dType = oldDType
			return false
		}
	}
	n.data = data
	return true
}

func (n *NBTTagList) Append(data interface{}) bool {
	if !n.checkType(data) {
		return false
	}
	n.data = append(n.data, data)
	return true
}

func (n *NBTTagList) TagList() *NBTTagList {
	return n
}

func (n *NBTTagList) Copy() NBTTag {
	b := NewTagList(n.Name(), n.dType, make([]interface{}, 0))
	for _, i := range n.data {
		if j, ok := i.(NBTTag); ok {
			b.Append(j.Copy())
		} else {
			b.Append(i)
		}
	}
	return b
}

func (n *NBTTagList) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagList); ok {
		if n.Name() == o.Name() && len(n.data) == len(o.data) {
			for k, v := range n.data {
				if w, ok := v.(equaler.Equaler); ok {
					if x, ok := o.data[k].(equaler.Equaler); ok {
						if !w.Equal(x) {
							return false
						}
					} else {
						panic("Malformed NBT_List Data")
					}
				} else {
					panic("Malformed NBT_List Data")
				}
			}
			return true
		}
	}
	return false
}

type NBTTagCompound struct {
	*nbttag
	data  []NBTTag
	level uint32
}

func (n *NBTTagCompound) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_Compound, file); err != nil {
		return d, err
	} else {
		c += d
	}
	d, err := n.readData(file)
	return c + d, err
}

func (n *NBTTagCompound) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	d, err := n.writeData(file)
	return c + d, err
}

func (n *NBTTagCompound) readData(file io.Reader) (int64, error) {
	var c int64
	if n.nbttag == nil {
		n.nbttag = new(nbttag)
		n.tagType = NBTTag_Compound
	}
	n.data = make([]NBTTag, 0)
	for {
		if tag, d, err := readTag(file); err != nil {
			return c + d, err
		} else {
			c += d
			n.data = append(n.data, tag)
			if _, ok := tag.(*NBTTagEnd); ok {
				break
			}
		}
	}
	return c, nil
}

func (n *NBTTagCompound) writeData(file io.Writer) (int64, error) {
	var c int64
	for i := 0; i < len(n.data); i++ {
		if d, err := n.data[i].WriteTo(file); err != nil {
			return c + d, err
		} else {
			c += d
		}
		if _, ok := n.data[i].(*NBTTagEnd); ok {
			return c, nil
		}
	}
	d, err := NewTagEnd().WriteTo(file)
	return c + d, err
}

func (n *NBTTagCompound) setLevel(level uint32) {
	n.level = level
	for i := 0; i < len(n.data); i++ {
		if tag, ok := n.data[i].(*NBTTagCompound); ok {
			tag.setLevel(level + 1)
			continue
		}
		if tag, ok := n.data[i].(*NBTTagList); ok {
			tag.setLevel(level + 1)
			continue
		}
	}
}

func (n *NBTTagCompound) String() string {
	toRet := "TAG_Compound(\"" + n.Name() + "\"): " + fmt.Sprint(len(n.data)) + " entries {\n"
	indent := ""
	for i := uint32(0); i < n.level; i++ {
		indent += "	"
	}
	for i := 0; i < len(n.data); i++ {
		toRet += indent + "	" + n.data[i].String() + "\n"
	}
	toRet += indent + "}"
	return toRet
}

func (n *NBTTagCompound) Length() int32 {
	return int32(len(n.data))
}

func (n *NBTTagCompound) Get(index int32) NBTTag {
	return n.data[index]
}

func (n *NBTTagCompound) GetArray() []NBTTag {
	return n.data
}

func (n *NBTTagCompound) SetArray(data []NBTTag) {
	n.data = data
}

func (n *NBTTagCompound) GetTag(tagName string) NBTTag {
	for i := 0; i < len(n.data); i++ {
		if n.data[i].Name() == tagName {
			return n.data[i]
		}
	}
	return nil
}

func (n *NBTTagCompound) RemoveTag(tagName string) NBTTag {
	t := make([]NBTTag, 0)
	for i := 0; i < len(n.data); i++ {
		if n.data[i].Name() != tagName {
			t = append(t, n.data[i])
		}
	}
	n.data = t
	return nil
}

func (n *NBTTagCompound) Set(index int32, data NBTTag) {
	if data == nil {
		t := make([]NBTTag, 0)
		i := int(index)
		for k, v := range n.data {
			if k != i {
				t = append(t, v)
			}
		}
		n.data = t
	} else {
		n.data[index] = data
		n.setLevel(n.level)
	}
}

func (n *NBTTagCompound) Append(data NBTTag) {
	n.data = append(n.data, data)
	n.data[len(n.data)-2], n.data[len(n.data)-1] = n.data[len(n.data)-1], n.data[len(n.data)-2]
}

func (n *NBTTagCompound) TagCompound() *NBTTagCompound {
	return n
}

func (n *NBTTagCompound) Copy() NBTTag {
	b := NewTagCompound(n.Name(), make([]NBTTag, 0))
	for _, i := range n.data {
		b.Append(i.Copy())
	}
	return b
}

func (n *NBTTagCompound) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagCompound); ok {
		if n.Name() == o.Name() && len(n.data) == len(o.data) {
			for k, v := range n.data {
				if w, ok := v.(equaler.Equaler); ok {
					if x, ok := o.data[k].(equaler.Equaler); ok {
						if !w.Equal(x) {
							return false
						}
					} else {
						panic("Malformed NBT_Compound Data")
					}
				} else {
					panic("Malformed NBT_Compound Data")
				}
			}
			return true
		}
	}
	return false
}

type NBTTagIntArray struct {
	*nbttag
	data []int32
}

func (n *NBTTagIntArray) ReadFrom(file io.Reader) (int64, error) {
	n.nbttag = new(nbttag)
	var c int64
	if d, err := n.readTag(NBTTag_IntArray, file); err != nil {
		return d, err
	} else {
		c = d
	}
	d, err := n.readData(file)
	return c + d, err
}

func (n *NBTTagIntArray) WriteTo(file io.Writer) (int64, error) {
	var c int64
	if d, err := n.writeTag(file); err != nil {
		return d, err
	} else {
		c = d
	}
	d, err := n.writeData(file)
	return c + d, err
}

func (n *NBTTagIntArray) readData(file io.Reader) (int64, error) {
	var length int32
	if n.nbttag == nil {
		n.nbttag = new(nbttag)
		n.tagType = NBTTag_IntArray
	}
	if err := binary.Read(file, binary.BigEndian, &length); err != nil {
		return 0, err
	}
	n.data = make([]int32, length)
	if err := binary.Read(file, binary.BigEndian, &n.data); err != nil {
		return 4, err
	}
	return int64(length) + 4, nil
}

func (n *NBTTagIntArray) writeData(file io.Writer) (int64, error) {
	length := int32(len(n.data))
	if err := binary.Write(file, binary.BigEndian, length); err != nil {
		return 0, err
	}
	if err := binary.Write(file, binary.BigEndian, n.data); err != nil {
		return 4, err
	}
	return int64(length) + 4, nil
}

func (n *NBTTagIntArray) String() string {
	return "TAG_IntArray(\"" + n.Name() + "\"): " + fmt.Sprint(n.data)
}

func (n *NBTTagIntArray) Length() int32 {
	return int32(len(n.data))
}

func (n *NBTTagIntArray) Get(index int32) int32 {
	return n.data[index]
}

func (n *NBTTagIntArray) Set(index, data int32) {
	n.data[index] = data
}

func (n *NBTTagIntArray) GetArray() []int32 {
	return n.data
}

func (n *NBTTagIntArray) Append(data int32) {
	n.data = append(n.data, data)
}

func (n *NBTTagIntArray) TagIntArray() *NBTTagIntArray {
	return n
}

func (n *NBTTagIntArray) Copy() NBTTag {
	b := NewTagIntArray(n.Name(), make([]int32, 0))
	for _, i := range n.data {
		b.Append(i)
	}
	return b
}

func (n *NBTTagIntArray) Equal(e equaler.Equaler) bool {
	if o, ok := e.(*NBTTagIntArray); ok {
		if n.Name() == o.Name() && len(n.data) == len(o.data) {
			for k, v := range n.data {
				if v != o.data[k] {
					return false
				}
			}
			return true
		}
	}
	return false
}

// ParseFile constructs an entire NBT Tree from a reader.
func ParseFile(file io.Reader) (*NBTTagCompound, int64, error) {
	var c int64
	if toRet, d, err := readTag(file); err != nil {
		return nil, d, err
	} else if t, ok := toRet.(*NBTTagCompound); ok {
		t.setLevel(0)
		return t, d, nil
	} else {
		c = d
	}
	return nil, c, fmt.Errorf("Invalid NBT File")
}

func readTag(file io.Reader) (NBTTag, int64, error) {
	var (
		id  uint8
		tag NBTTag
	)
	if err := binary.Read(file, binary.BigEndian, &id); err != nil {
		return nil, 0, err
	}
	switch id {
	case NBTTag_End:
		tag = new(NBTTagEnd)
	case NBTTag_Byte:
		tag = new(NBTTagByte)
	case NBTTag_Short:
		tag = new(NBTTagShort)
	case NBTTag_Int:
		tag = new(NBTTagInt)
	case NBTTag_Long:
		tag = new(NBTTagLong)
	case NBTTag_Float:
		tag = new(NBTTagFloat)
	case NBTTag_Double:
		tag = new(NBTTagDouble)
	case NBTTag_ByteArray:
		tag = new(NBTTagByteArray)
	case NBTTag_String:
		tag = new(NBTTagString)
	case NBTTag_List:
		tag = new(NBTTagList)
	case NBTTag_Compound:
		tag = new(NBTTagCompound)
	case NBTTag_IntArray:
		tag = new(NBTTagIntArray)
	default:
		return nil, 1, fmt.Errorf("Invalid tag found during parsing")
	}
	d, err := tag.ReadFrom(file)
	return tag, d + 1, err
}

func NewTagEnd() *NBTTagEnd {
	toRet := new(NBTTagEnd)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_End
	return toRet
}

func NewTagByte(name string, data int8) *NBTTagByte {
	toRet := new(NBTTagByte)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_Byte
	toRet.name = nbtString(name)
	toRet.data = data
	return toRet
}

func NewTagShort(name string, data int16) *NBTTagShort {
	toRet := new(NBTTagShort)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_Short
	toRet.name = nbtString(name)
	toRet.data = data
	return toRet
}

func NewTagInt(name string, data int32) *NBTTagInt {
	toRet := new(NBTTagInt)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_Int
	toRet.name = nbtString(name)
	toRet.data = data
	return toRet
}

func NewTagLong(name string, data int64) *NBTTagLong {
	toRet := new(NBTTagLong)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_Long
	toRet.name = nbtString(name)
	toRet.data = data
	return toRet
}

func NewTagFloat(name string, data float32) *NBTTagFloat {
	toRet := new(NBTTagFloat)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_Float
	toRet.name = nbtString(name)
	toRet.data = data
	return toRet
}

func NewTagDouble(name string, data float64) *NBTTagDouble {
	toRet := new(NBTTagDouble)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_Double
	toRet.name = nbtString(name)
	toRet.data = data
	return toRet
}

func NewTagByteArray(name string, data []byte) *NBTTagByteArray {
	toRet := new(NBTTagByteArray)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_ByteArray
	toRet.name = nbtString(name)
	toRet.data = data
	return toRet
}

func NewTagString(name string, data string) *NBTTagString {
	toRet := new(NBTTagString)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_String
	toRet.name = nbtString(name)
	toRet.data = nbtString(data)
	return toRet
}

func NewTagList(name string, dType uint8, data []interface{}) *NBTTagList {
	toRet := new(NBTTagList)
	if !toRet.SetArray(dType, data) {
		return nil
	}
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_List
	toRet.name = nbtString(name)
	return toRet
}

func NewTagCompound(name string, data []NBTTag) *NBTTagCompound {
	if data[len(data)-1].TagEnd() != nil {
		data = data[0 : len(data)-1]
	}
	toRet := new(NBTTagCompound)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_Compound
	toRet.name = nbtString(name)
	toRet.data = data
	toRet.setLevel(0)
	return toRet
}

func NewTagIntArray(name string, data []int32) *NBTTagIntArray {
	toRet := new(NBTTagIntArray)
	toRet.nbttag = new(nbttag)
	toRet.tagType = NBTTag_IntArray
	toRet.name = nbtString(name)
	toRet.data = data
	return toRet
}
