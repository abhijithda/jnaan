package faircandyswap

import (
	"log"
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	tests := []struct {
		description string
		A           []int
		B           []int
		output      []int
	}{
		{
			description: "first sample",
			A:           []int{1, 1},
			B:           []int{2, 2},
			output:      []int{1, 2},
		},
		{
			description: "second sample",
			A:           []int{1, 2},
			B:           []int{2, 3},
			output:      []int{2, 3},
		},
		{
			description: "third sample",
			A:           []int{2},
			B:           []int{1, 3},
			output:      []int{2, 3},
		},
		{
			description: "fourth sample",
			A:           []int{1, 2, 5},
			B:           []int{2, 4},
			output:      []int{5, 4},
		},
	}

	for _, tc := range tests {
		log.Println(tc.description)

		res := fairCandySwap(tc.A, tc.B)
		if reflect.DeepEqual(res, tc.output) == false {
			t.Errorf("%s:\n A: %+v B: %+v; got %+v, want %+v", tc.description, tc.A, tc.B, res, tc.output)
		}

	}

}
