// file automatically generated with listGen.sh

package nbt

import "strconv"

type ListByte []Byte

func (l ListByte) Equal(e interface{}) bool {
	m, ok := e.(ListByte)
	if !ok {
		var n *ListByte
		if n, ok = e.(*ListByte); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagByte && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListByte) Copy() Data {
	m := make(ListByte, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Byte)
	}
	return m
}

func (l ListByte) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Byte {"
	for _, d := range l {
		s += "\n        Byte: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListByte) Type() TagID {
	return TagList
}

func (ListByte) TagType() TagID {
	return TagByte
}

func (l ListByte) Set(i int, d Data) error {
	if m, ok := d.(Byte); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagByte, d.Type()}
	}
	return nil
}

func (l ListByte) Get(i int) Data {
	return l[i]
}

func (l ListByte) GetByte(i int) Byte {
	return l[i]
}

func (l *ListByte) Append(d ...Data) error {
	toAppend := make(ListByte, len(d))
	for n, e := range d {
		if f, ok := e.(Byte); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagByte, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListByte) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListByte, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Byte); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagByte, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListByte) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListByte) Len() int {
	return len(l)
}

func (l ListData) ListByte() ListByte {
	if l.tagType != TagByte {
		return nil
	}
	s := make(ListByte, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Byte)
	}
	return s
}

type ListShort []Short

func (l ListShort) Equal(e interface{}) bool {
	m, ok := e.(ListShort)
	if !ok {
		var n *ListShort
		if n, ok = e.(*ListShort); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagShort && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListShort) Copy() Data {
	m := make(ListShort, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Short)
	}
	return m
}

func (l ListShort) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Short {"
	for _, d := range l {
		s += "\n        Short: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListShort) Type() TagID {
	return TagList
}

func (ListShort) TagType() TagID {
	return TagShort
}

func (l ListShort) Set(i int, d Data) error {
	if m, ok := d.(Short); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagShort, d.Type()}
	}
	return nil
}

func (l ListShort) Get(i int) Data {
	return l[i]
}

func (l ListShort) GetShort(i int) Short {
	return l[i]
}

func (l *ListShort) Append(d ...Data) error {
	toAppend := make(ListShort, len(d))
	for n, e := range d {
		if f, ok := e.(Short); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagShort, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListShort) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListShort, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Short); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagShort, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListShort) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListShort) Len() int {
	return len(l)
}

func (l ListData) ListShort() ListShort {
	if l.tagType != TagShort {
		return nil
	}
	s := make(ListShort, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Short)
	}
	return s
}

type ListInt []Int

func (l ListInt) Equal(e interface{}) bool {
	m, ok := e.(ListInt)
	if !ok {
		var n *ListInt
		if n, ok = e.(*ListInt); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagInt && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListInt) Copy() Data {
	m := make(ListInt, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Int)
	}
	return m
}

func (l ListInt) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Int {"
	for _, d := range l {
		s += "\n        Int: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListInt) Type() TagID {
	return TagList
}

func (ListInt) TagType() TagID {
	return TagInt
}

func (l ListInt) Set(i int, d Data) error {
	if m, ok := d.(Int); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagInt, d.Type()}
	}
	return nil
}

func (l ListInt) Get(i int) Data {
	return l[i]
}

func (l ListInt) GetInt(i int) Int {
	return l[i]
}

func (l *ListInt) Append(d ...Data) error {
	toAppend := make(ListInt, len(d))
	for n, e := range d {
		if f, ok := e.(Int); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagInt, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListInt) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListInt, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Int); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagInt, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListInt) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListInt) Len() int {
	return len(l)
}

func (l ListData) ListInt() ListInt {
	if l.tagType != TagInt {
		return nil
	}
	s := make(ListInt, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Int)
	}
	return s
}

type ListLong []Long

func (l ListLong) Equal(e interface{}) bool {
	m, ok := e.(ListLong)
	if !ok {
		var n *ListLong
		if n, ok = e.(*ListLong); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagLong && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListLong) Copy() Data {
	m := make(ListLong, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Long)
	}
	return m
}

