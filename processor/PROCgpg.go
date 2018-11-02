package processor

import (
	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/algebra"
)

func gpg(ctx *context, tree *algebra.PatternTree, current algebra.Binding) error {
	join(ctx, tree.Blocks, current) // #TODO: children need to be part of join
	return nil
}

func join(ctx *context, blocks []*algebra.Block, current algebra.Binding) (bool, error) {
	if len(blocks) == 0 {
		return ctx.receiver(ctx.g, current)
	}

	block := blocks[0]
	sp := patternFromTerm(ctx.g, current, block.Subject)
	pp := patternFromTerm(ctx.g, current, block.Predicate)
	op := patternFromTerm(ctx.g, current, block.Object)
	tp := &algebra.TriplePattern{S: sp, P: pp, O: op}
	ms := Match(ctx.g, tp)
	for _, t := range ms {
		bs := current.BindTriple(tp, t)
		cont, err := join(ctx, blocks[1:], bs)
		if !cont || err != nil {
			return cont, err
		}
	}
	return true, nil // #TODO: is continue=true OK here?
}

func patternFromTerm(g *graph.Graph, current algebra.Binding, t algebra.Term) *algebra.Pattern {
	bt, virtual := current.BindTerm(t)
	if bt != nil {
		return algebra.PatternFromTerm(g, bt)
	}
	return algebra.PatternLiteral(virtual)
}