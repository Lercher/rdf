package algebra

import "fmt"

// Variable is just an index and therefore depends on Variables
type Variable int

// NotAVariable is the value for something that aggregates a Variable optionally
const NotAVariable = Variable(-1)

// BlankVariable is the representation of a blankNode _:PN_LOCAL
const BlankVariable = Variable(-2)

func (v Variable) String() string {
	if v == BlankVariable {
		return "_"
	}
	return fmt.Sprintf("$%d", int(v))
}
