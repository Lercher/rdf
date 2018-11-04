package graph

import (
	"github.com/lercher/rdf/values"
)

// VirtualValue pairs a Virtual (i.e. the hash value) with its real Value
type VirtualValue struct {
	vhash
	Virtual
	values.Value
	Size  int
	Known bool
}

func vvalue(g *Graph, seed byte, bytes []byte, v values.Value) VirtualValue {
	val := v
	h := hash(seed, bytes)
	idx1, ok := g.vhash2indexPlus1[h]
	if ok {
		val = g.valuelist[idx1-1]
	}
	return VirtualValue{
		vhash:   h,
		Virtual: idx1,
		Value:   val,
		Known:   ok,
		Size:    len(bytes),
	}
}
