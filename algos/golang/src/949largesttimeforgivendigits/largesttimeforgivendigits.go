package largesttimeforgivendigits

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

func getCombinations(ele string, A []int) []string {
	log.Printf("Got ele: %s; A: %+v", ele, A)
	if len(A) == 0 {
		log.Println("Got combination:", ele)
		return []string{ele}
	}
	var neweles []string
	newelelist := map[string]int{}

	for i := range A {
		newList := []int{}
		if i == 0 {
			newList = A[i+1:]
		} else if i == len(A) {
			newList = A[0:i]
		} else {
			newList = append(newList, A[0:i]...)
			newList = append(newList, A[i+1:len(A)]...)
		}

		newele1 := getCombinations(ele+strconv.Itoa(A[i]), newList)
		log.Println("Left combination:", newele1)
		if len(newele1) != 0 && len(newele1[0]) == 4 {
			log.Println("Appending Left combination:", newele1)
			for n := range newele1 {
				newelelist[newele1[n]] = 1
			}
		}
		newele2 := getCombinations(strconv.Itoa(A[i])+ele, newList)
		log.Println("Left-Right combination:", newele2)
		if len(newele2) != 0 && len(newele2[0]) == 4 {
			log.Println("Appending Left-Right combination:", newele2)
			for n := range newele2 {
				newelelist[newele2[n]] = 1
			}
		}
	}
	for k := range newelelist {
		neweles = append(neweles, k)
	}
	log.Printf("Returning %d combinations: %+v", len(neweles), neweles)
	return neweles
}

func largestTimeFromDigits(A []int) string {
	largestTime := ""
	largestTimeinmm := 0

	combs := getCombinations("", A)
	// combs := getDigitCombinations(A)

	for q := range combs {
		timenum, _ := strconv.ParseInt(combs[q], 10, 0)
		hh := int(timenum) / 100
		mm := int(timenum) % 100
		if hh > 23 || mm > 59 {
			continue
		}
		ntime := fmt.Sprintf("%02d:%02d", hh, mm)
		if largestTimeinmm <= (hh*60)+mm {
			largestTimeinmm = (hh * 60) + mm
			largestTime = ntime
		}
	}
	return largestTime
}

// func getDigitCombinations(A []int) []string {
// 	combs := []string{}
// 	combs = append(combs, getCombinations("", A)...)
// 	log.Println("All digit combinations:", combs)
// 	return combs
// }

// func getDigitCombinationsOld(A []int) []string {
// 	combs := []string{}

// 	for i := range A {
// 		newList := []int{}
// 		if i == 0 {
// 			newList = A[i+1:]
// 		} else if i == len(A) {
// 			newList = A[0 : i-1]
// 		} else {
// 			newList = append(newList, A[0:i]...)
// 			newList = append(newList, A[i+1:len(A)]...)
// 		}

// 		log.Println(">>>>> Sending new list:", newList)
// 		combs = append(combs, getCombinations(strconv.Itoa(A[i]), newList)...)
// 		// combs = append(combs, getComb(true, strconv.Itoa(A[i]), newList)...)
// 		// combs = append(combs, getComb(false, strconv.Itoa(A[i]), newList)...)
// 	}

// 	log.Println("All digit combinations:", combs)
// 	return combs
// }

/*
	1. Form all combinations of given digits as a string.
	2. Validate whether these strings are valid time, and determine their mins value.
	3. Return string with max minutes value.
*/
func largestTimeFromDigitsBug(A []int) string {
	largestTime := ""
	largestTimeinmm := 0
	log.Println("Given digits:", A)
	forms := []string{}
	for i := range A {
		queue := []string{strconv.Itoa(A[i])}
		log.Println("Initiliazing queue with", queue)
		for j := (i + 1) % len(A); j != i; j = (j + 1) % len(A) {
			eles := queue
			queue = []string{}
			for k := range eles {
				ele := eles[k]
				queue = append(queue, ele+strconv.Itoa(A[j]), strconv.Itoa(A[j])+ele)
				// queue = queue[1:]
				log.Println("Appended elements to queue:", queue)
			}
		}
		for q := range queue {
			timenum, _ := strconv.ParseInt(queue[q], 10, 0)
			hh := int(timenum) / 100
			mm := int(timenum) % 100
			if hh > 23 || mm > 59 {
				continue
			}
			ntime := fmt.Sprintf("%02d:%02d", hh, mm)
			forms = append(forms, ntime)
			if largestTimeinmm <= (hh*60)+mm {
				largestTimeinmm = (hh * 60) + mm
				largestTime = ntime
			}
		}
	}

	log.Printf("Valid Forms: %+v\n", forms)
	log.Printf("Largest time: %+v\n", largestTime)
	return largestTime
}

