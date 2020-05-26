package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	// if len(s) <= 1 {
	// 	return len(s)
	// }
	maxLen := 0
	insubstr := map[byte]bool{}
	b, e := 0, 0
	for ; e < len(s); e++ {
		if _, found := insubstr[s[e]]; found {
			if maxLen < e-b {
				maxLen = e - b
			}
			for found {
				delete(insubstr, s[b])
				b++
				_, found = insubstr[s[e]]
			}
		}
		insubstr[s[e]] = true
	}
	if maxLen < e-b {
		maxLen = e - b
	}
	return maxLen
}
