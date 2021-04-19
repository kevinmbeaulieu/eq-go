package eqgo

import (
	"fmt"
	"go/ast"
	"go/token"
)

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
//     A message describing any differences found
// )
func PackagesEquivalent(a *ast.Package, fsetA *token.FileSet, b *ast.Package, fsetB *token.FileSet, f Formatter) (bool, string) {
	if a == nil || b == nil {
		panic(fmt.Errorf("missing package"))
	}

	equivalentPackageNameA = a.Name
	equivalentPackageNameB = b.Name

	mergeMode := ast.FilterUnassociatedComments | ast.FilterImportDuplicates
	mergedFileA := ast.MergePackageFiles(a, mergeMode)
	mergedFileB := ast.MergePackageFiles(b, mergeMode)

	cmp, root := compareFiles(mergedFileA, mergedFileB)
	eq := cmp == 0

	if f == nil {
		f = DefaultFormatter{
			LeftFSet:  fsetA,
			RightFSet: fsetB,
		}
	}
	return eq, f.Format(eq, root)
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
//     A message describing the differences found
// )
func FilesEquivalent(a *ast.File, fsetA *token.FileSet, b *ast.File, fsetB *token.FileSet, f Formatter) (bool, string) {
	cmp, root := compareFiles(a, b)
	eq := cmp == 0

	if f == nil {
		f = DefaultFormatter{
			LeftFSet:  fsetA,
			RightFSet: fsetB,
		}
	}
	return eq, f.Format(eq, root)
}
