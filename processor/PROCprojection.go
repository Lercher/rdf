package processor

import (
	"fmt"

	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
)

func projection(ctx *context, tree *algebra.Tree, current algebra.Binding) error {
	proj, ok := tree.Term.(algebra.Projection)
	if !ok {
		return fmt.Errorf("malformed %q: term is %T instead of Projection", tree.Mode, tree.Term)
	}
	if proj.All {
		return propagate(ctx, tree, current)
	}
	projctx := ctx.WithReceiver(func(g *graph.Graph, vs *algebra.Variables, bs algebra.Binding) (bool, error) {
		clone := algebra.NewBinding(len(bs))
		for _, v := range proj.Variables {
			clone[v] = bs[v]
		}
		return ctx.receiver(g, vs, clone)
	})
	return propagate(projctx, tree, current)
}
