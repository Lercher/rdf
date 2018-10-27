package graph

import (
	"fmt"

	"github.com/zhenjl/cityhash"
)

// Virtual is the datatype for virtualizing values inside a Graph. 
// It consists of 16 bytes and is formed from the Cityhash of a binary representation
// of a Value 
type Virtual cityhash.Uint128

// Virtual2 is a pair of Virtuals
type Virtual2 struct {
	V1, V2 Virtual
}

// IsNil holds true, if the underlying datatype is zero, i.e. 16 zero bytes
func (v Virtual) IsNil() bool {
	return cityhash.Uint128(v).Higher64() == 0 && cityhash.Uint128(v).Lower64() == 0
}

func (v Virtual) String(g *Graph) string {
	val := v.Value(g)
	return fmt.Sprintf("(%T:%[1]v)", val)
}

// Value returns the real Value of a virtualized Triple part in subject/predicate/object position
// according to the records of the creating Graph. It is undefined behaviour to
// retrieve a Virtual's Value from a different Graph
func (v Virtual) Value(g *Graph) Value {
	val := g.valuemap[v]
	return val
}

func hash(block []byte) Virtual {
	s := cityhash.CityHash128(block, uint32(len(block)))
	return Virtual(s)
}
