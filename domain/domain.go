package domain

import "github.com/activelylazy/wash/syntax"

// Concept represents a named, typed value in the system
type Concept struct {
	Name  string
	Type  Type
	Value string
}

// Type represents a type of data in the system
type Type struct {
	Name     string
	TypeName string
}

// String converts a DomainConcept to a string representation to output into code
func (c Concept) String() string {
	if c.Type.TypeName == "string" {
		return "\"" + c.Value + "\""
	}
	return c.Value
}

// NewType adds a new domain type
func NewType(name string, typeName string) Type {
	return Type{
		Name:     name,
		TypeName: typeName,
	}
}

// NewInstance adds a new domain concept - a named instance of a domain type with a specific value
func (t Type) NewInstance(name string, value string) Concept {
	c := Concept{
		Name:  name,
		Type:  t,
		Value: value,
	}
	return c
}

// ConceptsToFields converts a slice of Concept to a slice of fields, using the Type's names
func ConceptsToFields(concepts []Concept) []syntax.Field {
	results := make([]syntax.Field, len(concepts))
	for i, c := range concepts {
		results[i] = syntax.NewField(c.Type.Name, c.Type.TypeName)
	}
	return results
}
