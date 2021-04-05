package eqgo

import (
	"go/ast"
	"sort"
)

// Helpers to sort language entities in a stable way so that collections can be compared without
// regard for their original order.

func sortGenDeclList(x *[]*ast.GenDecl) {
	for i := range *x {
		sortGenDecl((*x)[i])
	}

	sort.SliceStable(*x, func(i, j int) bool {
		cmp, _ := compareGenDecls((*x)[i], (*x)[j])
		return cmp <= 0
	})

	// Remove import specs
	y := (*x)[:0]
	for _, d := range *x {
		newSpecs := d.Specs[:0]
		for _, s := range d.Specs {
			if _, ok := s.(*ast.ImportSpec); ok {
				continue
			}

			newSpecs = append(newSpecs, s)
		}

		// Omit any declarations that are empty after import specs have been removed.
		if len(newSpecs) > 0 {
			d.Specs = newSpecs
			y = append(y, d)
		}
	}
	*x = y

	// Remove duplicate values
	y = (*x)[:0]
	for i, v := range *x {
		if i+1 >= len(*x) {
			y = append(y, v)
			continue
		}

		if cmp, _ := compareGenDecls(v, (*x)[i+1]); cmp != 0 {
			y = append(y, v)
		}
	}
	*x = y
}

func sortGenDecl(x *ast.GenDecl) {
	sortSpecList(&x.Specs)
}

func sortFuncDeclList(x *[]*ast.FuncDecl) {
	sort.SliceStable(*x, func(i, j int) bool {
		if cmp, _ := compareIdentifiers((*x)[i].Name, (*x)[j].Name); cmp != 0 {
			return cmp < 0
		}
		if cmp, _ := compareFunctionTypes((*x)[i].Type, (*x)[j].Type); cmp != 0 {
			return cmp < 0
		}
		if cmp, _ := compareFieldLists((*x)[i].Recv, (*x)[j].Recv); cmp != 0 {
			return cmp < 0
		}
		if cmp, _ := compareBlockStatements((*x)[i].Body, (*x)[j].Body); cmp != 0 {
			return cmp < 0
		}
		return true
	})

	// Remove duplicate values
	y := (*x)[:0]
	for i, v := range *x {
		if i+1 >= len(*x) {
			y = append(y, v)
			continue
		}

		if cmp, _ := compareFuncDecls(v, (*x)[i+1]); cmp != 0 {
			y = append(y, v)
		}
	}
	*x = y
}

func sortSpecList(x *[]ast.Spec) {
	for i := range *x {
		sortSpec((*x)[i])
	}

	sort.SliceStable(*x, func(i, j int) bool {
		if importA, ok := (*x)[i].(*ast.ImportSpec); ok {
			importB := (*x)[j].(*ast.ImportSpec)
			cmp, _ := compareImportSpecs(importA, importB)
			return cmp <= 0
		}

		if valA, ok := (*x)[i].(*ast.ValueSpec); ok {
			valB := (*x)[j].(*ast.ValueSpec)
			cmp, _ := compareValueSpecs(valA, valB)
			return cmp <= 0
		}

		typeA := (*x)[i].(*ast.TypeSpec)
		typeB := (*x)[j].(*ast.TypeSpec)
		cmp, _ := compareTypeSpecs(typeA, typeB)
		return cmp <= 0
	})

	// Remove duplicate values
	y := (*x)[:0]
	for i, v := range *x {
		if i+1 >= len(*x) {
			y = append(y, v)
			continue
		}

		if cmp, _ := compareSpecs(v, (*x)[i+1]); cmp != 0 {
			y = append(y, v)
		}
	}
	*x = y
}

func sortImportList(x *[]*ast.ImportSpec) {
	sort.SliceStable(*x, func(i, j int) bool {
		if cmp, _ := compareIdentifiers((*x)[i].Name, (*x)[j].Name); cmp != 0 {
			return cmp < 0
		}
		if cmp, _ := compareBasicLiterals((*x)[i].Path, (*x)[j].Path); cmp != 0 {
			return cmp < 0
		}
		return true
	})

	// Remove duplicate values
	y := (*x)[:0]
	for i, v := range *x {
		if i+1 >= len(*x) {
			y = append(y, v)
			continue
		}

		if cmp, _ := compareImportSpecs(v, (*x)[i+1]); cmp != 0 {
			y = append(y, v)
		}
	}
	*x = y
}

