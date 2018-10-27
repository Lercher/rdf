package binary

import (
	"bytes"
	"testing"

	"github.com/lercher/rdf/graph"
)

func grString() *graph.Graph {
	g := graph.New()
	g.Assert("martin", "telefon", "+49897482400")
	g.Assert("andreas", "telefon", "+49897482400")
	g.Assert("martin", "boss", "justus")
	return g
}

func grMixed() *graph.Graph {
	g := graph.New()
	g.Assert("martin", 123, float64(1.234567))
	g.Assert("andreas", -321, float64(1237.89012))
	g.Assert("august", 0, float64(-211.7677182))
	return g
}

func TestWriteString(t *testing.T) {
	g := grString()
	buf := new(bytes.Buffer)
	enc := NewEncoder(buf)
	err := enc.Encode(g)
	if err != nil {
		t.Error(err)
	}
	bytes := buf.Bytes()
	t.Logf("%q", bytes)
	if len(bytes) != 272 {
		t.Errorf("want %d bytes written, got %d", 272, len(bytes))
	}
}

func TestWriteMixed(t *testing.T) {
	g := grMixed()
	buf := new(bytes.Buffer)
	enc := NewEncoder(buf)
	err := enc.Encode(g)
	if err != nil {
		t.Error(err)
	}
	bytes := buf.Bytes()
	mbytes := []byte(mixedGraph)
	t.Logf("got %q", bytes)
	if len(bytes) != len(mbytes) {
		t.Errorf("want len %d, got len %d", len(mbytes), len(bytes))
	}
}
