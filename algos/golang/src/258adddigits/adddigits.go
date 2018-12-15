package adddigits

// Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

func addDigits(num int) int {
	digit := 0
	for sum := num; sum > 0; {
		digit += sum % 10
		sum = int(sum / 10)
		if digit > 9 {
			digit = int(digit/10) + digit%10
		}
	}
	return digit
}
