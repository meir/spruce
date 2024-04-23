package spruce

type Variable interface {
	Value() any
	Set(any)
	String() string
}
