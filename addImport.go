package wash

// AddImportRequest represents a request to add an import
type AddImportRequest struct {
	file *File
	name string
	path string
}

// NewAddImportRequest creates a new request to add an import
func NewAddImportRequest(file *File, name string, path string) AddImportRequest {
	return AddImportRequest{
		file: file,
		name: name,
		path: path,
	}
}

// Add adds the import to the file
func (r AddImportRequest) Add(washer *Washer) {
	addImport(r.file.file, r.name, r.path)
}
