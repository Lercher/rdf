package graph

import (
	"fmt"

	"github.com/zhenjl/cityhash"
)

type Virtual cityhash.Uint128

type Virtual2 struct {
	V1, V2 Virtual
}

func (v Virtual) IsNil() bool {
	return cityhash.Uint128(v).Higher64() == 0 && cityhash.Uint128(v).Lower64() == 0
}

func (v Virtual) String(g *Graph) string {
	val := v.Value(g)
	return fmt.Sprintf("(%T:%[1]v)", val)
}

func (v Virtual) Value(g *Graph) Value {
	val := g.valuemap[v]
	return val
}

func hash(block []byte) Virtual {
	s := cityhash.CityHash128(block, uint32(len(block)))
	return Virtual(s)
}
