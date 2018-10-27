package binary

import (
	stdbinary "encoding/binary"
	"io"

	"github.com/lercher/rdf/graph"
)

// Encoder is used to serialize a virtualized rdf.Graph
type Encoder struct {
	w io.Writer
}

// NewEncoder creates an Encoder and assoiciates it with a Writer
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

// Encode writes a Graph to the initialized writer
func (e *Encoder) Encode(g *graph.Graph) error {
	err := e.encodeDataset(g.Dataset)
	if err != nil {
		return err
	}
	return e.encodeValues(g.Values())
}

func (e *Encoder) encodeDataset(triples []graph.Triple) error {
	_, err := e.w.Write([]byte("RDFGRA01"))
	if err != nil {
		return err
	}
	err = write64(e.w, len(triples))
	if err != nil {
		return err
	}
	for _, tr := range triples {
		err = tr.S.Encode(e.w)
		if err != nil {
			return err
		}
		err = tr.P.Encode(e.w)
		if err != nil {
			return err
		}
		err = tr.O.Encode(e.w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Encoder) encodeValues(values []graph.Value) error {
	_, err := e.w.Write([]byte("RDFVAL01"))
	if err != nil {
		return err
	}
	err = write64(e.w, len(values))
	if err != nil {
		return err
	}
	for _, val := range values {
		err = graph.ValueEncode(e.w, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func write64(w io.Writer, u int) error {
	b := make([]byte, 8)
	stdbinary.LittleEndian.PutUint64(b, uint64(u))
	_, err := w.Write(b)
	return err
}
