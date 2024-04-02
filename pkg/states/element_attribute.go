package states

import (
	"github.com/meir/spruce/pkg/structure"
	"github.com/meir/spruce/pkg/variables"
)

type ElementAttributeAST struct {
	key   string
	scope *structure.Scope
}

func (e *ElementAttributeAST) Next(ts *structure.Tokenizer) bool {
	return true
}

func (e *ElementAttributeAST) String(children []*structure.ASTWrapper) string {
	if m, ok := e.scope.Get("attributes").(*variables.MapVariable); ok {
		if v, ok := m.Get().(map[string]any); ok {
			v[e.key] = structure.JoinChildren(children)
			m.Set(v)
		}
	}
	return ""
}

type ElementAttributeNode struct{}

func NewElementAttributeNode() *ElementAttributeNode {
	return &ElementAttributeNode{}
}

func (e *ElementAttributeNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
		structure.STATE_ELEMENT_CONTENT,
	}
}

func (e *ElementAttributeNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST) {
	t := ts.Current()
	if t.Equals("=") {
		key := ts.PeekActual(-1).Str
		return structure.STATE_ELEMENT_ATTRIBUTE, &ElementAttributeAST{
			key:   key,
			scope: scope,
		}
	}
	return 0, nil
}
