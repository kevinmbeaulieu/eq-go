package main

import (
	"fmt"
	"github.com/kevinmbeaulieu/echo/echo"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	// Compare two directories
	pathPkgA, err := filepath.Abs("example/package-a")
	panicIfError(err)
	pathPkgB, err := filepath.Abs("example/package-b")
	panicIfError(err)
	eq, msg := echo.DirectoriesEquivalent(pathPkgA, pathPkgB)
	fmt.Printf("Directories result: %t (%s)\n", eq, msg)

	// Compare two packages
	pkgA := loadPackage("package-a", pathPkgA)
	pkgB := loadPackage("package-b", pathPkgB)
	eq, msg = echo.PackagesEquivalent(pkgA, pkgB)
	fmt.Printf("Packages result: %t (%s)\n", eq, msg)

	// Compare two files
	fset := token.NewFileSet()
	pathFoo, err := filepath.Abs("example/package-a/foo.go")
	pathBar, err := filepath.Abs("example/package-b/bar.go")
	fileA, err := parser.ParseFile(fset, pathFoo, nil, parser.AllErrors)
	panicIfError(err)
	fileB, err := parser.ParseFile(fset, pathBar, nil, parser.AllErrors)
	panicIfError(err)
	eq, msg = echo.FilesEquivalent(fileA, fileB)
	fmt.Printf("Files result: %t (%s)\n", eq, msg)
}

func panicIfError(err error) {
	if err == nil {
		return
	}

	panic(err)
}

func loadPackage(name string, path string) *ast.Package {
	pkg := ast.Package{
		Name: name,
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

	return &pkg
}