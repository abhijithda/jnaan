package keysandrooms

import (
	"fmt"
	"testing"
)

func Test_canVisitAllRooms(t *testing.T) {
	tests := []struct {
		description string
		rooms       [][]int
		canVisit    bool
	}{
		{
			description: "Yes, can visit",
			rooms:       [][]int{{1}, {2}, {3}, {}},
			canVisit:    true,
		},
		{
			description: "No, can't visit",
			rooms:       [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			canVisit:    false,
		},
		{
			description: "Yes, but no rooms to visit",
			rooms:       [][]int{{}},
			canVisit:    true,
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)
		visited := canVisitAllRooms(tc.rooms)
		if visited != tc.canVisit {
			t.Errorf("got %v, want %v", visited, tc.canVisit)
		} else {
			fmt.Println(tc.description, ": Pass")
		}
	}
}
