package findandreplacepattern

import (
	"fmt"
	"log"
)

func findAndReplacePattern(words []string, pattern string) []string {
	matchingWords := []string{}
	patternchars := map[byte]int{}
	intPattern := []int{}
	intWordPattern := []int{}
	uniqueCharsCount := 0
	// var mapchars := map[string]int{}

	for c := range pattern {
		if val, ok := patternchars[pattern[c]]; ok {
			intPattern = append(intPattern, val)
		} else {
			intPattern = append(intPattern, uniqueCharsCount)
			patternchars[pattern[c]] = uniqueCharsCount
			uniqueCharsCount++
		}
	}
	log.Printf("Int pattern: %+v\n", intPattern)

	for w := range words {
		fmt.Println(words[w])
		uniqueCharsCount = 0
		index := 0
		match := true
		patternchars = map[byte]int{}
		log.Printf("Processing word: %s\n", words[w])
		word := words[w]
		intWordPattern = []int{}
		for c := range word {
			log.Printf("Processing character: %v\n", string(word[c]))
			if val, ok := patternchars[word[c]]; ok {
				log.Printf("%v is present in patternchars with value: %v\n", string(pattern[c]), val)
				intWordPattern = append(intWordPattern, val)
				if intWordPattern[index] != intPattern[index] {
					match = false
					continue
				}
				index++
			} else {
				log.Printf("Processing new character: %v\n", string(word[c]))
				intWordPattern = append(intWordPattern, uniqueCharsCount)
				patternchars[word[c]] = uniqueCharsCount
				uniqueCharsCount++
				if intWordPattern[index] != intPattern[index] {
					match = false
					continue
				}
				index++
			}
		}
		log.Printf("Word Int pattern: %+v\n", intWordPattern)

		if match {
			matchingWords = append(matchingWords, words[w])
		}
	}

	return matchingWords
}
