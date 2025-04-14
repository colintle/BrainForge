package bf

import "fmt"

type BfToken struct {
	Type BfType
	Value byte
}

func (token *BfToken) ToString() string {
	return fmt.Sprintf("Token Type: %s | Value: %q", bfTypeToString(token.Type), token.Value)
}

