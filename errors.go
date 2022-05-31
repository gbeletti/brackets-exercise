package main

import "errors"

var (
	// ErrLeftNodeExists cannot insert a left node to a node that already has a left node
	ErrLeftNodeExists = errors.New("left node already exists")
	// ErrRightNodeExists cannot insert a right node to a node that already has a right node
	ErrRightNodeExists = errors.New("right node already exists")
	// ErrLeftNodeNotExists cannot go left because it does not have a left node
	ErrLeftNodeNotExists = errors.New("left node does not exist")
	// ErrRightNodeNotExists cannot go right because it does not have a right node
	ErrRightNodeNotExists = errors.New("right node does not exist")
	// ErrParentNodeNotExists cannot go up because it does not have a parent node
	ErrParentNodeNotExists = errors.New("parent node does not exist")
	// ErrCharacterNotExpected found an expected character
	ErrCharacterNotExpected = errors.New("character not expected")
)
