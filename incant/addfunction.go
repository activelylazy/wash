package incant

import (
	"log"

	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/domain"
)

// NewFunctionBuilder allows the fluent creation of a new function
type NewFunctionBuilder struct {
	file         *wash.File
	testFile     *wash.File
	name         string
	returnValues []domain.Concept
	arguments    []domain.Concept
}

// NewFunction begins creation of a new function
func NewFunction(name string) *NewFunctionBuilder {
	return &NewFunctionBuilder{
		arguments: make([]domain.Concept, 0),
		name:      name,
	}
}

// In specifies the file for the new function to be placed in
func (b *NewFunctionBuilder) In(f *wash.File) *NewFunctionBuilder {
	b.file = f
	return b
}

// InFile specifies the file for the new function to be placed in by name/path
func (b *NewFunctionBuilder) InFile(name string) *NewFunctionBuilder {
	return b.In(wash.FindFile(name))
}

// WithTestIn specifies the file to write the test to
func (b *NewFunctionBuilder) WithTestIn(f *wash.File) *NewFunctionBuilder {
	b.testFile = f
	return b
}

// Build builds the new function
func (b *NewFunctionBuilder) Build() {
	fn := b.file.AddFunction(b.name, domain.ConceptsToFields(b.arguments), b.returnValues)

	if err := wash.WriteFunctionCallTest(b.testFile, fn, b.arguments, b.returnValues); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// Given specifies an initial set of arguments to pass to the new function in the first test
func (b *NewFunctionBuilder) Given(arguments ...domain.Concept) *NewFunctionBuilder {
	b.arguments = arguments
	return b
}

// ShouldReturn specifies the default return values for the new function
func (b *NewFunctionBuilder) ShouldReturn(values ...domain.Concept) *NewFunctionBuilder {
	b.returnValues = values
	return b
}
