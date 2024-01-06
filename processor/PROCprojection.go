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

	projctx := ctx.WithReceiver(func(bs algebra.Binding, vs *algebra.Variables) (bool, error) {
		clone := algebra.NewBinding(0)
		pvs := algebra.NewVariables()
		for _, v := range proj.Variables {
			pv := pvs.Variable(vs.NameOf(v))
			for int(pv) >= len(clone) {
				clone = append(clone, bs[v])
			}
		}
		return ctx.receiver(clone, pvs)
	})
	return propagate(projctx, tree, current)
}
