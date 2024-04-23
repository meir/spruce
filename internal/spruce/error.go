package spruce

import "fmt"

type ErrEOF struct {
	Line   int
	Column int
}

func NewErrEOF(line, column int) ErrEOF {
	return ErrEOF{Line: line, Column: column}
}

func (e ErrEOF) Error() string {
	return fmt.Sprintf("unexpected EOF at %d:%d", e.Line, e.Column)
}

var ErrNoNegativeSkip = fmt.Errorf("cannot skip negative number of tokens")
var ErrCannotTokenizeEmptyFile = fmt.Errorf("cannot tokenize empty file")
var ErrSurpassedToken = fmt.Errorf("surpassed token")

type UnexpectedToken struct {
	token *Token
}

func NewUnexpectedToken(token *Token) UnexpectedToken {
	return UnexpectedToken{token: token}
}

func (e UnexpectedToken) Error() string {
	return fmt.Sprintf("unexpected token '%s' at %s", e.token.Str, e.token.Location())
}
