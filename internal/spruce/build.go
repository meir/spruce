package spruce

import (
	"os"
	"path"
	"strings"

	"github.com/meir/spruce/pkg/structure"
)

type FileWrapper struct {
	file *structure.File

	url string
}

func get_files_from_wrappers(wrappers []*FileWrapper) []*structure.File {
	files := []*structure.File{}
	for _, wrapper := range wrappers {
		files = append(files, wrapper.file)
	}
	return files
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

		println(content)

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
		f, err := Parse(file)
		if err != nil {
			panic(err)
		}

		if url_var := f.Scope.Get("url"); url_var != nil {
			url := url_var.String()
			fileWrappers = append(fileWrappers, &FileWrapper{
				file: f,
				url:  url,
			})
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
