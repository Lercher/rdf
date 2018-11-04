package algebra

import (
	"fmt"
	"strings"
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
	ty := fmt.Sprintf("%T", t)
	ty = strings.TrimPrefix(ty, "*")
	ty = strings.TrimPrefix(ty, "algebra.")
	ty = strings.TrimPrefix(ty, "values.")
	return fmt.Sprintf("(%s:%s)", ty, t.String())
}
