package compiler

import (
	"fmt"
	"strings"

	"github.com/colintle/BrainForge/models/bf"
	"github.com/colintle/BrainForge/models/golang"
)

var visitor map[bf.BfType]golang.VisitorMethods = map[bf.BfType]golang.VisitorMethods{
	bf.BfIncrement: {
		Enter: func(node *bf.BfNode, parent *golang.GoNode) *golang.GoNode {
			n := len(parent.Body)
			if n > 0 {
				last := parent.Body[n - 1]
				if last.Type == golang.GoStatement && strings.HasPrefix(last.Statement, "memory[pointer] += ") {
					var existing int
					fmt.Sscanf(last.Statement, "memory[pointer] += %d", &existing)
					last.Statement = fmt.Sprintf("memory[pointer] += %d", existing+1)
					return nil
				}
			}
	
			parent.Body = append(parent.Body, &golang.GoNode{
				Type:      golang.GoStatement,
				Statement: "memory[pointer] += 1",
			})

			return nil
		},
	},
	bf.BfDecrement : {
		Enter: func(node *bf.BfNode, parent *golang.GoNode) *golang.GoNode {
			n := len(parent.Body)
			if n > 0 {
				last := parent.Body[n - 1]
				if last.Type == golang.GoStatement && strings.HasPrefix(last.Statement, "memory[pointer] -= ") {
					var existing int
					fmt.Sscanf(last.Statement, "memory[pointer] -= %d", &existing)
					last.Statement = fmt.Sprintf("memory[pointer] -= %d", existing - 1)
					return nil
				}
			}

			parent.Body = append(parent.Body, &golang.GoNode{
				Type:      golang.GoStatement,
				Statement: "memory[pointer] -= 1",
			})

			return nil
		},
	},
	bf.BfMoveLeft: {
		Enter: func(node *bf.BfNode, parent *golang.GoNode) *golang.GoNode{
			n := len(parent.Body)
			if n > 0 {
				last := parent.Body[n - 1]
				if last.Type == golang.GoStatement && strings.HasPrefix(last.Statement, "pointer -= ") {
					var existing int
					fmt.Sscanf(last.Statement, "pointer -= %d", &existing)
					last.Statement = fmt.Sprintf("pointer -= %d", existing - 1)
					return nil
				}
			}

			parent.Body = append(parent.Body, &golang.GoNode{
				Type: golang.GoStatement,
				Statement: "pointer -= 1",
			})

			return nil
		},
	},
	bf.BfMoveRight: {
		Enter: func(node *bf.BfNode, parent *golang.GoNode) *golang.GoNode{
			n := len(parent.Body)
			if n > 0 {
				last := parent.Body[n - 1]
				if last.Type == golang.GoStatement && strings.HasPrefix(last.Statement, "pointer += ") {
					var existing int
					fmt.Sscanf(last.Statement, "pointer += %d", &existing)
					last.Statement = fmt.Sprintf("pointer += %d", existing + 1)
					return nil
				}
			}

			parent.Body = append(parent.Body, &golang.GoNode{
				Type: golang.GoStatement,
				Statement: "pointer += 1",
			})

			return nil
		},
	},
	bf.BfOutput: {
		Enter: func(node *bf.BfNode, parent *golang.GoNode) *golang.GoNode{
			parent.Body = append(parent.Body, &golang.GoNode{
				Type: golang.GoStatement,
				Statement: "fmt.Printf(\"%c\", memory[pointer])",
			})

			return nil
		},
	},
	// in transformer, check if node.type is bfinput once and then initialize the variable
	bf.BfInput: {
		Enter: func(node *bf.BfNode, parent *golang.GoNode) *golang.GoNode{
			parent.Body = append(parent.Body, &golang.GoNode{
				Type: golang.GoStatement,
				Statement: "fmt.Scanf(\"%c\", &input)\nmemory[pointer] = input",
			})

			return nil
		},
	},
	bf.BfBeginLoop: {
		Enter: func(node *bf.BfNode, parent *golang.GoNode) *golang.GoNode{
			loopNode := &golang.GoNode{
				Type:      golang.GoFor,
				Condition: "memory[pointer] != 0",
				Body:      []*golang.GoNode{},
			}
	
			parent.Body = append(parent.Body, loopNode)
			return loopNode
		},
	},
}

func Transformer(ast []*bf.BfNode)[]*golang.GoNode{
	var newAst []*golang.GoNode = []*golang.GoNode{}

	programNode := &golang.GoNode{
		Type: golang.GoProgram,
		Body: []*golang.GoNode{},
	}

	var walk func(nodes []*bf.BfNode, parent *golang.GoNode)
	walk = func(nodes []*bf.BfNode, parent *golang.GoNode) {
		for _, node := range nodes {
			if visit, ok := visitor[node.Type]; ok {
				child := visit.Enter(node, parent)
				if len(node.Body) > 0 && child != nil {
					walk(node.Body, child)
				}
			}
		}
	}

	walk(ast, programNode)
	newAst = append(newAst, programNode)
	return newAst
}
