package binarytreepaths

import (
	"log"
	"reflect"
	"testing"
)

func popStack(stack []*TreeNode) []*TreeNode {
	if len(stack) == 1 {
		stack = []*TreeNode{}
	} else {
		stack = stack[1:]
	}
	return stack
}

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

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				root: prepareTree([]int{1, 2, 3, -1, 5}),
			},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "Example 2",
			args: args{
				root: prepareTree([]int{1, 2, 3, 4, 5, -1, 7}),
			},
			want: []string{"1->2->4", "1->2->5", "1->3->7"},
		},
		{
			name: "Example 3",
			args: args{
				root: prepareTree([]int{}),
			},
			want: []string{},
		},
		{
			name: "Example 4",
			args: args{
				root: prepareTree([]int{1}),
			},
			want: []string{"1"},
		},
		{
			name: "Example 5",
			args: args{
				root: prepareTree([]int{1, -1, 2}),
			},
			want: []string{"1->2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
