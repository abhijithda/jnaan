package UncommonWordsfromTwoSentences

import (
	"log"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	uncommonWords := []string{}
	log.Println("A word:", A)
	log.Println("B word:", B)

	wordCountA := map[string]int{}
	words := strings.Split(A, " ")
	log.Printf("words: %+v\n", words)
	for w := range words {
		wstr := string(words[w])
		log.Printf("Word[%d]: %s", w, wstr)
		if val, ok := wordCountA[wstr]; ok {
			wordCountA[wstr] = val + 1
		} else {
			wordCountA[wstr] = 1
		}
	}
	log.Printf("A word count: %+v\n", wordCountA)

	wordCountB := map[string]int{}
	words = strings.Split(B, " ")
	for w := range words {
		wstr := string(words[w])
		if val, ok := wordCountB[wstr]; ok {
			wordCountB[wstr] = val + 1
		} else {
			wordCountB[wstr] = 1
		}
	}
	log.Printf("B word count: %+v\n", wordCountB)

	for word, cnt := range wordCountA {
		if cnt == 1 {
			if _, ok := wordCountB[word]; !ok {
				uncommonWords = append(uncommonWords, word)
			}
		}
	}
	for word, cnt := range wordCountB {
		if cnt == 1 {
			if _, ok := wordCountA[word]; !ok {
				uncommonWords = append(uncommonWords, word)
			}
		}
	}

	log.Printf("Uncommon words: %+v\n", uncommonWords)
	return uncommonWords

}
