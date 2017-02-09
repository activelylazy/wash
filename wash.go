package wash

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Washer allows modification of Go code
type Washer struct {
	pkgs     map[string]*ast.Package
	fset     *token.FileSet
	basePath string
	concepts map[string]DomainConcept
}

// DomainConcept represents a named, typed value in the system
type DomainConcept struct {
	name     string
	typeName string
	value    string
}

// WashFile represents a file managed by washer
type File struct {
	targetFilename string
	file           *ast.File
	washer         *Washer
}

// FluentFileEdit is a fluent structure for editing files
type FluentFileEdit struct {
	washer *Washer
	file   *File
}

// NewWasher creates a new Washer
func NewWasher(basePath string) (*Washer, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, basePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return &Washer{
		basePath: basePath,
		pkgs:     pkgs,
		fset:     fset,
		concepts: make(map[string]DomainConcept),
	}, nil
}

// Edit edits a file
func (w *Washer) Edit(file *File) FluentFileEdit {
	return FluentFileEdit{
		file:   file,
		washer: w,
	}
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

// AddFunction begins adding a function to a file
func (e FluentFileEdit) AddFunction(name string) FluentAddFunction {
	return FluentAddFunction{
		washer:       e.washer,
		file:         e.file,
		functionName: name,
	}
}
