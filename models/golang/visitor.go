package golang

import (
	"github.com/colintle/BrainForge/models/bf"
)

type VisitorMethods struct{
	Enter func(node *bf.BfNode, parent *GoNode)
	Exit func (node *bf.BfNode, parent *GoNode)
}
