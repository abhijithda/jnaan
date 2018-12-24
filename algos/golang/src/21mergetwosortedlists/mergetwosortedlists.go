package mergetwosortedlists

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	res = nil

	t1 := l1
	t2 := l2
	if t1 != nil && t2 != nil {
		if t1.Val < t2.Val {
			res = t1
			t1 = t1.Next
		} else {
			res = t2
			t2 = t2.Next
		}
	} else if t1 != nil {
		res = t1
		t1 = t1.Next
	} else if t2 != nil {
		res = t2
		t2 = t2.Next
	}
	head := res
	for t1 != nil && t2 != nil {
		if t1.Val < t2.Val {
			res.Next = t1
			t1 = t1.Next
			res = res.Next
		} else {
			res.Next = t2
			t2 = t2.Next
			res = res.Next
		}
	}
	for t1 != nil {
		res.Next = t1
		t1 = t1.Next
		res = res.Next
	}
	for t2 != nil {
		res.Next = t2
		t2 = t2.Next
		res = res.Next
	}

	return head
}
