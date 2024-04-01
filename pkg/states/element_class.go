package states

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/meir/spruce/pkg/structure"
)

type ElementClassAST struct {
	class   string
	classes []string
}

func (e *ElementClassAST) Next(ts *structure.Tokenizer) bool {
	e.class = ts.Current().Str
	rexp := regexp.MustCompile(`[ ]*[.a-zA-Z-_]+`)
	for ts.Next() {
		curr := ts.Current().Str
		if rexp.MatchString(curr) {
			e.class += strings.TrimSpace(curr)
		} else {
			break
		}
	}

	ts.Skip(-1)
	e.classes = append(e.classes, strings.Split(e.class, ".")...)

	return true
}

func (e ElementClassAST) String(children []*structure.ASTWrapper) string {
	return fmt.Sprintf(" class=\"%s\"", strings.Join(e.classes, " "))
}

type ElementClassNode struct{}

func NewElementClassNode() *ElementClassNode {
	return &ElementClassNode{}
}

func (e *ElementClassNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
		structure.STATE_ELEMENT_ATTRIBUTE,
		structure.STATE_ELEMENT_ID,
		structure.STATE_ELEMENT_CLASS,
	}
}

func (e *ElementClassNode) Active(ts *structure.Tokenizer) (structure.State, structure.AST) {
	t := ts.PeekNext(2)
	if t == nil {
		return 0, nil
	}

	rexp := regexp.MustCompile(`\.[a-zA-Z]+`)
	if rexp.MatchString(t.Str) {
		return structure.STATE_ELEMENT_CLASS, &ElementClassAST{
			class:   "",
			classes: []string{},
		}
	}
	return 0, nil
}
