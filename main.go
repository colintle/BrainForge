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
		token.ToString()
	}

	ast := compiler.Parser(tokens)
	fmt.Println("\nBF AST:")
	for _, node := range ast {
		node.ToString()
	}

	goAst := compiler.Transformer(ast)
	fmt.Println("\nGo AST:")
	for _, node := range goAst {
		fmt.Println(node.ToString())
	}

	generatedCode := compiler.CodeGenerator(goAst)
	fmt.Println("\nGenerated Go Code:\n")
	fmt.Println(generatedCode)
}
