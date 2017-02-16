package operations

import (
	"go/ast"
	"log"
	"os"
	"path"

	"github.com/activelylazy/wash"
)

// CreateFileRequest allows creation of a new file
type CreateFileRequest struct {
	filename    string
	packageName string
}

// NewCreateFileRequest creates a new CreateFileRequest
func NewCreateFileRequest(filename string, packageName string) CreateFileRequest {
	return CreateFileRequest{
		filename:    filename,
		packageName: packageName,
	}
}

// Create creates a file
func (r CreateFileRequest) Create(washer *wash.Washer) (*wash.File, error) {
	targetFilename := path.Join(washer.BasePath, r.filename)
	log.Printf("Creating file %s in package %s", targetFilename, r.packageName)
	file := newFile(r.packageName)
	os.MkdirAll(path.Dir(targetFilename), 0700)
	washFile := wash.NewFile(targetFilename, file, washer)
	if err := washFile.Write(); err != nil {
		return nil, err
	}
	return washFile, nil
}

func newFile(packageName string) *ast.File {
	f := &ast.File{}
	f.Name = ast.NewIdent(packageName)
	return f
}
