package eqgo

import (
	"fmt"
	"go/ast"
	"go/token"
)

// Helpers for comparing language entities for equivalence.
//
// Trivia such as comments and source positions which do not affect the overall code's behavior
// are ignored in these comparisons.
//
// Each method returns an int representing the result of the comparison check and a string message
// providing more info about the result if the two inputs were not equivalent.

func compareInts(a int, b int) (int, string) {
	if a < b {
		return -1, fmt.Sprintf("%d < %d", a, b)
	}
	if a > b {
		return 1, fmt.Sprintf("%d > %d", a, b)
	}
	return 0, ""
}

func compareBools(a bool, b bool) (int, string) {
	if a == b {
		return 0, ""
	}
	if a {
		return 1, fmt.Sprintf("%t > %t", a, b)
	}
	return -1, fmt.Sprintf("%t < %t", a, b)
}

func compareStrings(a string, b string) (int, string) {
	if a < b {
		return -1, fmt.Sprintf("%s < %s", a, b)
	}
	if a > b {
		return 1, fmt.Sprintf("%s > %s", a, b)
	}
	return 0, ""
}

func compareTokens(a token.Token, b token.Token) (int, string) {
	return compareInts(int(a), int(b))
}

func compareIdentifiers(a *ast.Ident, b *ast.Ident) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareStrings(a.Name, b.Name); cmp != 0 {
		return cmp, msg
	}

	// TODO: (kevinb) should .Object be compared
	return 0, ""
}

func compareImportSpecs(a *ast.ImportSpec, b *ast.ImportSpec) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareIdentifiers(a.Name, b.Name); cmp != 0 {
		return cmp, fmt.Sprintf("import names do not match: %s", msg)
	}

	if cmp, msg := compareBasicLiteratures(a.Path, b.Path); cmp != 0 {
		return cmp, fmt.Sprintf("import paths do not match: %s", msg)
	}

	return 0, ""
}

func compareBasicLiteratures(a *ast.BasicLit, b *ast.BasicLit) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareTokens(a.Kind, b.Kind); cmp != 0 {
		return cmp, msg
	}

	return compareStrings(a.Value, b.Value)
}

func compareValueSpecs(a *ast.ValueSpec, b *ast.ValueSpec) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareIdentifierLists(a.Names, b.Names); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.Type, b.Type); cmp != 0 {
		return cmp, msg
	}

	return compareExpressionLists(a.Values, b.Values)
}

func compareIdentifierLists(a []*ast.Ident, b []*ast.Ident) (int, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return cmp, msg
	}

	for i := range a {
		if cmp, msg := compareIdentifiers(a[i], b[i]); cmp != 0 {
			return cmp, msg
		}
	}

	return 0, ""
}

func compareExpressionLists(a []ast.Expr, b []ast.Expr) (int, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return cmp, msg
	}

	for i := range a {
		if cmp, msg := compareExpressions(a[i], b[i]); cmp != 0 {
			return cmp, msg
		}
	}

	return 0, ""
}

