package structure

import (
	"regexp"
	"strings"
)

type Token struct {
	Str string

	Line  int
	Start int
	End   int
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

func (t *Token) IsEmpty() bool {
	if curr := t; curr != nil {
		return strings.TrimSpace(curr.Str) == ""
	}
	return true
}

func (t *Token) Equals(str string) bool {
	if curr := t; curr != nil {
		return curr.Str == str
	}
	return false
}

func (t *Token) EqualsRegexp(str string) bool {
	rexp := regexp.MustCompile(str)
	if curr := t; curr != nil {
		return rexp.MatchString(curr.Str)
	}
	return false
}
