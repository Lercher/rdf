package algebra

import (
	"github.com/lercher/rdf/graph"
	"testing"
)

func TestPatterns(t *testing.T) {
	pnil := (*Pattern)(nil)
	want(t, "nil is annonymous", pnil.IsAnonymous(), true)
	want(t, "nil is variable", pnil.IsVariable(), true)
	want(t, "nil won't match", pnil.WontMatch(), false)

	pa := PatternAnonymous()
	want(t, "PatternAnonymous is annonymous", pa.IsAnonymous(), true)
	want(t, "PatternAnonymous is variable", pa.IsVariable(), true)
	want(t, "pa won't match", pa.WontMatch(), false)

	wont := PatternWontMatch()
	want(t, "p0 is annonymous", wont.IsAnonymous(), false)
	want(t, "p0 is variable", wont.IsVariable(), false)
	want(t, "p0 won't match", wont.WontMatch(), true)

	p0 := PatternVariable(Variable(0))
	want(t, "p0 is annonymous", p0.IsAnonymous(), false)
	want(t, "p0 is variable", p0.IsVariable(), true)
	want(t, "p0 won't match", p0.WontMatch(), false)

	p1 := PatternVariable(Variable(1))
	want(t, "p1 is annonymous", p1.IsAnonymous(), false)
	want(t, "p1 is variable", p1.IsVariable(), true)
	want(t, "p1 won't match", p1.WontMatch(), false)

	ps := PatternLiteral(graph.Virtual(1))
	want(t, "ps is annonymous", ps.IsAnonymous(), false)
	want(t, "ps is variable", ps.IsVariable(), false)
	want(t, "ps won't match", ps.WontMatch(), false)

	p1b := PatternVariable(Variable(1))
	p2 := PatternVariable(Variable(2))
	pq := PatternLiteral(graph.Virtual(2))
	want(t, "nil and PatternAnonymous are same", PatternOfSameVar(pnil, pa), false)
	want(t, "nil and p0 are same", PatternOfSameVar(pnil, p0), false)
	want(t, "p0 and nil are same", PatternOfSameVar(p0, pnil), false)
	want(t, "p0 and p0 are same", PatternOfSameVar(p0, p0), true)
	want(t, "nil and p1 are same", PatternOfSameVar(pnil, p1), false)
	want(t, "p1 and nil are same", PatternOfSameVar(p1, pnil), false)
	want(t, "nil and p2 are same", PatternOfSameVar(pnil, p2), false)
	want(t, "p1 and p2 are same", PatternOfSameVar(p1, p2), false)
	want(t, "p1 and p1' are same", PatternOfSameVar(p1, p1b), true)
	want(t, "p2 and p2 are same", PatternOfSameVar(p2, p2), true)
	want(t, "p1 and ps are same", PatternOfSameVar(p1, ps), false)
	want(t, "ps and pq are same", PatternOfSameVar(ps, pq), false)
	want(t, "nil and pq are same", PatternOfSameVar(pnil, pq), false)
	want(t, "ps and pnil are same", PatternOfSameVar(ps, pnil), false)
}

func want(t *testing.T, what string, got, want bool) {
	t.Helper()
	if got == want {
		return
	}
	t.Errorf("%s: want %v, got %v", what, want, got)
}
