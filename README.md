Inspired by: https://github.com/jamiebuilds/the-super-tiny-compiler

# BrainForge

**BrainForge** is an all-in-one Brainfuck toolkit built in Go.  
It can compile existing Brainfuck programs into native Go code.

## Features
-  **Brainfuck Compiler**: Compile Brainfuck source files into Go programs.
-  **Pipeline Mode**: Chain generation and compilation in one step.
-  **CLI Tooling**: Easy-to-use command-line interface.

## Example Usage:
`
go run main.go compile --input "++>++[<+>-]."
`