func (l ListLong) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Long {"
	for _, d := range l {
		s += "\n        Long: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListLong) Type() TagID {
	return TagList
}

func (ListLong) TagType() TagID {
	return TagLong
}

func (l ListLong) Set(i int, d Data) error {
	if m, ok := d.(Long); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagLong, d.Type()}
	}
	return nil
}

func (l ListLong) Get(i int) Data {
	return l[i]
}

func (l ListLong) GetLong(i int) Long {
	return l[i]
}

func (l *ListLong) Append(d ...Data) error {
	toAppend := make(ListLong, len(d))
	for n, e := range d {
		if f, ok := e.(Long); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagLong, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListLong) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListLong, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Long); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagLong, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListLong) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListLong) Len() int {
	return len(l)
}

func (l ListData) ListLong() ListLong {
	if l.tagType != TagLong {
		return nil
	}
	s := make(ListLong, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Long)
	}
	return s
}

type ListFloat []Float

func (l ListFloat) Equal(e interface{}) bool {
	m, ok := e.(ListFloat)
	if !ok {
		var n *ListFloat
		if n, ok = e.(*ListFloat); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagFloat && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListFloat) Copy() Data {
	m := make(ListFloat, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Float)
	}
	return m
}

func (l ListFloat) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Float {"
	for _, d := range l {
		s += "\n        Float: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListFloat) Type() TagID {
	return TagList
}

func (ListFloat) TagType() TagID {
	return TagFloat
}

func (l ListFloat) Set(i int, d Data) error {
	if m, ok := d.(Float); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagFloat, d.Type()}
	}
	return nil
}

func (l ListFloat) Get(i int) Data {
	return l[i]
}

func (l ListFloat) GetFloat(i int) Float {
	return l[i]
}

func (l *ListFloat) Append(d ...Data) error {
	toAppend := make(ListFloat, len(d))
	for n, e := range d {
		if f, ok := e.(Float); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagFloat, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListFloat) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListFloat, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Float); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagFloat, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListFloat) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListFloat) Len() int {
	return len(l)
}

func (l ListData) ListFloat() ListFloat {
	if l.tagType != TagFloat {
		return nil
	}
	s := make(ListFloat, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Float)
	}
	return s
}

type ListDouble []Double

func (l ListDouble) Equal(e interface{}) bool {
	m, ok := e.(ListDouble)
	if !ok {
		var n *ListDouble
		if n, ok = e.(*ListDouble); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagDouble && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListDouble) Copy() Data {
	m := make(ListDouble, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Double)
	}
	return m
}

func (l ListDouble) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Double {"
	for _, d := range l {
		s += "\n        Double: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListDouble) Type() TagID {
	return TagList
}

func (ListDouble) TagType() TagID {
	return TagDouble
}

func (l ListDouble) Set(i int, d Data) error {
	if m, ok := d.(Double); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagDouble, d.Type()}
	}
	return nil
}

func (l ListDouble) Get(i int) Data {
	return l[i]
}

func (l ListDouble) GetDouble(i int) Double {
	return l[i]
}

func (l *ListDouble) Append(d ...Data) error {
	toAppend := make(ListDouble, len(d))
	for n, e := range d {
		if f, ok := e.(Double); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagDouble, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListDouble) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListDouble, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Double); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagDouble, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListDouble) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListDouble) Len() int {
	return len(l)
}

func (l ListData) ListDouble() ListDouble {
	if l.tagType != TagDouble {
		return nil
	}
	s := make(ListDouble, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Double)
	}
	return s
}

type ListIntArray []IntArray

func (l ListIntArray) Equal(e interface{}) bool {
	m, ok := e.(ListIntArray)
	if !ok {
		var n *ListIntArray
		if n, ok = e.(*ListIntArray); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagIntArray && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListIntArray) Copy() Data {
	m := make(ListIntArray, len(l))
	for n, e := range l {
		m[n] = e.Copy().(IntArray)
	}
	return m
}

func (l ListIntArray) String() string {
	s := strconv.Itoa(len(l)) + " entries of type IntArray {"
	for _, d := range l {
		s += "\n        IntArray: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListIntArray) Type() TagID {
	return TagList
}

func (ListIntArray) TagType() TagID {
	return TagIntArray
}

func (l ListIntArray) Set(i int, d Data) error {
	if m, ok := d.(IntArray); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagIntArray, d.Type()}
	}
	return nil
}

