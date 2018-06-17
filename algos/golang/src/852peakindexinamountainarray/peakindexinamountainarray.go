package peakindexinamountainarray

// 852. Peak Index in a Mountain Array
// User Accepted: 652
// User Tried: 707
// Total Accepted: 653
// Total Submissions: 767
// Difficulty: Easy
// Let's call an array A a mountain if the following properties hold:

// A.length >= 3
// There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
// Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

// Example 1:

// Input: [0,1,0]
// Output: 1
// Example 2:

// Input: [0,2,1,0]
// Output: 1
// Note:

// 3 <= A.length <= 10000
// 0 <= A[i] <= 10^6
// A is a mountain, as defined above.

func peakIndexInMountainArray(A []int) int {
	maxPeakIndex := -1
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			if maxPeakIndex == -1 || A[maxPeakIndex] < A[i] {
				maxPeakIndex = i
			}
		}
	}
	return maxPeakIndex
}
