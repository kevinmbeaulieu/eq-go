package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"

	eqgo "github.com/kevinmbeaulieu/eq-go/eq-go"
)

func main() {
	// Compare two packages
	lhsPkgPath, err := filepath.Abs("package-a")
	panicIfError(err)
	rhsPkgPath, err := filepath.Abs("package-b")
	panicIfError(err)
	lhsPkg, lhsFSet := loadPackage("package-a", lhsPkgPath)
	rhsPkg, rhsFSet := loadPackage("package-b", rhsPkgPath)
	eq, msg := eqgo.PackagesEquivalent(lhsPkg, lhsFSet, rhsPkg, rhsFSet, nil)
	fmt.Printf("Packages result: %t\n%s\n\n", eq, msg)

	// Compare two files
	fset := token.NewFileSet()
	lhsFilePath, err := filepath.Abs("package-a/foo.go")
	panicIfError(err)
	rhsFilePath, err := filepath.Abs("package-b/bar.go")
	panicIfError(err)
	lhsFile, err := parser.ParseFile(fset, lhsFilePath, nil, parser.AllErrors)
	panicIfError(err)
	rhsFile, err := parser.ParseFile(fset, rhsFilePath, nil, parser.AllErrors)
	panicIfError(err)
	eq, msg = eqgo.FilesEquivalent(lhsFile, fset, rhsFile, fset, nil)
	fmt.Printf("Files result: %t\n%s\n\n", eq, msg)
}

func panicIfError(err error) {
	if err == nil {
		return
	}

	panic(err)
}

func loadPackage(name string, path string) (*ast.Package, *token.FileSet) {
	pkg := ast.Package{
		Name:  name,
		Files: make(map[string]*ast.File),
	}

	files, err := ioutil.ReadDir(path)
	panicIfError(err)
	fset := token.NewFileSet()
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".go") {
			continue
		}

		filename := filepath.Join(path, f.Name())
		src, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
		panicIfError(err)
		pkg.Files[filename] = src
	}

	return &pkg, fset
}
