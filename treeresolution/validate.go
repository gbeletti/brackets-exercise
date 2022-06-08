package treeresolution

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

// ValidateInput validates the input runes
func ValidateInput(input []rune) (ok bool, index int) {
	if len(input) == 0 { // if input is empty
		return
	}
	var tr *tree
	var err error
	for i, r := range input {
		if tr == nil || tr.IsComplete() {
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
