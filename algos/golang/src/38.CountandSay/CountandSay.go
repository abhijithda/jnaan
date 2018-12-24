package countandsay

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}

	say := "1"
	for seq := 1; seq != n; seq++ {
		digits := say[:]
		say = ""
		sameDigit := 1
		for i := 0; i < len(digits)-1; i++ {
			if digits[i] == digits[i+1] {
				sameDigit++
			} else {
				num, _ := strconv.Atoi(string(digits[i]))
				say += fmt.Sprintf("%d%d", sameDigit, num)
				sameDigit = 1
			}
		}
		num, _ := strconv.Atoi(string(digits[len(digits)-1]))
		say += fmt.Sprintf("%d%d", sameDigit, num)
	}
	return say
}
