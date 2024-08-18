//go:build js
// +build js

package nbt

import "io"

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
