package nbt

import "strconv"

// ListEnd satisfies the List interface for a list of Ends.
type ListEnd uint32

// Equal satisfies the equaler.Equaler interface, allowing for types to be
// checked for equality.
func (l ListEnd) Equal(e interface{}) bool {
	m, ok := e.(List)

	return ok && m.TagType() == TagEnd && ListEnd(m.Len()) == l
}

// Copy simply returns a deep-copy the the data.
func (l ListEnd) Copy() Data {
	return &l
}

func (l ListEnd) String() string {
	return strconv.Itoa(int(l)) + " entries of type End"
}

// Type returns the TagID of the data.
func (ListEnd) Type() TagID {
	return TagList
}

// TagType returns the TagID of the type of tag this list contains.
func (ListEnd) TagType() TagID {
	return TagEnd
}

// Set does nothing as it's not applicable for ListEnd.
func (l ListEnd) Set(_ int, d Data) error {
	return nil
}

// Get returns an end{}.
func (ListEnd) Get(_ int) Data {
	return end{}
}

// Append adds to the list.
func (l *ListEnd) Append(d ...Data) error {
	*l += ListEnd(len(d))

	return nil
}

// Insert adds to the list.
func (l *ListEnd) Insert(_ int, d ...Data) error {
	*l += ListEnd(len(d))

	return nil
}

// Remove removes from the list.
func (l *ListEnd) Remove(i int) {
	if ListEnd(i) < *l {
		*l--
	}
}

// Len returns the length of the List.
func (l ListEnd) Len() int {
	return int(l)
}

// ListEnd returns the list as a specifically typed List.
func (l ListData) ListEnd() ListEnd {
	if l.tagType != TagEnd {
		return 0
	}

	return ListEnd(len(l.data))
}
