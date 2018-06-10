package shiftingletters

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

var allCh []string

func shiftLetter(c string, shiftCount int) string {
	// allCh := make(map[string]int)

	log.Println("Shifting letter: ", c)
	log.Println("Len of allCh: ", len(allCh))
	loc := 0
	for i := 0; i < len(allCh); i++ {
		if c == allCh[i] {
			loc = i
			break
		}
	}

	loc += (shiftCount % len(allCh))
	loc %= len(allCh)

	log.Printf("Shiftletter : %v -%v-> %v", c, shiftCount, allCh[loc])
	return allCh[loc]
}

func shiftingLetters(S string, shifts []int) string {
	log.Println()
	log.Println("New run...!!!!!!!!!!")
	log.Println()
	allCh = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var res string

	totalShifts := 0
	for i := 0; i < len(shifts); i++ {
		effShift := (shifts[i] % len(allCh))
		totalShifts += effShift
	}

	for i := 0; i < len(shifts); i++ {
		log.Printf("Total Shifts %v of %v: ", totalShifts, string(S[i]))
		if i == 0 {
			res += shiftLetter(string(S[i]), totalShifts) + S[i+1:]
		} else if i == len(S) {
			res = res[0:i] + shiftLetter(string(S[i]), totalShifts)
		} else {
			res = res[0:i] + shiftLetter(string(S[i]), totalShifts) + S[i+1:]
		}
		totalShifts -= (shifts[i] % len(allCh))
		log.Println("Result: ", res)
	}
	return res
}
