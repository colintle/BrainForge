package bf

import (
	"strings"
	"fmt"
)

type BfNode struct{
	Type BfType
	Body []*BfNode
}

func (node *BfNode) ToString() string {
	var builder strings.Builder
	node.buildString(&builder, 0)
	return builder.String()
}

func (node *BfNode) buildString(builder *strings.Builder, level int) {
	indent := strings.Repeat("  ", level)
	builder.WriteString(fmt.Sprintf("%sNode Type: %s\n", indent, bfTypeToString(node.Type)))

	for _, child := range node.Body {
		child.buildString(builder, level+1)
	}
}
