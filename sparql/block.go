package sparql

import (
	"fmt"

	"github.com/lercher/rdf/algebra"
)

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
