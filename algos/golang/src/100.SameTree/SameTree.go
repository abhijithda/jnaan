package sametree

import (
	"fmt"
	"internal/tree"
	"log"
	"os"
)

// TreeNode is definition of node in a binary tree node.
type TreeNode = tree.Node

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	log.Printf("\n\n\nGiven p:%v; q:%v", p, q)
	if (p == nil || q == nil) && p != q {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		// log.Printf("p.Val != q.Val...")
		return false
	}
	if (p.Left == nil || q.Left == nil) && p.Left != q.Left {
		// log.Printf("p.Left != q.Left...")
		return false
	}
	if (p.Right == nil || q.Right == nil) && p.Right != q.Right {
		// log.Printf("p.Right != q.Right...")
		return false
	}
	if p.Left != nil {
		if !isSameTree(p.Left, q.Left) {
			return false
		}
	}
	if p.Right != nil {
		if !isSameTree(p.Right, q.Right) {
			return false
		}
	}
	return true
}
