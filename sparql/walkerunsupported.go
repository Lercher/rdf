package sparql

import "github.com/lercher/rdf/sparql/parser"

const (
	e1001 = 1001 + iota
	e1002
	e1003
)

func (w *walker) EnterDatasetClause(ctx *parser.DatasetClauseContext) {
	w.SemanticErrorAt(ctx.GetStop(), e1001, "FROM clauses are not yet supported")
}
