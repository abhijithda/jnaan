package MaximumDepthofBinaryTree

import "internal/tree"

// TreeNode is definition for a binary tree node.
type TreeNode = tree.Node

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := 1 + maxDepth(root.Left)
	rd := 1 + maxDepth(root.Right)
	if ld < rd {
		return rd
	}
	return ld
}
