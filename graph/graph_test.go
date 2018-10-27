package graph

import "testing"

func TestGraphAsserts(t *testing.T) {
	g := New()
	g.Assert("martins", "telefon", "+49897482400")
	if len(g.Dataset) != 1 {
		t.Error("graph dataset length got", len(g.Dataset), ", want 1")
	}
}