func compareExpressions(a ast.Expr, b ast.Expr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareInts(sortIndexForExpressionType(a), sortIndexForExpressionType(b)); cmp != 0 {
		return cmp, msg
	}

	if badExprA, ok := a.(*ast.BadExpr); ok {
		panic(fmt.Sprintf("found bad expression: %v", badExprA))
	}

	if identA, ok := a.(*ast.Ident); ok {
		identB := b.(*ast.Ident)
		return compareIdentifiers(identA, identB)
	}

	if ellipsisA, ok := a.(*ast.Ellipsis); ok {
		ellipsisB := b.(*ast.Ellipsis)
		return compareEllipses(ellipsisA, ellipsisB)
	}

	if basicLitA, ok := a.(*ast.BasicLit); ok {
		basicLitB := b.(*ast.BasicLit)
		return compareBasicLiteratures(basicLitA, basicLitB)
	}

	if funcLitA, ok := a.(*ast.FuncLit); ok {
		funcLitB := b.(*ast.FuncLit)
		return compareFunctionLiteratures(funcLitA, funcLitB)
	}

	if compositeLitA, ok := a.(*ast.CompositeLit); ok {
		compositeLitB := b.(*ast.CompositeLit)
		return compareCompositeLiteratures(compositeLitA, compositeLitB)
	}

	if parenExprA, ok := a.(*ast.ParenExpr); ok {
		parenExprB := b.(*ast.ParenExpr)
		return compareParentheses(parenExprA, parenExprB)
	}

	if selectorExprA, ok := a.(*ast.SelectorExpr); ok {
		selectorExprB := b.(*ast.SelectorExpr)
		return compareSelectors(selectorExprA, selectorExprB)
	}

	if indexExprA, ok := a.(*ast.IndexExpr); ok {
		indexExprB := b.(*ast.IndexExpr)
		return compareIndexExpressions(indexExprA, indexExprB)
	}

	if sliceExprA, ok := a.(*ast.SliceExpr); ok {
		sliceExprB := b.(*ast.SliceExpr)
		return compareSliceExpressions(sliceExprA, sliceExprB)
	}

	if typeAssertA, ok := a.(*ast.TypeAssertExpr); ok {
		typeAssertB := b.(*ast.TypeAssertExpr)
		return compareTypeAssertions(typeAssertA, typeAssertB)
	}

	if callExprA, ok := a.(*ast.CallExpr); ok {
		callExprB := b.(*ast.CallExpr)
		return compareCallExpressions(callExprA, callExprB)
	}

	if starExprA, ok := a.(*ast.StarExpr); ok {
		starExprB := b.(*ast.StarExpr)
		return compareStarExpressions(starExprA, starExprB)
	}

	if unaryExprA, ok := a.(*ast.UnaryExpr); ok {
		unaryExprB := b.(*ast.UnaryExpr)
		return compareUnaryExpressions(unaryExprA, unaryExprB)
	}

	if binaryExprA, ok := a.(*ast.BinaryExpr); ok {
		binaryExprB := b.(*ast.BinaryExpr)
		return compareBinaryExpressions(binaryExprA, binaryExprB)
	}

	if keyValueExprA, ok := a.(*ast.KeyValueExpr); ok {
		keyValueExprB := b.(*ast.KeyValueExpr)
		return compareKeyValueExpressions(keyValueExprA, keyValueExprB)
	}

	if arrayTypeA, ok := a.(*ast.ArrayType); ok {
		arrayTypeB := b.(*ast.ArrayType)
		return compareArrayTypes(arrayTypeA, arrayTypeB)
	}

	if structTypeA, ok := a.(*ast.StructType); ok {
		structTypeB := b.(*ast.StructType)
		return compareStructTypes(structTypeA, structTypeB)
	}

	if funcTypeA, ok := a.(*ast.FuncType); ok {
		funcTypeB := b.(*ast.FuncType)
		return compareFunctionTypes(funcTypeA, funcTypeB)
	}

	if interfaceTypeA, ok := a.(*ast.InterfaceType); ok {
		interfaceTypeB := b.(*ast.InterfaceType)
		return compareInterfaceTypes(interfaceTypeA, interfaceTypeB)
	}

	if mapTypeA, ok := a.(*ast.MapType); ok {
		mapTypeB := b.(*ast.MapType)
		return compareMapTypes(mapTypeA, mapTypeB)
	}

	if chanTypeA, ok := a.(*ast.ChanType); ok {
		chanTypeB := b.(*ast.ChanType)
		return compareChannelTypes(chanTypeA, chanTypeB)
	}

	panic(fmt.Errorf("unrecognized expression type %v", a))
}

func compareTypeSpecs(a *ast.TypeSpec, b *ast.TypeSpec) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareIdentifiers(a.Name, b.Name); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.Type, b.Type)
}

func compareSpecs(a ast.Spec, b ast.Spec) (int, string) {
	if importSpecA, ok := a.(*ast.ImportSpec); ok {
		if importSpecB, ok := b.(*ast.ImportSpec); ok {
			return compareImportSpecs(importSpecA, importSpecB)
		}
		return -1, "mismatched spec types"
	}

	if valueSpecA, ok := a.(*ast.ValueSpec); ok {
		if valueSpecB, ok := b.(*ast.ValueSpec); ok {
			return compareValueSpecs(valueSpecA, valueSpecB)
		}
		if _, ok := b.(*ast.ImportSpec); ok {
			return 1, "mismatched spec types"
		}
		return -1, "mismatched spec types"
	}

	if typeSpecA, ok := a.(*ast.TypeSpec); ok {
		if typeSpecB, ok := b.(*ast.TypeSpec); ok {
			return compareTypeSpecs(typeSpecA, typeSpecB)
		}
		return 1, "mismatched spec types"
	}

	panic(fmt.Sprintf("unrecognized spec type: %v", a))
}

