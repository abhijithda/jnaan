package primepalindrome

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(N int) bool {
	oN := N
	rN := 0
	for N > 0 {
		rN = (rN * 10) + (N % 10)
		N = int(N / 10)
	}
	return rN == oN
}

func primePalindrome(N int) int {
	n := N
	if n <= 2 {
		return 2
	}
	if n%2 == 0 {
		n++
	}
	for ; ; n = n + 2 {
		if isPalindrome(n) {
			if isPrime(n) {
				break
			}
		}
	}
	return n
}
