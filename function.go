package wash

import (
	"go/ast"
	"log"
)

// Function represents a function managed by wash
type Function struct {
	File         File
	functionName string
	Decl         *ast.FuncDecl
}

// AppendStmt writes a statement to the end of the function body
func (fn Function) AppendStmt(stmt ast.Stmt) {
	fn.Decl.Body.List = append(fn.Decl.Body.List, stmt)
	fn.File.Write()
}

// Append writes the given go code to the end of the function body
func (fn Function) Append(s string) {
	stmt, err := ParseStatement(s)
	if err != nil {
		log.Fatalf("Error parsing statement: %v", err)
	}
	fn.AppendStmt(stmt)
}
