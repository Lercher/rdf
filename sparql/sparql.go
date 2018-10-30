package sparql

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lercher/rdf/sparql/parser"
)


func Parse(stream antlr.CharStream) (*AST, error) {
	lexer := parser.NewSparqlLexer(stream)
	tstream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSparqlParser(tstream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Query()
	w := &walker{}
	antlr.ParseTreeWalkerDefault.Walk(w, tree)
	return &w.ast, nil
}
