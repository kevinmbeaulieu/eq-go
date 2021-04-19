package eqgo

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

// Functions for comparing language entities for equivalence.
//
// Trivia such as comments and source positions which do not affect the overall code's behavior
// are ignored in these comparisons.
//
// Each method returns an int representing the result of the comparison check and a tree node
// providing more info about the differences if the two inputs were not equivalent.

// TODO: (kevinb) Refactor to avoid using global variables for these.
var equivalentPackageNameA string
var equivalentPackageNameB string

type node struct {
	msg      string
	leftPos  token.Pos
	rightPos token.Pos
	children []*node
}

// Developer-friendly string representation of a node.
func (n node) String() string {
	var builder strings.Builder
	fmt.Fprintf(
		&builder,
		"{%s (lhs offset: %d; rhs offset: %d); children: %v",
		n.msg,
		n.leftPos,
		n.rightPos,
		n.children,
	)
	return builder.String()
}

// Create and return a pointer to a new node.
func newNode(msg string, left ast.Node, right ast.Node, children *[]*node) *node {
	var c []*node
	if children != nil {
		c = *children
	}

	isNil := func(x ast.Node) bool {
		if x == nil {
			return true
		}
		switch reflect.TypeOf(x).Kind() {
		case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
			return reflect.ValueOf(x).IsNil()
		default:
			return false
		}
	}

	var leftPos, rightPos token.Pos
	if !isNil(left) {
		leftPos = left.Pos()
	}
	if !isNil(right) {
		rightPos = right.Pos()
	}
	n := node{
		msg:      msg,
		leftPos:  leftPos,
		rightPos: rightPos,
		children: c,
	}
	return &n
}

// Construct an (int, *node) tuple to return from one of the compare* functions.
func newRetVal(cmp int, errMsg string, left ast.Node, right ast.Node, children []*node) (int, *node) {
	if cmp == 0 {
		return cmp, nil
	}
	return cmp, newNode(errMsg, left, right, &children)
}

// Construct an (int, *node) tuple to return from one of the compare* functions.
// Precondition: at least one of `a` and `b` is nil.
func newNilRetVal(a interface{}, b interface{}, msg string) (int, *node) {
	isNil := func(x interface{}) bool {
		if x == nil {
			return true
		}
		switch reflect.TypeOf(x).Kind() {
		case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
			return reflect.ValueOf(x).IsNil()
		default:
			return false
		}
	}

	cmp, child := compareBools(isNil(a), isNil(b))
	return newRetVal(
		cmp,
		msg,
		nil,
		nil,
		[]*node{
			newNode("nil comparisons did not match", nil, nil, &[]*node{child}),
		},
	)
}

// Set the value of the pointer to `val` if the pointer is currently set to 0.
func setIfUnset(x *int, val int) {
	if *x == 0 {
		*x = val
	}
}

func compareInts(a int, b int) (int, *node) {
	if a < b {
		return -1, newNode(fmt.Sprintf("ints did not match: %d < %d", a, b), nil, nil, nil)
	} else if a > b {
		return 1, newNode(fmt.Sprintf("ints did not match: %d > %d", a, b), nil, nil, nil)
	}
	return 0, nil
}

func compareBools(a bool, b bool) (int, *node) {
	if !a && b {
		return -1, newNode(fmt.Sprintf("bools did not match: %t < %t", a, b), nil, nil, nil)
	} else if a && !b {
		return 1, newNode(fmt.Sprintf("bools did not match: %t > %t", a, b), nil, nil, nil)
	}
	return 0, nil
}

func compareStrings(a string, b string) (int, *node) {
	if equivalentPackageNameA != "" && equivalentPackageNameB != "" {
		a = strings.ReplaceAll(a, equivalentPackageNameB, equivalentPackageNameA)
		b = strings.ReplaceAll(b, equivalentPackageNameB, equivalentPackageNameA)
	}

	if a < b {
		return -1, newNode(fmt.Sprintf("strings did not match: %s < %s", a, b), nil, nil, nil)
	} else if a > b {
		return 1, newNode(fmt.Sprintf("strings did not match: %s > %s", a, b), nil, nil, nil)
	}
	return 0, nil
}

func compareTokens(a token.Token, b token.Token) (int, *node) {
	cmp, child := compareInts(int(a), int(b))
	return newRetVal(cmp, "tokens did not match", nil, nil, []*node{child})
}

func compareIdentifiers(a *ast.Ident, b *ast.Ident) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "identifiers did not match")
	}

	cmp, child := compareStrings(a.Name, b.Name)
	// TODO: (kevinb) should .Object be compared
	return newRetVal(cmp, "identifiers did not match", nil, nil, []*node{child})
}

func compareImportSpecs(a *ast.ImportSpec, b *ast.ImportSpec) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "import specs did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareIdentifiers(a.Name, b.Name); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("names did not match", a.Name, b.Name, &[]*node{child}))
	}

	if cmp, child := compareBasicLiterals(a.Path, b.Path); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("paths did not match", a.Path, b.Path, &[]*node{child}))
	}

	return newRetVal(retCmp, "import specs did not match", a, b, children)
}

