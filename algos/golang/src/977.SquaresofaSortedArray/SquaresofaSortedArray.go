package squaresofasortedarray

import (
	"log"
	"math"
)

func sortedSquares(A []int) []int {
	midI := 0
	for i, a := range A {
		if a >= 0 {
			midI = i
			break
		}
	}
	log.Println("MidI:", midI)

	sortedSq := []int{}
	p := midI
	n := midI - 1
	for cnt := 0; cnt < len(A); cnt++ {
		if p >= len(A) || n < 0 {
			break
		}
		log.Printf("p: %d; n: %d; cnt: %d", p, n, cnt)
		if A[p] <= int(math.Abs(float64(A[n]))) && p < len(A) {
			sortedSq = append(sortedSq, A[p]*A[p])
			p++
		} else {
			sortedSq = append(sortedSq, A[n]*A[n])
			n--
		}
	}
	for ; p < len(A); p++ {
		sortedSq = append(sortedSq, A[p]*A[p])
	}
	for ; n >= 0; n-- {
		sortedSq = append(sortedSq, A[n]*A[n])
	}

	return sortedSq
}
