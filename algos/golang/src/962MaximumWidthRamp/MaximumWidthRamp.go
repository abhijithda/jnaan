package maximumwidthramp

import (
	"fmt"
	"log"
	"os"
)

const logFile = "log.txt"

func init() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()
	log.SetOutput(f)
}

func maxWidthRamp(A []int) int {
	ramp := 0
	for i := 0; i < len(A); i++ {
		for j := len(A) - 1; j > i; j-- {
			if j-i < ramp {
				break
			}
			if A[i] <= A[j] && j-i > ramp {
				ramp = j - i
			}
		}
	}
	return ramp
}

func maxWidthRampTLE(A []int) int {
	log.Println("\n\n\nGiven A:", A)
	values := map[int][]int{}
	for i, a := range A {
		if _, ok := values[a]; ok {
			values[a] = append(values[a], i)
		} else {
			values[a] = []int{i}
		}

		for k, v := range values {
			log.Printf("Checking whether i(%d) < v[0](%d) && k(%d) <= a(%d)",
				i, v[0], k, a)
			// k contains A's values; and v contains list of indicies
			if v[0] < i && k <= a {
				values[k] = append(values[k], i)
				log.Printf("New list for %d is %v", k, values[k])
			}
		}
	}

	ramp := 0
	for k, v := range values {
		log.Printf("Values of %d: %v", k, values[k])
		if v[len(v)-1]-v[0] > ramp {
			ramp = v[len(v)-1] - v[0]
		}
	}
	return ramp
}