func compareBasicLiterals(a *ast.BasicLit, b *ast.BasicLit) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "basic literals did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareTokens(a.Kind, b.Kind); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("kinds did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareStrings(a.Value, b.Value); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("values did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "basic literals did not match", a, b, children)
}

func compareValueSpecs(a *ast.ValueSpec, b *ast.ValueSpec) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "value specs did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareIdentifierLists(a.Names, b.Names); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("name lists did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Type, b.Type); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("types did not match", a.Type, b.Type, &[]*node{child}))
	}

	if cmp, child := compareExpressionLists(a.Values, b.Values); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("values did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "value specs did not match", a, b, children)
}

func compareIdentifierLists(a []*ast.Ident, b []*ast.Ident) (int, *node) {
	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a), len(b)); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a {
		if i >= len(b) {
			break
		}
		if cmp, child := compareIdentifiers(a[i], b[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)

			children = append(children, newNode(fmt.Sprintf("identifiers at index %d did not match", i), a[i], b[i], &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "identifier lists did not match", nil, nil, children)
}

func compareExpressionLists(a []ast.Expr, b []ast.Expr) (int, *node) {
	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a), len(b)); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a {
		if i >= len(b) {
			break
		}
		if cmp, child := compareExpressions(a[i], b[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)

			children = append(children, newNode(fmt.Sprintf("expressions at index %d did not match", i), a[i], b[i], &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "expression lists did not match", nil, nil, children)
}

func compareExpressions(a ast.Expr, b ast.Expr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "expressions did not match")
	}

	if cmp, child := compareInts(sortIndexForExpressionType(a), sortIndexForExpressionType(b)); cmp != 0 {
		return newRetVal(
			cmp,
			"expressions did not match",
			a,
			b,
			[]*node{
				newNode("sort indices did not match", nil, nil, &[]*node{child}),
			},
		)
	}

	if badExprA, ok := a.(*ast.BadExpr); ok {
		panic(fmt.Sprintf("found bad expression: %v", badExprA))
	}

	retCmp := 0
	var children []*node

	if identA, ok := a.(*ast.Ident); ok {
		identB := b.(*ast.Ident)
		if cmp, child := compareIdentifiers(identA, identB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if ellipsisA, ok := a.(*ast.Ellipsis); ok {
		ellipsisB := b.(*ast.Ellipsis)
		if cmp, child := compareEllipses(ellipsisA, ellipsisB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if basicLitA, ok := a.(*ast.BasicLit); ok {
		basicLitB := b.(*ast.BasicLit)
		if cmp, child := compareBasicLiterals(basicLitA, basicLitB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if funcLitA, ok := a.(*ast.FuncLit); ok {
		funcLitB := b.(*ast.FuncLit)
		if cmp, child := compareFunctionLiterals(funcLitA, funcLitB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if compositeLitA, ok := a.(*ast.CompositeLit); ok {
		compositeLitB := b.(*ast.CompositeLit)
		if cmp, child := compareCompositeLiterals(compositeLitA, compositeLitB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if parenExprA, ok := a.(*ast.ParenExpr); ok {
		parenExprB := b.(*ast.ParenExpr)
		if cmp, child := compareParentheses(parenExprA, parenExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if selectorExprA, ok := a.(*ast.SelectorExpr); ok {
		selectorExprB := b.(*ast.SelectorExpr)
		if cmp, child := compareSelectors(selectorExprA, selectorExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if indexExprA, ok := a.(*ast.IndexExpr); ok {
		indexExprB := b.(*ast.IndexExpr)
		if cmp, child := compareIndexExpressions(indexExprA, indexExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if sliceExprA, ok := a.(*ast.SliceExpr); ok {
		sliceExprB := b.(*ast.SliceExpr)
		if cmp, child := compareSliceExpressions(sliceExprA, sliceExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if typeAssertA, ok := a.(*ast.TypeAssertExpr); ok {
		typeAssertB := b.(*ast.TypeAssertExpr)
		if cmp, child := compareTypeAssertions(typeAssertA, typeAssertB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if callExprA, ok := a.(*ast.CallExpr); ok {
		callExprB := b.(*ast.CallExpr)
		if cmp, child := compareCallExpressions(callExprA, callExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if starExprA, ok := a.(*ast.StarExpr); ok {
		starExprB := b.(*ast.StarExpr)
		if cmp, child := compareStarExpressions(starExprA, starExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if unaryExprA, ok := a.(*ast.UnaryExpr); ok {
		unaryExprB := b.(*ast.UnaryExpr)
		if cmp, child := compareUnaryExpressions(unaryExprA, unaryExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if binaryExprA, ok := a.(*ast.BinaryExpr); ok {
		binaryExprB := b.(*ast.BinaryExpr)
		if cmp, child := compareBinaryExpressions(binaryExprA, binaryExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if keyValueExprA, ok := a.(*ast.KeyValueExpr); ok {
		keyValueExprB := b.(*ast.KeyValueExpr)
		if cmp, child := compareKeyValueExpressions(keyValueExprA, keyValueExprB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if arrayTypeA, ok := a.(*ast.ArrayType); ok {
		arrayTypeB := b.(*ast.ArrayType)
		if cmp, child := compareArrayTypes(arrayTypeA, arrayTypeB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if structTypeA, ok := a.(*ast.StructType); ok {
		structTypeB := b.(*ast.StructType)
		if cmp, child := compareStructTypes(structTypeA, structTypeB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if funcTypeA, ok := a.(*ast.FuncType); ok {
		funcTypeB := b.(*ast.FuncType)
		if cmp, child := compareFunctionTypes(funcTypeA, funcTypeB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if interfaceTypeA, ok := a.(*ast.InterfaceType); ok {
		interfaceTypeB := b.(*ast.InterfaceType)
		if cmp, child := compareInterfaceTypes(interfaceTypeA, interfaceTypeB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if mapTypeA, ok := a.(*ast.MapType); ok {
		mapTypeB := b.(*ast.MapType)
		if cmp, child := compareMapTypes(mapTypeA, mapTypeB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	if chanTypeA, ok := a.(*ast.ChanType); ok {
		chanTypeB := b.(*ast.ChanType)
		if cmp, child := compareChannelTypes(chanTypeA, chanTypeB); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	return newRetVal(retCmp, "expressions did not match", a, b, children)
}

func compareTypeSpecs(a *ast.TypeSpec, b *ast.TypeSpec) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "type specs did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareIdentifiers(a.Name, b.Name); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("names did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Type, b.Type); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("types did not match", a.Type, b.Type, &[]*node{child}))
	}

	return newRetVal(retCmp, "type specs did not match", a, b, children)
}

func compareSpecs(a ast.Spec, b ast.Spec) (int, *node) {
	if importSpecA, ok := a.(*ast.ImportSpec); ok {
		if importSpecB, ok := b.(*ast.ImportSpec); ok {
			cmp, child := compareImportSpecs(importSpecA, importSpecB)
			return newRetVal(cmp, "specs did not match", a, b, []*node{child})
		}
		return newRetVal(-1, "spec types did not match", nil, nil, nil)
	}

	if valueSpecA, ok := a.(*ast.ValueSpec); ok {
		if valueSpecB, ok := b.(*ast.ValueSpec); ok {
			cmp, child := compareValueSpecs(valueSpecA, valueSpecB)
			return newRetVal(cmp, "specs did not match", a, b, []*node{child})
		}
		if _, ok := b.(*ast.ImportSpec); ok {
			return newRetVal(1, "spec types did not match", nil, nil, nil)
		}
		return newRetVal(-1, "spec types did not match", nil, nil, nil)
	}

	if typeSpecA, ok := a.(*ast.TypeSpec); ok {
		if typeSpecB, ok := b.(*ast.TypeSpec); ok {
			cmp, child := compareTypeSpecs(typeSpecA, typeSpecB)
			return newRetVal(cmp, "specs did not match", a, b, []*node{child})
		}
		return newRetVal(1, "spec types did not match", nil, nil, nil)
	}

	panic(fmt.Sprintf("unrecognized spec type: %v", a))
}

func compareSpecLists(a []ast.Spec, b []ast.Spec) (int, *node) {
	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a), len(b)); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a {
		if i >= len(b) {
			break
		}
		if cmp, child := compareSpecs(a[i], b[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)

			children = append(children, newNode(fmt.Sprintf("specs at index %d did not match", i), nil, nil, &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "spec lists did not match", nil, nil, children)
}

func compareEllipses(a *ast.Ellipsis, b *ast.Ellipsis) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "ellipses did not match")
	}

	cmp, child := compareExpressions(a.Elt, b.Elt)
	return newRetVal(
		cmp,
		"ellipses did not match",
		a,
		b,
		[]*node{
			newNode(
				"expressions did not match",
				a.Elt,
				b.Elt,
				&[]*node{child},
			),
		},
	)
}

func compareFunctionLiterals(a *ast.FuncLit, b *ast.FuncLit) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "function literals did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareFunctionTypes(a.Type, b.Type); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("types did not match", a.Type, b.Type, &[]*node{child}))
	}

	if cmp, child := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("bodies did not match", a.Body, b.Body, &[]*node{child}))
	}

	return newRetVal(retCmp, "function literals did not match", a, b, children)
}

func compareBlockStatements(a *ast.BlockStmt, b *ast.BlockStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "block statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a.List), len(b.List)); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a.List {
		if i >= len(b.List) {
			break
		}
		if cmp, child := compareStatements(a.List[i], b.List[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)

			children = append(children, newNode(fmt.Sprintf("statements at index %d did not match", i), a.List[i], b.List[i], &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "block statements did not match", a, b, children)
}

func compareStatements(a ast.Stmt, b ast.Stmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "statements did not match")
	}

	statementTypeA := sortIndexForStatementType(a)
	statementTypeB := sortIndexForStatementType(b)

	if cmp, child := compareInts(statementTypeA, statementTypeB); cmp != 0 {
		return newRetVal(
			cmp,
			"statements did not match",
			nil,
			nil,
			[]*node{
				newNode("statement types did not match", nil, nil, &[]*node{child}),
			},
		)
	}

	retCmp := 0
	var children []*node

	switch statementTypeA {
	case 0:
		if cmp, child := compareBadStatements(a.(*ast.BadStmt), b.(*ast.BadStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 1:
		if cmp, child := compareDeclStatements(a.(*ast.DeclStmt), b.(*ast.DeclStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 2:
		if cmp, child := compareEmptyStatements(a.(*ast.EmptyStmt), b.(*ast.EmptyStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 3:
		if cmp, child := compareLabeledStatements(a.(*ast.LabeledStmt), b.(*ast.LabeledStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 4:
		if cmp, child := compareExpressionStatements(a.(*ast.ExprStmt), b.(*ast.ExprStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 5:
		if cmp, child := compareSendStatements(a.(*ast.SendStmt), b.(*ast.SendStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 6:
		if cmp, child := compareIncDecStatements(a.(*ast.IncDecStmt), b.(*ast.IncDecStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 7:
		if cmp, child := compareAssignStatements(a.(*ast.AssignStmt), b.(*ast.AssignStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 8:
		if cmp, child := compareGoStatements(a.(*ast.GoStmt), b.(*ast.GoStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 9:
		if cmp, child := compareDeferStatements(a.(*ast.DeferStmt), b.(*ast.DeferStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 10:
		if cmp, child := compareReturnStatements(a.(*ast.ReturnStmt), b.(*ast.ReturnStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 11:
		if cmp, child := compareBranchStatements(a.(*ast.BranchStmt), b.(*ast.BranchStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 12:
		if cmp, child := compareBlockStatements(a.(*ast.BlockStmt), b.(*ast.BlockStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 13:
		if cmp, child := compareIfStatements(a.(*ast.IfStmt), b.(*ast.IfStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 14:
		if cmp, child := compareCaseClauses(a.(*ast.CaseClause), b.(*ast.CaseClause)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 15:
		if cmp, child := compareSwitchStatements(a.(*ast.SwitchStmt), b.(*ast.SwitchStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 16:
		if cmp, child := compareTypeSwitchStatements(a.(*ast.TypeSwitchStmt), b.(*ast.TypeSwitchStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 17:
		if cmp, child := compareCommClauses(a.(*ast.CommClause), b.(*ast.CommClause)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 18:
		if cmp, child := compareSelectStatements(a.(*ast.SelectStmt), b.(*ast.SelectStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 19:
		if cmp, child := compareForStatements(a.(*ast.ForStmt), b.(*ast.ForStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	case 20:
		if cmp, child := compareRangeStatements(a.(*ast.RangeStmt), b.(*ast.RangeStmt)); cmp != 0 {
			retCmp = cmp
			children = append(children, child)
		}
	}

	return newRetVal(retCmp, "statements did not match", a, b, children)
}

func compareCompositeLiterals(a *ast.CompositeLit, b *ast.CompositeLit) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "composite literals did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Type, b.Type); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("types did not match", a.Type, b.Type, &[]*node{child}))
	}

	if cmp, child := compareExpressionLists(a.Elts, b.Elts); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("expression lists did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareBools(a.Incomplete, b.Incomplete); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("incomplete values did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "composite literals did not match", a, b, children)
}

func compareParentheses(a *ast.ParenExpr, b *ast.ParenExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "parentheses did not match")
	}

	cmp, child := compareExpressions(a.X, b.X)
	return newRetVal(
		cmp,
		"parentheses did not match",
		a,
		b,
		[]*node{
			newNode("expressions did not match", a.X, b.X, &[]*node{child}),
		},
	)
}

func compareSelectors(a *ast.SelectorExpr, b *ast.SelectorExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "selector expressions did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareIdentifiers(a.Sel, b.Sel); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("selectors did not match", a.Sel, b.Sel, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("expressions did not match", a.X, b.X, &[]*node{child}))
	}

	return newRetVal(retCmp, "selector expressions did not match", a, b, children)
}

func compareIndexExpressions(a *ast.IndexExpr, b *ast.IndexExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "index expressions did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Index, b.Index); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("indices did not match", a.Index, b.Index, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)

		children = append(children, newNode("expressions did not match", a.X, b.X, &[]*node{child}))
	}

	return newRetVal(retCmp, "index expressions did not match", a, b, children)
}

func compareSliceExpressions(a *ast.SliceExpr, b *ast.SliceExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "slices did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("X expressions did not match", a.X, b.X, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Low, b.Low); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("low expressions did not match", a.Low, b.Low, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.High, b.High); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("high expressions did not match", a.High, b.High, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Max, b.Max); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("max expressions did not match", a.Max, b.Max, &[]*node{child}))
	}

	if cmp, child := compareBools(a.Slice3, b.Slice3); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("slice3 values did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "slices did not match", a, b, children)
}

func compareTypeAssertions(a *ast.TypeAssertExpr, b *ast.TypeAssertExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "type assertions did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Type, b.Type); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("types did not match", a.Type, b.Type, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("X expressions did not match", a.X, b.X, &[]*node{child}))
	}

	return newRetVal(retCmp, "type assertions did not match", a, b, children)
}

func compareCallExpressions(a *ast.CallExpr, b *ast.CallExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "call expressions did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Fun, b.Fun); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("functions did not match", a.Fun, b.Fun, &[]*node{child}))
	}

	if cmp, child := compareExpressionLists(a.Args, b.Args); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("arguments did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "call expressions did not match", a, b, children)
}

func compareStarExpressions(a *ast.StarExpr, b *ast.StarExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "star expressions did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("X expressions did not match", a.X, b.X, &[]*node{child}))
	}

	return newRetVal(retCmp, "star expressions did not match", a, b, children)
}

func compareUnaryExpressions(a *ast.UnaryExpr, b *ast.UnaryExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "unary expressions did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareTokens(a.Op, b.Op); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("operators did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("X expressions did not match", a.X, b.X, &[]*node{child}))
	}

	return newRetVal(retCmp, "unary expressions did not match", a, b, children)
}

func compareBinaryExpressions(a *ast.BinaryExpr, b *ast.BinaryExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "binary expressions did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareTokens(a.Op, b.Op); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("operators did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("X expressions did not match", a.X, b.X, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Y, b.Y); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("Y expressions did not match", a.Y, b.Y, &[]*node{child}))
	}

	return newRetVal(retCmp, "binary expressions did not match", a, b, children)
}

func compareKeyValueExpressions(a *ast.KeyValueExpr, b *ast.KeyValueExpr) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "key-value expressions did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Key, b.Key); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("keys did not match", a.Key, b.Key, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Value, b.Value); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("values did not match", a.Value, b.Value, &[]*node{child}))
	}

	return newRetVal(retCmp, "key-value expressions did not match", a, b, children)
}

func compareArrayTypes(a *ast.ArrayType, b *ast.ArrayType) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "array types did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Len, b.Len); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("length expressions did not match", a.Len, b.Len, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Elt, b.Elt); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("element type expressions did not match", a.Elt, b.Elt, &[]*node{child}))
	}

	return newRetVal(retCmp, "array types did not match", a, b, children)
}

func compareStructTypes(a *ast.StructType, b *ast.StructType) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "struct types did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareBools(a.Incomplete, b.Incomplete); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("incomplete values did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareFieldLists(a.Fields, b.Fields); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("field lists did not match", a.Fields, b.Fields, &[]*node{child}))
	}

	return newRetVal(retCmp, "struct types did not match", a, b, children)
}

func compareFunctionTypes(a *ast.FuncType, b *ast.FuncType) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "function types did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareFieldLists(a.Params, b.Params); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("parameter lists did not match", a.Params, b.Params, &[]*node{child}))
	}

	if cmp, child := compareFieldLists(a.Results, b.Results); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("result lists did not match", a.Results, b.Results, &[]*node{child}))
	}

	return newRetVal(retCmp, "function types did not match", a, b, children)
}

func compareFieldLists(a *ast.FieldList, b *ast.FieldList) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "field lists did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a.List), len(b.List)); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a.List {
		if i >= len(b.List) {
			break
		}
		if cmp, child := compareFields(a.List[i], b.List[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)
			children = append(children, newNode(fmt.Sprintf("fields at index %d did not match", i), a.List[i], b.List[i], &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "field lists did not match", a, b, children)
}

func compareFields(a *ast.Field, b *ast.Field) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "fields did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareIdentifierLists(a.Names, b.Names); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("name lists did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Type, b.Type); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("types did not match", a.Type, b.Type, &[]*node{child}))
	}

	if cmp, child := compareBasicLiterals(a.Tag, b.Tag); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("tags did not match", a.Tag, b.Tag, &[]*node{child}))
	}

	return newRetVal(retCmp, "fields did not match", a, b, children)
}

func compareInterfaceTypes(a *ast.InterfaceType, b *ast.InterfaceType) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "interface types did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareFieldLists(a.Methods, b.Methods); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("method lists did not match", a.Methods, b.Methods, &[]*node{child}))
	}

	if cmp, child := compareBools(a.Incomplete, b.Incomplete); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("incomplete values did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "interface types did not match", a, b, children)
}

func compareMapTypes(a *ast.MapType, b *ast.MapType) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "map types did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Key, b.Key); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("key expressions did not match", a.Key, b.Key, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Value, b.Value); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("value expressions did not match", a.Value, b.Value, &[]*node{child}))
	}

	return newRetVal(retCmp, "map types did not match", a, b, children)
}

func compareChannelTypes(a *ast.ChanType, b *ast.ChanType) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "channel types did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareChannelDirections(a.Dir, b.Dir); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("directions did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Value, b.Value); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("value expressions did not match", a.Value, b.Value, &[]*node{child}))
	}

	return newRetVal(retCmp, "channel types did not match", a, b, children)
}

func compareChannelDirections(a ast.ChanDir, b ast.ChanDir) (int, *node) {
	if cmp, child := compareInts(int(a), int(b)); cmp != 0 {
		return cmp, newNode("channel directions did not match", nil, nil, &[]*node{child})
	}
	return 0, nil
}

func compareBadStatements(a *ast.BadStmt, b *ast.BadStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "bad statements did not match")
	}

	return 0, nil
}

func compareDeclStatements(a *ast.DeclStmt, b *ast.DeclStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "declaration statements did not match")
	}

	cmp, child := compareDecls(a.Decl, b.Decl)
	return newRetVal(cmp, "declaration statements did not match", a, b, []*node{child})
}

func compareDecls(a ast.Decl, b ast.Decl) (int, *node) {
	if badDeclA, ok := a.(*ast.BadDecl); ok {
		if badDeclB, ok := b.(*ast.BadDecl); ok {
			cmp, child := compareBadDecls(badDeclA, badDeclB)
			return newRetVal(cmp, "declarations did not match", a, b, []*node{child})
		}
		return -1,
			newNode(
				"declarations did not match",
				nil,
				nil,
				&[]*node{
					newNode("declaration types did not match", nil, nil, nil),
				},
			)
	}
	if genDeclA, ok := a.(*ast.GenDecl); ok {
		if genDeclB, ok := b.(*ast.GenDecl); ok {
			cmp, child := compareGenDecls(genDeclA, genDeclB)
			return newRetVal(cmp, "declarations did not match", a, b, []*node{child})
		}
		if _, ok := b.(*ast.BadDecl); ok {
			return 1,
				newNode(
					"declarations did not match",
					nil,
					nil,
					&[]*node{
						newNode("declaration types did not match", nil, nil, nil),
					},
				)
		}
		return -1,
			newNode(
				"declarations did not match",
				nil,
				nil,
				&[]*node{
					newNode("declaration types did not match", nil, nil, nil),
				},
			)
	}
	if funcDeclA, ok := a.(*ast.FuncDecl); ok {
		if funcDeclB, ok := b.(*ast.FuncDecl); ok {
			cmp, child := compareFuncDecls(funcDeclA, funcDeclB)
			return newRetVal(cmp, "declarations did not match", a, b, []*node{child})
		}
		return 1,
			newNode(
				"declarations did not match",
				nil,
				nil,
				&[]*node{
					newNode("declaration types did not match", nil, nil, nil),
				},
			)
	}
	panic(fmt.Sprintf("unrecognized declaration type: %v", a))
}

func compareBadDecls(a *ast.BadDecl, b *ast.BadDecl) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "bad declarations did not match")
	}
	return 0, nil
}

func compareGenDecls(a *ast.GenDecl, b *ast.GenDecl) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "generic declarations did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareTokens(a.Tok, b.Tok); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("tokens did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareSpecLists(a.Specs, b.Specs); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("spec lists did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "generic declarations did not match", a, b, children)
}

func compareFuncDecls(a *ast.FuncDecl, b *ast.FuncDecl) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "function declarations did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareIdentifiers(a.Name, b.Name); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("names did not match", a.Name, b.Name, &[]*node{child}))
	}

	if cmp, child := compareFieldLists(a.Recv, b.Recv); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("receivers did not match", a.Recv, b.Recv, &[]*node{child}))
	}

	if cmp, child := compareFunctionTypes(a.Type, b.Type); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("types did not match", a.Type, b.Type, &[]*node{child}))
	}

	if cmp, child := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bodies did not match", a.Body, b.Body, &[]*node{child}))
	}

	return newRetVal(retCmp, "function declarations did not match", a, b, children)
}

func compareEmptyStatements(a *ast.EmptyStmt, b *ast.EmptyStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "empty statements did not match")
	}

	cmp, child := compareBools(a.Implicit, b.Implicit)
	return newRetVal(cmp, "empty statements did not match", a, b, []*node{child})
}

func compareLabeledStatements(a *ast.LabeledStmt, b *ast.LabeledStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "labeled statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareIdentifiers(a.Label, b.Label); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("labels did not match", a.Label, b.Label, &[]*node{child}))
	}

	if cmp, child := compareStatements(a.Stmt, b.Stmt); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("statements did not match", a.Stmt, b.Stmt, &[]*node{child}))
	}

	return newRetVal(retCmp, "labeled statements did not match", a, b, children)
}

func compareExpressionStatements(a *ast.ExprStmt, b *ast.ExprStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "expression statements did not match")
	}

	cmp, child := compareExpressions(a.X, b.X)
	return newRetVal(cmp, "expression statements did not match", a, b, []*node{child})
}

func compareSendStatements(a *ast.SendStmt, b *ast.SendStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "send statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Chan, b.Chan); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("channels did not match", a.Chan, b.Chan, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Value, b.Value); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("values did not match", a.Value, b.Value, &[]*node{child}))
	}

	return newRetVal(retCmp, "send statements did not match", a, b, children)
}

func compareIncDecStatements(a *ast.IncDecStmt, b *ast.IncDecStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "increment/decrement statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareTokens(a.Tok, b.Tok); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("tokens did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("X expressions did not match", a.X, b.X, &[]*node{child}))
	}

	return newRetVal(retCmp, "increment/decrement statements did not match", a, b, children)
}

func compareAssignStatements(a *ast.AssignStmt, b *ast.AssignStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "assign statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareTokens(a.Tok, b.Tok); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("tokens did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressionLists(a.Lhs, b.Lhs); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("lhs did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressionLists(a.Rhs, b.Rhs); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("rhs did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "assign statements did not match", a, b, children)
}

func compareGoStatements(a *ast.GoStmt, b *ast.GoStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "go statements did not match")
	}

	cmp, child := compareCallExpressions(a.Call, b.Call)
	return newRetVal(cmp, "go statements did not match", a, b, []*node{
		newNode("call expressions did not match", a.Call, b.Call, &[]*node{child}),
	})
}

func compareDeferStatements(a *ast.DeferStmt, b *ast.DeferStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "defer statements did not match")
	}

	cmp, child := compareCallExpressions(a.Call, b.Call)
	return newRetVal(cmp, "defer statements did not match", a, b, []*node{
		newNode("call expressions did not match", a.Call, b.Call, &[]*node{child}),
	})
}

func compareReturnStatements(a *ast.ReturnStmt, b *ast.ReturnStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "return statements did not match")
	}

	cmp, child := compareExpressionLists(a.Results, b.Results)
	return newRetVal(cmp, "return statements did not match", a, b, []*node{
		newNode("results did not match", nil, nil, &[]*node{child}),
	})
}

func compareBranchStatements(a *ast.BranchStmt, b *ast.BranchStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "branch statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareTokens(a.Tok, b.Tok); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("tokens did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareIdentifiers(a.Label, b.Label); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("labels did not match", a.Label, b.Label, &[]*node{child}))
	}

	return newRetVal(retCmp, "branch statements did not match", a, b, children)
}

func compareIfStatements(a *ast.IfStmt, b *ast.IfStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "if statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareStatements(a.Init, b.Init); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("init statements did not match", a.Init, b.Init, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Cond, b.Cond); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("conditions did not match", a.Cond, b.Cond, &[]*node{child}))
	}

	if cmp, child := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bodies did not match", a.Body, b.Body, &[]*node{child}))
	}

	if cmp, child := compareStatements(a.Else, b.Else); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("else statements did not match", a.Else, b.Else, &[]*node{child}))
	}

	return newRetVal(retCmp, "if statements did not match", a, b, children)
}

