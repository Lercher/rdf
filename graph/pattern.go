package graph

// Pattern is either a Variable or a Virtual, but not both.
type Pattern struct {
	variable Variable
	virtual  Virtual
}

func(p *Pattern) String(g *Graph) string{
	if p.IsAnonymous() {
		return "_"
	}
	if p.IsVariable() {
		return string(p.variable)
	}
	return p.virtual.String(g)
}

// IsVariable holds true, if there is no Virtual value
func (p *Pattern) IsVariable() bool {
	return p == nil || p.virtual.IsNil()
}

// IsAnonymous holds true, if the pointer is nil or the variable name is "_"
func (p *Pattern) IsAnonymous() bool {
	return p == nil || (p.virtual.IsNil() && p.variable == "_")
}

// Virtual returns this virtual value, unless it is a variable pattern, in this case the zero Virtual is returned.
// It is recommended to test IsVariable/IsAnonymous before using this method.
func (p *Pattern) Virtual() Virtual {
	if p.IsVariable() {
		return Virtual{}
	}
	return p.virtual
}

// PatternAnonymous is the Pattern, that matches anything and will never be bound
func PatternAnonymous() *Pattern {
	return nil
}

// PatternVariable creates a named Variable pattern that either matches anything or if bound, its value
func PatternVariable(s string) *Pattern {
	return &Pattern{variable: Variable(s)}
}

// PatternLiteral creates a constant matching pattern
func PatternLiteral(v Virtual) *Pattern {
	if v.IsNil() {
		panic("a virtual pattern must not be nil")
	}
	return &Pattern{virtual: v}
}
