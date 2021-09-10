package minecraft

import (
	"errors"
	"strconv"

	"vimagination.zapto.org/minecraft/nbt"
)

var (
	// ErrOOB is an error returned when sanity checking section data
	ErrOOB = errors.New("received out-of-bounds error")
	// ErrNoLock is an error returns by path types to indicate that the lock on the
	// minecraft level has been locked and needs reinstating to continue
	ErrNoLock = errors.New("lost lock on files")
)

// MissingTagError is an error type returned when an expected tag is not found
type MissingTagError struct {
	TagName string
}

func (m MissingTagError) Error() string {
	return "minecraft: missing " + strconv.Quote(m.TagName) + " tag"
}

// WrongTypeError is an error returned when a nbt tag has an unexpected type
type WrongTypeError struct {
	TagName        string
	Expecting, Got nbt.TagID
}

func (w WrongTypeError) Error() string {
	return "minecraft: tag " + strconv.Quote(w.TagName) + " is of incorrect type, expecting " + strconv.Quote(w.Expecting.String()) + ", got " + strconv.Quote(w.Got.String())
}

// UnexpectedValue is an error returned from chunk loading during sanity
// checking
type UnexpectedValue struct {
	TagName, Expecting, Got string
}

func (u UnexpectedValue) Error() string {
	return "minecraft: tag " + strconv.Quote(u.TagName) + " was expecting " + strconv.Quote(u.Expecting) + ", but got " + strconv.Quote(u.Got)
}

// UnknownCompression is an error returned by path types when it encounters a
// compression scheme it is not prepared to handle or an unkown compression
// scheme
type UnknownCompression struct {
	Code byte
}

func (u UnknownCompression) Error() string {
	return "unknown compression code: " + strconv.FormatUint(uint64(u.Code), 10)
}

// ConflictError is an error return by SetChunk when trying to save a single
// chunk multiple times during the same save operation
type ConflictError struct {
	X, Z int32
}

func (c ConflictError) Error() string {
	return "already setting chunk " + strconv.FormatInt(int64(c.X), 10) + "," + strconv.FormatInt(int64(c.Z), 10) + ", recall SetChunk to set again"
}

// FilePathSetError is an error returned from SetChunk when some error is
// returned either from converting the nbt or saving it
type FilePathSetError struct {
	X, Z int32
	Err  error
}

func (f FilePathSetError) Error() string {
	return "chunk " + strconv.FormatInt(int64(f.X), 10) + "," + strconv.FormatInt(int64(f.Z), 10) + " had the following error: " + f.Err.Error()
}

// MultiError is an error type that contains multiple errors
type MultiError struct {
	Errors []error
}

func (m MultiError) Error() string {
	if len(m.Errors) == 1 {
		return m.Errors[0].Error()
	}
	return "received " + strconv.FormatInt(int64(len(m.Errors)), 10) + " errors"
}
