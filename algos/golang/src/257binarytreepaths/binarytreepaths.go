package binarytreepaths

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// TreeNode definition of a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func getPath2Leaf(t *TreeNode, nodes []int, path *[]string) {
	if t == nil {
		return
	}
	if t.Left == nil && t.Right == nil {
		newPath := ""
		for _, n := range nodes {
			newPath += strconv.Itoa(n) + "->"
		}
		newPath += strconv.Itoa(t.Val)
		log.Println("New path:", newPath)
		*path = append(*path, newPath)
		log.Println("All paths for so far!:", path)
		return
	}
	if t.Left != nil {
		getPath2Leaf(t.Left, append(nodes, t.Val), path)
	}
	if t.Right != nil {
		getPath2Leaf(t.Right, append(nodes, t.Val), path)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	nodes := []int{}
	getPath2Leaf(root, nodes, &paths)

	log.Println("All paths", paths)
	return paths
}
