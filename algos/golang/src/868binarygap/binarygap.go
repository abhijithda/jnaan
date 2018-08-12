package binarygap

import (
	"log"
	"strconv"
)

func binaryGap(N int) int {
	maxDist := 0
	dist := -1
	binN := strconv.FormatInt(int64(N), 2)
	log.Printf("%v = %v\n", N, binN)
	for i := 0; i < len(binN); i++ {
		if binN[i] == '1' {
			log.Printf("%v(1) at %v", binN[i], i)
			// if dist == -1 {
			// 	// Found first '1'
			// 	dist = 0
			// } else {
			dist++
			if maxDist < dist {
				maxDist = dist
			}
			// Reset dist, as this is the new first '1'
			dist = 0
			// }
		} else {
			log.Printf("%v at %v", binN, i)
			// If '1' was found earlier, then increment dist.
			if dist != -1 {
				dist++
			}
		}
	}
	log.Printf("Max Distance: %v", maxDist)
	return maxDist
}
