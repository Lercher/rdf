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

func TestPrefixedIRI(t *testing.T) {
	pis := []*PrefixedIRI{
		&PrefixedIRI{"o", "http://ontology#"},
		&PrefixedIRI{"s", "http://subjects/"},
	}
	color, err := IRIPrefixed("o:color", pis)
	if err != nil {
		t.Errorf("want no error for o:color, got %v", err)
	} else if color.String() != "<http://ontology#color>" {
		t.Errorf("want <http://ontology#color> for o:color, got %v", color)
	}

	s5, err := IRIPrefixed("s:5", pis)
	if err != nil {
		t.Errorf("want no error for s:5, got %v", err)
	} else if s5.String() != "<http://subjects/5>" {
		t.Errorf("want <http://subjects/5> for s:5, got %v", s5)
	}

	_, err = IRIPrefixed("illegal/iri", pis)
	if err == nil {
		t.Errorf("want error for illegal/iri")
	}
	_, err = IRIPrefixed("unknown:iri", pis)
	if err == nil {
		t.Errorf("want error for unknown:iri")
	}
}
