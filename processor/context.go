package processor

import (
	"fmt"
	"log"

	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
)

type context struct {
	g         *graph.Graph
	variables *algebra.Variables
	receiver  Receiver
}

// WithReceiver returns a Copy of the context with a given Receiver
func (ctx *context) WithReceiver(r Receiver) *context {
	return &context{
		ctx.g,
		ctx.variables,
		r,
	}
}


type processingFunction func(ctx *context, tree *algebra.PatternTree, current algebra.Binding) error

var processorMap map[string]processingFunction

func init() {
	processorMap = map[string]processingFunction{
		"PROJECTION": projection,
		"GGP":        gpg,
		"DISTINCT":   nil,
		"REDUCED":    nil,
		"UNION":      nil,
		"OPTIONAL":   nil,
		"GRAPH":      nil,
		"FILTER":     logAndPropagate,
	}
}

func (ctx *context) process(tree *algebra.PatternTree, current algebra.Binding) error {
	proc, ok := processorMap[tree.Mode]
	if !ok {
		return fmt.Errorf("not a known ProcessingTree mode %q", tree.Mode)
	}
	if proc == nil {
		return fmt.Errorf("implementation: ProcessingTree mode %q is not yet implemented", tree.Mode)
	}
	return proc(ctx, tree, current)
}


func logAndPropagate(ctx *context, tree *algebra.PatternTree, current algebra.Binding) error {
	log.Println("#TODO just propagating", tree.Mode)
	return propagate(ctx, tree, current)
}

func propagate(ctx *context, tree *algebra.PatternTree, current algebra.Binding) error {
	for _, c := range tree.Children {
		err := ctx.process(c, current)
		if err != nil {
			return err
		}
	}
	return nil
}

