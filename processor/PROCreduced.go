package processor

import (
	"github.com/lercher/rdf/algebra"
)

/*
https://stackoverflow.com/questions/2990343/sparql-distinct-vs-reduced

What is the difference between DISTINCT and REDUCED in SPARQL?

REDUCED is like a 'best effort' DISTINCT.
Whereas DISTINCT guarantees no duplicated results,
REDUCED may eliminate some, all, or no duplicates.

What's the point? Well DISTINCT can be expensive;
REDUCED can do the straightforward de-duplication work
(e.g. remove immediately repeated results) without having
to remember every row. In many applications that's good enough.

Having said that I've never used REDUCE, I've never seen anyone
use REDUCED, and never seen REDUCED mentioned in a talk or tutorial.
*/

func reduced(ctx *context, tree *algebra.Tree, current algebra.Binding) error {
	gate := NewReducedGate(ctx.receiver)
	projctx := ctx.WithReceiver(gate.Receive)
	return propagate(projctx, tree, current)
}
