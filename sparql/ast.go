package sparql

import "github.com/lercher/rdf/graph"
import "github.com/lercher/rdf/algebra"

// AST is the abstract syntax tree of a query
type AST struct {
	temporary
	symbols
	QueryType  string // select|ask|construct|describe
	Modifier   string // nil|distinct|reduced
	Projection algebra.Projection
	Where      []*Block
}

type symbols struct {
	Base         graph.IRI
	PrefixedIRIs []*graph.PrefixedIRI
	Variables    *algebra.Variables
}

type temporary struct {
	groupGraphPattern []*Block
}

func newAST() *AST {
	return &AST{
		symbols: symbols{
			Variables: algebra.NewVariables(),
		},
	}
}
