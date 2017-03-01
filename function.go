package wash

import "go/ast"

// Function represents a function managed by wash
type Function struct {
	File         File
	functionName string
	Decl         *ast.FuncDecl
}

// Append writes a statement to the end of the function body
func (fn Function) Append(stmt ast.Stmt) {
	fn.Decl.Body.List = append(fn.Decl.Body.List, stmt)
	fn.File.Write()
}
