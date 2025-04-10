package bf

import "fmt"

type BfToken struct {
	Type BfType
	Occurences int
}

func (token *BfToken) ToString() {
	fmt.Printf("Token Type: %s | Occureances: %f\n", bfTypeToString(token.Type), token.Occurences)
}
