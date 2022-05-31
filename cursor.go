package main

type cursor struct {
	Current *node
}

func (c *cursor) goLeft() (err error) {
	if c.Current.left == nil {
		err = ErrLeftNodeNotExists
		return
	}
	c.Current = c.Current.left
	return
}

func (c *cursor) goRight() (err error) {
	if c.Current.right == nil {
		err = ErrRightNodeNotExists
		return
	}
	c.Current = c.Current.right
	return
}

func (c *cursor) goUp() (err error) {
	if c.Current.parent == nil {
		err = ErrParentNodeNotExists
		return
	}
	c.Current = c.Current.parent
	return
}

func (c *cursor) getLeft() (node *node) {
	return c.Current.left
}

func (c *cursor) getRight() (node *node) {
	return c.Current.right
}

func newCursor(n *node) *cursor {
	return &cursor{
		Current: n,
	}
}
