package nbt

import (
	"fmt"
	"io"
)

type countReader struct {
	io.Reader
	bytesRead int64
}

func (c *countReader) Read(d []byte) (total int, err error) {
	total, err = c.Reader.Read(d)
	c.bytesRead += int64(total)
	return
}

func (c countReader) BytesRead(d *int64) {
	*d = c.bytesRead
}

type countWriter struct {
	io.Writer
	bytesWritten int64
}

func (c *countWriter) Write(d []byte) (total int, err error) {
	total, err = c.Writer.Write(d)
	c.bytesWritten += int64(total)
	return
}

func (c countWriter) BytesWritten(d *int64) {
	*d = c.bytesWritten
}

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
			out += s[last:i] + "	"
			last = i
		}
	}
	out += s[last:]
	return s
}
