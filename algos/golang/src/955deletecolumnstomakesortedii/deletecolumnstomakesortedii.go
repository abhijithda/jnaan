package deletecolumnstomakesortedii

import (
	"log"
)

// One of the answers in LC
func minDeletionSizeLC(A []string) int {
	ans := 0
	cuts := make([]bool, len(A))

	for i := range A[0] {
		tmp := true
		for j := 0; j < len(A)-1; j++ {
			if !cuts[j] && A[j][i] > A[j+1][i] {
				ans++
				tmp = false
				break
			}
		}
		if tmp {
			for k := 0; k < len(A)-1; k++ {
				if A[k][i] < A[k+1][i] {
					cuts[k] = true
				}
			}
		}
	}
	return ans
}

func minDeletionSize(A []string) int {
	log.Println()
	defer log.Println()
	log.Println("Given A:", A)
	numD := 0

	wordlen := 0
	if len(A) != 0 {
		wordlen = len(A[0])
	}

	noCmp := map[string]bool{}
	for j := 0; j < wordlen; j++ {
		prevD := numD
		same := false
		toNoCmp := map[string]bool{}

		for i := 0; i < len(A)-1; i++ {
			log.Printf("A[%d][%d]: %s; A[%d][%d]: %s;",
				i, j, string(A[i][j]), i+1, j, string(A[i+1][j]))
			log.Printf("%v Compare? %v; %v Compare? %v",
				A[i], noCmp[A[i]], A[i+1], noCmp[A[i+1]])
			if noCmp[A[i]] {
				log.Printf("Not comparing!!!")
				continue
			}
			if A[i][j] > A[i+1][j] {
				numD++
				break
			} else if A[i][j] == A[i+1][j] {
				same = true
			} else {
				toNoCmp[A[i]] = true
			}
		}

		if numD == prevD {
			if !same {
				break
			}
			for kvs := range toNoCmp {
				if !noCmp[kvs] {
					noCmp[kvs] = toNoCmp[kvs]
				}
			}
		}
	}

	log.Println("Num of indices that need to be deleted:", numD)
	return numD
}
