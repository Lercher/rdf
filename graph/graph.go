package graph

type Graph struct {
	Dataset  []Triple
	values   map[Virtual]Value
	sindex   map[Virtual][]*Triple
	pindex   map[Virtual][]*Triple
	oindex   map[Virtual][]*Triple
	spindex  map[Virtual2][]*Triple
	soindex  map[Virtual2][]*Triple
	poindex  map[Virtual2][]*Triple
	spoindex map[Triple]bool
}

func New() *Graph {
	return &Graph{
		nil,
		make(map[Virtual]Value),
		make(map[Virtual][]*Triple),
		make(map[Virtual][]*Triple),
		make(map[Virtual][]*Triple),
		make(map[Virtual2][]*Triple),
		make(map[Virtual2][]*Triple),
		make(map[Virtual2][]*Triple),
		make(map[Triple]bool),
	}
}

func (g *Graph) Assert(s, p, o interface{}) Triple {
	vs := CreateVValue(g, s)
	vp := CreateVValue(g, p)
	vo := CreateVValue(g, o)
	t := Triple{vs.Virtual, vp.Virtual, vo.Virtual}
	duplicate := g.spoindex[t]
	if duplicate {
		return t
	}
	g.Dataset = append(g.Dataset, t)
	g.sindex[t.S] = append(g.sindex[t.S], &t)
	g.pindex[t.P] = append(g.pindex[t.P], &t)
	g.oindex[t.O] = append(g.oindex[t.O], &t)
	g.spindex[t.SP()] = append(g.spindex[t.SP()], &t)
	g.soindex[t.SO()] = append(g.soindex[t.SO()], &t)
	g.poindex[t.PO()] = append(g.poindex[t.PO()], &t)
	g.spoindex[t] = true
	return t
}
