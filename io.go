package wash

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func writeFile(targetFilename string, fset *token.FileSet, file *ast.File) error {
	outfile, err := os.Create(targetFilename)
	if err != nil {
		return err
	}
	defer outfile.Close()
	printer.Fprint(outfile, fset, file)
	return nil
}
