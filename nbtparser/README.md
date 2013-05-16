# nbt
--
    import "github.com/MJKWoolnough/minecraft/nbtparser"

Package nbt implements a full Named Binary Tag reader/writer, based on the specs at
http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt

## Usage

#### type Byte

```go
type Byte byte
```


#### func (Byte) Copy

```go
func (n Byte) Copy() Data
```

#### func (Byte) Equal

```go
func (n Byte) Equal(e equaler.Equaler) bool
```

#### func (*Byte) ReadFrom

```go
func (n *Byte) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (Byte) String

```go
func (n Byte) String() string
```

#### func (*Byte) WriteTo

```go
func (n *Byte) WriteTo(f io.Writer) (total int64, err error)
```

#### type ByteArray

```go
type ByteArray []byte
```


#### func (ByteArray) Copy

```go
func (n ByteArray) Copy() Data
```

#### func (ByteArray) Equal

```go
func (n ByteArray) Equal(e equaler.Equaler) bool
```

#### func (*ByteArray) ReadFrom

```go
func (n *ByteArray) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (ByteArray) String

```go
func (n ByteArray) String() string
```

#### func (*ByteArray) WriteTo

```go
func (n *ByteArray) WriteTo(f io.Writer) (total int64, err error)
```

#### type Compound

```go
type Compound []Tag
```


#### func (Compound) Copy

```go
func (n Compound) Copy() Data
```

#### func (Compound) Equal

```go
func (n Compound) Equal(e equaler.Equaler) bool
```

#### func (*Compound) ReadFrom

```go
func (n *Compound) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (Compound) String

```go
func (n Compound) String() string
```

#### func (*Compound) WriteTo

```go
func (n *Compound) WriteTo(f io.Writer) (total int64, err error)
```

#### type Data

```go
type Data interface {
	io.ReaderFrom
	io.WriterTo
	equaler.Equaler
	Copy() Data
	String() string
}
```


#### type Double

```go
type Double float64
```


#### func (Double) Copy

```go
func (n Double) Copy() Data
```

#### func (Double) Equal

```go
func (n Double) Equal(e equaler.Equaler) bool
```

#### func (*Double) ReadFrom

```go
func (n *Double) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (Double) String

```go
func (n Double) String() string
```

#### func (*Double) WriteTo

```go
func (n *Double) WriteTo(f io.Writer) (total int64, err error)
```

#### type Float

```go
type Float float32
```


#### func (Float) Copy

```go
func (n Float) Copy() Data
```

#### func (Float) Equal

```go
func (n Float) Equal(e equaler.Equaler) bool
```

#### func (*Float) ReadFrom

```go
func (n *Float) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (Float) String

```go
func (n Float) String() string
```

#### func (*Float) WriteTo

```go
func (n *Float) WriteTo(f io.Writer) (total int64, err error)
```

#### type Int

```go
type Int int32
```


#### func (Int) Copy

```go
func (n Int) Copy() Data
```

#### func (Int) Equal

```go
func (n Int) Equal(e equaler.Equaler) bool
```

#### func (*Int) ReadFrom

```go
func (n *Int) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (Int) String

```go
func (n Int) String() string
```

#### func (*Int) WriteTo

```go
func (n *Int) WriteTo(f io.Writer) (total int64, err error)
```

#### type IntArray

```go
type IntArray []int32
```


#### func (IntArray) Copy

```go
func (n IntArray) Copy() Data
```

#### func (IntArray) Equal

```go
func (n IntArray) Equal(e equaler.Equaler) bool
```

#### func (*IntArray) ReadFrom

```go
func (n *IntArray) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (IntArray) String

```go
func (n IntArray) String() string
```

#### func (*IntArray) WriteTo

```go
func (n *IntArray) WriteTo(f io.Writer) (total int64, err error)
```

#### type List

```go
type List struct {
}
```


#### func (List) Copy

```go
func (n List) Copy() Data
```

#### func (List) Equal

```go
func (n List) Equal(e equaler.Equaler) bool
```

#### func (*List) ReadFrom

```go
func (n *List) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (List) String

```go
func (n List) String() string
```

#### func (*List) WriteTo

```go
func (n *List) WriteTo(f io.Writer) (total int64, err error)
```

#### type Long

```go
type Long int64
```


#### func (Long) Copy

```go
func (n Long) Copy() Data
```

#### func (Long) Equal

```go
func (n Long) Equal(e equaler.Equaler) bool
```

#### func (*Long) ReadFrom

```go
func (n *Long) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (Long) String

```go
func (n Long) String() string
```

#### func (*Long) WriteTo

```go
func (n *Long) WriteTo(f io.Writer) (total int64, err error)
```

#### type Short

```go
type Short int16
```


#### func (Short) Copy

```go
func (n Short) Copy() Data
```

#### func (Short) Equal

```go
func (n Short) Equal(e equaler.Equaler) bool
```

#### func (*Short) ReadFrom

```go
func (n *Short) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (Short) String

```go
func (n Short) String() string
```

#### func (*Short) WriteTo

```go
func (n *Short) WriteTo(f io.Writer) (total int64, err error)
```

#### type String

```go
type String string
```


#### func (String) Copy

```go
func (n String) Copy() Data
```

#### func (String) Equal

```go
func (n String) Equal(e equaler.Equaler) bool
```

#### func (*String) ReadFrom

```go
func (n *String) ReadFrom(f io.Reader) (total int64, err error)
```

#### func (String) String

```go
func (n String) String() string
```

#### func (*String) WriteTo

```go
func (n *String) WriteTo(f io.Writer) (total int64, err error)
```

#### type Tag

```go
type Tag interface {
	io.ReaderFrom
	io.WriterTo
	equaler.Equaler
	Data() Data
	Name() String
	String() string
	Tag() TagId
	Copy() Tag
}
```


#### func  NewTag

```go
func NewTag(name string, d Data) (n Tag, err error)
```

#### type TagId

```go
type TagId uint8
```


```go
const (
	Tag_End TagId = iota
	Tag_Byte
	Tag_Short
	Tag_Int
	Tag_Long
	Tag_Float
	Tag_Double
	Tag_ByteArray
	Tag_String
	Tag_List
	Tag_Compound
	Tag_IntArray
)
```
Tag Types

#### func (TagId) String

```go
func (t TagId) String() string
```
