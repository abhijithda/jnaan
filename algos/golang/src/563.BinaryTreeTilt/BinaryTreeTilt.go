package binarytreetilt

import (
	"internal/tree"
	"math"
)

// TreeNode is a definition for a binary tree node.
type TreeNode = tree.Node

func getTilts(t *TreeNode, tilts map[*TreeNode]int) int {
	if t == nil {
		return 0
	}
	lsum := 0
	rsum := 0
	if t.Left != nil {
		lsum = getTilts(t.Left, tilts)
	}
	if t.Right != nil {
		rsum = getTilts(t.Right, tilts)
	}
	tilts[t] = int(math.Abs(float64(lsum) - float64(rsum)))

	return t.Val + lsum + rsum
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	totalTilt := 0
	tilts := map[*TreeNode]int{}
	getTilts(root, tilts)
	for _, t := range tilts {
		totalTilt += t
	}
	return totalTilt
}
