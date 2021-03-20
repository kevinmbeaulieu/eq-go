package eqgo

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCompareInts(t *testing.T) {
	testCases := []struct {
		a       int
		b       int
		want    int
		wantMsg string
	}{
		{
			a:       1,
			b:       2,
			want:    -1,
			wantMsg: "1 < 2",
		},
		{
			a:       2,
			b:       1,
			want:    1,
			wantMsg: "2 > 1",
		},
		{
			a:    1,
			b:    1,
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareInts(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareInts(%d, %d) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareBools(t *testing.T) {
	testCases := []struct {
		a       bool
		b       bool
		want    int
		wantMsg string
	}{
		{
			a:       false,
			b:       true,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a:       true,
			b:       false,
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:    false,
			b:    false,
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareBools(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareBools(%t, %t) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareStrings(t *testing.T) {
	testCases := []struct {
		a       string
		b       string
		want    int
		wantMsg string
	}{
		{
			a:       "a",
			b:       "b",
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a:       "b",
			b:       "a",
			want:    1,
			wantMsg: "b > a",
		},
		{
			a:    "a",
			b:    "a",
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareStrings(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareStrings(%s, %s) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareIdentifiers(t *testing.T) {
	testCases := []struct {
		a       *ast.Ident
		b       *ast.Ident
		want    int
		wantMsg string
	}{
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:       nil,
			b:       &ast.Ident{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.Ident{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a:       &ast.Ident{Name: "a"},
			b:       &ast.Ident{Name: "b"},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a:       &ast.Ident{Name: "b"},
			b:       &ast.Ident{Name: "a"},
			want:    1,
			wantMsg: "b > a",
		},
		{
			a:    &ast.Ident{Name: "a"},
			b:    &ast.Ident{Name: "a"},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareIdentifiers(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareIdentifiers(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareImportSpecs(t *testing.T) {
	testCases := []struct {
		a       *ast.ImportSpec
		b       *ast.ImportSpec
		want    int
		wantMsg string
	}{
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:       nil,
			b:       &ast.ImportSpec{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.ImportSpec{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
			},
			b: &ast.ImportSpec{
				Name: &ast.Ident{Name: "b"},
			},
			want:    -1,
			wantMsg: "import names do not match: a < b",
		},
		{
			a: &ast.ImportSpec{
				Name: &ast.Ident{Name: "b"},
			},
			b: &ast.ImportSpec{
				Name: &ast.Ident{Name: "a"},
			},
			want:    1,
			wantMsg: "import names do not match: b > a",
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
			want:    -1,
			wantMsg: "import paths do not match: 5 < 6",
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
			want:    1,
			wantMsg: "import paths do not match: 6 > 5",
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
		got, gotMsg := compareImportSpecs(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareImportSpecs(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareBasicLiteratures(t *testing.T) {
	testCases := []struct {
		a       *ast.BasicLit
		b       *ast.BasicLit
		want    int
		wantMsg string
	}{
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:       nil,
			b:       &ast.BasicLit{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.BasicLit{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a:       &ast.BasicLit{Kind: token.INT},
			b:       &ast.BasicLit{Kind: token.FLOAT},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a:       &ast.BasicLit{Kind: token.FLOAT},
			b:       &ast.BasicLit{Kind: token.INT},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a:       &ast.BasicLit{Value: "a"},
			b:       &ast.BasicLit{Value: "b"},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a:       &ast.BasicLit{Value: "b"},
			b:       &ast.BasicLit{Value: "a"},
			want:    1,
			wantMsg: "b > a",
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareBasicLiteratures(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareBasicLiteratures(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareValueSpecs(t *testing.T) {
	testCases := []struct {
		a       *ast.ValueSpec
		b       *ast.ValueSpec
		want    int
		wantMsg string
	}{
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:       nil,
			b:       &ast.ValueSpec{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.ValueSpec{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "b"}},
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "b"}},
			},
			b: &ast.ValueSpec{
				Names: []*ast.Ident{{Name: "a"}},
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "aType < bType",
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
			want:    1,
			wantMsg: "bType > aType",
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
			want:    -1,
			wantMsg: "5 < 6",
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
			want:    1,
			wantMsg: "6 > 5",
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
		got, gotMsg := compareValueSpecs(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareValueSpecs(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareIdentifierLists(t *testing.T) {
	testCases := []struct {
		a       []*ast.Ident
		b       []*ast.Ident
		want    int
		wantMsg string
	}{
		{
			a: []*ast.Ident{
				{},
			},
			b: []*ast.Ident{
				{},
				{},
			},
			want:    -1,
			wantMsg: "1 < 2",
		},
		{
			a: []*ast.Ident{
				{},
				{},
			},
			b: []*ast.Ident{
				{},
			},
			want:    1,
			wantMsg: "2 > 1",
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
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: []*ast.Ident{
				{Name: "b"},
			},
			b: []*ast.Ident{
				{Name: "a"},
			},
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareIdentifierLists(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareIdentifierLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareExpressionLists(t *testing.T) {
	testCases := []struct {
		a       []ast.Expr
		b       []ast.Expr
		want    int
		wantMsg string
	}{
		{
			a: []ast.Expr{
				ast.NewIdent(""),
			},
			b: []ast.Expr{
				ast.NewIdent(""),
				ast.NewIdent(""),
			},
			want:    -1,
			wantMsg: "1 < 2",
		},
		{
			a: []ast.Expr{
				ast.NewIdent(""),
				ast.NewIdent(""),
			},
			b: []ast.Expr{
				ast.NewIdent(""),
			},
			want:    1,
			wantMsg: "2 > 1",
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
			want:    -1,
			wantMsg: "b < c",
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
			want:    1,
			wantMsg: "c > b",
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
		got, gotMsg := compareExpressionLists(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareExpressionLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareExpressions(t *testing.T) {
	testCases := []struct {
		a       ast.Expr
		b       ast.Expr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       ast.NewIdent(""),
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       ast.NewIdent(""),
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			a:       ast.NewIdent(""),
			b:       &ast.Ellipsis{},
			want:    -1,
			wantMsg: "1 < 2",
		},
		{
			a:       &ast.Ellipsis{},
			b:       ast.NewIdent(""),
			want:    1,
			wantMsg: "2 > 1",
		},
		{
			a:    ast.NewIdent(""),
			b:    ast.NewIdent(""),
			want: 0,
		},
		{
			a:       ast.NewIdent("a"),
			b:       ast.NewIdent("b"),
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "5 < 6",
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
			want:    -1,
			wantMsg: "function type parameters do not match: fields at index 0 do not match: field names do not match: a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "5 < 6",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "function type parameters do not match: fields at index 0 do not match: field names do not match: a < b",
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
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "1 < 2",
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareExpressions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareTypeSpecs(t *testing.T) {
	testCases := []struct {
		a       *ast.TypeSpec
		b       *ast.TypeSpec
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.TypeSpec{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.TypeSpec{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.TypeSpec{
				Name: ast.NewIdent("b"),
			},
			b: &ast.TypeSpec{
				Name: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareTypeSpecs(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareTypeSpecs(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareSpecs(t *testing.T) {
	testCases := []struct {
		a       ast.Spec
		b       ast.Spec
		want    int
		wantMsg string
	}{
		{
			a:       &ast.ImportSpec{},
			b:       &ast.ValueSpec{},
			want:    -1,
			wantMsg: "mismatched spec types",
		},
		{
			a:       &ast.ImportSpec{},
			b:       &ast.TypeSpec{},
			want:    -1,
			wantMsg: "mismatched spec types",
		},
		{
			a:       &ast.ValueSpec{},
			b:       &ast.ImportSpec{},
			want:    1,
			wantMsg: "mismatched spec types",
		},
		{
			a:       &ast.TypeSpec{},
			b:       &ast.ImportSpec{},
			want:    1,
			wantMsg: "mismatched spec types",
		},
		{
			a:       &ast.ValueSpec{},
			b:       &ast.TypeSpec{},
			want:    -1,
			wantMsg: "mismatched spec types",
		},
		{
			a:       &ast.TypeSpec{},
			b:       &ast.ValueSpec{},
			want:    1,
			wantMsg: "mismatched spec types",
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
			want:    -1,
			wantMsg: "import names do not match: a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareSpecs(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareSpecs(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareSpecLists(t *testing.T) {
	testCases := []struct {
		a       []ast.Spec
		b       []ast.Spec
		want    int
		wantMsg string
	}{
		{
			a: []ast.Spec{},
			b: []ast.Spec{
				&ast.ImportSpec{},
			},
			want:    -1,
			wantMsg: "0 < 1",
		},
		{
			a: []ast.Spec{
				&ast.ImportSpec{},
			},
			b:       []ast.Spec{},
			want:    1,
			wantMsg: "1 > 0",
		},
		{
			a: []ast.Spec{
				&ast.ImportSpec{},
			},
			b: []ast.Spec{
				&ast.ValueSpec{},
			},
			want:    -1,
			wantMsg: "mismatched spec types",
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
		got, gotMsg := compareSpecLists(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareSpecLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareEllipses(t *testing.T) {
	testCases := []struct {
		a       *ast.Ellipsis
		b       *ast.Ellipsis
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.Ellipsis{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.Ellipsis{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.Ellipsis{
				Elt: ast.NewIdent("a"),
			},
			b: &ast.Ellipsis{
				Elt: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.Ellipsis{
				Elt: ast.NewIdent("b"),
			},
			b: &ast.Ellipsis{
				Elt: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareEllipses(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareEllipses(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareFunctionLiteratures(t *testing.T) {
	testCases := []struct {
		a       *ast.FuncLit
		b       *ast.FuncLit
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.FuncLit{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.FuncLit{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "function type parameters do not match: fields at index 0 do not match: field names do not match: a < b",
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
			want:    1,
			wantMsg: "function type parameters do not match: fields at index 0 do not match: field names do not match: b > a",
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
							Tok: token.FLOAT,
						},
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
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
							Tok: token.INT,
						},
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
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
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareFunctionLiteratures(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareFunctionLiteratures(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareBlockStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.BlockStmt
		b       *ast.BlockStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.BlockStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.BlockStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{},
					&ast.BranchStmt{},
				},
			},
			want:    -1,
			wantMsg: "1 < 2",
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{},
					&ast.BranchStmt{},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{},
				},
			},
			want:    1,
			wantMsg: "2 > 1",
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Tok: token.INT,
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Tok: token.FLOAT,
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Tok: token.FLOAT,
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Tok: token.INT,
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Tok: token.INT,
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
					&ast.AssignStmt{
						Tok: token.INT,
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareBlockStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareBlockStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareStatements(t *testing.T) {
	testCases := []struct {
		a       ast.Stmt
		b       ast.Stmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.AssignStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.AssignStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a:       &ast.DeclStmt{},
			b:       &ast.EmptyStmt{},
			want:    -1,
			wantMsg: "1 < 2",
		},
		{
			a:       &ast.EmptyStmt{},
			b:       &ast.DeclStmt{},
			want:    1,
			wantMsg: "2 > 1",
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
				Decl: &ast.FuncDecl{},
			},
			want:    -1,
			wantMsg: "mismatched declaration types",
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.FuncDecl{},
			},
			b: &ast.DeclStmt{
				Decl: &ast.FuncDecl{},
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
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a:    &ast.ExprStmt{},
			b:    &ast.ExprStmt{},
			want: 0,
		},
		{
			a: &ast.SendStmt{
				Value: ast.NewIdent("a"),
			},
			b: &ast.SendStmt{
				Value: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a:    &ast.IncDecStmt{},
			b:    &ast.IncDecStmt{},
			want: 0,
		},
		{
			a: &ast.AssignStmt{
				Tok: token.INT,
			},
			b: &ast.AssignStmt{
				Tok: token.FLOAT,
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a:    &ast.AssignStmt{},
			b:    &ast.AssignStmt{},
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    -1,
			wantMsg: "5 < 6",
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
						Tok: token.INT,
					},
				},
			},
			b: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.FLOAT,
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a:    &ast.BlockStmt{},
			b:    &ast.BlockStmt{},
			want: 0,
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
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
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a:    &ast.CaseClause{},
			b:    &ast.CaseClause{},
			want: 0,
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a:    &ast.SwitchStmt{},
			b:    &ast.SwitchStmt{},
			want: 0,
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a:    &ast.TypeSwitchStmt{},
			b:    &ast.TypeSwitchStmt{},
			want: 0,
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
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
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a:    &ast.SelectStmt{},
			b:    &ast.SelectStmt{},
			want: 0,
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
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
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a:    &ast.RangeStmt{},
			b:    &ast.RangeStmt{},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareCompositeLiteratures(t *testing.T) {
	testCases := []struct {
		a       *ast.CompositeLit
		b       *ast.CompositeLit
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.CompositeLit{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.CompositeLit{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
			},
			b: &ast.CompositeLit{
				Type: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.CompositeLit{
				Type: ast.NewIdent("b"),
			},
			b: &ast.CompositeLit{
				Type: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "1 < 2",
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
			want:    1,
			wantMsg: "2 > 1",
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
			want:    -1,
			wantMsg: "false < true",
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
			want:    1,
			wantMsg: "true > false",
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
		got, gotMsg := compareCompositeLiteratures(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareCompositeLiteratures(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareParentheses(t *testing.T) {
	testCases := []struct {
		a       *ast.ParenExpr
		b       *ast.ParenExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.ParenExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.ParenExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.ParenExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.ParenExpr{
				X: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.ParenExpr{
				X: ast.NewIdent("b"),
			},
			b: &ast.ParenExpr{
				X: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareParentheses(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareParentheses(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareSelectors(t *testing.T) {
	testCases := []struct {
		a       *ast.SelectorExpr
		b       *ast.SelectorExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.SelectorExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.SelectorExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
			},
			b: &ast.SelectorExpr{
				Sel: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.SelectorExpr{
				Sel: ast.NewIdent("b"),
			},
			b: &ast.SelectorExpr{
				Sel: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareSelectors(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareSelectors(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareIndexExpressions(t *testing.T) {
	testCases := []struct {
		a       *ast.IndexExpr
		b       *ast.IndexExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.IndexExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.IndexExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
			},
			b: &ast.IndexExpr{
				Index: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.IndexExpr{
				Index: ast.NewIdent("b"),
			},
			b: &ast.IndexExpr{
				Index: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareIndexExpressions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareIndexExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareSliceExpressions(t *testing.T) {
	testCases := []struct {
		a       *ast.SliceExpr
		b       *ast.SliceExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.SliceExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.SliceExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.SliceExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.SliceExpr{
				X: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.SliceExpr{
				X: ast.NewIdent("b"),
			},
			b: &ast.SliceExpr{
				X: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
			want:    -1,
			wantMsg: "i < j",
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
			want:    1,
			wantMsg: "j > i",
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
			want:    -1,
			wantMsg: "m < n",
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
			want:    1,
			wantMsg: "n > m",
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
			want:    -1,
			wantMsg: "false < true",
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
			want:    1,
			wantMsg: "true > false",
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
		got, gotMsg := compareSliceExpressions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareSliceExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareTypeAssertions(t *testing.T) {
	testCases := []struct {
		a       *ast.TypeAssertExpr
		b       *ast.TypeAssertExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.TypeAssertExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.TypeAssertExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
			},
			b: &ast.TypeAssertExpr{
				Type: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.TypeAssertExpr{
				Type: ast.NewIdent("b"),
			},
			b: &ast.TypeAssertExpr{
				Type: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareTypeAssertions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareTypeAssertions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareCallExpressions(t *testing.T) {
	testCases := []struct {
		a       *ast.CallExpr
		b       *ast.CallExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.CallExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.CallExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
			},
			b: &ast.CallExpr{
				Fun: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.CallExpr{
				Fun: ast.NewIdent("b"),
			},
			b: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "0 < 1",
		},
		{
			a: &ast.CallExpr{
				Fun: ast.NewIdent("a"),
				Args: []ast.Expr{
					ast.NewIdent("x"),
				},
			},
			b: &ast.CallExpr{
				Fun:  ast.NewIdent("a"),
				Args: []ast.Expr{},
			},
			want:    1,
			wantMsg: "1 > 0",
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
		got, gotMsg := compareCallExpressions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareCallExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareStarExpressions(t *testing.T) {
	testCases := []struct {
		a       *ast.StarExpr
		b       *ast.StarExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.StarExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.StarExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.StarExpr{
				X: ast.NewIdent("a"),
			},
			b: &ast.StarExpr{
				X: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.StarExpr{
				X: ast.NewIdent("b"),
			},
			b: &ast.StarExpr{
				X: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareStarExpressions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareStarExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareUnaryExpressions(t *testing.T) {
	testCases := []struct {
		a       *ast.UnaryExpr
		b       *ast.UnaryExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.UnaryExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.UnaryExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.UnaryExpr{
				Op: token.INT,
			},
			b: &ast.UnaryExpr{
				Op: token.FLOAT,
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.UnaryExpr{
				Op: token.FLOAT,
			},
			b: &ast.UnaryExpr{
				Op: token.INT,
			},
			want:    1,
			wantMsg: "6 > 5",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareUnaryExpressions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareUnaryExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareBinaryExpressions(t *testing.T) {
	testCases := []struct {
		a       *ast.BinaryExpr
		b       *ast.BinaryExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.BinaryExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.BinaryExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.BinaryExpr{
				Op: token.INT,
			},
			b: &ast.BinaryExpr{
				Op: token.FLOAT,
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.BinaryExpr{
				Op: token.FLOAT,
			},
			b: &ast.BinaryExpr{
				Op: token.INT,
			},
			want:    1,
			wantMsg: "6 > 5",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareBinaryExpressions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareBinaryExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareKeyValueExpressions(t *testing.T) {
	testCases := []struct {
		a       *ast.KeyValueExpr
		b       *ast.KeyValueExpr
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.KeyValueExpr{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.KeyValueExpr{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.KeyValueExpr{
				Key: ast.NewIdent("a"),
			},
			b: &ast.KeyValueExpr{
				Key: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.KeyValueExpr{
				Key: ast.NewIdent("b"),
			},
			b: &ast.KeyValueExpr{
				Key: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareKeyValueExpressions(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareKeyValueExpressions(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareArrayTypes(t *testing.T) {
	testCases := []struct {
		a       *ast.ArrayType
		b       *ast.ArrayType
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.ArrayType{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.ArrayType{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.ArrayType{
				Len: ast.NewIdent("a"),
			},
			b: &ast.ArrayType{
				Len: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.ArrayType{
				Len: ast.NewIdent("b"),
			},
			b: &ast.ArrayType{
				Len: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareArrayTypes(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareArrayTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareStructTypes(t *testing.T) {
	testCases := []struct {
		a       *ast.StructType
		b       *ast.StructType
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.StructType{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.StructType{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.StructType{
				Incomplete: false,
			},
			b: &ast.StructType{
				Incomplete: true,
			},
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.StructType{
				Incomplete: true,
			},
			b: &ast.StructType{
				Incomplete: false,
			},
			want:    1,
			wantMsg: "true > false",
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
			want:    -1,
			wantMsg: "fields at index 0 do not match: field names do not match: a < b",
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
			want:    1,
			wantMsg: "fields at index 0 do not match: field names do not match: b > a",
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
		got, gotMsg := compareStructTypes(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareStructTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareFunctionTypes(t *testing.T) {
	testCases := []struct {
		a       *ast.FuncType
		b       *ast.FuncType
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.FuncType{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.FuncType{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "function type parameters do not match: fields at index 0 do not match: field names do not match: a < b",
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
			want:    1,
			wantMsg: "function type parameters do not match: fields at index 0 do not match: field names do not match: b > a",
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
			want:    -1,
			wantMsg: "function type results do not match: fields at index 0 do not match: field names do not match: x < y",
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
			want:    1,
			wantMsg: "function type results do not match: fields at index 0 do not match: field names do not match: y > x",
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
		got, gotMsg := compareFunctionTypes(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareFunctionTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareFieldLists(t *testing.T) {
	testCases := []struct {
		a       *ast.FieldList
		b       *ast.FieldList
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.FieldList{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.FieldList{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.FieldList{
				List: []*ast.Field{
					{},
				},
			},
			b: &ast.FieldList{
				List: []*ast.Field{
					{},
					{},
				},
			},
			want:    -1,
			wantMsg: "length of field lists do not match: 1 < 2",
		},
		{
			a: &ast.FieldList{
				List: []*ast.Field{
					{},
					{},
				},
			},
			b: &ast.FieldList{
				List: []*ast.Field{
					{},
				},
			},
			want:    1,
			wantMsg: "length of field lists do not match: 2 > 1",
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
			want:    -1,
			wantMsg: "fields at index 0 do not match: field names do not match: a < b",
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
			want:    1,
			wantMsg: "fields at index 0 do not match: field names do not match: b > a",
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
		got, gotMsg := compareFieldLists(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareFieldLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareFields(t *testing.T) {
	testCases := []struct {
		a       *ast.Field
		b       *ast.Field
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.Field{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.Field{},
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "field names do not match: a < b",
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
			want:    1,
			wantMsg: "field names do not match: b > a",
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
			want:    -1,
			wantMsg: "field types do not match: x < y",
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
			want:    1,
			wantMsg: "field types do not match: y > x",
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
			want:    -1,
			wantMsg: "field tags do not match: 5 < 6",
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
			want:    1,
			wantMsg: "field tags do not match: 6 > 5",
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
		got, gotMsg := compareFields(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareFields(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareInterfaceTypes(t *testing.T) {
	testCases := []struct {
		a       *ast.InterfaceType
		b       *ast.InterfaceType
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.InterfaceType{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.InterfaceType{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "fields at index 0 do not match: field names do not match: a < b",
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
			want:    1,
			wantMsg: "fields at index 0 do not match: field names do not match: b > a",
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
			want:    -1,
			wantMsg: "false < true",
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
			want:    1,
			wantMsg: "true > false",
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
		got, gotMsg := compareInterfaceTypes(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareInterfaceTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareMapTypes(t *testing.T) {
	testCases := []struct {
		a       *ast.MapType
		b       *ast.MapType
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.MapType{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.MapType{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.MapType{
				Key: ast.NewIdent("a"),
			},
			b: &ast.MapType{
				Key: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.MapType{
				Key: ast.NewIdent("b"),
			},
			b: &ast.MapType{
				Key: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareMapTypes(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareMapTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareChannelTypes(t *testing.T) {
	testCases := []struct {
		a       *ast.ChanType
		b       *ast.ChanType
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.ChanType{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.ChanType{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.ChanType{
				Dir: ast.SEND,
			},
			b: &ast.ChanType{
				Dir: ast.RECV,
			},
			want:    -1,
			wantMsg: "1 < 2",
		},
		{
			a: &ast.ChanType{
				Dir: ast.RECV,
			},
			b: &ast.ChanType{
				Dir: ast.SEND,
			},
			want:    1,
			wantMsg: "2 > 1",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareChannelTypes(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareChannelTypes(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareChannelDirections(t *testing.T) {
	testCases := []struct {
		a       ast.ChanDir
		b       ast.ChanDir
		want    int
		wantMsg string
	}{
		{
			a:       ast.SEND,
			b:       ast.RECV,
			want:    -1,
			wantMsg: "1 < 2",
		},
		{
			a:       ast.RECV,
			b:       ast.SEND,
			want:    1,
			wantMsg: "2 > 1",
		},
		{
			a:    ast.RECV,
			b:    ast.RECV,
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareChannelDirections(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareChannelDirections(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareBadStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.BadStmt
		b       *ast.BadStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.BadStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.BadStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a:    &ast.BadStmt{},
			b:    &ast.BadStmt{},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareBadStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareBadStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareDeclStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.DeclStmt
		b       *ast.DeclStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.DeclStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.DeclStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("a"),
				},
			},
			b: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("b"),
				},
			},
			want:    -1,
			wantMsg: "function names do not match: a < b",
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("b"),
				},
			},
			b: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("a"),
				},
			},
			want:    1,
			wantMsg: "function names do not match: b > a",
		},
		{
			a: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("a"),
				},
			},
			b: &ast.DeclStmt{
				Decl: &ast.FuncDecl{
					Name: ast.NewIdent("a"),
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareDeclStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareDeclStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareDecls(t *testing.T) {
	testCases := []struct {
		a       ast.Decl
		b       ast.Decl
		want    int
		wantMsg string
	}{
		{
			a:       &ast.BadDecl{},
			b:       &ast.GenDecl{},
			want:    -1,
			wantMsg: "mismatched declaration types",
		},
		{
			a:       &ast.BadDecl{},
			b:       &ast.FuncDecl{},
			want:    -1,
			wantMsg: "mismatched declaration types",
		},
		{
			a:       &ast.GenDecl{},
			b:       &ast.FuncDecl{},
			want:    -1,
			wantMsg: "mismatched declaration types",
		},
		{
			a:       &ast.GenDecl{},
			b:       &ast.BadDecl{},
			want:    1,
			wantMsg: "mismatched declaration types",
		},
		{
			a:       &ast.FuncDecl{},
			b:       &ast.BadDecl{},
			want:    1,
			wantMsg: "mismatched declaration types",
		},
		{
			a:       &ast.FuncDecl{},
			b:       &ast.GenDecl{},
			want:    1,
			wantMsg: "mismatched declaration types",
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
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.GenDecl{
				Tok: token.FLOAT,
			},
			b: &ast.GenDecl{
				Tok: token.INT,
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a:    &ast.GenDecl{},
			b:    &ast.GenDecl{},
			want: 0,
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "function names do not match: a < b",
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("b"),
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "function names do not match: b > a",
		},
		{
			a:    &ast.FuncDecl{},
			b:    &ast.FuncDecl{},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareDecls(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareDecls(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareBadDecls(t *testing.T) {
	testCases := []struct {
		a       *ast.BadDecl
		b       *ast.BadDecl
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.BadDecl{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.BadDecl{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a:    &ast.BadDecl{},
			b:    &ast.BadDecl{},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareBadDecls(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareBadDecls(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareGenDecls(t *testing.T) {
	testCases := []struct {
		a       *ast.GenDecl
		b       *ast.GenDecl
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.GenDecl{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.GenDecl{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.GenDecl{
				Tok: token.INT,
			},
			b: &ast.GenDecl{
				Tok: token.FLOAT,
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.GenDecl{
				Tok: token.FLOAT,
			},
			b: &ast.GenDecl{
				Tok: token.INT,
			},
			want:    1,
			wantMsg: "6 > 5",
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
			want:    -1,
			wantMsg: "1 < 2",
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
			want:    1,
			wantMsg: "2 > 1",
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
		got, gotMsg := compareGenDecls(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareGenDecls(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareFuncDecls(t *testing.T) {
	testCases := []struct {
		a       *ast.FuncDecl
		b       *ast.FuncDecl
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.FuncDecl{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.FuncDecl{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "function names do not match: a < b",
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("b"),
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "function names do not match: b > a",
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: nil,
			},
			want:    -1,
			wantMsg: "function receivers do not match: false < true",
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: nil,
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
			},
			want:    1,
			wantMsg: "function receivers do not match: true > false",
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: nil,
			},
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: nil,
			},
			b: &ast.FuncDecl{
				Name: ast.NewIdent("a"),
				Recv: &ast.FieldList{},
				Type: &ast.FuncType{},
			},
			want:    1,
			wantMsg: "true > false",
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
			want:    -1,
			wantMsg: "function bodies do not match: false < true",
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
			want:    1,
			wantMsg: "function bodies do not match: true > false",
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
		got, gotMsg := compareFuncDecls(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareFuncDecls(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareEmptyStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.EmptyStmt
		b       *ast.EmptyStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.EmptyStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.EmptyStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.EmptyStmt{
				Implicit: false,
			},
			b: &ast.EmptyStmt{
				Implicit: true,
			},
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.EmptyStmt{
				Implicit: true,
			},
			b: &ast.EmptyStmt{
				Implicit: false,
			},
			want:    1,
			wantMsg: "true > false",
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
		got, gotMsg := compareEmptyStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareEmptyStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareLabeledStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.LabeledStmt
		b       *ast.LabeledStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.LabeledStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.LabeledStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("b"),
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.LabeledStmt{
				Label: ast.NewIdent("a"),
				Stmt: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareLabeledStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareLabeledStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareExpressionStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.ExprStmt
		b       *ast.ExprStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.ExprStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.ExprStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.ExprStmt{
				X: ast.NewIdent("a"),
			},
			b: &ast.ExprStmt{
				X: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.ExprStmt{
				X: ast.NewIdent("b"),
			},
			b: &ast.ExprStmt{
				X: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareExpressionStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareExpressionStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareSendStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.SendStmt
		b       *ast.SendStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.SendStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.SendStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.SendStmt{
				Chan: ast.NewIdent("a"),
			},
			b: &ast.SendStmt{
				Chan: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.SendStmt{
				Chan: ast.NewIdent("b"),
			},
			b: &ast.SendStmt{
				Chan: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
		got, gotMsg := compareSendStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareSendStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareIncDecStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.IncDecStmt
		b       *ast.IncDecStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.IncDecStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.IncDecStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.IncDecStmt{
				Tok: token.INT,
			},
			b: &ast.IncDecStmt{
				Tok: token.FLOAT,
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.IncDecStmt{
				Tok: token.FLOAT,
			},
			b: &ast.IncDecStmt{
				Tok: token.INT,
			},
			want:    1,
			wantMsg: "6 > 5",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareIncDecStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareIncDecStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareAssignStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.AssignStmt
		b       *ast.AssignStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.AssignStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.AssignStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.AssignStmt{
				Tok: token.INT,
			},
			b: &ast.AssignStmt{
				Tok: token.FLOAT,
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.AssignStmt{
				Tok: token.FLOAT,
			},
			b: &ast.AssignStmt{
				Tok: token.INT,
			},
			want:    1,
			wantMsg: "6 > 5",
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
			want:    -1,
			wantMsg: "1 < 2",
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
			want:    1,
			wantMsg: "2 > 1",
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
			want:    -1,
			wantMsg: "1 < 2",
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
			want:    1,
			wantMsg: "2 > 1",
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
		got, gotMsg := compareAssignStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareAssignStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareGoStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.GoStmt
		b       *ast.GoStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.GoStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.GoStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareGoStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareGoStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareDeferStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.DeferStmt
		b       *ast.DeferStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.DeferStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.DeferStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareDeferStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareDeferStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareReturnStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.ReturnStmt
		b       *ast.ReturnStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.ReturnStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.ReturnStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareReturnStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareReturnStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareBranchStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.BranchStmt
		b       *ast.BranchStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.BranchStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.BranchStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.BranchStmt{
				Tok: token.INT,
			},
			b: &ast.BranchStmt{
				Tok: token.FLOAT,
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.BranchStmt{
				Tok: token.FLOAT,
			},
			b: &ast.BranchStmt{
				Tok: token.INT,
			},
			want:    1,
			wantMsg: "6 > 5",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    1,
			wantMsg: "b > a",
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
		got, gotMsg := compareBranchStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareBranchStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareIfStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.IfStmt
		b       *ast.IfStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.IfStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.IfStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("b"),
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
				Else: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareIfStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareIfStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareCaseClauses(t *testing.T) {
	testCases := []struct {
		a       *ast.CaseClause
		b       *ast.CaseClause
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.CaseClause{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.CaseClause{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
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
			want:    -1,
			wantMsg: "a < b",
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
			want:    1,
			wantMsg: "b > a",
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
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
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.FLOAT,
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
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.CaseClause{
				List: []ast.Expr{
					ast.NewIdent("a"),
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
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
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareCaseClauses(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareCaseClauses(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareStatementLists(t *testing.T) {
	testCases := []struct {
		a       []ast.Stmt
		b       []ast.Stmt
		want    int
		wantMsg string
	}{
		{
			a: []ast.Stmt{},
			b: []ast.Stmt{
				&ast.AssignStmt{},
			},
			want:    -1,
			wantMsg: "0 < 1",
		},
		{
			a: []ast.Stmt{
				&ast.AssignStmt{},
			},
			b:       []ast.Stmt{},
			want:    1,
			wantMsg: "1 > 0",
		},
		{
			a: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.INT,
				},
				&ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.FLOAT,
				},
				&ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.FLOAT,
				},
				&ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.INT,
				},
				&ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.INT,
				},
				&ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.INT,
				},
				&ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareStatementLists(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareStatementLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareSwitchStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.SwitchStmt
		b       *ast.SwitchStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.SwitchStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.SwitchStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("a"),
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("b"),
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.SwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Tag: ast.NewIdent("a"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareSwitchStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareSwitchStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareTypeSwitchStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.TypeSwitchStmt
		b       *ast.TypeSwitchStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.TypeSwitchStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.TypeSwitchStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.TypeSwitchStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Assign: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareTypeSwitchStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareTypeSwitchStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareCommClauses(t *testing.T) {
	testCases := []struct {
		a       *ast.CommClause
		b       *ast.CommClause
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.CommClause{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.CommClause{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.FLOAT,
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.FLOAT,
					},
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
				},
			},
			b: &ast.CommClause{
				Comm: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: []ast.Stmt{
					&ast.AssignStmt{
						Tok: token.INT,
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareCommClauses(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareCommClauses(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareSelectStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.SelectStmt
		b       *ast.SelectStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.SelectStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.SelectStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			b: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.SelectStmt{
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareSelectStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareSelectStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareForStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.ForStmt
		b       *ast.ForStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.ForStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.ForStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("b"),
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.FLOAT,
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.FLOAT,
						},
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
		},
		{
			a: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			b: &ast.ForStmt{
				Init: &ast.AssignStmt{
					Tok: token.INT,
				},
				Cond: ast.NewIdent("a"),
				Post: &ast.AssignStmt{
					Tok: token.INT,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.INT,
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareForStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareForStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareRangeStatements(t *testing.T) {
	testCases := []struct {
		a       *ast.RangeStmt
		b       *ast.RangeStmt
		want    int
		wantMsg string
	}{
		{
			a:       nil,
			b:       &ast.RangeStmt{},
			want:    1,
			wantMsg: "true > false",
		},
		{
			a:       &ast.RangeStmt{},
			b:       nil,
			want:    -1,
			wantMsg: "false < true",
		},
		{
			a: &ast.RangeStmt{
				Key: ast.NewIdent("a"),
			},
			b: &ast.RangeStmt{
				Key: ast.NewIdent("b"),
			},
			want:    -1,
			wantMsg: "a < b",
		},
		{
			a: &ast.RangeStmt{
				Key: ast.NewIdent("b"),
			},
			b: &ast.RangeStmt{
				Key: ast.NewIdent("a"),
			},
			want:    1,
			wantMsg: "b > a",
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
			want:    -1,
			wantMsg: "x < y",
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
			want:    1,
			wantMsg: "y > x",
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
			want:    -1,
			wantMsg: "5 < 6",
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
			want:    1,
			wantMsg: "6 > 5",
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
			want:    -1,
			wantMsg: "m < n",
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
			want:    1,
			wantMsg: "n > m",
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
						},
					},
				},
			},
			want:    -1,
			wantMsg: "5 < 6",
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
						},
					},
				},
			},
			want:    1,
			wantMsg: "6 > 5",
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
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, c := range testCases {
		got, gotMsg := compareRangeStatements(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareRangeStatements(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareGenDeclLists(t *testing.T) {
	testCases := []struct {
		a       []*ast.GenDecl
		b       []*ast.GenDecl
		want    int
		wantMsg string
	}{
		{
			a: []*ast.GenDecl{},
			b: []*ast.GenDecl{
				{},
			},
			want:    -1,
			wantMsg: "0 < 1",
		},
		{
			a: []*ast.GenDecl{
				{},
			},
			b:       []*ast.GenDecl{},
			want:    1,
			wantMsg: "1 > 0",
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
			want:    -1,
			wantMsg: "5 < 6",
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
			want:    1,
			wantMsg: "6 > 5",
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
		got, gotMsg := compareGenDeclLists(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareGenDeclLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareFuncDeclLists(t *testing.T) {
	testCases := []struct {
		a       []*ast.FuncDecl
		b       []*ast.FuncDecl
		want    int
		wantMsg string
	}{
		{
			a: []*ast.FuncDecl{},
			b: []*ast.FuncDecl{
				{},
			},
			want:    -1,
			wantMsg: "length of function declaration lists do not match: 0 < 1",
		},
		{
			a: []*ast.FuncDecl{
				{},
			},
			b:       []*ast.FuncDecl{},
			want:    1,
			wantMsg: "length of function declaration lists do not match: 1 > 0",
		},
		{
			a: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("a"),
				},
				{
					Name: ast.NewIdent("b"),
				},
			},
			b: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("b"),
				},
				{
					Name: ast.NewIdent("a"),
				},
			},
			want:    -1,
			wantMsg: "function declarations at index 0 do not match: function names do not match: a < b",
		},
		{
			a: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("b"),
				},
				{
					Name: ast.NewIdent("a"),
				},
			},
			b: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("a"),
				},
				{
					Name: ast.NewIdent("b"),
				},
			},
			want:    1,
			wantMsg: "function declarations at index 0 do not match: function names do not match: b > a",
		},
		{
			a: []*ast.FuncDecl{
				{
					Name: ast.NewIdent("a"),
				},
				{
					Name: ast.NewIdent("b"),
				},
			},
			b: []*ast.FuncDecl{
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
		got, gotMsg := compareFuncDeclLists(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareFuncDeclLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}

func TestCompareImportSpecLists(t *testing.T) {
	testCases := []struct {
		a       []*ast.ImportSpec
		b       []*ast.ImportSpec
		want    int
		wantMsg string
	}{
		{
			a: []*ast.ImportSpec{},
			b: []*ast.ImportSpec{
				{},
			},
			want:    -1,
			wantMsg: "length of import lists do not match: 0 < 1",
		},
		{
			a: []*ast.ImportSpec{
				{},
			},
			b:       []*ast.ImportSpec{},
			want:    1,
			wantMsg: "length of import lists do not match: 1 > 0",
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
			want:    -1,
			wantMsg: "imports at index 0 do not match: import names do not match: a < b",
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
			want:    1,
			wantMsg: "imports at index 0 do not match: import names do not match: b > a",
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
		got, gotMsg := compareImportSpecLists(c.a, c.b)
		if got != c.want || gotMsg != c.wantMsg {
			t.Errorf(
				"compareImportSpecLists(%v, %v) == (%d, %s), want (%d, %s)",
				c.a,
				c.b,
				got,
				gotMsg,
				c.want,
				c.wantMsg,
			)
		}
	}
}
