package operations

import (
	"go/ast"

	"github.com/activelylazy/wash"
)

// AppendToFunctionBodyRequest represents a request to append to a function body's statement list
type AppendToFunctionBodyRequest struct {
	fn   wash.Function
	stmt ast.Stmt
}

// NewAppendToFunctionBodyRequest creates a new request to append to a function body
func NewAppendToFunctionBodyRequest(fn wash.Function, stmt ast.Stmt) AppendToFunctionBodyRequest {
	return AppendToFunctionBodyRequest{
		fn:   fn,
		stmt: stmt,
	}
}

// Apply adds the statement to the function body
func (r AppendToFunctionBodyRequest) Apply(washer *wash.Washer) {
	r.fn.Decl.Body.List = append(r.fn.Decl.Body.List, r.stmt)
	r.fn.File.Write()
}
