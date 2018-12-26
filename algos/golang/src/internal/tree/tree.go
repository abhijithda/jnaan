package tree

import "log"

// Node is definition of node in a binary tree.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// DisplayTree displays tree from the given node
func DisplayTree(t *Node) {
	log.Println("\n\n ---Display")
	visited := map[int]bool{}
	if t == nil {
		log.Println("Empty tree!")
		return
	}
	stack := []*Node{t}
	log.Println("Root:", t.Val)
	for len(stack) != 0 {
		log.Println("Stack: ", stack)
		n := stack[0]
		tstack := stack
		if n.Left != nil && !visited[n.Left.Val] {
			log.Println("Left: ", n.Left.Val)
			stack = []*Node{n.Left}
			// stack = append(stack, n.Left)
			stack = append(stack, tstack...)
		} else if n.Right != nil && !visited[n.Right.Val] {
			log.Println("Right: ", n.Right.Val)
			// Pop n
			tstack = popStack(tstack)
			visited[n.Val] = true
			stack = []*Node{n.Right}
			stack = append(stack, tstack...)
		} else if !visited[n.Val] {
			log.Println("Leaf: ", n.Val)
			// Pop n
			stack = popStack(tstack)
			visited[n.Val] = true
		}
	}
}

// popStack TODO: removes & returns the first element if present from the stack.
func popStack(stack []*Node) []*Node {
	if len(stack) > 1 {
		return stack[1:]
	}
	return []*Node{}
}

// PrepareTree prepares a tree from an array.
// 	To specify nil in integer list, specify `-1` in it's place.
// TODO:
// 	Currently doesn't handle case of all left nodes or all right nodes scenario!
func PrepareTree(nodesval []int) *Node {
	log.Println("\n\n ---Preparing Tree for", nodesval)
	var root, trav *Node

	level := []*Node{}
	for i, nv := range nodesval {
		if nv == -1 {
			if (i+1)%2 == 1 {
				// Pop element as both left & right have been processed to node.
				level = popStack(level)
			}
			continue
		}
		node := Node{Val: nv}
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

	DisplayTree(root)
	return root
}
