package domain

// DomainConcept represents a named, typed value in the system
type DomainConcept struct {
	Name  string
	Type  DomainType
	Value string
}

// DomainType represents a type of data in the system
type DomainType struct {
	Name     string
	TypeName string
}

// String converts a DomainConcept to a string representation to output into code
func (c DomainConcept) String() string {
	if c.Type.TypeName == "string" {
		return "\"" + c.Value + "\""
	}
	return c.Value
}

// NewDomainType adds a new domain type
func NewDomainType(name string, typeName string) DomainType {
	return DomainType{
		Name:     name,
		TypeName: typeName,
	}
}

// NewInstance adds a new domain concept - a named instance of a domain type with a specific value
func (t DomainType) NewInstance(name string, value string) DomainConcept {
	c := DomainConcept{
		Name:  name,
		Type:  t,
		Value: value,
	}
	return c
}
