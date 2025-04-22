package golang

type GoNodeType int

const (
	GoProgram GoNodeType = iota
	GoFunc
	GoFor
	GoStatement
)

func goTypeToString(t GoNodeType) string {
	switch t {
	case GoProgram:
		return "GoProgram"
	case GoFunc:
		return "GoFunc"
	case GoFor:
		return "GoFor"
	case GoStatement:
		return "BfMoveLeft"
	default:
		return "Unknown"
	}
}
