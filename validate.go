package main

import "log"

var leftBrackets = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

var rightBrackets = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func validateInput(input []rune) (ok bool, index int) {
	inputLen := len(input)
	if inputLen == 0 { // if input is empty
		return
	}
	tr := newTree()
	var err error
	for i, r := range input {
		if tr.IsComplete() {
			log.Printf("tree is complete, going to next character i [%d] rune [%c]", i, r)
			tr = newTree()
		}
		nextNode := newNode(tr.cursor.Current, r, i)
		_, okLeft := leftBrackets[r]
		if okLeft { // character is left, should go left
			err = shouldGoLeft(tr, nextNode)
			if err != nil {
				index = i
				return
			}
			continue
		}
		leftBracket, okRight := rightBrackets[r]
		if !okRight {
			log.Printf("invalid character:[%c]\n", r)
			index = i
			return
		}
		err = shouldGoUpAndRight(tr, leftBracket, nextNode)
		if err != nil {
			index = i
			return
		}
	}
	ok = tr.IsComplete()
	if !ok {
		index = len(input) - 1
	}
	return
}

func shouldGoLeft(tr *tree, nextNode *node) (err error) {
	err = tr.cursor.Current.SetLeft(nextNode)
	if err != nil {
		return
	}
	err = tr.cursor.goLeft()
	return
}

func shouldGoUpAndRight(tr *tree, leftBracket rune, nextNode *node) (err error) {
	tr.cursor.goUp()
	// character is right, maybe go right
	lNode := tr.cursor.getLeft()
	if lNode == nil {
		err = ErrLeftNodeNotExists
		return
	}
	if lNode.getCharacter() != leftBracket {
		err = ErrCharacterNotExpected
		return
	}
	err = tr.cursor.Current.SetRight(nextNode)
	return
}
