package ransomnote

import "log"

func canConstruct(ransomNote string, magazine string) bool {
	log.Printf("\nGiven:- ransomNote: %s, magazine: %s\n", ransomNote, magazine)
	noteLetters := map[rune]int{}
	for _, nl := range ransomNote {
		if _, ok := noteLetters[nl]; !ok {
			noteLetters[nl] = 0
		}
		noteLetters[nl]++
	}
	log.Printf("Note Letters: %+v", noteLetters)

	for _, ml := range magazine {
		if _, ok := noteLetters[ml]; ok {
			noteLetters[ml]--
		}
	}

	log.Printf("Note Letters remaining after using magazine: %+v", noteLetters)
	for _, count := range noteLetters {
		if count > 0 {
			return false
		}
	}
	return true
}
