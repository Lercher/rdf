package binary

import (
	"bytes"
	"testing"

	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/values"
)

func grString() *graph.Graph {
	g := graph.New()
	g.AssertLiterally("martin", "telefon", "+49897482400")
	g.AssertLiterally("andreas", "telefon", "+49897482400")
	g.AssertLiterally("martin", "boss", "justus")
	return g
}

func grMixed() *graph.Graph {
	g := graph.New()
	g.Assert(values.LiteralString("martin"), values.Int(123), values.Float(1.234567))
	g.Assert(values.LiteralString("andreas"), values.Int(-321), values.Float(1237.89012))
	g.Assert(values.LiteralString("august"), values.Int(0), values.Float(-211.7677182))
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
	sbytes := []byte(simpleGraph)
	t.Logf("%q", bytes)
	if len(bytes) != len(sbytes) {
		t.Errorf("want %d bytes written, got %d", len(sbytes), len(bytes))
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
