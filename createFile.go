package wash

import (
	"log"
	"os"
	"path"
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
func (r CreateFileRequest) Create(washer *Washer) (*File, error) {
	targetFilename := path.Join(washer.basePath, r.filename)
	log.Printf("Creating file %s in package %s", targetFilename, r.packageName)
	file := newFile(r.packageName)
	os.MkdirAll(path.Dir(targetFilename), 0700)
	washFile := &File{
		targetFilename: targetFilename,
		file:           file,
		washer:         washer,
	}
	if err := washFile.write(); err != nil {
		return nil, err
	}
	return washFile, nil
}
