package structure

import (
	"regexp"
)

type Tokenizer struct {
	Tokens []Token

	index int
}

const tokenizeRegex = `([^a-zA-Z ]|[a-zA-Z]+[a-zA-Z0-9-_]*| +)`

var tokenizeRegexp = regexp.MustCompile(tokenizeRegex)

func NewTokens(s string) *Tokenizer {
	return &Tokenizer{
		Tokens: tokenize(s),

		index: -1,
	}
}

func tokenize(s string) []Token {
	tokens := tokenizeRegexp.FindAllString(s, -1)
	var result []Token
	var line int
	var start int

	for _, tok := range tokens {
		result = append(result, Token{
			Str:   tok,
			Line:  line,
			Start: start,
			End:   start + len(tok) - 1,
		})
		start += len(tok)

		if tok == "\n" {
			line++
			start = 0
		}
	}

	return result
}

func (t *Tokenizer) Current() *Token {
	if t.index < len(t.Tokens) {
		token := t.Tokens[t.index]
		return &token
	}
	return nil
}

func (t *Tokenizer) Next() bool {
	if t.index+1 < len(t.Tokens) {
		t.index++
		return true
	}
	return false
}

func (t *Tokenizer) NextActual() bool {
	if t.index+1 < len(t.Tokens) {
		x := 1
		for {
			pt := t.Peek(x)
			if pt != nil {
				if pt.IsEmpty() {
					x++
					continue
				}
			}
			break
		}
		t.index += x
		return true
	}
	return false
}

func (t *Tokenizer) Peek(i int) *Token {
	if t.index+i < len(t.Tokens) {
		token := t.Tokens[t.index+i]
		return &token
	}
	return nil
}

func (t *Tokenizer) PeekActual(i int) *Token {
	x := 1
	if i < 0 {
		x = -1
	}

Start:
	if t.index+i < len(t.Tokens) {
		if t.Peek(i).IsEmpty() {
			i += x
			goto Start
		}
		return t.Peek(i)
	}
	return nil
}

func (t *Tokenizer) PeekNext(i int) *Token {
	if t.index+i < len(t.Tokens) {
		current := t.Current()
		for j := 1; j < i; j++ {
			peek := t.Peek(j)
			if !current.Join(peek) {
				return nil
			}
		}
		return current
	}
	return nil
}

func (t *Tokenizer) PeekNextActual(i int) *Token {
	x := 1
	if i < 0 {
		x = -1
	}

	if t.index+i < len(t.Tokens) {
		current := t.Current()
		for j, offset := 1, 0; j < i+offset; j++ {
			peek := t.Peek(j)
			if peek.IsEmpty() {
				offset += x
				continue
			}
			if !current.Join(peek) {
				return nil
			}
		}
		return current
	}
	return nil
}

func (t *Tokenizer) Skip(i int) {
	if t.index+i < len(t.Tokens) {
		t.index += i
	}
}

func (t *Tokenizer) SkipTo(tx *Token) {
	for i, token := range t.Tokens {
		if token.Line == tx.Line && token.Start == tx.Start {
			t.index = i
			break
		}
	}
}
