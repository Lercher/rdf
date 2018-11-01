package algebra

import (
	"strings"
)

// Variablename is the name of a variable without the prefixes $ or ?
type Variablename string

// VariablenameParse creates a Variablename from a string
func VariablenameParse(s string) Variablename {
	if strings.HasPrefix(s, "?") {
		return Variablename(strings.TrimPrefix(s, "?"))
	}
	if strings.HasPrefix(s, "$") {
		return Variablename(strings.TrimPrefix(s, "$"))
	}
	return Variablename(s)
}
