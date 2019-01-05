package sqrtx

func mySqrt(x int) int {
	i := 0
	for ; i*i < x; i++ {
	}
	if i*i > x {
		i--
	}
	return i
}
