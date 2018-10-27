package graph

import "testing"

func TestGraphAsserts(t *testing.T) {
	g := New()
	g.Assert("martin", "telefon", "+49897482400")
	g.Assert("andreas", "telefon", "+49897482400")
	g.Assert("martin", "boss", "justus")
	if len(g.Dataset) != 3 {
		t.Fatal("graph dataset length want 3, got", len(g.Dataset))
	}

	tr0 := g.Dataset[0]
	s := tr0.S.Value(g)
	if s != "martin" {
		t.Error("subject[0] want martin, got", s)
	}

	tr1 := g.Dataset[1]
	p := tr1.P.Value(g)
	if p != "telefon" {
		t.Error("predicate[1] want telefon, got", p)
	}

	tr2 := g.Dataset[2]
	o := tr2.O.Value(g)
	if o != "justus" {
		t.Error("object[2] want justus, got", o)
	}
}
