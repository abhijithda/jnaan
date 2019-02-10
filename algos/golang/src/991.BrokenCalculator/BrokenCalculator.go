package brokencalculator

import "log"

func brokenCalc(X int, Y int) int {
	log.Printf("Given X: %d; Y: %d", X, Y)
	opCnt := 0
	for Y > X {
		opCnt++
		if Y%2 == 1 {
			// Note: 11/2 should be set to 6, and
			// opCnt should be incremented.
			Y = (Y + 1) / 2
			opCnt++
		} else {
			Y /= 2
		}
	}
	log.Printf("Y: %d; opCnt: %d", Y, opCnt)
	opCnt += X - Y
	return opCnt
}
