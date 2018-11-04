package sparql

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lercher/rdf/sparql/parser"
)

// Parse parses a SPARQL query to an abstract syntax tree
func Parse(stream antlr.CharStream) (*AST, error) {
	upperstream := NewCaseChangingStream(stream, true)
	lexer := parser.NewSparqlLexer(upperstream)
	tstream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSparqlParser(tstream)
	p.RemoveErrorListeners() // get rid of the "line 1:3 mismatched input '<EOF>' expecting ..." outputs

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
