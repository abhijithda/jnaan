package sumwithmultiplicity

// Given an integer array A, and an integer target, return the number of tuples i, j, k  such that i < j < k and A[i] + A[j] + A[k] == target.

// As the answer can be very large, return it modulo 10^9 + 7.

import (
	"log"
	"math"
	"sort"

	mu "internal/mathutils"
)

func threeSumMulti(A []int, target int) int {
	// log.Printf("\n\n\nGiven A: %v; target: %d", A, target)
	count := 0

	cntD := map[int]int{}
	uniqueD := []int{}
	for _, n := range A {
		cntD[n]++
		if cntD[n] == 1 {
			uniqueD = append(uniqueD, n)
		}
	}
	// log.Println("cntD:", cntD)

	for d, c := range cntD {
		if c >= 3 {
			if 3*d == target {
				count += mu.NCR(c, 3)
				log.Printf("d1:%d; d2:%d; d3:%d; Count:%d", d, d, d, count)
			}
		}
		if c >= 2 {
			d3 := target - 2*d
			if d3 != d && cntD[d3] > 0 {
				count += mu.NCR(c, 2) * mu.NCR(cntD[d3], 1)
				log.Printf("d1:%d; d2:%d; d3:%d; Count:%d", d, d, d3, count)
			}
		}
	}

	sort.Ints(uniqueD)
	for i := 0; i < len(uniqueD)-1; i++ {
		d1 := uniqueD[i]
		for j := i + 1; j < len(uniqueD); j++ {
			d2 := uniqueD[j]
			d3 := target - d1 - d2
			if d3 > d2 && cntD[d3] > 0 {
				// if d1 != d && d != d3 && cntD[d3] > 0 {
				count += mu.NCR(cntD[d1], 1) * mu.NCR(cntD[d2], 1) * mu.NCR(cntD[d3], 1)
				// log.Printf("d1:%d; d2:%d; d3:%d; Count:%d", d1, d2, d3, count)
			}
		}
	}
	// log.Println("Count:", count)
	return count % (int(math.Pow(10, 9)) + 7)
}
