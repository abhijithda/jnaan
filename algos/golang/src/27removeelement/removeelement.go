package removeelement

import "log"

func removeElement(nums []int, val int) int {
	log.Printf("Given nums: %v, val: %d", nums, val)
	s := 0
	e := len(nums) - 1
	for s <= e {
		for nums[e] == val {
			e--
			if s > e {
				return s
			}
		}
		log.Println("E:", e)
		if nums[s] == val {
			nums[s] = nums[e]
			nums[e] = val
			log.Println("nums after swap:", nums)
		}
		s++
		log.Println("S:", s)
	}
	log.Println("nums:", nums)
	log.Println("Length:", s)
	return s
}
