package states

import (
	"fmt"
	"regexp"

	"github.com/meir/spruce/pkg/structure"
)

type ElementIdAST struct {
	id string
}

func (e *ElementIdAST) Next(ts *structure.Tokenizer) bool {
	e.id = ts.Current().Str
	return true
}

func (e ElementIdAST) String(children []*structure.ASTWrapper) string {
	return fmt.Sprintf(" id=\"%s\"", e.id)
}

type ElementIdNode struct{}

func NewElementIdNode() *ElementIdNode {
	return &ElementIdNode{}
}

func (e *ElementIdNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
		structure.STATE_ELEMENT_ATTRIBUTE,
		structure.STATE_ELEMENT_ID,
		structure.STATE_ELEMENT_CLASS,
	}
}

func (e *ElementIdNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST) {
	t := ts.PeekNext(2)
	if t == nil {
		return 0, nil
	}

	rexp := regexp.MustCompile(`#[a-zA-Z]+`)
	if rexp.MatchString(t.Str) {
		return structure.STATE_ELEMENT_ID, &ElementIdAST{id: ""}
	}
	return 0, nil
}
