package graph

import "testing"

func TestVirtualBytes(t *testing.T) {
	v := hash([]byte("something else"))
	b := v.Bytes()
	vv := VirtualFrom(b)
	bb := vv.Bytes()
	if v != vv {
		t.Errorf("encode then decode: want %v, got %v", b, bb)
	}
	if bb[0] != 220 {
		t.Log(bb)
		t.Errorf("first encoded byte: want %d, got %d", 220, bb[0])
	}
}
