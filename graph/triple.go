package graph

import (
	"fmt"
	"github.com/lercher/rdf/values"
)

// Triple stores an assertation in a Graph in virtual form,
// i.e. instead of the values it holds 3x16 bytes of hash values.
// Therefore a Triple must always be coupled with a graph from witch
// it was created. It is undefined behaviour to use a triple with a different Graph.
type Triple struct {
	S, P, O Virtual
}

func (t *Triple) String(g *Graph) string {
	s := t.S.String(g)
	p := t.P.String(g)
	o := t.O.String(g)
	return fmt.Sprintf("[%s %s %s]", s, p, o)
}

// NewTriple crates a new Triple with Values of a Graph. However, new Values are not added to the graph.
func NewTriple(g *Graph, s, p, o values.Value) *Triple {
	vs := g.VirtualValue(s)
	vp := g.VirtualValue(p)
	vo := g.VirtualValue(o)
	return &Triple{vs.Virtual, vp.Virtual, vo.Virtual}
}

// SP creates a pair from the subject/predicate Virtuals
func (t *Triple) SP() Virtual2 {
	return Virtual2{t.S, t.P}
}

// SO creates a pair from the subject/object Virtuals
func (t *Triple) SO() Virtual2 {
	return Virtual2{t.S, t.O}
}

// PO creates a pair from the predicate/object Virtuals
func (t *Triple) PO() Virtual2 {
	return Virtual2{t.P, t.O}
}

// TripleFilter filters a slice of Triples by a predicate of Triple
func TripleFilter(ts []*Triple, pred func(*Triple) bool) (list []*Triple) {
	for _, t := range ts {
		if pred(t) {
			list = append(list, t)
		}
	}
	return list
}
