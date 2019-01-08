package linkedlistcycle

import (
	"fmt"
	"internal/list"
	"log"
	"os"
)

// ListNode is definition of a node in a singly-linked list.
type ListNode = list.Node

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	trav := head.Next
	check := head
	for check.Next != trav.Next {
		check = check.Next
		log.Printf("Check %v -> <- %v Trav\n", check, trav)
		if check == trav {
			log.Println("Moving trav to next node:", trav.Next)
			if trav.Next == nil {
				fmt.Println("No loop found!")
				return false
			}
			trav = trav.Next
			check = head
		}
	}

	return true
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
