package compiler

import (
	"fmt"
	"strings"

	"github.com/colintle/BrainForge/models/bf"
	"github.com/colintle/BrainForge/models/golang"
)

var visitor map[bf.BfType]golang.VisitorMethods = map[bf.BfType]golang.VisitorMethods{
	bf.BfIncrement: {
		Enter: func(node *bf.BfNode, parent *golang.GoNode) {
			n := len(parent.Body)
			if n > 0 {
				last := parent.Body[n-1]
				if last.Type == golang.GoStatement && strings.HasPrefix(last.Statement, "memory[pointer] += ") {
					var existing int
					fmt.Sscanf(last.Statement, "memory[pointer] += %d", &existing)
					last.Statement = fmt.Sprintf("memory[pointer] += %d", existing+1)
					return
				}
			}
	
			parent.Body = append(parent.Body, &golang.GoNode{
				Type:      golang.GoStatement,
				Statement: "memory[pointer] += 1",
			})
		},
	},
}


func Transformer(ast []*bf.BfNode)[]*golang.GoNode{
	var newAst []*golang.GoNode = []*golang.GoNode{}
	return newAst
}