package operations

import (
	"go/ast"
	"log"

	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/syntax"
)

// AddFunctionRequest represents a request to add a new function to a file
type AddFunctionRequest struct {
	file         *wash.File
	functionName string
	params       []syntax.Field
	returnFields []syntax.Field
	returnValues []string
}

// Function represents a function managed by wash
type Function struct {
	file         *wash.File
	functionName string
	decl         *ast.FuncDecl
}

// NewAddFunctionRequest creates a new request to add a function
func NewAddFunctionRequest(file *wash.File, functionName string, params []syntax.Field, returnFields []syntax.Field, returnValues []string) AddFunctionRequest {
	return AddFunctionRequest{
		file:         file,
		functionName: functionName,
		params:       params,
		returnFields: returnFields,
		returnValues: returnValues,
	}
}

// Apply adds a new function to a file
func (r AddFunctionRequest) Apply(washer *wash.Washer) Function {
	log.Printf("Adding function %s to %s", r.functionName, r.file.TargetFilename)
	params := r.params
	results := r.returnFields
	statements := []ast.Stmt{}
	if len(r.returnValues) > 0 {
		statements = append(statements, newReturnStmt(r.returnValues))
	}
	decl := addFunction(r.file.File, r.functionName, params, results, statements)
	r.file.Write()
	return Function{
		file:         r.file,
		functionName: r.functionName,
		decl:         decl,
	}
}

func addFunction(f *ast.File, name string, params []syntax.Field, results []syntax.Field, statementList []ast.Stmt) *ast.FuncDecl {
	newDecl := &ast.FuncDecl{
		Name: syntax.NewIdent(name),
		Type: newFuncType(params, results),
		Body: &ast.BlockStmt{
			List: statementList,
		},
	}
	f.Decls = append(f.Decls, newDecl)
	return newDecl
}

func newFuncType(params []syntax.Field, results []syntax.Field) *ast.FuncType {
	return &ast.FuncType{
		Params:  syntax.NewFieldList(params),
		Results: syntax.NewFieldList(results),
	}
}

func newReturnStmt(returnValues []string) *ast.ReturnStmt {
	results := []ast.Expr{}
	for _, s := range returnValues {
		results = append(results, syntax.NewBasicLit(s))
	}
	return &ast.ReturnStmt{
		Results: results,
	}
}