func (l ListIntArray) Get(i int) Data {
	return l[i]
}

func (l ListIntArray) GetIntArray(i int) IntArray {
	return l[i]
}

func (l *ListIntArray) Append(d ...Data) error {
	toAppend := make(ListIntArray, len(d))
	for n, e := range d {
		if f, ok := e.(IntArray); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagIntArray, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListIntArray) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListIntArray, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(IntArray); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagIntArray, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListIntArray) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListIntArray) Len() int {
	return len(l)
}

func (l ListData) ListIntArray() ListIntArray {
	if l.tagType != TagIntArray {
		return nil
	}
	s := make(ListIntArray, len(l.data))
	for n, v := range l.data {
		s[n] = v.(IntArray)
	}
	return s
}

type ListBool []Bool

func (l ListBool) Equal(e interface{}) bool {
	m, ok := e.(ListBool)
	if !ok {
		var n *ListBool
		if n, ok = e.(*ListBool); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagBool && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListBool) Copy() Data {
	m := make(ListBool, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Bool)
	}
	return m
}

func (l ListBool) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Bool {"
	for _, d := range l {
		s += "\n        Bool: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListBool) Type() TagID {
	return TagList
}

func (ListBool) TagType() TagID {
	return TagBool
}

func (l ListBool) Set(i int, d Data) error {
	if m, ok := d.(Bool); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagBool, d.Type()}
	}
	return nil
}

func (l ListBool) Get(i int) Data {
	return l[i]
}

func (l ListBool) GetBool(i int) Bool {
	return l[i]
}

func (l *ListBool) Append(d ...Data) error {
	toAppend := make(ListBool, len(d))
	for n, e := range d {
		if f, ok := e.(Bool); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagBool, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListBool) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListBool, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Bool); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagBool, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListBool) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListBool) Len() int {
	return len(l)
}

func (l ListData) ListBool() ListBool {
	if l.tagType != TagBool {
		return nil
	}
	s := make(ListBool, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Bool)
	}
	return s
}

type ListUint8 []Uint8

func (l ListUint8) Equal(e interface{}) bool {
	m, ok := e.(ListUint8)
	if !ok {
		var n *ListUint8
		if n, ok = e.(*ListUint8); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagUint8 && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListUint8) Copy() Data {
	m := make(ListUint8, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Uint8)
	}
	return m
}

func (l ListUint8) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Uint8 {"
	for _, d := range l {
		s += "\n        Uint8: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListUint8) Type() TagID {
	return TagList
}

func (ListUint8) TagType() TagID {
	return TagUint8
}

func (l ListUint8) Set(i int, d Data) error {
	if m, ok := d.(Uint8); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagUint8, d.Type()}
	}
	return nil
}

func (l ListUint8) Get(i int) Data {
	return l[i]
}

func (l ListUint8) GetUint8(i int) Uint8 {
	return l[i]
}

func (l *ListUint8) Append(d ...Data) error {
	toAppend := make(ListUint8, len(d))
	for n, e := range d {
		if f, ok := e.(Uint8); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagUint8, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListUint8) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListUint8, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Uint8); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagUint8, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListUint8) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListUint8) Len() int {
	return len(l)
}

func (l ListData) ListUint8() ListUint8 {
	if l.tagType != TagUint8 {
		return nil
	}
	s := make(ListUint8, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Uint8)
	}
	return s
}

type ListUint16 []Uint16

func (l ListUint16) Equal(e interface{}) bool {
	m, ok := e.(ListUint16)
	if !ok {
		var n *ListUint16
		if n, ok = e.(*ListUint16); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagUint16 && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListUint16) Copy() Data {
	m := make(ListUint16, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Uint16)
	}
	return m
}