func compareSpecLists(a []ast.Spec, b []ast.Spec) (int, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return cmp, msg
	}

	for i := range a {
		if cmp, msg := compareSpecs(a[i], b[i]); cmp != 0 {
			return cmp, msg
		}
	}

	return 0, ""
}

func compareEllipses(a *ast.Ellipsis, b *ast.Ellipsis) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareExpressions(a.Elt, b.Elt)
}

func compareFunctionLiteratures(a *ast.FuncLit, b *ast.FuncLit) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareFunctionTypes(a.Type, b.Type); cmp != 0 {
		return cmp, msg
	}

	return compareBlockStatements(a.Body, b.Body)
}

func compareBlockStatements(a *ast.BlockStmt, b *ast.BlockStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareInts(len(a.List), len(b.List)); cmp != 0 {
		return cmp, msg
	}

	for i := range a.List {
		if cmp, msg := compareStatements(a.List[i], b.List[i]); cmp != 0 {
			return cmp, msg
		}
	}

	return 0, ""
}

func compareStatements(a ast.Stmt, b ast.Stmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	statementTypeA := sortIndexForStatementType(a)
	statementTypeB := sortIndexForStatementType(b)

	if cmp, msg := compareInts(statementTypeA, statementTypeB); cmp != 0 {
		return cmp, msg
	}

	switch statementTypeA {
	case 0:
		return compareBadStatements(a.(*ast.BadStmt), b.(*ast.BadStmt))
	case 1:
		return compareDeclStatements(a.(*ast.DeclStmt), b.(*ast.DeclStmt))
	case 2:
		return compareEmptyStatements(a.(*ast.EmptyStmt), b.(*ast.EmptyStmt))
	case 3:
		return compareLabeledStatements(a.(*ast.LabeledStmt), b.(*ast.LabeledStmt))
	case 4:
		return compareExpressionStatements(a.(*ast.ExprStmt), b.(*ast.ExprStmt))
	case 5:
		return compareSendStatements(a.(*ast.SendStmt), b.(*ast.SendStmt))
	case 6:
		return compareIncDecStatements(a.(*ast.IncDecStmt), b.(*ast.IncDecStmt))
	case 7:
		return compareAssignStatements(a.(*ast.AssignStmt), b.(*ast.AssignStmt))
	case 8:
		return compareGoStatements(a.(*ast.GoStmt), b.(*ast.GoStmt))
	case 9:
		return compareDeferStatements(a.(*ast.DeferStmt), b.(*ast.DeferStmt))
	case 10:
		return compareReturnStatements(a.(*ast.ReturnStmt), b.(*ast.ReturnStmt))
	case 11:
		return compareBranchStatements(a.(*ast.BranchStmt), b.(*ast.BranchStmt))
	case 12:
		return compareBlockStatements(a.(*ast.BlockStmt), b.(*ast.BlockStmt))
	case 13:
		return compareIfStatements(a.(*ast.IfStmt), b.(*ast.IfStmt))
	case 14:
		return compareCaseClauses(a.(*ast.CaseClause), b.(*ast.CaseClause))
	case 15:
		return compareSwitchStatements(a.(*ast.SwitchStmt), b.(*ast.SwitchStmt))
	case 16:
		return compareTypeSwitchStatements(a.(*ast.TypeSwitchStmt), b.(*ast.TypeSwitchStmt))
	case 17:
		return compareCommClauses(a.(*ast.CommClause), b.(*ast.CommClause))
	case 18:
		return compareSelectStatements(a.(*ast.SelectStmt), b.(*ast.SelectStmt))
	case 19:
		return compareForStatements(a.(*ast.ForStmt), b.(*ast.ForStmt))
	case 20:
		return compareRangeStatements(a.(*ast.RangeStmt), b.(*ast.RangeStmt))
	}
	panic(fmt.Sprintf("unrecognized statement type: %v", a))
}

