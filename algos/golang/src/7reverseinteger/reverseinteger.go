package reverseinteger

import (
	"log"
	"strconv"
)

func reverse(x int) int {
	sx := strconv.Itoa(x)
	asx := sx[:]
	ispositive := true
	if asx[0] == '-' {
		ispositive = false
		asx = asx[1:]
	}

	l := len(asx)
	var rsx = make([]byte, l)
	for i := range asx {
		log.Println("Index:", i, ":", string(asx[i]))
		rsx[l-i-1] = asx[i]
		log.Printf("r[%d] = %s", l-i-1, string(asx[i]))
	}

	maxint := int(^uint32(0) >> 1)
	log.Println("Max int:", maxint)
	minint := (-1*int(^uint32(0)>>1) - 1)
	log.Println("Min int:", minint)
	rx, _ := strconv.Atoi(string(rsx))
	if !ispositive {
		rx *= -1
	}
	if rx < minint || rx > maxint {
		return 0
	}

	return rx
}
