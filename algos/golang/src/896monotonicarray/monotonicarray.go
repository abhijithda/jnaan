package monotonicarray

func isMonotonic(A []int) bool {
	isMonotonic := true
	incrOrDecr := 0

	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			if incrOrDecr == 1 {
				isMonotonic = false
				break
			}
			incrOrDecr = -1
		} else if A[i-1] > A[i] {
			if incrOrDecr == -1 {
				isMonotonic = false
				break
			}
			incrOrDecr = 1
		}
	}

	return isMonotonic
}
