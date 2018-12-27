package transposematrix

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	tests := []struct {
		description string
		A           [][]int
		tA          [][]int
	}{
		{
			description: "N x N matrix",
			A:           [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			tA:          [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			description: "N x M matrix",
			A:           [][]int{{1, 2, 3}, {4, 5, 6}},
			tA:          [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}
	for _, tc := range tests {
		fmt.Println(tc.description)
		log.Println(tc.description)

		resA := transpose(tc.A)
		if reflect.DeepEqual(resA, tc.tA) == false {
			t.Errorf("got %v, want %v", resA, tc.tA)
		}
	}
}
