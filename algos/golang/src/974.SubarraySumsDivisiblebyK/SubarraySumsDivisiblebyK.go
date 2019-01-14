package subarraysumsdivisiblebyk

import "log"

// Solution: https://leetcode.com/problems/subarray-sums-divisible-by-k/discuss/217985/JavaC%2B%2B-Count-the-Remainder
func subarraysDivByK(A []int, K int) int {
	count := map[int]int{0: 1}
	prefix := 0
	res := 0
	for _, a := range A {
		prefix = (prefix + a%K + K) % K
		res += count[prefix]
		count[prefix]++
		log.Println("Result:", res)
		log.Printf("count[%d] = %d", prefix, count[prefix])
	}
	return res
}

func subarraysDivByKTake1(A []int, K int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		sum := 0
		for j := i; j < len(A); j++ {
			sum += A[j]
			if sum%K == 0 {
				count++
			}
		}
	}
	return count
}
