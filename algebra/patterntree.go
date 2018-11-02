package algebra

import (
	"bytes"
	"strings"
)

// PatternTree results from parsing a groupGraphPattern
type PatternTree struct {
	Mode     string
	Term     Term // like the varOrIRIref of a graphGraphPattern
	Children []*PatternTree
	Blocks   []*Block
}

func (tree *PatternTree) String() string {
	buf := new(bytes.Buffer)
	tree.stringIndent(buf, 0)
	return buf.String()
}

func (tree *PatternTree) stringIndent(sw stringwriter, level int) {
	const indent = "  "
	pre := strings.Repeat(indent, level)
	sw.WriteString(pre)
	sw.WriteString("(")
	sw.WriteString(tree.Mode)
	if tree.Term != nil {
		t := DescribeTerm(tree.Term)
		sw.WriteString(" ")
		sw.WriteString(t)
	}

	for _, b := range tree.Blocks {
		sw.WriteString("\n")
		sw.WriteString(pre)
		sw.WriteString(indent)
		sw.WriteString(b.String())
	}
	sw.WriteString("\n")

	for _, child := range tree.Children {
		child.stringIndent(sw, level+1)
	}

	sw.WriteString(pre)
	sw.WriteString(")\n")
}

type stringwriter interface {
	WriteString(s string) (int, error)
}
