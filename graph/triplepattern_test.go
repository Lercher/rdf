package graph

import "testing"

func TestPatterns0Criteria(t *testing.T) {
	g := gr()

	all := g.Match(&TriplePattern{})
	if len(all) != len(g.Dataset) {
		t.Error("want full length", len(g.Dataset), "got", len(all))
	}
	t.Log()
}

func TestPatterns1Criteria(t *testing.T) {
	g := gr()

	s0 := g.Dataset[0].S
	allS := &TriplePattern{S: PatternLiteral(s0)}
	t.Log("match", allS.String(g))
	ms := g.Match(allS)
	for _, tr := range ms {
		t.Log(tr.String(g))
	}
	if len(ms) != 2 {
		t.Error("want length 2, got", len(ms))
	}
	t.Log()

	p0 := g.Dataset[0].P
	allP := &TriplePattern{P: PatternLiteral(p0)}
	t.Log("match", allP.String(g))
	mp := g.Match(allP)
	for _, tr := range mp {
		t.Log(tr.String(g))
	}
	if len(mp) != 2 {
		t.Error("want length 2, got", len(mp))
	}
	t.Log()

	o0 := g.Dataset[0].O
	allO := &TriplePattern{O: PatternLiteral(o0)}
	t.Log("match", allO.String(g))
	mo := g.Match(allO)
	for _, tr := range mo {
		t.Log(tr.String(g))
	}
	if len(mo) != 2 {
		t.Error("want length 2, got", len(mo))
	}
	t.Log()
}

func TestPatterns2Criteria(t *testing.T) {
	g := gr()

	s0 := g.Dataset[0].S
	p0 := g.Dataset[0].P
	o0 := g.Dataset[0].O

	allSP := &TriplePattern{S: PatternLiteral(s0), P: PatternLiteral(p0), O: PatternVariable("$o")}
	t.Log("match", allSP.String(g))
	msp := g.Match(allSP)
	for _, tr := range msp {
		t.Log(tr.String(g))
	}
	if len(msp) != 1 {
		t.Error("want length 1, got", len(msp))
	}
	t.Log()

	allSO := &TriplePattern{S: PatternLiteral(s0), P: PatternVariable("$p"), O: PatternLiteral(o0)}
	t.Log("match", allSO.String(g))
	mso := g.Match(allSO)
	for _, tr := range mso {
		t.Log(tr.String(g))
	}
	if len(mso) != 1 {
		t.Error("want length 1, got", len(mso))
	}
	t.Log()

	allPO := &TriplePattern{S: PatternVariable("$s"), P: PatternLiteral(p0), O: PatternLiteral(o0)}
	t.Log("match", allPO.String(g))
	mpo := g.Match(allPO)
	for _, tr := range mpo {
		t.Log(tr.String(g))
	}
	if len(mpo) != 2 {
		t.Error("want length 2, got", len(mpo))
	}
	t.Log()
}

func TestPatterns3Criteria(t *testing.T) {
	g := gr()

	s0 := g.Dataset[0].S
	p0 := g.Dataset[0].P
	o0 := g.Dataset[0].O
	allSPO := &TriplePattern{S: PatternLiteral(s0), P: PatternLiteral(p0), O: PatternLiteral(o0)}
	t.Log("match", allSPO.String(g))
	mspo := g.Match(allSPO)
	for _, tr := range mspo {
		t.Log(tr.String(g))
	}
	if len(mspo) != 1 {
		t.Error("want length 1, got", len(mspo))
	}
	t.Log()
}