func compareCaseClauses(a *ast.CaseClause, b *ast.CaseClause) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "case clauses did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressionLists(a.List, b.List); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("lists did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareStatementLists(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bodies did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "case clauses did not match", a, b, children)
}

func compareStatementLists(a []ast.Stmt, b []ast.Stmt) (int, *node) {
	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a), len(b)); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a {
		if i >= len(b) {
			break
		}
		if cmp, child := compareStatements(a[i], b[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)
			children = append(children, newNode(fmt.Sprintf("statements at index %d did not match", i), a[i], b[i], &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "statement lists did not match", nil, nil, children)
}

func compareSwitchStatements(a *ast.SwitchStmt, b *ast.SwitchStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "switch statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareStatements(a.Init, b.Init); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("init statements did not match", a.Init, b.Init, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Tag, b.Tag); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("tags did not match", a.Tag, b.Tag, &[]*node{child}))
	}

	if cmp, child := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bodies did not match", a.Body, b.Body, &[]*node{child}))
	}

	return newRetVal(retCmp, "switch statements did not match", a, b, children)
}

func compareTypeSwitchStatements(a *ast.TypeSwitchStmt, b *ast.TypeSwitchStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "type switch statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareStatements(a.Init, b.Init); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("init statements did not match", a.Init, b.Init, &[]*node{child}))
	}

	if cmp, child := compareStatements(a.Assign, b.Assign); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("assign statements did not match", a.Assign, b.Assign, &[]*node{child}))
	}

	if cmp, child := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bodies did not match", a.Body, b.Body, &[]*node{child}))
	}

	return newRetVal(retCmp, "type switch statements did not match", a, b, children)
}

