package symmetrictree

import (
	"fmt"
	"internal/tree"
	"log"
	"os"
)

// TreeNode is Definition for a binary tree node.
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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return cmpSubTrees(root.Left, root.Right)
}

func cmpSubTrees(leftST, rightST *TreeNode) bool {
	if leftST == nil && rightST == nil {
		return true
	}
	if (leftST == nil && rightST != nil) ||
		(leftST != nil && rightST == nil) {
		return false
	}
	if leftST.Val != rightST.Val {
		return false
	}
	return cmpSubTrees(leftST.Left, rightST.Right) &&
	cmpSubTrees(leftST.Right, rightST.Left)
	
}
