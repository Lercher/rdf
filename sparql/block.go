package sparql

import (
	"fmt"

	"github.com/lercher/rdf/algebra"
)

// Block is a mode tagged triple of Terms.
// A list of Blocks forms a where-clause.
type Block struct {
	Mode      string
	Subject   algebra.Term
	Predicate algebra.Term
	Object    algebra.Term
}

func (b *Block) String() string {
	return fmt.Sprintf(
		"{%s: %s %s %s}",
		b.Mode,
		algebra.DescribeTerm(b.Subject),
		algebra.DescribeTerm(b.Predicate),
		algebra.DescribeTerm(b.Object),
	)
}
