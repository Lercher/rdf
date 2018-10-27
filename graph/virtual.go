package graph

import "github.com/zhenjl/cityhash"

type Virtual cityhash.Uint128

type Virtual2 struct {
	V1, V2 Virtual
}

func (v Virtual) Value(g *Graph) Value {
	val := g.valuemap[v]
	return val
}

func hash(block []byte) Virtual {
	s := cityhash.CityHash128(block, uint32(len(block)))
	return Virtual(s)
}
