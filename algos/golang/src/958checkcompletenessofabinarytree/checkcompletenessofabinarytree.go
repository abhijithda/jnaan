package checkcompletenessofabinarytree

import (
	"log"
	"math"
)

// TreeNode structure below is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	log.Println("\n\n\nGiven root:", root)
	t := root

	if t == nil {
		return true
	}
	queue := []*TreeNode{t}
	level := 0
	numOfNodes := 1
	noMoreNodes := false

	for len(queue) != 0 {
		n := queue[0]
		log.Println("Queue:", queue)
		if numOfNodes == int(math.Pow(2, float64(level))) {
			level++
			log.Printf("Incremented level to %d as max no. of nodes reached.\n", level)
			numOfNodes = 0
		}

		log.Println("Processing node:", n)
		// Shift element
		if len(queue) == 1 {
			queue = []*TreeNode{}
		} else {
			queue = queue[1:]
		}

		if n.Left == nil && n.Right != nil {
			log.Println("Left is null while right is not!")
			return false
		}
		if n.Left != nil && n.Right == nil {
			log.Println("Left is not null while right is!")
			if noMoreNodes {
				return false
			}
			log.Println("Left is not null while right is! Setting no more nodes!!!")
			queue = append(queue, n.Left)
			noMoreNodes = true
			continue
		}
		if n.Left != nil && n.Right != nil {
			log.Println("Both nodes are not null for", n)
			if noMoreNodes {
				log.Println("But no more nodes expected!")
				return false
			}
			queue = append(queue, n.Left, n.Right)
			numOfNodes += 2
		}
		if n.Left == nil && n.Right == nil {
			log.Println("Left and righ are null. Leaf node. Setting no more new level nodes!!!")
			noMoreNodes = true
		}

	}

	return true
}
