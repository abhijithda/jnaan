package removeduplicatesfromsortedlist

import (
	"internal/list"
)

// ListNode is a node in a binary linked list.
type ListNode = list.Node

func deleteDuplicates(head *ListNode) *ListNode {
	trav := head
	if trav == nil {
		return nil
	}
	prev := trav
	prevVal := trav.Val
	trav = trav.Next
	for trav != nil {
		if trav.Val == prevVal {
			// Node already present
			prev.Next = trav.Next
		} else {
			prevVal = trav.Val
			prev = trav
		}
		trav = trav.Next
	}
	return head
}

func deleteDuplicatesWithVisit(head *ListNode) *ListNode {
	trav := head
	prev := trav
	present := map[int]bool{}
	for trav != nil {
		if _, ok := present[trav.Val]; ok {
			// Node already present
			prev.Next = trav.Next
		} else {
			present[trav.Val] = true
			prev = trav
		}
		trav = trav.Next
	}
	return head
}
