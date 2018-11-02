package processor

import (
	"testing"

	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
)

func gr() (*graph.Graph, *graph.Triple, *graph.Triple, *graph.Triple) {
	g := graph.New()
	t0 := g.Assert("martin", "telefon", "+49897482400")
	t1 := g.Assert("andreas", "telefon", "+49897482400")
	t2 := g.Assert("martin", "boss", "justus")
	g.Assert("martin", "telefon", "+49897482400")
	return g, t0, t1, t2
}

func TestPatterns0Criteria(t *testing.T) {
	g, _, _, _ := gr()

	all := Match(g, &algebra.TriplePattern{})
	if len(all) != g.CountTriples() {
		t.Error("want full length", g.CountTriples(), "got", len(all))
	}
	t.Log()
}

func TestPatterns1Criteria(t *testing.T) {
	g, t0, _, _ := gr()

	s0 := t0.S
	allS := &algebra.TriplePattern{S: algebra.PatternLiteral(s0)}
	t.Log("match", allS.String(g))
	ms := Match(g, allS)
	for _, tr := range ms {
		t.Log(tr.String(g))
	}
	if len(ms) != 2 {
		t.Error("want length 2, got", len(ms))
	}
	t.Log()

	p0 := t0.P
	allP := &algebra.TriplePattern{P: algebra.PatternLiteral(p0)}
	t.Log("match", allP.String(g))
	mp := Match(g, allP)
	for _, tr := range mp {
		t.Log(tr.String(g))
	}
	if len(mp) != 2 {
		t.Error("want length 2, got", len(mp))
	}
	t.Log()

	o0 := t0.O
	allO := &algebra.TriplePattern{O: algebra.PatternLiteral(o0)}
	t.Log("match", allO.String(g))
	mo := Match(g, allO)
	for _, tr := range mo {
		t.Log(tr.String(g))
	}
	if len(mo) != 2 {
		t.Error("want length 2, got", len(mo))
	}
	t.Log()
}

func TestPatterns2Criteria(t *testing.T) {
	g, t0, _, _ := gr()

	s0 := t0.S
	p0 := t0.P
	o0 := t0.O

	allSP := &algebra.TriplePattern{S: algebra.PatternLiteral(s0), P: algebra.PatternLiteral(p0)}
	t.Log("match", allSP.String(g))
	msp := Match(g, allSP)
	for _, tr := range msp {
		t.Log(tr.String(g))
	}
	if len(msp) != 1 {
		t.Error("want length 1, got", len(msp))
	}
	t.Log()

	allSO := &algebra.TriplePattern{S: algebra.PatternLiteral(s0), O: algebra.PatternLiteral(o0)}
	t.Log("match", allSO.String(g))
	mso := Match(g, allSO)
	for _, tr := range mso {
		t.Log(tr.String(g))
	}
	if len(mso) != 1 {
		t.Error("want length 1, got", len(mso))
	}
	t.Log()

	allPO := &algebra.TriplePattern{P: algebra.PatternLiteral(p0), O: algebra.PatternLiteral(o0)}
	t.Log("match", allPO.String(g))
	mpo := Match(g, allPO)
	for _, tr := range mpo {
		t.Log(tr.String(g))
	}
	if len(mpo) != 2 {
		t.Error("want length 2, got", len(mpo))
	}
	t.Log()
}

func TestPatterns3Criteria(t *testing.T) {
	g, t0, _, _ := gr()

	allSPO := &algebra.TriplePattern{S: algebra.PatternLiteral(t0.S), P: algebra.PatternLiteral(t0.P), O: algebra.PatternLiteral(t0.O)}
	t.Log("match", allSPO.String(g))
	mspo := Match(g, allSPO)
	for _, tr := range mspo {
		t.Log(tr.String(g))
	}
	if len(mspo) != 1 {
		t.Error("want length 1, got", len(mspo))
	}
	t.Log()
}
