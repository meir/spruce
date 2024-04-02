package states

import (
	"fmt"
	"strings"

	"github.com/meir/spruce/pkg/structure"
)

type ElementContentAST struct{}

func (e *ElementContentAST) Next(ts *structure.Tokenizer) bool {
	t := ts.Current()
	switch t.Str {
	case "}":
		ts.Skip(-1)
		return true
	default:
		return false
	}
}

func (e ElementContentAST) String(children []*structure.ASTWrapper) string {
	attr := []string{}
	content := []string{}
	for _, c := range children {
		switch c.Ast.(type) {
		case *ElementAttributeAST:
			attr = append(attr, c.Ast.String(c.Children))
			continue
		}
		content = append(content, c.Ast.String(c.Children))
	}

	return fmt.Sprintf("%s>%s</", strings.Join(attr, ""), strings.Join(content, ""))
}

type ElementContentNode struct{}

func NewElementContentNode() *ElementContentNode {
	return &ElementContentNode{}
}

func (e *ElementContentNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
		structure.STATE_ELEMENT_ATTRIBUTE,
		structure.STATE_ELEMENT_ID,
		structure.STATE_ELEMENT_CLASS,
	}
}

func (e *ElementContentNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST) {
	t := ts.Current()
	switch t.Str {
	case "{":
		return structure.STATE_ELEMENT_CONTENT, &ElementContentAST{}
	default:
		return 0, nil
	}
}
