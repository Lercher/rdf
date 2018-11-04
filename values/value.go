package values

// Value is virtualized and can be stored in a Graph's assertions
type Value interface {
	String() string
	Serialize() (byte, []byte)
	Inner() interface{}
}
