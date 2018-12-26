package binarytreepreordertraversal

import (
	"fmt"
	"log"
	"os"

	tree "internal/tree"
)

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

// TreeNode is the defintion of node in a binary tree.
type TreeNode = tree.Node

// Recursive
func preorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	trav := root
	preorder := []int{}
	preorder = append(preorder, trav.Val)
	if trav.Left != nil {
		preorder = append(preorder, preorderTraversal(trav.Left)...)
	}
	if trav.Right != nil {
		preorder = append(preorder, preorderTraversal(trav.Right)...)
	}
	return preorder
}

// Iterative
func preorderTraversal(root *TreeNode) []int {
	log.Println("\n\n\nGiven root:", root)
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	visited := map[*TreeNode]bool{}
	preorder := []int{}
	log.Println("preorder:", preorder)
	log.Println("stack:", stack)
	for len(stack) != 0 {
		trav := stack[0]
		if visited[trav] != true {
			preorder = append(preorder, trav.Val)
			visited[trav] = true
		}
		if trav.Left != nil && visited[trav.Left] != true {
			tstack := stack
			stack = []*TreeNode{trav.Left}
			stack = append(stack, tstack...)
		} else {
			if len(stack) > 1 {
				stack = stack[1:]
			} else {
				stack = []*TreeNode{}
			}
			if trav.Right != nil &&
				visited[trav.Right] != true {
				tstack := stack
				stack = []*TreeNode{trav.Right}
				stack = append(stack, tstack...)
			}
		}
		log.Println("preorder:", preorder)
		log.Println("stack:", stack)
	}
	return preorder
}
