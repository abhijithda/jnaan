package backspacestringcompare

import (
	"fmt"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		description string
		S           string
		T           string
		res         bool
	}{
		{
			description: "Equal string with single backspace",
			S:           "ab#c",
			T:           "ad#c",
			res:         true,
		},
		{
			description: "Equal string with two backspaces at end",
			S:           "ab##",
			T:           "c#d#",
			res:         true,
		},
		{
			description: "Equal string with two backspaces in b/n",
			S:           "a##c",
			T:           "#a#c",
			res:         true,
		},
		{
			description: "Strings not equal with one backspace in b/n",
			S:           "a#c",
			T:           "b",
			res:         false,
		},
		{
			description: "Failing TC",
			S:           "xywrrmp",
			T:           "xywrrmu#p",
			res:         true,
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)

		res := backspaceCompare(tc.S, tc.T)
		if res != tc.res {
			t.Errorf("got %v, want %v", res, tc.res)
		}
	}
}
