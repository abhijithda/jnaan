package greatestcommondivisorofstrings

import (
	"fmt"
	"log"
	"os"
)

// ABAABAABA
// ABA

func init() {
	const logFile = "log.txt"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()
	log.SetOutput(f)
}

func isDivisible(str, gcdstr string) bool {
	log.Printf("isDivisible(%s, %s)", str, gcdstr)
	gcdLen := len(gcdstr)
	if len(str)%gcdLen != 0 {
		return false
	}
	for i := 0; i < len(str); i++ {
		if str[i] != gcdstr[i%gcdLen] {
			return false
		}
	}
	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	log.Printf("Given str1: %s, str2: %s", str1, str2)
	defer log.Println()
	// Initialize gcd string to the shorter string
	lstr := str1
	sstr := str2
	l1 := len(lstr)
	l2 := len(sstr)
	if l1 < l2 {
		lstr = str2
		sstr = str1
	}

	for d := 1; d < len(sstr); d++ {
		rem := len(sstr) % d
		if rem != 0 {
			log.Printf("d: %d, sstr Len: %d", d, len(sstr))
			continue
		}

		newgl := len(sstr) / d
		log.Println("New GCD len:", newgl)
		if newgl == 0 {
			return ""
		}

		gcdstr := sstr[:newgl]
		if isDivisible(sstr, gcdstr) && isDivisible(lstr, gcdstr) {
			log.Printf("GCD str: %s", gcdstr)
			return gcdstr
		}
	}

	return ""
}
