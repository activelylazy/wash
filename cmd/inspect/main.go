package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	fset := token.NewFileSet()
	path := "C:\\Users\\Dave\\Documents\\Projects\\Go\\src\\github.com\\activelylazy\\vending-kata"
	pkgs, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Error parsing: %v", err)
	}

	packageName := "vending"
	fileName := "vending_test.go"
	f := findFile(pkgs, packageName, fileName)

	if f == nil {
		log.Fatalf("Failed to find file %s/%s", packageName, fileName)
	}

	for _, s := range f.Decls {
		switch v := s.(type) {
		default:
			fmt.Printf("Read %T\n", v)
		case *ast.GenDecl:
			g := s.(*ast.GenDecl)
			fmt.Printf("Gen decl of type %s\n", g.Tok)
			if g.Tok == token.IMPORT {
				for _, spec := range g.Specs {
					is := spec.(*ast.ImportSpec)
					fmt.Printf("..Import spec name=%s, path=%s\n", is.Name, is.Path.Value)
				}
			}
		case *ast.FuncDecl:
			f := s.(*ast.FuncDecl)
			fmt.Printf("Read function %s\n", f.Name.Name)
			for _, stmt := range f.Body.List {
				fmt.Printf("..Statement: %s\n", reflect.TypeOf(stmt))
			}
		}
	}

	printer.Fprint(os.Stdout, fset, f)
}

func findFile(pkgs map[string]*ast.Package, packageName string, fileName string) *ast.File {
	for k, pkg := range pkgs {
		if k != packageName {
			continue
		}
		for n, f := range pkg.Files {
			if strings.HasSuffix(n, fileName) {
				return f
			}
		}
	}
	return nil
}
