package faircandyswap

func fairCandySwap(A []int, B []int) []int {
	exch := []int{0, 0}

	sumA := 0
	sumB := 0
	for i := range A {
		sumA += A[i]
	}
	for i := range B {
		sumB += B[i]
	}

	diff := sumA - sumB
	swapDiff := diff / 2

	for a := range A {
		for b := range B {
			if A[a]-B[b] == swapDiff {
				exch[0] = A[a]
				exch[1] = B[b]
			}
		}
	}
	return exch
}
