package stackresolution

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

// ValidateInput validates the input string
func ValidateInput(input []rune) (ok bool, index int) {
	if len(input) == 0 { // if input is empty
		return
	}
	stk := newStack()
	for i, r := range input {
		_, okLeft := leftBrackets[r]
		if okLeft {
			stk.push(r)
			continue
		}
		leftBracket, okRight := rightBrackets[r]
		if !okRight {
			index = i
			return
		}
		rStk, err := stk.pop()
		if err != nil || rStk != leftBracket {
			index = i
			return
		}
	}
	ok = stk.isEmpty()
	if !ok {
		index = len(input) - 1
	}
	return
}
