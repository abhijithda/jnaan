package linkedlistloop

import "testing"

func Test_DisplayNNodes(t *testing.T) {
	tests := []struct {
		description string
		maxNodes    int
		iPoint      int
	}{
		{
			description: "0 Nodes",
		},
		{
			description: "1 Node no loop",
			maxNodes:    1,
		},
		{
			description: "1 Node",
			maxNodes:    1,
			iPoint:      1,
		},
		{
			description: "5 Nodes no loop",
			maxNodes:    5,
			iPoint:      0,
		},
		{
			description: "5 Nodes",
			maxNodes:    5,
			iPoint:      3,
		},
		{
			description: "100 Nodes",
			maxNodes:    100,
			iPoint:      9,
		},
	}

	for _, tc := range tests {
		t.Log(tc.description)
		maxNodes := tc.maxNodes
		iPoint := tc.iPoint

		head, last := prepareList(maxNodes, iPoint)
		// displayNNodes(head, maxNodes)
		lastNode := findLastNode(head)
		// displayNNodes(head, maxNodes)
		if last != lastNode {
			t.Errorf("want %v, got %v", last, lastNode)
		}
	}
}
