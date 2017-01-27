package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	path := "C:\\Users\\Dave\\Documents\\Projects\\Go\\src\\github.com\\activelylazy\\generated-vending"
	_, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Error parsing: %v", err)
	}

	packageName := "vending"
	// fileName := "vending.go"
	f := &ast.File{}
	f.Name = ast.NewIdent(packageName)
	addImport(f, "", "github.com/moo")

	printer.Fprint(os.Stdout, fset, f)
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

func newBasicLit(value string) *ast.BasicLit {
	return &ast.BasicLit{
		Value: value,
	}
}
