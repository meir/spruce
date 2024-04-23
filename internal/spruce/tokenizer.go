package spruce

import "regexp"

// Tokenizer is used to tokenize a file and keep track of the token index while running through the program
type Tokenizer struct {
	file *File

	line  int
	index int
}

const tokenizerRegex = `([^a-zA-Z ]|[a-zA-Z]+[a-zA-Z0-9-_]*| +)`

var tokenizeRegexp = regexp.MustCompile(tokenizerRegex)

// NewTokenizer will tokenize the file content and return a new Tokenizer
// This will error if the content is empty
func NewTokenizer(file *File) (*Tokenizer, error) {
	if file.raw == nil || len(file.raw) == 0 {
		return nil, ErrCannotTokenizeEmptyFile
	}

	file.Tokens = tokenize(string(file.raw))
	return &Tokenizer{file: file}, nil
}

// tokenize will convert the raw content into a slice of tokens
func tokenize(content string) []Token {
	tokens := tokenizeRegexp.FindAllString(content, -1)
	var result []Token
	line := 1
	start := 1

	for _, token := range tokens {
		result = append(result, Token{
			Str:   token,
			Line:  line,
			Start: start,
			End:   start + len(token),
		})

		start += len(token)

		if token == "\n" {
			line++
			start = 1
		}
	}

	return result
}

// HasNext returns if there is a next token
func (t *Tokenizer) HasNext() bool {
	return t.index+1 < len(t.file.Tokens) && t.index+1 >= 0
}

// Next returns the next token
func (t *Tokenizer) Next() error {
	if t.index+1 < len(t.file.Tokens) && t.index+1 >= 0 {
		t.index++
		return nil
	}

	current := t.file.Tokens[t.index]
	return NewErrEOF(current.Line, current.Start+1)
}

// Current returns the current token
func (t *Tokenizer) Current() (*Token, error) {
	if t.index < len(t.file.Tokens) && t.index >= 0 {
		return &t.file.Tokens[t.index], nil
	}

	return nil, NewErrEOF(0, t.index)
}

// Skip will skip the next [n] tokens
func (t *Tokenizer) Skip(n int) (*Token, error) {
	if n < 0 {
		return nil, ErrNoNegativeSkip
	}
	if t.index+n < len(t.file.Tokens) && t.index+n >= 0 {
		t.index += n
		return &t.file.Tokens[t.index], nil
	}

	return nil, NewErrEOF(0, t.index)
}

// Peek will peek at the [n]th next token
func (t *Tokenizer) Peek(n int) (*Token, error) {
	if t.index+n < len(t.file.Tokens) && t.index+n >= 0 {
		return &t.file.Tokens[t.index+n], nil
	}

	return nil, NewErrEOF(0, t.index)
}

// PeekCheck returns the next token that matches the given regular expression
func (t *Tokenizer) PeekCheck(n int, match regexp.Regexp) (*Token, error) {
	token, err := t.Peek(n)
	if err != nil {
		return nil, err
	}

	for !match.MatchString(token.Str) {
		n++
		token, err = t.Peek(n)
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

// SkipTo skips to a given token. This token has to be an exact token gotten from the tokenizer.
// if custom token is used or a modified token, then error "surpassed token" will be expected.
func (t *Tokenizer) SkipTo(token Token) error {
	current := t.file.Tokens[t.index]

	for !(current.Line == token.Line && current.Start == token.Start) {
		if current.Line > token.Line || (current.Line >= token.Line && current.Start > token.Start) {
			return ErrSurpassedToken
		}

		if err := t.Next(); err != nil {
			return err
		}

		current = t.file.Tokens[t.index]
	}

	return nil
}
