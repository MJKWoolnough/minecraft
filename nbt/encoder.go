package nbt // import "vimagination.zapto.org/minecraft/nbt"

import (
	"io"

	"vimagination.zapto.org/byteio"
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
	case *ListData:
		err = e.encodeList(d)
	case Compound:
		err = e.encodeCompound(d)
	case IntArray:
		err = e.encodeIntArray(d)
	case Bool:
		err = e.encodeBool(d)
	case Uint8:
		err = e.encodeUint8(d)
	case Uint16:
		err = e.encodeUint16(d)
	case Uint32:
		err = e.encodeUint32(d)
	case Uint64:
		err = e.encodeUint64(d)
	case Complex64:
		err = e.encodeComplex64(d)
	case Complex128:
		err = e.encodeComplex128(d)
	default:
		if l, ok := d.(List); ok {
			e.encodeList(l)
		} else {
			err = UnknownTag{d.Type()}
		}
	}
	return err
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
	_, err := e.w.WriteString16(string(s))
	return err
}

// EncodeList will write a List Data
func (e Encoder) encodeList(l List) error {
	tagType := l.TagType()
	_, err := e.w.WriteUint8(uint8(tagType))
	if err != nil {
		return err
	}
	_, err = e.w.WriteUint32(uint32(l.Len()))
	if err != nil {
		return err
	}
	if l.TagType() != TagEnd {
		for i := 0; i < l.Len(); i++ {
			data := l.Get(i)
			if tagID := data.Type(); tagID != tagType {
				return WrongTag{tagType, tagID}
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
	if err != nil {
		return err
	}
	for _, i := range ints {
		_, err = e.w.WriteInt32(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Encoder) encodeBool(b Bool) error {
	var err error
	if b {
		_, err = e.w.WriteUint8(1)
	} else {
		_, err = e.w.WriteUint8(0)
	}
	return err
}

func (e Encoder) encodeUint8(u Uint8) error {
	_, err := e.w.WriteUint8(uint8(u))
	return err
}

func (e Encoder) encodeUint16(u Uint16) error {
	_, err := e.w.WriteUint16(uint16(u))
	return err
}

func (e Encoder) encodeUint32(u Uint32) error {
	_, err := e.w.WriteUint32(uint32(u))
	return err
}

func (e Encoder) encodeUint64(u Uint64) error {
	_, err := e.w.WriteUint64(uint64(u))
	return err
}

func (e Encoder) encodeComplex64(c Complex64) error {
	_, err := e.w.WriteFloat32(real(c))
	if err != nil {
		return err
	}
	_, err = e.w.WriteFloat32(imag(c))
	return err
}

func (e Encoder) encodeComplex128(c Complex128) error {
	_, err := e.w.WriteFloat64(real(c))
	if err != nil {
		return err
	}
	_, err = e.w.WriteFloat64(imag(c))
	return err
}
