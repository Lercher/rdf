package algebra

// Algebra is the query algebra for a SPARQL select, ask, construct and describe
type Algebra struct {
	*Variables
	*Tree
}

// Optimize optimizes the query algebra. It is a NOP, currently.
func (a *Algebra) Optimize() *Algebra {
	return a
}
