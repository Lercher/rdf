package values

func init() {
	RegisterConstructor(0, constructNil)
	RegisterConstructor('i', constructInt)
	RegisterConstructor('f', constructFloat)
	RegisterConstructor('s', constructString)
	RegisterConstructor('T', constructTrue)
	RegisterConstructor('F', constructFalse)
	RegisterConstructor('I', constructIRI)
	RegisterConstructor('L', constructLiteral)
}
