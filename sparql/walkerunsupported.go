package sparql

import "github.com/lercher/rdf/sparql/parser"

const (
	e1001 = 1001 + iota
	e1002
)

func (w *walker) EnterDatasetClause(ctx *parser.DatasetClauseContext) {
	w.SemanticErrorAt(ctx.GetStop(), e1001, "FROM clauses are not yet supported")
}

func (w *walker) EnterGraphGraphPattern(ctx *parser.GraphGraphPatternContext) {
	w.SemanticErrorAt(ctx.GetStart(), e1001, "external GRAPH references are not yet supported")
}
