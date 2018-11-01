package algebra

import (
	"testing"

	"github.com/lercher/rdf/graph"
)

func TestVerb(t *testing.T) {
	v := Verb(graph.IRIParse("<http://IRI>"))

	got := DescribeVerb(v)
	want := "graph.IRI <http://IRI>"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
