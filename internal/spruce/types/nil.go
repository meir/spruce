package types

type Nil struct{}

func (n *Nil) Value() any {
	return nil
}

func (n *Nil) Set(v any) {
}

func (n *Nil) String() string {
	return ""
}
