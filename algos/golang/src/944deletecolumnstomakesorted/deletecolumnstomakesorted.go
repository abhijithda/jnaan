package deletecolumnstomakesorted

func minDeletionSize(A []string) int {
	if len(A) == 0 {
		return 0
	}
	del := 0
	for c := 0; c < len(A[0]); c++ {
		for w := 0; w < len(A)-1; w++ {
			if A[w][c] > A[w+1][c] {
				del++
				break
			}
		}
	}

	return del
}
