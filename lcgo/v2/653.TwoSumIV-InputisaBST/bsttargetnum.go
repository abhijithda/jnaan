package bsttarget

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// type TreeNode tree.Node

func findTarget(root *TreeNode, k int) bool {
	stack := []*TreeNode{root}
	reqNum := map[int]bool{}
	for len(stack) != 0 {
		trav := stack[0]
		if len(stack) == 1 {
			stack = []*TreeNode{}
		} else {
			stack = stack[1:]
		}
		if reqNum[trav.Val] {
			return true
		}

		reqNum[k-trav.Val] = true
		if trav.Left != nil {
			stack = append(stack, trav.Left)
		}
		if trav.Right != nil {
			stack = append(stack, trav.Right)
		}

	}
	return false
}
