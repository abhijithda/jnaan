package subarraysumsdivisiblebyk

func subarraysDivByK(A []int, K int) int {
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
