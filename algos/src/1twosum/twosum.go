package twosum

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// func main() {
// 	list1 := []int{1, 2, 3}
// 	target := 3
// 	fmt.Println(twoSum(list1, target))
// }
