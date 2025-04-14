package compiler

import (
	"github.com/colintle/BrainForge/models/bf"
)

func Parser(tokens []*bf.BfToken) []*bf.BfNode{
	var current_index int = 0
	var ast []*bf.BfNode = []*bf.BfNode{}
	var token *bf.BfToken
	var walk func() *bf.BfNode

	walk = func() *bf.BfNode{
		token = tokens[current_index]

		if token.Type != bf.BfBeginLoop && token.Type != bf.BfEndLoop{
			current_index++
			return &bf.BfNode{
				Type: token.Type,
			}
		} else {
			if token.Type == bf.BfBeginLoop{
				current_index++
				var node bf.BfNode = bf.BfNode{
					Type: token.Type,
				}

				for current_index < len(tokens) && token.Type != bf.BfEndLoop {
					node.Body = append(node.Body, walk())
					token = tokens[current_index]
				}
				
				current_index++
				return &node
			}
		}

		return nil
	}

	for current_index < len(tokens){
		ast = append(ast, walk())
	}

	return ast
}
