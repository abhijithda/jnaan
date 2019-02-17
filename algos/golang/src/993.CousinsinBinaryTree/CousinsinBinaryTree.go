package cousinsinbinarytree

import "internal/tree"

// TreeNode is definition of a binary tree node.
type TreeNode = tree.Node

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	qu := []*TreeNode{root}

	var foundX, foundY bool
	for len(qu) != 0 {
		newQ := []*TreeNode{}
		for _, p := range qu {
			var tfoundX, tfoundY bool
			if p.Left != nil {
				if p.Left.Val == x {
					foundX = true
					tfoundX = true
				}
				if p.Left.Val == y {
					foundY = true
					tfoundY = true
				}
				newQ = append(newQ, p.Left)
			}
			if p.Right != nil {
				if p.Right.Val == x {
					foundX = true
					tfoundX = true
				}
				if p.Right.Val == y {
					foundY = true
					tfoundY = true
				}
				newQ = append(newQ, p.Right)
			}

			if tfoundX == true && tfoundY == true {
				return false
			}
		}
		if foundX != foundY {
			return false
		} else if foundX == true {
			// Since foundX should be equal to foundY as previous condition failed.
			return true
		}
		qu = newQ
	}
	return false
}
