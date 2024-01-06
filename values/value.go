package values

// Value is virtualized and can be stored in a Graph's assertions
type Value interface {
	String() string
	Serialize() (byte, []byte)
	Inner() interface{}
	IsSameTypeAndLessThan(Value) (bool, bool)
}

// Less compares two Values.
// If they are of the same type it is delegated to IsSameTypeAndLessThan of the type.
// Float and Int are compared as Floats. String and Literal as strings.
// If other different types, the type's serialization type byte is relevant for the ordering.
func Less(v1, v2 Value) bool {
	if v1 == nil && v2 == nil {
		return false // nil is always less than anyting else
	}
	if v1 == nil {
		return true // nil is always less than anyting else
	}
	if v2 == nil {
		return false // nil is always less than anyting else
	}
	ok, less := v1.IsSameTypeAndLessThan(v2)
	if ok {
		return less
	}
	b1, _ := v1.Serialize()
	b2, _ := v2.Serialize()
	return b1 < b2
}

// LessSlice orders lexicographically
func LessSlice(vs1, vs2 []Value) bool {
	for i := range vs2 {
		if i >= len(vs1) {
			return true // vs1 is shorter than vs2, but we are here, so everything else was not Less
		}
		if Less(vs1[i], vs2[i]) {
			return true
		}
	}
	return false
}
