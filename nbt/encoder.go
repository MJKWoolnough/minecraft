package nbt

import (
	"io"

	"github.com/MJKWoolnough/byteio"
)

type Encoder struct {
	w byteio.EndianWriter
}

func NewEncoder(w io.Writer) Encoder {
	return NewEncoderEndian(&byteio.BigEndianWriter{Writer: w})
}

func NewEncoderEndian(e byteio.EndianWriter) Encoder {
	return Encoder{w: e}
}

func (e Encoder) EncodeTag(t Tag) error {
	tagType := t.data.Type()
	_, err := e.w.WriteUint8(uint8(tagType))
	if err != nil {
		return WriteError{"named TagId", err}
	}
	if tagType == TagEnd {
		return nil
	}
	s := String(t.name)
	err = e.EncodeString(s)
	if err != nil {
		return err
	}
	return e.encodeData(t.data)
}

func (e Encoder) encodeData(d Data) error {
	var err error
	switch d := d.(type) {
	case Byte:
		err = e.EncodeByte(d)
	case Short:
		err = e.EncodeShort(d)
	case Int:
		err = e.EncodeInt(d)
	case Long:
		err = e.EncodeLong(d)
	case Float:
		err = e.EncodeFloat(d)
	case Double:
		err = e.EncodeDouble(d)
	case ByteArray:
		err = e.EncodeByteArray(d)
	case String:
		err = e.EncodeString(d)
	case *List:
		err = e.EncodeList(d)
	case Compound:
		err = e.EncodeCompound(d)
	case IntArray:
		err = e.EncodeIntArray(d)
	default:
		err = UnknownTag{d.Type()}
	}
	if err != nil {
		return err
	}
	return nil
}

func (e Encoder) EncodeByte(b Byte) error {
	_, err := e.w.WriteInt8(int8(b))
	return err
}

func (e Encoder) EncodeShort(s Short) error {
	_, err := e.w.WriteInt16(int16(s))
	return err
}

func (e Encoder) EncodeInt(i Int) error {
	_, err := e.w.WriteInt32(int32(i))
	return err
}

func (e Encoder) EncodeLong(l Long) error {
	_, err := e.w.WriteInt64(int64(l))
	return err
}

func (e Encoder) EncodeFloat(f Float) error {
	_, err := e.w.WriteFloat32(float32(f))
	return err
}

func (e Encoder) EncodeDouble(do Double) error {
	_, err := e.w.WriteFloat64(float64(do))
	return err
}

func (e Encoder) EncodeByteArray(ba ByteArray) error {
	_, err := e.w.WriteUint32(uint32(len(ba)))
	if err != nil {
		return err
	}
	data := make([]byte, len(ba))
	for i := 0; i < len(ba); i++ {
		data[i] = byte(ba[i])
	}
	_, err = e.w.Write(data)
	return err
}

func (e Encoder) EncodeString(s String) error {
	_, err := e.w.WriteUint16(uint16(len(s)))
	if err != nil {
		return err
	}
	_, err = e.w.Write([]byte(s))
	return err
}

func (e Encoder) EncodeList(l *List) error {
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
				return &WrongTag{l.tagType, tagID}
			}
			err = e.encodeData(data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e Encoder) EncodeCompound(c Compound) error {
	for _, data := range c {
		if data.TagID() == TagEnd {
			break
		}
		err := e.EncodeTag(data)
		if err != nil {
			return err
		}
	}
	_, err := e.w.Write([]byte{byte(TagEnd)})
	return err
}

func (e Encoder) EncodeIntArray(ints IntArray) error {
	_, err := e.w.WriteUint32(uint32(len(ints)))
	for _, i := range ints {
		_, err = e.w.WriteInt32(i)
		if err != nil {
			return err
		}
	}
	return nil
}
