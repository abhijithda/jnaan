package findandreplaceinstring

import (
	"fmt"
	"testing"
)

func Test_findReplaceString(t *testing.T) {
	tests := []struct {
		description string
		S           string
		indexes     []int
		sources     []string
		targets     []string
		res         string
	}{
		{
			description: "1st Image",
			S:           "abcd",
			indexes:     []int{0, 2},
			sources:     []string{"a", "cd"},
			targets:     []string{"eee", "ffff"},
			res:         "eeebffff",
		},
		{
			description: "2nd Image",
			S:           "abcd",
			indexes:     []int{0, 2},
			sources:     []string{"ab", "ec"},
			targets:     []string{"eee", "ffff"},
			res:         "eeecd",
		},
		{
			description: "Non-sorted index",
			S:           "vmokgggqzp",
			indexes:     []int{3, 5, 1},
			sources:     []string{"kg", "ggq", "mo"},
			targets:     []string{"s", "so", "bfr"},
			res:         "vbfrssozp",
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)
		res := findReplaceString(tc.S, tc.indexes, tc.sources, tc.targets)
		t.Logf("Result: %v; Expected: %v", res, tc.res)
		if res != tc.res {
			t.Errorf("got %v, want %v", res, tc.res)
		}
		fmt.Println()
	}

}
