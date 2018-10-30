package graph


import (
	"fmt"
	"strings"
)

// Prefix encapsules a rdf query prefix like ontology:
type Prefix string

// String returns the prefix representation ontology:
func (p Prefix) String() string {
	return fmt.Sprintf("%s:", string(p)) 
}

// Inner returns the inner representation ontology
func (p Prefix) Inner() string {
	return string(p)
}

// PrefixParse converts ontology: and ontology to a Prefix
func PrefixParse(s string) Prefix {
	if strings.HasSuffix(s, ":") {
		s = strings.TrimSuffix(s, ":")
	}
	return Prefix(s)
}
