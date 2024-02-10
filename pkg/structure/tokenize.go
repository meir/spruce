package structure

import "regexp"

type Token struct {
	Str string

	Line  int
	Start int
	End   int
}

type Tokenizer struct {
	Tokens []*Token

	index int
}

const tokenizeRegex = `([^a-zA-Z0-9-_ ]|[a-zA-Z]+[a-zA-Z0-9-_]*| +)`

var tokenizeRegexp = regexp.MustCompile(tokenizeRegex)

func NewTokens(s string) *Tokenizer {
	return &Tokenizer{
		Tokens: tokenize(s),

		index: 0,
	}
}

func tokenize(s string) []*Token {
	tokens := tokenizeRegexp.FindAllString(s, -1)
	var result []*Token
	var line int
	var start int

	for _, tok := range tokens {
		result = append(result, &Token{
			Str:   tok,
			Line:  line,
			Start: start,
			End:   start + len(tok),
		})
		start += len(tok)

		if tok == "\n" {
			line++
			start = 0
		}
	}

	return result
}

func (t *Tokenizer) Next() *Token {
	if t.index >= len(t.Tokens) {
		return nil
	}
	tok := t.Tokens[t.index]
	t.index++
	return tok
}

func (t *Tokenizer) Peek() *Token {
	if t.index >= len(t.Tokens) {
		return nil
	}
	return t.Tokens[t.index]
}
