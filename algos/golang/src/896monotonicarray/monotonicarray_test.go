package monotonicarray

import "testing"

func Test_isMonotonic(t *testing.T) {
	tests := []struct {
		Description string
		Input       []int
		Output      bool
	}{
		{
			Description: "First",
			Input:       []int{1, 2, 2, 3},
			Output:      true,
		},
		{
			Description: "Second",
			Input:       []int{6, 5, 4, 4},
			Output:      true,
		},
		{
			Description: "Third",
			Input:       []int{1, 3, 2},
			Output:      false,
		},
		{
			Description: "Fourth",
			Input:       []int{1, 2, 4, 5},
			Output:      true,
		},
		{
			Description: "Fifth",
			Input:       []int{1, 1, 1},
			Output:      true,
		},
		{
			Description: "Sixth",
			Input:       []int{},
			Output:      true,
		},
	}

	for _, tc := range tests {
		t.Log(tc.Description)
		res := isMonotonic(tc.Input)
		if res != tc.Output {
			t.Errorf("got %v, want %v\n", res, tc.Output)
		}
	}
}
