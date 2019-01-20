package longestturbulentsubarray

func maxTurbulenceSize(A []int) int {
	if len(A) == 0 {
		return 0
	}

	maxSize := 1
	size := 1
	for k := 0; k < len(A)-1; k++ {
		if k%2 == 0 && A[k] < A[k+1] {
			// log.Printf("Even: %d(A[%d]) < %d(A[%d])", A[k], k, A[k+1], k+1)
			size++
		} else if k%2 != 0 && A[k] > A[k+1] {
			// log.Printf("Odd: %d(A[%d]) > %d(A[%d])", A[k], k, A[k+1], k+1)
			size++
		} else {
			if maxSize < size {
				maxSize = size
			}
			size = 1
			// log.Println("Resetting size...")
		}
	}
	if maxSize < size {
		maxSize = size
	}

	size = 1
	for k := 0; k < len(A)-1; k++ {
		// log.Printf("%d(A[%d]) <=> %d(A[%d])", A[k], k, A[k+1], k+1)
		if k%2 == 0 && A[k] > A[k+1] {
			// log.Printf("Even: %d(A[%d]) > %d(A[%d])", A[k], k, A[k+1], k+1)
			size++
		} else if k%2 != 0 && A[k] < A[k+1] {
			// log.Printf("Odd: %d(A[%d]) < %d(A[%d])", A[k], k, A[k+1], k+1)
			size++
		} else {
			// log.Printf("Resetting size from %d to 1...", size)
			if maxSize < size {
				maxSize = size
			}
			size = 1
		}
	}
	if maxSize < size {
		maxSize = size
	}
	return maxSize
}
