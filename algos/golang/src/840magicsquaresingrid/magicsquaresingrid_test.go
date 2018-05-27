package magicsquaresingrid

import (
	"fmt"
	"testing"
)

func Test_numMagicSquaresInside(t *testing.T) {
	tests := []struct {
		description string
		grid        [][]int
		res         int
	}{
		{
			description: "One magic grid!",
			grid:        [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}},
			res:         1,
		},
		{
			description: "10 in magic grid!",
			grid:        [][]int{{10, 3, 5}, {1, 6, 11}, {7, 9, 2}},
			res:         0,
		},
		{
			description: "Zero in magic grid!",
			grid:        [][]int{{7, 0, 5}, {2, 4, 6}, {3, 8, 1}},
			res:         0,
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)
		res := numMagicSquaresInside(tc.grid)
		if res != tc.res {
			t.Errorf("got %v, want %v", res, tc.res)
		} else {
			fmt.Println(tc.description, ": Pass")
		}
	}
}
