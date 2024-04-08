package states

import (
	"os"
	"path"

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
		NewAtStatementNode(),
		NewImportNode(),
		NewElementIdNode(),
		NewElementClassNode(),
		NewVariableInsertNode(),
		NewTemplateNode(),
		NewElementAttributeNode(),
		NewVariableNode(),
		NewContainerNode(),
		NewElementNode(),
		NewStringNode(),
	})

	scope := structure.NewScope()
	sfile := &structure.File{
		Tokens: tokens,
		Lexer:  lexer,
		Scope:  scope,

		Dir: dir,
		Raw: data,
	}

	sfile.Asts = lexer.Parse(sfile)

	return sfile, nil
}
