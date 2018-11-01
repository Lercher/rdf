package algebra

import (
	"testing"

	"github.com/lercher/rdf/graph"
)

func TestTerm(t *testing.T) {
	term := Term(graph.IRIParse("<http://IRI>"))

	got := DescribeTerm(term)
	want := "(graph.IRI <http://IRI>)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
