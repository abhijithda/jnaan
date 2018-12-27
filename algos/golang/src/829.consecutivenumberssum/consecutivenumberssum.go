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
	for i := 2; (i*i)/2 <= N; i++ {
		midDigit := int(N / i)
		sum := midDigit
		nums := []int{midDigit}
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
			nums = append(nums, d)
		}
		for j := 3; j <= i; j += 2 {
			d := midDigit - (j / 2)
			sum += d
			nums = append(nums, d)
		}
		if sum == N {
			log.Println("New way", i, ":", nums)
			ways++
		}
	}
	return ways
}

func consecutiveNumbersSumOld(N int) int {
	variations := 1

	startNum := 1
	sum := startNum
	nextNum := startNum + 1
	log.Printf("Sum: %v, Next Num: %v\n", sum, nextNum)
	for nextNum <= N/2+1 {
		if sum+nextNum > N {
			log.Printf("Sum(%v) + nextNum(%v) is greater than N (%v)\n", sum, nextNum, N)
			sum -= startNum
			startNum++
			log.Printf("Decremented sum(%v), and Incremented startNum(%v)\n",
				sum, startNum)
			if startNum == nextNum {
				nextNum++
				log.Printf("Start Num was same as nextNum, so incremented nextNum (%v)\n", nextNum)
			}
			// continue
		} else {
			sum += nextNum
			log.Printf("Sum: %v, Next Num: %v\n", sum, nextNum)

			if sum == N {
				variations++
				log.Printf("Variations (start: %v, end: %v): %v\n",
					startNum, nextNum, variations)
				// continue
			}
			nextNum++
		}
	}

	log.Printf("\n\nTotal Variations for N(%v): %v\n\n\n", N, variations)
	return variations

	// variations := 1

	// for start := 1; start < N; start++ {
	// 	sum := start
	// 	nextNum := start + 1
	// 	log.Printf("Sum: %v, Next Num: %v", sum, nextNum)
	// 	for sum <= N {
	// 		if sum == N {
	// 			variations++
	// 			break
	// 		}
	// 		sum += nextNum
	// 		nextNum++
	// 	}
	// }
	// return variations
}

func mainOld() {
	tests := []struct {
		description   string
		num           int
		sumVariations int
	}{
		{
			description:   "5 is 2",
			num:           5,
			sumVariations: 2,
		},
		{
			description:   "9 is 3",
			num:           9,
			sumVariations: 3,
		},
		{
			description:   "15 is 4",
			num:           15,
			sumVariations: 4,
		},
	}

	for _, tc := range tests {
		log.Println(tc.description)
		res := consecutiveNumbersSumOld(tc.num)
		if res != tc.sumVariations {
			log.Printf("got %v, want %v", res, tc.sumVariations)
		}
	}

}
