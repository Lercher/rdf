package processor

import (
	"log"

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
	spsame := algebra.PatternOfSameVar(tp.S, tp.P)
	sosame := algebra.PatternOfSameVar(tp.S, tp.O)
	posame := algebra.PatternOfSameVar(tp.P, tp.O)
	log.Println(tp.S, tp.P, tp.O, spsame, sosame, posame)
	t := tp.Triple()
	return graph.Indexed(g, &t, givens, givenp, giveno, spsame, sosame, posame)
}
