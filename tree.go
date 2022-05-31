package main

type tree struct {
	root   *node
	cursor *cursor
}

// IsComplete returns true if the root has a left and right nodes
func (t *tree) IsComplete() bool {
	return t.root.IsComplete()
}

func newTree() *tree {
	root := newNode(nil, ' ', 0)
	return &tree{
		root:   root,
		cursor: newCursor(root),
	}
}
