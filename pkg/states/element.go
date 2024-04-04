package states

import (
	"fmt"

	"github.com/meir/spruce/pkg/structure"
	"github.com/meir/spruce/pkg/variables"
)

type ElementAST struct {
	tag string
}

func (e ElementAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	return ts.Current().Str == "}"
}

func (e *ElementAST) String(self *structure.ASTWrapper) string {
	// <%s%s%s> might seem weird but element_content will provide the >%s</

	attrStr := ""

	if id, ok := structure.Get[string](self.Scope.Get("id")); ok && id != "" {
		attrStr += fmt.Sprintf(" id=\"%s\"", id)
	}

	if class, ok := structure.Get[string](self.Scope.Get("class")); ok && class != "" {
		attrStr += fmt.Sprintf(" class=\"%s\"", class)
	}

	attributes := self.Scope.Get("attributes").String()
	if attributes != "" {
		attrStr += " " + attributes
	}

	return fmt.Sprintf("<%s%s>%s</%s>", e.tag, attrStr, self.JoinChildren(), e.tag)
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
		scope.Set("attributes", variables.NewMapVariable(map[string]any{}))
		return structure.STATE_ELEMENT, &ElementAST{
			tag: t.Str,
		}
	}
	return 0, nil
}
