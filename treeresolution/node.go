package treeresolution

// node is a struct that represents a node in a tree.
type node struct {
	parent    *node
	left      *node
	right     *node
	index     int
	character rune
}

// SetLeft sets the left node of the current node if it has no left node
func (n *node) SetLeft(left *node) (err error) {
	if n.left != nil {
		err = ErrLeftNodeExists
		return
	}
	n.left = left
	return
}

// SetRight sets the right node of the current node if it has no right node
func (n *node) SetRight(right *node) (err error) {
	if n.right != nil {
		err = ErrRightNodeExists
		return
	}
	n.right = right
	return
}

// IsComplete returns true if the node has a left and right nodes
func (n *node) IsComplete() bool {
	return n.left != nil && n.right != nil
}

func (n *node) getCharacter() rune {
	return n.character
}

func newNode(parent *node, character rune, index int) *node {
	return &node{
		parent:    parent,
		character: character,
		index:     index,
	}
}
