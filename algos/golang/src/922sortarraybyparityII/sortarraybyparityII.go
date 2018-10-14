package sortarraybyparityii

import (
	"log"
)

func sortArrayByParityII(A []int) []int {
	sA := make([]int, len(A))

	log.Println("Input A:", A)
	defer log.Println("Return A:", sA)
	ei := 0
	oi := 1
	for i := 0; i < len(A); i++ {
		log.Printf("A[%d]: %d\n", i, A[i])
		if 0 == A[i]%2 {
			sA[ei] = A[i]
			ei += 2
		} else {
			sA[oi] = A[i]
			oi += 2
		}
	}
	return sA
}
