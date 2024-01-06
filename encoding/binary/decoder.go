package binary

import (
	stdbinary "encoding/binary"
	"fmt"
	"github.com/lercher/rdf/values"
	"io"

	"github.com/lercher/rdf/graph"
)

// Decoder deserializes a binary serialized graph.Graph
type Decoder struct {
	r  io.Reader
	br io.ByteReader
}

// NewDecoder creates a new Decoder for binary serialized graph.Graphs
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r, toByteReader(r)}
}

// Decode deserializes a graph.Graph
func (d *Decoder) Decode() (*graph.Graph, error) {
	g := graph.New()
	err := d.decodeValues(g)
	if err != nil {
		return nil, err
	}
	err = d.decodeDataset(g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (d *Decoder) decodeDataset(g *graph.Graph) error {
	err := checkHeader(d.r, headerDataset)
	if err != nil {
		return err
	}
	l, err := readVarint(d.br)
	if err != nil {
		return err
	}
	g.PrepareVirtualSpace(l)
	for i := 0; i < l; i++ {
		s, err := readVirtual(d.br)
		if err != nil {
			return err
		}
		p, err := readVirtual(d.br)
		if err != nil {
			return err
		}
		o, err := readVirtual(d.br)
		if err != nil {
			return err
		}
		t := &graph.Triple{S: s, P: p, O: o}
		g.BulkAddTriple(t, false)
	}
	g.RebuildIndex()
	return nil
}

func (d *Decoder) decodeValues(g *graph.Graph) error {
	err := checkHeader(d.r, headerValues)
	if err != nil {
		return err
	}
	l, err := readVarint(d.br)
	if err != nil {
		return err
	}
	g.PrepareValueSpace(l)
	for i := 0; i < l; i++ {
		val, err := values.ConstructNext(d.r)
		if err != nil {
			return err
		}
		g.AddValue(val)
	}
	return nil
}

func checkHeader(r io.Reader, s string) error {
	b := make([]byte, 8)
	n, err := io.ReadFull(r, b)
	if err != nil {
		return err
	}
	if n != len(b) {
		return fmt.Errorf("wanted to read %d bytes header, got only %d", len(b), n)
	}
	rs := string(b)
	if rs != s {
		return fmt.Errorf("expected header %q %v but got %q %v", s, []byte(s), rs, b)
	}
	return nil
}

func readVarint(br io.ByteReader) (int, error) {
	i64, err := stdbinary.ReadVarint(br)
	return int(i64), err
}

func readVirtual(br io.ByteReader) (graph.Virtual, error) {
	i, err := readVarint(br)
	return graph.Virtual(i), err
}
