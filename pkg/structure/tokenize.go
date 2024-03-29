package structure

import "regexp"

type Token struct {
	Str string

	Line  int
	Start int
	End   int
}

type Tokenizer struct {
	Tokens []Token

	index int
}

func (t *Token) Join(t2 *Token) bool {
	if t.Line != t2.Line {
		return false
	}

	if t.Start+len(t.Str) != t2.Start {
		return false
	}

	t.Str += t2.Str
	t.End = t2.End
	return true
}

const tokenizeRegex = `([^a-zA-Z ]|[a-zA-Z]+[a-zA-Z0-9-_]*| +)`

var tokenizeRegexp = regexp.MustCompile(tokenizeRegex)

func NewTokens(s string) *Tokenizer {
	return &Tokenizer{
		Tokens: tokenize(s),

		index: 0,
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

func (t *Tokenizer) Peek(i int) *Token {
	if t.index+i < len(t.Tokens) {
		token := t.Tokens[t.index+i]
		return &token
	}
	return nil
}

func (t *Tokenizer) PeekNext(i int) *Token {
	if t.index+i < len(t.Tokens) {
		current := t.Current()
		for j := 1; j <= i; j++ {
			peek := t.Peek(j)
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
