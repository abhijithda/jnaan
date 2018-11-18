package validmountainarray

func validMountainArray(A []int) bool {
	var incr, tip bool
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return false
		} else if A[i] < A[i+1] {
			incr = true
			if tip == true {
				return false
			}
		} else if A[i] > A[i+1] {
			if incr == false {
				return false
			}
			tip = true
		}
	}
	return tip
}
