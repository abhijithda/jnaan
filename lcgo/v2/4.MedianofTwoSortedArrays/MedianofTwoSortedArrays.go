package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	median := 0.0

	maxLen := len(nums1)
	minLen := len(nums1)
	if minLen > len(nums2) {
		minLen = len(nums2)
	} else {
		maxLen = len(nums2)
	}
	i := 0
	for ; i < minLen; i++ {
		median += float64((nums1[i] + nums2[i]) / 2)
	}
	for ; i < len(nums1); i++ {
		median += float64(nums1[i])
	}
	for ; i < len(nums2); i++ {
		median += float64(nums2[i])
	}
	median /= float64(maxLen)
	return median
}
