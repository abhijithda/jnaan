package implementstrstr

import "log"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	n := 0
	for i := 0; i < len(haystack); i++ {
		log.Printf("Checking H[%d](%s) == N[%d](%s)", i, string(haystack[i]), n, string(needle[n]))
		if haystack[i] == needle[n] {
			log.Printf("Matched character %v at indexes H(%d) & N(%d)", string(haystack[i]), i, n)
			if len(needle)-1 == n {
				log.Println("Returning position:", i-n)
				return i - n
			}
			n++
		} else {
			// Reset index so as to check from next character...
			log.Println("Resetting indexes...")
			i = i - n
			n = 0
		}
	}
	return -1
}
