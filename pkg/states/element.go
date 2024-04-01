package states

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/meir/spruce/pkg/structure"
)

type ElementAST struct {
	tag string
}

func (e ElementAST) Next(ts *structure.Tokenizer) bool {
	return ts.Current().Str == "}"
}

func (e *ElementAST) String(children []*structure.ASTWrapper) string {
	content := []string{}
	for _, c := range children {
		content = append(content, c.Ast.String(c.Children))
	}

	// <%s%s%s> might seem weird but element_content will provide the >%s</
	return fmt.Sprintf("<%s%s%s>", e.tag, strings.Join(content, ""), e.tag)
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

func (e *ElementNode) Active(ts *structure.Tokenizer) (structure.State, structure.AST) {
	t := ts.Current()
	rexp := regexp.MustCompile(`[a-z]+`)
	if rexp.MatchString(t.Str) {
		return structure.STATE_ELEMENT, &ElementAST{tag: t.Str}
	}
	return 0, nil
}
