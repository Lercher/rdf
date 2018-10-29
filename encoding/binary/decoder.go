package binary

import (
	stdbinary "encoding/binary"
	"fmt"
	"io"

	"github.com/lercher/rdf/graph"
)

// Decoder deserializes a binary serialized graph.Graph
type Decoder struct {
	r io.Reader
}

// NewDecoder creates a new Decoder for binary serialized graph.Graphs
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}

// Decode deserializes a graph.Graph
func (d *Decoder) Decode() (*graph.Graph, error) {
	g := graph.New()
	err := d.decodeDataset(g)
	if err != nil {
		return nil, err
	}
	err = d.decodeValues(g)
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
	l, err := readUint64AsInt(d.r)
	if err != nil {
		return err
	}
	g.PrepareVirtualSpace(l)
	for i := 0; i < l; i++ {
		s, err := graph.VirtualDecode(d.r)
		if err != nil {
			return err
		}
		p, err := graph.VirtualDecode(d.r)
		if err != nil {
			return err
		}
		o, err := graph.VirtualDecode(d.r)
		if err != nil {
			return err
		}
		t := graph.Triple{S: s, P: p, O: o}
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
	l, err := readUint64AsInt(d.r)
	if err != nil {
		return err
	}
	g.PrepareValueSpace(l)
	for i := 0; i < l; i++ {
		val, err := graph.ValueDecode(d.r)
		if err != nil {
			return err
		}
		g.BulkAddValue(val)
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

func readUint64AsInt(r io.Reader) (int, error) {
	b := make([]byte, 8)
	n, err := io.ReadFull(r, b)
	if err != nil {
		return 0, err
	}
	if n != len(b) {
		return 0, fmt.Errorf("wanted to read %d bytes int, got only %d", len(b), n)
	}
	ui := stdbinary.LittleEndian.Uint64(b)
	return int(ui), nil
}
