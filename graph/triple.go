package graph

import "fmt"

type Triple struct {
	S, P, O Virtual
}

func (t *Triple) String(g *Graph) string {
	s := t.S.String(g)
	p := t.P.String(g)
	o := t.O.String(g)
	return fmt.Sprintf("[%s %s %s]", s, p, o)
}

func NewTriple(g *Graph, s, p, o interface{}) Triple {
	vs := NewVValue(g, s)
	vp := NewVValue(g, p)
	vo := NewVValue(g, o)
	return Triple{vs.Virtual, vp.Virtual, vo.Virtual}
}

func (t *Triple) SP() Virtual2 {
	return Virtual2{t.S, t.P}
}

func (t *Triple) SO() Virtual2 {
	return Virtual2{t.S, t.O}
}

func (t *Triple) PO() Virtual2 {
	return Virtual2{t.P, t.O}
}
