package tree

import "log"

// Node is definition of node in a binary tree.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Display displays tree from the given node
func Display(t *Node) {
	log.Println("\n\n\n----- Display -----")
	if t == nil {
		log.Println("Empty tree!")
		return
	}
	log.Println("Root:", t.Val)
	stack := []*Node{t}
	visited := map[*Node]bool{}
	for len(stack) != 0 {
		log.Println("Stack: ", stack)
		n := stack[0]
		visited[n] = true
		if n.Left != nil && !visited[n.Left] {
			log.Println("Left: ", n.Left)
			tstack := stack
			stack = []*Node{n.Left}
			stack = append(stack, tstack...)
		} else {
			// Pop n
			stack = popStack(stack)
			if n.Right != nil && !visited[n.Right] {
				log.Println("Right: ", n.Right)
				tstack := stack
				stack = []*Node{n.Right}
				stack = append(stack, tstack...)
			}
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

// Create prepares a tree from an array which is in BFS order.
// 	To specify nil in integer list, specify `-1` in it's place.
func Create(nodesval []int) *Node {
	log.Println("\n\n----- Creating Tree for", nodesval)
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

	Display(root)
	return root
}
