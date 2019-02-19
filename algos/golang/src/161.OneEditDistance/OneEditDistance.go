package oneeditdistance

import (
	"log"
	"math"
)

func isOneEditDistance(s, t string) bool {
	log.Printf("Given s: %s; t: %s", s, t)
	if math.Abs(float64(len(s))-float64(len(t))) > 1 {
		return false
	}
	isOne := false
	si := -1
	ti := -1
	for si < len(s)-1 && ti < len(t)-1 {
		if s[si+1] == t[ti+1] {
			si++
			ti++
		} else {
			if len(s) != len(t) {
				if ti+2 < len(t) && s[si+1] == t[ti+2] {
					si++
					ti += 2
					if isOne {
						return false
					}
					isOne = true
				} else if si+2 < len(s) && s[si+2] == t[ti+1] {
					si += 2
					ti++
					if isOne {
						return false
					}
					isOne = true
				} else {
					return false
				}
			} else {
				if si+2 < len(s) && ti+2 < len(t) && s[si+2] == t[ti+2] {
					si += 2
					ti += 2
					if isOne {
						return false
					}
					isOne = true
				} else {
					return false
				}
			}
		}
	}

	if si == len(s)-1 && ti == len(t)-1 {
		return isOne
	}
	// If len is not same, then it means last element differs
	return true
}