func compareCommClauses(a *ast.CommClause, b *ast.CommClause) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "comm clauses did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareStatements(a.Comm, b.Comm); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("comm statements did not match", a.Comm, b.Comm, &[]*node{child}))
	}

	if cmp, child := compareStatementLists(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bodies did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "comm clauses did not match", a, b, children)
}

func compareSelectStatements(a *ast.SelectStmt, b *ast.SelectStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "select statements did not match")
	}

	cmp, child := compareBlockStatements(a.Body, b.Body)
	return newRetVal(cmp, "select statements did not match", a, b, []*node{
		newNode("bodies did not match", a.Body, b.Body, &[]*node{child}),
	})
}

func compareForStatements(a *ast.ForStmt, b *ast.ForStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "for statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareStatements(a.Init, b.Init); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("init statements did not match", a.Init, b.Init, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Cond, b.Cond); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("conditions did not match", a.Cond, b.Cond, &[]*node{child}))
	}

	if cmp, child := compareStatements(a.Post, b.Post); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("post statements did not match", a.Post, b.Post, &[]*node{child}))
	}

	if cmp, child := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bodies did not match", a.Body, b.Body, &[]*node{child}))
	}

	return newRetVal(retCmp, "for statements did not match", a, b, children)
}

func compareRangeStatements(a *ast.RangeStmt, b *ast.RangeStmt) (int, *node) {
	if a == nil || b == nil {
		return newNilRetVal(a, b, "range statements did not match")
	}

	retCmp := 0
	var children []*node

	if cmp, child := compareExpressions(a.Key, b.Key); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("key statements did not match", a.Key, b.Key, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.Value, b.Value); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("value statements did not match", a.Value, b.Value, &[]*node{child}))
	}

	if cmp, child := compareTokens(a.Tok, b.Tok); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("tokens did not match", nil, nil, &[]*node{child}))
	}

	if cmp, child := compareExpressions(a.X, b.X); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("X expressions did not match", a.X, b.X, &[]*node{child}))
	}

	if cmp, child := compareBlockStatements(a.Body, b.Body); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bodies did not match", a.Body, b.Body, &[]*node{child}))
	}

	return newRetVal(retCmp, "range statements did not match", a, b, children)
}

