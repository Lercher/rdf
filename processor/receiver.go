package processor

import (
	"github.com/lercher/rdf/algebra"
)

// Receiver is a func for clients of Execute, it receives the results of the query one by one, the bool is true if the client wants to continue
type Receiver func(bs algebra.Binding) (bool, error)
