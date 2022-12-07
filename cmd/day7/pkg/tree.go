package pkg

type Node struct {
	Key      string
	Value    int
	Parent   *Node
	Children map[string]*Node
}

func (n *Node) AddChild(child *Node) {

	n.Children[child.Key] = child
}

// func (n *Node) GetParent() *Node {
// 	return n.parent
// }
