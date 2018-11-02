package algebra

import (
	"fmt"
)

// Variables is the repository of all Variables of a query
type Variables struct {
	m map[Variablename]Variable
	n []Variablename
}

// NewVariables creates a new repository of Variables
func NewVariables() *Variables {
	return &Variables{
		m: make(map[Variablename]Variable),
	}
}

// Variable gets and registers a Variable by its name
func (vs *Variables) Variable(vn Variablename) Variable {
	v, ok := vs.m[vn]
	if !ok {
		v = Variable(len(vs.n))
		vs.m[vn] = v
		vs.n = append(vs.n, vn)
	}
	return v
}

// Register parses and then gets and registers a Variablename 
func (vs *Variables) Register(s string) Variable {
	vn := VariablenameParse(s)
	return vs.Variable(vn)
}

// NameOf retrieves the name of a Variable in this repository
func (vs *Variables) NameOf(v Variable) Variablename {
	return vs.n[v]
}

func (vs *Variables) String() string {	
	var list []string
	for i, vn := range vs.n {
		list = append(list, fmt.Sprintf("%s:$%d", vn, i))
	}
	return fmt.Sprint(list)
}
