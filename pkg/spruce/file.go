package spruce

import (
	"os"
	"path"

	"github.com/meir/spruce/pkg/states"
	"github.com/meir/spruce/pkg/structure"
)

func Parse(file string) (*structure.File, error) {
	dir := path.Dir(file)
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	tokens := structure.NewTokens(string(data))
	lexer := structure.NewLexer(tokens, []structure.Node{
		states.NewImportNode(),
		states.NewElementIdNode(),
		states.NewElementClassNode(),
		states.NewVariableInsertNode(),
		states.NewTemplateNode(),
		states.NewElementAttributeNode(),
		states.NewVariableNode(),
		states.NewContainerNode(),
		states.NewElementNode(),
		states.NewStringNode(),
	})
	scope := structure.NewScope()
	asts := lexer.Parse()

	return &structure.File{
		Tokens: tokens,
		Asts:   asts,
		Lexer:  lexer,
		Scope:  scope,

		Dir: dir,
		Raw: data,
	}, nil
}
