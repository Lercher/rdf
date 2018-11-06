package processor

import (
	"github.com/lercher/rdf/algebra"
)

func distinct(ctx *context, tree *algebra.Tree, current algebra.Binding) error {
	gate := NewDistinctGate(ctx.receiver)
	projctx := ctx.WithReceiver(gate.Receive)
	return propagate(projctx, tree, current)
}
