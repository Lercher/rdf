package graph

import (
	"errors"
	"sync"
)

// Graph forms an RDF graph including several hidden indices
type Graph struct {
	Dataset   []*Triple
	valuemap  map[Virtual]Value
	valuesize int
	sindex    map[Virtual][]*Triple
	pindex    map[Virtual][]*Triple
	oindex    map[Virtual][]*Triple
	spindex   map[Virtual2][]*Triple
	soindex   map[Virtual2][]*Triple
	poindex   map[Virtual2][]*Triple
	spoindex  map[Triple]bool
}

// CountValues counts all distinct values in s/p/o position of the Graph
func (g *Graph) CountValues() int {
	return len(g.valuemap)
}

// Values returns an unsorted list of all distinct valus in the Graph
func (g *Graph) Values() []Value {
	list := make([]Value, 0, g.CountValues())
	for _, v := range g.valuemap {
		list = append(list, v)
	}
	return list
}

// New creates a new Graph
func New() *Graph {
	return &Graph{
		nil,
		make(map[Virtual]Value),
		0,
		make(map[Virtual][]*Triple),
		make(map[Virtual][]*Triple),
		make(map[Virtual][]*Triple),
		make(map[Virtual2][]*Triple),
		make(map[Virtual2][]*Triple),
		make(map[Virtual2][]*Triple),
		make(map[Triple]bool),
	}
}

// AddValue adds an unknown primitive to the graph's values. Returns either this new VirtualValue or the known one.
func (g *Graph) AddValue(primitive interface{}) VirtualValue {
	vv := g.VirtualValue(primitive)
	if !vv.Known {
		g.valuemap[vv.Virtual] = vv.Value
		g.valuesize += vv.Size
	}
	return vv
}

// Assert appends a new Triple to the Graph, unless it is already present
func (g *Graph) Assert(s, p, o interface{}) Triple {
	vs := g.AddValue(s)
	vp := g.AddValue(p)
	vo := g.AddValue(o)
	t := Triple{vs.Virtual, vp.Virtual, vo.Virtual}
	duplicate := g.spoindex[t]
	if duplicate {
		return t
	}
	g.BulkAddTriple(&t, true)
	return t
}

