package spiralarray

import "log"

func spiral(N int) [][]int {
	res := make([][]int, N)
	for i := range res {
		res[i] = make([]int, N)
	}
	log.Printf("Initialize array: %+v", res)

	loop := 0
	cnt := 1
	for loop < N/2 {
		i, j := loop, loop
		// log.Printf("loop: %d; i:%d; j:%d, cnt:%d", loop, i, j, cnt)
		// Right
		for ; j < N-1-loop; j++ {
			res[i][j] = cnt
			cnt++
		}
		// log.Printf("Loop %d; Right: %+v", loop, res)
		// Down
		for ; i < N-1-loop; i++ {
			res[i][j] = cnt
			cnt++
		}
		// log.Printf("Loop %d; Down: %+v", loop, res)
		// Left
		for ; j > loop; j-- {

			res[i][j] = cnt
			cnt++
		}
		// log.Printf("Loop %d; Left: %+v", loop, res)
		// Up
		for ; i > loop; i-- {
			res[i][j] = cnt
			cnt++
		}
		// log.Printf("Loop %d; Up: %+v", loop, res)
		loop++
	}
	if N%2 == 1 {
		res[N/2][N/2] = cnt
	}
	return res
}
