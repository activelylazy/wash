package wash

import (
	"go/ast"
	"go/token"
)

type field struct {
	fieldName string
	typeName  string
}

func newFile(packageName string) *ast.File {
	f := &ast.File{}
	f.Name = ast.NewIdent(packageName)
	return f
}

func addFunction(f *ast.File, name string, params []field, results []field, statementList []ast.Stmt) {
	newDecl := &ast.FuncDecl{
		Name: newIdent(name),
		Type: newFuncType(params, results),
		Body: &ast.BlockStmt{
			List: statementList,
		},
	}
	f.Decls = append(f.Decls, newDecl)
}

func newFuncType(params []field, results []field) *ast.FuncType {
	return &ast.FuncType{
		Params:  newFieldList(params),
		Results: newFieldList(results),
	}
}

func addImport(f *ast.File, name string, path string) {
	newDecl := &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			newImportSpec(name, path),
		},
	}
	f.Decls = append([]ast.Decl{newDecl}, f.Decls...)
}

func newImportSpec(name string, path string) *ast.ImportSpec {
	return &ast.ImportSpec{
		Name: newIdent(name),
		Path: newBasicLit(path),
	}
}

func newIdent(name string) *ast.Ident {
	if name == "" {
		return nil
	}
	return &ast.Ident{
		Name: name,
	}
}

func newIdentList(name string) []*ast.Ident {
	if name == "" {
		return make([]*ast.Ident, 0)
	}
	return []*ast.Ident{newIdent(name)}
}

func newBasicLit(value string) *ast.BasicLit {
	return &ast.BasicLit{
		Value: value,
	}
}

func newFieldList(fields []field) *ast.FieldList {
	l := &ast.FieldList{}
	l.List = make([]*ast.Field, len(fields))
	for i, p := range fields {
		l.List[i] = &ast.Field{
			Names: newIdentList(p.fieldName),
			Type:  newBasicLit(p.typeName),
		}
	}
	return l
}
