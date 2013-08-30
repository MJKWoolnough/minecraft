package minecraft

import (
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbt"
)

type MissingTagError struct {
	tagName string
}

func (m MissingTagError) Error() string {
	return fmt.Sprintf("minecraft: missing %q tag", m.tagName)
}

type WrongTypeError struct {
	tagName        string
	expecting, got nbt.TagId
}

func (w WrongTypeError) Error() string {
	return fmt.Sprintf("minecraft: tag %q is of incorrect type, expecting %q, got %q", w.tagName, w.expecting, w.got)
}

type OOB struct{}

func (o OOB) Error() string {
	return "Received Out-of-bounds error"
}

type UnexpectedValue struct {
	tagName, expecting, got string
}

func (u UnexpectedValue) Error() string {
	return fmt.Sprintf("minecraft: tag %q was expecting %s, but got %q", u.tagName, u.expecting, u.got)
}

type UnknownCompression struct {
	code byte
}

func (u UnknownCompression) Error() string {
	return fmt.Sprintf("unknown compression code: %d", u.code)
}

type ExpectedData struct{}

func (e ExpectedData) Error() string {
	return "received nil data stream where non-nil was expected."
}

type NoLock struct{}

func (n NoLock) Error() string {
	return "lost lock on files"
}

type ConflictError struct {
	X, Z int32
}

func (c ConflictError) Error() string {
	return fmt.Sprintf("already setting chunk %d,%d, recall SetChunk to set again", c.X, c.Z)
}
