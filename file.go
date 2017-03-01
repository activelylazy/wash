package wash

import (
	"go/ast"
	"go/token"

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
