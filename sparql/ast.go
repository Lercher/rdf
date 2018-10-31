package sparql

import "github.com/lercher/rdf/graph"
import "github.com/lercher/rdf/algebra"

// AST is the abstract syntax tree of a query
type AST struct {
	Base         graph.IRI
	PrefixedIRIs []*graph.PrefixedIRI
	QueryType    string // select|ask|construct|describe
	Modifier     string // nil|distinct|reduced
	Variables    *algebra.Variables
	Projection   algebra.Projection
}

func NewAST() *AST {
	return &AST{
		Variables: algebra.NewVariables(),
	}
}
