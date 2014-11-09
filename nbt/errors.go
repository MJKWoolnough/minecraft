package nbt

import "fmt"

type ReadError struct {
	Where string
	Err   error
}

func (r ReadError) Error() string {
	return "encountered an error while trying to read a " + r.Where + ": " + r.Err.Error()
}

type WriteError struct {
	Where string
	Err   error
}

func (w WriteError) Error() string {
	return "encountered an error while trying to write a " + w.Where + ": " + w.Err.Error()
}

type UnknownTag struct {
	TagId
}

func (u UnknownTag) Error() string {
	return fmt.Sprintf("discovered unknown TagId with id %d", u.TagId)
}

type WrongTag struct {
	Expecting, Got TagId
}

func (w WrongTag) Error() string {
	return fmt.Sprintf("expecting tag id %d, got %d", w.Expecting, w.Got)
}

type BadRange struct{}

func (BadRange) Error() string {
	return "given index was out-of-range"
}
