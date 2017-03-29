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
}

// DomainConcept represents a named, typed value in the system
type DomainConcept struct {
	Name  string
	Type  DomainType
	value string
}

// DomainType represents a type of data in the system
type DomainType struct {
	Name     string
	TypeName string
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
	}, nil
}

// String converts a DomainConcept to a string representation to output into code
func (c DomainConcept) String() string {
	if c.Type.TypeName == "string" {
		return "\"" + c.value + "\""
	}
	return c.value
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

// NewDomainType adds a new domain type
func NewDomainType(name string, typeName string) DomainType {
	return DomainType{
		Name:     name,
		TypeName: typeName,
	}
}

// NewInstance adds a new domain concept - a named instance of a domain type with a specific value
func (t DomainType) NewInstance(name string, value string) DomainConcept {
	c := DomainConcept{
		Name:  name,
		Type:  t,
		value: value,
	}
	return c
}
