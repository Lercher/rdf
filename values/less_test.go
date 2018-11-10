package values

import (
	"testing"
)

func TestLess(t *testing.T) {
	wantLess(t, A, A, false)
	wantLess(t, XS, A, false)
	wantLess(t, A, XS, true)
	wantLess(t, Int(0), Int(1), true)
	wantLess(t, Int(1), Int(-1), false)
	wantLess(t, Int(1), Int(1), false)
	wantLess(t, Float(0), Int(1), true)
	wantLess(t, Float(1), Int(-1), false)
	wantLess(t, Float(1), Int(1), false)
	wantLess(t, Int(0), Float(1), true)
	wantLess(t, Int(1), Float(-1), false)
	wantLess(t, Int(1), Float(1), false)
	wantLess(t, Float(1.5), Float(1.6), true)
	wantLess(t, String("a"), String("aa"), true)
	wantLess(t, String("a"), String("AA"), false)
	wantLess(t, Bool(false), Bool(true), true)
	wantLess(t, Bool(false), Bool(false), false)
	wantLess(t, IRI("a"), String("a"), true)
	wantLess(t, NIL, String("a"), true)
	wantLess(t, nil, String("a"), true)
	wantLess(t, nil, nil, false)
	wantLess(t, Int(1), nil, false)
	wantLess(t, LiteralString("xx"), String("xx"), false)
	wantLess(t, LiteralString("xx"), LiteralString("xx"), false)
	wantLess(t, String("xx"), LiteralString("xx"), false)
	wantLess(t, LiteralString("12"), String("xx"), true)
	wantLess(t, LiteralString("12"), LiteralString("xx"), true)
	wantLess(t, String("12"), LiteralString("xx"), true)
	wantLess(t, LiteralFrom("12", "DE", NotAnIRI), LiteralFrom("12", "EN", NotAnIRI), true)
	wantLess(t, LiteralFrom("12", "EN", NotAnIRI), LiteralFrom("12", "DE", NotAnIRI), false)
	wantLess(t, LiteralFrom("12", "", IRI("xs:string")), LiteralFrom("12", "", IRI("xs:int")), false)
	wantLess(t, LiteralFrom("12", "", IRI("xs:int")), LiteralFrom("12", "", IRI("xs:string")), true)
	wantLess(t, LiteralFrom("a", "", IRI("xs:string")), LiteralFrom("aa", "", IRI("xs:int")), true)
	wantLess(t, LiteralFrom("a", "", IRI("xs:int")), LiteralFrom("aa", "", IRI("xs:string")), true)
}

func TestLessSlice(t *testing.T) {
	s123 := []Value{Int(1), Int(2), Int(3)}
	s124 := []Value{Int(1), Int(2), Int(4)}
	s120 := []Value{Int(1), Int(2), Int(0)}
	s469 := []Value{Int(4), Int(6), Int(9)}
	s1234 := []Value{Int(1), Int(2), Int(3), Int(4)}

	wantLessS(t, nil, nil, false)
	wantLessS(t, nil, s123, true)
	wantLessS(t, s123, nil, false)
	wantLessS(t, s123, s123, false)
	wantLessS(t, s123, s124, true)
	wantLessS(t, s123, s120, false)
	wantLessS(t, s123, s469, true)
	wantLessS(t, s123, s1234, true)
	wantLessS(t, s1234, s123, false)
	wantLessS(t, s1234, s469, true)
	wantLessS(t, s469, s1234, true) // sic! 469<123 and the second is longer. i.e. don't compare slices of different length
}

func wantLessS(t *testing.T, vs1, vs2 []Value, wantb bool) {
	t.Helper()
	gotb := LessSlice(vs1, vs2)
	if gotb != wantb {
		t.Errorf("want LessSlice(%v, %v) = %v, got %v", vs1, vs2, wantb, gotb)
	}
}

func wantLess(t *testing.T, v1, v2 Value, wantb bool) {
	t.Helper()
	gotb := Less(v1, v2)
	if gotb != wantb {
		t.Errorf("want Less(%s, %s) = %v, got %v", v1, v2, wantb, gotb)
	}
}
