package algebra

import (
	"fmt"
)

type Variables struct {
	m map[Variablename]Variable
	n []Variablename
}

func NewVariables() *Variables {
	return &Variables{
		m: make(map[Variablename]Variable),
	}
}

func (vs *Variables) Variable(vn Variablename) Variable {
	v, ok := vs.m[vn]
	if !ok {
		v = Variable(len(vs.n))
		vs.m[vn] = v
		vs.n = append(vs.n, vn)
	}
	return v
}

func (vs *Variables) Register(s string) Variable {
	vn := VariablenameParse(s)
	return vs.Variable(vn)
}

func (vs *Variables) NameOf(v Variable) Variablename {
	return vs.n[v]
}

func (vs *Variables) String() string {	
	return fmt.Sprint(vs.m)
}
