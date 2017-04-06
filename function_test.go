package wash

import (
	"go/ast"
	"testing"

	"github.com/activelylazy/wash/domain"
	"github.com/activelylazy/wash/syntax"
	"github.com/stretchr/testify/mock"
)

type mockFile struct {
	mock.Mock
}

func (f *mockFile) AddImport(name string, path string) {
	_ = f.Called(name, path)
}

func (f *mockFile) AddFunction(functionName string, params []syntax.Field, returnValues []domain.Concept) Function {
	args := f.Called(functionName, params, returnValues)
	return args.Get(0).(Function)
}

func (f *mockFile) Write() error {
	args := f.Called()
	return args.Error(0)
}

func (f *mockFile) RelPath() (string, error) {
	args := f.Called()
	return args.String(0), args.Error(1)
}

func TestAppendStmtAppendsToStatementListAndWritesFile(t *testing.T) {
	mockFile := &mockFile{}
	mockFile.On("Write").Return(nil)

	fn := Function{
		File: mockFile,
		Decl: &ast.FuncDecl{
			Body: &ast.BlockStmt{
				List: make([]ast.Stmt, 0),
			},
		},
	}
	stmt := &ast.IfStmt{}

	fn.AppendStmt(stmt)

	mockFile.AssertExpectations(t)
	if len(fn.Decl.Body.List) != 1 {
		t.Errorf("Expected body list to have 1 statement but has %d", len(fn.Decl.Body.List))
	}
	if fn.Decl.Body.List[0] != stmt {
		t.Errorf("Expected dummy statement but was %v", fn.Decl.Body.List[0])
	}
}
