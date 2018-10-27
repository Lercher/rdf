package binary

import (
	"bytes"
	"testing"
)

const emptyGraph = "RDFGRA01\000\000\000\000\000\000\000\000RDFVAL01\000\000\000\000\000\000\000\000"
const simpleGraph = "RDFGRA01\x03\x00\x00\x00\x00\x00\x00\x00Z\xe4\x94?\x963tp|*\xe9\x1d\xd1\x0f8\x9dk(\x02\xfb\xf6\xf2vr\xd4j!\nH\xf7\x8f\x83ƣ\fM,zf\xef\x82\xea\xdc\x17\xb5d:\x82\xee\xf9\x81x\x04\xb9\xc7~&\x04\x80\x9d#\xb7W(k(\x02\xfb\xf6\xf2vr\xd4j!\nH\xf7\x8f\x83ƣ\fM,zf\xef\x82\xea\xdc\x17\xb5d:\x82Z\xe4\x94?\x963tp|*\xe9\x1d\xd1\x0f8\x9d\x1fu\xae&)\x00@ \xb7k\as\xee\xad\x1a#'\xe4#\x1b\x8c\u007f\x93\xc2?F\xf8:\x8e\xac%\xe0RDFVAL01\x06\x00\x00\x00\x00\x00\x00\x00\x00\x06\x00\x00\x00\x00\x00\x00\x00martin\x00\a\x00\x00\x00\x00\x00\x00\x00telefon\x00\f\x00\x00\x00\x00\x00\x00\x00+49897482400\x00\a\x00\x00\x00\x00\x00\x00\x00andreas\x00\x04\x00\x00\x00\x00\x00\x00\x00boss\x00\x06\x00\x00\x00\x00\x00\x00\x00justus"

func TestReadEmpty(t *testing.T) {
	buf := bytes.NewBuffer([]byte(emptyGraph))
	dec := NewDecoder(buf)
	g, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if g.Dataset == nil {
		t.Errorf("want dataset not nil, got nil")
	}
	if len(g.Dataset) != 0 {
		t.Errorf("want dataset %d len, got %d", 0, len(g.Dataset))
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
	for _, tr := range g.Dataset {
		t.Log(tr.String(g))
	}
	if g.CountValues() != 6 {
		t.Errorf("want values %d len, got %d", 0, g.CountValues())
	}
	if len(g.Dataset) != 3 {
		t.Fatalf("want dataset %d len, got %d", 0, len(g.Dataset))
	}
	t.Log(g.Dataset[0].S.Bytes(), g.Dataset[0].S.String(g))
	s0 := g.Dataset[0].String(g)
	want0 := "[(string:martin) (string:telefon) (string:+49897482400)]"
	if s0 != want0 {
		t.Errorf("want triple0 %s, got %s", want0, s0)
	}
}