/*
	1. Initialize Queue with specified digits (A[]).
	2. Pop an element from Queue (qele).
	3. Append one of the specified digits that's not already part of the popped element.
	4. See if the new element after append is still lesser than maxtime value.
	4a. Compare the no. of digits in maxTime
	5. If new element is lesser, and if it doesn't contain 4 digits, then add it back to Queue.
	6. If new element is lesser, and if it contain 4 digits, then return that element.
*/
func largestTimeFromDigitsMLE(A []int) string {
	log.Println("Given digits:", A)
	time := ""

	// All zeros?
	allzeros := true
	for a := range A {
		if 0 != A[a] {
			allzeros = false
		}
	}
	if allzeros {
		time = "00:00"
		log.Printf("Time String: %s\n\n\n", time)
		return time
	}

	maxTime := []int{2, 23, 5, 59}

	// 1. Initialize Queue with specified digits (A[]).
	qtime := A
	sort.Sort(sort.Reverse(sort.IntSlice(qtime)))

	for len(qtime) != 0 {
		// 2. Pop an element from Queue (qele).
		qele := qtime[0]
		qtime = qtime[1:]
		log.Println("> Queue ele:", qele)

		// 3. Append one of the specified digits that's not already part of the popped element.
		for a := range A {
			log.Println(">> A[a]:", A[a])

			// Does qele already contains A[a]?
			qcount := 0
			tqele := qele
			var digit int
			for tqele > 0 {
				digit = tqele % 10
				tqele /= 10
				if digit == A[a] {
					// log.Println("Index:", i)
					log.Printf("Contains digit: %d; Queue ele: %d; A[a]: %d;\n", digit, qele, A[a])
					qcount++
					break
				}
			}
			// If qele already contains all of same digits from A[a], then skip appending
			digitFreq := 0
			for ta := range A {
				if A[ta] == digit {
					digitFreq++
				}
			}
			log.Printf("Digit: %d; frequency: %d;\n", digit, digitFreq)
			// if qele already contains same digit with more no. of occurrences, then skip
			if qcount >= digitFreq {
				continue
			}

			// Append digit as it's not part of qele already.
			nqele := qele*10 + A[a]

			// 4. See if the new element after append is still lesser than maxtime value.
			tqele = nqele
			// cd is count digits
			cd := 0
			// For scenario where 0 is the digit, 1 should be answer!
			if tqele == 0 {
				cd = 1
			} else {
				for ; tqele > 0; cd++ {
					tqele /= 10
				}
			}
			log.Printf("%d has %d digits\n", nqele, cd)

			// If length is 4, then any digit would be valid, so return.
			if cd == 4 {
				d1 := nqele % 10
				nqele /= 10
				d2 := nqele % 10
				nqele /= 10
				d3 := nqele % 10
				nqele /= 10
				d4 := nqele % 10
				time = fmt.Sprintf("%d%d:%d%d", d4, d3, d2, d1)
				log.Printf("Time String: %s\n\n\n", time)
				return time
			} else if cd > 4 {
				log.Fatalf("%d is supposed to have just 4 digits and not %d", qele, cd)
			}
			// 4a. Compare the no. of digits in maxTime
			if (cd < 3 && nqele <= maxTime[cd-1]) ||
				(cd == 3 && nqele%10 <= maxTime[cd-1]) {
				qtime = append(qtime, nqele)
				log.Printf("Added %d back to queue %+v\n", nqele, qtime)
			} else {
				log.Printf("%d not added back to queue %+v\n", nqele, qtime)
			}
		}
	}

	log.Printf("Time String: %s\n\n\n", time)
	return time
}
