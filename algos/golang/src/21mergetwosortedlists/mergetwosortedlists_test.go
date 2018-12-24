package mergetwosortedlists

import (
	"reflect"
	"testing"
)

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

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Empty lists",
			args: args{
				l1: prepareList([]int{}),
				l2: prepareList([]int{}),
			},
			want: nil,
		},
		{
			name: "First single node list and second empty lists",
			args: args{
				l1: prepareList([]int{1}),
				l2: prepareList([]int{}),
			},
			want: prepareList([]int{1}),
		},
		{
			name: "First empty list and second single node lists",
			args: args{
				l1: nil,
				l2: &ListNode{},
			},
			want: &ListNode{},
		},
		{
			name: "First list has smaller element",
			args: args{
				l1: prepareList([]int{1}),
				l2: prepareList([]int{2}),
			},
			want: prepareList([]int{1, 2}),
		},
		{
			name: "Second list has smaller element",
			args: args{
				l1: prepareList([]int{9}),
				l2: prepareList([]int{5}),
			},
			want: prepareList([]int{5, 9}),
		},
		{
			name: "Odd & even lists",
			args: args{
				l1: prepareList([]int{1, 3, 5, 7, 9}),
				l2: prepareList([]int{0, 2, 4, 6, 8}),
			},
			want: prepareList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
