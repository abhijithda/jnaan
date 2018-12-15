package leafsimilartrees

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

// TreeNode is the definition for a binary tree node.
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

// func getLeafNodes(t *TreeNode) []int {
// 	leaves := []int{}
// 	stack := []*TreeNode{t}
// 	for len(stack) != 0 {
// 		log.Println("Stack: ", stack)
// 		tstack := stack
// 		n := tstack[0]
// 		if n.Left != nil {
// 			stack = append(stack, n.Left)
// 			stack = append(stack, tstack...)
// 		} else if n.Right != nil {
// 			// Pop n
// 			stack = popStack(stack)
// 			stack = append(stack, n.Right)
// 		} else {
// 			// Pop n
// 			stack = popStack(stack)
// 			leaves = append(leaves, n.Val)
// 		}
// 	}
// 	log.Println("All leaves:", leaves)
// 	return leaves
// }

// func leafSimilar_MLE(root1 *TreeNode, root2 *TreeNode) bool {
// 	leaf1 := getLeafNodes(root1)
// 	leaf2 := getLeafNodes(root2)
// 	return reflect.DeepEqual(leaf1, leaf2)
// }

func getLeafValues(t *TreeNode, lv *[]int) {
	log.Printf("Given t: %+v; lv: %+v", t, lv)
	if t == nil {
		return
	}
	if t.Left == nil && t.Right == nil {
		*lv = append(*lv, t.Val)
	}
	if t.Left != nil {
		getLeafValues(t.Left, lv)
	}
	if t.Right != nil {
		getLeafValues(t.Right, lv)
	}
	log.Printf("Result t: %+v; lv: %+v", t, lv)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	lv1, lv2 := []int{}, []int{}
	getLeafValues(root1, &lv1)
	getLeafValues(root2, &lv2)

	log.Println("Leaves1:", lv1)
	log.Println("Leaves2:", lv2)
	return reflect.DeepEqual(lv1, lv2)
}
