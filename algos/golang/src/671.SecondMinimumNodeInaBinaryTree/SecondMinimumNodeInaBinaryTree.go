package secondminimumnodeinabinarytree

import (
	"internal/tree"
	"log"
)

// TreeNode is a definition of a binary tree node.
type TreeNode = tree.Node

func findSecondMinimumValue(root *TreeNode) int {
	log.Println("\n\n\nGiven root:", root)
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}

	qu := []*TreeNode{root}
	if root.Left != nil {
		qu = append(qu, root.Left)
	}
	if root.Right != nil {
		qu = append(qu, root.Right)
	}
	min := root.Val
	secMin := int(^uint(0) >> 1)
	for len(qu) != 0 {
		e := qu[0]
		log.Printf("Ele: %v; Queue: %v", e, qu)
		if len(qu) == 1 {
			qu = []*TreeNode{}
		} else {
			qu = qu[1:]
		}

		log.Printf("secMin: %d; e.Val: %d; min: %d", secMin, e.Val, min)
		if e.Val <= secMin && e.Val >= min {
			if e.Val != min {
				log.Println("Setting secMin:", e.Val)
				secMin = e.Val
			}
			if e.Left != nil {
				qu = append(qu, e.Left)
			}
			if e.Right != nil {
				qu = append(qu, e.Right)
			}
		}
	}

	if secMin == int(^uint(0)>>1) {
		return -1
	}
	return secMin
}
