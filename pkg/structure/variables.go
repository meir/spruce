package structure

type Variable interface {
	Get() any
	Set(any)
	String() string
}
