package graph

import (
	"fmt"
	"strings"
)

// IRI encapsules rdf uri like <http://identifiers>
type IRI string

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
