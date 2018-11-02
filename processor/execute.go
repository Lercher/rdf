package processor

import (
	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/algebra"
)

// Execute executes an Algebra on a Graph, each resulting Binding will be propagated to the Receiver
func Execute(a *algebra.Algebra, g *graph.Graph, receiver Receiver) error {
	a = a.Optimize()
	ctx := &context{g: g, receiver: receiver, variables: a.Variables}
	empty := algebra.NewBinding(ctx.variables.Count())
	return ctx.process(a.PatternTree, empty)
}
