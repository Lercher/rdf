package graph

import (
	"encoding/binary"
	"fmt"
)

type VirtualType int

type VirtualValue struct {
	Virtual
	Value
}

func NewVValue(g *Graph, primitive interface{}) VirtualValue {
	switch v := primitive.(type) {
	case int:
		return VValueInt(g, v)
	case float64:
		return VValueFloat(g, v)
	case string:
		return VValueString(g, v)
	default:
		panic(fmt.Errorf("unsupported value type %T", primitive))
	}
}

func VValueString(g *Graph, s string) VirtualValue {
	return vvalue(g, []byte(s), s)
}

func VValueInt(g *Graph, i int) VirtualValue {
	i64 := int64(i)
	b := make([]byte, 8, 9)
	binary.LittleEndian.PutUint64(b, uint64(i64))
	b = append(b, byte(ValueTypeInt))
	return vvalue(g, b, i)
}

func VValueFloat(g *Graph, f float64) VirtualValue {
	b := make([]byte, 8, 9)
	binary.LittleEndian.PutUint64(b, uint64(f))
	b = append(b, byte(ValueTypeFloat))
	return vvalue(g, b, f)
}

func vvalue(g *Graph, bytes []byte, v interface{}) VirtualValue {
	h := hash(bytes)
	val, ok := g.valuemap[h]
	if !ok {
		val = Value(v)
		g.valuemap[h] = val
	}
	return VirtualValue{
		Virtual: h,
		Value:   val,
	}
}

const (
	ValueTypeString = VirtualType(iota)
	ValueTypeInt
	ValueTypeFloat
)
