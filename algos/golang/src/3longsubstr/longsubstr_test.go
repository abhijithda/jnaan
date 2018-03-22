package longsubstr

import "testing"
import "fmt"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		description string
		str         string
		len         int
	}{
		{
			description: "Null string",
		},
		{
			description: "Single char",
			str:         "a",
			len:         1,
		},
		{
			description: "3 continous char",
			str:         "abcabcbb",
			len:         3,
		},
		{
			description: "One char repeated",
			str:         "bbbbb",
			len:         1,
		},
		{
			description: "3 chars in b/n",
			str:         "pwwkew",
			len:         3,
		},
		{
			description: "3 chars in b/n starting previous match",
			str:         "dvdf",
			len:         3,
		},
	}
	for _, tc := range tests {
		fmt.Println(tc.description)
		l := lengthOfLongestSubstring(tc.str)
		if l != tc.len {
			t.Errorf("got %v, want %v", l, tc.len)
		}
	}
}