func (l ListUint16) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Uint16 {"
	for _, d := range l {
		s += "\n        Uint16: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListUint16) Type() TagID {
	return TagList
}

func (ListUint16) TagType() TagID {
	return TagUint16
}

func (l ListUint16) Set(i int, d Data) error {
	if m, ok := d.(Uint16); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagUint16, d.Type()}
	}
	return nil
}

func (l ListUint16) Get(i int) Data {
	return l[i]
}

func (l ListUint16) GetUint16(i int) Uint16 {
	return l[i]
}

func (l *ListUint16) Append(d ...Data) error {
	toAppend := make(ListUint16, len(d))
	for n, e := range d {
		if f, ok := e.(Uint16); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagUint16, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListUint16) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListUint16, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Uint16); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagUint16, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListUint16) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListUint16) Len() int {
	return len(l)
}

func (l ListData) ListUint16() ListUint16 {
	if l.tagType != TagUint16 {
		return nil
	}
	s := make(ListUint16, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Uint16)
	}
	return s
}

type ListUint32 []Uint32

func (l ListUint32) Equal(e interface{}) bool {
	m, ok := e.(ListUint32)
	if !ok {
		var n *ListUint32
		if n, ok = e.(*ListUint32); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagUint32 && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListUint32) Copy() Data {
	m := make(ListUint32, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Uint32)
	}
	return m
}

func (l ListUint32) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Uint32 {"
	for _, d := range l {
		s += "\n        Uint32: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListUint32) Type() TagID {
	return TagList
}

func (ListUint32) TagType() TagID {
	return TagUint32
}

func (l ListUint32) Set(i int, d Data) error {
	if m, ok := d.(Uint32); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagUint32, d.Type()}
	}
	return nil
}

func (l ListUint32) Get(i int) Data {
	return l[i]
}

func (l ListUint32) GetUint32(i int) Uint32 {
	return l[i]
}

func (l *ListUint32) Append(d ...Data) error {
	toAppend := make(ListUint32, len(d))
	for n, e := range d {
		if f, ok := e.(Uint32); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagUint32, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListUint32) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListUint32, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Uint32); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagUint32, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListUint32) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListUint32) Len() int {
	return len(l)
}

func (l ListData) ListUint32() ListUint32 {
	if l.tagType != TagUint32 {
		return nil
	}
	s := make(ListUint32, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Uint32)
	}
	return s
}

type ListUint64 []Uint64

func (l ListUint64) Equal(e interface{}) bool {
	m, ok := e.(ListUint64)
	if !ok {
		var n *ListUint64
		if n, ok = e.(*ListUint64); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagUint64 && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListUint64) Copy() Data {
	m := make(ListUint64, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Uint64)
	}
	return m
}

func (l ListUint64) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Uint64 {"
	for _, d := range l {
		s += "\n        Uint64: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListUint64) Type() TagID {
	return TagList
}

func (ListUint64) TagType() TagID {
	return TagUint64
}

func (l ListUint64) Set(i int, d Data) error {
	if m, ok := d.(Uint64); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagUint64, d.Type()}
	}
	return nil
}

func (l ListUint64) Get(i int) Data {
	return l[i]
}

func (l ListUint64) GetUint64(i int) Uint64 {
	return l[i]
}

func (l *ListUint64) Append(d ...Data) error {
	toAppend := make(ListUint64, len(d))
	for n, e := range d {
		if f, ok := e.(Uint64); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagUint64, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListUint64) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListUint64, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Uint64); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagUint64, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListUint64) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListUint64) Len() int {
	return len(l)
}

func (l ListData) ListUint64() ListUint64 {
	if l.tagType != TagUint64 {
		return nil
	}
	s := make(ListUint64, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Uint64)
	}
	return s
}

type ListComplex64 []Complex64

func (l ListComplex64) Equal(e interface{}) bool {
	m, ok := e.(ListComplex64)
	if !ok {
		var n *ListComplex64
		if n, ok = e.(*ListComplex64); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagComplex64 && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListComplex64) Copy() Data {
	m := make(ListComplex64, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Complex64)
	}
	return m
}

