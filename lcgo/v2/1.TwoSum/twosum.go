package twosum

func twoSum(nums []int, target int) []int {
	calc := map[int]int{}
	for i, n := range nums {
		if j, found := calc[n]; found {
			return []int{j, i}
		}
		calc[target-nums[i]] = i
	}
	return []int{}
}