func sortSpec(x ast.Spec) {
	if importSpec, ok := x.(*ast.ImportSpec); ok {
		sortIdentifier(importSpec.Name)
		return
	}

	if valSpec, ok := x.(*ast.ValueSpec); ok {
		sortIdentifierList(&valSpec.Names, false)
		sortExpression(valSpec.Type)
		sortExpressionList(&valSpec.Values)
		return
	}

	if typeSpec, ok := x.(*ast.TypeSpec); ok {
		sortIdentifier(typeSpec.Name)
		sortExpression(typeSpec.Type)
		return
	}
}

func sortIdentifier(x *ast.Ident) {
	// TODO: (kevinb) should .Object be sorted?
}

func sortIdentifierList(x *[]*ast.Ident, removeDuplicates bool) {
	if removeDuplicates {
		for i := range *x {
			sortIdentifier((*x)[i])
		}
	}

	sort.SliceStable(*x, func(i, j int) bool {
		cmp, _ := compareIdentifiers((*x)[i], (*x)[j])
		return cmp <= 0
	})

	// Remove duplicate values
	y := (*x)[:0]
	for i, v := range *x {
		if i+1 >= len(*x) {
			y = append(y, v)
			continue
		}

		if cmp, _ := compareIdentifiers(v, (*x)[i+1]); cmp != 0 {
			y = append(y, v)
		}
	}
	*x = y
}

func sortExpressionList(x *[]ast.Expr) {
	for i := range *x {
		sortExpression((*x)[i])
	}

	sort.SliceStable(*x, func(i, j int) bool {
		cmp, _ := compareExpressions((*x)[i], (*x)[j])
		return cmp <= 0
	})

	// Remove duplicate values
	y := (*x)[:0]
	for i, v := range *x {
		if i+1 >= len(*x) {
			y = append(y, v)
			continue
		}

		if cmp, _ := compareExpressions(v, (*x)[i+1]); cmp != 0 {
			y = append(y, v)
		}
	}
	*x = y
}

func sortExpression(x ast.Expr) {
	if _, ok := x.(*ast.BadExpr); ok {
		return
	}
	if ident, ok := x.(*ast.Ident); ok {
		sortIdentifier(ident)
		return
	}
	if ellipsis, ok := x.(*ast.Ellipsis); ok {
		sortExpression(ellipsis.Elt)
		return
	}
	if _, ok := x.(*ast.BasicLit); ok {
		return
	}
	if _, ok := x.(*ast.FuncLit); ok {
		return
	}
	if compositeLit, ok := x.(*ast.CompositeLit); ok {
		sortExpression(compositeLit.Type)
		sortExpressionList(&compositeLit.Elts)
		return
	}
	if parenExpr, ok := x.(*ast.ParenExpr); ok {
		sortExpression(parenExpr.X)
		return
	}
	if selectorExpr, ok := x.(*ast.SelectorExpr); ok {
		sortExpression(selectorExpr.X)
		sortIdentifier(selectorExpr.Sel)
		return
	}
	if indexExpr, ok := x.(*ast.IndexExpr); ok {
		sortExpression(indexExpr.X)
		sortExpression(indexExpr.Index)
		return
	}
	if sliceExpr, ok := x.(*ast.SliceExpr); ok {
		sortExpression(sliceExpr.X)
		sortExpression(sliceExpr.Low)
		sortExpression(sliceExpr.High)
		sortExpression(sliceExpr.Max)
		return
	}
	if typeAssert, ok := x.(*ast.TypeAssertExpr); ok {
		sortExpression(typeAssert.X)
		sortExpression(typeAssert.Type)
		return
	}
	if callExpr, ok := x.(*ast.CallExpr); ok {
		sortExpression(callExpr.Fun)
		sortExpressionList(&callExpr.Args)
		return
	}
	if starExpr, ok := x.(*ast.StarExpr); ok {
		sortExpression(starExpr.X)
		return
	}
	if unaryExpr, ok := x.(*ast.UnaryExpr); ok {
		sortExpression(unaryExpr.X)
		return
	}
	if binaryExpr, ok := x.(*ast.BinaryExpr); ok {
		sortExpression(binaryExpr.X)
		return
	}
	if keyValueExpr, ok := x.(*ast.KeyValueExpr); ok {
		sortExpression(keyValueExpr.Key)
		sortExpression(keyValueExpr.Value)
		return
	}
	if arrayType, ok := x.(*ast.ArrayType); ok {
		sortExpression(arrayType.Len)
		sortExpression(arrayType.Elt)
		return
	}
	if _, ok := x.(*ast.StructType); ok {
		return
	}
	if _, ok := x.(*ast.FuncType); ok {
		return
	}
	if _, ok := x.(*ast.InterfaceType); ok {
		return
	}
	if mapType, ok := x.(*ast.MapType); ok {
		sortExpression(mapType.Key)
		sortExpression(mapType.Value)
		return
	}
	if chanType, ok := x.(*ast.ChanType); ok {
		sortExpression(chanType.Value)
		return
	}
}

