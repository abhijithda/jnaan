package increasingordersearchtree

import (
	"fmt"
	it "internal/tree"
	"log"
	"os"
)

func init() {
	const logFile = "log.txt"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()
	log.SetOutput(f)
}

// TreeNode : Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
type TreeNode = it.Node

// Iterative solution
func increasingBSTOld(root *TreeNode) *TreeNode {
	log.Println("\n\n\nGiven root:", root)
	var newroot, nroot *TreeNode
	var rootFound bool

	stack := []*TreeNode{root}
	visited := map[*TreeNode]bool{}
	for len(stack) != 0 {
		trav := stack[0]
		visited[trav] = true

		if trav.Left != nil && !visited[trav.Left] {
			tstack := stack
			stack = []*TreeNode{trav.Left}
			stack = append(stack, tstack...)
		} else {
			if len(stack) > 1 {
				stack = stack[1:]
			} else {
				stack = []*TreeNode{}
			}
			log.Println("New trav:", nroot)

			if !rootFound {
				// First leaf node.
				newroot = trav
				nroot = newroot
				rootFound = true
				// Leaf node nothing to process after popping from stack,
				// 	so continue...
				continue
			}

			nroot.Right = trav
			nroot = nroot.Right

			if trav.Right != nil && !visited[trav.Right] {
				tstack := stack
				stack = []*TreeNode{trav.Right}
				stack = append(stack, tstack...)
			}
		}
		log.Println("Stack:", stack)
		// log.Println("New trav:", nroot)
	}

	it.DisplayTree(newroot)
	return newroot
}

// Recursive solution
func increasingBST(root *TreeNode) *TreeNode {
	log.Println("\n\n\nGiven root:", root)
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	var nroot *TreeNode
	if root.Left != nil {
		log.Println("Left:", root.Left)
		nroot = increasingBST(root.Left)
		trav := nroot
		for trav.Right != nil {
			trav = trav.Right
		}
		trav.Right = root
		root.Left = nil
	}

	if root.Right != nil {
		log.Println("Right:", root.Right)
		// Info: Reassign cur node's (i.e., root's) left cover cases
		//   where left(L) node becomes root after recursive call!
		// 	 I.e., 7<-L-(8)-R->9 becomes (7)-R->8-R->9
		ttrav := increasingBST(root.Right)
		root.Right = ttrav
		if nroot == nil {
			nroot = root
			log.Println("nroot was nil! Now:", nroot)
			log.Println("ttrav:", ttrav)
		}
	}

	it.DisplayTree(nroot)
	log.Println("\n-----")
	return nroot
}
