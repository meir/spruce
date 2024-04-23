package spruce

import (
	"os"
	"path"
)

type File struct {
	Dir  string
	File string

	raw    []byte
	Tokens []Token
}

func NewFile(file string) (*File, error) {
	dir := path.Base(file)

	raw, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return &File{
		Dir:  dir,
		File: file,

		raw:    raw,
		Tokens: []Token{},
	}, nil
}
