package algo

// TODO: do the whole thing

type Node struct {
	LeftNode  *Node
	RightNode *Node
	Data any
}

func NewNode(data any, left *Node, right *Node) *Node {
	return &Node{
		LeftNode: left,
		RightNode: right,
		Data: data,
	}
}

func InsertNode(rootNode *Node, newNode *Node) int {
	return 0
}

func FindElement(rootNode *Node) *Node {
	return nil
}

