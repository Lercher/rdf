package algebra

import (
	"fmt"
)

// Projection stores the list of variables selected
type Projection struct {
	All bool
	Variables []Variable
}

func (p Projection) String(vs *Variables) string {
	if p.All {
		return "*"
	}
	var list []Variablename
	for _, v := range p.Variables {
		list = append(list, vs.NameOf(v))
	}
	return fmt.Sprint(list)
}