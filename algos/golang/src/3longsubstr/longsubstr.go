package longsubstr

import "log"

func lengthOfLongestSubstring(s string) int {
	log.Println("Given s:", s)
	maxlen := 0
	for i := 0; i < len(s); i++ {
		present := map[string]bool{string(s[i]): true}
		j := i + 1
		for ; j < len(s); j++ {
			if _, ok := present[string(s[j])]; ok {
				log.Printf("%s repeated at %d & %d", string(s[i]), i, j)
				break
			} else {
				present[string(s[j])] = true
			}
		}
		if maxlen < j-i {
			maxlen = j - i
			log.Println("Longest string:", s[i:j], i, j)
		}
	}
	return maxlen
}
