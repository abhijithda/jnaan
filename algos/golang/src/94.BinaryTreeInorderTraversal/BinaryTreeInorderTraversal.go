package binarytreeinordertraversal

import (
	"fmt"
	"log"
	"os"

	it "internal/tree"
)

// Given a binary tree, return the inorder traversal of its nodes' values.
// Follow up: Recursive solution is trivial, could you do it iteratively?

// TreeNode is definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
type TreeNode = it.Node

const logFile = "log.txt"

func init() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()
	log.SetOutput(f)
}

func popStack(stack []*TreeNode) []*TreeNode {
	if len(stack) > 1 {
		return stack[1:]
	}
	return []*TreeNode{}
}

func inorderTraversal(root *TreeNode) []int {
	log.Println("\n\n\nGiven root:", root)
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	visited := map[*TreeNode]bool{}

	inorder := []int{}
	for len(stack) != 0 {
		n := stack[0]
		log.Println("Processing node:", n)

		// Leaf
		if n.Left == nil && n.Right == nil && !visited[n] {
			visited[n] = true
			inorder = append(inorder, n.Val)
			log.Println("Added leaf. Inorder:", inorder)
			stack = popStack(stack)
			log.Println("Removed leaf from stack:", stack)
		} else if ((n.Left != nil && visited[n.Left]) || n.Left == nil) && !visited[n] {
			// Left child is already visited, or there is no left node,
			// and current node is not visited

			// Mark current node as visited, and add it to inorder
			visited[n] = true
			inorder = append(inorder, n.Val)
			log.Println("Added current node!. Inorder:", inorder)
			stack = popStack(stack)

			if n.Right != nil {
				tstack := []*TreeNode{n.Right}
				tstack = append(tstack, stack...)
				stack = tstack
				log.Println("Added right to stack:", stack)
			}
		} else if n.Left != nil {
			if visited[n.Left] {
				visited[n] = true
				log.Printf("Marked node %+v as visited", n.Left)
			} else {
				tstack := []*TreeNode{n.Left}
				tstack = append(tstack, stack...)
				stack = tstack
				log.Println("Added left to stack:", stack)
			}
		}
	}

	return inorder
}
