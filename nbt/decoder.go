package nbt

import (
	"io"

	"github.com/MJKWoolnough/byteio"
)

// Decoder is a type used to decode NBT streams
type Decoder struct {
	r byteio.EndianReader
}

// NewDecoder returns a Decoder using Big Endian
func NewDecoder(r io.Reader) Decoder {
	return NewDecoderEndian(&byteio.BigEndianReader{Reader: r})
}

// NewDecoderEndian allows you to specify your own Endian Reader
func NewDecoderEndian(e byteio.EndianReader) Decoder {
	return Decoder{r: e}
}

// Decode will eencode a single tag from the reader using the default settings
func Decode(r io.Reader) (Tag, error) {
	return NewDecoder(r).Decode()
}

// Decode will read a whole tag out of the decoding stream
func (d Decoder) Decode() (Tag, error) {
	t, _, err := d.r.ReadUint8()
	if err != nil {
		err = ReadError{"named TagId", err}
		return Tag{}, err
	}
	tagID := TagID(t)
	if tagID == TagEnd {
		return Tag{data: end{}}, nil
	}
	n, err := d.decodeString()
	if err != nil {
		err = ReadError{"name", err}
		return Tag{}, err
	}
	data, err := d.decodeData(tagID)
	if err != nil {
		return Tag{}, err
	}
	return Tag{name: string(n), data: data}, nil
}

func (d Decoder) decodeData(tagID TagID) (Data, error) {
	var (
		data Data
		err  error
	)
	switch tagID {
	case TagByte:
		data, err = d.decodeByte()
	case TagShort:
		data, err = d.decodeShort()
	case TagInt:
		data, err = d.decodeInt()
	case TagLong:
		data, err = d.decodeLong()
	case TagFloat:
		data, err = d.decodeFloat()
	case TagDouble:
		data, err = d.decodeDouble()
	case TagByteArray:
		data, err = d.decodeByteArray()
	case TagString:
		data, err = d.decodeString()
	case TagList:
		data, err = d.decodeList()
	case TagCompound:
		data, err = d.decodeCompound()
	case TagIntArray:
		data, err = d.decodeIntArray()
	case TagBool:
		data, err = d.decodeBool()
	case TagUint8:
		data, err = d.decodeUint8()
	case TagUint16:
		data, err = d.decodeUint16()
	case TagUint32:
		data, err = d.decodeUint32()
	case TagUint64:
		data, err = d.decodeUint64()
	case TagComplex64:
		data, err = d.decodeComplex64()
	case TagComplex128:
		data, err = d.decodeComplex128()
	default:
		err = UnknownTag{tagID}
	}
	if err != nil {
		if _, ok := err.(ReadError); !ok {
			err = ReadError{tagID.String(), err}
		}
		return nil, err
	}
	return data, nil
}

// DecodeByte will read a single Byte Data
func (d Decoder) decodeByte() (Byte, error) {
	b, _, err := d.r.ReadInt8()
	return Byte(b), err
}

// DecodeShort will read a single Short Data
func (d Decoder) decodeShort() (Short, error) {
	s, _, err := d.r.ReadInt16()
	return Short(s), err
}

// DecodeInt will read a single Int Data
func (d Decoder) decodeInt() (Int, error) {
	i, _, err := d.r.ReadInt32()
	return Int(i), err
}

// DecodeLong will read a single Long Data
func (d Decoder) decodeLong() (Long, error) {
	l, _, err := d.r.ReadInt64()
	return Long(l), err
}

// DecodeFloat will read a single Float Data
func (d Decoder) decodeFloat() (Float, error) {
	f, _, err := d.r.ReadFloat32()
	return Float(f), err
}

// DecodeDouble will read a single Double Data
func (d Decoder) decodeDouble() (Double, error) {
	do, _, err := d.r.ReadFloat64()
	return Double(do), err
}

// DecodeByteArray will read a ByteArray Data
func (d Decoder) decodeByteArray() (ByteArray, error) {
	l, _, err := d.r.ReadUint32()
	if err != nil {
		return nil, err
	}
	data := make(ByteArray, l)
	err = data.readFrom(d.r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DecodeString will read a String Data
func (d Decoder) decodeString() (String, error) {
	l, _, err := d.r.ReadUint16()
	if err != nil {
		return "", err
	}
	data := make([]byte, l)
	_, err = io.ReadFull(d.r, data)
	if err != nil {
		return "", err
	}
	return String(data), nil
}

// DecodeList will read a List Data
func (d Decoder) decodeList() (List, error) {
	t, _, err := d.r.ReadUint8()
	if err != nil {
		return nil, err
	}
	tagID := TagID(t)
	length, _, err := d.r.ReadUint32()
	if err != nil {
		return nil, err
	}
	l := newListWithLength(tagID, length)
	var data Data
	for i := uint32(0); i < length; i++ {
		data, err = d.decodeData(tagID)
		if err != nil {
			return nil, err
		}
		l.Append(data)
	}
	return l, nil
}

// DecodeCompound will read a Compound Data
func (d Decoder) decodeCompound() (Compound, error) {
	data := make(Compound, 0)
	for {
		t, err := d.Decode()
		if err != nil {
			return nil, err
		}
		if t.TagID() == TagEnd {
			break
		}
		data = append(data, t)
	}
	return data, nil
}

// DecodeIntArray will read an IntArray Data
func (d Decoder) decodeIntArray() (IntArray, error) {
	l, _, err := d.r.ReadUint32()
	if err != nil {
		return nil, err
	}
	ints := make(IntArray, l)
	for i := uint32(0); i < l; i++ {
		ints[i], _, err = d.r.ReadInt32()
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func (d Decoder) decodeBool() (Bool, error) {
	b, _, err := d.r.ReadUint8()
	return b == 1, err
}

func (d Decoder) decodeUint8() (Uint8, error) {
	u, _, err := d.r.ReadUint8()
	if err != nil {
		return 0, err
	}
	return Uint8(u), err
}

func (d Decoder) decodeUint16() (Uint16, error) {
	u, _, err := d.r.ReadUint16()
	if err != nil {
		return 0, err
	}
	return Uint16(u), err
}

func (d Decoder) decodeUint32() (Uint32, error) {
	u, _, err := d.r.ReadUint32()
	if err != nil {
		return 0, err
	}
	return Uint32(u), err
}

func (d Decoder) decodeUint64() (Uint64, error) {
	u, _, err := d.r.ReadUint64()
	if err != nil {
		return 0, err
	}
	return Uint64(u), err
}

func (d Decoder) decodeComplex64() (Complex64, error) {
	r, _, err := d.r.ReadFloat32()
	if err != nil {
		return 0, err
	}
	i, _, err := d.r.ReadFloat32()
	if err != nil {
		return 0, err
	}
	return Complex64(complex(r, i)), nil
}

func (d Decoder) decodeComplex128() (Complex128, error) {
	r, _, err := d.r.ReadFloat64()
	if err != nil {
		return 0, err
	}
	i, _, err := d.r.ReadFloat64()
	if err != nil {
		return 0, err
	}
	return Complex128(complex(r, i)), nil
}
