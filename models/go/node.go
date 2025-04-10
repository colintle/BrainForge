package models

type GoNodeType int

const (
	GoProgram GoNodeType = iota
	GoFunc
	GoFor
	GoStatement
)

type GoNode struct {
	Type GoNodeType
	Body []*GoNode
	Statement string
	Condition string
}
