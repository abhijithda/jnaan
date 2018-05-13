package flippinganimage

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {

	tests := []struct {
		description string
		img         [][]int
		flipImg     [][]int
	}{
		{
			description: "1st Image",
			img:         [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			flipImg:     [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			description: "2nd Image",
			img:         [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			flipImg:     [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)
		resImg := flipAndInvertImage(tc.img)
		if reflect.DeepEqual(resImg, tc.flipImg) {
			t.Errorf("got %v, want %v", resImg, tc.flipImg)
			// t.Errorf("got %v, want %v", resImg[:][:], tc.flipImg[:][:])
		}
	}
}
