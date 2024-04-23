package spruce

type State uint64

const (
	STATE_ROOT State = iota
	STATE_IDENTIFIER
	STATE_SCOPE
	STATE_STRING
	STATE_AT_TAG

	STATE_CUSTOM
)
