Inspired by: https://github.com/jamiebuilds/the-super-tiny-compiler

# BrainForge

**BrainForge** is an all-in-one Brainfuck toolkit built in Go.  
It can compile existing Brainfuck programs into native Go code.

## Features
-  **Brainfuck Compiler**: Compile Brainfuck source files into Go programs.
-  **Pipeline Mode**: Chain generation and compilation in one step.
-  **CLI Tooling**: Easy-to-use command-line interface.

## How It Works

### 1. Tokenizer

- Scans the input Brainfuck code character-by-character.
- Maps valid Brainfuck instructions (`+ - > < [ ] . ,`) into token types.
- Ignores all other characters.
- **Output**: A list of `Token` structs.

### 2. Parser

- Converts tokens into an Abstract Syntax Tree (AST).
- Handles nested loops using recursive `LoopNode` structures.
- Maps each token into a node (e.g., `+` → `BfIncrement`, `[` → `BfBeginLoop`).
- **Output**: A structured AST representing the full program.

### 3. Transformer

- Optimizes the AST for performance.
- Merges consecutive operations like multiple `+` into a single `BfIncrement(count)`.
- Reduces Go code verbosity and improves runtime efficiency.
- **Output**: An optimized AST.

### 4. Code Generator

- Translates the optimized AST into valid Go source code.
- Handles tape initialization, pointer movement, IO, and loops.
- Uses standard Go packages like `fmt` for input/output.
- **Output**: A string containing complete Go code.

## Example Usage:
`
go run main.go compile --input "++>++[<+>-]."
`
