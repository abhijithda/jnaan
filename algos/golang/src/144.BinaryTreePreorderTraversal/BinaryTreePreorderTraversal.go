package binarytreepreordertraversal

import (
	"fmt"
	itu "internal/treeutils"
	"log"
	"os"
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

// Recursive
func preorderTraversalRecursive(root *itu.TreeNode) []int {
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
func preorderTraversal(root *itu.TreeNode) []int {
	log.Println("\n\n\nGiven root:", root)
	if root == nil {
		return []int{}
	}
	stack := []*itu.TreeNode{root}
	visited := map[*itu.TreeNode]bool{}
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
			stack = []*itu.TreeNode{trav.Left}
			stack = append(stack, tstack...)
		} else {
			if len(stack) > 1 {
				stack = stack[1:]
			} else {
				stack = []*itu.TreeNode{}
			}
			if trav.Right != nil &&
				visited[trav.Right] != true {
				tstack := stack
				stack = []*itu.TreeNode{trav.Right}
				stack = append(stack, tstack...)
			}
		}
		log.Println("preorder:", preorder)
		log.Println("stack:", stack)
	}
	return preorder
}
