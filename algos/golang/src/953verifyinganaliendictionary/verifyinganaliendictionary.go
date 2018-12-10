package verifyinganaliendictionary

import "log"

func isAlienSorted(words []string, order string) bool {
	log.Printf("Given words: %+v; order: %s", words, order)

	ordervalue := map[string]int{" ": 0}
	for i := range order {
		ordervalue[string(order[i])] = i + 1
	}
	log.Printf("Ordered values: %+v", ordervalue)

	wordVal := map[string]int{}

	maxworldlen := 0
	for w := range words {
		if maxworldlen < len(words[w]) {
			maxworldlen = len(words[w])
		}
	}
	for w := range words {
		word := words[w]
		n := 0
		for ; n < len(word); n++ {
			wordVal[word] = wordVal[word]*10 + ordervalue[string(word[n])]
		}
		for ; n < maxworldlen; n++ {
			wordVal[word] *= 10
		}
	}
	log.Println("Word values: ", wordVal)

	prev := -1
	for w := range words {
		if prev <= wordVal[words[w]] {
			prev = wordVal[words[w]]
		} else {
			return false
		}
	}

	return true
}
