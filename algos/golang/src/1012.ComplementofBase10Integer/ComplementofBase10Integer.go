package complementofbase10integer

import (
	"math"
)

func bitwiseComplement(N int) int {
	// rbits: bits in reverse order
	rbits := []int{}
	if N == 0 {
		rbits = append(rbits, 0)
	}
	for N > 0 {
		r := N % 2
		rbits = append(rbits, r)
		N /= 2
	}

	for i, b := range rbits {
		if b == 1 {
			rbits[i] = 0
		} else {
			rbits[i] = 1
		}
	}

	res := 0
	for i, b := range rbits {
		if b == 1 {
			res += int(math.Pow(2, float64(i)))
		}
	}
	return res
}
