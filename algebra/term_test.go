package algebra

import (
	"testing"

	"github.com/lercher/rdf/values"
)

func TestTerm(t *testing.T) {
	term := Term(values.IRIParse("<http://IRI>"))

	got := DescribeTerm(term)
	want := "(IRI:<http://IRI>)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
