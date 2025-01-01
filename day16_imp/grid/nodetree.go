package grid

type NodeTree struct {
	root *Node
}

func NewNodeTree(root *Node) *NodeTree {
	return &NodeTree{root: root}
}
