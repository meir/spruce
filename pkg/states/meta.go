package states

import (
	"github.com/meir/spruce/pkg/structure"
	"github.com/meir/spruce/pkg/variables"
)

type MetaAST struct {
	rootScope *structure.Scope
}

func (m *MetaAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	if len(self.Children) == 1 {
		if attributes, ok := structure.Get[map[string]any](self.Scope.Get("attributes")); ok {
			if url, ok := attributes["url"]; ok {
				m.rootScope.Set("url", variables.NewStringVariable(url))
			}

			if title, ok := attributes["title"]; ok {
				m.rootScope.Set("title", variables.NewStringVariable(title))
			}

			if description, ok := attributes["description"]; ok {
				m.rootScope.Set("description", variables.NewStringVariable(description))
			}

			// TODO: Create array variable for keywords and tags
		}
		return true
	}
	return false
}

func (m *MetaAST) String(self *structure.ASTWrapper) string {
	return ""
}

type MetaNode struct{}

func NewMetaNode() *MetaNode {
	return &MetaNode{}
}

func (m *MetaNode) States() []structure.State {
	return []structure.State{
		structure.STATE_AT_STATEMENT,
	}
}

func (m *MetaNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	if t == nil {
		return 0, nil, nil
	}

	meta_scope := structure.NewScope()
	meta_scope.Set("attributes", variables.NewMapVariable(map[string]any{}))

	if t.Str == "meta" {
		return structure.STATE_META, &MetaAST{
			rootScope: scope,
		}, meta_scope
	}
	return 0, nil, nil
}
