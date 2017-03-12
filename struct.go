package wash

import "go/ast"

// Struct represents a struct managed by wash
type Struct struct {
	File       File
	structName string
	Decl       *ast.GenDecl
}
