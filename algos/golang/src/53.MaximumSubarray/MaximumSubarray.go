package maximumsubarray

import "log"

func maxSubArray(nums []int) int {
	log.Println("Given nums:", nums)
	minVal := -((int(^uint(0) >> 1)) - 1)
	// log.Println("Minimum value:", minVal)
	maxsa := minVal
	sum := 0
	for _, n := range nums {
		if sum+n > n {
			sum += n
		} else {
			sum = n
		}

		log.Println("Sum:", sum)
		if maxsa < sum {
			maxsa = sum
			log.Println("Max subarray so far:", maxsa)
		}
	}
	return maxsa
}

func maxSubArrayMLE(nums []int) int {
	log.Println("Given nums:", nums)
	// minVal := -((int(^uint(0) >> 1)) - 1)
	// log.Println("Minimum value:", minVal)
	sums := make([][]int, len(nums))
	for n := range nums {
		sums[n] = make([]int, len(nums))
		sums[n][n] = nums[n]
	}

	// log.Println("Sums:", sums)
	maxsa := nums[0]
	r := 0
	for c := 1; c < len(nums); c++ {
		// log.Printf("%d (sums[%d][%d]) + %d (nums[%d]) <= %d (nums[%d])?",
		// 	sums[r][c-1], r, c-1, nums[c], c, nums[c], c)
		if sums[r][c-1]+nums[c] <= nums[c] {
			if sums[r][c-1] > nums[c] {
				sums[r][c] = sums[r][c-1]
			} else {
				r = c
				// log.Println("Updated row:", r)
			}
		} else {
			sums[r][c] = sums[r][c-1] + nums[c]
			// log.Printf("Updated %d (sums[%d][%d])", r, c, sums[r][c])
		}
		if maxsa < sums[r][c] {
			maxsa = sums[r][c]
			// log.Println("Max subarray so far:", maxsa)
		}
	}
	return maxsa
}
