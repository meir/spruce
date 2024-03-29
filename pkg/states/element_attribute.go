package states

import (
	"fmt"
	"regexp"

	"github.com/meir/spruce/pkg/structure"
)

type ElementAttributeAST struct {
	key   string
	value structure.AST
}

func (e *ElementAttributeAST) Next(ts *structure.Tokenizer) bool {
	return e.value.Next(ts)
}

func (e *ElementAttributeAST) String(children []*structure.ASTWrapper) string {
	return fmt.Sprintf(" %s=\"%s\"", e.key, e.value.String(children))
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

func (e *ElementAttributeNode) Active(ts *structure.Tokenizer) (structure.State, structure.AST) {
	t := ts.PeekNext(1)
	if t == nil {
		return 0, nil
	}

	rexp := regexp.MustCompile(`[a-zA-Z]+=`)
	if rexp.MatchString(t.Str) {
		key := ts.Current().Str
		ts.Skip(2)
		_, stringAst := NewStringNode().Active(ts)
		if stringAst == nil {
			panic(fmt.Errorf("expected string to start at %v:%v", ts.Current().Line, ts.Current().Start))
		}

		return structure.STATE_ELEMENT_ATTRIBUTE, &ElementAttributeAST{
			key:   key,
			value: stringAst,
		}
	}
	return 0, nil
}
