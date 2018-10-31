package sparql

import "github.com/lercher/rdf/graph"
import "github.com/lercher/rdf/algebra"

// AST is the abstract syntax tree of a query
type AST struct {
	temporary
	Base         graph.IRI
	PrefixedIRIs []*graph.PrefixedIRI
	QueryType    string // select|ask|construct|describe
	Modifier     string // nil|distinct|reduced
	Variables    *algebra.Variables
	Projection   algebra.Projection
	Where        []*Block
}

type temporary struct {
	groupGraphPattern []*Block
}

func NewAST() *AST {
	return &AST{
		Variables: algebra.NewVariables(),
	}
}
