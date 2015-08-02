package nbt

import (
	"errors"
	"io"
	"reflect"
)

type rEncoder struct {
	Encoder
}

// REncode will write the structure (with the given name) to the writer
func REncode(w io.Writer, name string, v interface{}) error {
	return NewEncoder(w).REncode(name, v)
}

// REncode will write the structure (with the given name) to the writer
func (e Encoder) REncode(name string, v interface{}) error {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return ErrIncorrectValue
	}
	re := rEncoder{e}
	_, err := re.w.WriteUint8(uint8(TagCompound))
	if err != nil {
		return err
	}
	err = re.encodeString(String(name))
	if err != nil {
		return err
	}
	return re.encodeCompound(rv)
}

//func (re rEncoder) encodeData(tagType TagID, rv reflect.Value) error // defined in either safe.go or unsafe.go

func (re rEncoder) encodeList(rv reflect.Value) error {
	tagType := tagFromType(rv.Type().Elem())
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
		t := rv.Index(i)
		for t.Kind() == reflect.Ptr {
			if t.IsNil() {
				return ErrInvalidPointer
			}
			t = t.Elem()
		}
		err = re.encodeData(tagType, t)
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
			for f.Kind() == reflect.Ptr {
				if f.IsNil() {
					return ErrInvalidPointer
				}
				f = f.Elem()
			}
			tf := rv.Type().Field(i)
			n := tf.Name
			if m := tf.Tag.Get("nbt"); n == "-" {
				continue
			} else if n != "" {
				n = m
			}
			tagType := tagFromType(f.Type())
			_, err := re.w.WriteUint8(uint8(tagType))
			if err != nil {
				return err
			}
			err = re.encodeString(String(n))
			if err != nil {
				return err
			}
			err = re.encodeData(tagType, f)
			if err != nil {
				return err
			}
		}
	}
	_, err := re.w.Write([]byte{byte(TagEnd)})
	return err
}

func tagFromType(rv reflect.Type) TagID {
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	switch rv {
	case nbtListType:
		return TagList
	case nbtCompoundType:
		return TagCompound
	default:
		switch rv.Kind() {
		case reflect.Int8:
			return TagByte
		case reflect.Int16:
			return TagShort
		case reflect.Int32:
			return TagInt
		case reflect.Int64, reflect.Int:
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
			switch rv.Elem().Kind() {
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
}

// Errors
var (
	ErrIncorrectValue = errors.New("incorrect value")
	ErrInvalidPointer = errors.New("invalid pointer")
)
