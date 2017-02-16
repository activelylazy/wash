package operations

import "github.com/activelylazy/wash"

// FluentFileEdit is a fluent structure for editing files
type FluentFileEdit struct {
	washer *wash.Washer
	file   *wash.File
}

// // Edit edits a file
// func (w *Washer) Edit(file *File) operations.FluentFileEdit {
// 	return operations.FluentFileEdit{
// 		file:   file,
// 		washer: w,
// 	}
// }
