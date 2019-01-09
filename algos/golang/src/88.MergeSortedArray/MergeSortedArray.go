package mergesortedarray

import "log"

func merge(nums1 []int, m int, nums2 []int, n int) {
	log.Printf("Given nums1: %v, m: %d; nums2: %v; n: %d",
		nums1, m, nums2, n)
	i := m + n - 1
	m--
	n--
	for ; i >= 0; i-- {
		log.Printf("Index: %d; m: %d; n: %d", i, m, n)
		if n == -1 || m == -1 {
			break
		}
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		log.Println("nums1:", nums1)
	}

	if m >= 0 {
		for ; i >= 0; i-- {
			nums1[i] = nums1[m]
			m--
		}
	}

	if n >= 0 {
		for ; i >= 0; i-- {
			nums1[i] = nums2[n]
			n--
		}
	}
}
