package wash

import (
	"go/ast"
	"go/token"
	"log"

	"github.com/activelylazy/wash/syntax"
)

// File represents a file managed by washer
type File struct {
	TargetFilename string
	File           *ast.File
	washer         *Washer
}

// AddImport adds a new import to this file
func (f File) AddImport(name string, path string) {
	addImport(f.File, name, path)
}

// AddFunction adds a new function to this file
func (f File) AddFunction(functionName string, params []syntax.Field, returnFields []syntax.Field, returnValues []string) Function {
	log.Printf("Adding function %s to %s", functionName, f.TargetFilename)
	statements := []ast.Stmt{}
	if len(returnValues) > 0 {
		statements = append(statements, newReturnStmt(returnValues))
	}
	decl := addFunction(f.File, functionName, params, returnFields, statements)
	f.Write()
	return Function{
		File:         f,
		functionName: functionName,
		Decl:         decl,
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
		Name: syntax.NewIdent(name),
		Path: syntax.NewBasicLit("\"" + path + "\""),
	}
}

func addFunction(f *ast.File, name string, params []syntax.Field, results []syntax.Field, statementList []ast.Stmt) *ast.FuncDecl {
	newDecl := &ast.FuncDecl{
		Name: syntax.NewIdent(name),
		Type: newFuncType(params, results),
		Body: &ast.BlockStmt{
			List: statementList,
		},
	}
	f.Decls = append(f.Decls, newDecl)
	return newDecl
}

func newFuncType(params []syntax.Field, results []syntax.Field) *ast.FuncType {
	return &ast.FuncType{
		Params:  syntax.NewFieldList(params),
		Results: syntax.NewFieldList(results),
	}
}

func newReturnStmt(returnValues []string) *ast.ReturnStmt {
	results := []ast.Expr{}
	for _, s := range returnValues {
		results = append(results, syntax.NewBasicLit(s))
	}
	return &ast.ReturnStmt{
		Results: results,
	}
}
