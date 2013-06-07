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

package nbt

import (
	"fmt"
)

type readError struct {
	where string
	err   error
}

func (r readError) Error() string {
	return fmt.Sprintf("encountered an error while trying to read a %s: %s", r.where, r.err)
}

type writeError struct {
	where string
	err   error
}

func (w writeError) Error() string {
	return fmt.Sprintf("encountered an error while trying to write a %s: %s", w.where, w.err)
}

type unknownTag struct {
	TagId
}

func (u unknownTag) Error() string {
	return fmt.Sprintf("discovered unknown TagId with id %d", u.TagId)
}

type wrongTag struct {
	expecting, got TagId
}

func (w wrongTag) Error() string {
	return fmt.Sprintf("expecting tag id %d, got %d", w.expecting, w.got)
}

type badRange struct{}

func (b badRange) Error() string {
	return "given index was out-of-range"
}

func newFromTag(id TagId) (d Data, err error) {
	switch id {
	case Tag_Byte:
		d = new(Byte)
	case Tag_Short:
		d = new(Short)
	case Tag_Int:
		d = new(Int)
	case Tag_Long:
		d = new(Long)
	case Tag_Float:
		d = new(Float)
	case Tag_Double:
		d = new(Double)
	case Tag_ByteArray:
		d = new(ByteArray)
	case Tag_String:
		d = new(String)
	case Tag_List:
		d = new(List)
	case Tag_Compound:
		d = new(Compound)
	case Tag_IntArray:
		d = new(IntArray)
	default:
		err = &unknownTag{id}
	}
	return
}

func idFromData(d Data) (t TagId, err error) {
	switch d.(type) {
	case *Byte:
		t = Tag_Byte
	case *Short:
		t = Tag_Short
	case *Int:
		t = Tag_Int
	case *Long:
		t = Tag_Long
	case *Float:
		t = Tag_Float
	case *Double:
		t = Tag_Double
	case *ByteArray:
		t = Tag_ByteArray
	case *String:
		t = Tag_String
	case *List:
		t = Tag_List
	case *Compound:
		t = Tag_Compound
	case *IntArray:
		t = Tag_IntArray
	default:
		err = fmt.Errorf("couldn't determine tag type")
	}
	return
}

func indent(s string) (out string) {
	last := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			out += s[last:i+1] + "	"
			last = i + 1
		}
	}
	out += s[last:]
	return out
}
