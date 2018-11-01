package sparql

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/sparql/parser"
)

const (
	// RDF is the rdf: namespace, rdf syntax
	RDF  = graph.IRI(`http://www.w3.org/1999/02/22-rdf-syntax-ns#`)

	// RDFS is the rdfs: namespace, rdf schema
	RDFS = graph.IRI(`http://www.w3.org/2000/01/rdf-schema#`)

	// XS is the xs: namespace, XML schema
	XS  = graph.IRI(`http://www.w3.org/2001/XMLSchema#`)
	
	// FOAF is the foaf: namespace, friend of a friend found in many rdf examples
	FOAF = graph.IRI(`http://xmlns.com/foaf/0.1/`)

	// A is rdf:type, the semantic of 'a' in a sparql query
	A    = graph.IRI(`http://www.w3.org/1999/02/22-rdf-syntax-ns#type`)
)

// Parse parses a SPARQL query to an abstract syntax tree
func Parse(stream antlr.CharStream) (*AST, error) {
	upperstream := NewCaseChangingStream(stream, true)
	lexer := parser.NewSparqlLexer(upperstream)
	tstream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSparqlParser(tstream)

	el := &ErrorListener{DiagnosticErrorListener: antlr.NewDiagnosticErrorListener(true), errors: new(errors)}
	p.AddErrorListener(el)

	p.BuildParseTrees = true
	tree := p.Query()
	w := &walker{ast: newAST(), ErrorListener: el}
	if el.errors.ErrorCount > 0 {
		return nil, fmt.Errorf("%d syntax error(s)", el.errors.ErrorCount)
	}
	antlr.ParseTreeWalkerDefault.Walk(w, tree)
	if el.errors.ErrorCount > 0 || el.errors.WarningCount > 0 {
		return w.ast, fmt.Errorf("%d semantic error(s) and %d warning(s)", el.errors.ErrorCount, el.errors.WarningCount)
	}
	return w.ast, nil
}
