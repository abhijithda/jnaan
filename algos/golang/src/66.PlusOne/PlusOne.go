package plusone

func plusOne(digits []int) []int {
	res := make([]int, len(digits)+1)
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := carry + digits[i]
		res[i+1] = sum % 10
		carry = sum / 10
	}

	if carry != 0 {
		res[0] = carry
	} else {
		res = res[1:len(res)]
	}
	return res
}
