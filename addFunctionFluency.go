package wash

import (
	"go/ast"
	"log"
)

// FluentAddFunction is a fluent structure for adding functions
type FluentAddFunction struct {
	washer       *Washer
	file         *File
	functionName string
}

// FluentAddFunctionParameter is a fluent structure for adding parameters to a new function
type FluentAddFunctionParameter struct {
	washer       *Washer
	file         *File
	functionName string
	params       []field
}

// FluentAddFunctionWithReturn is a fluent structure for defining return types of a new function
type FluentAddFunctionWithReturn struct {
	washer       *Washer
	file         *File
	functionName string
	params       []field
	returnFields []field
}

// WithParameter adds a parameter to the function
func (e FluentAddFunction) WithParameter(name string, typeName string) FluentAddFunctionParameter {
	params := []field{newField(name, typeName)}
	return FluentAddFunctionParameter{
		washer:       e.washer,
		file:         e.file,
		functionName: e.functionName,
		params:       params,
	}
}

// Returning declares the return type(s) of the new function
func (e FluentAddFunctionParameter) Returning(typeNames ...string) FluentAddFunctionWithReturn {
	returnFields := make([]field, len(typeNames))
	for i := 0; i < len(typeNames); i++ {
		returnFields[i] = newField("", typeNames[i])
	}
	return FluentAddFunctionWithReturn{
		washer:       e.washer,
		file:         e.file,
		functionName: e.functionName,
		params:       e.params,
		returnFields: returnFields,
	}
}

// Build creates the new function and writes the updated file to disk
func (e FluentAddFunctionWithReturn) Build() {
	log.Printf("Adding function %s to %s", e.functionName, e.file.targetFilename)
	params := e.params
	results := e.returnFields
	addFunction(e.file.file, e.functionName, params, results,
		[]ast.Stmt{
		// &ast.ReturnStmt{
		// 	Results: []ast.Expr{
		// 		newBasicLit("0"),
		// 		newBasicLit("false"),
		// 	},
		// },
		})
	e.file.write()
}

func newField(name string, typeName string) field {
	return field{
		fieldName: name,
		typeName:  typeName,
	}
}
