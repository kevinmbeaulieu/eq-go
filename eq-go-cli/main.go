package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"

	eqgo "github.com/kevinmbeaulieu/eq-go/eq-go"
)

type stringSliceArg []string

func (a *stringSliceArg) String() string {
	if a == nil {
		return "<nil>"
	}
	return fmt.Sprintf("[%s]", strings.Join(*a, ", "))
}

func (a *stringSliceArg) Set(s string) error {
	values := strings.Split(s, ",")
	*a = append(*a, values...)
	return nil
}

// Usage: `go run path/to/eq-go-cli --pkgs foo,bar --paths path/to/package/foo,path/to/package/bar`
func main() {
	var pkgNamesArg stringSliceArg
	flag.Var(&pkgNamesArg, "pkgs", "Comma-separated pair of input packages' names")

	var pkgPathsArg stringSliceArg
	flag.Var(&pkgPathsArg, "paths", "Comma-separated pair of input packages' paths")

	flag.Parse()

	lhsPkgName := pkgNamesArg[0]
	rhsPkgName := pkgNamesArg[1]

	lhsPkgPath := pkgPathsArg[0]
	rhsPkgPath := pkgPathsArg[1]

	lhsPkg, lhsFSet := loadPackage(lhsPkgName, lhsPkgPath)
	rhsPkg, rhsFSet := loadPackage(rhsPkgName, rhsPkgPath)

	eq, msg := eqgo.PackagesEquivalent(lhsPkg, lhsFSet, rhsPkg, rhsFSet, nil)
	if eq {
		fmt.Printf("%s (%s) and %s (%s) are equivalent.\n", lhsPkgName, lhsPkgPath, rhsPkgName, rhsPkgPath)
	} else {
		fmt.Printf("%s (%s) and %s (%s) are not equivalent.\n\n%s\n", lhsPkgName, lhsPkgPath, rhsPkgName, rhsPkgPath, msg)
	}
}

func loadPackage(name string, path string) (*ast.Package, *token.FileSet) {
	pkg := ast.Package{
		Name:  name,
		Files: make(map[string]*ast.File),
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".go") {
			continue
		}

		filename := filepath.Join(path, f.Name())
		src, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
		if err != nil {
			panic(err)
		}

		pkg.Files[filename] = src
	}

	return &pkg, fset
}