func compareCompositeLiteratures(a *ast.CompositeLit, b *ast.CompositeLit) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Type, b.Type); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressionLists(a.Elts, b.Elts); cmp != 0 {
		return cmp, msg
	}

	return compareBools(a.Incomplete, b.Incomplete)
}

func compareParentheses(a *ast.ParenExpr, b *ast.ParenExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareExpressions(a.X, b.X)
}

func compareSelectors(a *ast.SelectorExpr, b *ast.SelectorExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareIdentifiers(a.Sel, b.Sel); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.X, b.X)
}

func compareIndexExpressions(a *ast.IndexExpr, b *ast.IndexExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Index, b.Index); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.X, b.X)
}

func compareSliceExpressions(a *ast.SliceExpr, b *ast.SliceExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.X, b.X); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.Low, b.Low); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.High, b.High); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.Max, b.Max); cmp != 0 {
		return cmp, msg
	}

	return compareBools(a.Slice3, b.Slice3)
}

func compareTypeAssertions(a *ast.TypeAssertExpr, b *ast.TypeAssertExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Type, b.Type); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.X, b.X)
}

func compareCallExpressions(a *ast.CallExpr, b *ast.CallExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Fun, b.Fun); cmp != 0 {
		return cmp, msg
	}

	return compareExpressionLists(a.Args, b.Args)
}

func compareStarExpressions(a *ast.StarExpr, b *ast.StarExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareExpressions(a.X, b.X)
}

func compareUnaryExpressions(a *ast.UnaryExpr, b *ast.UnaryExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareTokens(a.Op, b.Op); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.X, b.X)
}

func compareBinaryExpressions(a *ast.BinaryExpr, b *ast.BinaryExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareTokens(a.Op, b.Op); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.X, b.X); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.Y, b.Y)
}

func compareKeyValueExpressions(a *ast.KeyValueExpr, b *ast.KeyValueExpr) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Key, b.Key); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.Value, b.Value)
}

func compareArrayTypes(a *ast.ArrayType, b *ast.ArrayType) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Len, b.Len); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.Elt, b.Elt)
}

func compareStructTypes(a *ast.StructType, b *ast.StructType) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareBools(a.Incomplete, b.Incomplete); cmp != 0 {
		return cmp, msg
	}

	return compareFieldLists(a.Fields, b.Fields)
}

func compareFunctionTypes(a *ast.FuncType, b *ast.FuncType) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareFieldLists(a.Params, b.Params); cmp != 0 {
		return cmp, fmt.Sprintf("function type parameters do not match: %s", msg)
	}

	if cmp, msg := compareFieldLists(a.Results, b.Results); cmp != 0 {
		return cmp, fmt.Sprintf("function type results do not match: %s", msg)
	}

	return 0, ""
}

func compareFieldLists(a *ast.FieldList, b *ast.FieldList) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareInts(len(a.List), len(b.List)); cmp != 0 {
		return cmp, fmt.Sprintf("length of field lists do not match: %s", msg)
	}

	for i := range a.List {
		if cmp, msg := compareFields(a.List[i], b.List[i]); cmp != 0 {
			return cmp, fmt.Sprintf("fields at index %d do not match: %s", i, msg)
		}
	}

	return 0, ""
}

func compareFields(a *ast.Field, b *ast.Field) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareIdentifierLists(a.Names, b.Names); cmp != 0 {
		return cmp, fmt.Sprintf("field names do not match: %s", msg)
	}

	if cmp, msg := compareExpressions(a.Type, b.Type); cmp != 0 {
		return cmp, fmt.Sprintf("field types do not match: %s", msg)
	}

	if cmp, msg := compareBasicLiteratures(a.Tag, b.Tag); cmp != 0 {
		return cmp, fmt.Sprintf("field tags do not match: %s", msg)
	}

	return 0, ""
}

func compareInterfaceTypes(a *ast.InterfaceType, b *ast.InterfaceType) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareFieldLists(a.Methods, b.Methods); cmp != 0 {
		return cmp, msg
	}

	return compareBools(a.Incomplete, b.Incomplete)
}

