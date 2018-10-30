package graph

import "testing"

const (
	iri = `<http://identifiers>`
	uri = `http://identifiers`
)

func TestIRI(t *testing.T) {
	i0 := IRI(uri)
	i1 := IRIParse(iri)
	i2 := IRIParse(uri)

	if i0.String() != iri {
		t.Errorf("i0 String: want %q, got %q", iri, i0.String())
	}
	if i1.String() != iri {
		t.Errorf("i1 String: want %q, got %q", iri, i1.String())
	}
	if i2.String() != iri {
		t.Errorf("i2 String: want %q, got %q", iri, i2.String())
	}

	if i0.URI() != uri {
		t.Errorf("i0 URI: want %q, got %q", uri, i0.URI())
	}
}