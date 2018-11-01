package algebra

import (
	"fmt"
)

// Term is a variable or any variant of a literal
type Term interface {
	String() string
}

// DescribeTerm returns kind and content of a Term as a string
func DescribeTerm(t Term) string {
	if t == nil {
		return "nil-term"
	}
	return fmt.Sprintf("%T %s", t, t.String())
}
