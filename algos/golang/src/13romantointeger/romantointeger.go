package romantointeger

import (
	"fmt"
	"log"
)

func romanToInt(s string) int {
	romanValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	subtractRule := map[string]string{
		"V": "I",
		"X": "I",
		"L": "X",
		"C": "X",
		"D": "C",
		"M": "C",
	}

	num := 0
	// sc := s[:]
	// fmt.Printf("%+v", s)
	for c := 0; c < len(s); c++ {
		fmt.Printf("%v -> %v\n", s[c], romanValues[string(s[c])])
		if c+1 < len(s) && subtractRule[string(s[c+1])] == string(s[c]) {
			log.Printf("Subtracting as %v before %v\n", string(s[c]), string(s[c+1]))
			num = num + (romanValues[string(s[c+1])] - romanValues[string(s[c])])
			c++
		} else {
			num += romanValues[string(s[c])]
		}
	}
	return num
}
