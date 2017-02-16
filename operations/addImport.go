package operations

import (
	"go/ast"
	"go/token"

	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/syntax"
)

// AddImportRequest represents a request to add an import
type AddImportRequest struct {
	file *wash.File
	name string
	path string
}

// NewAddImportRequest creates a new request to add an import
func NewAddImportRequest(file *wash.File, name string, path string) AddImportRequest {
	return AddImportRequest{
		file: file,
		name: name,
		path: path,
	}
}

// Apply adds the import to the file
func (r AddImportRequest) Apply(washer *wash.Washer) {
	addImport(r.file.File, r.name, r.path)
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
