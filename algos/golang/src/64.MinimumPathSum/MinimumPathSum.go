package MinimumPathSum

import (
	"fmt"
	"log"
	"math"
)

func minPathSum(grid [][]int) int {
	fmt.Println("Given grid:", grid)
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				if grid[i][j-1] < grid[i-1][j] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}
	log.Println("Grid:", grid)
	return grid[m-1][n-1]
}

func minPathSumWithMathMinFunc(grid [][]int) int {
	fmt.Println("Given grid:", grid)
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += int(math.Min(float64(grid[i][j-1]), float64(grid[i-1][j])))
			}
		}
	}
	log.Println("Grid:", grid)
	return grid[m-1][n-1]
}
