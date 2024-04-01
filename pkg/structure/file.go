package structure

type File struct {
	Tokens *Tokenizer
	Asts   []*ASTWrapper
	Lexer  *Lexer

	Scope *Scope

	Dir string
	Raw []byte
}

func (f *File) String() string {
	return f.Lexer.Format(f.Asts)
}
