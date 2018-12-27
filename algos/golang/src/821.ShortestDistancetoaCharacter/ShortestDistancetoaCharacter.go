// 821. Shortest Distance to a Character

// Given a string S and a character C, return an array of integers representing the shortest distance from the character C in the string.

// Example 1:

// Input: S = "loveleetcode", C = 'e'
// Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]

// Note:

// S string length is in [1, 10000].
// C is a single character, and guaranteed to be in string S.
// All letters in S and C are lowercase.

package shortesttocharacter

import (
	"fmt"
	"log"
	"os"
)

const logFile = "log.txt"

func init() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()

	log.SetOutput(f)
}

func shortestToChar(S string, C byte) []int {
	D := make([]int, len(S))

	// If only one char, it has to be C, so set d to len-1.
	d := len(S) - 1
	for i, c := range S {
		if byte(c) == C {
			d = 0
		}
		D[i] = d
		d++
	}
	// log.Println("Pass 1:", D)

	d = len(S) - 1
	for i := len(S) - 1; i >= 0; i-- {
		if d < D[i] {
			D[i] = d
		} else {
			d = D[i]
		}
		d++
	}

	// log.Println("Pass 2:", D)

	return D
}

func shortestToCharOld(S string, C byte) []int {
	var indexes []int
	for i, p := 0, len(S); i < len(S); i++ {
		log.Printf("Char at %v is %v\n", i, S[i:i+1])
		indexes[i] = i - p
		if i-p < 0 {
			indexes[i] = p - i
		}
		if C == S[i] {
			for t := p; t <= i; t++ {
				if indexes[i] > i-t {
					indexes[t] = i - t
				}
			}
			p = i
		}
	}

	for i := range indexes {
		log.Printf("Indexes: %v", i)
	}
	return indexes
}
