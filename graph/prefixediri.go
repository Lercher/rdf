package graph

// PrefixedIRI pairs a Prefix with its IRI
type PrefixedIRI struct {
	Prefix
	IRI
}

func (pi *PrefixedIRI) String() string {
	return pi.Prefix.String() + pi.IRI.String()
}