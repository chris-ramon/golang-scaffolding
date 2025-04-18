package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rp := ReplacePaths{}
	result, err := rp.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}

const rootDir string = "./"

type ReplacePaths struct{}

func (rp *ReplacePaths) Run() ([]string, error) {
	goFiles, err := rp.readGoFiles()
	if err != nil {
		return nil, err
	}

	fileImports, err := rp.readAllFileImports(goFiles)
	if err != nil {
		return nil, err
	}

	importPaths := []string{}
	for _, fileImport := range fileImports {
		importPaths = append(importPaths, fileImport)
	}

	return importPaths, nil
}

func (rp *ReplacePaths) readGoFiles() ([]string, error) {
	files := []string{}

	walk := func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(s, ".go") {
			files = append(files, s)
		}

		return nil
	}

	filepath.WalkDir(rootDir, walk)

	return files, nil
}

func (rp *ReplacePaths) readAllFileImports(testFiles []string) ([]string, error) {
	result := []string{}

	for _, filePath := range testFiles {
		funcNames, err := rp.readFileImports(filePath)
		if err != nil {
			return result, err
		}

		result = append(result, funcNames...)
	}

	return result, nil
}

func (rp *ReplacePaths) readFileImports(filePath string) ([]string, error) {
	goFile, err := rp.readFile(filePath)
	if err != nil {
		return nil, err
	}
	defer goFile.Close()

	funcNames := []string{}
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", goFile, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, decl := range astFile.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			funcNames = append(funcNames, t.Name.Name)
		}
	}

	return funcNames, nil
}

func (rp *ReplacePaths) readFile(filePath string) (*os.File, error) {
	goFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return goFile, nil
}
