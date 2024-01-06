package processor

import (
	"testing"

	"github.com/lercher/rdf/algebra"
)

func TestDistinctBinding(t *testing.T) {
	b1 := algebra.NewBinding(3).Set(1, 2, 3)

	b1clone := b1.Clone()
	b2 := b1.Clone().Set(4)

	n := 0
	testr := func(bs algebra.Binding, vs *algebra.Variables) (bool, error) {
		n++
		t.Log(bs)
		return true, nil
	}
	gate := NewDistinctGate(testr)

	gate.Receive(b1, nil)
	gate.Receive(b1clone, nil)
	gate.Receive(b2, nil)
	gate.Receive(b1, nil)
	gate.Receive(b1clone, nil)
	gate.Receive(b2, nil)

	if n != 2 {
		t.Errorf("want %d distincts, got %d", 2, n)
	}
}