func sortIndexForStatementType(x ast.Stmt) int {
	if _, ok := x.(*ast.BadStmt); ok {
		return 0
	}
	if _, ok := x.(*ast.DeclStmt); ok {
		return 1
	}
	if _, ok := x.(*ast.EmptyStmt); ok {
		return 2
	}
	if _, ok := x.(*ast.LabeledStmt); ok {
		return 3
	}
	if _, ok := x.(*ast.ExprStmt); ok {
		return 4
	}
	if _, ok := x.(*ast.SendStmt); ok {
		return 5
	}
	if _, ok := x.(*ast.IncDecStmt); ok {
		return 6
	}
	if _, ok := x.(*ast.AssignStmt); ok {
		return 7
	}
	if _, ok := x.(*ast.GoStmt); ok {
		return 8
	}
	if _, ok := x.(*ast.DeferStmt); ok {
		return 9
	}
	if _, ok := x.(*ast.ReturnStmt); ok {
		return 10
	}
	if _, ok := x.(*ast.BranchStmt); ok {
		return 11
	}
	if _, ok := x.(*ast.BlockStmt); ok {
		return 12
	}
	if _, ok := x.(*ast.IfStmt); ok {
		return 13
	}
	if _, ok := x.(*ast.CaseClause); ok {
		return 14
	}
	if _, ok := x.(*ast.SwitchStmt); ok {
		return 15
	}
	if _, ok := x.(*ast.TypeSwitchStmt); ok {
		return 16
	}
	if _, ok := x.(*ast.CommClause); ok {
		return 17
	}
	if _, ok := x.(*ast.SelectStmt); ok {
		return 18
	}
	if _, ok := x.(*ast.ForStmt); ok {
		return 19
	}
	if _, ok := x.(*ast.RangeStmt); ok {
		return 20
	}
	panic("unrecognized statement type")
}

func sortIndexForExpressionType(x ast.Expr) int {
	if _, ok := x.(*ast.BadExpr); ok {
		return 0
	}
	if _, ok := x.(*ast.Ident); ok {
		return 1
	}
	if _, ok := x.(*ast.Ellipsis); ok {
		return 2
	}
	if _, ok := x.(*ast.BasicLit); ok {
		return 3
	}
	if _, ok := x.(*ast.FuncLit); ok {
		return 4
	}
	if _, ok := x.(*ast.CompositeLit); ok {
		return 5
	}
	if _, ok := x.(*ast.ParenExpr); ok {
		return 6
	}
	if _, ok := x.(*ast.SelectorExpr); ok {
		return 7
	}
	if _, ok := x.(*ast.IndexExpr); ok {
		return 8
	}
	if _, ok := x.(*ast.SliceExpr); ok {
		return 9
	}
	if _, ok := x.(*ast.TypeAssertExpr); ok {
		return 10
	}
	if _, ok := x.(*ast.CallExpr); ok {
		return 11
	}
	if _, ok := x.(*ast.StarExpr); ok {
		return 12
	}
	if _, ok := x.(*ast.UnaryExpr); ok {
		return 13
	}
	if _, ok := x.(*ast.BinaryExpr); ok {
		return 14
	}
	if _, ok := x.(*ast.KeyValueExpr); ok {
		return 15
	}
	if _, ok := x.(*ast.ArrayType); ok {
		return 16
	}
	if _, ok := x.(*ast.StructType); ok {
		return 17
	}
	if _, ok := x.(*ast.FuncType); ok {
		return 18
	}
	if _, ok := x.(*ast.InterfaceType); ok {
		return 19
	}
	if _, ok := x.(*ast.MapType); ok {
		return 20
	}
	if _, ok := x.(*ast.ChanType); ok {
		return 21
	}
	panic("unrecognized expression type")
}
