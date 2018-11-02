package algebra

import (
	"fmt"
)

// Block is a mode tagged triple of Terms. Results from a triplesBlock
type Block struct {
	Subject   Term
	Predicate Term
	Object    Term
}

func (b *Block) String() string {
	return fmt.Sprintf(
		"{%s %s %s}",
		DescribeTerm(b.Subject),
		DescribeTerm(b.Predicate),
		DescribeTerm(b.Object),
	)
}
