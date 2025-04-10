package bf

import "fmt"

type BfToken struct {
	Type BfType
	Value byte
}

func (token *BfToken) ToString() {
	fmt.Printf("Token Type: %s | Value: %s\n", bfTypeToString(token.Type), token.Value)
}
