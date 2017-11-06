package longsubstr

func lengthOfLongestSubstring(s string) int {
	l := 0
	ll := 0
	sp := &s

	var charExists = map[string]int{}
	for i := 0; i < len(s); i++ {
		c := (*sp)[i : i+1]
		// log.Printf("Char at %v is %v\n", i, c)
		if _, ok := charExists[c]; ok {
			if ll < l {
				ll = l
			}
			for k := range charExists {
				delete(charExists, k)
			}
			l = 0
		}
		charExists[c] = 1
		l++
	}
	if ll < l {
		ll = l
	}

	return ll
}
