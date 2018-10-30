package sparql

import (
	"github.com/lercher/rdf/graph"
	"log"

	"github.com/lercher/rdf/sparql/parser"
)

// Walker represents the AST of a Query
type walker struct {
	*parser.BaseSparqlListener
	ast AST
}

//func (w *walker) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText())
//}

func (w *walker) EnterBaseDecl(ctx *parser.BaseDeclContext) {
	log.Println("EnterBaseDecl")
	iri := ctx.GetIri()
	w.ast.Base = graph.IRIParse(iri.GetText())
}
