package flippinganimage

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

func flipAndInvertImage(A [][]int) [][]int {
	n := len(A)
	log.Printf("Array length: %v", n)
	n--
	for r := 0; r < len(A); r++ {
		for c := 0; c <= n/2; c++ {
			if c != n-c {
				A[r][c] ^= A[r][n-c]
				A[r][n-c] ^= A[r][c]
				A[r][c] ^= A[r][n-c]
				A[r][n-c] ^= 1
				log.Printf("[%v,%v]: %v", r, n-c, A[r][n-c])
			}
			A[r][c] ^= 1
			log.Printf("[%v,%v]: %v", r, c, A[r][c])
		}
	}
	return A
}
