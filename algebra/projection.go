package algebra

import (
	"fmt"
)

// Projection stores the list of variables selected
type Projection struct {
	All       bool
	Variables []Variable
}

// Names returns a list of Variablenames
func (p Projection) Names(vs *Variables) []Variablename {
	var list []Variablename
	if p.All {
		list = append(list, Variablename("*"))
	} else {
		for _, v := range p.Variables {
			list = append(list, vs.NameOf(v))
		}
	}
	return list
}

// StringNames returns a list of names as a string
func (p Projection) StringNames(vs *Variables) string {
	return fmt.Sprint(p.Names(vs))
}

func (p Projection) String() string {
	if p.All {
		return "*"
	}
	var list []string
	for _, v := range p.Variables {
		list = append(list, v.String())
	}
	return fmt.Sprint(list)
}
