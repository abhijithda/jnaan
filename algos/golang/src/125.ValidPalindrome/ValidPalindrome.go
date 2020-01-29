package ValidPalindrome

import "log"

import "strings"

import "unicode"

func isPuncChar(c byte) bool {
	log.Printf("isPuncChar(%v)", c)
	if unicode.IsPunct(rune(c)) || c == ' ' || c == '`' {
		log.Printf("Punc character: %v", c)
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	log.Println("Given s:", s)
	s = strings.ToLower(s)
	b := 0
	e := len(s) - 1
	for b <= e {
		if isPuncChar(s[b]) {
			b++
		} else if isPuncChar(s[e]) {
			e--
		} else if s[b] == s[e] {
			log.Printf("s[%d] equals s[%d]: %v", b, e, string(s[b]))
			b++
			e--
		} else {
			return false
		}
	}
	return true
}
