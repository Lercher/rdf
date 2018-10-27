package graph

import "fmt"

// TriplePattern holds patterns for Subject, Predicate and Objekt
type TriplePattern struct {
	S, P, O *Pattern
}

func (tp *TriplePattern) String(g *Graph) string {
	s := tp.S.String(g)
	p := tp.P.String(g)
	o := tp.O.String(g)
	return fmt.Sprintf("{%s %s %s}", s, p, o)
}

// Triple constructs a Triple from the three patterns regardless of VariablePatterns, they produce the zero Virtual
func (tp *TriplePattern) Triple() Triple {
	return Triple{
		S: tp.S.Virtual(),
		P: tp.P.Virtual(),
		O: tp.O.Virtual(),
	}
}
