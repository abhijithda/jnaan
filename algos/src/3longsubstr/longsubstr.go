package main

import (
	"fmt"
)

// Given a string, find the length of the longest substring without repeating characters.

// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.

// Given "bbbbb", the answer is "b", with the length of 1.

// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	l := 0
	ll := 0
	sp := &s

	var charExists = map[string]int{}
	for i := 0; i < len(s); i++ {
		c := (*sp)[i : i+1]
		// fmt.Printf("Char at %v is %v\n", i, c)
		if _, ok := charExists[c]; ok {
			if ll < l {
				ll = l
			}
			for k := range charExists {
				delete(charExists, k)
			}
			l = 0
		}
		charExists[c] = 1
		l++
	}
	if ll < l {
		ll = l
	}

	return ll
}

func main() {
	s := "abcabcbb"
	fmt.Printf("Length of longest substring of %v: %v\n", s, lengthOfLongestSubstring(s))
	s = "bbbbb"
	fmt.Printf("Length of longest substring of %v: %v\n", s, lengthOfLongestSubstring(s))
	s = "pwwkew"
	fmt.Printf("Length of longest substring of %v: %v\n", s, lengthOfLongestSubstring(s))
	s = "abhijith"
	fmt.Printf("Length of longest substring of %v: %v\n", s, lengthOfLongestSubstring(s))
	s = "dvdf"
	fmt.Printf("Length of longest substring of %v: %v\n", s, lengthOfLongestSubstring(s))
}
