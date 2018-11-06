package processor

import (
	"github.com/lercher/rdf/algebra"
)

// ReducedGate forms a Receiver that passes no immediatly following equal bindings only
type ReducedGate struct {
	inner Receiver
	last  []byte
}

// NewReducedGate creates a new ReducedGate based on a Receiver
func NewReducedGate(inner Receiver) *ReducedGate {
	return &ReducedGate{
		inner: inner,
	}
}

// Receive implements Receiver
func (gate *ReducedGate) Receive(bs algebra.Binding, vs *algebra.Variables) (bool, error) {
	as := bs.Bytes()
	defer func() {
		gate.last = as
	}()
	if len(gate.last) != len(as) {
		return gate.inner(bs, vs)
	}
	for i, b := range gate.last {
		if as[i] != b {
			return gate.inner(bs, vs)
		}
	}
	return true, nil // duplicate ignored
}
