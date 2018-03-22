package main

// Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

// Do not allocate extra space for another array, you must do this in place with constant memory.

// For example,
// Given input array nums = [1,1,2],

// Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

import "fmt"

func removeDuplicates(nums []int) int {
	l := 0
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			l++
			nums[l] = nums[i]
		}
	}
	return l + 1
}

func main() {
	l := 0
	a := []int{1, 1}
	l = removeDuplicates(a)
	fmt.Println(a[:l])

	a = []int{1, 1, 2}
	l = removeDuplicates(a)
	fmt.Println(a[:l])

	a = []int{1, 2, 2, 3, 5, 5}
	l = removeDuplicates(a)
	fmt.Println(a[:l])
}
