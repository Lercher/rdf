package algebra

import (
	"strings"
)

type Variablename string

func VariablenameParse(s string) Variablename {
	if strings.HasPrefix(s, "?") {
		return Variablename(strings.TrimPrefix(s, "?"))
	}
	if strings.HasPrefix(s, "$") {
		return Variablename(strings.TrimPrefix(s, "$"))
	}
	return Variablename(s)
}
