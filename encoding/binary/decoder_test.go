package binary

import (
	"bytes"
	"testing"
)

const (
	emptyGraph  = "RDFGRA01\000\000\000\000\000\000\000\000RDFVAL01\000\000\000\000\000\000\000\000"
	simpleGraph = "RDFGRA01\x03\x00\x00\x00\x00\x00\x00\x00Z\xe4\x94?\x963tp|*\xe9\x1d\xd1\x0f8\x9dk(\x02\xfb\xf6\xf2vr\xd4j!\nH\xf7\x8f\x83ƣ\fM,zf\xef\x82\xea\xdc\x17\xb5d:\x82\xee\xf9\x81x\x04\xb9\xc7~&\x04\x80\x9d#\xb7W(k(\x02\xfb\xf6\xf2vr\xd4j!\nH\xf7\x8f\x83ƣ\fM,zf\xef\x82\xea\xdc\x17\xb5d:\x82Z\xe4\x94?\x963tp|*\xe9\x1d\xd1\x0f8\x9d\x1fu\xae&)\x00@ \xb7k\as\xee\xad\x1a#'\xe4#\x1b\x8c\u007f\x93\xc2?F\xf8:\x8e\xac%\xe0RDFVAL01\x06\x00\x00\x00\x00\x00\x00\x00S\x04\x00\x00\x00\x00\x00\x00\x00bossS\x06\x00\x00\x00\x00\x00\x00\x00justusS\x06\x00\x00\x00\x00\x00\x00\x00martinS\a\x00\x00\x00\x00\x00\x00\x00telefonS\f\x00\x00\x00\x00\x00\x00\x00+49897482400S\a\x00\x00\x00\x00\x00\x00\x00andreas"
	mixedGraph  = "RDFGRA01\x03\x00\x00\x00\x00\x00\x00\x00Z\xe4\x94?\x963tp|*\xe9\x1d\xd1\x0f8\x9d3\x89*\x9ec\x9c\xb5\\\xf1nPBN\xff\xf2\xee\x81ؔ&o\xb0\x86\x14r->2\x94\xa3\xd9\b\xee\xf9\x81x\x04\xb9\xc7~&\x04\x80\x9d#\xb7W(\xe3\x01\xd0\x16=EQD<\xb3\xb7ׅ:\xb7U\xd9]o܃\x1a\u007f\x8a\xff\xccٴ\x96\x15w\xf1\x06\x05\xa7\xa8\xf0\xd6\x0fs\x81\x9d\n\xec&\x1a1\xa2\x0f\xb1\f\xb1\xb0R\x91q$\x01k\x9e8\xcf\xd1If\xea\x99a\xc1\x1czm\x1a\x8e\xef;\xe6\x8dD\x82RDFVAL01\t\x00\x00\x00\x00\x00\x00\x00Fh1\xc2%\x91xj\xc0F\x87\x88\x9bS\xc9\xc0\xf3?S\a\x00\x00\x00\x00\x00\x00\x00andreasI\xbf\xfe\xff\xff\xff\xff\xff\xffS\x06\x00\x00\x00\x00\x00\x00\x00augustI\x00\x00\x00\x00\x00\x00\x00\x00S\x06\x00\x00\x00\x00\x00\x00\x00martinI{\x00\x00\x00\x00\x00\x00\x00F\x10\x06\x9e{\x8fW\x93@"
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
