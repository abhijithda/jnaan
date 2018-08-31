package longestcommonprefix

import "log"

func longestCommonPrefix(strs []string) string {
	prefix := ""

	maxcmplen := int(^uint(0) >> 1)
	log.Printf("Compare length: %d\n", maxcmplen)

	indexUpdated := false
	for s := range strs {
		if maxcmplen > len(strs[s]) {
			maxcmplen = len(strs[s])
			indexUpdated = true
		}
	}
	log.Printf("Compare length: %d\n", maxcmplen)

	if !indexUpdated {
		return prefix
	}

	for i := 0; i < maxcmplen; i++ {
		tStr := strs[0]
		curchar := tStr[i]
		match := true
		log.Printf("Current char to compare: %v", string(curchar))
		for s := range strs {
			str := strs[s]
			log.Printf("comparing cur char %v with %v\n", string(curchar), string(str[i]))
			if str[i] != curchar {
				match = false
				break
			}
		}
		if match {
			prefix += string(tStr[i])
		} else {
			break
		}
	}
	return prefix
}
