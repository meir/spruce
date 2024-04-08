package states

import (
	"github.com/meir/spruce/pkg/structure"
)

type StringAST struct {
	content string

	quote string
}

func (s *StringAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	t := ts.Current()
	switch t.Str {
	case "\\":
		if ps := ts.PeekNext(len(s.quote) + 1); ps != nil && ps.Str == "\\"+s.quote {
			s.content += "\\" + s.quote
			ts.Skip(len(s.quote))
		} else {
			s.content += t.Str
			ts.Next()
			s.content += ts.Current().Str
		}
		return false
	case "`":
		if ps := ts.PeekNext(len(s.quote)); ps != nil && ps.Str != s.quote {
			return false
		}

		fallthrough
	case s.quote:
		return true
	default:
		s.content += t.Str
		return false
	}
}

func (s *StringAST) String(self *structure.ASTWrapper) string {
	return s.content
}

type StringNode struct{}

func NewStringNode() *StringNode {
	return &StringNode{}
}

func (s *StringNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
		structure.STATE_CONTAINER,
		structure.STATE_ELEMENT_ATTRIBUTE,
		structure.STATE_IMPORT,
	}
}

func (e *StringNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	switch t.Str {
	case "\"", "'", "`":
		quote := t.Str
		if ps := ts.PeekNext(2); ps != nil && ps.Str == "```" {
			quote = ps.Str
			ts.Skip(2)
		}
		return structure.STATE_STRING, &StringAST{quote: quote}, scope
	default:
		return 0, nil, nil
	}
}
