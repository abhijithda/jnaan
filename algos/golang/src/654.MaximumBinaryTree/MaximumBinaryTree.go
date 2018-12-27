package maximumbinarytree

import (
	"fmt"
	"log"
	"os"

	"internal/tree"
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

// TreeNode is the definition of a node in a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
type TreeNode = tree.Node

func createMBTree(nums []int, root *TreeNode, left bool) *TreeNode {
	log.Printf("Arguments nums: %+v, root: %+v", nums, root)
	if len(nums) == 0 {
		return nil
	}
	maxi := 0
	for i, n := range nums {
		if nums[maxi] < n {
			maxi = i
		}
	}

	nn := &TreeNode{Val: nums[maxi]}
	if root == nil {
		root = nn
	} else if left {
		root.Left = nn
		root = root.Left
	} else {
		root.Right = nn
		root = root.Right
	}
	log.Printf("New sub-root (%d) node: %v", nums[maxi], root)
	if maxi > 0 {
		log.Printf("START: Left of sub-root (%d): %v", nums[maxi], root)
		// root.Left = createMBTree(nums[0:maxi], root, true)
		createMBTree(nums[0:maxi], root, true)
		log.Printf("END: Left of sub-root (%d): %v", nums[maxi], root.Left)
	}
	if maxi < len(nums)-1 {
		log.Printf("START: Right of sub-root (%d): %v", nums[maxi], root)
		// root.Right = createMBTree(nums[maxi+1:len(nums)], root, false)
		createMBTree(nums[maxi+1:len(nums)], root, false)
		log.Printf("END: Right of sub-root (%d): %v", nums[maxi], root.Right)
	}
	return root
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	log.Println("\n\n\nGiven nums:", nums)
	if len(nums) == 0 {
		return nil
	}

	root := createMBTree(nums, nil, true)
	log.Println("Root: ", root)
	// tree.Display(root)
	return root
}
