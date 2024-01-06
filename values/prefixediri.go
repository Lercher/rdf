package values

// PrefixedIRI pairs a Prefix with its IRI
type PrefixedIRI struct {
	Prefix
	IRI
}

func (pi *PrefixedIRI) String() string {
	return pi.Prefix.String() + pi.IRI.String()
}

// Combine forms a full IRI from the given tail
func (pi *PrefixedIRI) Combine(tail string) IRI {
	return IRI(string(pi.IRI) + tail)
}
