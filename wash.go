package wash

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
)

var ctx = context{}

type context struct {
	washer *Washer
}

// Washer allows modification of Go code
type Washer struct {
	pkgs     map[string]*ast.Package
	Fset     *token.FileSet
	BasePath string
}

// SetBasePath sets the base path to use when writing source files
func SetBasePath(packageName string) {
	var gopath = os.Getenv("GOPATH")
	var gorootSrc = filepath.Join(filepath.Clean(gopath), "src")
	var basePath = filepath.Join(gorootSrc, packageName)

	if _, err := os.Stat(basePath); err != nil {
		fmt.Printf("[ERROR] Could not find package path %s\n", basePath)
	}

	washer, err := NewWasher(basePath)
	if err != nil {
		fmt.Printf("[ERROR] Could not set base path: %s\n", err)
	}
	ctx.washer = washer
}

func CreateFile(filename string, packageName string) *File {
	f, err := ctx.washer.CreateFile(filename, packageName)
	if err != nil {
		fmt.Printf("[ERROR] Could not create file: %s\n", err)
		return nil
	}
	return f
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
