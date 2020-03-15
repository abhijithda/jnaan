package list

import (
	"fmt"
	"log"
)

// Node is a definition of node in a singly-linked list.
type Node struct {
	Val  int
	Next *Node
}

// Display the given linked list.
func Display(trav *Node) {
	list := ""
	for ; trav != nil; trav = trav.Next {
		list += fmt.Sprintf("%d->", trav.Val)
	}
	list += fmt.Sprintf("nil")
	log.Println("List:", list)
}

// Prepare a linked list.
func Prepare(nodes []int) *Node {
	var head, trav *Node

	for i := 0; i < len(nodes); i++ {
		nn := &Node{
			Val:  nodes[i],
			Next: nil,
		}
		if head == nil {
			head = nn
			trav = head
			continue
		}
		trav.Next = nn
		trav = trav.Next
	}

	return head
}
