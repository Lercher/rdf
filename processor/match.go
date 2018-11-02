package processor

import (
	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
)

// Match locates a list of Triples that match a given triple of patterns
func Match(g *graph.Graph, tp *algebra.TriplePattern) []*graph.Triple {
	if tp.S.WontMatch() || tp.P.WontMatch() || tp.O.WontMatch() {
		return nil
	}
	givens := !tp.S.IsVariable()
	givenp := !tp.P.IsVariable()
	giveno := !tp.O.IsVariable()
	t := tp.Triple()
	return graph.MatchPrimitive(g, &t, givens, givenp, giveno)
}
