package pkg

type Node struct {
	key      string
	value    int
	parent   *Node
	children map[string]Node
}

func (n Node) AddChild(child Node) {

	n.children[child.key] = child
}

// func (n *Node) GetParent() *Node {
// 	return n.parent
// }
