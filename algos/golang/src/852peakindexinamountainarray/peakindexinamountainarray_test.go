package peakindexinamountainarray

import (
	"fmt"
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		description  string
		A            []int
		maxPeakIndex int
	}{
		{
			description:  "mountain in 3",
			A:            []int{0, 1, 0},
			maxPeakIndex: 1,
		},
		{
			description:  "mountain in 4",
			A:            []int{0, 2, 1, 0},
			maxPeakIndex: 1,
		},
		{
			description:  "2 mountains in 8",
			A:            []int{0, 2, 1, 0, 3, 5, 1, 4},
			maxPeakIndex: 5,
		},
	}
	for _, tc := range tests {
		fmt.Println(tc.description)

		res := peakIndexInMountainArray(tc.A)
		if res != tc.maxPeakIndex {
			t.Errorf("got %v, want %v", res, tc.maxPeakIndex)
		}
	}
}
