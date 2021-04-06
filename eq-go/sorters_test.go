package eqgo

// func TestSortGenDeclList(t *testing.T) {
// 	testCases := []struct {
// 		x    []*ast.GenDecl
// 		want []*ast.GenDecl
// 	}{
// 		{
// 			x:    []*ast.GenDecl{},
// 			want: []*ast.GenDecl{},
// 		},
// 		{
// 			x: []*ast.GenDecl{
// 				{
// 					Tok: token.FLOAT,
// 					Specs: []ast.Spec{
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("b"),
// 						},
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("a"),
// 						},
// 					},
// 				},
// 				{
// 					Tok: token.INT,
// 					Specs: []ast.Spec{
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("a"),
// 						},
// 					},
// 				},
// 				{
// 					Tok: token.FLOAT,
// 					Specs: []ast.Spec{
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("a"),
// 						},
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("c"),
// 						},
// 					},
// 				},
// 				{
// 					Tok: token.FLOAT,
// 					Specs: []ast.Spec{
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("a"),
// 						},
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("b"),
// 						},
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("c"),
// 						},
// 					},
// 				},
// 			},
// 			want: []*ast.GenDecl{
// 				{
// 					Tok: token.INT,
// 					Specs: []ast.Spec{
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("a"),
// 						},
// 					},
// 				},
// 				{
// 					Tok: token.FLOAT,
// 					Specs: []ast.Spec{
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("a"),
// 						},
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("b"),
// 						},
// 					},
// 				},
// 				{
// 					Tok: token.FLOAT,
// 					Specs: []ast.Spec{
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("a"),
// 						},
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("c"),
// 						},
// 					},
// 				},
// 				{
// 					Tok: token.FLOAT,
// 					Specs: []ast.Spec{
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("a"),
// 						},
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("b"),
// 						},
// 						&ast.ImportSpec{
// 							Name: ast.NewIdent("c"),
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, c := range testCases {
// 		got := c.x
// 		sortGenDeclList(got)
// 		if cmp, msg := compareGenDeclLists(got, c.want); cmp != 0 {
// 			t.Errorf(
// 				"sortGenDeclList: got %v, want %v: %s",
// 				got,
// 				c.want,
// 				msg,
// 			)
// 		}
// 	}
// }

// func TestSortGenDecl(t *testing.T) {
// 	testCases := []struct {
// 		x    *ast.GenDecl
// 		want *ast.GenDecl
// 	}{
// 		{
// 			x:    &ast.GenDecl{},
// 			want: &ast.GenDecl{},
// 		},
// 		{
// 			x: &ast.GenDecl{
// 				Tok: token.INT,
// 				Specs: []ast.Spec{
// 					&ast.TypeSpec{
// 						Name: ast.NewIdent("b"),
// 					},
// 					&ast.TypeSpec{
// 						Name: ast.NewIdent("a"),
// 					},
// 					&ast.TypeSpec{
// 						Name: ast.NewIdent("c"),
// 					},
// 				},
// 			},
// 			want: &ast.GenDecl{
// 				Tok: token.INT,
// 				Specs: []ast.Spec{
// 					&ast.TypeSpec{
// 						Name: ast.NewIdent("a"),
// 					},
// 					&ast.TypeSpec{
// 						Name: ast.NewIdent("b"),
// 					},
// 					&ast.TypeSpec{
// 						Name: ast.NewIdent("c"),
// 					},
// 				},
// 			},
// 		},
// 		{
// 			x: &ast.GenDecl{
// 				Tok: token.INT,
// 				Specs: []ast.Spec{
// 					&ast.ImportSpec{
// 						Name: ast.NewIdent("b"),
// 					},
// 					&ast.ImportSpec{
// 						Name: ast.NewIdent("a"),
// 					},
// 					&ast.ImportSpec{
// 						Name: ast.NewIdent("c"),
// 					},
// 				},
// 			},
// 			want: &ast.GenDecl{
// 				Tok: token.INT,
// 				Specs: []ast.Spec{
// 					&ast.ImportSpec{
// 						Name: ast.NewIdent("a"),
// 					},
// 					&ast.ImportSpec{
// 						Name: ast.NewIdent("b"),
// 					},
// 					&ast.ImportSpec{
// 						Name: ast.NewIdent("c"),
// 					},
// 				},
// 			},
// 		},
// 		{
// 			x: &ast.GenDecl{
// 				Tok: token.INT,
// 				Specs: []ast.Spec{
// 					&ast.ValueSpec{
// 						Names: []*ast.Ident{
// 							ast.NewIdent("b"),
// 						},
// 					},
// 					&ast.ValueSpec{
// 						Names: []*ast.Ident{
// 							ast.NewIdent("a"),
// 						},
// 					},
// 					&ast.ValueSpec{
// 						Names: []*ast.Ident{
// 							ast.NewIdent("c"),
// 						},
// 					},
// 				},
// 			},
// 			want: &ast.GenDecl{
// 				Tok: token.INT,
// 				Specs: []ast.Spec{
// 					&ast.ValueSpec{
// 						Names: []*ast.Ident{
// 							ast.NewIdent("a"),
// 						},
// 					},
// 					&ast.ValueSpec{
// 						Names: []*ast.Ident{
// 							ast.NewIdent("b"),
// 						},
// 					},
// 					&ast.ValueSpec{
// 						Names: []*ast.Ident{
// 							ast.NewIdent("c"),
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, c := range testCases {
// 		got := c.x
// 		sortGenDecl(got)
// 		if cmp, msg := compareGenDecls(got, c.want); cmp != 0 {
// 			t.Errorf(
// 				"sortGenDecl: got %v, want %v: %s",
// 				got,
// 				c.want,
// 				msg,
// 			)
// 		}
// 	}
// }

// func TestSortFuncDeclLists(t *testing.T) {
// 	testCases := []struct {
// 		x    []*ast.FuncDecl
// 		want []*ast.FuncDecl
// 	}{
// 		{
// 			x:    []*ast.FuncDecl{},
// 			want: []*ast.FuncDecl{},
// 		},
// 		{
// 			x: []*ast.FuncDecl{
// 				{
// 					Name: ast.NewIdent("b"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 				{
// 					Name: ast.NewIdent("a"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 				{
// 					Name: ast.NewIdent("a"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 				{
// 					Name: ast.NewIdent("a"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 				{
// 					Name: ast.NewIdent("a"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 			},
// 			want: []*ast.FuncDecl{
// 				{
// 					Name: ast.NewIdent("a"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 				{
// 					Name: ast.NewIdent("a"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 				{
// 					Name: ast.NewIdent("a"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 				{
// 					Name: ast.NewIdent("a"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 				{
// 					Name: ast.NewIdent("b"),
// 					Type: &ast.FuncType{
// 						Params: &ast.FieldList{
// 							List: []*ast.Field{
// 								{},
// 							},
// 						},
// 					},
// 					Recv: &ast.FieldList{
// 						List: []*ast.Field{
// 							{},
// 						},
// 					},
// 					Body: &ast.BlockStmt{
// 						List: []ast.Stmt{
// 							&ast.AssignStmt{},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, c := range testCases {
// 		got := c.x
// 		sortFuncDeclList(got)
// 		if cmp, msg := compareFuncDeclLists(got, c.want); cmp != 0 {
// 			t.Errorf(
// 				"sortFuncDeclList: got %v, want %v: %s",
// 				got,
// 				c.want,
// 				msg,
// 			)
// 		}
// 	}
// }

// func TestSortSpecList(t *testing.T) {
// 	testCases := []struct {
// 		x    []ast.Spec
// 		want []ast.Spec
// 	}{
// 		{
// 			x:    []ast.Spec{},
// 			want: []ast.Spec{},
// 		},
// 		{
// 			x: []ast.Spec{
// 				&ast.ImportSpec{
// 					Name: ast.NewIdent("b"),
// 				},
// 				&ast.ImportSpec{
// 					Name: ast.NewIdent("a"),
// 				},
// 				&ast.ImportSpec{
// 					Name: ast.NewIdent("c"),
// 				},
// 			},
// 			want: []ast.Spec{
// 				&ast.ImportSpec{
// 					Name: ast.NewIdent("a"),
// 				},
// 				&ast.ImportSpec{
// 					Name: ast.NewIdent("b"),
// 				},
// 				&ast.ImportSpec{
// 					Name: ast.NewIdent("c"),
// 				},
// 			},
// 		},
// 		{
// 			x: []ast.Spec{
// 				&ast.TypeSpec{
// 					Name: ast.NewIdent("b"),
// 				},
// 				&ast.TypeSpec{
// 					Name: ast.NewIdent("a"),
// 				},
// 				&ast.TypeSpec{
// 					Name: ast.NewIdent("c"),
// 				},
// 			},
// 			want: []ast.Spec{
// 				&ast.TypeSpec{
// 					Name: ast.NewIdent("a"),
// 				},
// 				&ast.TypeSpec{
// 					Name: ast.NewIdent("b"),
// 				},
// 				&ast.TypeSpec{
// 					Name: ast.NewIdent("c"),
// 				},
// 			},
// 		},
// 		{
// 			x: []ast.Spec{
// 				&ast.ValueSpec{
// 					Names: []*ast.Ident{
// 						ast.NewIdent("b"),
// 					},
// 				},
// 				&ast.ValueSpec{
// 					Names: []*ast.Ident{
// 						ast.NewIdent("a"),
// 					},
// 				},
// 				&ast.ValueSpec{
// 					Names: []*ast.Ident{
// 						ast.NewIdent("c"),
// 					},
// 				},
// 			},
// 			want: []ast.Spec{
// 				&ast.ValueSpec{
// 					Names: []*ast.Ident{
// 						ast.NewIdent("a"),
// 					},
// 				},
// 				&ast.ValueSpec{
// 					Names: []*ast.Ident{
// 						ast.NewIdent("b"),
// 					},
// 				},
// 				&ast.ValueSpec{
// 					Names: []*ast.Ident{
// 						ast.NewIdent("c"),
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, c := range testCases {
// 		got := c.x
// 		sortSpecList(got)
// 		if cmp, msg := compareSpecLists(got, c.want); cmp != 0 {
// 			t.Errorf(
// 				"sortSpecList: got %v, want %v: %s",
// 				got,
// 				c.want,
// 				msg,
// 			)
// 		}
// 	}
// }
