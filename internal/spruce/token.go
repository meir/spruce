package spruce

import (
	"fmt"
	"regexp"
	"strings"
)

type Token struct {
	Str string

	Line int

	Start int
	End   int
}

func (t Token) String() string {
	return t.Str
}

func (t Token) IsEmpty() bool {
	return strings.TrimSpace(t.Str) == ""
}

func (t Token) MatchString(s string) bool {
	return t.Str == s
}

func (t Token) Match(rgx string) bool {
	return regexp.MustCompile(rgx).MatchString(t.Str)
}

func (t Token) Location() string {
	return fmt.Sprintf("%d:%d", t.Line, t.Start)
}
