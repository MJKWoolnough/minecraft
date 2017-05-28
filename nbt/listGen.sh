#!/bin/bash

types=( Byte Short Int Long Float Double IntArray Bool Uint8 Uint16 Uint32 Uint64 Complex64 Complex128 );

{
	echo "// file automatically generated with listGen.sh";
	echo;
	echo "package nbt";
	echo;
	echo "import \"strconv\"";

	for type in ${types[@]}; do
		echo;

		echo "type List$type []$type";

		echo;

		echo "func (l List$type) Equal(e interface{}) bool {";
		echo "	m, ok := e.(List$type)";
		echo "	if !ok {";
		echo "		var n *List$type";
		echo "		if n, ok = e.(*List$type); ok {";
		echo "			m = *n";
		echo "		}";
		echo "	}";
		echo "	if ok {";
		echo "		if len(l) == len(m) {";
		echo "			for n, t := range m {";
		echo "				if !t.Equal(l[n]) {"
		echo "					return false";
		echo "				}";
		echo "			}";
		echo "			return true";
		echo "		}";
		echo "	} else if d, ok := e.(List); ok && d.TagType() == Tag$type && d.Len() == len(l) {";
		echo "		for i := 0; i < d.Len(); i++ {";
		echo "			if !d.Get(i).Equal(l[i]) {";
		echo "				return false";
		echo "			}";
		echo "		}";
		echo "		return true";
		echo "	}";
		echo "	return false";
		echo "}";

		echo;

		echo "func (l List$type) Copy() Data {";
		echo "	m := make(List$type, len(l))";
		echo "	for n, e := range l {";
		echo "		m[n] = e.Copy().($type)";
		echo "	}";
		echo "	return m";
		echo "}";

		echo;

		echo "func (l List$type) String() string {";
		echo "	s := strconv.Itoa(len(l)) + \" entries of type $type {\"";
		echo "	for _, d := range l {";
		echo "		s += \"\\n        $type: \" + indent(d.String())";
		echo "	}";
		echo "	return s + \"\\n}\"";
		echo "}";

		echo;

		echo "func (List$type) Type() TagID {";
		echo "	return TagList";
		echo "}";

		echo;

		echo "func (List$type) TagType() TagID {";
		echo "	return Tag$type";
		echo "}";

		echo;

		echo "func (l List$type) Set(i int, d Data) error {";
		echo "	if m, ok := d.($type); ok {";
		echo "		if i <= 0 || i >= int(len(l)) {";
		echo "			return ErrBadRange";
		echo "		}";
		echo "		l[i] = m";
		echo "	} else {";
		echo "		return &WrongTag{Tag$type, d.Type()}";
		echo "	}";
		echo "	return nil";
		echo "}";

		echo;

		echo "func (l List$type) Get(i int) Data {";
		echo "	return l[i]";
		echo "}";

		echo;

		echo "func (l List$type) Get$type(i int) $type {";
		echo "	return l[i]";
		echo "}";

		echo;

		echo "func (l *List$type) Append(d ...Data) error {";
		echo "	toAppend := make(List$type, len(d))";
		echo "	for n, e := range d {";
		echo "		if f, ok := e.($type); ok {";
		echo "			toAppend[n] = f";	
		echo "		} else {";
		echo "			return &WrongTag{Tag$type, e.Type()}";
		echo "		}";
		echo "	}";
		echo "	*l = append(*l, toAppend...)";
		echo "	return nil";
		echo "}";

		echo;

		echo "func (l *List$type) Insert(i int, d ...Data) error {";
		echo "	if i >= len(*l) {";
		echo "		return l.Append(d...)";
		echo "	}";
		echo "	toInsert := make(List$type, len(d), len(d)+len(*l)-int(i))";
		echo "	for n, e := range d {";
		echo "		if f, ok := e.($type); ok {";
		echo "			toInsert[n] = f";	
		echo "		} else {";
		echo "			return &WrongTag{Tag$type, e.Type()}";
		echo "		}";
		echo "	}";
		echo "	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)";
		echo "	return nil";
		echo "}";

		echo;

		echo "func (l *List$type) Remove(i int) {";
		echo "	if i >= len(*l) {";
		echo "		return";
		echo "	}";
		echo "	copy((*l)[i:], (*l)[i+1:])";
		if [ "$type" = "List" -o "$type" = "Compound" ]; then
			echo "	(*l)[i] = nil";
		fi;
		echo "	*l = (*l)[:len(*l)-1]";
		echo "	return";
		echo "}";

		echo;

		echo "func (l List$type) Len() int {";
		echo "	return len(l)";
		echo "}";

		echo;

		echo "func (l ListData) List$type() List$type {";
		echo "	if l.tagType != Tag$type {";
		echo "		return nil";
		echo "	}";
		echo "	s := make(List$type, len(l.data))";
		echo "	for n, v := range l.data {";
		echo "		s[n] = v.($type)";
		echo "	}";
		echo "	return s";
		echo "}";
	done;

} > "lists.go"
