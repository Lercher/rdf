package graph

import "testing"

func gr() *Graph {
	g := New()
	g.Assert("martin", "telefon", "+49897482400")
	g.Assert("andreas", "telefon", "+49897482400")
	g.Assert("martin", "boss", "justus")
	g.Assert("martin", "telefon", "+49897482400")
	return g
}

func TestGraphLength(t *testing.T) {
	g := gr()
	if len(g.Dataset) != 3 {
		t.Fatal("graph dataset length want 3, got", len(g.Dataset))
	}
}

func TestGraphSimpleProps(t *testing.T) {
	g := gr()

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

func TestTriple(t *testing.T) {
	g := gr()
	l := len(g.Dataset)
	c := g.CountValues()
	tr0 := NewTriple(g, 0, 0, 0)
	trfloat := NewTriple(g, float64(0), float64(0), float64(0))
	trmixed := NewTriple(g, 0, float64(0), "0")
	if len(g.Dataset) != l {
		t.Fatal("graph dataset length want ", l, ", got", len(g.Dataset))
	}
	if g.CountValues() != c+3 {
		t.Errorf("want %d distinct values, got %d",c+3, g.CountValues())
	}
	for _, v := range g.Values() {
		t.Logf("%T: %[1]v", v)
	}
	int0 := tr0.O.Value(g)
	if int0 != int(0) {
		t.Errorf("want int 0, got %T %[1]v", int0)
	}
	fl0 := trfloat.O.Value(g)
	if fl0 != float64(0) {
		t.Errorf("want float 0, got %T %[1]v", fl0)
	}
	s0 := trmixed.O.Value(g)
	if s0 != "0" {
		t.Errorf("want string 0, got %T %[1]v", s0)
	}
}
