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
	fset     *token.FileSet
	basePath string
}

// FluentFileCreator allows creation of a new file
type FluentFileCreator struct {
	filename string
	washer   *Washer
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
	}, nil
}

// CreateFile creates a new file with the given name
func (w *Washer) CreateFile(filename string) FluentFileCreator {
	return FluentFileCreator{
		filename: filename,
		washer:   w,
	}
}

// InPackage creates a new file in a given package
func (f FluentFileCreator) InPackage(packageName string) error {
	targetFilename := path.Join(f.washer.basePath, f.filename)
	log.Printf("Creating file %s in package %s", targetFilename, packageName)
	file := newFile(packageName)
	os.MkdirAll(path.Dir(targetFilename), 0700)
	outfile, err := os.Create(targetFilename)
	if err != nil {
		return err
	}
	defer outfile.Close()
	printer.Fprint(outfile, f.washer.fset, file)
	return nil
}

func newFile(packageName string) *ast.File {
	f := &ast.File{}
	f.Name = ast.NewIdent(packageName)
	return f
}
