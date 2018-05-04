package palindromenumber

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		description string
		num         int
		expRes      bool
	}{
		{
			description: "Positive Palindrome number",
			num:         121,
			expRes:      true,
		},
		{
			description: "Negative number",
			num:         -121,
			expRes:      false,
		},
		{
			description: "Positive non-Palindrome number",
			num:         1211,
			expRes:      false,
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)
		res := isPalindrome(tc.num)
		if res != tc.expRes {
			t.Errorf("got %v, want %v", res, tc.expRes)
		}
	}
}
