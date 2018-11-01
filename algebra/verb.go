package algebra

import (
	"fmt"
)

// Verb is either one of its components
type Verb interface {
	String() string
}

// DescribeVerb returns kind and content of a Verb as a string
func DescribeVerb(v Verb) string {
	if v == nil {
		return "nil"
	}
	return fmt.Sprintf("%T %s", v, v.String())
}
