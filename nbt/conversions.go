package nbt

//copied/modified from library packages math and encoding/binary

import "unsafe"

func float32bits(f float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&f))
}

func float32frombits(b uint32) float32 {
	return *(*float32)(unsafe.Pointer(&b))
}

func float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func float64frombits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}

func getUint16(b []byte) uint16 {
	return uint16(b[1]) | uint16(b[0])<<8
}

func putUint16(v uint16) []byte {
	var b [2]byte
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	return b[:]
}

func getUint32(b []byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func putUint32(v uint32) []byte {
	var b [4]byte
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	return b[:]
}

func getUint64(b []byte) uint64 {
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 | uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}

func putUint64(v uint64) []byte {
	var b [8]byte
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b[:]
}
