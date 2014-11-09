package nbt

import (
	"fmt"
)

type ReadError struct {
	Where string
	Err   error
}

func (r ReadError) Error() string {
	return fmt.Sprintf("encountered an error while trying to read a %s: %s", r.Where, r.Err)
}

type WriteError struct {
	Where string
	Err   error
}

func (w WriteError) Error() string {
	return fmt.Sprintf("encountered an error while trying to write a %s: %s", w.Where, w.Err)
}

type UnknownTag struct {
	TagId
}

func (u UnknownTag) Error() string {
	return fmt.Sprintf("discovered unknown TagId with id %d", u.TagId)
}

type WrongTag struct {
	Expecting, Got TagId
}

func (w WrongTag) Error() string {
	return fmt.Sprintf("expecting tag id %d, got %d", w.Expecting, w.Got)
}

type BadRange struct{}

func (BadRange) Error() string {
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
		err = &UnknownTag{id}
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
