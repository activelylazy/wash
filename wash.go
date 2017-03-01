package wash

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path"
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

// CreateFile creates a new go file
func (washer *Washer) CreateFile(filename string, packageName string) (*File, error) {
	targetFilename := path.Join(washer.BasePath, filename)
	log.Printf("Creating file %s in package %s", targetFilename, packageName)
	file := newFile(packageName)
	os.MkdirAll(path.Dir(targetFilename), 0700)
	washFile := newWashFile(targetFilename, file, washer)
	if err := washFile.Write(); err != nil {
		return nil, err
	}
	return washFile, nil
}

func newWashFile(targetFilename string, file *ast.File, washer *Washer) *File {
	return &File{
		TargetFilename: targetFilename,
		File:           file,
		washer:         washer,
	}
}

func newFile(packageName string) *ast.File {
	f := &ast.File{}
	f.Name = ast.NewIdent(packageName)
	return f
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
func (washer *Washer) NewDomainConcept(name string, typeName string, value string) DomainConcept {
	c := DomainConcept{
		name:     name,
		typeName: typeName,
		value:    value,
	}
	washer.concepts[name] = c
	return c
}
