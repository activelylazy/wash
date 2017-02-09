package wash

import (
	"go/ast"
	"go/token"
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

// AppendToFunctionBodyRequest represents a request to append to a function body's statement list
type AppendToFunctionBodyRequest struct {
	fn   Function
	stmt ast.Stmt
}

// Function represents a function managed by wash
type Function struct {
	file         *File
	functionName string
	decl         *ast.FuncDecl
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
func (r AddFunctionRequest) Add(washer *Washer) Function {
	log.Printf("Adding function %s to %s", r.functionName, r.file.targetFilename)
	params := r.params
	results := r.returnFields
	statements := []ast.Stmt{}
	if len(r.returnValues) > 0 {
		statements = append(statements, newReturnStmt(r.returnValues))
	}
	decl := addFunction(r.file.file, r.functionName, params, results, statements)
	r.file.write()
	return Function{
		file:         r.file,
		functionName: r.functionName,
		decl:         decl,
	}
}

func newReturnStmt(returnValues []string) *ast.ReturnStmt {
	results := []ast.Expr{}
	for _, s := range returnValues {
		results = append(results, NewBasicLit(s))
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

// NewDefineAssignStmt creates a new assign statement defining at least one new variable :=
func NewDefineAssignStmt(targetVarNames []string, rhs ...ast.Expr) *ast.AssignStmt {
	lhs := []ast.Expr{}
	for _, s := range targetVarNames {
		lhs = append(lhs, newIdent(s))
	}
	return &ast.AssignStmt{
		Lhs: lhs,
		Tok: token.DEFINE,
		Rhs: rhs,
	}
}

// NewCallExpr creates a new function call expression
func NewCallExpr(functionName string, args ...ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  newIdent(functionName),
		Args: args,
	}
}

// NewAppendToFunctionBodyRequest creates a new request to append to a function body
func NewAppendToFunctionBodyRequest(fn Function, stmt ast.Stmt) AppendToFunctionBodyRequest {
	return AppendToFunctionBodyRequest{
		fn:   fn,
		stmt: stmt,
	}
}

// Add adds the statement to the function body
func (r AppendToFunctionBodyRequest) Add(washer *Washer) {
	r.fn.decl.Body.List = append(r.fn.decl.Body.List, r.stmt)
	r.fn.file.write()
}
