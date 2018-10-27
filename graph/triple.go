package graph

type Triple struct {
	S, P, O Virtual
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
