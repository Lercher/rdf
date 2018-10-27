package binary

import (
	"bytes"
	"testing"

	"github.com/lercher/rdf/graph"
)

func gr() *graph.Graph {
	g := graph.New()
	g.Assert("martin", "telefon", "+49897482400")
	g.Assert("andreas", "telefon", "+49897482400")
	g.Assert("martin", "boss", "justus")
	return g
}

func TestWrite(t *testing.T) {
	g := gr()
	buf := new(bytes.Buffer)
	enc := NewEncoder(buf)
	err := enc.Encode(g)
	if err != nil {
		t.Error(err)
	}
	bytes := buf.Bytes()
	if len(bytes) != 272 {
		t.Errorf("want %d bytes written, got %d", 272, len(bytes))
	}
}
