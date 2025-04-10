package bf

type BfType int

const (
	BfIncrement BfType = iota
	BfDecrement
	BfMoveRight
	BfMoveLeft
	BfOutput
	BfInput
	BfBeginLoop
	BfEndLoop
)

func bfTypeToString(t BfType) string {
	switch t {
	case BfIncrement:
		return "BfIncrement"
	case BfDecrement:
		return "BfDecrement"
	case BfMoveRight:
		return "BfMoveRight"
	case BfMoveLeft:
		return "BfMoveLeft"
	case BfOutput:
		return "BfOutput"
	case BfInput:
		return "BfInput"
	case BfBeginLoop:
		return "BfBeginLoop"
	case BfEndLoop:
		return "BfEndLoop"
	default:
		return "Unknown"
	}
}