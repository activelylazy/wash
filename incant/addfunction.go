package incant

import (
	"log"

	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/syntax"
)

// NewFunctionBuilder allows the fluent creation of a new function
type NewFunctionBuilder struct {
	file         *wash.File
	testFile     *wash.File
	name         string
	returnValues []wash.DomainConcept
	arguments    []wash.DomainConcept
}

// NewFunction begins creation of a new function
func NewFunction() *NewFunctionBuilder {
	return &NewFunctionBuilder{
		arguments: make([]wash.DomainConcept, 0),
	}
}

// In specifies the file for the new function to be placed in
func (b *NewFunctionBuilder) In(f *wash.File) *NewFunctionBuilder {
	b.file = f
	return b
}

// WithTestIn specifies the file to write the test to
func (b *NewFunctionBuilder) WithTestIn(f *wash.File) *NewFunctionBuilder {
	b.testFile = f
	return b
}

// Named sets the name of the new function
func (b *NewFunctionBuilder) Named(s string) *NewFunctionBuilder {
	b.name = s
	return b
}

// Build builds the new function
func (b *NewFunctionBuilder) Build() {
	fn := b.file.AddFunction(b.name, conceptsToFields(b.arguments), b.returnValues)

	if err := wash.WriteFunctionCallTest(b.testFile, fn, b.arguments, b.returnValues); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// WhenGiven specifies an initial set of arguments to pass to the new function in the first test
func (b *NewFunctionBuilder) WhenGiven(arguments ...wash.DomainConcept) *NewFunctionBuilder {
	b.arguments = arguments
	return b
}

// Returns specifies the default return values for the new function
func (b *NewFunctionBuilder) Returns(values ...wash.DomainConcept) *NewFunctionBuilder {
	b.returnValues = values
	return b
}

func conceptsToFields(concepts []wash.DomainConcept) []syntax.Field {
	results := make([]syntax.Field, len(concepts))
	for i, c := range concepts {
		results[i] = syntax.NewField(c.Name, c.TypeName)
	}
	return results
}
