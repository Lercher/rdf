package graph

type Graph struct {
	Dataset  []Triple
	valuemap map[Virtual]Value
	sindex   map[Virtual][]*Triple
	pindex   map[Virtual][]*Triple
	oindex   map[Virtual][]*Triple
	spindex  map[Virtual2][]*Triple
	soindex  map[Virtual2][]*Triple
	poindex  map[Virtual2][]*Triple
	spoindex map[Triple]bool
}

func (g *Graph) CountValues() int {
	return len(g.valuemap)
}

func (g *Graph) Values() []Value {
	list := make([]Value, 0, g.CountValues())
	for _, v := range g.valuemap {
		list = append(list, v)
	}
	return list
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
	vs := NewVValue(g, s)
	vp := NewVValue(g, p)
	vo := NewVValue(g, o)
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

func (g *Graph) Match(p *TriplePattern) []*Triple {
	givens := !p.S.IsVariable()
	givenp := !p.P.IsVariable()
	giveno := !p.O.IsVariable()
	
	// 3 crit, no variables
	if givens && givenp && giveno {
		t := Triple{S: p.S.virtual, P: p.P.virtual, O: p.O.virtual}
		if g.spoindex[t] {
			return []*Triple{&t}
		}
		return nil
	}

	// 2 crit
	if givens && givenp {
		v2 := Virtual2{V1: p.S.virtual, V2: p.P.virtual}
		return g.spindex[v2]
	}
	if givens && giveno {
		v2 := Virtual2{V1: p.S.virtual, V2: p.O.virtual}
		return g.soindex[v2]
	}
	if givenp && giveno {
		v2 := Virtual2{V1: p.P.virtual, V2: p.O.virtual}
		return g.poindex[v2]
	}

	// 1 crit
	if givens {
		return g.sindex[p.S.virtual]
	}
	if givenp {
		return g.pindex[p.P.virtual]
	}
	if giveno {
		return g.oindex[p.O.virtual]
	}

	// 0 crit, all variables
	list := make([]*Triple, len(g.Dataset))
	for i := range g.Dataset {
		list[i] = &g.Dataset[i]
	}
	return list
}
