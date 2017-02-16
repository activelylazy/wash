package operations

import (
	"go/ast"

	"github.com/activelylazy/wash"
)

// NewAppendToFunctionBodyRequest creates a new request to append to a function body
func NewAppendToFunctionBodyRequest(fn Function, stmt ast.Stmt) AppendToFunctionBodyRequest {
	return AppendToFunctionBodyRequest{
		fn:   fn,
		stmt: stmt,
	}
}

// Apply adds the statement to the function body
func (r AppendToFunctionBodyRequest) Apply(washer *wash.Washer) {
	r.fn.decl.Body.List = append(r.fn.decl.Body.List, r.stmt)
	r.fn.file.Write()
}
