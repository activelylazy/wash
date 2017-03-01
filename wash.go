package wash

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

// Washer allows modification of Go code
type Washer struct {
	pkgs     map[string]*ast.Package
	Fset     *token.FileSet
	BasePath string
	concepts map[string]DomainConcept
}

// DomainConcept represents a named, typed value in the system
type DomainConcept struct {
	name     string
	typeName string
	value    string
}

// NewFile creates a new wash file
func NewFile(targetFilename string, file *ast.File, washer *Washer) *File {
	return &File{
		TargetFilename: targetFilename,
		File:           file,
		washer:         washer,
	}
}

// Function represents a function managed by wash
type Function struct {
	File         File
	functionName string
	Decl         *ast.FuncDecl
}

// NewWasher creates a new Washer
func NewWasher(basePath string) (*Washer, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, basePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return &Washer{
		BasePath: basePath,
		pkgs:     pkgs,
		Fset:     fset,
		concepts: make(map[string]DomainConcept),
	}, nil
}

func (w *File) Write() error {
	outfile, err := os.Create(w.TargetFilename)
	if err != nil {
		return err
	}
	defer outfile.Close()
	printer.Fprint(outfile, w.washer.Fset, w.File)
	return nil
}

// NewDomainConcept adds a new domain concept - a named, typed value
func (w *Washer) NewDomainConcept(name string, typeName string, value string) DomainConcept {
	c := DomainConcept{
		name:     name,
		typeName: typeName,
		value:    value,
	}
	w.concepts[name] = c
	return c
}
