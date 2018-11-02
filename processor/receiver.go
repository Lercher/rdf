package processor

import (
	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
)

// Receiver is a func for clients of Execute, it receives the results of the query one by one, the bool is true if the client wants to continue
type Receiver func(g *graph.Graph, bs algebra.Binding) (bool, error)
