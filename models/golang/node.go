package golang

import (
	"strings"
	"fmt"
)

type GoNode struct {
	Type GoNodeType
	Statement string
	Condition string
	Body []*GoNode
}

func (node *GoNode) ToString() string {
	var builder strings.Builder
	node.buildString(&builder, 0)
	return builder.String()
}

func (node *GoNode) buildString(builder *strings.Builder, level int) {
	indent := strings.Repeat("  ", level)
	builder.WriteString(fmt.Sprintf("%sNode Type: %s", indent, goTypeToString(node.Type)))

	if node.Statement != "" {
		builder.WriteString(fmt.Sprintf(", Statement: %s", node.Statement))
	}
	if node.Condition != "" {
		builder.WriteString(fmt.Sprintf(", Condition: %s", node.Condition))
	}
	builder.WriteString("\n")

	for _, child := range node.Body {
		child.buildString(builder, level+1)
	}
}
