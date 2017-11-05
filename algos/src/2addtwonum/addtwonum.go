package main

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	ln := res
	t1 := l1
	t2 := l2
	carry := 0
	for t1 != nil && t2 != nil {
		fmt.Printf("l1.Val: %v, l2.Val: %v\n", t1.Val, t2.Val)
		tot := t1.Val + t2.Val + carry
		carry = int((tot) / 10)
		if carry == 1 {
			tot = tot - 10
		}
		newnode := &ListNode{
			Val:  tot,
			Next: nil,
		}
		if res == nil {
			res = newnode
			ln = res
		} else {
			ln.Next = newnode
			ln = ln.Next
		}
		t1 = t1.Next
		t2 = t2.Next
	}
	exT := t1
	// If t2 is not nil, then t1 would be nil as above for loop exited.
	if t2 != nil {
		exT = t2
	}
	for exT != nil {
		tot := exT.Val + carry
		carry = int((tot) / 10)
		if carry == 1 {
			tot = tot - 10
		}
		newnode := &ListNode{
			Val:  tot,
			Next: nil,
		}
		if res == nil {
			res = newnode
			ln = res
		} else {
			ln.Next = newnode
			ln = ln.Next
		}
		exT = exT.Next
	}

	if carry == 1 {
		newnode := &ListNode{
			Val:  carry,
			Next: nil,
		}
		ln.Next = newnode
	}
	return res
}

func test1() {
	l1n1 := ListNode{
		Val:  3,
		Next: nil,
	}
	l1n2 := ListNode{
		Val:  4,
		Next: &l1n1,
	}
	l1n3 := ListNode{
		Val:  2,
		Next: &l1n2,
	}

	l1 := &l1n3

	l2n1 := ListNode{
		Val:  4,
		Next: nil,
	}
	l2n2 := ListNode{
		Val:  6,
		Next: &l2n1,
	}
	l2n3 := ListNode{
		Val:  5,
		Next: &l2n2,
	}

	l2 := &l2n3

	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}

func test2() {
	l1n1 := ListNode{
		Val:  8,
		Next: nil,
	}
	l1n2 := ListNode{
		Val:  1,
		Next: &l1n1,
	}
	l1 := &l1n2

	l2n1 := ListNode{
		Val:  0,
		Next: nil,
	}
	l2 := &l2n1

	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

func main() {
	test1()
	test2()
}
