package nbt

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/MJKWoolnough/byteio"
)

type rDecoder struct {
	Decoder
	io.Seeker
}

func (d Decoder) RDecode(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		return ErrInvalidValue
	}
	rd := rDecoder{d, makeSeeker(d.r)}
	t, _, err := d.r.ReadUint8()
	if err != nil {
		return err
	}
	tagID := TagID(t)
	if tagID != TagCompound {
		return ErrIncorrectType
	}
	_, err = d.decodeString()
	if err != nil {
		return err
	}
	if !checkType(tagID, rv.Kind()) {
		return ErrIncorrectType
	}
	return rd.decodeCompound(rv.Elem())
}

func (rd rDecoder) decodeData(tagID TagID, rv reflect.Value) error {
	switch tagID {
	case TagByte:
		data, err := rd.decodeByte()
		if err != nil {
			return err
		}
		rv.SetInt(int64(data))
	case TagShort:
		data, err := rd.decodeShort()
		if err != nil {
			return err
		}
		rv.SetInt(int64(data))
	case TagInt:
		data, err := rd.decodeInt()
		if err != nil {
			return err
		}
		rv.SetInt(int64(data))
	case TagLong:
		data, err := rd.decodeLong()
		if err != nil {
			return err
		}
		rv.SetInt(int64(data))
	case TagFloat:
		data, err := rd.decodeFloat()
		if err != nil {
			return err
		}
		rv.SetFloat(float64(data))
	case TagDouble:
		data, err := rd.decodeDouble()
		if err != nil {
			return err
		}
		rv.SetFloat(float64(data))
	case TagByteArray:
		data, err := rd.decodeByteArray()
		if err != nil {
			return err
		}
		if k := rv.Type().Elem().Kind(); k == reflect.Uint8 {
			rv.SetBytes(data.Bytes())
		} else if k == reflect.Int8 {
			rv.Set(reflect.ValueOf([]int8(data)))
		} else {
			return ErrInvalidType
		}
	case TagString:
		data, err := rd.decodeString()
		if err != nil {
			return err
		}
		rv.SetString(string(data))
	case TagList:
		if err := rd.decodeList(rv); err != nil {
			return err
		}
	case TagCompound:
		if err := rd.decodeCompound(rv); err != nil {
			return err
		}
	case TagIntArray:
		if rv.Type().Elem().Kind() != reflect.Int32 {
			return ErrInvalidType
		}
		data, err := rd.decodeIntArray()
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf([]int32(data)))
	case TagBool:
		data, err := rd.decodeBool()
		if err != nil {
			return err
		}
		rv.SetBool(bool(data))
	case TagUint8:
		data, err := rd.decodeUint8()
		if err != nil {
			return err
		}
		rv.SetUint(uint64(data))
	case TagUint16:
		data, err := rd.decodeUint16()
		if err != nil {
			return err
		}
		rv.SetUint(uint64(data))
	case TagUint32:
		data, err := rd.decodeUint32()
		if err != nil {
			return err
		}
		rv.SetUint(uint64(data))
	case TagUint64:
		data, err := rd.decodeUint64()
		if err != nil {
			return err
		}
		rv.SetUint(uint64(data))
	case TagComplex64:
		data, err := rd.decodeComplex64()
		if err != nil {
			return err
		}
		rv.SetComplex(complex128(data))
	case TagComplex128:
		data, err := rd.decodeComplex128()
		if err != nil {
			return err
		}
		rv.SetComplex(complex128(data))
	default:
		return UnknownTag{tagID}
	}
	return nil
}

var typeToKind = [...]reflect.Kind{
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Float32,
	reflect.Float64,
	reflect.Slice,
	reflect.String,
	reflect.Slice,
	reflect.Struct,
	reflect.Slice,
	reflect.Bool,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Complex64,
	reflect.Complex128,
}

func checkType(tagID TagID, rk reflect.Kind) bool {
	if int(tagID) > len(typeToKind) {
		return false
	}
	return typeToKind[tagID] == rk
}

type Seeker struct {
	io.Reader
}

func makeSeeker(r byteio.EndianReader) io.Seeker {
	if ber, ok := r.(byteio.BigEndianReader); ok {
		if s, ok := ber.Reader.(io.Seeker); ok {
			return s
		}
	} else if ler, ok := r.(byteio.LittleEndianReader); ok {
		if s, ok := ler.Reader.(io.Seeker); ok {
			return s
		}
	}
	return Seeker{r}
}

func (s Seeker) Seek(offset int64, whence int) (int64, error) {
	if whence != os.SEEK_CUR || whence < 0 {
		return 0, ErrUnsupportedWhence
	}
	return io.CopyN(ioutil.Discard, s, offset)
}

func (rd rDecoder) decodeList(rv reflect.Value) error {
	t, _, err := rd.r.ReadUint8()
	if err != nil {
		return err
	}
	tagID := TagID(t)
	if !checkType(tagID, rv.Type().Elem().Kind()) {
		return ErrInvalidType
	}
	l, _, err := rd.r.ReadUint32()
	if err != nil {
		return err
	}
	data := reflect.MakeSlice(rv.Elem().Type(), int(l), int(l))
	for i := uint32(0); i < l; i++ {
		err = rd.decodeData(tagID, data.Index(int(i)))
		if err != nil {
			return err
		}
	}
	rv.Set(data)
	return nil
}

