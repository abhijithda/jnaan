package largesttimeforgivendigits

import (
	"fmt"
	"log"
	"sort"
)

/*
	1. Initialize Queue with specified digits (A[]).
	2. Pop an element from Queue (qele).
	3. Append one of the specified digits that's not already part of the popped element.
	4. See if the new element after append is still lesser than maxtime value.
	4a. Compare the no. of digits in maxTime
	5. If new element is lesser, and if it doesn't contain 4 digits, then add it back to Queue.
	6. If new element is lesser, and if it contain 4 digits, then return that element.
*/
func largestTimeFromDigits(A []int) string {
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
