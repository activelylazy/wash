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
	returnValue  string
}

// NewAddFunctionRequest creates a new request to add a function
func NewAddFunctionRequest(file *File, functionName string, params []Field, returnFields []Field, returnValue string) AddFunctionRequest {
	return AddFunctionRequest{
		file:         file,
		functionName: functionName,
		params:       params,
		returnFields: returnFields,
		returnValue:  returnValue,
	}
}

// Add adds a new function to a file
func (r AddFunctionRequest) Add(washer *Washer) {
	log.Printf("Adding function %s to %s", r.functionName, r.file.targetFilename)
	params := r.params
	results := r.returnFields
	addFunction(r.file.file, r.functionName, params, results,
		[]ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{
					newBasicLit(r.returnValue),
				},
			},
		})
	r.file.write()
}

// NewField creates a new field
func NewField(name string, typeName string) Field {
	return Field{
		fieldName: name,
		typeName:  typeName,
	}
}
