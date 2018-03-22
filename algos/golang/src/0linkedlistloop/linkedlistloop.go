package linkedlistloop

import (
	"fmt"
	"log"
	"os"
)

// Node of a single linked list.
type Node struct {
	data int
	next *Node
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

func displayNNodes(head *Node, n int) {
	trav := head
	for i := 0; i < n && trav != nil; i++ {
		log.Printf("%d Node address %v\n", i+1, trav)
		fmt.Printf("-> %v ", trav)
		trav = trav.next
	}
	fmt.Println()
}

func prepareList(maxNodes, iPoint int) (*Node, *Node) {
	var head, trav, iNode *Node
	for i := 1; i <= maxNodes; i++ {
		nn := &Node{
			data: i,
			next: nil,
		}
		if i == iPoint {
			iNode = nn
		}
		if head == nil {
			head = nn
			trav = head
			continue
		}
		trav.next = nn
		trav = trav.next
	}

	if iNode != nil && trav != nil {
		trav.next = iNode
	}
	return head, trav
}

func findLastNode(head *Node) *Node {
	if head == nil {
		fmt.Println("Empty head!")
		return head
	}
	if head.next == nil {
		fmt.Println("Only one node present in the list. No loop found!")
		return head
	}
	trav := head.next
	check := head

	for check.next != trav.next {
		check = check.next
		log.Printf("Check %v -> <- %v Trav\n", check, trav)
		if check == trav {
			log.Println("Moving trav to next node:", trav.next)
			if trav.next == nil {
				fmt.Println("No loop found!")
				return trav
			}
			trav = trav.next
			check = head
		}
	}
	fmt.Printf("Intersecting node is %v, and contains %d\n", trav.next, trav.next.data)
	trav.next = nil
	return trav
}
