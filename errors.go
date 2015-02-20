package minecraft

import (
	"fmt"

	"github.com/MJKWoolnough/minecraft/nbt"
)

// MissingTagError is an error type returned when an expected tag is not found
type MissingTagError struct {
	TagName string
}

func (m *MissingTagError) Error() string {
	return fmt.Sprintf("minecraft: missing %q tag", m.TagName)
}

// WrongTypeError is an error returned when a nbt tag has an unexpected type
type WrongTypeError struct {
	TagName        string
	Expecting, Got nbt.TagID
}

func (w *WrongTypeError) Error() string {
	return fmt.Sprintf("minecraft: tag %q is of incorrect type, expecting %q, got %q", w.TagName, w.Expecting, w.Got)
}

// OOB is an error returned when sanity checking section data
type OOB struct{}

func (OOB) Error() string {
	return "Received Out-of-bounds error"
}

// UnexpectedValue is an error returned from chunk loading during sanity
// checking
type UnexpectedValue struct {
	TagName, Expecting, Got string
}

func (u *UnexpectedValue) Error() string {
	return fmt.Sprintf("minecraft: tag %q was expecting %s, but got %q", u.TagName, u.Expecting, u.Got)
}

// UnknownCompression is an error returned by path types when it encounters a
// compression scheme it is not prepared to handle or an unkown compression
// scheme
type UnknownCompression struct {
	Code byte
}

func (u UnknownCompression) Error() string {
	return fmt.Sprintf("unknown compression code: %d", u.Code)
}

// NoLock is an error returns by path types to indicate that the lock on the
// minecraft level has been locked and needs reinstating to continue
type NoLock struct{}

func (NoLock) Error() string {
	return "lost lock on files"
}

// ConflictError is an error return by SetChunk when trying to save a single
// chunk multiple times during the same save operation
type ConflictError struct {
	X, Z int32
}

func (c *ConflictError) Error() string {
	return fmt.Sprintf("already setting chunk %d,%d, recall SetChunk to set again", c.X, c.Z)
}

// FilePathSetError is an error returned from SetChunk when some error is
// returned either from converting the nbt or saving it
type FilePathSetError struct {
	X, Z int32
	Err  error
}

func (f *FilePathSetError) Error() string {
	return fmt.Sprintf("chunk %d, %d had the following error: %s", f.X, f.Z, f.Err.Error())
}

// MultiError is an error type that contains multiple errors
type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	if len(m.Errors) == 1 {
		return m.Errors[0].Error()
	}
	return fmt.Sprintf("received %d errors", len(m.Errors))
}
