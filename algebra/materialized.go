package algebra

import (
	"fmt"

	"github.com/lercher/rdf/values"
)

// Materialized is a Variable that has a Value
type Materialized struct {
	Name  Variablename
	Value values.Value
}

func (m *Materialized) String() string {
	return fmt.Sprintf("%s=%[2]T:%[2]v", m.Name, m.Value)
}
