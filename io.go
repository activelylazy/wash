package wash

import (
	"go/printer"
	"os"
)

func (w *File) write() error {
	outfile, err := os.Create(w.targetFilename)
	if err != nil {
		return err
	}
	defer outfile.Close()
	printer.Fprint(outfile, w.washer.fset, w.file)
	return nil
}
