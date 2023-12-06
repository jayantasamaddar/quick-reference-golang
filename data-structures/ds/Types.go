package ds

type Signed interface {
	~int8 | ~int16 | ~int | ~int32 | ~int64
}

type Unsigned interface {
	~uint8 | ~uint16 | ~uint | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Number interface {
	Signed | Unsigned | Float
}

type Type interface {
	Number | Complex | ~string | ~rune
}

// Can apply to Queues and Stacks that use and underlying Slice. Any Slice based linear data structure you build may also implement this.
type SlicekMethods[T Type] interface {
	Get(index int) T
	Set(index int, value T)
	Size() int
	IndexOf(item T) int
	LastIndexOf(item T) int
	Includes(item T) bool
	Find(callback func(index int, value T) bool) T
	FindIndex(callback func(index int, value T) bool) int
	Some(callback func(index int, value T) bool) bool
	Every(callback func(index int, value T) bool) bool
	Join(sep string) string
	String() string
}