func compareMapTypes(a *ast.MapType, b *ast.MapType) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Key, b.Key); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.Value, b.Value)
}

func compareChannelTypes(a *ast.ChanType, b *ast.ChanType) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareChannelDirections(a.Dir, b.Dir); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.Value, b.Value)
}

func compareChannelDirections(a ast.ChanDir, b ast.ChanDir) (int, string) {
	return compareInts(int(a), int(b))
}

func compareBadStatements(a *ast.BadStmt, b *ast.BadStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return 0, ""
}

func compareDeclStatements(a *ast.DeclStmt, b *ast.DeclStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareDecls(a.Decl, b.Decl)
}

func compareDecls(a ast.Decl, b ast.Decl) (int, string) {
	if badDeclA, ok := a.(*ast.BadDecl); ok {
		if badDeclB, ok := b.(*ast.BadDecl); ok {
			return compareBadDecls(badDeclA, badDeclB)
		}
		return -1, "mismatched declaration types"
	}
	if genDeclA, ok := a.(*ast.GenDecl); ok {
		if genDeclB, ok := b.(*ast.GenDecl); ok {
			return compareGenDecls(genDeclA, genDeclB)
		}
		if _, ok := b.(*ast.BadDecl); ok {
			return 1, "mismatched declaration types"
		}
		return -1, "mismatched declaration types"
	}
	if funcDeclA, ok := a.(*ast.FuncDecl); ok {
		if funcDeclB, ok := b.(*ast.FuncDecl); ok {
			return compareFuncDecls(funcDeclA, funcDeclB)
		}
		return 1, "mismatched declaration types"
	}
	panic(fmt.Sprintf("unrecognized declaration type: %v", a))
}

func compareBadDecls(a *ast.BadDecl, b *ast.BadDecl) (int, string) {
	return compareBools(a == nil, b == nil)
}

func compareGenDecls(a *ast.GenDecl, b *ast.GenDecl) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareTokens(a.Tok, b.Tok); cmp != 0 {
		return cmp, msg
	}

	return compareSpecLists(a.Specs, b.Specs)
}

func compareFuncDecls(a *ast.FuncDecl, b *ast.FuncDecl) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareIdentifiers(a.Name, b.Name); cmp != 0 {
		return cmp, fmt.Sprintf("function names do not match: %s", msg)
	}

	if cmp, msg := compareFieldLists(a.Recv, b.Recv); cmp != 0 {
		return cmp, fmt.Sprintf("function receivers do not match: %s", msg)
	}

	if cmp, msg := compareFunctionTypes(a.Type, b.Type); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		return cmp, fmt.Sprintf("function bodies do not match: %s", msg)
	}

	return 0, ""
}

func compareEmptyStatements(a *ast.EmptyStmt, b *ast.EmptyStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareBools(a.Implicit, b.Implicit)
}

func compareLabeledStatements(a *ast.LabeledStmt, b *ast.LabeledStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareIdentifiers(a.Label, b.Label); cmp != 0 {
		return cmp, msg
	}

	return compareStatements(a.Stmt, b.Stmt)
}

func compareExpressionStatements(a *ast.ExprStmt, b *ast.ExprStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareExpressions(a.X, b.X)
}

func compareSendStatements(a *ast.SendStmt, b *ast.SendStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Chan, b.Chan); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.Value, b.Value)
}

func compareIncDecStatements(a *ast.IncDecStmt, b *ast.IncDecStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareTokens(a.Tok, b.Tok); cmp != 0 {
		return cmp, msg
	}

	return compareExpressions(a.X, b.X)
}

func compareAssignStatements(a *ast.AssignStmt, b *ast.AssignStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareTokens(a.Tok, b.Tok); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressionLists(a.Lhs, b.Lhs); cmp != 0 {
		return cmp, msg
	}

	return compareExpressionLists(a.Rhs, b.Rhs)
}

func compareGoStatements(a *ast.GoStmt, b *ast.GoStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareCallExpressions(a.Call, b.Call)
}

func compareDeferStatements(a *ast.DeferStmt, b *ast.DeferStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareCallExpressions(a.Call, b.Call)
}

