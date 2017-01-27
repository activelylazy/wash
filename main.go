package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

type field struct {
	fieldName string
	typeName  string
}

func main() {
	fset := token.NewFileSet()
	path := "C:\\Users\\Dave\\Documents\\Projects\\Go\\src\\github.com\\activelylazy\\generated-vending"
	_, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Error parsing: %v", err)
	}

	packageName := "vending"
	// fileName := "vending.go"
	f := newFile(packageName)
	addImport(f, "", "\"github.com/moo\"")
	addFunction(f, "validateCoin", []field{
		field{
			fieldName: "s",
			typeName:  "string",
		}},
		[]field{
			field{
				fieldName: "",
				typeName:  "int",
			},
			field{
				fieldName: "",
				typeName:  "bool",
			},
		},
		[]ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{
					newBasicLit("0"),
					newBasicLit("false"),
				},
			},
		})

	printer.Fprint(os.Stdout, fset, f)
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
