package wash

import (
	"go/ast"
	"log"
)

// AddFunctionRequest represents a request to add a new function to a file
type AddFunctionRequest struct {
	file         *File
	functionName string
	params       []Field
	returnFields []Field
	returnValues []string
}

// NewAddFunctionRequest creates a new request to add a function
func NewAddFunctionRequest(file *File, functionName string, params []Field, returnFields []Field, returnValues []string) AddFunctionRequest {
	return AddFunctionRequest{
		file:         file,
		functionName: functionName,
		params:       params,
		returnFields: returnFields,
		returnValues: returnValues,
	}
}

// Add adds a new function to a file
func (r AddFunctionRequest) Add(washer *Washer) {
	log.Printf("Adding function %s to %s", r.functionName, r.file.targetFilename)
	params := r.params
	results := r.returnFields
	addFunction(r.file.file, r.functionName, params, results,
		[]ast.Stmt{newReturnStmt(r.returnValues)})
	r.file.write()
}

func newReturnStmt(returnValues []string) *ast.ReturnStmt {
	results := []ast.Expr{}
	for _, s := range returnValues {
		results = append(results, newBasicLit(s))
	}
	return &ast.ReturnStmt{
		Results: results,
	}
}

// NewField creates a new field
func NewField(name string, typeName string) Field {
	return Field{
		fieldName: name,
		typeName:  typeName,
	}
}
