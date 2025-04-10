package models

type BfType int

const (
	BfIncrement BfType = iota
	BfDecrement
	BfMoveRight
	BfMoveLeft
	BfOutput
	BfInput
	BfLoop
)

type BfNode struct {
	Type BfType
	Occurences int
	Body []*BfNode
}
