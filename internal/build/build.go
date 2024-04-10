package build

import (
	"os"
	"path"
	"strings"

	"github.com/meir/spruce/internal/spruce"
	"github.com/meir/spruce/pkg/states"
	"github.com/meir/spruce/pkg/structure"
	"github.com/meir/spruce/pkg/variables"
)

type FileWrapper struct {
	file *structure.File

	url string
}

func Build(dir string, output_dir string) {
	pages := find_pages(dir)
	println("found", len(pages), "pages")
	for _, page := range pages {
		println("building", page.url)
		file_dir := path.Join(output_dir, page.url)
		err := os.MkdirAll(file_dir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		file := path.Join(file_dir, "index.html")
		content := page.file.Lexer.Format(page.file.Asts)

		err = os.WriteFile(file, []byte(content), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func find_pages(dir string) []*FileWrapper {
	files, err := recursively_get_files(dir)
	if err != nil {
		panic(err)
	}

	println("found", len(files), "files")

	fileWrappers := []*FileWrapper{}
	for _, file := range files {
		f, err := spruce.Parse(file)
		if err != nil {
			panic(err)
		}

		// find MetaAST
		for _, ast := range f.Asts {
			if _, ok := ast.Ast.(*states.AtStatementAST); ok {
				if ast.Children == nil || len(ast.Children) == 0 {
					continue
				}

				if _, ok := ast.Children[0].Ast.(*states.MetaAST); ok {
					ast = ast.Children[0]
				}

				if ast.Scope.Get("attributes") == nil {
					continue
				}

				attributes := ast.Scope.Get("attributes").(*variables.MapVariable)

				if attributes.Get() == nil {
					continue
				}

				attr := attributes.Get().(map[string]interface{})

				url := ""
				for k, v := range attr {
					if k == "url" {
						url = v.(string)
					}
				}

				if url == "" {
					continue
				}

				fileWrappers = append(fileWrappers, &FileWrapper{
					file: f,
					url:  url,
				})
			}
		}
	}
	return fileWrappers
}

const EXT = ".spr"

func recursively_get_files(dir string) ([]string, error) {
	files := []string{}
	dir_files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range dir_files {
		if file.IsDir() {
			sub_files, err := recursively_get_files(path.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			files = append(files, sub_files...)
		} else {
			if strings.HasSuffix(file.Name(), EXT) {
				files = append(files, path.Join(dir, file.Name()))
			}
		}
	}

	return files, nil
}
