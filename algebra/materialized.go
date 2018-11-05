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
	if m.Value == nil {
		// happens, if an optional did bind nothing
		return fmt.Sprintf("%s=nil", m.Name)
	}
	return fmt.Sprintf("%s=%s", m.Name, m.Value)
}
