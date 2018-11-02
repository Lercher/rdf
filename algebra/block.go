package algebra

import (
	"fmt"
)

// Block is a triple of Terms. Results from a triplesBlock.
// A block term can be a Variable or iriRef|rdfLiteral|numericLiteral|booleanLiteral|blankNode|NIL
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
