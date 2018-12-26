package maximumbinarytree

import (
	"log"
	"reflect"
	"testing"
	// itu "internal/treeutils"
)

func displayTree(t *TreeNode) {
	log.Println("\n\n ---Display")
	visited := map[int]bool{}
	if t == nil {
		log.Println("Empty tree!")
		return
	}
	stack := []*TreeNode{t}
	log.Println("Root:", t.Val)
	for len(stack) != 0 {
		log.Println("Stack: ", stack)
		n := stack[0]
		tstack := stack
		if n.Left != nil && !visited[n.Left.Val] {
			log.Println("Left: ", n.Left.Val)
			stack = []*TreeNode{n.Left}
			// stack = append(stack, n.Left)
			stack = append(stack, tstack...)
		} else if n.Right != nil && !visited[n.Right.Val] {
			log.Println("Right: ", n.Right.Val)
			// Pop n
			tstack = popStack(tstack)
			visited[n.Val] = true
			stack = []*TreeNode{n.Right}
			stack = append(stack, tstack...)
		} else if !visited[n.Val] {
			log.Println("Leaf: ", n.Val)
			// Pop n
			stack = popStack(tstack)
			visited[n.Val] = true
		}
	}
}

func popStack(stack []*TreeNode) []*TreeNode {
	if len(stack) > 1 {
		return stack[1:]
	}
	return []*TreeNode{}
}

// TODO:
// 	Currently doesn't handle case of all left nodes or all right nodes scenario!
func prepareTree(nodesval []int) *TreeNode {
	log.Println("\n\n ---Preparing Tree for", nodesval)
	var root, trav *TreeNode

	level := []*TreeNode{}
	for i, nv := range nodesval {
		if nv == -1 {
			if (i+1)%2 == 1 {
				// Pop element as both left & right have been processed to node.
				level = popStack(level)
			}
			continue
		}
		node := TreeNode{Val: nv}
		if len(level) == 0 {
			root = &node
			log.Println("Root:", nv)
		} else {
			trav = level[0]
			if (i+1)%2 == 0 {
				trav.Left = &node
				log.Println("Left:", nv)
			} else {
				trav.Right = &node
				log.Println("Right:", nv)
				// Pop element as both left & right have been assigned to node.
				level = popStack(level)
			}
		}
		level = append(level, &node)
	}

	displayTree(root)
	return root
}

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{3, 2, 1, 6, 0, 5},
			},
			// want: itu.PrepareTree([]int{6, 3, 5, -1, 2, 0, -1, -1, -1, -1, 1}),
			want: prepareTree([]int{6, 3, 5, -1, 2, 0, -1, -1, 1}),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 3},
			},
			// want: itu.PrepareTree([]int{3, 2}),
			want: prepareTree([]int{3, 2}),
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{3, 2},
			},
			// want: itu.PrepareTree([]int{3, -1, 2}),
			want: prepareTree([]int{3, -1, 2}),
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{1},
			},
			// want: itu.PrepareTree([]int{1}),
			want: prepareTree([]int{1}),
		},
		{
			name: "Example 5",
			args: args{
				nums: []int{},
			},
			// want: itu.PrepareTree([]int{}),
			want: prepareTree([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
