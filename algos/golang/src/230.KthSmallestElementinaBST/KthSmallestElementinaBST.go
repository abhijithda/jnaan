package kthsmallestelementinabst

import (
	"internal/tree"
)

// TreeNode is definition of a binary tree node.
type TreeNode = tree.Node

func kthSmallest(root *TreeNode, k int) int {
	// log.Printf("\n\n\nroot: %v, k: %d", root, k)
	var kth int

	st := []*TreeNode{root}
	v := map[*TreeNode]bool{}
	for len(st) != 0 {
		// log.Println("Stack:", st)
		e := st[0]
		v[e] = true

		if e.Left != nil && v[e.Left] != true {
			// log.Println("Push left on to stack:", e.Left)
			tst := []*TreeNode{e.Left}
			tst = append(tst, st...)
			st = tst
		} else {
			k--
			// log.Printf("k: %d; e: %v", k, e)
			if k == 0 {
				kth = e.Val
				break
			}

			// Pop `e` from stack.
			if len(st) == 1 {
				st = []*TreeNode{}
			} else {
				st = st[1:]
			}

			if e.Right != nil {
				// Push on to stack
				// log.Println("Push right on to stack:", e.Right)
				tst := []*TreeNode{e.Right}
				tst = append(tst, st...)
				st = tst
			}
		}
	}
	return kth
}
