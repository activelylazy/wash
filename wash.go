package wash

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
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

	washer, err := newWasher(basePath)
	if err != nil {
		fmt.Printf("[ERROR] Could not set base path: %s\n", err)
	}
	ctx.washer = washer
}

// CreateFile creates a new, empty file
func CreateFile(filename string, packageName string) *File {
	f, err := ctx.washer.createFile(filename, packageName)
	if err != nil {
		fmt.Printf("[ERROR] Could not create file: %s\n", err)
		return nil
	}
	return f
}

// FindFile finds a file by name / path
func FindFile(filename string) *File {
	f, err := ctx.washer.findFile(filename)
	if err != nil {
		fmt.Printf("[ERROR] Could not find file: %s\n", err)
	}
	return f
}

// NewWasher creates a new Washer
func newWasher(basePath string) (*Washer, error) {
	fset := token.NewFileSet()
	pkgs, err := getAllPackages(fset, basePath)
	if err != nil {
		return nil, err
	}
	return &Washer{
		BasePath: basePath,
		pkgs:     pkgs,
		Fset:     fset,
	}, nil
}

func getAllPackages(fset *token.FileSet, basePath string) (map[string]*ast.Package, error) {
	pkgs := make(map[string]*ast.Package)
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}
	if err = getPackages(fset, basePath, pkgs); err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			if err = getPackages(fset, path.Join(basePath, f.Name()), pkgs); err != nil {
				return nil, err
			}
		}
	}

	return pkgs, nil
}

func getPackages(fset *token.FileSet, path string, pkgs map[string]*ast.Package) error {
	p, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	for k, v := range p {
		pkgs[k] = v
	}
	return nil
}

func (washer *Washer) createFile(filename string, packageName string) (*File, error) {
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

func (washer *Washer) findFile(filename string) (*File, error) {
	targetFilename := path.Join(washer.BasePath, filename)
	targetFilename = strings.Replace(targetFilename, "\\", "/", -1)
	if _, err := os.Stat(targetFilename); err != nil {
		return nil, err
	}
	dir, _ := path.Split(targetFilename)
	dir = path.Clean(dir)
	lastSlash := strings.LastIndex(dir, "/")
	packageName := dir[lastSlash+1:]

	for name, pkg := range ctx.washer.pkgs {
		if name == packageName {
			for k, f := range pkg.Files {
				if strings.Replace(k, "\\", "/", -1) == targetFilename {
					return newWashFile(targetFilename, f, washer), nil
				}
			}
		}
	}

	return nil, fmt.Errorf("Faild to find %s in the parsed sources", filename)
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
