package bitwiseorsofsubarrays

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

func subarrayBitwiseORs(A []int) int {
	count := 0
	mbitwiseOr := map[int]int{}

	for i := 0; i < len(A); i++ {
		if i+1 < len(A) && A[i+1] == A[i] {
			continue
		}
		bitwiseOr := A[i]
		mbitwiseOr[bitwiseOr] = bitwiseOr
		for j := i + 1; j < len(A); j++ {
			bitwiseOr |= A[j]
			mbitwiseOr[bitwiseOr] = bitwiseOr
		}
	}

	count = len(mbitwiseOr)
	log.Printf("Count: %d based on length of %v\n", count, mbitwiseOr)
	return count
}

// TLEsubarrayBitwiseORs is time limit exceeded !
func TLEsubarrayBitwiseORs(A []int) int {
	count := 0
	mbitwiseOr := map[int]int{}

	for ws := 1; ws <= len(A); ws++ {
		// log.Println("Window size:", ws)
		for i := 0; i < len(A); i++ {
			// log.Println("Index:", i)
			if i+ws <= len(A) {
				curSubArray := A[i : i+ws]
				bitwiseOr := 0
				for s := 0; s < len(curSubArray); s++ {
					bitwiseOr |= curSubArray[s]
				}
				mbitwiseOr[bitwiseOr] = bitwiseOr
				log.Printf("Bitwise OR of %v: %d\n", curSubArray, mbitwiseOr[i])
			}
		}
	}

	count = len(mbitwiseOr)
	log.Printf("Count: %d based on length of %v\n", count, mbitwiseOr)
	return count
}

// MLEsubarrayBitwiseORs is memory limit exceeded !
func MLEsubarrayBitwiseORs(A []int) int {
	count := 0
	// ws := 1

	allSubArrays := [][]int{}

	for ws := 1; ws <= len(A); ws++ {
		// log.Println("Window size:", ws)
		for i := 0; i < len(A); i++ {
			// log.Println("Index:", i)
			if i+ws <= len(A) {
				allSubArrays = append(allSubArrays, A[i:i+ws])
			}
		}
	}
	log.Printf("All Sub Arrays: %+v\n", allSubArrays)

	mbitwiseOr := map[int]int{}
	for i := 0; i < len(allSubArrays); i++ {
		curSubArray := allSubArrays[i]
		bitwiseOr := 0
		for s := 0; s < len(curSubArray); s++ {
			bitwiseOr |= curSubArray[s]
		}
		mbitwiseOr[bitwiseOr] = bitwiseOr
		log.Printf("Bitwise OR of %v: %d\n", curSubArray, mbitwiseOr[i])
	}

	count = len(mbitwiseOr)
	log.Printf("Count: %d based on length of %v\n", count, mbitwiseOr)
	return count
}
