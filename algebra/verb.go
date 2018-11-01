package algebra

import (
	"fmt"

	"github.com/lercher/rdf/graph"
)

// Verb is either a Variable or a literal IRI
type Verb struct {
	Variable
	Literal graph.IRI
}

// VerbVariable creates a Variable Verb
func VerbVariable(v Variable) *Verb {
	return &Verb{
		Variable: v,
		Literal:  graph.NotAnIRI,
	}
}

// VerbLiteral creates a literal Verb
func VerbLiteral(iri graph.IRI) *Verb {
	return &Verb{
		Variable: NotAVariable,
		Literal:  iri,
	}
}

// IsLiteral holds true if the Verb is a Literal IRI
func (v *Verb) IsLiteral() bool {
	return v != nil && v.Variable == NotAVariable
}

// IsVariable holds true if the Verb is a Variable
func (v *Verb) IsVariable() bool {
	return v != nil && v.Variable != NotAVariable
}

func (v *Verb) String() string {
	if v.IsLiteral() {
		return fmt.Sprintf("Literal %s", v.Literal)
	}
	if v.IsVariable() {
		return fmt.Sprintf("Variable %d", v.Variable)
	}
	return "nil"
}
