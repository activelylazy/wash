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
	existing := existingImportSpecs(f.File)
	for _, s := range existing {
		if nilSafeToString(s.Name) == name && s.Path.Value == "\""+path+"\"" {
			return
		}
	}
	addImport(f.File, name, path)
}

// AddFunction adds a new function to this file
func (f File) AddFunction(functionName string, params []syntax.Field, returnValues []DomainConcept) Function {
	log.Printf("Adding function %s to %s", functionName, f.TargetFilename)
	statements := []ast.Stmt{}
	if len(returnValues) > 0 {
		statements = append(statements, newReturnStmt(returnValues))
	}
	returnFields := conceptsToFields(returnValues)
	decl := addFunction(f.File, functionName, params, returnFields, statements)
	f.Write()
	return Function{
		File:         f,
		FunctionName: functionName,
		Decl:         decl,
		Params:       params,
		ReturnValues: returnFields,
	}
}

func conceptsToFields(concepts []DomainConcept) []syntax.Field {
	results := make([]syntax.Field, len(concepts))
	for i, c := range concepts {
		results[i] = syntax.NewField(c.name, c.typeName)
	}
	return results
}

// AddStruct adds a new struct to this file
func (f File) AddStruct(file *File, structName string, fieldDeclarations []syntax.Field) Struct {
	log.Printf("Adding struct %s to %s", structName, f.TargetFilename)
	decl := addStruct(f.File, structName, fieldDeclarations)
	f.Write()
	return Struct{
		File:       f,
		structName: structName,
		Decl:       decl,
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

func existingImportSpecs(f *ast.File) []*ast.ImportSpec {
	specs := make([]*ast.ImportSpec, 0)
	for _, d := range f.Decls {
		switch d.(type) {
		case *ast.GenDecl:
			g := d.(*ast.GenDecl)
			if g.Tok == token.IMPORT {
				for _, s := range g.Specs {
					switch s.(type) {
					case *ast.ImportSpec:
						i := s.(*ast.ImportSpec)
						specs = append(specs, i)
					}
				}
			}
		}
	}

	return specs
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

func addStruct(f *ast.File, structName string, fieldDeclarations []syntax.Field) *ast.GenDecl {
	newDecl := &ast.GenDecl{
		Tok: token.STRUCT,
		Specs: []ast.Spec{
			newTypeSpec(structName, fieldDeclarations),
		},
	}
	f.Decls = append([]ast.Decl{newDecl}, f.Decls...)
	return newDecl
}

func newTypeSpec(structName string, fieldDeclarations []syntax.Field) *ast.TypeSpec {
	return &ast.TypeSpec{
		Name: syntax.NewIdent(structName),
		Type: newStructType(fieldDeclarations),
	}
}

func newStructType(fieldDeclarations []syntax.Field) *ast.StructType {
	return &ast.StructType{
		Fields: syntax.NewFieldList(fieldDeclarations),
	}
}

func newReturnStmt(returnValues []DomainConcept) *ast.ReturnStmt {
	results := []ast.Expr{}
	for _, s := range returnValues {
		results = append(results, syntax.NewBasicLit(s.value))
	}
	return &ast.ReturnStmt{
		Results: results,
	}
}

func nilSafeToString(i *ast.Ident) string {
	if i == nil {
		return ""
	}
	return i.String()
}
