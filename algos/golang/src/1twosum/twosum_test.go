package twosum

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		description string
		inpNums     []int
		inpTarget   int
		expTarget   []int
	}{
		{
			description: "Three elements",
			inpNums:     []int{1, 2, 3},
			inpTarget:   3,
			expTarget:   []int{0, 1},
		},
		{
			description: "Four elements",
			inpNums:     []int{2, 7, 11, 15},
			inpTarget:   9,
			expTarget:   []int{0, 1},
		},
	}

	for _, tc := range tests {
		t.Log("Test:", tc.description)
		inpNums := tc.inpNums
		inpTarget := tc.inpTarget
		expTarget := tc.expTarget

		target := twoSum(inpNums, inpTarget)
		if target[0] != expTarget[0] {
			t.Errorf("\nwant %v, got %v", expTarget, target)
		}
	}

}
