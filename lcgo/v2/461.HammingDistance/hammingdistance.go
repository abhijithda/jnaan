package hammingdistance

func hammingDistance(x int, y int) int {
	bits := x ^ y
	dist := 0
	for bits != 0 {
		dist++
		bits &= (bits - 1)
	}
	return dist
}
