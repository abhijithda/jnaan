package arraypartition

import (
	"log"
	"sort"
)

func arrayPairSum(nums []int) int {
	log.Println("Given:", nums)
	sort.Ints(nums)

	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
