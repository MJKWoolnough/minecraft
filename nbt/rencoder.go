package nbt

import (
	"errors"
	"reflect"
	"unsafe"
)

type rEncoder struct {
	Encoder
}

func (e Encoder) REncode(v interface{}) error {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct || rv.IsNil() {
		return ErrIncorrectValue
	}
	re := rEncoder{e}
	_, err := re.w.WriteUint8(uint8(TagCompound))
	if err != nil {
		return err
	}
	err = re.encodeString("")
	if err != nil {
		return err
	}
	return re.encodeCompound(rv)
}

func (re rEncoder) encodeData(tagType TagID, rv reflect.Value) error {
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
		return re.encodeByteArray(*(*[]int8)(unsafe.Pointer(ptr)))
	case TagString:
		return re.encodeString(*(*String)(unsafe.Pointer(ptr)))
	case TagList:
		return re.encodeList(rv)
	case TagCompound:
		return re.encodeCompound(rv)
	case TagIntArray:
		return re.encodeIntArray(*(*[]int32)(unsafe.Pointer(ptr)))
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

func (re rEncoder) encodeList(rv reflect.Value) error {
	tagType := tagFromValue(rv.Elem())
	_, err := re.w.WriteUint8(uint8(tagType))
	if err != nil {
		return err
	}
	l := rv.Len()
	_, err = re.w.WriteUint32(uint32(l))
	if err != nil {
		return err
	}
	for i := 0; i < l; i++ {
		err = re.encodeData(tagType, rv.Index(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func (re rEncoder) encodeCompound(rv reflect.Value) error {
	for i := 0; i < rv.Type().NumField(); i++ {
		f := rv.Field(i)
		if f.CanSet() {
			tf := rv.Type().Field(i)
			n := tf.Name
			if m := tf.Tag.Get("nbt"); n == "-" {
				continue
			} else if n != "" {
				n = m
			}
			tagType := tagFromValue(f)
			_, err := re.w.WriteUint8(uint8(tagType))
			if err != nil {
				return err
			}
			err = re.encodeString(String(n))
			if err != nil {
				return err
			}
			err = re.encodeData(tagType, rv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func tagFromValue(rv reflect.Value) TagID {
	switch rv.Kind() {
	case reflect.Int8:
		return TagByte
	case reflect.Int16:
		return TagShort
	case reflect.Int32:
		return TagInt
	case reflect.Int64:
		return TagLong
	case reflect.Float32:
		return TagFloat
	case reflect.Float64:
		return TagDouble
	case reflect.String:
		return TagString
	case reflect.Struct:
		return TagCompound
	case reflect.Bool:
		return TagBool
	case reflect.Uint8:
		return TagUint8
	case reflect.Uint16:
		return TagUint16
	case reflect.Uint32:
		return TagUint32
	case reflect.Uint64:
		return TagUint64
	case reflect.Complex64:
		return TagComplex64
	case reflect.Complex128:
		return TagComplex128
	case reflect.Slice, reflect.Array:
		switch rv.Type().Elem().Kind() {
		case reflect.Int8:
			return TagByteArray
		case reflect.Int32:
			return TagIntArray
		default:
			return TagList
		}
	default:
		return TagEnd
	}
}

// Errors

var ErrIncorrectValue = errors.New("incorrect value")
