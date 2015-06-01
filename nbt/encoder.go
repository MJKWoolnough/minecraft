package nbt

import (
	"io"

	"github.com/MJKWoolnough/byteio"
)

// Encoder is a type used to encode NBT streams
type Encoder struct {
	w byteio.EndianWriter
}

// NewEncoder returns an Encoder using Big Endian
func NewEncoder(w io.Writer) Encoder {
	return NewEncoderEndian(&byteio.BigEndianWriter{Writer: w})
}

// NewEncoderEndian allows you to specify your own Endian Writer
func NewEncoderEndian(e byteio.EndianWriter) Encoder {
	return Encoder{w: e}
}

// Encode will encode a single tag to the writer using the default settings
func Encode(w io.Writer, t Tag) error {
	return NewEncoder(w).Encode(t)
}

// Encode will encode a whole tag to the encoding stream
func (e Encoder) Encode(t Tag) error {
	tagType := t.TagID()
	_, err := e.w.WriteUint8(uint8(tagType))
	if err != nil {
		return WriteError{"named TagId", err}
	}
	if tagType == TagEnd {
		return nil
	}
	s := String(t.name)
	err = e.encodeString(s)
	if err != nil {
		return err
	}
	return e.encodeData(t.data)
}

func (e Encoder) encodeData(d Data) error {
	var err error
	switch d := d.(type) {
	case Byte:
		err = e.encodeByte(d)
	case Short:
		err = e.encodeShort(d)
	case Int:
		err = e.encodeInt(d)
	case Long:
		err = e.encodeLong(d)
	case Float:
		err = e.encodeFloat(d)
	case Double:
		err = e.encodeDouble(d)
	case ByteArray:
		err = e.encodeByteArray(d)
	case String:
		err = e.encodeString(d)
	case *List:
		err = e.encodeList(d)
	case Compound:
		err = e.encodeCompound(d)
	case IntArray:
		err = e.encodeIntArray(d)
	default:
		err = UnknownTag{d.Type()}
	}
	if err != nil {
		return err
	}
	return nil
}

// EncodeByte will write a single Byte Data
func (e Encoder) encodeByte(b Byte) error {
	_, err := e.w.WriteInt8(int8(b))
	return err
}

// EncodeShort will write a single Short Data
func (e Encoder) encodeShort(s Short) error {
	_, err := e.w.WriteInt16(int16(s))
	return err
}

// EncodeInt will write a single Int Data
func (e Encoder) encodeInt(i Int) error {
	_, err := e.w.WriteInt32(int32(i))
	return err
}

// EncodeLong will write a single Long Data
func (e Encoder) encodeLong(l Long) error {
	_, err := e.w.WriteInt64(int64(l))
	return err
}

// EncodeFloat will write a single Float Data
func (e Encoder) encodeFloat(f Float) error {
	_, err := e.w.WriteFloat32(float32(f))
	return err
}

// EncodeDouble will write a single Double Data
func (e Encoder) encodeDouble(do Double) error {
	_, err := e.w.WriteFloat64(float64(do))
	return err
}

// EncodeByteArray will write a ByteArray Data
func (e Encoder) encodeByteArray(ba ByteArray) error {
	_, err := e.w.WriteUint32(uint32(len(ba)))
	if err != nil {
		return err
	}
	_, err = e.w.Write(ba.Bytes())
	return err
}

// EncodeString will write a String Data
func (e Encoder) encodeString(s String) error {
	_, err := e.w.WriteUint16(uint16(len(s)))
	if err != nil {
		return err
	}
	_, err = e.w.Write([]byte(s))
	return err
}

// EncodeList will write a List Data
func (e Encoder) encodeList(l *List) error {
	_, err := e.w.WriteUint8(uint8(l.tagType))
	if err != nil {
		return err
	}
	_, err = e.w.WriteUint32(uint32(len(l.data)))
	if err != nil {
		return err
	}
	if l.TagType() != TagEnd {
		for _, data := range l.data {
			if tagID := data.Type(); tagID != l.tagType {
				return WrongTag{l.tagType, tagID}
			}
			err = e.encodeData(data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// EncodeCompound will write a Compound Data
func (e Encoder) encodeCompound(c Compound) error {
	for _, data := range c {
		if data.TagID() == TagEnd {
			break
		}
		err := e.Encode(data)
		if err != nil {
			return err
		}
	}
	_, err := e.w.Write([]byte{byte(TagEnd)})
	return err
}

// EncodeIntArray will write a IntArray Data
func (e Encoder) encodeIntArray(ints IntArray) error {
	_, err := e.w.WriteUint32(uint32(len(ints)))
	for _, i := range ints {
		_, err = e.w.WriteInt32(i)
		if err != nil {
			return err
		}
	}
	return nil
}