package longestpalindromicsubstring

import "log"

func longestPalindromeSubstring(s string) string {
	log.Println("Given s:", s)
	if len(s) < 2 {
		return s
	} else if len(s) == 2 {
		if s[1] != s[0] {
			return string(s[0])
		}
		return s
	}

	maxLen := 0
	centerIndex := 0
	for i := 1; i < len(s)-1; i++ {
		b := i
		e := i
		// For case when palindrome length is even.
		if s[i] == s[i-1] {
			b = i - 1
			l := getLongestPalindrome(s, b, e)
			if maxLen < l {
				maxLen = l
				centerIndex = b
			}
			b = i
		}
		// For case when palindrome length is even.
		if s[i] == s[i+1] {
			e = i + 1
			l := getLongestPalindrome(s, b, e)
			if maxLen < l {
				maxLen = l
				centerIndex = e
			}
			e = i
		}
		// For case when palindrome length is odd.
		// I.e., only one middle element.
		l := getLongestPalindrome(s, b, e)
		if maxLen < l {
			maxLen = l
			centerIndex = b
		}
	}
	log.Printf("centerIndex: %v, maxLen: %v", centerIndex, maxLen)
	b := centerIndex + (maxLen % 2) - (maxLen / 2)
	e := b + maxLen
	if e > len(s) {
		b--
		e--
	}
	return s[b:e]
}

func getLongestPalindrome(s string, b, e int) int {

	for b > 0 && e+1 < len(s) {
		if s[b-1] != s[e+1] {
			break
		}
		b--
		e++
	}
	return e - b + 1
}
