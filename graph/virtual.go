package graph

import (
	"fmt"
	"strings"

	"github.com/lercher/rdf/values"
)

// Virtual is the datatype for virtualizing values inside a Graph. It's the index+1 into a slice ov Values, so 0 is Nil
type Virtual int32

// WontMatchVirtual represents a Virtual that won't match, because it is no member of the graph
const WontMatchVirtual = Virtual(-1)

// NotAVirtual is a Virtual Value that doesn't represent a Virtual
const NotAVirtual = Virtual(0)

// Virtual2 is a pair of Virtuals
type Virtual2 struct {
	V1, V2 Virtual
}

// IsNil holds true, if the underlying datatype is zero, i.e. 4 zero bytes
func (v Virtual) IsNil() bool {
	return v == 0
}

// WontMatch holds true, if the Virtual Represents something that is not a member of a graphs Values
func (v Virtual) WontMatch() bool {
	return v == -1
}

func (v Virtual) String(g *Graph) string {
	val := v.Value(g)
	ty := fmt.Sprintf("%T", val)
	ty = strings.TrimPrefix(ty, "*")
	ty = strings.TrimPrefix(ty, "values.")
	return fmt.Sprintf("(%s:%v)", ty, val)
}

// Value returns the real Value of a virtualized Triple part in subject/predicate/object position
// according to the records of the creating Graph. It is undefined behaviour to
// retrieve a Virtual's Value from a different Graph
func (v Virtual) Value(g *Graph) values.Value {
	index := int(v) - 1
	if 0 <= index && index < len(g.valuelist) {
		val := g.valuelist[index]
		return val
	}
	return nil
}
