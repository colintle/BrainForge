package compiler

import (
	"strings"
	"github.com/colintle/BrainForge/models/golang"
)

func CodeGenerator(ast []*golang.GoNode) string {
	var builder strings.Builder
	builder.WriteString("package main\n\n")
	builder.WriteString("import (\n\t\"fmt\"\n)\n\n")
	builder.WriteString("func main() {\n")
	builder.WriteString("\tvar memory [30000]byte\n")
	builder.WriteString("\tvar pointer int\n\n")
	builder.WriteString("\tvar input byte \n\n")
	builder.WriteString("\t_ = input \t")

	// traverse ast
	
	builder.WriteString("}\n")

	return builder.String()
}
