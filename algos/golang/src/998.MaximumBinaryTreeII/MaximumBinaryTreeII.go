package maximumbinarytreeii

import (
	"internal/tree"
	"log"
)

// TreeNode is Definition of a binary tree node.
type TreeNode = tree.Node

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	log.Printf("\n\n\nGiven root: %v, val: %d", root, val)
	nn := &TreeNode{Val: val}
	if root == nil {
		return nn
	}

	cur := root
	prev := cur
	for cur != nil {
		log.Println("Current:", cur)
		if cur.Val < val {
			nn.Left = cur
			log.Printf("Added %v to left of new node", cur)
			if prev == cur {
				// I.e., root node
				root = nn
			} else {
				prev.Right = nn
			}
			log.Printf("Returning root: %v", root)
			return root
		}
		prev = cur
		cur = cur.Right
	}
	if cur == nil {
		prev.Right = nn
		log.Printf("Added right of prev: %v", prev)
	}
	log.Printf("Returning root: %v", root)
	return root
}
