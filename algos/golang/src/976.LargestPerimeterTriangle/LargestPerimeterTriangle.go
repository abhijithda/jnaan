package largestperimetertriangle

import (
	"log"
	"sort"
)

func largestPerimeter(A []int) int {
	log.Println("A:", A)
	sort.Ints(A)
	log.Println("Sorted A:", A)
	for i := len(A) - 1; i >= 2; i-- {
		log.Printf("A[i]: %d; A[i-1]: %d; A[i-2]: %d", A[i], A[i-1], A[i-2])
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}