// Match locates a list of Triples that match a given triple of patterns
func (g *Graph) Match(p *TriplePattern) []*Triple {
	givens := !p.S.IsVariable()
	givenp := !p.P.IsVariable()
	giveno := !p.O.IsVariable()
	t := p.Triple()

	// 3 crit, no variables
	if givens && givenp && giveno {
		if g.spoindex[t] {
			return []*Triple{&t}
		}
		return nil
	}

	// 2 crit
	if givens && givenp {
		return g.spindex[t.SP()]
	}
	if givens && giveno {
		return g.soindex[t.SO()]
	}
	if givenp && giveno {
		return g.poindex[t.PO()]
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
	return g.Dataset
}

// PrepareValueSpace reserves memory for the given count of distinct values, it panics if the Graph is not empty
func (g *Graph) PrepareValueSpace(size int) {
	if len(g.valuemap) != 0 {
		panic(errors.New("value space can only prepared if the graph is empty"))
	}
	g.valuemap = make(map[Virtual]Value, size)
}

// PrepareVirtualSpace reserves memory for 3 times size hashes of 16 bytes, it panics if tha Graph is not empty
func (g *Graph) PrepareVirtualSpace(size int) {
	if g.Dataset != nil {
		panic(errors.New("virtual space can only prepared if the graph is empty"))
	}
	g.Dataset = make([]*Triple, 0, size)
}

// VirtualValue retrieves a primitive keyed by its hash from the graph or creates a new VirtualValue
func (g *Graph) VirtualValue(primitive interface{}) VirtualValue {
	return newVirtualValue(g, primitive)
}

// BulkAddTriple adds a Triple without checking to the Graph, see Assert instead
func (g *Graph) BulkAddTriple(t *Triple, withindex bool) *Triple {
	g.Dataset = append(g.Dataset, t)
	if withindex {
		g.sindex[t.S] = append(g.sindex[t.S], t)
		g.pindex[t.P] = append(g.pindex[t.P], t)
		g.oindex[t.O] = append(g.oindex[t.O], t)
		g.spindex[t.SP()] = append(g.spindex[t.SP()], t)
		g.soindex[t.SO()] = append(g.soindex[t.SO()], t)
		g.poindex[t.PO()] = append(g.poindex[t.PO()], t)
		g.spoindex[*t] = true
	}
	return t
}

// RebuildIndex rebuilds all the indices on the dataset
func (g *Graph) RebuildIndex() {
	sindex := make(map[Virtual]int)
	pindex := make(map[Virtual]int)
	oindex := make(map[Virtual]int)
	spindex := make(map[Virtual2]int)
	soindex := make(map[Virtual2]int)
	poindex := make(map[Virtual2]int)

	var wg sync.WaitGroup
	wg.Add(6)

	go func() {
		for _, t := range g.Dataset {
			sindex[t.S]++
		}
		wg.Done()
	}()

	go func() {
		for _, t := range g.Dataset {
			pindex[t.P]++
		}
		wg.Done()
	}()

	go func() {
		for _, t := range g.Dataset {
			oindex[t.O]++
		}
		wg.Done()
	}()

	go func() {
		for _, t := range g.Dataset {
			spindex[t.SP()]++
		}
		wg.Done()
	}()

	go func() {
		for _, t := range g.Dataset {
			soindex[t.SO()]++
		}
		wg.Done()
	}()

	go func() {
		for _, t := range g.Dataset {
			poindex[t.PO()]++
		}
		wg.Done()
	}()

	wg.Wait()
	// log.Println("index count done")

	wg.Add(7)

	go func() {
		g.sindex = make(map[Virtual][]*Triple, len(sindex))
		for k, v := range sindex {
			g.sindex[k] = make([]*Triple, 0, v)
			// log.Println("sindex", v, k.Value(g))
		}
		for _, tt := range g.Dataset {
			t := tt
			g.sindex[t.S] = append(g.sindex[t.S], t)
		}
		// log.Println("sindex done", len(g.sindex))
		wg.Done()
	}()

	go func() {
		g.pindex = make(map[Virtual][]*Triple, len(pindex))
		for k, v := range pindex {
			g.pindex[k] = make([]*Triple, 0, v)
		}
		for _, t := range g.Dataset {
			g.pindex[t.P] = append(g.pindex[t.P], t)
		}
		// log.Println("pindex done", len(g.pindex))
		wg.Done()
	}()

	go func() {
		g.oindex = make(map[Virtual][]*Triple, len(oindex))
		for k, v := range oindex {
			g.oindex[k] = make([]*Triple, 0, v)
			// log.Println("oindex", v, k.Value(g))
		}
		for _, t := range g.Dataset {
			g.oindex[t.O] = append(g.oindex[t.O], t)
		}
		// log.Println("oindex done", len(g.oindex))
		wg.Done()
	}()

	go func() {
		g.spoindex = make(map[Triple]bool, len(g.Dataset))
		for _, t := range g.Dataset {
			g.spoindex[*t] = true
		}
		// log.Println("spoindex done", len(g.spoindex))
		wg.Done()
	}()

	go func() {
		g.spindex = make(map[Virtual2][]*Triple, len(spindex))
		for k, v := range spindex {
			g.spindex[k] = make([]*Triple, 0, v)
		}
		for _, t := range g.Dataset {
			g.spindex[t.SP()] = append(g.spindex[t.SP()], t)
		}
		// log.Println("spindex done", len(g.spindex))
		wg.Done()
	}()

	go func() {
		g.soindex = make(map[Virtual2][]*Triple, len(soindex))
		for k, v := range soindex {
			g.soindex[k] = make([]*Triple, 0, v)
		}
		for _, t := range g.Dataset {
			g.soindex[t.SO()] = append(g.soindex[t.SO()], t)
		}
		// log.Println("soindex done", len(g.soindex))
		wg.Done()
	}()

	go func() {
		g.poindex = make(map[Virtual2][]*Triple, len(poindex))
		for k, v := range poindex {
			g.poindex[k] = make([]*Triple, 0, v)
		}
		for _, t := range g.Dataset {
			g.poindex[t.PO()] = append(g.poindex[t.PO()], t)
		}
		// log.Println("poindex done", len(g.poindex))
		wg.Done()
	}()

	wg.Wait()
}

func distinct(index map[Virtual][]*Triple) []Virtual {
	list := make([]Virtual, 0, len(index))
	for k := range index {
		list = append(list, k)
	}
	return list
}

// DistinctS returns all distinct P values in the Graph
func (g *Graph) DistinctS() []Virtual {
	return distinct(g.sindex)
}

// DistinctP returns all distinct P values in the Graph
func (g *Graph) DistinctP() []Virtual {
	return distinct(g.pindex)
}

// DistinctO returns all distinct P values in the Graph
func (g *Graph) DistinctO() []Virtual {
	return distinct(g.oindex)
}

// ByteSizes returns estimated sizes of Dataset, Values and Index
func (g *Graph) ByteSizes() (int, int, int) {
	ds := (3*16 + 8) * len(g.Dataset)

	idx := (3*16 + 8) * len(g.Dataset)
	idx += calcsize1(g.sindex)
	idx += calcsize1(g.pindex)
	idx += calcsize1(g.oindex)
	idx += calcsize2(g.spindex)
	idx += calcsize2(g.poindex)
	idx += calcsize2(g.soindex)

	return ds, g.valuesize, idx
}

func calcsize1(index map[Virtual][]*Triple) int {
	l := len(index)
	for _, list := range index {
		l += len(list)
	}
	return l * 8
}

func calcsize2(index map[Virtual2][]*Triple) int {
	l := len(index)
	for _, list := range index {
		l += len(list)
	}
	return l * 8
}
