package eqgo

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func newTestNode(msg string, child *node) *node {
	var children []*node
	if child != nil {
		children = append(children, child)
	}
	return &node{
		msg:      msg,
		children: children,
	}
}

func TestCompareInts(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		want     int
		wantNode *node
	}{
		{
			a:        1,
			b:        2,
			want:     -1,
			wantNode: newTestNode("ints did not match: 1 < 2", nil),
		},
		{
			a:        2,
			b:        1,
			want:     1,
			wantNode: newTestNode("ints did not match: 2 > 1", nil),
		},
		{
			a:    1,
			b:    1,
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareInts(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareInts(%d, %d) == (%d, %v), want (%d, %v)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareBools(t *testing.T) {
	testCases := []struct {
		a        bool
		b        bool
		want     int
		wantNode *node
	}{
		{
			a:        false,
			b:        true,
			want:     -1,
			wantNode: newTestNode("bools did not match: false < true", nil),
		},
		{
			a:        true,
			b:        false,
			want:     1,
			wantNode: newTestNode("bools did not match: true > false", nil),
		},
		{
			a:    false,
			b:    false,
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareBools(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareBools(%t, %t) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareStrings(t *testing.T) {
	testCases := []struct {
		a        string
		b        string
		want     int
		wantNode *node
	}{
		{
			a:        "a",
			b:        "b",
			want:     -1,
			wantNode: newTestNode("strings did not match: a < b", nil),
		},
		{
			a:        "b",
			b:        "a",
			want:     1,
			wantNode: newTestNode("strings did not match: b > a", nil),
		},
		{
			a:    "a",
			b:    "a",
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareStrings(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareStrings(%s, %s) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareIdentifiers(t *testing.T) {
	testCases := []struct {
		a        *ast.Ident
		b        *ast.Ident
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:    nil,
			b:    &ast.Ident{},
			want: 1,
			wantNode: newTestNode(
				"identifiers did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.Ident{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"identifiers did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a:    &ast.Ident{Name: "a"},
			b:    &ast.Ident{Name: "b"},
			want: -1,
			wantNode: newTestNode(
				"identifiers did not match",
				newTestNode("strings did not match: a < b", nil),
			),
		},
		{
			a:    &ast.Ident{Name: "b"},
			b:    &ast.Ident{Name: "a"},
			want: 1,
			wantNode: newTestNode(
				"identifiers did not match",
				newTestNode("strings did not match: b > a", nil),
			),
		},
		{
			a:    &ast.Ident{Name: "a"},
			b:    &ast.Ident{Name: "a"},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareIdentifiers(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareIdentifiers(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareImportSpecs(t *testing.T) {
	testCases := []struct {
		a        *ast.ImportSpec
		b        *ast.ImportSpec
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:    nil,
			b:    &ast.ImportSpec{},
			want: 1,
			wantNode: newTestNode(
				"import specs did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.ImportSpec{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"import specs did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
			},
			b: &ast.ImportSpec{
				Name: &ast.Ident{Name: "b"},
			},
			want: -1,
			wantNode: newTestNode(
				"import specs did not match",
				newTestNode(
					"names did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: a < b", nil),
					),
				),
			),
		},
		{
			a: &ast.ImportSpec{
				Name: &ast.Ident{Name: "b"},
			},
			b: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
			},
			want: 1,
			wantNode: newTestNode(
				"import specs did not match",
				newTestNode(
					"names did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: b > a", nil),
					),
				),
			),
		},
		{
			a: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
				Path: &ast.BasicLit{Kind: token.INT},
			},
			b: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
				Path: &ast.BasicLit{Kind: token.FLOAT},
			},
			want: -1,
			wantNode: newTestNode(
				"import specs did not match",

				newTestNode(
					"paths did not match",

					newTestNode(
						"basic literals did not match",

						newTestNode(
							"kinds did not match",

							newTestNode(
								"tokens did not match",

								newTestNode("ints did not match: 5 < 6", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
				Path: &ast.BasicLit{Kind: token.FLOAT},
			},
			b: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
				Path: &ast.BasicLit{Kind: token.INT},
			},
			want: 1,
			wantNode: newTestNode(
				"import specs did not match",

				newTestNode(
					"paths did not match",

					newTestNode(
						"basic literals did not match",

						newTestNode(
							"kinds did not match",

							newTestNode(
								"tokens did not match",

								newTestNode("ints did not match: 6 > 5", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
				Path: &ast.BasicLit{Kind: token.FLOAT},
			},
			b: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
				Path: &ast.BasicLit{Kind: token.FLOAT},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareImportSpecs(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareImportSpecs(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareBasicLiterals(t *testing.T) {
	testCases := []struct {
		a        *ast.BasicLit
		b        *ast.BasicLit
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:    nil,
			b:    &ast.BasicLit{},
			want: 1,
			wantNode: newTestNode(
				"basic literals did not match",

				newTestNode(
					"nil comparisons did not match",

					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.BasicLit{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"basic literals did not match",

				newTestNode(
					"nil comparisons did not match",

					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a:    &ast.BasicLit{Kind: token.INT},
			b:    &ast.BasicLit{Kind: token.FLOAT},
			want: -1,
			wantNode: newTestNode(
				"basic literals did not match",

				newTestNode(
					"kinds did not match",

					newTestNode(
						"tokens did not match",

						newTestNode("ints did not match: 5 < 6", nil),
					),
				),
			),
		},
		{
			a:    &ast.BasicLit{Kind: token.FLOAT},
			b:    &ast.BasicLit{Kind: token.INT},
			want: 1,
			wantNode: newTestNode(
				"basic literals did not match",

				newTestNode(
					"kinds did not match",

					newTestNode(
						"tokens did not match",

						newTestNode("ints did not match: 6 > 5", nil),
					),
				),
			),
		},
		{
			a:    &ast.BasicLit{Value: "a"},
			b:    &ast.BasicLit{Value: "b"},
			want: -1,
			wantNode: newTestNode(
				"basic literals did not match",

				newTestNode(
					"values did not match",

					newTestNode("strings did not match: a < b", nil),
				),
			),
		},
		{
			a:    &ast.BasicLit{Value: "b"},
			b:    &ast.BasicLit{Value: "a"},
			want: 1,
			wantNode: newTestNode(
				"basic literals did not match",

				newTestNode(
					"values did not match",

					newTestNode("strings did not match: b > a", nil),
				),
			),
		},
		{
			a: &ast.BasicLit{
				Kind:  token.INT,
				Value: "a",
			},
			b: &ast.BasicLit{
				Kind:  token.INT,
				Value: "a",
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareBasicLiterals(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareBasicLiterals(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareValueSpecs(t *testing.T) {
	testCases := []struct {
		a        *ast.ValueSpec
		b        *ast.ValueSpec
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:    nil,
			b:    &ast.ValueSpec{},
			want: 1,
			wantNode: newTestNode(
				"value specs did not match",

				newTestNode(
					"nil comparisons did not match",

					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.ValueSpec{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"value specs did not match",

				newTestNode(
					"nil comparisons did not match",

					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "b"}},
			},
			want: -1,
			wantNode: newTestNode(
				"value specs did not match",

				newTestNode(
					"name lists did not match",

					newTestNode(
						"identifier lists did not match",

						newTestNode(
							"identifiers at index 0 did not match",

							newTestNode(
								"identifiers did not match",

								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "b"}},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
			},
			want: 1,
			wantNode: newTestNode(
				"value specs did not match",

				newTestNode(
					"name lists did not match",

					newTestNode(
						"identifier lists did not match",

						newTestNode(
							"identifiers at index 0 did not match",

							newTestNode(
								"identifiers did not match",

								newTestNode("strings did not match: b > a", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
			},
			want: 0,
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Type:  &ast.Ident{Name: "aType"},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Type:  &ast.Ident{Name: "bType"},
			},
			want: -1,
			wantNode: newTestNode(
				"value specs did not match",

				newTestNode(
					"types did not match",

					newTestNode(
						"expressions did not match",

						newTestNode(
							"identifiers did not match",

							newTestNode("strings did not match: aType < bType", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Type:  &ast.Ident{Name: "bType"},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Type:  &ast.Ident{Name: "aType"},
			},
			want: 1,
			wantNode: newTestNode(
				"value specs did not match",

				newTestNode(
					"types did not match",

					newTestNode(
						"expressions did not match",

						newTestNode(
							"identifiers did not match",

							newTestNode("strings did not match: bType > aType", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Type:  &ast.Ident{Name: "aType"},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Type:  &ast.Ident{Name: "aType"},
			},
			want: 0,
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "5",
					},
				},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "6",
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"value specs did not match",

				newTestNode(
					"values did not match",

					newTestNode(
						"expression lists did not match",

						newTestNode(
							"expressions at index 0 did not match",

							newTestNode(
								"expressions did not match",

								newTestNode(
									"basic literals did not match",

									newTestNode(
										"values did not match",

										newTestNode("strings did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "6",
					},
				},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "5",
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"value specs did not match",

				newTestNode(
					"values did not match",

					newTestNode(
						"expression lists did not match",

						newTestNode(
							"expressions at index 0 did not match",

							newTestNode(
								"expressions did not match",

								newTestNode(
									"basic literals did not match",

									newTestNode(
										"values did not match",

										newTestNode("strings did not match: 6 > 5", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "5",
					},
				},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "5",
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareValueSpecs(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareValueSpecs(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareIdentifierLists(t *testing.T) {
	testCases := []struct {
		a        []*ast.Ident
		b        []*ast.Ident
		want     int
		wantNode *node
	}{
		{
			a: []*ast.Ident{
				{},
			},
			b: []*ast.Ident{
				{},
				{},
			},
			want: -1,
			wantNode: newTestNode(
				"identifier lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode(
						"ints did not match: 1 < 2",
						nil,
					),
				),
			),
		},
		{
			a: []*ast.Ident{
				{},
				{},
			},
			b: []*ast.Ident{
				{},
			},
			want: 1,
			wantNode: newTestNode(
				"identifier lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode(
						"ints did not match: 2 > 1",
						nil,
					),
				),
			),
		},
		{
			a: []*ast.Ident{
				{},
				{},
			},
			b: []*ast.Ident{
				{},
				{},
			},
			want: 0,
		},
		{
			a: []*ast.Ident{
				{Name: "a"},
			},
			b: []*ast.Ident{
				{Name: "b"},
			},
			want: -1,
			wantNode: newTestNode(
				"identifier lists did not match",
				newTestNode(
					"identifiers at index 0 did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: a < b", nil),
					),
				),
			),
		},
		{
			a: []*ast.Ident{
				{Name: "b"},
			},
			b: []*ast.Ident{
				{Name: "a"},
			},
			want: 1,
			wantNode: newTestNode(
				"identifier lists did not match",
				newTestNode(
					"identifiers at index 0 did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: b > a", nil),
					),
				),
			),
		},
		{
			a: []*ast.Ident{
				{Name: "a"},
			},
			b: []*ast.Ident{
				{Name: "a"},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareIdentifierLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareIdentifierLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareExpressionLists(t *testing.T) {
	testCases := []struct {
		a        []ast.Expr
		b        []ast.Expr
		want     int
		wantNode *node
	}{
		{
			a: []ast.Expr{
				ast.NewIdent(""),
			},
			b: []ast.Expr{
				ast.NewIdent(""),
				ast.NewIdent(""),
			},
			want: -1,
			wantNode: newTestNode(
				"expression lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 1 < 2", nil),
				),
			),
		},
		{
			a: []ast.Expr{
				ast.NewIdent(""),
				ast.NewIdent(""),
			},
			b: []ast.Expr{
				ast.NewIdent(""),
			},
			want: 1,
			wantNode: newTestNode(
				"expression lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 2 > 1", nil),
				),
			),
		},
		{
			a: []ast.Expr{
				ast.NewIdent("a"),
				ast.NewIdent("b"),
			},
			b: []ast.Expr{
				ast.NewIdent("a"),
				ast.NewIdent("c"),
			},
			want: -1,
			wantNode: newTestNode(
				"expression lists did not match",
				newTestNode(
					"expressions at index 1 did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b < c", nil),
						),
					),
				),
			),
		},
		{
			a: []ast.Expr{
				ast.NewIdent("a"),
				ast.NewIdent("c"),
			},
			b: []ast.Expr{
				ast.NewIdent("a"),
				ast.NewIdent("b"),
			},
			want: 1,
			wantNode: newTestNode(
				"expression lists did not match",
				newTestNode(
					"expressions at index 1 did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: c > b", nil),
						),
					),
				),
			),
		},
		{
			a: []ast.Expr{
				ast.NewIdent("a"),
				ast.NewIdent("b"),
			},
			b: []ast.Expr{
				ast.NewIdent("a"),
				ast.NewIdent("b"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareExpressionLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareExpressionLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareExpressions(t *testing.T) {
	testCases := []struct {
		a        ast.Expr
		b        ast.Expr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    ast.NewIdent(""),
			want: 1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode(
						"bools did not match: true > false",
						nil,
					),
				),
			),
		},
		{
			a:    ast.NewIdent(""),
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode(
						"bools did not match: false < true",
						nil,
					),
				),
			),
		},
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:    ast.NewIdent(""),
			b:    &ast.Ellipsis{},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"sort indices did not match",
					newTestNode(
						"ints did not match: 1 < 2",
						nil,
					),
				),
			),
		},
		{
			a:    &ast.Ellipsis{},
			b:    ast.NewIdent(""),
			want: 1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"sort indices did not match",
					newTestNode(
						"ints did not match: 2 > 1",
						nil,
					),
				),
			),
		},
		{
			a:    ast.NewIdent(""),
			b:    ast.NewIdent(""),
			want: 0,
		},
		{
			a:    ast.NewIdent("a"),
			b:    ast.NewIdent("b"),
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"identifiers did not match",
					newTestNode(
						"strings did not match: a < b",
						nil,
					),
				),
			),
		},
		{
			a:    &ast.Ellipsis{},
			b:    &ast.Ellipsis{},
			want: 0,
		},
		{
			a: &ast.Ellipsis{
				Elt: ast.NewIdent("a"),
			},
			b: &ast.Ellipsis{
				Elt: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"ellipses did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.BasicLit{},
			b:    &ast.BasicLit{},
			want: 0,
		},
		{
			a: &ast.BasicLit{
				Kind: token.INT,
			},
			b: &ast.BasicLit{
				Kind: token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"basic literals did not match",
					newTestNode(
						"kinds did not match",
						newTestNode(
							"tokens did not match",
							newTestNode(
								"ints did not match: 5 < 6",
								nil,
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.FuncLit{},
			b:    &ast.FuncLit{},
			want: 0,
		},
		{
			a: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},
			},
			b: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("b"),
								},
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"function literals did not match",
					newTestNode(
						"types did not match",
						newTestNode(
							"function types did not match",
							newTestNode(
								"parameter lists did not match",
								newTestNode(
									"field lists did not match",
									newTestNode(
										"fields at index 0 did not match",
										newTestNode(
											"fields did not match",
											newTestNode(
												"name lists did not match",
												newTestNode(
													"identifier lists did not match",
													newTestNode(
														"identifiers at index 0 did not match",
														newTestNode(
															"identifiers did not match",
															newTestNode("strings did not match: a < b", nil),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.CompositeLit{},
			b:    &ast.CompositeLit{},
			want: 0,
		},
		{
			a: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
			},
			b: &ast.CompositeLit{
				Type: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"composite literals did not match",
					newTestNode(
						"types did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.ParenExpr{},
			b:    &ast.ParenExpr{},
			want: 0,
		},
		{
			a: &ast.ParenExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.ParenExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"parentheses did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.SelectorExpr{},
			b:    &ast.SelectorExpr{},
			want: 0,
		},
		{
			a: &ast.SelectorExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.SelectorExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"selector expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.IndexExpr{},
			b:    &ast.IndexExpr{},
			want: 0,
		},
		{
			a: &ast.IndexExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.IndexExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"index expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.SliceExpr{},
			b:    &ast.SliceExpr{},
			want: 0,
		},
		{
			a: &ast.SliceExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.SliceExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"slices did not match",
					newTestNode(
						"X expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.TypeAssertExpr{},
			b:    &ast.TypeAssertExpr{},
			want: 0,
		},
		{
			a: &ast.TypeAssertExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.TypeAssertExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"type assertions did not match",
					newTestNode(
						"X expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.CallExpr{},
			b:    &ast.CallExpr{},
			want: 0,
		},
		{
			a: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
			},
			b: &ast.CallExpr{
				Fun: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"call expressions did not match",
					newTestNode(
						"functions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.StarExpr{},
			b:    &ast.StarExpr{},
			want: 0,
		},
		{
			a: &ast.StarExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.StarExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"star expressions did not match",
					newTestNode(
						"X expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.UnaryExpr{},
			b:    &ast.UnaryExpr{},
			want: 0,
		},
		{
			a: &ast.UnaryExpr{
				Op: token.INT,
			},
			b: &ast.UnaryExpr{
				Op: token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"unary expressions did not match",
					newTestNode(
						"operators did not match",
						newTestNode(
							"tokens did not match",
							newTestNode("ints did not match: 5 < 6", nil),
						),
					),
				),
			),
		},
		{
			a:    &ast.BinaryExpr{},
			b:    &ast.BinaryExpr{},
			want: 0,
		},
		{
			a: &ast.BinaryExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.BinaryExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"binary expressions did not match",
					newTestNode(
						"X expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.KeyValueExpr{},
			b:    &ast.KeyValueExpr{},
			want: 0,
		},
		{
			a: &ast.KeyValueExpr{
				Key: ast.NewIdent("a"),
			},
			b: &ast.KeyValueExpr{
				Key: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"key-value expressions did not match",
					newTestNode(
						"keys did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.ArrayType{},
			b:    &ast.ArrayType{},
			want: 0,
		},
		{
			a: &ast.ArrayType{
				Len: ast.NewIdent("a"),
			},
			b: &ast.ArrayType{
				Len: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"array types did not match",
					newTestNode(
						"length expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.StructType{},
			b:    &ast.StructType{},
			want: 0,
		},
		{
			a: &ast.StructType{
				Incomplete: false,
			},
			b: &ast.StructType{
				Incomplete: true,
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"struct types did not match",
					newTestNode(
						"incomplete values did not match",
						newTestNode("bools did not match: false < true", nil),
					),
				),
			),
		},
		{
			a:    &ast.FuncType{},
			b:    &ast.FuncType{},
			want: 0,
		},
		{
			a: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			b: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("b"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"function types did not match",
					newTestNode(
						"parameter lists did not match",
						newTestNode(
							"field lists did not match",
							newTestNode(
								"fields at index 0 did not match",
								newTestNode(
									"fields did not match",
									newTestNode(
										"name lists did not match",
										newTestNode(
											"identifier lists did not match",
											newTestNode(
												"identifiers at index 0 did not match",
												newTestNode(
													"identifiers did not match",
													newTestNode("strings did not match: a < b", nil),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.InterfaceType{},
			b:    &ast.InterfaceType{},
			want: 0,
		},
		{
			a: &ast.InterfaceType{
				Incomplete: false,
			},
			b: &ast.InterfaceType{
				Incomplete: true,
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"interface types did not match",
					newTestNode(
						"incomplete values did not match",
						newTestNode("bools did not match: false < true", nil),
					),
				),
			),
		},
		{
			a:    &ast.MapType{},
			b:    &ast.MapType{},
			want: 0,
		},
		{
			a: &ast.MapType{
				Key: ast.NewIdent("a"),
			},
			b: &ast.MapType{
				Key: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"map types did not match",
					newTestNode(
						"key expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.ChanType{},
			b:    &ast.ChanType{},
			want: 0,
		},
		{
			a: &ast.ChanType{
				Dir: ast.SEND,
			},
			b: &ast.ChanType{
				Dir: ast.RECV,
			},
			want: -1,
			wantNode: newTestNode(
				"expressions did not match",
				newTestNode(
					"channel types did not match",
					newTestNode(
						"directions did not match",
						newTestNode(
							"channel directions did not match",
							newTestNode("ints did not match: 1 < 2", nil),
						),
					),
				),
			),
		},
	}
	for _, c := range testCases {
		got, gotNode := compareExpressions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareTypeSpecs(t *testing.T) {
	testCases := []struct {
		a        *ast.TypeSpec
		b        *ast.TypeSpec
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.TypeSpec{},
			want: 1,
			wantNode: newTestNode(
				"type specs did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.TypeSpec{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"type specs did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a:    &ast.TypeSpec{},
			b:    &ast.TypeSpec{},
			want: 0,
		},
		{
			a: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
			},
			b: &ast.TypeSpec{
				Name: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"type specs did not match",
				newTestNode(
					"names did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: a < b", nil),
					),
				),
			),
		},
		{
			a: &ast.TypeSpec{
				Name: ast.NewIdent("b"),
			},
			b: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"type specs did not match",
				newTestNode(
					"names did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: b > a", nil),
					),
				),
			),
		},
		{
			a: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
			},
			b: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
			},
			want: 0,
		},
		{
			a: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
				Type: ast.NewIdent("x"),
			},
			b: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
				Type: ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"type specs did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
				Type: ast.NewIdent("y"),
			},
			b: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
				Type: ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"type specs did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
				Type: ast.NewIdent("x"),
			},
			b: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
				Type: ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareTypeSpecs(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareTypeSpecs(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareSpecs(t *testing.T) {
	testCases := []struct {
		a        ast.Spec
		b        ast.Spec
		want     int
		wantNode *node
	}{
		{
			a:        &ast.ImportSpec{},
			b:        &ast.ValueSpec{},
			want:     -1,
			wantNode: newTestNode("spec types did not match", nil),
		},
		{
			a:        &ast.ImportSpec{},
			b:        &ast.TypeSpec{},
			want:     -1,
			wantNode: newTestNode("spec types did not match", nil),
		},
		{
			a:        &ast.ValueSpec{},
			b:        &ast.ImportSpec{},
			want:     1,
			wantNode: newTestNode("spec types did not match", nil),
		},
		{
			a:        &ast.TypeSpec{},
			b:        &ast.ImportSpec{},
			want:     1,
			wantNode: newTestNode("spec types did not match", nil),
		},
		{
			a:        &ast.ValueSpec{},
			b:        &ast.TypeSpec{},
			want:     -1,
			wantNode: newTestNode("spec types did not match", nil),
		},
		{
			a:        &ast.TypeSpec{},
			b:        &ast.ValueSpec{},
			want:     1,
			wantNode: newTestNode("spec types did not match", nil),
		},
		{
			a:    &ast.ImportSpec{},
			b:    &ast.ImportSpec{},
			want: 0,
		},
		{
			a: &ast.ImportSpec{
				Name: ast.NewIdent("a"),
			},
			b: &ast.ImportSpec{
				Name: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"specs did not match",
				newTestNode(
					"import specs did not match",
					newTestNode(
						"names did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a:    &ast.ValueSpec{},
			b:    &ast.ValueSpec{},
			want: 0,
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{
					ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"specs did not match",
				newTestNode(
					"value specs did not match",
					newTestNode(
						"name lists did not match",
						newTestNode(
							"identifier lists did not match",
							newTestNode(
								"identifiers at index 0 did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.TypeSpec{},
			b:    &ast.TypeSpec{},
			want: 0,
		},
		{
			a: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
			},
			b: &ast.TypeSpec{
				Name: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"specs did not match",
				newTestNode(
					"type specs did not match",
					newTestNode(
						"names did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
	}
	for _, c := range testCases {
		got, gotNode := compareSpecs(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareSpecs(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareSpecLists(t *testing.T) {
	testCases := []struct {
		a        []ast.Spec
		b        []ast.Spec
		want     int
		wantNode *node
	}{
		{
			a: []ast.Spec{},
			b: []ast.Spec{
				&ast.ImportSpec{},
			},
			want: -1,
			wantNode: newTestNode(
				"spec lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 0 < 1", nil),
				),
			),
		},
		{
			a: []ast.Spec{
				&ast.ImportSpec{},
			},
			b:    []ast.Spec{},
			want: 1,
			wantNode: newTestNode(
				"spec lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 1 > 0", nil),
				),
			),
		},
		{
			a: []ast.Spec{
				&ast.ImportSpec{},
			},
			b: []ast.Spec{
				&ast.ValueSpec{},
			},
			want: -1,
			wantNode: newTestNode(
				"spec lists did not match",
				newTestNode(
					"specs at index 0 did not match",
					newTestNode("spec types did not match", nil),
				),
			),
		},
		{
			a: []ast.Spec{
				&ast.ImportSpec{},
			},
			b: []ast.Spec{
				&ast.ImportSpec{},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareSpecLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareSpecLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareEllipses(t *testing.T) {
	testCases := []struct {
		a        *ast.Ellipsis
		b        *ast.Ellipsis
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.Ellipsis{},
			want: 1,
			wantNode: newTestNode(
				"ellipses did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.Ellipsis{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"ellipses did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.Ellipsis{
				Elt: ast.NewIdent("a"),
			},
			b: &ast.Ellipsis{
				Elt: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"ellipses did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.Ellipsis{
				Elt: ast.NewIdent("b"),
			},
			b: &ast.Ellipsis{
				Elt: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"ellipses did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.Ellipsis{
				Elt: ast.NewIdent("a"),
			},
			b: &ast.Ellipsis{
				Elt: ast.NewIdent("a"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareEllipses(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareEllipses(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareFunctionLiterals(t *testing.T) {
	testCases := []struct {
		a        *ast.FuncLit
		b        *ast.FuncLit
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.FuncLit{},
			want: 1,
			wantNode: newTestNode(
				"function literals did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.FuncLit{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"function literals did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},
			},
			b: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("b"),
								},
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"function literals did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"function types did not match",
						newTestNode(
							"parameter lists did not match",
							newTestNode(
								"field lists did not match",
								newTestNode(
									"fields at index 0 did not match",
									newTestNode(
										"fields did not match",
										newTestNode(
											"name lists did not match",
											newTestNode(
												"identifier lists did not match",
												newTestNode(
													"identifiers at index 0 did not match",
													newTestNode(
														"identifiers did not match",
														newTestNode("strings did not match: a < b", nil),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("b"),
								},
							},
						},
					},
				},
			},
			b: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"function literals did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"function types did not match",
						newTestNode(
							"parameter lists did not match",
							newTestNode(
								"field lists did not match",
								newTestNode(
									"fields at index 0 did not match",
									newTestNode(
										"fields did not match",
										newTestNode(
											"name lists did not match",
											newTestNode(
												"identifier lists did not match",
												newTestNode(
													"identifiers at index 0 did not match",
													newTestNode(
														"identifiers did not match",
														newTestNode("strings did not match: b > a", nil),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},

				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("x"),
							},
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},

				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("x"),
							},
							Tok: token.FLOAT,
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"function literals did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},

				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("x"),
							},
							Tok: token.FLOAT,
						},
					},
				},
			},
			b: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},

				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("x"),
							},
							Tok: token.INT,
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"function literals did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},

				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("a"),
								},
							},
						},
					},
				},

				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareFunctionLiterals(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareFunctionLiterals(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareBlockStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.BlockStmt
		b        *ast.BlockStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.BlockStmt{},
			want: 1,
			wantNode: newTestNode(
				"block statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.BlockStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"block statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
					&ast.BranchStmt{},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"block statements did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 1 < 2", nil),
				),
			),
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
					&ast.BranchStmt{},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"block statements did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 2 > 1", nil),
				),
			),
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.FLOAT,
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"block statements did not match",
				newTestNode(
					"statements at index 1 did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.FLOAT,
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"block statements did not match",
				newTestNode(
					"statements at index 1 did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareBlockStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareBlockStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareStatements(t *testing.T) {
	testCases := []struct {
		a        ast.Stmt
		b        ast.Stmt
		want     int
		wantNode *node
	}{
		{
			a: nil,
			b: &ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("i"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("i"),
				},
			},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.FuncDecl{},
			},
			b:    &ast.EmptyStmt{},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"statement types did not match",
					newTestNode("ints did not match: 1 < 2", nil),
				),
			),
		},
		{
			a:    &ast.EmptyStmt{},
			b:    &ast.DeclStmt{},
			want: 1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"statement types did not match",
					newTestNode("ints did not match: 2 > 1", nil),
				),
			),
		},
		{
			a:    &ast.BadStmt{},
			b:    &ast.BadStmt{},
			want: 0,
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.BadDecl{},
			},
			b: &ast.DeclStmt{
				Decl: &ast.GenDecl{},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"declaration statements did not match",
					newTestNode(
						"declarations did not match",
						newTestNode("declaration types did not match", nil),
					),
				),
			),
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.GenDecl{},
			},
			b: &ast.DeclStmt{
				Decl: &ast.GenDecl{},
			},
			want: 0,
		},
		{
			a: &ast.EmptyStmt{
				Implicit: false,
			},
			b: &ast.EmptyStmt{
				Implicit: true,
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"empty statements did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a:    &ast.EmptyStmt{},
			b:    &ast.EmptyStmt{},
			want: 0,
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"labeled statements did not match",
					newTestNode(
						"labels did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a:    &ast.LabeledStmt{},
			b:    &ast.LabeledStmt{},
			want: 0,
		},
		{
			a: &ast.ExprStmt{
				X: ast.NewIdent("a"),
			},
			b: &ast.ExprStmt{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"expression statements did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a:    &ast.ExprStmt{},
			b:    &ast.ExprStmt{},
			want: 0,
		},
		{
			a: &ast.SendStmt{
				Chan:  &ast.ChanType{},
				Value: ast.NewIdent("a"),
			},
			b: &ast.SendStmt{
				Chan:  &ast.ChanType{},
				Value: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"send statements did not match",
					newTestNode(
						"values did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.SendStmt{},
			b:    &ast.SendStmt{},
			want: 0,
		},
		{
			a: &ast.IncDecStmt{
				X: ast.NewIdent("a"),
			},
			b: &ast.IncDecStmt{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"increment/decrement statements did not match",
					newTestNode(
						"X expressions did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.IncDecStmt{},
			b:    &ast.IncDecStmt{},
			want: 0,
		},
		{
			a: &ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
				Tok: token.INT,
			},
			b: &ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
				Tok: token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"assign statements did not match",
					newTestNode(
						"tokens did not match",
						newTestNode(
							"tokens did not match",
							newTestNode("ints did not match: 5 < 6", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("i"),
				},
			},
			b: &ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("i"),
				},
			},
			want: 0,
		},
		{
			a: &ast.GoStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			b: &ast.GoStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"go statements did not match",
					newTestNode(
						"call expressions did not match",
						newTestNode(
							"call expressions did not match",
							newTestNode(
								"functions did not match",
								newTestNode(
									"expressions did not match",
									newTestNode(
										"identifiers did not match",
										newTestNode("strings did not match: a < b", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.GoStmt{},
			b:    &ast.GoStmt{},
			want: 0,
		},
		{
			a: &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			b: &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"defer statements did not match",
					newTestNode(
						"call expressions did not match",
						newTestNode(
							"call expressions did not match",
							newTestNode(
								"functions did not match",
								newTestNode(
									"expressions did not match",
									newTestNode(
										"identifiers did not match",
										newTestNode("strings did not match: a < b", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.DeferStmt{},
			b:    &ast.DeferStmt{},
			want: 0,
		},
		{
			a: &ast.ReturnStmt{
				Results: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			b: &ast.ReturnStmt{
				Results: []ast.Expr{
					ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"return statements did not match",
					newTestNode(
						"results did not match",
						newTestNode(
							"expression lists did not match",
							newTestNode(
								"expressions at index 0 did not match",
								newTestNode(
									"expressions did not match",
									newTestNode(
										"identifiers did not match",
										newTestNode("strings did not match: a < b", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.ReturnStmt{},
			b:    &ast.ReturnStmt{},
			want: 0,
		},
		{
			a: &ast.BranchStmt{
				Tok: token.INT,
			},
			b: &ast.BranchStmt{
				Tok: token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"branch statements did not match",
					newTestNode(
						"tokens did not match",
						newTestNode(
							"tokens did not match",
							newTestNode("ints did not match: 5 < 6", nil),
						),
					),
				),
			),
		},
		{
			a:    &ast.BranchStmt{},
			b:    &ast.BranchStmt{},
			want: 0,
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.INT,
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("a"),
						},
						Tok: token.FLOAT,
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"block statements did not match",
					newTestNode(
						"statements at index 0 did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.BlockStmt{},
			b:    &ast.BlockStmt{},
			want: 0,
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.INT,
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.FLOAT,
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"if statements did not match",
					newTestNode(
						"init statements did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.IfStmt{},
			b:    &ast.IfStmt{},
			want: 0,
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			b: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"case clauses did not match",
					newTestNode(
						"lists did not match",
						newTestNode(
							"expression lists did not match",
							newTestNode(
								"expressions at index 0 did not match",
								newTestNode(
									"expressions did not match",
									newTestNode(
										"identifiers did not match",
										newTestNode("strings did not match: a < b", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.CaseClause{},
			b:    &ast.CaseClause{},
			want: 0,
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.INT,
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.FLOAT,
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"switch statements did not match",
					newTestNode(
						"init statements did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.SwitchStmt{},
			b:    &ast.SwitchStmt{},
			want: 0,
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
					Tok: token.INT,
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.FLOAT,
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"type switch statements did not match",
					newTestNode(
						"init statements did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.TypeSwitchStmt{},
			b:    &ast.TypeSwitchStmt{},
			want: 0,
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.INT,
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.FLOAT,
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"comm clauses did not match",
					newTestNode(
						"comm statements did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.CommClause{},
			b:    &ast.CommClause{},
			want: 0,
		},
		{
			a: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("a"),
							},
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("a"),
							},
							Tok: token.FLOAT,
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"select statements did not match",
					newTestNode(
						"bodies did not match",
						newTestNode(
							"block statements did not match",
							newTestNode(
								"statements at index 0 did not match",
								newTestNode(
									"statements did not match",
									newTestNode(
										"assign statements did not match",
										newTestNode(
											"tokens did not match",
											newTestNode(
												"tokens did not match",
												newTestNode("ints did not match: 5 < 6", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.SelectStmt{},
			b:    &ast.SelectStmt{},
			want: 0,
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.INT,
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("a"),
					},
					Tok: token.FLOAT,
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"for statements did not match",
					newTestNode(
						"init statements did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.ForStmt{},
			b:    &ast.ForStmt{},
			want: 0,
		},
		{
			a: &ast.RangeStmt{
				Key: ast.NewIdent("a"),
			},
			b: &ast.RangeStmt{
				Key: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"statements did not match",
				newTestNode(
					"range statements did not match",
					newTestNode(
						"key statements did not match",
						newTestNode(
							"expressions did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a:    &ast.RangeStmt{},
			b:    &ast.RangeStmt{},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareCompositeLiterals(t *testing.T) {
	testCases := []struct {
		a        *ast.CompositeLit
		b        *ast.CompositeLit
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.CompositeLit{},
			want: 1,
			wantNode: newTestNode(
				"composite literals did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.CompositeLit{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"composite literals did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
			},
			b: &ast.CompositeLit{
				Type: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"composite literals did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.CompositeLit{
				Type: ast.NewIdent("b"),
			},
			b: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"composite literals did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
				Elts: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			b: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
				Elts: []ast.Expr{
					ast.NewIdent("x"),
					ast.NewIdent("y"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"composite literals did not match",
				newTestNode(
					"expression lists did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 1 < 2", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
				Elts: []ast.Expr{
					ast.NewIdent("x"),
					ast.NewIdent("y"),
				},
			},
			b: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
				Elts: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"composite literals did not match",
				newTestNode(
					"expression lists did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 2 > 1", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.CompositeLit{
				Type:       ast.NewIdent("a"),
				Elts:       []ast.Expr{},
				Incomplete: false,
			},
			b: &ast.CompositeLit{
				Type:       ast.NewIdent("a"),
				Elts:       []ast.Expr{},
				Incomplete: true,
			},
			want: -1,
			wantNode: newTestNode(
				"composite literals did not match",
				newTestNode(
					"incomplete values did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.CompositeLit{
				Type:       ast.NewIdent("a"),
				Elts:       []ast.Expr{},
				Incomplete: true,
			},
			b: &ast.CompositeLit{
				Type:       ast.NewIdent("a"),
				Elts:       []ast.Expr{},
				Incomplete: false,
			},
			want: 1,
			wantNode: newTestNode(
				"composite literals did not match",
				newTestNode(
					"incomplete values did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a: &ast.CompositeLit{
				Type:       ast.NewIdent("a"),
				Elts:       []ast.Expr{},
				Incomplete: false,
			},
			b: &ast.CompositeLit{
				Type:       ast.NewIdent("a"),
				Elts:       []ast.Expr{},
				Incomplete: false,
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareCompositeLiterals(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareCompositeLiterals(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareParentheses(t *testing.T) {
	testCases := []struct {
		a        *ast.ParenExpr
		b        *ast.ParenExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.ParenExpr{},
			want: 1,
			wantNode: newTestNode(
				"parentheses did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.ParenExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"parentheses did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.ParenExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.ParenExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"parentheses did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ParenExpr{
				X: ast.NewIdent("b"),
			},
			b: &ast.ParenExpr{
				X: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"parentheses did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ParenExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.ParenExpr{
				X: ast.NewIdent("a"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareParentheses(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareParentheses(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareSelectors(t *testing.T) {
	testCases := []struct {
		a        *ast.SelectorExpr
		b        *ast.SelectorExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.SelectorExpr{},
			want: 1,
			wantNode: newTestNode(
				"selector expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.SelectorExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"selector expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
				X:   ast.NewIdent("x"),
			},
			b: &ast.SelectorExpr{
				Sel: ast.NewIdent("b"),
				X:   ast.NewIdent("x"),
			},
			want: -1,
			wantNode: newTestNode(
				"selector expressions did not match",
				newTestNode(
					"selectors did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: a < b", nil),
					),
				),
			),
		},
		{
			a: &ast.SelectorExpr{
				Sel: ast.NewIdent("b"),
				X:   ast.NewIdent("x"),
			},
			b: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
				X:   ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"selector expressions did not match",
				newTestNode(
					"selectors did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: b > a", nil),
					),
				),
			),
		},
		{
			a: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
				X:   ast.NewIdent("x"),
			},
			b: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
				X:   ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"selector expressions did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
				X:   ast.NewIdent("y"),
			},
			b: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
				X:   ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"selector expressions did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
				X:   ast.NewIdent("x"),
			},
			b: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
				X:   ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareSelectors(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareSelectors(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareIndexExpressions(t *testing.T) {
	testCases := []struct {
		a        *ast.IndexExpr
		b        *ast.IndexExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.IndexExpr{},
			want: 1,
			wantNode: newTestNode(
				"index expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.IndexExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"index expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
				X:     ast.NewIdent("x"),
			},
			b: &ast.IndexExpr{
				Index: ast.NewIdent("b"),
				X:     ast.NewIdent("x"),
			},
			want: -1,
			wantNode: newTestNode(
				"index expressions did not match",
				newTestNode(
					"indices did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.IndexExpr{
				Index: ast.NewIdent("b"),
				X:     ast.NewIdent("x"),
			},
			b: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
				X:     ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"index expressions did not match",
				newTestNode(
					"indices did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
				X:     ast.NewIdent("x"),
			},
			b: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
				X:     ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"index expressions did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
				X:     ast.NewIdent("y"),
			},
			b: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
				X:     ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"index expressions did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
				X:     ast.NewIdent("x"),
			},
			b: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
				X:     ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareIndexExpressions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareIndexExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareSliceExpressions(t *testing.T) {
	testCases := []struct {
		a        *ast.SliceExpr
		b        *ast.SliceExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.SliceExpr{},
			want: 1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.SliceExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.SliceExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X: ast.NewIdent("b"),
			},
			b: &ast.SliceExpr{
				X: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:   ast.NewIdent("a"),
				Low: ast.NewIdent("x"),
			},
			b: &ast.SliceExpr{
				X:   ast.NewIdent("a"),
				Low: ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"low expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:   ast.NewIdent("a"),
				Low: ast.NewIdent("y"),
			},
			b: &ast.SliceExpr{
				X:   ast.NewIdent("a"),
				Low: ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"low expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:    ast.NewIdent("a"),
				Low:  ast.NewIdent("x"),
				High: ast.NewIdent("i"),
			},
			b: &ast.SliceExpr{
				X:    ast.NewIdent("a"),
				Low:  ast.NewIdent("x"),
				High: ast.NewIdent("j"),
			},
			want: -1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"high expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: i < j", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:    ast.NewIdent("a"),
				Low:  ast.NewIdent("x"),
				High: ast.NewIdent("j"),
			},
			b: &ast.SliceExpr{
				X:    ast.NewIdent("a"),
				Low:  ast.NewIdent("x"),
				High: ast.NewIdent("i"),
			},
			want: 1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"high expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: j > i", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:    ast.NewIdent("a"),
				Low:  ast.NewIdent("x"),
				High: ast.NewIdent("i"),
				Max:  ast.NewIdent("m"),
			},
			b: &ast.SliceExpr{
				X:    ast.NewIdent("a"),
				Low:  ast.NewIdent("x"),
				High: ast.NewIdent("i"),
				Max:  ast.NewIdent("n"),
			},
			want: -1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"max expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: m < n", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:    ast.NewIdent("a"),
				Low:  ast.NewIdent("x"),
				High: ast.NewIdent("i"),
				Max:  ast.NewIdent("n"),
			},
			b: &ast.SliceExpr{
				X:    ast.NewIdent("a"),
				Low:  ast.NewIdent("x"),
				High: ast.NewIdent("i"),
				Max:  ast.NewIdent("m"),
			},
			want: 1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"max expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: n > m", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:      ast.NewIdent("a"),
				Low:    ast.NewIdent("x"),
				High:   ast.NewIdent("i"),
				Max:    ast.NewIdent("m"),
				Slice3: false,
			},
			b: &ast.SliceExpr{
				X:      ast.NewIdent("a"),
				Low:    ast.NewIdent("x"),
				High:   ast.NewIdent("i"),
				Max:    ast.NewIdent("m"),
				Slice3: true,
			},
			want: -1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"slice3 values did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:      ast.NewIdent("a"),
				Low:    ast.NewIdent("x"),
				High:   ast.NewIdent("i"),
				Max:    ast.NewIdent("m"),
				Slice3: true,
			},
			b: &ast.SliceExpr{
				X:      ast.NewIdent("a"),
				Low:    ast.NewIdent("x"),
				High:   ast.NewIdent("i"),
				Max:    ast.NewIdent("m"),
				Slice3: false,
			},
			want: 1,
			wantNode: newTestNode(
				"slices did not match",
				newTestNode(
					"slice3 values did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a: &ast.SliceExpr{
				X:      ast.NewIdent("a"),
				Low:    ast.NewIdent("x"),
				High:   ast.NewIdent("i"),
				Max:    ast.NewIdent("m"),
				Slice3: true,
			},
			b: &ast.SliceExpr{
				X:      ast.NewIdent("a"),
				Low:    ast.NewIdent("x"),
				High:   ast.NewIdent("i"),
				Max:    ast.NewIdent("m"),
				Slice3: true,
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareSliceExpressions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareSliceExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareTypeAssertions(t *testing.T) {
	testCases := []struct {
		a        *ast.TypeAssertExpr
		b        *ast.TypeAssertExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.TypeAssertExpr{},
			want: 1,
			wantNode: newTestNode(
				"type assertions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.TypeAssertExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"type assertions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
				X:    ast.NewIdent("x"),
			},
			b: &ast.TypeAssertExpr{
				Type: ast.NewIdent("b"),
				X:    ast.NewIdent("x"),
			},
			want: -1,
			wantNode: newTestNode(
				"type assertions did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeAssertExpr{
				Type: ast.NewIdent("b"),
				X:    ast.NewIdent("x"),
			},
			b: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
				X:    ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"type assertions did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
				X:    ast.NewIdent("x"),
			},
			b: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
				X:    ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"type assertions did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
				X:    ast.NewIdent("y"),
			},
			b: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
				X:    ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"type assertions did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
				X:    ast.NewIdent("x"),
			},
			b: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
				X:    ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareTypeAssertions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareTypeAssertions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareCallExpressions(t *testing.T) {
	testCases := []struct {
		a        *ast.CallExpr
		b        *ast.CallExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.CallExpr{},
			want: 1,
			wantNode: newTestNode(
				"call expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.CallExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"call expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
			},
			b: &ast.CallExpr{
				Fun: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"call expressions did not match",
				newTestNode(
					"functions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.CallExpr{
				Fun: ast.NewIdent("b"),
			},
			b: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"call expressions did not match",
				newTestNode(
					"functions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.CallExpr{
				Fun:  ast.NewIdent("a"),
				Args: []ast.Expr{},
			},
			b: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
				Args: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"call expressions did not match",
				newTestNode(
					"arguments did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 0 < 1", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
				Args: []ast.Expr{
					ast.NewIdent("x"),
				}},
			b: &ast.CallExpr{
				Fun:  ast.NewIdent("a"),
				Args: []ast.Expr{},
			},
			want: 1,
			wantNode: newTestNode(
				"call expressions did not match",
				newTestNode(
					"arguments did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 1 > 0", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
				Args: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			b: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
				Args: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareCallExpressions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareCallExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareStarExpressions(t *testing.T) {
	testCases := []struct {
		a        *ast.StarExpr
		b        *ast.StarExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.StarExpr{},
			want: 1,
			wantNode: newTestNode(
				"star expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.StarExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"star expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.StarExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.StarExpr{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"star expressions did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.StarExpr{
				X: ast.NewIdent("b"),
			},
			b: &ast.StarExpr{
				X: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"star expressions did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.StarExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.StarExpr{
				X: ast.NewIdent("a"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareStarExpressions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareStarExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareUnaryExpressions(t *testing.T) {
	testCases := []struct {
		a        *ast.UnaryExpr
		b        *ast.UnaryExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.UnaryExpr{},
			want: 1,
			wantNode: newTestNode(
				"unary expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.UnaryExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"unary expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.UnaryExpr{
				Op: token.INT,
			},
			b: &ast.UnaryExpr{
				Op: token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"unary expressions did not match",
				newTestNode(
					"operators did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 5 < 6", nil),
					),
				),
			),
		},
		{
			a: &ast.UnaryExpr{
				Op: token.FLOAT,
			},
			b: &ast.UnaryExpr{
				Op: token.INT,
			},
			want: 1,
			wantNode: newTestNode(
				"unary expressions did not match",
				newTestNode(
					"operators did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 6 > 5", nil),
					),
				),
			),
		},
		{
			a: &ast.UnaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("x"),
			},
			b: &ast.UnaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"unary expressions did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.UnaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("y"),
			},
			b: &ast.UnaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"unary expressions did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.UnaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("x"),
			},
			b: &ast.UnaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareUnaryExpressions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareUnaryExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareBinaryExpressions(t *testing.T) {
	testCases := []struct {
		a        *ast.BinaryExpr
		b        *ast.BinaryExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.BinaryExpr{},
			want: 1,
			wantNode: newTestNode(
				"binary expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.BinaryExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"binary expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
			},
			b: &ast.BinaryExpr{
				Op: token.FLOAT,
				X:  ast.NewIdent("a"),
			},
			want: -1,
			wantNode: newTestNode(
				"binary expressions did not match",
				newTestNode(
					"operators did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 5 < 6", nil),
					),
				),
			),
		},
		{
			a: &ast.BinaryExpr{
				Op: token.FLOAT,
				X:  ast.NewIdent("a"),
			},
			b: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"binary expressions did not match",
				newTestNode(
					"operators did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 6 > 5", nil),
					),
				),
			),
		},
		{
			a: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
			},
			b: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"binary expressions did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("b"),
			},
			b: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"binary expressions did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
				Y:  ast.NewIdent("x"),
			},
			b: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
				Y:  ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"binary expressions did not match",
				newTestNode(
					"Y expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
				Y:  ast.NewIdent("y"),
			},
			b: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
				Y:  ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"binary expressions did not match",
				newTestNode(
					"Y expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
				Y:  ast.NewIdent("x"),
			},
			b: &ast.BinaryExpr{
				Op: token.INT,
				X:  ast.NewIdent("a"),
				Y:  ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareBinaryExpressions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareBinaryExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareKeyValueExpressions(t *testing.T) {
	testCases := []struct {
		a        *ast.KeyValueExpr
		b        *ast.KeyValueExpr
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.KeyValueExpr{},
			want: 1,
			wantNode: newTestNode(
				"key-value expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.KeyValueExpr{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"key-value expressions did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.KeyValueExpr{
				Key: ast.NewIdent("a"),
			},
			b: &ast.KeyValueExpr{
				Key: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"key-value expressions did not match",
				newTestNode(
					"keys did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.KeyValueExpr{
				Key: ast.NewIdent("b"),
			},
			b: &ast.KeyValueExpr{
				Key: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"key-value expressions did not match",
				newTestNode(
					"keys did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.KeyValueExpr{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			b: &ast.KeyValueExpr{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"key-value expressions did not match",
				newTestNode(
					"values did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.KeyValueExpr{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("y"),
			},
			b: &ast.KeyValueExpr{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"key-value expressions did not match",
				newTestNode(
					"values did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.KeyValueExpr{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			b: &ast.KeyValueExpr{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareKeyValueExpressions(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareKeyValueExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareArrayTypes(t *testing.T) {
	testCases := []struct {
		a        *ast.ArrayType
		b        *ast.ArrayType
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.ArrayType{},
			want: 1,
			wantNode: newTestNode(
				"array types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.ArrayType{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"array types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.ArrayType{
				Len: ast.NewIdent("a"),
			},
			b: &ast.ArrayType{
				Len: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"array types did not match",
				newTestNode(
					"length expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ArrayType{
				Len: ast.NewIdent("b"),
			},
			b: &ast.ArrayType{
				Len: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"array types did not match",
				newTestNode(
					"length expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ArrayType{
				Len: ast.NewIdent("a"),
				Elt: ast.NewIdent("x"),
			},
			b: &ast.ArrayType{
				Len: ast.NewIdent("a"),
				Elt: ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"array types did not match",
				newTestNode(
					"element type expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ArrayType{
				Len: ast.NewIdent("a"),
				Elt: ast.NewIdent("y"),
			},
			b: &ast.ArrayType{
				Len: ast.NewIdent("a"),
				Elt: ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"array types did not match",
				newTestNode(
					"element type expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ArrayType{
				Len: ast.NewIdent("a"),
				Elt: ast.NewIdent("x"),
			},
			b: &ast.ArrayType{
				Len: ast.NewIdent("a"),
				Elt: ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareArrayTypes(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareArrayTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareStructTypes(t *testing.T) {
	testCases := []struct {
		a        *ast.StructType
		b        *ast.StructType
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.StructType{},
			want: 1,
			wantNode: newTestNode(
				"struct types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.StructType{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"struct types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.StructType{
				Incomplete: false,
			},
			b: &ast.StructType{
				Incomplete: true,
			},
			want: -1,
			wantNode: newTestNode(
				"struct types did not match",
				newTestNode(
					"incomplete values did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.StructType{
				Incomplete: true,
			},
			b: &ast.StructType{
				Incomplete: false,
			},
			want: 1,
			wantNode: newTestNode(
				"struct types did not match",
				newTestNode(
					"incomplete values did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a: &ast.StructType{
				Incomplete: true,
				Fields: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			b: &ast.StructType{
				Incomplete: true,
				Fields: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("b"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"struct types did not match",
				newTestNode(
					"field lists did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"fields at index 0 did not match",
							newTestNode(
								"fields did not match",
								newTestNode(
									"name lists did not match",
									newTestNode(
										"identifier lists did not match",
										newTestNode(
											"identifiers at index 0 did not match",
											newTestNode(
												"identifiers did not match",
												newTestNode("strings did not match: a < b", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.StructType{
				Incomplete: true,
				Fields: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("b"),
							},
						},
					},
				},
			},
			b: &ast.StructType{
				Incomplete: true,
				Fields: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"struct types did not match",
				newTestNode(
					"field lists did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"fields at index 0 did not match",
							newTestNode(
								"fields did not match",
								newTestNode(
									"name lists did not match",
									newTestNode(
										"identifier lists did not match",
										newTestNode(
											"identifiers at index 0 did not match",
											newTestNode(
												"identifiers did not match",
												newTestNode("strings did not match: b > a", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.StructType{
				Incomplete: true,
				Fields: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			b: &ast.StructType{
				Incomplete: true,
				Fields: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareStructTypes(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareStructTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareFunctionTypes(t *testing.T) {
	testCases := []struct {
		a        *ast.FuncType
		b        *ast.FuncType
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.FuncType{},
			want: 1,
			wantNode: newTestNode(
				"function types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.FuncType{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"function types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			b: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("b"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"function types did not match",
				newTestNode(
					"parameter lists did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"fields at index 0 did not match",
							newTestNode(
								"fields did not match",
								newTestNode(
									"name lists did not match",
									newTestNode(
										"identifier lists did not match",
										newTestNode(
											"identifiers at index 0 did not match",
											newTestNode(
												"identifiers did not match",
												newTestNode("strings did not match: a < b", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("b"),
							},
						},
					},
				},
			},
			b: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"function types did not match",
				newTestNode(
					"parameter lists did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"fields at index 0 did not match",
							newTestNode(
								"fields did not match",
								newTestNode(
									"name lists did not match",
									newTestNode(
										"identifier lists did not match",
										newTestNode(
											"identifiers at index 0 did not match",
											newTestNode(
												"identifiers did not match",
												newTestNode("strings did not match: b > a", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("x"),
							},
						},
					},
				},
			},
			b: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("y"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"function types did not match",
				newTestNode(
					"result lists did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"fields at index 0 did not match",
							newTestNode(
								"fields did not match",
								newTestNode(
									"name lists did not match",
									newTestNode(
										"identifier lists did not match",
										newTestNode(
											"identifiers at index 0 did not match",
											newTestNode(
												"identifiers did not match",
												newTestNode("strings did not match: x < y", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("y"),
							},
						},
					},
				},
			},
			b: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("x"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"function types did not match",
				newTestNode(
					"result lists did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"fields at index 0 did not match",
							newTestNode(
								"fields did not match",
								newTestNode(
									"name lists did not match",
									newTestNode(
										"identifier lists did not match",
										newTestNode(
											"identifiers at index 0 did not match",
											newTestNode(
												"identifiers did not match",
												newTestNode("strings did not match: y > x", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("x"),
							},
						},
					},
				},
			},
			b: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("x"),
							},
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareFunctionTypes(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareFunctionTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareFieldLists(t *testing.T) {
	testCases := []struct {
		a        *ast.FieldList
		b        *ast.FieldList
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.FieldList{},
			want: 1,
			wantNode: newTestNode(
				"field lists did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.FieldList{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"field lists did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("a"),
					},
				},
			},
			b: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("a"),
					},
					{
						Type: ast.NewIdent("a"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"field lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 1 < 2", nil),
				),
			),
		},
		{
			a: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("a"),
					},
					{
						Type: ast.NewIdent("a"),
					},
				},
			},
			b: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("a"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"field lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 2 > 1", nil),
				),
			),
		},
		{
			a: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("a"),
						},
					},
				},
			},
			b: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("b"),
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"field lists did not match",
				newTestNode(
					"fields at index 0 did not match",
					newTestNode(
						"fields did not match",
						newTestNode(
							"name lists did not match",
							newTestNode(
								"identifier lists did not match",
								newTestNode(
									"identifiers at index 0 did not match",
									newTestNode(
										"identifiers did not match",
										newTestNode("strings did not match: a < b", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("b"),
						},
					},
				},
			},
			b: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("a"),
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"field lists did not match",
				newTestNode(
					"fields at index 0 did not match",
					newTestNode(
						"fields did not match",
						newTestNode(
							"name lists did not match",
							newTestNode(
								"identifier lists did not match",
								newTestNode(
									"identifiers at index 0 did not match",
									newTestNode(
										"identifiers did not match",
										newTestNode("strings did not match: b > a", nil),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("a"),
						},
					},
				},
			},
			b: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("a"),
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareFieldLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareFieldLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareFields(t *testing.T) {
	testCases := []struct {
		a        *ast.Field
		b        *ast.Field
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.Field{},
			want: 1,
			wantNode: newTestNode(
				"fields did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.Field{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"fields did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
			},
			b: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"fields did not match",
				newTestNode(
					"name lists did not match",
					newTestNode(
						"identifier lists did not match",
						newTestNode(
							"identifiers at index 0 did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("b"),
				},
			},
			b: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"fields did not match",
				newTestNode(
					"name lists did not match",
					newTestNode(
						"identifier lists did not match",
						newTestNode(
							"identifiers at index 0 did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: b > a", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("x"),
			},
			b: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"fields did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("y"),
			},
			b: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"fields did not match",
				newTestNode(
					"types did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("x"),
				Tag: &ast.BasicLit{
					Kind: token.INT,
				},
			},
			b: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("x"),
				Tag: &ast.BasicLit{
					Kind: token.FLOAT,
				},
			},
			want: -1,
			wantNode: newTestNode(
				"fields did not match",
				newTestNode(
					"tags did not match",
					newTestNode(
						"basic literals did not match",
						newTestNode(
							"kinds did not match",
							newTestNode(
								"tokens did not match",
								newTestNode("ints did not match: 5 < 6", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("x"),
				Tag: &ast.BasicLit{
					Kind: token.FLOAT,
				},
			},
			b: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("x"),
				Tag: &ast.BasicLit{
					Kind: token.INT,
				},
			},
			want: 1,
			wantNode: newTestNode(
				"fields did not match",
				newTestNode(
					"tags did not match",
					newTestNode(
						"basic literals did not match",
						newTestNode(
							"kinds did not match",
							newTestNode(
								"tokens did not match",
								newTestNode("ints did not match: 6 > 5", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("x"),
				Tag: &ast.BasicLit{
					Kind: token.INT,
				},
			},
			b: &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("a"),
				},
				Type: ast.NewIdent("x"),
				Tag: &ast.BasicLit{
					Kind: token.INT,
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareFields(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareFields(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareInterfaceTypes(t *testing.T) {
	testCases := []struct {
		a        *ast.InterfaceType
		b        *ast.InterfaceType
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.InterfaceType{},
			want: 1,
			wantNode: newTestNode(
				"interface types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.InterfaceType{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"interface types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			b: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("b"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"interface types did not match",
				newTestNode(
					"method lists did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"fields at index 0 did not match",
							newTestNode(
								"fields did not match",
								newTestNode(
									"name lists did not match",
									newTestNode(
										"identifier lists did not match",
										newTestNode(
											"identifiers at index 0 did not match",
											newTestNode(
												"identifiers did not match",
												newTestNode("strings did not match: a < b", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("b"),
							},
						},
					},
				},
			},
			b: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"interface types did not match",
				newTestNode(
					"method lists did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"fields at index 0 did not match",
							newTestNode(
								"fields did not match",
								newTestNode(
									"name lists did not match",
									newTestNode(
										"identifier lists did not match",
										newTestNode(
											"identifiers at index 0 did not match",
											newTestNode(
												"identifiers did not match",
												newTestNode("strings did not match: b > a", nil),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Incomplete: false,
			},
			b: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Incomplete: true,
			},
			want: -1,
			wantNode: newTestNode(
				"interface types did not match",
				newTestNode(
					"incomplete values did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Incomplete: true,
			},
			b: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Incomplete: false,
			},
			want: 1,
			wantNode: newTestNode(
				"interface types did not match",
				newTestNode(
					"incomplete values did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Incomplete: true,
			},
			b: &ast.InterfaceType{
				Methods: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("a"),
							},
						},
					},
				},

				Incomplete: true,
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareInterfaceTypes(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareInterfaceTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareMapTypes(t *testing.T) {
	testCases := []struct {
		a        *ast.MapType
		b        *ast.MapType
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.MapType{},
			want: 1,
			wantNode: newTestNode(
				"map types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.MapType{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"map types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.MapType{
				Key: ast.NewIdent("a"),
			},
			b: &ast.MapType{
				Key: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"map types did not match",
				newTestNode(
					"key expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.MapType{
				Key: ast.NewIdent("b"),
			},
			b: &ast.MapType{
				Key: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"map types did not match",
				newTestNode(
					"key expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.MapType{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			b: &ast.MapType{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"map types did not match",
				newTestNode(
					"value expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.MapType{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("y"),
			},
			b: &ast.MapType{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"map types did not match",
				newTestNode(
					"value expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.MapType{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			b: &ast.MapType{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareMapTypes(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareMapTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareChannelTypes(t *testing.T) {
	testCases := []struct {
		a        *ast.ChanType
		b        *ast.ChanType
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.ChanType{},
			want: 1,
			wantNode: newTestNode(
				"channel types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.ChanType{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"channel types did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.ChanType{
				Dir: ast.SEND,
			},
			b: &ast.ChanType{
				Dir: ast.RECV,
			},
			want: -1,
			wantNode: newTestNode(
				"channel types did not match",
				newTestNode(
					"directions did not match",
					newTestNode(
						"channel directions did not match",
						newTestNode("ints did not match: 1 < 2", nil),
					),
				),
			),
		},
		{
			a: &ast.ChanType{
				Dir: ast.RECV,
			},
			b: &ast.ChanType{
				Dir: ast.SEND,
			},
			want: 1,
			wantNode: newTestNode(
				"channel types did not match",
				newTestNode(
					"directions did not match",
					newTestNode(
						"channel directions did not match",
						newTestNode("ints did not match: 2 > 1", nil),
					),
				),
			),
		},
		{
			a: &ast.ChanType{
				Dir:   ast.SEND,
				Value: ast.NewIdent("a"),
			},
			b: &ast.ChanType{
				Dir:   ast.SEND,
				Value: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"channel types did not match",
				newTestNode(
					"value expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ChanType{
				Dir:   ast.SEND,
				Value: ast.NewIdent("b"),
			},
			b: &ast.ChanType{
				Dir:   ast.SEND,
				Value: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"channel types did not match",
				newTestNode(
					"value expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ChanType{
				Dir:   ast.SEND,
				Value: ast.NewIdent("a"),
			},
			b: &ast.ChanType{
				Dir:   ast.SEND,
				Value: ast.NewIdent("a"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareChannelTypes(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareChannelTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareChannelDirections(t *testing.T) {
	testCases := []struct {
		a        ast.ChanDir
		b        ast.ChanDir
		want     int
		wantNode *node
	}{
		{
			a:    ast.SEND,
			b:    ast.RECV,
			want: -1,
			wantNode: newTestNode(
				"channel directions did not match",
				newTestNode("ints did not match: 1 < 2", nil),
			),
		},
		{
			a:    ast.RECV,
			b:    ast.SEND,
			want: 1,
			wantNode: newTestNode(
				"channel directions did not match",
				newTestNode("ints did not match: 2 > 1", nil),
			),
		},
		{
			a:    ast.RECV,
			b:    ast.RECV,
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareChannelDirections(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareChannelDirections(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareBadStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.BadStmt
		b        *ast.BadStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.BadStmt{},
			want: 1,
			wantNode: newTestNode(
				"bad statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.BadStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"bad statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a:    &ast.BadStmt{},
			b:    &ast.BadStmt{},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareBadStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareBadStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareDeclStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.DeclStmt
		b        *ast.DeclStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.DeclStmt{},
			want: 1,
			wantNode: newTestNode(
				"declaration statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.DeclStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"declaration statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
			},
			b: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"declaration statements did not match",
				newTestNode(
					"declarations did not match",
					newTestNode(
						"function declarations did not match",
						newTestNode(
							"names did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: a < b", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			b: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"declaration statements did not match",
				newTestNode(
					"declarations did not match",
					newTestNode(
						"function declarations did not match",
						newTestNode(
							"names did not match",
							newTestNode(
								"identifiers did not match",
								newTestNode("strings did not match: b > a", nil),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
			},
			b: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareDeclStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareDeclStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareDecls(t *testing.T) {
	testCases := []struct {
		a        ast.Decl
		b        ast.Decl
		want     int
		wantNode *node
	}{
		{
			a:    &ast.BadDecl{},
			b:    &ast.GenDecl{},
			want: -1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode("declaration types did not match", nil),
			),
		},
		{
			a:    &ast.BadDecl{},
			b:    &ast.FuncDecl{},
			want: -1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode("declaration types did not match", nil),
			),
		},
		{
			a:    &ast.GenDecl{},
			b:    &ast.FuncDecl{},
			want: -1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode("declaration types did not match", nil),
			),
		},
		{
			a:    &ast.GenDecl{},
			b:    &ast.BadDecl{},
			want: 1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode("declaration types did not match", nil),
			),
		},
		{
			a:    &ast.FuncDecl{},
			b:    &ast.BadDecl{},
			want: 1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode("declaration types did not match", nil),
			),
		},
		{
			a:    &ast.FuncDecl{},
			b:    &ast.GenDecl{},
			want: 1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode("declaration types did not match", nil),
			),
		},
		{
			a:    &ast.BadDecl{},
			b:    &ast.BadDecl{},
			want: 0,
		},
		{
			a: &ast.GenDecl{
				Tok: token.INT,
			},
			b: &ast.GenDecl{
				Tok: token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode(
					"generic declarations did not match",
					newTestNode(
						"tokens did not match",
						newTestNode(
							"tokens did not match",
							newTestNode("ints did not match: 5 < 6", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.GenDecl{
				Tok: token.FLOAT,
			},
			b: &ast.GenDecl{
				Tok: token.INT,
			},
			want: 1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode(
					"generic declarations did not match",
					newTestNode(
						"tokens did not match",
						newTestNode(
							"tokens did not match",
							newTestNode("ints did not match: 6 > 5", nil),
						),
					),
				),
			),
		},
		{
			a:    &ast.GenDecl{},
			b:    &ast.GenDecl{},
			want: 0,
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Type: &ast.FuncType{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("b"),
				Type: &ast.FuncType{},
			},
			want: -1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode(
					"function declarations did not match",
					newTestNode(
						"names did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("b"),
				Type: &ast.FuncType{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Type: &ast.FuncType{},
			},
			want: 1,
			wantNode: newTestNode(
				"declarations did not match",
				newTestNode(
					"function declarations did not match",
					newTestNode(
						"names did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Type: &ast.FuncType{},
			},
			b: &ast.FuncDecl{
				Type: &ast.FuncType{},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareDecls(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareDecls(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareBadDecls(t *testing.T) {
	testCases := []struct {
		a        *ast.BadDecl
		b        *ast.BadDecl
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.BadDecl{},
			want: 1,
			wantNode: newTestNode(
				"bad declarations did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.BadDecl{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"bad declarations did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a:    &ast.BadDecl{},
			b:    &ast.BadDecl{},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareBadDecls(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareBadDecls(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareGenDecls(t *testing.T) {
	testCases := []struct {
		a        *ast.GenDecl
		b        *ast.GenDecl
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.GenDecl{},
			want: 1,
			wantNode: newTestNode(
				"generic declarations did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.GenDecl{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"generic declarations did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.GenDecl{
				Tok: token.INT,
			},
			b: &ast.GenDecl{
				Tok: token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"generic declarations did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 5 < 6", nil),
					),
				),
			),
		},
		{
			a: &ast.GenDecl{
				Tok: token.FLOAT,
			},
			b: &ast.GenDecl{
				Tok: token.INT,
			},
			want: 1,
			wantNode: newTestNode(
				"generic declarations did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 6 > 5", nil),
					),
				),
			),
		},
		{
			a: &ast.GenDecl{
				Tok: token.INT,
				Specs: []ast.Spec{
					&ast.ImportSpec{},
				},
			},
			b: &ast.GenDecl{
				Tok: token.INT,
				Specs: []ast.Spec{
					&ast.ImportSpec{},
					&ast.ImportSpec{},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"generic declarations did not match",
				newTestNode(
					"spec lists did not match",
					newTestNode(
						"spec lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 1 < 2", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.GenDecl{
				Tok: token.INT,
				Specs: []ast.Spec{
					&ast.ImportSpec{},
					&ast.ImportSpec{},
				},
			},
			b: &ast.GenDecl{
				Tok: token.INT,
				Specs: []ast.Spec{
					&ast.ImportSpec{},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"generic declarations did not match",
				newTestNode(
					"spec lists did not match",
					newTestNode(
						"spec lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 2 > 1", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.GenDecl{
				Tok: token.INT,
				Specs: []ast.Spec{
					&ast.ImportSpec{},
				},
			},
			b: &ast.GenDecl{
				Tok: token.INT,
				Specs: []ast.Spec{
					&ast.ImportSpec{},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareGenDecls(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareGenDecls(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareFuncDecls(t *testing.T) {
	testCases := []struct {
		a        *ast.FuncDecl
		b        *ast.FuncDecl
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.FuncDecl{},
			want: 1,
			wantNode: newTestNode(
				"function declarations did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.FuncDecl{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"function declarations did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Type: &ast.FuncType{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("b"),
				Type: &ast.FuncType{},
			},
			want: -1,
			wantNode: newTestNode(
				"function declarations did not match",
				newTestNode(
					"names did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: a < b", nil),
					),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("b"),
				Type: &ast.FuncType{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Type: &ast.FuncType{},
			},
			want: 1,
			wantNode: newTestNode(
				"function declarations did not match",
				newTestNode(
					"names did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: b > a", nil),
					),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: nil,
				Type: &ast.FuncType{},
			},
			want: -1,
			wantNode: newTestNode(
				"function declarations did not match",
				newTestNode(
					"receivers did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"nil comparisons did not match",
							newTestNode("bools did not match: false < true", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: nil,
				Type: &ast.FuncType{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
			},
			want: 1,
			wantNode: newTestNode(
				"function declarations did not match",
				newTestNode(
					"receivers did not match",
					newTestNode(
						"field lists did not match",
						newTestNode(
							"nil comparisons did not match",
							newTestNode("bools did not match: true > false", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
				Body: nil,
			},
			want: -1,
			wantNode: newTestNode(
				"function declarations did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"nil comparisons did not match",
							newTestNode("bools did not match: false < true", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
				Body: nil,
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{},
			},
			want: 1,
			wantNode: newTestNode(
				"function declarations did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"nil comparisons did not match",
							newTestNode("bools did not match: true > false", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareFuncDecls(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareFuncDecls(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareEmptyStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.EmptyStmt
		b        *ast.EmptyStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.EmptyStmt{},
			want: 1,
			wantNode: newTestNode(
				"empty statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.EmptyStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"empty statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.EmptyStmt{
				Implicit: false,
			},
			b: &ast.EmptyStmt{
				Implicit: true,
			},
			want: -1,
			wantNode: newTestNode(
				"empty statements did not match",
				newTestNode("bools did not match: false < true", nil),
			),
		},
		{
			a: &ast.EmptyStmt{
				Implicit: true,
			},
			b: &ast.EmptyStmt{
				Implicit: false,
			},
			want: 1,
			wantNode: newTestNode(
				"empty statements did not match",
				newTestNode("bools did not match: true > false", nil),
			),
		},
		{
			a: &ast.EmptyStmt{
				Implicit: true,
			},
			b: &ast.EmptyStmt{
				Implicit: true,
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareEmptyStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareEmptyStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareLabeledStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.LabeledStmt
		b        *ast.LabeledStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.LabeledStmt{},
			want: 1,
			wantNode: newTestNode(
				"labeled statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.LabeledStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"labeled statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"labeled statements did not match",
				newTestNode(
					"labels did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: a < b", nil),
					),
				),
			),
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("b"),
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"labeled statements did not match",
				newTestNode(
					"labels did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: b > a", nil),
					),
				),
			),
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"labeled statements did not match",
				newTestNode(
					"statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"labeled statements did not match",
				newTestNode(
					"statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareLabeledStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareLabeledStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareExpressionStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.ExprStmt
		b        *ast.ExprStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.ExprStmt{},
			want: 1,
			wantNode: newTestNode(
				"expression statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.ExprStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"expression statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.ExprStmt{
				X: ast.NewIdent("a"),
			},
			b: &ast.ExprStmt{
				X: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"expression statements did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: a < b", nil),
					),
				),
			),
		},
		{
			a: &ast.ExprStmt{
				X: ast.NewIdent("b"),
			},
			b: &ast.ExprStmt{
				X: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"expression statements did not match",
				newTestNode(
					"expressions did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: b > a", nil),
					),
				),
			),
		},
		{
			a: &ast.ExprStmt{
				X: ast.NewIdent("a"),
			},
			b: &ast.ExprStmt{
				X: ast.NewIdent("a"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareExpressionStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareExpressionStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareSendStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.SendStmt
		b        *ast.SendStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.SendStmt{},
			want: 1,
			wantNode: newTestNode(
				"send statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.SendStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"send statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.SendStmt{
				Chan: ast.NewIdent("a"),
			},
			b: &ast.SendStmt{
				Chan: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"send statements did not match",
				newTestNode(
					"channels did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SendStmt{
				Chan: ast.NewIdent("b"),
			},
			b: &ast.SendStmt{
				Chan: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"send statements did not match",
				newTestNode(
					"channels did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SendStmt{
				Chan:  ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			b: &ast.SendStmt{
				Chan:  ast.NewIdent("a"),
				Value: ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"send statements did not match",
				newTestNode(
					"values did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SendStmt{
				Chan:  ast.NewIdent("a"),
				Value: ast.NewIdent("y"),
			},
			b: &ast.SendStmt{
				Chan:  ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"send statements did not match",
				newTestNode(
					"values did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SendStmt{
				Chan:  ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			b: &ast.SendStmt{
				Chan:  ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareSendStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareSendStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareIncDecStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.IncDecStmt
		b        *ast.IncDecStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.IncDecStmt{},
			want: 1,
			wantNode: newTestNode(
				"increment/decrement statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.IncDecStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"increment/decrement statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.IncDecStmt{
				Tok: token.INT,
				X:   ast.NewIdent("a"),
			},
			b: &ast.IncDecStmt{
				Tok: token.FLOAT,
				X:   ast.NewIdent("a"),
			},
			want: -1,
			wantNode: newTestNode(
				"increment/decrement statements did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 5 < 6", nil),
					),
				),
			),
		},
		{
			a: &ast.IncDecStmt{
				Tok: token.FLOAT,
				X:   ast.NewIdent("a"),
			},
			b: &ast.IncDecStmt{
				Tok: token.INT,
				X:   ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"increment/decrement statements did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 6 > 5", nil),
					),
				),
			),
		},
		{
			a: &ast.IncDecStmt{
				Tok: token.INT,
				X:   ast.NewIdent("a"),
			},
			b: &ast.IncDecStmt{
				Tok: token.INT,
				X:   ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"increment/decrement statements did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.IncDecStmt{
				Tok: token.INT,
				X:   ast.NewIdent("b"),
			},
			b: &ast.IncDecStmt{
				Tok: token.INT,
				X:   ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"increment/decrement statements did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.IncDecStmt{
				Tok: token.INT,
				X:   ast.NewIdent("a"),
			},
			b: &ast.IncDecStmt{
				Tok: token.INT,
				X:   ast.NewIdent("a"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareIncDecStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareIncDecStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareAssignStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.AssignStmt
		b        *ast.AssignStmt
		want     int
		wantNode *node
	}{
		{
			a: nil,
			b: &ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("i"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"assign statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("i"),
				},
			},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"assign statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			b: &ast.AssignStmt{
				Tok: token.FLOAT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"assign statements did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 5 < 6", nil),
					),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Tok: token.FLOAT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			b: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"assign statements did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 6 > 5", nil),
					),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			b: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
					ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"assign statements did not match",
				newTestNode(
					"lhs did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 1 < 2", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
					ast.NewIdent("b"),
				},
			},
			b: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"assign statements did not match",
				newTestNode(
					"lhs did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 2 > 1", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
				Rhs: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			b: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
				Rhs: []ast.Expr{
					ast.NewIdent("x"),
					ast.NewIdent("y"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"assign statements did not match",
				newTestNode(
					"rhs did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 1 < 2", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
				Rhs: []ast.Expr{
					ast.NewIdent("x"),
					ast.NewIdent("y"),
				},
			},
			b: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
				Rhs: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"assign statements did not match",
				newTestNode(
					"rhs did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 2 > 1", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
				Rhs: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			b: &ast.AssignStmt{
				Tok: token.INT,
				Lhs: []ast.Expr{
					ast.NewIdent("a"),
				},
				Rhs: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareAssignStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareAssignStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareGoStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.GoStmt
		b        *ast.GoStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.GoStmt{},
			want: 1,
			wantNode: newTestNode(
				"go statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.GoStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"go statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.GoStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			b: &ast.GoStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"go statements did not match",
				newTestNode(
					"call expressions did not match",
					newTestNode(
						"call expressions did not match",
						newTestNode(
							"functions did not match",
							newTestNode(
								"expressions did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.GoStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("b"),
				},
			},
			b: &ast.GoStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"go statements did not match",
				newTestNode(
					"call expressions did not match",
					newTestNode(
						"call expressions did not match",
						newTestNode(
							"functions did not match",
							newTestNode(
								"expressions did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: b > a", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.GoStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			b: &ast.GoStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareGoStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareGoStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareDeferStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.DeferStmt
		b        *ast.DeferStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.DeferStmt{},
			want: 1,
			wantNode: newTestNode(
				"defer statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.DeferStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"defer statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			b: &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"defer statements did not match",
				newTestNode(
					"call expressions did not match",
					newTestNode(
						"call expressions did not match",
						newTestNode(
							"functions did not match",
							newTestNode(
								"expressions did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("b"),
				},
			},
			b: &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"defer statements did not match",
				newTestNode(
					"call expressions did not match",
					newTestNode(
						"call expressions did not match",
						newTestNode(
							"functions did not match",
							newTestNode(
								"expressions did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: b > a", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			b: &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: ast.NewIdent("a"),
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareDeferStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareDeferStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareReturnStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.ReturnStmt
		b        *ast.ReturnStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.ReturnStmt{},
			want: 1,
			wantNode: newTestNode(
				"return statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.ReturnStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"return statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.ReturnStmt{
				Results: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			b: &ast.ReturnStmt{
				Results: []ast.Expr{
					ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"return statements did not match",
				newTestNode(
					"results did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"expressions at index 0 did not match",
							newTestNode(
								"expressions did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ReturnStmt{
				Results: []ast.Expr{
					ast.NewIdent("b"),
				},
			},
			b: &ast.ReturnStmt{
				Results: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"return statements did not match",
				newTestNode(
					"results did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"expressions at index 0 did not match",
							newTestNode(
								"expressions did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: b > a", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ReturnStmt{
				Results: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			b: &ast.ReturnStmt{
				Results: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareReturnStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareReturnStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareBranchStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.BranchStmt
		b        *ast.BranchStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.BranchStmt{},
			want: 1,
			wantNode: newTestNode(
				"branch statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.BranchStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"branch statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.BranchStmt{
				Tok: token.INT,
			},
			b: &ast.BranchStmt{
				Tok: token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"branch statements did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 5 < 6", nil),
					),
				),
			),
		},
		{
			a: &ast.BranchStmt{
				Tok: token.FLOAT,
			},
			b: &ast.BranchStmt{
				Tok: token.INT,
			},
			want: 1,
			wantNode: newTestNode(
				"branch statements did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 6 > 5", nil),
					),
				),
			),
		},
		{
			a: &ast.BranchStmt{
				Tok:   token.INT,
				Label: ast.NewIdent("a"),
			},
			b: &ast.BranchStmt{
				Tok:   token.INT,
				Label: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"branch statements did not match",
				newTestNode(
					"labels did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: a < b", nil),
					),
				),
			),
		},
		{
			a: &ast.BranchStmt{
				Tok:   token.INT,
				Label: ast.NewIdent("b"),
			},
			b: &ast.BranchStmt{
				Tok:   token.INT,
				Label: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"branch statements did not match",
				newTestNode(
					"labels did not match",
					newTestNode(
						"identifiers did not match",
						newTestNode("strings did not match: b > a", nil),
					),
				),
			),
		},
		{
			a: &ast.BranchStmt{
				Tok:   token.INT,
				Label: ast.NewIdent("a"),
			},
			b: &ast.BranchStmt{
				Tok:   token.INT,
				Label: ast.NewIdent("a"),
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareBranchStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareBranchStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareIfStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.IfStmt
		b        *ast.IfStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.IfStmt{},
			want: 1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.IfStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"init statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"init statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"conditions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("b"),
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"conditions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},

				Else: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("z"),
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("z"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"else statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("z"),
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("z"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"if statements did not match",
				newTestNode(
					"else statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("z"),
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("x"),
					},
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("y"),
							},
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("z"),
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareIfStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareIfStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareCaseClauses(t *testing.T) {
	testCases := []struct {
		a        *ast.CaseClause
		b        *ast.CaseClause
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.CaseClause{},
			want: 1,
			wantNode: newTestNode(
				"case clauses did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.CaseClause{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"case clauses did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			b: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("b"),
				},
			},
			want: -1,
			wantNode: newTestNode(
				"case clauses did not match",
				newTestNode(
					"lists did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"expressions at index 0 did not match",
							newTestNode(
								"expressions did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("b"),
				},
			},
			b: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
			},
			want: 1,
			wantNode: newTestNode(
				"case clauses did not match",
				newTestNode(
					"lists did not match",
					newTestNode(
						"expression lists did not match",
						newTestNode(
							"expressions at index 0 did not match",
							newTestNode(
								"expressions did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: b > a", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			b: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.FLOAT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"case clauses did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"statement lists did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.FLOAT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			b: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"case clauses did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"statement lists did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			b: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareCaseClauses(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareCaseClauses(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareStatementLists(t *testing.T) {
	testCases := []struct {
		a        []ast.Stmt
		b        []ast.Stmt
		want     int
		wantNode *node
	}{
		{
			a: []ast.Stmt{},
			b: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"statement lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 0 < 1", nil),
				),
			),
		},
		{
			a: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b:    []ast.Stmt{},
			want: 1,
			wantNode: newTestNode(
				"statement lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 1 > 0", nil),
				),
			),
		},
		{
			a: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				&ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				&ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: -1,
			wantNode: newNode(
				"statement lists did not match",
				nil,
				nil,
				&[]*node{
					newTestNode(
						"statements at index 0 did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
					newTestNode(
						"statements at index 1 did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 6 > 5", nil),
									),
								),
							),
						),
					),
				},
			),
		},
		{
			a: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				&ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				&ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: 1,
			wantNode: newNode(
				"statement lists did not match",
				nil,
				nil,
				&[]*node{
					newTestNode(
						"statements at index 0 did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 6 > 5", nil),
									),
								),
							),
						),
					),
					newTestNode(
						"statements at index 1 did not match",
						newTestNode(
							"statements did not match",
							newTestNode(
								"assign statements did not match",
								newTestNode(
									"tokens did not match",
									newTestNode(
										"tokens did not match",
										newTestNode("ints did not match: 5 < 6", nil),
									),
								),
							),
						),
					),
				},
			),
		},
		{
			a: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				&ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				&ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareStatementLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareStatementLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareSwitchStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.SwitchStmt
		b        *ast.SwitchStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.SwitchStmt{},
			want: 1,
			wantNode: newTestNode(
				"switch statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.SwitchStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"switch statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"switch statements did not match",
				newTestNode(
					"init statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"switch statements did not match",
				newTestNode(
					"init statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("a"),
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"switch statements did not match",
				newTestNode(
					"tags did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("b"),
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"switch statements did not match",
				newTestNode(
					"tags did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"switch statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"switch statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareSwitchStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareSwitchStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareTypeSwitchStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.TypeSwitchStmt
		b        *ast.TypeSwitchStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.TypeSwitchStmt{},
			want: 1,
			wantNode: newTestNode(
				"type switch statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.TypeSwitchStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"type switch statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"type switch statements did not match",
				newTestNode(
					"init statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"type switch statements did not match",
				newTestNode(
					"init statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"type switch statements did not match",
				newTestNode(
					"assign statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"type switch statements did not match",
				newTestNode(
					"assign statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"type switch statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"type switch statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareTypeSwitchStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareTypeSwitchStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareCommClauses(t *testing.T) {
	testCases := []struct {
		a        *ast.CommClause
		b        *ast.CommClause
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.CommClause{},
			want: 1,
			wantNode: newTestNode(
				"comm clauses did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.CommClause{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"comm clauses did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"comm clauses did not match",
				newTestNode(
					"comm statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"comm clauses did not match",
				newTestNode(
					"comm statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.FLOAT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"comm clauses did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"statement lists did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.FLOAT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"comm clauses did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"statement lists did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
						Lhs: []ast.Expr{
							ast.NewIdent("i"),
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareCommClauses(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareCommClauses(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareSelectStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.SelectStmt
		b        *ast.SelectStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.SelectStmt{},
			want: 1,
			wantNode: newTestNode(
				"select statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.SelectStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"select statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"select statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"select statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareSelectStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareSelectStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareForStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.ForStmt
		b        *ast.ForStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.ForStmt{},
			want: 1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.ForStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"init statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"init statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"conditions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("b"),
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"conditions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"post statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.FLOAT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"post statements did not match",
					newTestNode(
						"statements did not match",
						newTestNode(
							"assign statements did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"for statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
					Lhs: []ast.Expr{
						ast.NewIdent("i"),
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareForStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareForStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareRangeStatements(t *testing.T) {
	testCases := []struct {
		a        *ast.RangeStmt
		b        *ast.RangeStmt
		want     int
		wantNode *node
	}{
		{
			a:    nil,
			b:    &ast.RangeStmt{},
			want: 1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: true > false", nil),
				),
			),
		},
		{
			a:    &ast.RangeStmt{},
			b:    nil,
			want: -1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"nil comparisons did not match",
					newTestNode("bools did not match: false < true", nil),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key: ast.NewIdent("a"),
			},
			b: &ast.RangeStmt{
				Key: ast.NewIdent("b"),
			},
			want: -1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"key statements did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: a < b", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key: ast.NewIdent("b"),
			},
			b: &ast.RangeStmt{
				Key: ast.NewIdent("a"),
			},
			want: 1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"key statements did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: b > a", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("y"),
			},
			want: -1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"value statements did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: x < y", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("y"),
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
			},
			want: 1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"value statements did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: y > x", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.FLOAT,
			},
			want: -1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 5 < 6", nil),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.FLOAT,
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
			},
			want: 1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"tokens did not match",
					newTestNode(
						"tokens did not match",
						newTestNode("ints did not match: 6 > 5", nil),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("m"),
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("n"),
			},
			want: -1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: m < n", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("n"),
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("m"),
			},
			want: 1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"X expressions did not match",
					newTestNode(
						"expressions did not match",
						newTestNode(
							"identifiers did not match",
							newTestNode("strings did not match: n > m", nil),
						),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("m"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("m"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 5 < 6", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("m"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("m"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 1,
			wantNode: newTestNode(
				"range statements did not match",
				newTestNode(
					"bodies did not match",
					newTestNode(
						"block statements did not match",
						newTestNode(
							"statements at index 0 did not match",
							newTestNode(
								"statements did not match",
								newTestNode(
									"assign statements did not match",
									newTestNode(
										"tokens did not match",
										newTestNode(
											"tokens did not match",
											newTestNode("ints did not match: 6 > 5", nil),
										),
									),
								),
							),
						),
					),
				),
			),
		},
		{
			a: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("m"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			b: &ast.RangeStmt{
				Key:   ast.NewIdent("a"),
				Value: ast.NewIdent("x"),
				Tok:   token.INT,
				X:     ast.NewIdent("m"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
							Lhs: []ast.Expr{
								ast.NewIdent("i"),
							},
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareRangeStatements(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareRangeStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareGenDeclLists(t *testing.T) {
	testCases := []struct {
		a        []*ast.GenDecl
		b        []*ast.GenDecl
		want     int
		wantNode *node
	}{
		{
			a: []*ast.GenDecl{},
			b: []*ast.GenDecl{
				{},
			},
			want: -1,
			wantNode: newTestNode(
				"generic declaration lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 0 < 1", nil),
				),
			),
		},
		{
			a: []*ast.GenDecl{
				{},
			},
			b:    []*ast.GenDecl{},
			want: 1,
			wantNode: newTestNode(
				"generic declaration lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 1 > 0", nil),
				),
			),
		},
		{
			a: []*ast.GenDecl{
				{
					Tok: token.INT,
				},
				{
					Tok: token.FLOAT,
				},
			},
			b: []*ast.GenDecl{
				{
					Tok: token.FLOAT,
				},
				{
					Tok: token.INT,
				},
			},
			want: -1,
			wantNode: newNode(
				"generic declaration lists did not match",
				nil,
				nil,
				&[]*node{
					newTestNode(
						"generic declarations at index 0 did not match",
						newTestNode(
							"generic declarations did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
					newTestNode(
						"generic declarations at index 1 did not match",
						newTestNode(
							"generic declarations did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
				},
			),
		},
		{
			a: []*ast.GenDecl{
				{
					Tok: token.FLOAT,
				},
				{
					Tok: token.INT,
				},
			},
			b: []*ast.GenDecl{
				{
					Tok: token.INT,
				},
				{
					Tok: token.FLOAT,
				},
			},
			want: 1,
			wantNode: newNode(
				"generic declaration lists did not match",
				nil,
				nil,
				&[]*node{
					newTestNode(
						"generic declarations at index 0 did not match",
						newTestNode(
							"generic declarations did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 6 > 5", nil),
								),
							),
						),
					),
					newTestNode(
						"generic declarations at index 1 did not match",
						newTestNode(
							"generic declarations did not match",
							newTestNode(
								"tokens did not match",
								newTestNode(
									"tokens did not match",
									newTestNode("ints did not match: 5 < 6", nil),
								),
							),
						),
					),
				},
			),
		},
		{
			a: []*ast.GenDecl{
				{
					Tok: token.INT,
				},
				{
					Tok: token.FLOAT,
				},
			},
			b: []*ast.GenDecl{
				{
					Tok: token.INT,
				},
				{
					Tok: token.FLOAT,
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareGenDeclLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareGenDeclLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareFuncDeclLists(t *testing.T) {
	testCases := []struct {
		a        []*ast.FuncDecl
		b        []*ast.FuncDecl
		want     int
		wantNode *node
	}{
		{
			a: []*ast.FuncDecl{},
			b: []*ast.FuncDecl{
				{
					Type: &ast.FuncType{},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"function declaration lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 0 < 1", nil),
				),
			),
		},
		{
			a: []*ast.FuncDecl{
				{
					Type: &ast.FuncType{},
				},
			},
			b:    []*ast.FuncDecl{},
			want: 1,
			wantNode: newTestNode(
				"function declaration lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 1 > 0", nil),
				),
			),
		},
		{
			a: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
				{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			b: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
				{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
			},
			want: -1,
			wantNode: newNode(
				"function declaration lists did not match",
				nil,
				nil,
				&[]*node{
					newTestNode(
						"function declarations at index 0 did not match",
						newTestNode(
							"function declarations did not match",
							newTestNode(
								"names did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
					newTestNode(
						"function declarations at index 1 did not match",
						newTestNode(
							"function declarations did not match",
							newTestNode(
								"names did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: b > a", nil),
								),
							),
						),
					),
				},
			),
		},
		{
			a: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
				{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
			},
			b: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
				{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			want: 1,
			wantNode: newNode(
				"function declaration lists did not match",
				nil,
				nil,
				&[]*node{
					newTestNode(
						"function declarations at index 0 did not match",
						newTestNode(
							"function declarations did not match",
							newTestNode(
								"names did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: b > a", nil),
								),
							),
						),
					),
					newTestNode(
						"function declarations at index 1 did not match",
						newTestNode(
							"function declarations did not match",
							newTestNode(
								"names did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
				},
			),
		},
		{
			a: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
				{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			b: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
				{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareFuncDeclLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareFuncDeclLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareImportSpecLists(t *testing.T) {
	testCases := []struct {
		a        []*ast.ImportSpec
		b        []*ast.ImportSpec
		want     int
		wantNode *node
	}{
		{
			a: []*ast.ImportSpec{},
			b: []*ast.ImportSpec{
				{},
			},
			want: -1,
			wantNode: newTestNode(
				"import spec lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 0 < 1", nil),
				),
			),
		},
		{
			a: []*ast.ImportSpec{
				{},
			},
			b:    []*ast.ImportSpec{},
			want: 1,
			wantNode: newTestNode(
				"import spec lists did not match",
				newTestNode(
					"length of lists did not match",
					newTestNode("ints did not match: 1 > 0", nil),
				),
			),
		},
		{
			a: []*ast.ImportSpec{
				{
					Name: ast.NewIdent("a"),
				},
				{
					Name: ast.NewIdent("b"),
				},
			},
			b: []*ast.ImportSpec{
				{
					Name: ast.NewIdent("b"),
				},
				{
					Name: ast.NewIdent("a"),
				},
			},
			want: -1,
			wantNode: newNode(
				"import spec lists did not match",
				nil,
				nil,
				&[]*node{
					newTestNode(
						"import specs at index 0 did not match",
						newTestNode(
							"import specs did not match",
							newTestNode(
								"names did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
					newTestNode(
						"import specs at index 1 did not match",
						newTestNode(
							"import specs did not match",
							newTestNode(
								"names did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: b > a", nil),
								),
							),
						),
					),
				},
			),
		},
		{
			a: []*ast.ImportSpec{
				{
					Name: ast.NewIdent("b"),
				},
				{
					Name: ast.NewIdent("a"),
				},
			},
			b: []*ast.ImportSpec{
				{
					Name: ast.NewIdent("a"),
				},
				{
					Name: ast.NewIdent("b"),
				},
			},
			want: 1,
			wantNode: newNode(
				"import spec lists did not match",
				nil,
				nil,
				&[]*node{
					newTestNode(
						"import specs at index 0 did not match",
						newTestNode(
							"import specs did not match",
							newTestNode(
								"names did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: b > a", nil),
								),
							),
						),
					),
					newTestNode(
						"import specs at index 1 did not match",
						newTestNode(
							"import specs did not match",
							newTestNode(
								"names did not match",
								newTestNode(
									"identifiers did not match",
									newTestNode("strings did not match: a < b", nil),
								),
							),
						),
					),
				},
			),
		},
		{
			a: []*ast.ImportSpec{
				{
					Name: ast.NewIdent("a"),
				},
				{
					Name: ast.NewIdent("b"),
				},
			},
			b: []*ast.ImportSpec{
				{
					Name: ast.NewIdent("a"),
				},
				{
					Name: ast.NewIdent("b"),
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareImportSpecLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareImportSpecLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}

func TestCompareDeclLists(t *testing.T) {
	testCases := []struct {
		a        []ast.Decl
		b        []ast.Decl
		want     int
		wantNode *node
	}{
		{
			a: []ast.Decl{},
			b: []ast.Decl{
				&ast.FuncDecl{
					Type: &ast.FuncType{},
				},
			},
			want: -1,
			wantNode: newTestNode(
				"declaration lists did not match",
				newTestNode(
					"function declaration lists did not match",
					newTestNode(
						"function declaration lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 0 < 1", nil),
						),
					),
				),
			),
		},
		{
			a: []ast.Decl{
				&ast.FuncDecl{
					Type: &ast.FuncType{},
				},
			},
			b:    []ast.Decl{},
			want: 1,
			wantNode: newTestNode(
				"declaration lists did not match",
				newTestNode(
					"function declaration lists did not match",
					newTestNode(
						"function declaration lists did not match",
						newTestNode(
							"length of lists did not match",
							newTestNode("ints did not match: 1 > 0", nil),
						),
					),
				),
			),
		},
		{
			a: []ast.Decl{
				&ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
				&ast.FuncDecl{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			b: []ast.Decl{
				&ast.FuncDecl{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
				&ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
			},
			want: 0, // Order of declarations in list does not matter
		},
		{
			a: []ast.Decl{
				&ast.FuncDecl{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
				&ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
			},
			b: []ast.Decl{
				&ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
				&ast.FuncDecl{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			want: 0, // Order of declarations in list does not matter
		},
		{
			a: []ast.Decl{
				&ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
				&ast.FuncDecl{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			b: []ast.Decl{
				&ast.FuncDecl{
					Name: ast.NewIdent("a"),
					Type: &ast.FuncType{},
				},
				&ast.FuncDecl{
					Name: ast.NewIdent("b"),
					Type: &ast.FuncType{},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotNode := compareDeclLists(c.a, c.b)
		if got != c.want || !reflect.DeepEqual(gotNode, c.wantNode) {
			t.Errorf(
				"compareDeclLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotNode,
				c.want,
				c.wantNode,
			)
		}
	}
}
