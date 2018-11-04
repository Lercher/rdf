package values

import (
	"fmt"
	"io"
	"strings"
)

// IRI encapsules rdf uri like <http://identifiers>
type IRI string

// NotAnIRI denotes the value for something that aggregates an optional IRI
const NotAnIRI = IRI("")

// String returns the rdf representation <http://identifier>
func (i IRI) String() string {
	return fmt.Sprintf("<%s>", string(i))
}

// URI returns the inner representation http://identifier of <http://identifier>
func (i IRI) URI() string {
	return string(i)
}

// IRIParse converts http://identifier and <http://identifier> to an IRI
func IRIParse(s string) IRI {
	if strings.HasPrefix(s, "<") && strings.HasSuffix(s, ">") {
		s = strings.TrimPrefix(s, "<")
		s = strings.TrimSuffix(s, ">")
	}
	return IRI(s)
}

// IRIPrefixed converts a prefixed IRIRef to full IRI
func IRIPrefixed(prefixed string, prefixes []*PrefixedIRI) (IRI, error) {
	sa := strings.SplitN(prefixed, ":", 2)
	if len(sa) != 2 {
		return "", fmt.Errorf("IRI %q is not prefixed", prefixed)
	}
	prefix := Prefix(sa[0])
	tail := sa[1]
	for _, pi := range prefixes {
		if pi.Prefix == prefix {
			return pi.Combine(tail), nil
		}
	}
	return "", fmt.Errorf("%q has unknown prefix %q", prefixed, prefix)
}

// Serialize to a typebyte and byte slice
func (i IRI) Serialize() (byte, []byte) {
	return 'I', serializeString(string(i))
}

func constructIRI(r io.Reader) (Value, error) {
	s, err := readString(r)
	return IRI(s), err
}


const (
	// RDF is the rdf: namespace, rdf syntax
	RDF  = IRI(`http://www.w3.org/1999/02/22-rdf-syntax-ns#`)

	// RDFS is the rdfs: namespace, rdf schema
	RDFS = IRI(`http://www.w3.org/2000/01/rdf-schema#`)

	// XS is the xs: namespace, XML schema
	XS  = IRI(`http://www.w3.org/2001/XMLSchema#`)
	
	// FOAF is the foaf: namespace, friend of a friend found in many rdf examples
	FOAF = IRI(`http://xmlns.com/foaf/0.1/`)

	// A is rdf:type, the semantic of 'a' in a sparql query
	A    = IRI(RDF + `type`)
)

// Inner returns the inner primitive of this Value
func (i IRI) Inner() interface{} {
	return string(i)
}
