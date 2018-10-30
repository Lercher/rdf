package sparql

import "github.com/lercher/rdf/graph"

// AST is the abstract syntax tree of a query
type AST struct {
	Base         graph.IRI
	PrefixedIRIs []*graph.PrefixedIRI
	QueryType    string // select|ask|construct|describe
	Modifier     string // nil|distinct|reduced
}
