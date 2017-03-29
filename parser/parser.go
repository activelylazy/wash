package parser

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
)

// ParseStatement parses inline go code and returns the single statement it contains
func ParseStatement(s string) (ast.Stmt, error) {
	fset := token.NewFileSet()
	src := "package main\n\nfunc __here() {\n" + s + "\n}"
	f, err := parser.ParseFile(fset, "inline.go", src, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for _, d := range f.Decls {
		switch d.(type) {
		case *ast.FuncDecl:
			fd := d.(*ast.FuncDecl)
			if len(fd.Body.List) == 0 {
				return nil, errors.New("Inline go code did not result in a statement")
			}
			if len(fd.Body.List) > 1 {
				return nil, errors.New("Inline go code resulted in multiple statements")
			}
			return fd.Body.List[0], nil
		}
	}

	return nil, nil
}
