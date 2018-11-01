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
}

func newAST() *AST {
	 ast :=&AST{
		symbols: symbols{
			Variables: algebra.NewVariables(),
		},
	}
	ast.symbols.AddPrefix("rdf", RDF)
	ast.symbols.AddPrefix("rdfs", RDFS)
	ast.symbols.AddPrefix("xs", XS)
	return ast
}

func (s *symbols) AddPrefix(p string, iri graph.IRI) *graph.PrefixedIRI {
	pi := &graph.PrefixedIRI{
		Prefix: graph.PrefixParse(p),
		IRI: iri,
	}
	s.PrefixedIRIs = append(s.PrefixedIRIs, pi)
	return pi
}