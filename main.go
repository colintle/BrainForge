package main

import (
    "fmt"
    "log"
    "os"

    "github.com/colintle/BrainForge/compiler"
    "github.com/urfave/cli/v2"
)

func compile(c *cli.Context) error {
    src := c.String("input")
    if src == "" {
        return fmt.Errorf("please provide --input")
    }

    tokens := compiler.Tokenizer(src)
    // fmt.Println("Tokens:")
    // for _, tok := range tokens {
    //     fmt.Print(tok.ToString())
    // }

    ast := compiler.Parser(tokens)
    // fmt.Println("\nBF AST:")
    // for _, node := range ast {
    //     fmt.Print(node.ToString())
    // }

    goAst := compiler.Transformer(ast)
    // fmt.Println("\nGo AST:")
    // for _, node := range goAst {
    //     fmt.Print(node.ToString())
    // }

    code := compiler.CodeGenerator(goAst)
    fmt.Print("\nGenerated Go Code:\n", code)
    return nil
}

// Example Usage
// go run main.go compile --input "++>++[<+>-]."
func main() {
    app := &cli.App{
        Name:  "brainforge",
        Usage: "compile Brainfuck to Go",
        Commands: []*cli.Command{
            {
                Name:   "compile",
                Usage:  "tokenize, parse, transform, and codegen",
                Flags: []cli.Flag{
                    &cli.StringFlag{
                        Name:     "input, i",
                        Usage:    "Brainfuck source string or file",
                        Required: true,
                    },
                },
                Action: compile,
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
