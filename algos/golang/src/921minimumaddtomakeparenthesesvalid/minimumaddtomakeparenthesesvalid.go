package minimumaddtomakeparenthesesvalid

import (
	"log"
	"math"
)

func minAddToMakeValid(S string) int {
	newneeded := 0
	count := 0

	for i := 0; i < len(S); i++ {
		log.Printf("S[%d]: %s", i, string(S[i]))
		if string(S[i]) == string('(') {
			count++
		} else if string(S[i]) == string(')') {
			count--
			if count < 0 {
				newneeded++
				count = 0
			}
		}
	}
	return int(math.Abs(float64(count))) + newneeded
}
