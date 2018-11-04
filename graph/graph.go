package graph

import (
	"errors"
	"sync"

	"github.com/lercher/rdf/values"
)

// Graph forms an RDF graph including several hidden indices
type Graph struct {
	dataset          map[Triple]*Triple
	vhash2indexPlus1 map[vhash]Virtual
	valuelist        []values.Value
	valuesize        int
	sindex           map[Virtual][]*Triple
	pindex           map[Virtual][]*Triple
	oindex           map[Virtual][]*Triple
	spindex          map[Virtual2][]*Triple
	soindex          map[Virtual2][]*Triple
	poindex          map[Virtual2][]*Triple
}

// New creates a new Graph
func New() *Graph {
	return &Graph{
		make(map[Triple]*Triple),
		make(map[vhash]Virtual),
		nil,
		0,
		make(map[Virtual][]*Triple),
		make(map[Virtual][]*Triple),
		make(map[Virtual][]*Triple),
		make(map[Virtual2][]*Triple),
		make(map[Virtual2][]*Triple),
		make(map[Virtual2][]*Triple),
	}
}

// CountValues counts all distinct values in s/p/o position of the Graph
func (g *Graph) CountValues() int {
	return len(g.valuelist)
}

// Values returns an unsorted list of all distinct valus in the Graph
func (g *Graph) Values() []values.Value {
	return g.valuelist
}

// AssertLiterally asserts s and o as Literal and p as IRI
func (g *Graph) AssertLiterally(s, p, o string) *Triple {
	return g.Assert(values.LiteralString(s), values.IRIParse(p), values.LiteralString(o))
}

// Assert appends a new Triple to the Graph, unless it is already present
func (g *Graph) Assert(s, p, o values.Value) *Triple {
	vs := g.AddValue(s)
	vp := g.AddValue(p)
	vo := g.AddValue(o)
	t := &Triple{vs.Virtual, vp.Virtual, vo.Virtual}
	tt, duplicate := g.dataset[*t]
	if duplicate {
		return tt
	}
	g.BulkAddTriple(t, true)
	return t
}

// DatasetMap returns the asserted Triples
func (g *Graph) DatasetMap() map[Triple]*Triple {
	return g.dataset
}

// Dataset returns a randomly sorted list of asserted Triples
func (g *Graph) Dataset() []*Triple {
	list := make([]*Triple, 0, len(g.dataset))
	for k := range g.dataset {
		kk := k
		list = append(list, &kk)
	}
	return list
}

// CountTriples returns the count of assertions
func (g *Graph) CountTriples() int {
	return len(g.dataset)
}

// PrepareValueSpace reserves memory for the given count of distinct values, it panics if the Graph is not empty
func (g *Graph) PrepareValueSpace(size int) {
	if len(g.valuelist) != 0 {
		panic(errors.New("value space can only prepared if the graph is empty"))
	}
	g.valuelist = make([]values.Value, 0, size)
	g.vhash2indexPlus1 = make(map[vhash]Virtual, size)
}

// PrepareVirtualSpace reserves memory for 3 times size hashes of 16 bytes, it panics if tha Graph is not empty
func (g *Graph) PrepareVirtualSpace(size int) {
	if len(g.dataset) != 0 {
		panic(errors.New("virtual space can only prepared if the graph is empty"))
	}
	g.dataset = make(map[Triple]*Triple, size)
}

// AddValueString adds a string Literal to the Values. See AddValue for Details.
func (g *Graph) AddValueString(s string) VirtualValue {
	return g.AddValue(values.LiteralString(s))
}

// AddValueIRI adds an IRI to the Values. See AddValue for Details.
func (g *Graph) AddValueIRI(iri string) VirtualValue {
	return g.AddValue(values.IRIParse(iri))
}

// AddValue adds an unknown primitive to the graph's values. Returns either this new VirtualValue or the known one.
func (g *Graph) AddValue(v values.Value) VirtualValue {
	vv := g.VirtualValue(v)
	if !vv.Known {
		g.valuelist = append(g.valuelist, vv.Value)
		vv.Virtual = Virtual(len(g.valuelist))
		g.vhash2indexPlus1[vv.vhash] = vv.Virtual
		g.valuesize += vv.Size + 8 + 16 + 4
	}
	return vv
}

// VirtualValue retrieves a primitive keyed by its hash from the graph or creates a new VirtualValue
func (g *Graph) VirtualValue(v values.Value) VirtualValue {
	tb, b := v.Serialize()
	return vvalue(g, tb, b, v)
}

