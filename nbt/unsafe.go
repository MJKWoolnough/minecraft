// +build !js

package nbt

import (
	"io"
	"reflect"
	"unsafe"
)

func byteArrayToByteSlice(s ByteArray) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func (b ByteArray) readFrom(r io.Reader) error {
	_, err := io.ReadFull(r, byteArrayToByteSlice(b))
	return err
}

func (re rEncoder) encodeData(tagType TagID, rv reflect.Value) error {
	switch t := rv.Interface().(type) {
	case List:
		return re.Encoder.encodeList(&t)
	case Compound:
		return re.Encoder.encodeCompound(t)
	default:
		ptr := rv.UnsafeAddr()
		switch tagType {
		case TagByte:
			return re.encodeByte(*(*Byte)(unsafe.Pointer(ptr)))
		case TagShort:
			return re.encodeShort(*(*Short)(unsafe.Pointer(ptr)))
		case TagInt:
			return re.encodeInt(*(*Int)(unsafe.Pointer(ptr)))
		case TagLong:
			return re.encodeLong(*(*Long)(unsafe.Pointer(ptr)))
		case TagFloat:
			return re.encodeFloat(*(*Float)(unsafe.Pointer(ptr)))
		case TagDouble:
			return re.encodeDouble(*(*Double)(unsafe.Pointer(ptr)))
		case TagByteArray:
			return re.encodeByteArray(ByteArray(*(*[]int8)(unsafe.Pointer(ptr))))
		case TagString:
			return re.encodeString(*(*String)(unsafe.Pointer(ptr)))
		case TagList:
			return re.encodeList(rv)
		case TagCompound:
			return re.encodeCompound(rv)
		case TagIntArray:
			return re.encodeIntArray(IntArray(*(*[]int32)(unsafe.Pointer(ptr))))
		case TagBool:
			return re.encodeBool(*(*Bool)(unsafe.Pointer(ptr)))
		case TagUint8:
			return re.encodeUint8(*(*Uint8)(unsafe.Pointer(ptr)))
		case TagUint16:
			return re.encodeUint16(*(*Uint16)(unsafe.Pointer(ptr)))
		case TagUint32:
			return re.encodeUint32(*(*Uint32)(unsafe.Pointer(ptr)))
		case TagUint64:
			return re.encodeUint64(*(*Uint64)(unsafe.Pointer(ptr)))
		case TagComplex64:
			return re.encodeComplex64(*(*Complex64)(unsafe.Pointer(ptr)))
		case TagComplex128:
			return re.encodeComplex128(*(*Complex128)(unsafe.Pointer(ptr)))
		default:
			return UnknownTag{tagType}
		}
	}
}
