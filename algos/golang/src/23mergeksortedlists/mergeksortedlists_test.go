package mergeksortedlists

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeKListsOld(t *testing.T) {
	a1n3 := &ListNode{
		Val:  5,
		Next: nil,
	}
	a1n2 := &ListNode{
		Val:  4,
		Next: a1n3,
	}
	a1n1 := &ListNode{
		Val:  1,
		Next: a1n2,
	}
	l1 := a1n1
	a2n3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	a2n2 := &ListNode{
		Val:  3,
		Next: a2n3,
	}
	a2n1 := &ListNode{
		Val:  1,
		Next: a2n2,
	}
	l2 := a2n1
	a3n2 := &ListNode{
		Val:  6,
		Next: nil,
	}
	a3n1 := &ListNode{
		Val:  2,
		Next: a3n2,
	}
	l3 := a3n1

	listA := []*ListNode{l1, l2, l3}

	mergedList := mergeKLists(listA)
	trav := mergedList
	for trav != nil {
		t.Log(trav.Val)
		trav = trav.Next
	}

}

func displayList(t *testing.T, trav *ListNode) {
	list := ""
	for ; trav != nil; trav = trav.Next {
		list += fmt.Sprintf("%d->", trav.Val)
	}
	list += fmt.Sprintf("nil")
	t.Log("List:", list)
}

func prepareList(nodes []int) *ListNode {
	var head, trav *ListNode

	for i := 0; i < len(nodes); i++ {
		nn := &ListNode{
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

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				lists: []*ListNode{
					prepareList([]int{1, 4, 5}),
					prepareList([]int{1, 3, 4}),
					prepareList([]int{2, 6}),
				},
			},
			want: prepareList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "Example 2",
			args: args{
				lists: []*ListNode{
					prepareList([]int{}),
					prepareList([]int{1, 3, 4}),
					prepareList([]int{2, 6}),
					prepareList([]int{-1, 0, 9}),
					prepareList([]int{0, 2, 4, 5, 9}),
					prepareList([]int{-5}),
					prepareList([]int{}),
				},
			},
			want: prepareList([]int{-5, -1, 0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 9, 9}),
		},
		{
			name: "Example 3 Bug",
			args: args{
				lists: []*ListNode{},
			},
			want: prepareList([]int{}),
		},
		{
			name: "Example 4",
			args: args{
				lists: []*ListNode{nil, nil},
			},
			want: prepareList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Log("Got:")
				displayList(t, got)
				t.Log("Want:")
				displayList(t, tt.want)
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
