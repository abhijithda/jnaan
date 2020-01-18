package UniquePaths

import "log"

func uniquePaths(m int, n int) int {
	paths := make([][]int, m)
	for i := range paths {
		paths[i] = make([]int, n)
		paths[i][0] = 1
	}
	for i := 0; i < n; i++ {
		paths[0][i] = 1
	}

	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			paths[x][y] = paths[x-1][y] + paths[x][y-1]
		}
	}

	log.Printf("Paths: %++v", paths)

	return paths[m-1][n-1]
}
