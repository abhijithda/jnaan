package findandreplaceinstring

import (
	"log"
	"os"
	"sort"
)

const logFile = "log.txt"

func init() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()
	log.SetOutput(f)
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	log.Printf("Given string: %v", S)
	startInd := 0
	var resS string

	// TODO: Sort indexes, and based on the new values update the source and the targets.
	mappedIndexes := make(map[int]int)
	sortedIndexes := make([]int, len(indexes))
	copy(sortedIndexes, indexes)
	sort.Ints(sortedIndexes)

	for i := range indexes {
		mappedIndexes[indexes[i]] = i
	}

	for k, v := range mappedIndexes {
		log.Printf("MI: %v -> %v", k, v)
		// log.Printf("MI: %v -> %v", mi, mappedIndexes[mi])
	}
	for si := range sortedIndexes {
		i := mappedIndexes[sortedIndexes[si]]
		log.Printf("Index (si=%v): %v", si, i)
		resS += S[startInd:indexes[i]]
		startInd = indexes[i]
		if S[indexes[i]:indexes[i]+len(sources[i])] == sources[i] {
			log.Printf("String %v matches with source %v at index: %v.\n Replacing with %v",
				S[indexes[i]:indexes[i]+len(sources[i])], sources[i], indexes[i], targets[i])
			resS += targets[i]
		} else {
			log.Printf("String %v doesn't match with source %v at index: %v.\n Replacing with %v",
				S[indexes[i]:indexes[i]+len(sources[i])], sources[i], indexes[i], S[indexes[i]:indexes[i]+len(sources[i])])
			resS += S[indexes[i] : indexes[i]+len(sources[i])]
		}
		startInd += len(sources[i])
		log.Printf("Resulting String: %v", resS)
	}
	if startInd < len(S) {
		resS += S[startInd:]
	}
	log.Printf("Final resulting string: %v", resS)
	return resS
}
