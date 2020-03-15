package mathutils

// NCR returns the combination results of n & r.
// I.e., nCr = n!/(r!(n-r)!)
func NCR(n, r int) int {
	val := 1
	// In this problem, r is one of 1,2,3... so, use n-r as that'll bigger
	// 	number compared r to avoid lot of multiplications. (Otherwise,
	//  	we would have to find whether r or n-r is bigger.)
	for c := (n - r) + 1; c <= n; c++ {
		val *= c
	}
	for r > 1 {
		val /= r
		r--
	}
	return val
}