func (l ListComplex64) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Complex64 {"
	for _, d := range l {
		s += "\n        Complex64: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListComplex64) Type() TagID {
	return TagList
}

func (ListComplex64) TagType() TagID {
	return TagComplex64
}

func (l ListComplex64) Set(i int, d Data) error {
	if m, ok := d.(Complex64); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagComplex64, d.Type()}
	}
	return nil
}

func (l ListComplex64) Get(i int) Data {
	return l[i]
}

func (l ListComplex64) GetComplex64(i int) Complex64 {
	return l[i]
}

func (l *ListComplex64) Append(d ...Data) error {
	toAppend := make(ListComplex64, len(d))
	for n, e := range d {
		if f, ok := e.(Complex64); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagComplex64, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListComplex64) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListComplex64, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Complex64); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagComplex64, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListComplex64) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListComplex64) Len() int {
	return len(l)
}

func (l ListData) ListComplex64() ListComplex64 {
	if l.tagType != TagComplex64 {
		return nil
	}
	s := make(ListComplex64, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Complex64)
	}
	return s
}

type ListComplex128 []Complex128

func (l ListComplex128) Equal(e interface{}) bool {
	m, ok := e.(ListComplex128)
	if !ok {
		var n *ListComplex128
		if n, ok = e.(*ListComplex128); ok {
			m = *n
		}
	}
	if ok {
		if len(l) == len(m) {
			for n, t := range m {
				if !t.Equal(l[n]) {
					return false
				}
			}
			return true
		}
	} else if d, ok := e.(List); ok && d.TagType() == TagComplex128 && d.Len() == len(l) {
		for i := 0; i < d.Len(); i++ {
			if !d.Get(i).Equal(l[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l ListComplex128) Copy() Data {
	m := make(ListComplex128, len(l))
	for n, e := range l {
		m[n] = e.Copy().(Complex128)
	}
	return m
}

func (l ListComplex128) String() string {
	s := strconv.Itoa(len(l)) + " entries of type Complex128 {"
	for _, d := range l {
		s += "\n        Complex128: " + indent(d.String())
	}
	return s + "\n}"
}

func (ListComplex128) Type() TagID {
	return TagList
}

func (ListComplex128) TagType() TagID {
	return TagComplex128
}

func (l ListComplex128) Set(i int, d Data) error {
	if m, ok := d.(Complex128); ok {
		if i <= 0 || i >= int(len(l)) {
			return ErrBadRange
		}
		l[i] = m
	} else {
		return &WrongTag{TagComplex128, d.Type()}
	}
	return nil
}

func (l ListComplex128) Get(i int) Data {
	return l[i]
}

func (l ListComplex128) GetComplex128(i int) Complex128 {
	return l[i]
}

func (l *ListComplex128) Append(d ...Data) error {
	toAppend := make(ListComplex128, len(d))
	for n, e := range d {
		if f, ok := e.(Complex128); ok {
			toAppend[n] = f
		} else {
			return &WrongTag{TagComplex128, e.Type()}
		}
	}
	*l = append(*l, toAppend...)
	return nil
}

func (l *ListComplex128) Insert(i int, d ...Data) error {
	if i >= len(*l) {
		return l.Append(d...)
	}
	toInsert := make(ListComplex128, len(d), len(d)+len(*l)-int(i))
	for n, e := range d {
		if f, ok := e.(Complex128); ok {
			toInsert[n] = f
		} else {
			return &WrongTag{TagComplex128, e.Type()}
		}
	}
	*l = append((*l)[:i], append(toInsert, (*l)[i:]...)...)
	return nil
}

func (l *ListComplex128) Remove(i int) {
	if i >= len(*l) {
		return
	}
	copy((*l)[i:], (*l)[i+1:])
	*l = (*l)[:len(*l)-1]
	return
}

func (l ListComplex128) Len() int {
	return len(l)
}

func (l ListData) ListComplex128() ListComplex128 {
	if l.tagType != TagComplex128 {
		return nil
	}
	s := make(ListComplex128, len(l.data))
	for n, v := range l.data {
		s[n] = v.(Complex128)
	}
	return s
}
