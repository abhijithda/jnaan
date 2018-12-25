package printbinarytree

import (
	"fmt"
	"log"
	"math"
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

// TreeNode is definition of node in a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getHeight(root *TreeNode) int {
	if root == nil {
		// Empty/No tree!
		return 0
	} else if root.Left == nil && root.Right == nil {
		// Leaf node
		return 1
	}

	if root.Left != nil && root.Right != nil {
		if 1+getHeight(root.Left) > 1+getHeight(root.Right) {
			return 1 + getHeight(root.Left)
		}
		return 1 + getHeight(root.Right)
	}
	if root.Left != nil {
		return 1 + getHeight(root.Left)
	}
	if root.Right != nil {
		return 1 + getHeight(root.Right)
	}

	return 0
}

func updateNodeEntry(node *TreeNode, ar [][]string, r, sc, ec int) {
	log.Printf("\n\n\nGiven node: %+v; ar: %+v; r:%d; sc:%d; ec:%d",
		node, ar, r, sc, ec)

	if node == nil || r < 0 {
		log.Println("Node is null or r < 0")
		return
	}

	c := (sc + ec) / 2
	log.Printf("Row: %d; Col: %d; Node: %+v", r, c, node)
	ar[r][c] = fmt.Sprintf("%d", node.Val)
	log.Printf("Row: %d, Col: %d updated to %d; ar: %+v",
		r, c, node.Val, ar)
	if node.Left == nil && node.Right == nil {
		// Leaf node
		log.Println("Node is leaf node:", node)
		return
	}

	if node.Left != nil {
		updateNodeEntry(node.Left, ar, r+1, sc, c-1)
	}
	if node.Right != nil {
		updateNodeEntry(node.Right, ar, r+1, c+1, ec)
	}
}

func printTree(root *TreeNode) [][]string {
	log.Println("Given root:", root)
	h := getHeight(root)
	log.Println("Height of tree:", h)

	cols := int(math.Pow(2, float64(h))) - 1
	log.Println("Number of cols:", cols)
	ar := make([][]string, h)
	for c := range ar {
		ar[c] = make([]string, cols)
	}

	updateNodeEntry(root, ar, 0, 0, cols-1)
	return ar
}
