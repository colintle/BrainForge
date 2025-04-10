package compiler

import (
	"github.com/colintle/BrainForge/models/bf"
)

var tokenMap = map[byte]bf.BfType{
	'>': bf.BfMoveRight,
	'<': bf.BfMoveLeft,
	'+': bf.BfIncrement,
	'-': bf.BfDecrement,
	'.': bf.BfOutput,
	',': bf.BfInput,
	'[': bf.BfBeginLoop,
	']': bf.BfEndLoop,
}

func Tokenizer(input string) []*bf.BfToken {
	var bfTokens []*bf.BfToken

	for i := 0; i < len(input); i++ {
		char := input[i]

		tokenType, ok := tokenMap[char]
		if !ok {
			continue
		}

		bfTokens = append(bfTokens, &bf.BfToken{
			Type:       tokenType,
			Occurences: 1,
		})
	}

	return bfTokens
}
