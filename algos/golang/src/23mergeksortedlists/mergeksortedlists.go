package mergeksortedlists

import (
	"fmt"
	"log"
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

// ListNode is definition of a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var mergedListH *ListNode

	min := int(^uint(0) >> 1)
	log.Println("Minimum value:", min)
	minNodeIndex := -1
	trav := make([]*ListNode, len(lists))
	for i := range lists {
		log.Printf("List %d: %+v", i, lists[i])
		trav[i] = lists[i]
		if trav[i] == nil {
			log.Printf("Skipping list %d as it's empty!", i)
			continue
		}
		if min > trav[i].Val {
			min = trav[i].Val
			minNodeIndex = i
		}
	}
	if minNodeIndex == -1 {
		log.Println("List of pointers is empty or contains only list of empty lists")
		return mergedListH
	}
	log.Println("Min node:", trav[minNodeIndex])
	mergedListH = trav[minNodeIndex]
	trav[minNodeIndex] = trav[minNodeIndex].Next

	mergedList := mergedListH
	traversed := 0
	for traversed != len(lists) {
		min := int(^uint(0) >> 1)
		minNodeIndex = -1
		traversed = 0
		for i := range trav {
			if trav[i] == nil {
				traversed++
				continue
			}
			if min > trav[i].Val {
				min = trav[i].Val
				minNodeIndex = i
			}
		}
		if minNodeIndex == -1 {
			break
		}
		log.Println("Min node:", trav[minNodeIndex])
		mergedList.Next = trav[minNodeIndex]
		mergedList = mergedList.Next
		trav[minNodeIndex] = trav[minNodeIndex].Next
	}

	return mergedListH
}
