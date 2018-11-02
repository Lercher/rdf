package algebra

import (
	"testing"

	"github.com/lercher/rdf/graph"
)

func TestBinding(t *testing.T) {
	bs := NewBinding(3)
	if len(bs) != 3 {
		t.Fatalf("wanted %d bindings, got %d", 3, len(bs))
	}
	for i, b := range bs {
		if !b.IsNil() {
			t.Errorf("wanted all Bindings nil, but [%d] is not: %v", i, b)
		}
	}
	if !bs.HasNil() {
		t.Errorf("wanted Bindings HasNil true, got false")
	}

	clone := bs.Clone()
	for i := range clone {
		clone[i] = graph.Virtual(10+i) // not Nil
	}
	if clone.HasNil() {
		t.Errorf("wanted cloned Bindings not HasNil, got true")
	}
	for i, b := range clone {
		if b.IsNil() {
			t.Errorf("wanted all cloned Bindings not nil, but [%d] is: %v", i, b)
		}
	}
	for i, b := range bs {
		if !b.IsNil() {
			t.Errorf("wanted all Bindings after clone nil, but [%d] is not: %v", i, b)
		}
	}
}
