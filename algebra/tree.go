package algebra

import (
	"bytes"
	"strings"
)

// Tree results from parsing a groupGraphPattern
type Tree struct {
	Mode     string
	Term     Term // like the varOrIRIref of a graphGraphPattern
	Children []*Tree
}

// WrapIn wraps this tree in a new Tree with the given mode
func (tree *Tree) WrapIn(mode string) *Tree {
	t := &Tree{
		Mode:     mode,
		Children: []*Tree{tree},
	}
	return t
}

func (tree *Tree) String() string {
	buf := new(bytes.Buffer)
	tree.stringIndent(buf, 0)
	return buf.String()
}

func (tree *Tree) stringIndent(sw stringwriter, level int) {
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