// BulkAddTriple adds a Triple without checking to the Graph, see Assert instead
func (g *Graph) BulkAddTriple(t *Triple, withindex bool) *Triple {
	g.dataset[*t] = t
	if withindex {
		t = g.dataset[*t]
		g.sindex[t.S] = append(g.sindex[t.S], t)
		g.pindex[t.P] = append(g.pindex[t.P], t)
		g.oindex[t.O] = append(g.oindex[t.O], t)
		g.spindex[t.SP()] = append(g.spindex[t.SP()], t)
		g.soindex[t.SO()] = append(g.soindex[t.SO()], t)
		g.poindex[t.PO()] = append(g.poindex[t.PO()], t)
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
		for t := range g.dataset {
			sindex[t.S]++
		}
		wg.Done()
	}()

	go func() {
		for t := range g.dataset {
			pindex[t.P]++
		}
		wg.Done()
	}()

	go func() {
		for t := range g.dataset {
			oindex[t.O]++
		}
		wg.Done()
	}()

	go func() {
		for t := range g.dataset {
			spindex[t.SP()]++
		}
		wg.Done()
	}()

	go func() {
		for t := range g.dataset {
			soindex[t.SO()]++
		}
		wg.Done()
	}()

	go func() {
		for t := range g.dataset {
			poindex[t.PO()]++
		}
		wg.Done()
	}()

	wg.Wait()
	// log.Println("index count done")

	wg.Add(6)

	go func() {
		g.sindex = make(map[Virtual][]*Triple, len(sindex))
		for k, v := range sindex {
			g.sindex[k] = make([]*Triple, 0, v)
			// log.Println("sindex", v, k.Value(g))
		}
		for _, t := range g.dataset {
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
		for _, t := range g.dataset {
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
		for _, t := range g.dataset {
			g.oindex[t.O] = append(g.oindex[t.O], t)
		}
		// log.Println("oindex done", len(g.oindex))
		wg.Done()
	}()

	go func() {
		g.spindex = make(map[Virtual2][]*Triple, len(spindex))
		for k, v := range spindex {
			g.spindex[k] = make([]*Triple, 0, v)
		}
		for _, t := range g.dataset {
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
		for _, t := range g.dataset {
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
		for _, t := range g.dataset {
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
	ds := (3*4 + 8) * len(g.dataset)

	idx := 0
	idx += calcsize1(g.sindex)
	idx += calcsize1(g.pindex)
	idx += calcsize1(g.oindex)
	idx += calcsize2(g.spindex)
	idx += calcsize2(g.poindex)
	idx += calcsize2(g.soindex)

	return ds, g.valuesize, idx
}

func calcsize1(index map[Virtual][]*Triple) int {
	l := (3 * 4) * len(index)
	for _, list := range index {
		l += 8 * len(list)
	}
	return l
}

func calcsize2(index map[Virtual2][]*Triple) int {
	l := (3 * 4) * len(index)
	for _, list := range index {
		l += 8 * len(list)
	}
	return l
}

// Indexed locates a list of Triples that match a given triple of patterns
func Indexed(g *Graph, t *Triple, givens, givenp, giveno, spsame, sosame, posame bool) []*Triple {
	// 3 crit, no variables
	if givens && givenp && giveno {
		if _, ok := g.dataset[*t]; ok {
			return []*Triple{t}
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
		ts := g.sindex[t.S]
		if posame {
			ts = TripleFilter(ts, func(t *Triple) bool { return t.P == t.O })
		}
		return ts
	}
	if givenp {
		ts := g.pindex[t.P]
		if sosame {
			ts = TripleFilter(ts, func(t *Triple) bool { return t.S == t.O })
		}
		return ts
	}
	if giveno {
		ts := g.oindex[t.O]
		if spsame {
			ts = TripleFilter(ts, func(t *Triple) bool { return t.S == t.P })
		}
		return ts
	}

	// 0 crit, all variables
	ts := g.Dataset()
	if spsame {
		ts = TripleFilter(ts, func(t *Triple) bool { return t.S == t.P })
	}
	if sosame {
		ts = TripleFilter(ts, func(t *Triple) bool { return t.S == t.O })
	}
	if posame {
		ts = TripleFilter(ts, func(t *Triple) bool { return t.P == t.O })
	}
	return ts
}
