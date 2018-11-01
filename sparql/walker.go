package sparql

import (
	"strings"

	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/sparql/parser"
)

// Walker represents the AST of a Query
type walker struct {
	*parser.BaseSparqlListener
	*ErrorListener
	ast *AST
}

const (
	e2001 = 2001 + iota
	e2002
	e2003
	e2004
	e2005
	e2006
)

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

func (w *walker) EnterSelectQuery(ctx *parser.SelectQueryContext) {
	mod := ctx.GetMod()
	if mod != nil {
		w.ast.Modifier = strings.ToLower(mod.GetText())
	}
	w.ast.QueryType = "select"
	if ctx.GetVarstar() != nil {
		w.ast.Projection.All = true
	} else {
		w.ast.Projection.All = false
		for _, vtctx := range ctx.GetVars() {
			vt := vtctx.GetText()
			v := w.ast.Variables.Register(vt)
			w.ast.Projection.Variables = append(w.ast.Projection.Variables, v)
		}
	}
}

func (w *walker) EnterConstructQuery(ctx *parser.ConstructQueryContext) {
	w.ast.QueryType = "construct"
}

func (w *walker) EnterDescribeQuery(ctx *parser.DescribeQueryContext) {
	w.ast.QueryType = "describe"
}

func (w *walker) EnterAskQuery(ctx *parser.AskQueryContext) {
	w.ast.QueryType = "ask"
}

func (w *walker) EnterTriplesSameSubject(ctx *parser.TriplesSameSubjectContext) {
	subjectC := ctx.GetSubject()
	if subjectC != nil {
		subject, err := convertVarOrTermContext(subjectC, &w.ast.symbols)
		if err != nil {
			w.SemanticErrorAt(subjectC.GetStart(), e2004, err.Error())
		}
		propertiesC := ctx.GetProperties()
		verbs := propertiesC.GetVerbs()
		for i, v := range verbs {
			predicate, err := convertVerbContext(v, &w.ast.symbols)
			if err != nil {
				w.SemanticErrorAt(v.GetStart(), e2005, err.Error())
			}
			// https://www.w3.org/TR/rdf-sparql-query/#objLists
			objectlistC := propertiesC.GetOl()[i]
			for _, objC := range objectlistC.GetOb() {
				graphNodeC := objC.GetGn()
				vtC := graphNodeC.GetVt()
				if vtC != nil {
					object, err := convertVarOrTermContext(vtC, &w.ast.symbols)
					if err != nil {
						w.SemanticErrorAt(vtC.GetStart(), e2006, err.Error())
					}
					block := &Block{
						Mode:      "BGP",
						Subject:   subject,
						Predicate: predicate,
						Object:    object,
					}
					w.ast.Where = append(w.ast.Where, block)
				} else {
					tnC := graphNodeC.GetTn() // triplesNode
					w.SemanticErrorAt(tnC.GetStart(), e1003, "triplesNode as object not yet supported")
				}
			}
		}
	}
}