func compareReturnStatements(a *ast.ReturnStmt, b *ast.ReturnStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareExpressionLists(a.Results, b.Results)
}

func compareBranchStatements(a *ast.BranchStmt, b *ast.BranchStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareTokens(a.Tok, b.Tok); cmp != 0 {
		return cmp, msg
	}

	return compareIdentifiers(a.Label, b.Label)
}

func compareIfStatements(a *ast.IfStmt, b *ast.IfStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareStatements(a.Init, b.Init); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.Cond, b.Cond); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		return cmp, msg
	}

	return compareStatements(a.Else, b.Else)
}

func compareCaseClauses(a *ast.CaseClause, b *ast.CaseClause) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressionLists(a.List, b.List); cmp != 0 {
		return cmp, msg
	}

	return compareStatementLists(a.Body, b.Body)
}

func compareStatementLists(a []ast.Stmt, b []ast.Stmt) (int, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return cmp, msg
	}

	for i := range a {
		if cmp, msg := compareStatements(a[i], b[i]); cmp != 0 {
			return cmp, msg
		}
	}

	return 0, ""
}

func compareSwitchStatements(a *ast.SwitchStmt, b *ast.SwitchStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareStatements(a.Init, b.Init); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.Tag, b.Tag); cmp != 0 {
		return cmp, msg
	}

	return compareBlockStatements(a.Body, b.Body)
}

func compareTypeSwitchStatements(a *ast.TypeSwitchStmt, b *ast.TypeSwitchStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareStatements(a.Init, b.Init); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareStatements(a.Assign, b.Assign); cmp != 0 {
		return cmp, msg
	}

	return compareBlockStatements(a.Body, b.Body)
}

func compareCommClauses(a *ast.CommClause, b *ast.CommClause) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareStatements(a.Comm, b.Comm); cmp != 0 {
		return cmp, msg
	}

	return compareStatementLists(a.Body, b.Body)
}

func compareSelectStatements(a *ast.SelectStmt, b *ast.SelectStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	return compareBlockStatements(a.Body, b.Body)
}

func compareForStatements(a *ast.ForStmt, b *ast.ForStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareStatements(a.Init, b.Init); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.Cond, b.Cond); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareStatements(a.Post, b.Post); cmp != 0 {
		return cmp, msg
	}

	return compareBlockStatements(a.Body, b.Body)
}

func compareRangeStatements(a *ast.RangeStmt, b *ast.RangeStmt) (int, string) {
	if a == nil || b == nil {
		return compareBools(a == nil, b == nil)
	}

	if cmp, msg := compareExpressions(a.Key, b.Key); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.Value, b.Value); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareTokens(a.Tok, b.Tok); cmp != 0 {
		return cmp, msg
	}

	if cmp, msg := compareExpressions(a.X, b.X); cmp != 0 {
		return cmp, msg
	}

	return compareBlockStatements(a.Body, b.Body)
}

func compareGenDeclLists(a []*ast.GenDecl, b []*ast.GenDecl) (int, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return cmp, msg
	}

	for i := range a {
		if cmp, msg := compareGenDecls(a[i], b[i]); cmp != 0 {
			return cmp, msg
		}
	}

	return 0, ""
}

func compareFuncDeclLists(a []*ast.FuncDecl, b []*ast.FuncDecl) (int, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return cmp, fmt.Sprintf("length of function declaration lists do not match: %s", msg)
	}

	for i := range a {
		if cmp, msg := compareFuncDecls(a[i], b[i]); cmp != 0 {
			return cmp, fmt.Sprintf("function declarations at index %d do not match: %s", i, msg)
		}
	}

	return 0, ""
}

func compareImportSpecLists(a []*ast.ImportSpec, b []*ast.ImportSpec) (int, string) {
	if cmp, msg := compareInts(len(a), len(b)); cmp != 0 {
		return cmp, fmt.Sprintf("length of import lists do not match: %s", msg)
	}

	for i := range a {
		if cmp, msg := compareImportSpecs(a[i], b[i]); cmp != 0 {
			return cmp, fmt.Sprintf("imports at index %d do not match: %s", i, msg)
		}
	}

	return 0, ""
}
