package linkedlistcycle

import (
	"fmt"
	"log"
	"testing"
)

func prepareList(maxListNodes, iPoint int) *ListNode {
	var head, trav, iListNode *ListNode
	for i := 1; i <= maxListNodes; i++ {
		nn := &ListNode{
			Val:  i,
			Next: nil,
		}
		if i == iPoint {
			iListNode = nn
		}
		if head == nil {
			head = nn
			trav = head
			continue
		}
		trav.Next = nn
		trav = trav.Next
	}

	if iListNode != nil && trav != nil {
		trav.Next = iListNode
	}
	return head
}

func displayNListNodes(head *ListNode, n int) {
	trav := head
	for i := 0; i < n && trav != nil; i++ {
		log.Printf("%d ListNode address %v\n", i+1, trav)
		fmt.Printf("-> %v ", trav)
		trav = trav.Next
	}
	fmt.Println()
}

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				head: prepareList(1, 0),
			},
			want: false,
		},
		{
			name: "Example 2",
			args: args{
				head: prepareList(1, 1),
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				head: prepareList(5, 0),
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				head: prepareList(5, 3),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
