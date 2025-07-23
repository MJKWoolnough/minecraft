#!/bin/bash

types=( Byte Short Int Long Float Double Compound IntArray Bool Uint8 Uint16 Uint32 Uint64 Complex64 Complex128 );

{
	echo "// file automatically generated with listGen.sh.";
	echo;
	echo "package nbt";
	echo;
	echo "import \"strconv\"";

	for type in ${types[@]}; do
		echo;

		echo "// List$type satisfies the List interface for a list of ${type}s.";
		echo "type List$type []$type";

		echo;

		echo "// Equal satisfies the equaler.Equaler interface, allowing for types to be";
		echo "// checked for equality";
		echo "func (l List$type) Equal(e interface{}) bool {";
		echo "	m, ok := e.(List$type)";
		echo "	if !ok {";
		echo "		var n *List$type";
		echo;
		echo "		if n, ok = e.(*List$type); ok {";
		echo "			m = *n";
		echo "		}";
		echo "	}";
		echo;
		echo "	if ok {";
		echo "		if len(l) == len(m) {";
		echo "			for n, t := range m {";
		echo "				if !t.Equal(l[n]) {";
		echo "					return false";
		echo "				}";
		echo "			}";
		echo;
		echo "			return true";
		echo "		}";
		echo "	} else if d, ok := e.(List); ok && d.TagType() == Tag$type && d.Len() == len(l) {";
		echo "		for i := 0; i < d.Len(); i++ {";
		echo "			if !d.Get(i).Equal(l[i]) {";
		echo "				return false";
		echo "			}";
		echo "		}";
		echo;
		echo "		return true";
		echo "	}";
		echo;
		echo "	return false";
		echo "}";

		echo;

		echo "// Copy simply returns a deep-copy of the data.";
		echo "func (l List$type) Copy() Data {";
		echo "	m := make(List$type, len(l))";
		echo "	for n, e := range l {";
		echo "		m[n] = e.Copy().($type)";
		echo "	}";
		echo;
		echo "	return &m";
		echo "}";

		echo;

		echo "func (l List$type) String() string {";
		echo "	s := strconv.Itoa(len(l)) + \" entries of type $type {\"";
		echo;
		echo "	for _, d := range l {";
		echo "		s += \"\\n        $type: \" + indent(d.String())";
		echo "	}";
		echo;
		echo "	return s + \"\\n}\"";
		echo "}";

		echo;

		echo "// Type returns the TagID of the data.";
		echo "func (List$type) Type() TagID {";
		echo "	return TagList";
		echo "}";

		echo;

		echo "// TagType returns the TagID of the type of tag this list contains.";
		echo "func (List$type) TagType() TagID {";
		echo "	return Tag$type";
		echo "}";

		echo;

		echo "// Set sets the data at the given position. It does not append.";
		echo "func (l List$type) Set(i int, d Data) error {";
		echo "	if m, ok := d.($type); ok {";
		echo "		if i <= 0 || i >= len(l) {";
		echo "			return ErrBadRange";
		echo "		}";
		echo;
		echo "		l[i] = m";
		echo "	} else {";
		echo "		return &WrongTag{Tag$type, d.Type()}";
		echo "	}";
		echo;
		echo "	return nil";
		echo "}";

		echo;

		echo "// Get returns the data at the given position.";
		echo "func (l List$type) Get(i int) Data {";
		echo "	return l[i]";
		echo "}";

		echo;

		echo "// Append adds data to the list";
		echo "func (l *List$type) Append(d ...Data) error {";
		echo "	toAppend := make(List$type, len(d))";
		echo;
		echo "	for n, e := range d {";
		echo "		if f, ok := e.($type); ok {";
		echo "			toAppend[n] = f";
		echo "		} else {";
		echo "			return &WrongTag{Tag$type, e.Type()}";
		echo "		}";
		echo "	}";
		echo;
		echo "	*l = append(*l, toAppend...)";
		echo;
		echo "	return nil";
		echo "}";

		echo;

		echo "// Insert will add the given data at the specified position, moving other";
		echo "// up.";
		echo "func (l *List$type) Insert(i int, d ...Data) error {";
		echo "	if i >= len(*l) {";
		echo "		return l.Append(d...)";
		echo "	}";
		echo;
		echo "	toInsert := make(List$type, len(d), len(d)+len(*l)-i)";
		echo;
		echo "	for n, e := range d {";
		echo "		if f, ok := e.($type); ok {";
		echo "			toInsert[n] = f";
		echo "		} else {";
		echo "			return &WrongTag{Tag$type, e.Type()}";
		echo "		}";
		echo "	}";
		echo;
		echo "	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)";
		echo;
		echo "	return nil";
		echo "}";

		echo;

		echo "// Remove deletes the specified position and shifts remaining data down.";
		echo "func (l *List$type) Remove(i int) {";
		echo "	if i >= len(*l) {";
		echo "		return";
		echo "	}";
		echo;
		echo "	copy((*l)[i:], (*l)[i+1:])";
		echo;
		if [ "$type" = "List" -o "$type" = "Compound" ]; then
			echo "	(*l)[i] = nil";
		fi;
		echo "	*l = (*l)[:len(*l)-1]";
		echo;
		echo "}";

		echo;

		echo "// Len returns the length of the list.";
		echo "func (l List$type) Len() int {";
		echo "	return len(l)";
		echo "}";

		echo;

		echo "// List$type returns the list as a specifically typed List.";
		echo "func (l ListData) List$type() List$type {";
		echo "	if l.tagType != Tag$type {";
		echo "		return nil";
		echo "	}";
		echo;
		echo "	s := make(List$type, len(l.data))";
		echo;
		echo "	for n, v := range l.data {";
		echo "		s[n] = v.($type)";
		echo "	}";
		echo;
		echo "	return s";
		echo "}";
	done;
} > "lists.go";
