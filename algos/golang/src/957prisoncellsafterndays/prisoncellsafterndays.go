package prisoncellsafterndays

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

func prisonAfterNDays(cells []int, N int) []int {
	log.Printf("\n\nGiven cells: %+v, N: %d\n", cells, N)
	cellDays := map[int][]int{}
	// While the modulo base is chosen as 14 based on observation in output.
	// As to determining modulo base is not very clear, we can check the map to see if
	// the value already exists. If so, the cycle has started, and we can break the loop...
	n := N % 14
	if n == 0 {
		n = 14
	}
	for d := 1; d <= n; d++ {
		tcells := make([]int, len(cells))
		for i := 1; i != len(cells)-1; i++ {
			if cells[i-1] == cells[i+1] {
				// log.Printf("%d (%d) <-> %d(%d), so occuppied (1)\n", cells[i-1], i-1, cells[i+1], i+1)
				tcells[i] = 1
			} else {
				// log.Printf("%d (%d) <-> %d(%d), so vacant (0)\n", cells[i-1], i-1, cells[i+1], i+1)
				tcells[i] = 0
			}
		}
		cells = tcells
		log.Println("Day", d, cells)
		cellDays[d%14] = cells
	}

	log.Println("Max 14 day cells:", cellDays)

	log.Println("Returning cells:", cellDays[N%14])
	return cellDays[N%14]
}
