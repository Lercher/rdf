package binary

import (
	"bytes"
	"testing"
)

const (
	emptyGraph  = "RDFGRA01\000\000\000\000\000\000\000\000RDFVAL01\000\000\000\000\000\000\000\000"
	simpleGraph = "RDFGRA01\x03\x00\x00\x00\x00\x00\x00\x00\xb9a\x1c\xd3˅\xf1U.Q\xdaE\xe3\x8csׇ\xb4\x1d\xca\xca\xf4\x9e\xaf\x9b\x14'1\xf5\x11\xfdG\xa7\x92\xa4\xee]\x06}1\xacU\x1f\x82s\xbfP\xe06\x06\xa6\xc3&\x9e2\x15X\t?\xf5\\\xac&u\x87\xb4\x1d\xca\xca\xf4\x9e\xaf\x9b\x14'1\xf5\x11\xfdG\xa7\x92\xa4\xee]\x06}1\xacU\x1f\x82s\xbfP\xe0\xb9a\x1c\xd3˅\xf1U.Q\xdaE\xe3\x8cs\xd7\xf3\x9f\xc1\xba \xca\xdf\xfb\x82K\x80=\xe40\x1f+b\xe2\xbeWq\f\xaa\n=\xa8C\xd6x\n\xb6\fRDFVAL01\x06\x00\x00\x00\x00\x00\x00\x00S\x06\x00\x00\x00\x00\x00\x00\x00justusS\x06\x00\x00\x00\x00\x00\x00\x00martinS\a\x00\x00\x00\x00\x00\x00\x00telefonS\f\x00\x00\x00\x00\x00\x00\x00+49897482400S\a\x00\x00\x00\x00\x00\x00\x00andreasS\x04\x00\x00\x00\x00\x00\x00\x00boss"
	mixedGraph  = "RDFGRA01\x03\x00\x00\x00\x00\x00\x00\x00\xb9a\x1c\xd3˅\xf1U.Q\xdaE\xe3\x8csׅ\x1a\x92\xe4\x96L0\x0f\xa1\x96pBڄ\xf8\xbdaP\x00r\x8f55M\x022^\xbb\xe6\xb15\xe16\x06\xa6\xc3&\x9e2\x15X\t?\xf5\\\xac&u#%\x95C\xf2\r\xd2e\x8a\xff\xe6\xfa\xb1\xb5\xd6\x16\x17\xc4\xc3\xf2\xac<\xa5\x00\xaaIY1I\xe8\xa7%\xb6L\x87/\xc3F\xe5\f\xcc\xc1y]\x93\xccL\xef* e}\xccZ~\xf5o_\xfe\xc2\xf6$\xaf\f\xa6\x86i,Ζ\x0e\xe2e\xd3\xf3\x93n\a\x9e\x9fRDFVAL01\t\x00\x00\x00\x00\x00\x00\x00S\a\x00\x00\x00\x00\x00\x00\x00andreasS\x06\x00\x00\x00\x00\x00\x00\x00augustFh1\xc2%\x91xj\xc0S\x06\x00\x00\x00\x00\x00\x00\x00martinI{\x00\x00\x00\x00\x00\x00\x00F\x87\x88\x9bS\xc9\xc0\xf3?I\xbf\xfe\xff\xff\xff\xff\xff\xffF\x10\x06\x9e{\x8fW\x93@I\x00\x00\x00\x00\x00\x00\x00\x00"
)

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
		t.Errorf("want values %d len, got %d", 6, g.CountValues())
	}
	
	if len(g.Dataset) != 3 {
		t.Fatalf("want dataset %d len, got %d", 3, len(g.Dataset))
	}
	
	t.Log(g.Dataset[0].S.Bytes(), g.Dataset[0].S.String(g))
	s0 := g.Dataset[0].String(g)
	want0 := "[(string:martin) (string:telefon) (string:+49897482400)]"
	if s0 != want0 {
		t.Errorf("want triple0 %s, got %s", want0, s0)
	}
}

func TestReadMixedGraph(t *testing.T) {
	buf := bytes.NewBuffer([]byte(mixedGraph))
	dec := NewDecoder(buf)
	g, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	
	for _, tr := range g.Dataset {
		t.Log(tr.String(g))
	}
	
	if g.CountValues() != 9 {
		t.Errorf("want values %d len, got %d", 9, g.CountValues())
	}
	
	if len(g.Dataset) != 3 {
		t.Fatalf("want dataset %d len, got %d", 3, len(g.Dataset))
	}

	s2 := g.Dataset[2].String(g)
	want2 := "[(string:august) (int:0) (float64:-211.7677182)]"
	if s2 != want2 {
		t.Errorf("want triple2 %s, got %s", want2, s2)
	}
}
