package fibonaccinumber

func fib(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}
	prev2 := 0
	prev1 := 1
	var sum int
	for i := 2; i <= N; i++ {
		sum = prev1 + prev2
		prev2 = prev1
		prev1 = sum
	}
	return sum
}
