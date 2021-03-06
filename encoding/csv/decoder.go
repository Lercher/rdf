package csv

import (
	stdcsv "encoding/csv"
	"fmt"
	"io"
	"net/url"

	"github.com/lercher/rdf/graph"
)

// Decoder deserializes csv data to graph.Graph
type Decoder struct {
	csvr           *stdcsv.Reader
	Entityprefix   string
	Propertyprefix string
	Reindex        bool
}

// NewDecoder creates a new csv Decoder for graph.Graphs
func NewDecoder(r io.Reader, entityprefix, propertyprefix string) *Decoder {
	csvr := stdcsv.NewReader(r)
	return &Decoder{csvr, entityprefix, propertyprefix, true}
}

// Decode deserializes a graph.Graph
func (d *Decoder) Decode() (*graph.Graph, error) {
	g := graph.New()

	hdr, err := d.csvr.Read()
	if err != nil {
		return nil, err
	}
	for i := range hdr {
		hdr[i] = url.PathEscape(hdr[i])
	}
	// log.Printf("Header read, %d columns", len(hdr))

	n := 0
	for {
		line, err := d.csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		n++

		if len(line) != len(hdr) {
			return nil, fmt.Errorf("csv line %d expected %d columns, got %d", n+1, len(hdr), len(line))
		}
		s := g.AddValueIRI(d.Entityprefix + url.PathEscape(line[0]))
		for col := 1; col < len(line); col++ {
			p := g.AddValueIRI(d.Propertyprefix + hdr[col])
			o := g.AddValueString(line[col])
			// log.Println(n, col, p.Value, o.Value)
			t := &graph.Triple{S: s.Virtual, P: p.Virtual, O: o.Virtual}
			g.BulkAddTriple(t, false)
			// log.Println(n, col, tr.P.String(g), tr.O.String(g))
		}
		// if n%1000 == 0 {
		//     log.Println(n)
		// }
	}
	// log.Println(n)
	if d.Reindex {
		g.RebuildIndex()
		// log.Println("Reindexed")
	}
	return g, nil
}
