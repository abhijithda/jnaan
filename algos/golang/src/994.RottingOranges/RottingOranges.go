package rottingoranges

import (
	"log"
	"reflect"
)

func orangesRotting(grid [][]int) int {
	log.Println("\n\n\nGiven grid:", grid)

	freshPresent := false
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				log.Printf("Fresh oranges are present at (%d,%d)", i, j)
				freshPresent = true
			}
		}
	}
	if !freshPresent {
		return 0
	}

	for min := 0; min != len(grid)*len(grid[0]); min++ {
		log.Println("\nMinute:", min)
		ngrid := make([][]int, len(grid))
		for i := range grid {
			ngrid[i] = make([]int, len(grid[i]))
			copy(ngrid[i], grid[i])
		}
		for i := range grid {
			log.Printf("grid[%d]: %v", i, grid[i])
			for j := range grid[i] {
				log.Printf("grid[%d][%d]: %d", i, j, grid[i][j])
				if grid[i][j] == 2 {
					// Up
					if i != 0 && grid[i-1][j] == 1 {
						ngrid[i-1][j] = 2
						log.Printf("Rotting: ngrid[%d][%d]: %d", i-1, j, ngrid[i-1][j])
					}
					// Right
					if j != len(grid[0])-1 && grid[i][j+1] == 1 {
						ngrid[i][j+1] = 2
						log.Printf("Rotting: ngrid[%d][%d]: %d", i, j+1, ngrid[i][j+1])
					}
					// Down
					if i != len(grid)-1 && grid[i+1][j] == 1 {
						ngrid[i+1][j] = 2
						log.Printf("Rotting: ngrid[%d][%d]: %d", i+1, j, ngrid[i+1][j])
					}
					// Left
					if j != 0 && grid[i][j-1] == 1 {
						ngrid[i][j-1] = 2
						log.Printf("Rotting: ngrid[%d][%d]: %d", i, j-1, ngrid[i][j-1])
					}
				}
			}
		}
		if reflect.DeepEqual(ngrid, grid) == true {
			log.Printf("No fresh oranges were rotten this time:(%d)!", min)
			for i := range ngrid {
				for j := range ngrid[i] {
					if ngrid[i][j] == 1 {
						log.Printf("Fresh oranges still left at (%d,%d)", i, j)
						return -1
					}
				}
			}
			return min
		}
		for i := range grid {
			copy(grid[i], ngrid[i])
		}
	}
	return 0
}
