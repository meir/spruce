package states

import (
	"github.com/meir/spruce/pkg/structure"
)

type ElementClassAST struct {
	class   string
	classes []string
}

func (e *ElementClassAST) Next(ts *structure.Tokenizer) bool {
	// add class to scope
	return true
}

func (e ElementClassAST) String(children []*structure.ASTWrapper) string {
	return ""
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

func (e *ElementClassNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST) {
	t := ts.PeekNext(2)
	if t.EqualsRegexp(`\.[a-zA-Z]+`) {
		return structure.STATE_ELEMENT_CLASS, &ElementClassAST{
			class:   "",
			classes: []string{},
		}
	}
	return 0, nil
}
