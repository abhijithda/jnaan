package magicsquaresingrid

import "fmt"

/* A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.

Given an N x N grid of integers, how many 3 x 3 "magic square" subgrids are there?  (Each subgrid is contiguous).

Example 1:

Input: [[4,3,8,4],
        [9,5,1,9],
        [2,7,6,2]]
Output: 1
Explanation:
The following subgrid is a 3 x 3 magic square:
438
951
276

while this one is not:
384
519
762

In total, there is only one magic square inside the given grid.
Note:

1 <= grid.length = grid[0].length <= 10
0 <= grid[i][j] <= 15
*/
func numMagicSquaresInside(grid [][]int) int {
	magicGrids := 0
	rowLen := len(grid)
	if rowLen <= 0 {
		return magicGrids
	}

	fmt.Println("Row Len: ", rowLen)
	colLen := len(grid[0])
	fmt.Println("Col Len: ", colLen)
	for r := 0; r <= rowLen-3; r++ {
		for c := 0; c <= colLen-3; c++ {

			if grid[r][c] < 1 || grid[r][c] > 9 ||
				grid[r][c+1] < 1 || grid[r][c+1] > 9 ||
				grid[r][c+2] < 1 || grid[r][c+2] > 9 ||
				grid[r+1][c] < 1 || grid[r+1][c] > 9 ||
				grid[r+1][c+1] < 1 || grid[r+1][c+1] > 9 ||
				grid[r+1][c+2] < 1 || grid[r+1][c+2] > 9 ||
				grid[r+2][c] < 1 || grid[r+2][c] > 9 ||
				grid[r+2][c+1] < 1 || grid[r+2][c+1] > 9 ||
				grid[r+2][c+2] < 1 || grid[r+2][c+2] > 9 {
				continue
			}

			r1 := grid[r][c] + grid[r][c+1] + grid[r][c+2]
			r2 := grid[r+1][c] + grid[r+1][c+1] + grid[r+1][c+2]
			r3 := grid[r+2][c] + grid[r+2][c+1] + grid[r+2][c+2]
			if r1 != r2 || r1 != r3 {
				continue
			}

			c1 := grid[r][c] + grid[r+1][c] + grid[r+2][c]
			c2 := grid[r][c+1] + grid[r+1][c+1] + grid[r+2][c+1]
			c3 := grid[r][c+2] + grid[r+1][c+2] + grid[r+2][c+2]
			if r1 != c1 || c1 != c2 || c1 != c3 {
				continue
			}

			d1 := grid[r][c] + grid[r+1][c+1] + grid[r+2][c+2]
			d2 := grid[r][c+2] + grid[r+1][c+1] + grid[r+2][c]

			if r1 != d1 || d1 != d2 {
				continue
			}
			magicGrids++
		}
	}

	return magicGrids
}
