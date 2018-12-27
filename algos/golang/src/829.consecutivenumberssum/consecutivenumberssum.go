package consecutivenumberssum

import (
	"log"
	"os"
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

func consecutiveNumbersSum(N int) int {
	log.Println("\n\n\nGiven N:", N)
	log.Println("New way 1 :", N)

	ways := 1
	for i := 2; (i*(i+1))/2 <= N; i++ {
		if (N-(i*(i+1)/2))%i == 0 {
			log.Println("New way", i)
			ways++
		}
	}
	return ways
}

func consecutiveNumbersSumTLE(N int) int {
	log.Println("\n\n\nGiven N:", N)
	log.Println("New way 1 :", N)

	ways := 1
	for i := 2; (i*i)/2 <= N; i++ {
		midDigit := int(N / i)
		sum := midDigit
		// nums := []int{midDigit}
		// start j from 2 as sum already has one digit (i.e., midDigit).
		// for j := 2; j <= i; j++ {
		// 	var d int
		// 	if j%2 == 0 {
		// 		d = midDigit + (j / 2)
		// 	} else {
		// 		d = midDigit - (j / 2)
		// 	}
		// 	sum += d
		// 	// nums = append(nums, d)
		// }

		// INFO: Same logic as above, but is little bit more efficient as
		// 	the mod is not being done.
		for j := 2; j <= i; j += 2 {
			d := midDigit + (j / 2)
			sum += d
			// nums = append(nums, d)
		}
		for j := 3; j <= i; j += 2 {
			d := midDigit - (j / 2)
			sum += d
			// nums = append(nums, d)
		}
		if sum == N {
			// log.Println("New way", i, ":", nums)
			ways++
		}
	}
	return ways
}
