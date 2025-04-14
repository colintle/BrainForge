package main

import (
	"fmt"
	"github.com/colintle/BrainForge/compiler"
)

func main() {
	input := "++>++[<+>-]."

	tokens := compiler.Tokenizer(input)

	fmt.Println("Tokens:")
	for _, token := range tokens {
		fmt.Println(token.ToString())
	}

	ast := compiler.Parser(tokens)
	fmt.Println("\nAST:")
	for _, node := range ast {
		fmt.Println(node.ToString())
	}
}
