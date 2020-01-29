package main

import (
	"log"
	"math"
)

func reverse(x int) int {
	log.Println("Given:", x)
	rev := 0
	maxint := math.MaxInt32
	log.Println("Max int:", maxint)

	positive := true
	if x < 0 {
		positive = false
		x = int(math.Abs(float64(x)))
	}

	for x >= 1 {
		rem := x % 10
		x = x / 10
		rev = (rev * 10) + rem
	}

	if positive == false {
		if rev-1 > int(maxint) {
			return 0
		}
		rev *= -1
	} else if rev > int(maxint) {
		return 0
	}

	return rev
}
