package binarygap

import (
	"fmt"
	"log"
	"testing"
)

func Test_binaryGap(t *testing.T) {

	tests := []struct {
		description string
		N           int
		dist        int
	}{
		{
			description: "22 -> 2",
			N:           22,
			dist:        2,
		},
		{
			description: "5 -> 2",
			N:           5,
			dist:        2,
		},
		{
			description: "6 -> 1",
			N:           6,
			dist:        1,
		},
		{
			description: "8 -> 0",
			N:           8,
			dist:        0,
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)
		log.Println(tc.description)

		resDist := binaryGap(tc.N)
		if resDist != tc.dist {
			t.Errorf("got %v, want %v", resDist, tc.dist)
		}
	}

}
