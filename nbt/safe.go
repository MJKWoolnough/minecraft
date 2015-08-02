// +build js

package nbt

import (
	"io"
	"reflect"
)

func byteArrayToByteSlice(s ByteArray) []byte {
	i := make([]byte, len(s))
	for n, b := range s {
		i[n] = byte(b)
	}
	return i
}

func (b ByteArray) readFrom(r io.Reader) error {
	bytes := make([]byte, len(b))
	_, err := io.ReadFull(r, bytes)
	for n, by := range bytes {
		b[n] = int8(by)
	}
	return err
}

func (re rEncoder) encodeData(tagType TagID, rv reflect.Value) error {
	switch t := rv.Interface().(type) {
	case List:
		return re.Encoder.encodeList(&t)
	case Compound:
		return re.Encoder.encodeCompound(t)
	default:
		switch tagType {
		case TagByte:
			return re.encodeByte(Byte(rv.Int()))
		case TagShort:
			return re.encodeShort(Short(rv.Int()))
		case TagInt:
			return re.encodeInt(Int(rv.Int()))
		case TagLong:
			return re.encodeLong(Long(rv.Int()))
		case TagFloat:
			return re.encodeFloat(Float(rv.Float()))
		case TagDouble:
			return re.encodeDouble(Double(rv.Float()))
		case TagByteArray:
			bytes := make(ByteArray, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				bytes[i] = int8(rv.Index(i).Int())
			}
			return re.encodeByteArray(ByteArray(bytes))
		case TagString:
			return re.encodeString(String(rv.String()))
		case TagList:
			return re.encodeList(rv)
		case TagCompound:
			return re.encodeCompound(rv)
		case TagIntArray:
			ints := make(IntArray, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				ints[i] = int32(rv.Index(i).Int())
			}
			return re.encodeIntArray(ints)
		case TagBool:
			return re.encodeBool(Bool(rv.Bool()))
		case TagUint8:
			return re.encodeUint8(Uint8(rv.Uint()))
		case TagUint16:
			return re.encodeUint16(Uint16(rv.Uint()))
		case TagUint32:
			return re.encodeUint32(Uint32(rv.Uint()))
		case TagUint64:
			return re.encodeUint64(Uint64(rv.Uint()))
		case TagComplex64:
			return re.encodeComplex64(Complex64(rv.Complex()))
		case TagComplex128:
			return re.encodeComplex128(Complex128(rv.Complex()))
		default:
			return UnknownTag{tagType}
		}
	}
}
