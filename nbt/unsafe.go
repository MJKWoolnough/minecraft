// +build !js

package nbt

import (
	"io"
	"unsafe"
)

func byteArrayToByteSlice(s ByteArray) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func (b ByteArray) readFrom(r io.Reader) error {
	_, err := io.ReadFull(r, byteArrayToByteSlice(b))
	return err
}
