package transposematrix

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

func transpose(A [][]int) [][]int {
	rowLen := len(A)
	colLen := len(A[0])
	log.Println("A Row len: ", rowLen)
	log.Println("A Col len: ", colLen)
	var tA = make([][]int, colLen)
	for i := range tA {
		tA[i] = make([]int, rowLen)
	}
	log.Println("A' Row len: ", len(tA))
	log.Println("A' Col len: ", len(tA[0]))

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			log.Printf("A[%v][%v] = %v", i, j, A[i][j])
			tA[j][i] = A[i][j]
			log.Printf("A'[%v][%v] = %v", j, i, tA[j][i])
		}
	}
	return tA
}
