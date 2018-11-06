package algebra

import (
	"encoding/binary"
	"bytes"
	"github.com/lercher/rdf/graph"
)

// Binding items bind a Variable (index into Variables) to a Virtual (index into Values)
type Binding []graph.Virtual

// NewBinding returns Bindings that are all IsNil
func NewBinding(n int) Binding {
	return make(Binding, n)
}

// Set initializes the first elements from ints. 
// Useful for testing, but hardly for anything else.
func (bs Binding) Set(vs ...int) Binding {
	for i, v := range vs {
		if i < len(bs) {
			bs[i] = graph.Virtual(v)
		}
	}
	return bs
}

// Bytes is a byteslice containing all values
func (bs Binding) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, bs)
	return buf.Bytes()
}

// IsUnbound is true if the given Variable has not yet a Value
func (bs Binding) IsUnbound(v Variable) bool {
	b := bs[v]
	return b.IsNil()
}

// HasNil is true if one item IsNil
func (bs Binding) HasNil() bool {
	for _, b := range bs {
		if b.IsNil() {
			return true
		}
	}
	return false
}

// Clone makes a copy of this Binding
func (bs Binding) Clone() Binding {
	clone := make(Binding, len(bs))
	copy(clone, bs)
	return clone
}

// BindTriple makes a copy of this Binding, replacing Variables with Values
func (bs Binding) BindTriple(tp *TriplePattern, t *graph.Triple) Binding {
	set := func(clone Binding, p *Pattern, v graph.Virtual) {
		if p.IsVariable() {
			if !p.IsAnonymous() {
				clone[p.Variable()] = v
			}
		}
	}

	clone := bs.Clone()
	set(clone, tp.S, t.S)
	set(clone, tp.P, t.P)
	set(clone, tp.O, t.O)
	return clone
}

// Materialize retrieves names and values of bound variables
func (bs Binding) Materialize(g *graph.Graph, variables *Variables) (list []*Materialized) {
	for i, b := range bs {
		vn := variables.NameOf(Variable(i))
		val := b.Value(g)
		list = append(list, &Materialized{Name: vn, Value: val})
	}
	return list
}

// BindTerm transforms an unbound Term to a Literal Term
func (bs Binding) BindTerm(t Term) (Term, graph.Virtual) {
	if b, ok := t.(binder); ok {
		return b.bind(bs)
	}
	return t, graph.NotAVirtual
}

type binder interface {
	bind(bs Binding) (Term, graph.Virtual)
}
