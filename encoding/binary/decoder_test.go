package binary

import (
	"bytes"
	"testing"
)

const (
	emptyGraph  = "RDFVAL01\000\000\000\000RDFGRA01\000\000\000\000"
	simpleGraph = "RDFVAL01\x06\x00\x00\x00S\x06\x00\x00\x00\x00\x00\x00\x00martinS\a\x00\x00\x00\x00\x00\x00\x00telefonS\f\x00\x00\x00\x00\x00\x00\x00+49897482400S\a\x00\x00\x00\x00\x00\x00\x00andreasS\x04\x00\x00\x00\x00\x00\x00\x00bossS\x06\x00\x00\x00\x00\x00\x00\x00justusRDFGRA01\x03\x00\x00\x00\x01\x00\x00\x00\x02\x00\x00\x00\x03\x00\x00\x00\x04\x00\x00\x00\x02\x00\x00\x00\x03\x00\x00\x00\x01\x00\x00\x00\x05\x00\x00\x00\x06\x00\x00\x00"
	mixedGraph  = "RDFVAL01\t\x00\x00\x00S\x06\x00\x00\x00\x00\x00\x00\x00martinI{\x00\x00\x00\x00\x00\x00\x00F\x87\x88\x9bS\xc9\xc0\xf3?S\a\x00\x00\x00\x00\x00\x00\x00andreasI\xbf\xfe\xff\xff\xff\xff\xff\xffF\x10\x06\x9e{\x8fW\x93@S\x06\x00\x00\x00\x00\x00\x00\x00augustI\x00\x00\x00\x00\x00\x00\x00\x00Fh1\xc2%\x91xj\xc0RDFGRA01\x03\x00\x00\x00\x01\x00\x00\x00\x02\x00\x00\x00\x03\x00\x00\x00\x04\x00\x00\x00\x05\x00\x00\x00\x06\x00\x00\x00\a\x00\x00\x00\b\x00\x00\x00\t\x00\x00\x00"
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

	want := "[(string:martin) (string:telefon) (string:+49897482400)]"
	for tr := range g.DatasetMap() {
		t.Log(tr.S.Bytes(), tr.S.String(g))
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

	want := "[(string:august) (int:0) (float64:-211.7677182)]"
	for tr := range g.DatasetMap() {
		t.Log(tr.S.Bytes(), tr.S.String(g))
		s := tr.String(g)
		if s == want {
			return
		}
	}
	t.Errorf("want triple %s, but not found", want)
}
