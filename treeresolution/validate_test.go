package treeresolution_test

import (
	"testing"

	"github.com/gbeletti/brackets-exercise/treeresolution"
)

func TestValidateInput(t *testing.T) {
	tcases := []struct {
		name  string
		input string
		ok    bool
		index int
	}{
		{
			name:  "01 - empty",
			input: "",
			ok:    false,
			index: 0,
		},
		{
			name:  "02 - invalid brackets",
			input: "{",
			ok:    false,
			index: 0,
		},
		{
			name:  "03 - valid brackets",
			input: "{[]}",
			ok:    true,
			index: 0,
		},
		{
			name:  "04 - valid brackets sequential",
			input: "{[]}(){}[{()}]",
			ok:    true,
			index: 0,
		},
		{
			name:  "05 - incomplete brackets",
			input: "(([]",
			ok:    false,
			index: 3,
		},
		{
			name:  "06 - invalid brackets",
			input: "(((]",
			ok:    false,
			index: 3,
		},
	}
	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			testValidateInput(t, tcase.input, tcase.ok, tcase.index)
		})
	}
}

func testValidateInput(t *testing.T, input string, expectedOK bool, expectedIndex int) {
	inp := []rune(input)
	ok, index := treeresolution.ValidateInput(inp)
	if ok != expectedOK {
		t.Errorf("expected ok to be %t, but got %t", expectedOK, ok)
	}
	if index != expectedIndex {
		t.Errorf("expected index to be %d, but got %d", expectedIndex, index)
	}
}
