package states

import (
	"path"
	"strings"

	"github.com/meir/spruce/pkg/structure"
)

type ImportAST struct {
	value structure.AST
}

func (e *ImportAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	if len(self.Children) >= 1 {
		child := self.Children[0]
		importStr := child.Ast.String(child)

		if !strings.HasSuffix(importStr, ".spr") {
			importStr += ".spr"
		}

		importStr = path.Join(self.File.Dir, importStr)

		file, err := Parse(importStr)
		if err != nil {
			panic(err)
		}

		self.Scope.Merge(file.Scope)
	}
	return len(self.Children) > 0
}

func (e *ImportAST) String(self *structure.ASTWrapper) string {
	return ""
}

type ImportNode struct{}

func NewImportNode() *ImportNode {
	return &ImportNode{}
}

func (e *ImportNode) States() []structure.State {
	return []structure.State{
		structure.STATE_AT_STATEMENT,
	}
}

func (e *ImportNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	if t == nil {
		return 0, nil, nil
	}

	if t.Str == "import" {
		return structure.STATE_IMPORT, &ImportAST{}, scope
	}
	return 0, nil, nil
}
