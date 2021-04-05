package eqgo

import (
	"fmt"
	"go/token"
	"strings"
)

type Formatter interface {
	Format(eq bool, n *node) string
}

type DefaultFormatter struct {
	LeftFSet, RightFSet *token.FileSet
}

func (f DefaultFormatter) Format(eq bool, n *node) string {
	if eq {
		return "equivalent"
	}

	return "not equivalent:\n" + f.formatWithLevel(n, 0)
}

func (f DefaultFormatter) formatWithLevel(n *node, level int) string {
	var builder strings.Builder
	fmt.Fprint(&builder, strings.Repeat("    ", level), n.msg)

	leftPosition := f.LeftFSet.Position(n.leftPos)
	rightPosition := f.RightFSet.Position(n.rightPos)

	if leftPosition.IsValid() || rightPosition.IsValid() {
		fmt.Fprintf(&builder, " (%v != %v)", leftPosition, rightPosition)
	}
	fmt.Fprint(&builder, "\n")

	for _, c := range n.children {
		fmt.Fprint(&builder, f.formatWithLevel(c, level+1), "\n")
	}
	return strings.TrimRight(builder.String(), "\n")
}
