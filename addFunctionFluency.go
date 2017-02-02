package wash

import (
	"go/ast"
	"log"
)

// FluentAddFunction is a fluent structure for adding functions
type FluentAddFunction struct {
	washer       *Washer
	file         *WashFile
	functionName string
}

// FluentAddFunctionParameter is a fluent structure for adding parameters to a new function
type FluentAddFunctionParameter struct {
	washer       *Washer
	file         *WashFile
	functionName string
	params       map[string]string
}

// FluentAddFunctionWithReturn is a fluent structure for defining return types of a new function
type FluentAddFunctionWithReturn struct {
	washer          *Washer
	file            *WashFile
	functionName    string
	params          map[string]string
	returnTypeNames []string
}

// WithParameter adds a parameter to the function
func (e FluentAddFunction) WithParameter(name string, typeName string) FluentAddFunctionParameter {
	params := make(map[string]string)
	params[name] = typeName
	return FluentAddFunctionParameter{
		washer:       e.washer,
		file:         e.file,
		functionName: e.functionName,
		params:       params,
	}
}

// Returning declares the return type(s) of the new function
func (e FluentAddFunctionParameter) Returning(typeNames ...string) FluentAddFunctionWithReturn {
	return FluentAddFunctionWithReturn{
		washer:          e.washer,
		file:            e.file,
		functionName:    e.functionName,
		params:          e.params,
		returnTypeNames: typeNames,
	}
}

func (e FluentAddFunctionWithReturn) getParamFields() []field {
	return []field{
		field{
			fieldName: "s",
			typeName:  "string",
		}}
}

func (e FluentAddFunctionWithReturn) getResultFields() []field {
	return []field{
		field{
			fieldName: "",
			typeName:  "int",
		},
		field{
			fieldName: "",
			typeName:  "bool",
		},
	}
}

// Build creates the new function and writes the updated file to disk
func (e FluentAddFunctionWithReturn) Build() {
	log.Printf("Adding function %s to %s", e.functionName, e.file.targetFilename)
	params := e.getParamFields()
	results := e.getResultFields()
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
