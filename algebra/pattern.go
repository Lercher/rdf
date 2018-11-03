package algebra

import "github.com/lercher/rdf/graph"

// Pattern is either a Variable or a Virtual, but not both.
type Pattern struct {
	variable Variable
	virtual  graph.Virtual
}

func (p *Pattern) String(g *graph.Graph) string {
	if p.IsAnonymous() {
		return "_"
	}
	if p.IsVariable() {
		return string(p.variable)
	}
	return p.virtual.String(g)
}

// PatternOfSameVar holds true, if both are Variables and both are not annonymous and have the same vars
func PatternOfSameVar(p1, p2 *Pattern) bool {
	return p1 != nil && p2 != nil && p1.Variable() == p2.Variable() && p1.Variable() >= 0
}

// IsVariable holds true, if there is no Virtual value
func (p *Pattern) IsVariable() bool {
	return p == nil || p.virtual.IsNil()
}

// IsAnonymous holds true, if the virtual is nil or the variable name is "_"
func (p *Pattern) IsAnonymous() bool {
	return p == nil || (p.virtual.IsNil() && p.variable == BlankVariable)
}

// WontMatch holds true, if the virtual WontMatch
func (p *Pattern) WontMatch() bool {
	return p != nil && p.virtual.WontMatch()
}

// Virtual returns this virtual value, unless it is a variable pattern, in this case the zero Virtual is returned.
// It is recommended to test IsVariable/IsAnonymous before using this method.
func (p *Pattern) Virtual() graph.Virtual {
	if p.IsVariable() {
		return graph.NotAVirtual
	}
	return p.virtual
}

// Variable returns the Variable, if it is one.
// It is recommended to test IsVariable/IsAnonymous before using this method.
func (p *Pattern) Variable() Variable {
	return p.variable
}

// PatternAnonymous is the Pattern, that matches anything and will never be bound
func PatternAnonymous() *Pattern {
	return nil
}

// PatternWontMatch is the Pattern, that matches nothing
func PatternWontMatch() *Pattern {
	return PatternLiteral(graph.WontMatchVirtual)
}

// PatternVariable creates a Variable pattern that either matches anything or if bound, its value
func PatternVariable(v Variable) *Pattern {
	return &Pattern{variable: v, virtual: graph.NotAVirtual}
}

// PatternLiteral creates a constant matching pattern
func PatternLiteral(v graph.Virtual) *Pattern {
	if v.IsNil() {
		panic("a virtual pattern must not be nil")
	}
	return &Pattern{virtual: v, variable: NotAVariable}
}

// PatternFromTerm returns a Pattern from a Term
func PatternFromTerm(g *graph.Graph, t Term) *Pattern {
	if variable, ok := t.(Variable); ok {
		return PatternVariable(variable)
	}
	vv := g.VirtualValue(t)
	if !vv.Known {
		return PatternWontMatch()
	}
	return PatternLiteral(vv.Virtual)
}
