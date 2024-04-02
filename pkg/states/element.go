package states

import (
	"fmt"

	"github.com/meir/spruce/pkg/structure"
	"github.com/meir/spruce/pkg/variables"
)

type ElementAST struct {
	tag   string
	scope *structure.Scope
}

func (e ElementAST) Next(ts *structure.Tokenizer) bool {
	return ts.Current().Str == "}"
}

func (e *ElementAST) String(children []*structure.ASTWrapper) string {
	// <%s%s%s> might seem weird but element_content will provide the >%s</

	attributes := e.scope.Get("attributes").String()
	println(attributes)
	return fmt.Sprintf("<%s%s%s%s>", e.tag, attributes, structure.JoinChildren(children), e.tag)
}

type ElementNode struct {
	ast *ElementAST
}

func NewElementNode() *ElementNode {
	return &ElementNode{}
}

func (e *ElementNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
		structure.STATE_ELEMENT_CONTENT,
	}
}

func (e *ElementNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST) {
	t := ts.Current()
	if t.EqualsRegexp(`[a-z]+`) {
		scope.Set("class", variables.NewStringVariable(""))
		scope.Set("id", variables.NewStringVariable(""))
		scope.Set("attributes", variables.NewMapVariable(map[string]interface{}{"test": 1}))
		return structure.STATE_ELEMENT, &ElementAST{
			tag:   t.Str,
			scope: scope,
		}
	}
	return 0, nil
}