func (rd rDecoder) decodeCompound(rv reflect.Value) error {
	for {
		t, _, err := rd.r.ReadUint8()
		if err != nil {
			return err
		}
		tagID := TagID(t)
		if tagID == TagEnd {
			return nil
		}
		n, err := rd.decodeString()
		if err != nil {
			return err
		}
		nv := getValueKey(rv, string(n))
		if nv.IsNil() {
			rd.skipData(tagID)
		} else {
			rd.decodeData(tagID, nv)
		}
	}
}

func getValueKey(rv reflect.Value, k string) reflect.Value {
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		f := t.Field(i)
		n := f.Name
		if st := f.Tag.Get("nbt"); st != "" {
			n = st
		}
		if n == k {
			tr := rv.Field(i)
			if tr.CanSet() {
				for tr.Kind() == reflect.Ptr {
					if tr.IsNil() {
						return reflect.Value{}
					}
					tr = tr.Elem()
				}
				return rv.Field(i)
			}
		}
	}
	return reflect.Value{}
}

func (rd rDecoder) skipData(tagID TagID) error {
	var err error
	switch tagID {
	case TagByte:
		_, err = rd.decodeByte()
	case TagShort:
		_, err = rd.decodeShort()
	case TagInt:
		_, err = rd.decodeInt()
	case TagLong:
		_, err = rd.decodeLong()
	case TagFloat:
		_, err = rd.decodeFloat()
	case TagDouble:
		_, err = rd.decodeDouble()
	case TagByteArray:
		err = rd.skipByteArray()
	case TagString:
		err = rd.skipString()
	case TagList:
		err = rd.skipList()
	case TagCompound:
		err = rd.skipCompound()
	case TagIntArray:
		err = rd.skipIntArray()
	case TagBool:
		_, err = rd.decodeBool()
	case TagUint8:
		_, err = rd.decodeUint8()
	case TagUint16:
		_, err = rd.decodeUint16()
	case TagUint32:
		_, err = rd.decodeUint32()
	case TagUint64:
		_, err = rd.decodeUint64()
	case TagComplex64:
		_, err = rd.decodeComplex64()
	case TagComplex128:
		_, err = rd.decodeComplex128()
	default:
		return UnknownTag{tagID}
	}
	return err
}

func (rd rDecoder) skipByteArray() error {
	l, _, err := rd.r.ReadUint32()
	if err != nil {
		return err
	}
	_, err = rd.Seek(int64(l), os.SEEK_CUR)
	return err
}

func (rd rDecoder) skipString() error {
	l, _, err := rd.r.ReadUint16()
	if err != nil {
		return err
	}
	_, err = rd.Seek(int64(l), os.SEEK_CUR)
	return err
}

func (rd rDecoder) skipList() error {
	t, _, err := rd.r.ReadUint8()
	if err != nil {
		return err
	}
	tagID := TagID(t)
	l, _, err := rd.r.ReadUint32()
	if err != nil {
		return err
	}
	var toSkip int64
	switch tagID {
	case TagEnd:
	case TagByte:
		toSkip = 1
	case TagShort:
		toSkip = 2
	case TagInt:
		toSkip = 4
	case TagLong:
		toSkip = 8
	case TagFloat:
		toSkip = 4
	case TagDouble:
		toSkip = 8
	case TagByteArray:
		for i := uint32(0); i < l; i++ {
			if err = rd.skipByteArray(); err != nil {
				return err
			}
		}
		return nil
	case TagString:
		for i := uint32(0); i < l; i++ {
			if err = rd.skipString(); err != nil {
				return err
			}
		}
		return nil
	case TagList:
		for i := uint32(0); i < l; i++ {
			if err = rd.skipList(); err != nil {
				return err
			}
		}
		return nil
	case TagCompound:
		for i := uint32(0); i < l; i++ {
			if err = rd.skipCompound(); err != nil {
				return err
			}
		}
		return nil
	case TagIntArray:
		for i := uint32(0); i < l; i++ {
			if err = rd.skipIntArray(); err != nil {
				return err
			}
		}
		return nil
	default:
		return UnknownTag{tagID}
	}
	_, err = rd.Seek(toSkip*int64(l), os.SEEK_CUR)
	return err
}

func (rd rDecoder) skipCompound() error {
	for {
		t, _, err := rd.r.ReadUint8()
		if err != nil {
			return ReadError{"named TagId", err}
		}
		tagID := TagID(t)
		if tagID == TagEnd {
			return nil
		}
		if err = rd.skipString(); err != nil {
			return err
		}
		if err = rd.skipData(tagID); err != nil {
			return err
		}
	}
}

func (rd rDecoder) skipIntArray() error {
	l, _, err := rd.r.ReadUint32()
	if err != nil {
		return err
	}
	_, err = rd.Seek(int64(l)*4, os.SEEK_CUR)
	return err
}

// Errors

var (
	ErrInvalidValue      = errors.New("invalid value type")
	ErrInvalidType       = errors.New("invalid type for NBT data")
	ErrUnsupportedWhence = errors.New("unsupported seek whence")
	ErrIncorrectType     = errors.New("incorrect type")
)
