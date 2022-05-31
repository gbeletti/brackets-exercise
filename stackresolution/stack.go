package stackresolution

import (
	"errors"
)

var (
	// ErrStackEmpty is returned when the stack is empty
	ErrStackEmpty = errors.New("stack is empty")
)

type stack struct {
	data []rune
}

func (s *stack) push(r rune) {
	s.data = append(s.data, r)
}

func (s *stack) pop() (rune, error) {
	if len(s.data) == 0 {
		return ' ', ErrStackEmpty
	}
	last := len(s.data) - 1
	r := s.data[last]
	s.data = s.data[:last]
	return r, nil
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func newStack() *stack {
	return &stack{}
}