func compareGenDeclLists(a []*ast.GenDecl, b []*ast.GenDecl) (int, *node) {
	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a), len(b)); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a {
		if i >= len(b) {
			break
		}
		if cmp, child := compareGenDecls(a[i], b[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)
			children = append(children, newNode(fmt.Sprintf("generic declarations at index %d did not match", i), a[i], b[i], &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "generic declaration lists did not match", nil, nil, children)
}

func compareFuncDeclLists(a []*ast.FuncDecl, b []*ast.FuncDecl) (int, *node) {
	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a), len(b)); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a {
		if i >= len(b) {
			break
		}
		if cmp, child := compareFuncDecls(a[i], b[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)
			children = append(children, newNode(fmt.Sprintf("function declarations at index %d did not match", i), a[i], b[i], &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "function declaration lists did not match", nil, nil, children)
}

func compareImportSpecLists(a []*ast.ImportSpec, b []*ast.ImportSpec) (int, *node) {
	retCmp := 0
	var children []*node

	if cmp, child := compareInts(len(a), len(b)); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("length of lists did not match", nil, nil, &[]*node{child}))
	}

	for i := range a {
		if i >= len(b) {
			break
		}
		if cmp, child := compareImportSpecs(a[i], b[i]); cmp != 0 {
			setIfUnset(&retCmp, cmp)
			children = append(children, newNode(fmt.Sprintf("import specs at index %d did not match", i), a[i], b[i], &[]*node{child}))
		}
	}

	return newRetVal(retCmp, "import spec lists did not match", nil, nil, children)
}

func compareDeclLists(a []ast.Decl, b []ast.Decl) (int, *node) {
	retCmp := 0
	var children []*node

	badDeclsA, genDeclsA, funcDeclsA := splitDecls(a)
	badDeclsB, genDeclsB, funcDeclsB := splitDecls(b)

	if cmp, child := compareBadDeclLists(badDeclsA, badDeclsB); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("bad declaration lists did not match", nil, nil, &[]*node{child}))
	}

	sortGenDeclList(&genDeclsA)
	sortGenDeclList(&genDeclsB)
	if cmp, child := compareGenDeclLists(genDeclsA, genDeclsB); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("generic declaration lists did not match", nil, nil, &[]*node{child}))
	}

	sortFuncDeclList(&funcDeclsA)
	sortFuncDeclList(&funcDeclsB)
	if cmp, child := compareFuncDeclLists(funcDeclsA, funcDeclsB); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("function declaration lists did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "declaration lists did not match", nil, nil, children)
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

func compareBadDeclLists(a []*ast.BadDecl, b []*ast.BadDecl) (int, *node) {
	if len(a) > 0 {
		panic(fmt.Errorf("first source contained bad declaration: %v", a[0]))
	}
	if len(b) > 0 {
		panic(fmt.Errorf("second source contained bad declaration: %v", b[0]))
	}

	return 0, nil
}

func compareFiles(a *ast.File, b *ast.File) (int, *node) {
	retCmp := 0
	var children []*node

	if cmp, child := compareDeclLists(a.Decls, b.Decls); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("declaration lists did not match", nil, nil, &[]*node{child}))
	}

	sortImportList(&a.Imports)
	sortImportList(&b.Imports)
	if cmp, child := compareImportSpecLists(a.Imports, b.Imports); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("imports did not match", nil, nil, &[]*node{child}))
	}

	sortIdentifierList(&a.Unresolved, true)
	sortIdentifierList(&b.Unresolved, true)
	if cmp, child := compareIdentifierLists(a.Unresolved, b.Unresolved); cmp != 0 {
		setIfUnset(&retCmp, cmp)
		children = append(children, newNode("unresolved identifiers did not match", nil, nil, &[]*node{child}))
	}

	return newRetVal(retCmp, "files did not match", a, b, children)
}
