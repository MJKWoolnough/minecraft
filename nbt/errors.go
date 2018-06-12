package nbt

import (
	"strconv"

	"vimagination.zapto.org/errors"
)

// ReadError is an error returned when a read error occurs
type ReadError struct {
	Where string
	Err   error
}

func (r ReadError) Error() string {
	return "encountered an error while trying to read a " + r.Where + ": " + r.Err.Error()
}

// WriteError is an error returned when a write error occurs
type WriteError struct {
	Where string
	Err   error
}

func (w WriteError) Error() string {
	return "encountered an error while trying to write a " + w.Where + ": " + w.Err.Error()
}

// UnknownTag is an error that occurs when an unknown tag id is discovered.
// This could also indicate corrupted or non-compliant data.
type UnknownTag struct {
	TagID
}

func (u UnknownTag) Error() string {
	return "discovered unknown TagId with id " + strconv.FormatUint(uint64(u.TagID), 10)
}

// WrongTag is an error returned when a tag of the incorrect type was intended
// to be added to a list.
type WrongTag struct {
	Expecting, Got TagID
}

func (w WrongTag) Error() string {
	return "expecting tag id " + w.Expecting.String() + ", got " + w.Got.String()
}

// ErrBadRange is an error that occurs when trying to set an item on a list which
// is outside of the current limits of the list.
const ErrBadRange errors.Error = "given index was out-of-range"
