package values

func init() {
	// constants
	RegisterConstructor(0, constructNil)
	RegisterConstructor('T', constructTrue)
	RegisterConstructor('F', constructFalse)
	RegisterConstructor('I', constructIRI)
	// datatypes with stored values
	RegisterConstructor('i', constructInt)
	RegisterConstructor('f', constructFloat)
	RegisterConstructor('s', constructString)
	RegisterConstructor('l', constructLiteral)
}
