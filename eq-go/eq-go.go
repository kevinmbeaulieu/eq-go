package eqgo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// DirectoriesEquivalent reports whether the source directories a and b are equivalent.
// Directories are considered equivalent if they contain packages/subdirectories which are
// all equivalent.
//
// Packages are equivalent if their sets of declarations are invariant under reordering,
// adding/removing spacing/indentation, and adding/removing comments.
//
// This equivalence operator is a relaxation of strict equality which can be useful for verifying
// whether two differently formatted copies of a codebase are interchangeable from a functional
// perspective. For example, if you have two code generators which arrange their outputs differently,
// DirectoriesEquivalent can verify whether, given the same input, the output of one could be
// swapped with the output of another with no change in behavior to external callers.
//
// Returns: (
//     A boolean indicating whether the directories are equivalent
//     A message describing the first difference found (or "")
// )
func DirectoriesEquivalent(a string, b string) (bool, string) {
	fsetA := token.NewFileSet()
	pkgMapA, err := parser.ParseDir(fsetA, a, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	fsetB := token.NewFileSet()
	pkgMapB, err := parser.ParseDir(fsetB, b, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	fileListA := packageMapToFileList(pkgMapA)
	fileListB := packageMapToFileList(pkgMapB)
	return fileListsEquivalent(fileListA, fileListB)
}

// PackagesEquivalent reports whether the Go packages represented by a and b are equivalent.
// Packages are equivalent if their sets of declarations are invariant under reordering,
// adding/removing spacing/indentation, and adding/removing comments.
//
// This equivalence operator is a relaxation of strict equality which can be useful for verifying
// whether two differently formatted copies of a codebase are interchangeable from a functional
// perspective. For example, if you have two code generators which arrange their outputs differently,
// PackagesEquivalent can verify whether, given the same input, the output of one could be swapped
// with the output of another with no change in behavior to external callers.
//
// Returns: (
//     A boolean indicating whether the packages are equivalent
//     A message describing the first difference found (or "")
// )
func PackagesEquivalent(a *ast.Package, b *ast.Package) (bool, string) {
	if a == nil || b == nil {
		panic(fmt.Errorf("missing package"))
	}

	fileListA := fileMapToList(a.Files)
	fileListB := fileMapToList(b.Files)
	return fileListsEquivalent(fileListA, fileListB)
}

// FilesEquivalent reports whether the Go source files represented by a and b are equivalent.
// Files are equivalent if their sets of declarations/imports are invariant under reordering,
// adding/removing spacing/indentation, and adding/removing comments.
//
// This equivalence operator is a relaxation of strict equality which can be useful for verifying
// whether two differently formatted copies of a codebase are interchangeable from a functional
// perspective. For example, if you have two code generators which arrange their outputs differently,
// FilesEquivalent can verify whether, given the same input, the output of one could be swapped
// with the output of another with no change in behavior to external callers.
//
// Returns: (
//     A boolean indicating whether the files are equivalent
//     A message describing the first difference found (or "")
// )
func FilesEquivalent(a *ast.File, b *ast.File) (bool, string) {
	if eq, msg := declListsEquivalent(a.Decls, b.Decls); !eq {
		return eq, "files' declaration lists not equivalent: " + msg
	}

	if eq, msg := importsEquivalent(a.Imports, b.Imports); !eq {
		return eq, "files' imports not equivalent: " + msg
	}

	if eq, msg := unresolvedIdentifiersEquivalent(a.Unresolved, b.Unresolved); !eq {
		return eq, "files' unresolved identifiers not equivalent: " + msg
	}

	return true, ""
}

// ------- Private Helpers -------

func packageMapToFileList(f map[string]*ast.Package) []*ast.File {
	var result []*ast.File
	for _, pkg := range f {
		result = append(result, fileMapToList(pkg.Files)...)
	}
	return result
}

func fileMapToList(f map[string]*ast.File) []*ast.File {
	var result []*ast.File
	for _, file := range f {
		result = append(result, file)
	}
	return result
}

// merge returns a File node representing the merged contents of all input source files.
// For example, if a list of three files, each declaring their own function, is passed in,
// the return value would be a single File node containing all three functions.
//
// Trivia such as comments, source positions, and file/package names may be trimmed from
// the output file. Only declarations/imports should be assumed to be propagated to the output.
//
// Precondition: None of the input files contain conflicting/duplicate definitions.
func merge(files []*ast.File) *ast.File {
	var decls []ast.Decl
	var imports []*ast.ImportSpec
	var unresolved []*ast.Ident

	for _, file := range files {
		decls = append(decls, file.Decls...)
		imports = append(imports, file.Imports...)
		unresolved = append(unresolved, file.Unresolved...)
	}

	return &ast.File{
		Doc:        nil,
		Package:    0,
		Name:       nil,
		Decls:      decls,
		Scope:      nil,
		Imports:    imports,
		Unresolved: unresolved,
		Comments:   nil,
	}
}

func fileListsEquivalent(a []*ast.File, b []*ast.File) (bool, string) {
	mergedFileA := merge(a)
	mergedFileB := merge(b)

	if eq, msg := FilesEquivalent(mergedFileA, mergedFileB); !eq {
		return eq, "file lists not equivalent: " + msg
	}
	return true, ""
}

func declListsEquivalent(a []ast.Decl, b []ast.Decl) (bool, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return false, fmt.Sprintf("length of declaration lists did not match: %s", msg)
	}

	badDeclsA, genDeclsA, funcDeclsA := splitDecls(a)
	badDeclsB, genDeclsB, funcDeclsB := splitDecls(b)

	if eq, msg := badDeclListsEquivalent(badDeclsA, badDeclsB); !eq {
		return eq, "list of bad declarations did not match: " + msg
	}

	if eq, msg := genDeclListsEquivalent(genDeclsA, genDeclsB); !eq {
		return eq, "list of generic declarations did not match: " + msg
	}

	if eq, msg := funcDeclListsEquivalent(funcDeclsA, funcDeclsB); !eq {
		return eq, "list of function declarations did not match: " + msg
	}

	return true, ""
}

func splitDecls(decls []ast.Decl) ([]*ast.BadDecl, []*ast.GenDecl, []*ast.FuncDecl) {
	var badDecls []*ast.BadDecl
	var genDecls []*ast.GenDecl
	var funcDecls []*ast.FuncDecl

	for _, decl := range decls {
		badDecl, ok := decl.(*ast.BadDecl)
		if ok {
			badDecls = append(badDecls, badDecl)
			continue
		}

		genDecl, ok := decl.(*ast.GenDecl)
		if ok {
			genDecls = append(genDecls, genDecl)
			continue
		}

		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok {
			funcDecls = append(funcDecls, funcDecl)
		}
	}

	return badDecls, genDecls, funcDecls
}

func badDeclListsEquivalent(a []*ast.BadDecl, b []*ast.BadDecl) (bool, string) {
	if len(a) > 0 {
		panic(fmt.Errorf("first source contained bad declaration: %v", a[0]))
	}
	if len(b) > 0 {
		panic(fmt.Errorf("second source contained bad declaration: %v", b[0]))
	}

	return true, ""
}

func genDeclListsEquivalent(a []*ast.GenDecl, b []*ast.GenDecl) (bool, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return false, fmt.Sprintf("number of generic declarations did not match: %s", msg)
	}

	sortGenDeclList(a)
	sortGenDeclList(b)

	if cmp, msg := compareGenDeclLists(a, b); cmp != 0 {
		return false, "generic declaration lists did not match: " + msg
	}

	return true, ""
}

func funcDeclListsEquivalent(a []*ast.FuncDecl, b []*ast.FuncDecl) (bool, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return false, fmt.Sprintf("number of function declarations did not match: %s", msg)
	}

	sortFuncDeclList(a)
	sortFuncDeclList(b)

	if cmp, msg := compareFuncDeclLists(a, b); cmp != 0 {
		return false, "function declaration lists did not match: " + msg
	}

	return true, ""
}

func unresolvedIdentifiersEquivalent(a []*ast.Ident, b []*ast.Ident) (bool, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return false, fmt.Sprintf("number of unresolved identifiers did not match: %s", msg)
	}

	sortIdentifierList(a)
	sortIdentifierList(b)

	if cmp, msg := compareIdentifierLists(a, b); cmp != 0 {
		return false, "unresolved identifier lists did not match: " + msg
	}

	return true, ""
}

func importsEquivalent(a []*ast.ImportSpec, b []*ast.ImportSpec) (bool, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return false, fmt.Sprintf("number of imports did not match: %s", msg)
	}

	sortImportList(a)
	sortImportList(b)

	if cmp, msg := compareImportSpecLists(a, b); cmp != 0 {
		return false, "import lists did not match: " + msg
	}

	return true, ""
}
