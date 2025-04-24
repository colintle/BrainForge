package compiler

import (
	"strings"
	"fmt"
	"github.com/colintle/BrainForge/models/golang"
)

func generateCode(builder *strings.Builder, node *golang.GoNode, level int) {
	indent := strings.Repeat("\t", level)

	switch node.Type {
	case golang.GoStatement:
		builder.WriteString(fmt.Sprintf("%s%s\n", indent, node.Statement))

	case golang.GoFor:
		builder.WriteString(fmt.Sprintf("%sfor %s {\n", indent, node.Condition))
		for _, child := range node.Body {
			generateCode(builder, child, level+1)
		}
		builder.WriteString(fmt.Sprintf("%s}\n", indent))

	case golang.GoProgram:
		for _, child := range node.Body {
			generateCode(builder, child, level)
		}
	}
}

func CodeGenerator(ast []*golang.GoNode) string {
	var builder strings.Builder
	builder.WriteString("package main\n\n")
	builder.WriteString("import (\n\t\"fmt\"\n)\n\n")
	builder.WriteString("func main() {\n")
	builder.WriteString("\tvar memory [30000]byte\n")
	builder.WriteString("\tvar pointer int\n")
	builder.WriteString("\tvar input byte \n")
	builder.WriteString("\t_ = input \n\n")

	// traverse ast
	for _, node := range ast {
		generateCode(&builder, node, 1)
	}

	builder.WriteString("}\n")

	return builder.String()
}
