package univaluedbinarytree

import "internal/tree"

// TreeNode is a node of a binary tree
type TreeNode = tree.Node

func isNodeValue(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	if node.Left != nil {
		if isNodeValue(node.Left, val) == false {
			return false
		}
	}
	if node.Right != nil {
		if isNodeValue(node.Right, val) == false {
			return false
		}
	}
	return true
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isNodeValue(root, root.Val)
}
