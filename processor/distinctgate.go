package processor

import (
	"github.com/zhenjl/cityhash"
	"github.com/lercher/rdf/algebra"
)

// DistinctGate forms a Receiver that passes unique bindings only
type DistinctGate struct {
	inner Receiver
	m map[cityhash.Uint128]bool
}

// NewDistinctGate creates a new DistinctGate based on a Receiver
func NewDistinctGate(inner Receiver) *DistinctGate {
	return &DistinctGate{
		inner: inner,
		m: make(map[cityhash.Uint128]bool),
	}
}

// Receive implements Receiver
func (gate *DistinctGate) Receive(bs algebra.Binding, vs *algebra.Variables) (bool, error) {
	b := bs.Bytes()
	h := cityhash.CityHash128(b, uint32(len(b)))
	if gate.m[h] {
		return true, nil // duplicate ignored
	}
	gate.m[h] = true
	return gate.inner(bs, vs)
}
