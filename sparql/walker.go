package sparql

import (
	"github.com/lercher/rdf/graph"
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
	iri := ctx.GetIri()
	w.ast.Base = graph.IRIParse(iri.GetText())
}

func (w *walker) EnterPrefixDecl(ctx *parser.PrefixDeclContext) {
	p := graph.PrefixParse(ctx.GetPrefix().GetText())
	i := graph.IRIParse(ctx.GetIri().GetText())
	pi := &graph.PrefixedIRI{Prefix: p, IRI: i}
	w.ast.PrefixedIRIs = append(w.ast.PrefixedIRIs, pi)
}
