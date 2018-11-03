package processor

import (
	"fmt"

	"github.com/lercher/rdf/algebra"
)

func projection(ctx *context, tree *algebra.Tree, current algebra.Binding) error {
	proj, ok := tree.Term.(algebra.Projection)
	if !ok {
		return fmt.Errorf("malformed %q: term is %T instead of Projection", tree.Mode, tree.Term)
	}
	if proj.All {
		return propagate(ctx, tree, current)
	}
	projctx := ctx.WithReceiver(func(bs algebra.Binding) (bool, error) {
		clone := algebra.NewBinding(len(bs))
		for _, v := range proj.Variables {
			clone[v] = bs[v]
		}
		return ctx.receiver(clone)
	})
	return propagate(projctx, tree, current)
}
