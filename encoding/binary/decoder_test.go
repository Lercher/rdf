package binary

import (
	"bytes"
	"testing"
)

const (
	emptyGraph  = "RDFVAL01\000RDFGRA01\000"
	simpleGraph = "RDFVAL01\fl\fmartin\x00\x00I\x0etelefonl\x18+49897482400\x00\x00l\x0eandreas\x00\x00I\bbossl\fjustus\x00\x00RDFGRA01\x06\x02\x04\x06\b\x04\x06\x02\n\f"
	mixedGraph  = "RDFVAL01\x12l\fmartin\x00\x00i\xf6\x01f\x87\x91\ue715\x99\xf0\xf9?l\x0eandreas\x00\x00i\x81\x05f\x90\x8c\xf8\xdc\xf7\xf1\xd5\xc9@l\faugust\x00\x00i\x00f\xe8∮\x92\x92\x9e\xb5\xc0\x01RDFGRA01\x06\x02\x04\x06\b\n\f\x0e\x10\x12"
)

func TestReadEmpty(t *testing.T) {
	buf := bytes.NewBuffer([]byte(emptyGraph))
	dec := NewDecoder(buf)
	g, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if g.CountTriples() != 0 {
		t.Errorf("want dataset %d len, got %d", 0, g.CountTriples())
	}

	if g.CountValues() != 0 {
		t.Errorf("want values %d len, got %d", 0, g.CountValues())
	}
}

func TestReadSimpleGraph(t *testing.T) {
	buf := bytes.NewBuffer([]byte(simpleGraph))
	dec := NewDecoder(buf)
	g, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}

	for tr := range g.DatasetMap() {
		t.Log(tr.String(g))
	}

	if g.CountValues() != 6 {
		t.Errorf("want values %d len, got %d", 6, g.CountValues())
	}

	if g.CountTriples() != 3 {
		t.Fatalf("want dataset %d len, got %d", 3, g.CountTriples())
	}

	want := `[(Literal:"andreas") (IRI:<telefon>) (Literal:"+49897482400")]`
	for tr := range g.DatasetMap() {
		t.Log(tr.S, tr.S.String(g))
		s := tr.String(g)
		if s == want {
			return
		}
	}
	t.Errorf("want triple %s, but not found", want)
}

func TestReadMixedGraph(t *testing.T) {
	buf := bytes.NewBuffer([]byte(mixedGraph))
	dec := NewDecoder(buf)
	g, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}

	for tr := range g.DatasetMap() {
		t.Log(tr.String(g))
	}

	if g.CountValues() != 9 {
		t.Errorf("want values %d len, got %d", 9, g.CountValues())
	}

	if g.CountTriples() != 3 {
		t.Fatalf("want dataset %d len, got %d", 3, g.CountTriples())
	}

	want := `[(Literal:"august") (Int:0) (Float:-211.7677182)]`
	for tr := range g.DatasetMap() {
		t.Log(tr.S, tr.S.String(g))
		s := tr.String(g)
		if s == want {
			return
		}
	}
	t.Errorf("want triple %s, but not found", want)
}
