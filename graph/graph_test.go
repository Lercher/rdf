package graph

import (
	"testing"

	"github.com/lercher/rdf/values"
)

func gr() (*Graph, *Triple, *Triple, *Triple) {
	g := New()
	t0 := g.AssertLiterally("martin", "telefon", "+49897482400")
	t1 := g.AssertLiterally("andreas", "telefon", "+49897482400")
	t2 := g.AssertLiterally("martin", "boss", "justus")
	g.AssertLiterally("martin", "telefon", "+49897482400")
	return g, t0, t1, t2
}

func TestGraphLength(t *testing.T) {
	g, _, _, _ := gr()
	if g.CountTriples() != 3 {
		t.Fatal("graph dataset length want 3, got", g.CountTriples())
	}
}

func TestGraphSimpleProps(t *testing.T) {
	g, tr0, tr1, tr2 := gr()

	s := tr0.S.Value(g)
	if s.Inner() != `martin` {
		t.Error(`subject[0] want martin, got`, s)
	}

	p := tr1.P.Value(g)
	if p.Inner() != "telefon" {
		t.Error("predicate[1] want telefon, got", p)
	}

	o := tr2.O.Value(g)
	if o.Inner() != `justus` {
		t.Error(`object[2] want justus, got`, o)
	}
}

func TestTriple(t *testing.T) {
	g, _, _, _ := gr()
	l := g.CountTriples()
	c := g.CountValues()
	tr0 := g.Assert(values.Int(0), values.Int(0), values.Int(0))
	trfloat := g.Assert(values.Float(0), values.Float(0), values.Float(0))
	trmixed := g.Assert(values.Int(0), values.Float(0), values.LiteralString("0"))
	if g.CountTriples() != l+3 {
		t.Fatalf("graph dataset length want %d, got %d", l+3, g.CountTriples())
	}
	for _, v := range g.Values() {
		t.Logf("%T: %[1]v", v)
	}
	if g.CountValues() != c+3 {
		t.Errorf("want %d distinct values, got %d", c+3, g.CountValues())
	}
	int0 := tr0.O.Value(g).Inner()
	if int0 != 0 {
		t.Errorf("want int 0, got %T %[1]v", int0)
	}
	fl0 := trfloat.O.Value(g).Inner()
	if fl0 != float64(0) {
		t.Errorf("want float 0, got %T %[1]v", fl0)
	}
	s0 := trmixed.O.Value(g).Inner()
	if s0 != "0" {
		t.Errorf("want string 0, got %T %[1]v", s0)
	}
}
