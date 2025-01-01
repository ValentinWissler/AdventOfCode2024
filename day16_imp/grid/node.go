package grid

// A node is a junction in the maze
type Node struct {
	pos  Pos         // The starting position of our node
	edge []*NodeEdge // adjacent nodes
}

func NewNode(curr Pos) *Node {
	return &Node{pos: curr, edge: make([]*NodeEdge, 0)}
}

func (n *Node) addEdge(edge *NodeEdge) {
	n.edge = append(n.edge, edge)
}
