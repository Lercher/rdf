package binary

import (
	stdbinary "encoding/binary"
	"fmt"
	"io"

	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/values"
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
	err = writeVarint(e.w, len(triples))
	if err != nil {
		return err
	}
	for tr := range triples {
		err = writeVarint(e.w, int(tr.S))
		if err != nil {
			return err
		}
		err = writeVarint(e.w, int(tr.P))
		if err != nil {
			return err
		}
		err = writeVarint(e.w, int(tr.O))
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Encoder) encodeValues(values []values.Value) error {
	_, err := e.w.Write([]byte(headerValues))
	if err != nil {
		return err
	}
	err = writeVarint(e.w, len(values))
	if err != nil {
		return err
	}
	for _, val := range values {
		tb, bs := val.Serialize()
		err = write(e.w, tb, bs)
		if err != nil {
			return err
		}
	}
	return nil
}

func write(w io.Writer, tb byte, bs[] byte) error {
	_, err:= w.Write([]byte{tb})
	if err != nil {
		return err
	}
	n, err := w.Write(bs)
	if err != nil {
		return err
	}
	if n != len(bs) {
		return fmt.Errorf("wanted to write %d content bytes, but only %d were written", len(bs), n)
	}
	return nil
}

func writeVarint(w io.Writer, u int) error {
	b := make([]byte, 10)
	n := stdbinary.PutVarint(b, int64(u))
	nn, err := w.Write(b[:n])
	if err != nil {
		return err
	}
	if n != nn {
		return fmt.Errorf("wanted to write %d length bytes, but only %d were written", n, nn)
	}
	return nil
}
