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
	hasContainer := false
	if len(self.Children) == 0 {
		return false
	}

	if _, ok := self.Children[len(self.Children)-1].Ast.(*ContainerAST); ok {
		hasContainer = true
	}
	return hasContainer
}

func (e *ElementAST) String(self *structure.ASTWrapper) string {
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

	content := self.JoinChildren()
	if content != "" {
		return fmt.Sprintf("<%s%s>%s</%s>", e.tag, attrStr, content, e.tag)
	}
	return fmt.Sprintf("<%s%s />", e.tag, attrStr)
}

type ElementNode struct{}

func NewElementNode() *ElementNode {
	return &ElementNode{}
}

func (e *ElementNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
		structure.STATE_CONTAINER,
	}
}

func (e *ElementNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	if t.EqualsRegexp(`[a-z]+`) {
		scope.Set("class", variables.NewStringVariable(""))
		scope.Set("id", variables.NewStringVariable(""))
		scope.Set("attributes", variables.NewMapVariable(map[string]any{}))
		return structure.STATE_ELEMENT, &ElementAST{
			tag: t.Str,
		}, structure.NewScopeWithParent(scope)
	}
	return 0, nil, nil
}
