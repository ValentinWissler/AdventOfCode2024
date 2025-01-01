package grid

type NodeEdge struct {
	distance int
	next     *Node
}

func NewNodeEdge(distance int, next *Node) *NodeEdge {
	return &NodeEdge{distance: distance, next: next}
}
