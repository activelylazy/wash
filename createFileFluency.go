package wash

import (
	"log"
	"os"
	"path"
)

// FluentFileCreator allows creation of a new file
type FluentFileCreator struct {
	filename string
	washer   *Washer
}

// InPackage creates a new file in a given package
func (f FluentFileCreator) InPackage(packageName string) (*WashFile, error) {
	targetFilename := path.Join(f.washer.basePath, f.filename)
	log.Printf("Creating file %s in package %s", targetFilename, packageName)
	file := newFile(packageName)
	os.MkdirAll(path.Dir(targetFilename), 0700)
	err := writeFile(targetFilename, f.washer.fset, file)
	if err != nil {
		return nil, err
	}
	return &WashFile{
		targetFilename: targetFilename,
		file:           file,
	}, nil
}
