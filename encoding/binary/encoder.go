package binary

import (
	stdbinary "encoding/binary"
	"io"

	"github.com/lercher/rdf/graph"
)

const (
	headerDataset = "RDFGRA01"
	headerValues  = "RDFVAL01"
)

// Encoder is used to serialize a virtualized rdf.Graph in binary form
type Encoder struct {
	w io.Writer
}

// NewEncoder creates a new binary Encoder for rdf.Graph and assoiciates it with a Writer
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

// Encode writes a Graph to the initialized writer
func (e *Encoder) Encode(g *graph.Graph) error {
	err := e.encodeValues(g.Values())
	if err != nil {
		return err
	}
	return e.encodeDataset(g.DatasetMap())
}

func (e *Encoder) encodeDataset(triples map[graph.Triple]*graph.Triple) error {
	_, err := e.w.Write([]byte(headerDataset))
	if err != nil {
		return err
	}
	err = write32(e.w, len(triples))
	if err != nil {
		return err
	}
	for tr := range triples {
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
	_, err := e.w.Write([]byte(headerValues))
	if err != nil {
		return err
	}
	err = write32(e.w, len(values))
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

func write32(w io.Writer, u int) error {
	b := make([]byte, 4)
	stdbinary.LittleEndian.PutUint32(b, uint32(u))
	_, err := w.Write(b)
	return err
}
