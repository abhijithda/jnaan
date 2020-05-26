package longestpalindromicsubstring

import (
	"testing"
)

func Test_longestPalindromeSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantLen int
	}{
		{
			name: "Example 1",
			args: args{
				s: "babad",
			},
			want:    "aba",
			wantLen: 3,
		},
		{
			name: "Example 1-a",
			args: args{
				s: "babbad",
			},
			want:    "abba",
			wantLen: 4,
		},
		{
			name: "Example 2",
			args: args{
				s: "cbbd",
			},
			want:    "bb",
			wantLen: 2,
		},
		{
			name: "Example 3",
			args: args{
				s: "1",
			},
			want:    "1",
			wantLen: 1,
		},
		{
			name: "Example 4",
			args: args{
				s: "",
			},
			want:    "",
			wantLen: 0,
		},
		{
			name: "Example 5",
			args: args{
				s: "babadabab",
			},
			want:    "babadabab",
			wantLen: 9,
		},
		{
			name: "Example 6",
			args: args{
				s: "xyzbabadabab",
			},
			want:    "babadabab",
			wantLen: 9,
		},
		{
			name: "Example 7",
			args: args{
				s: "ac",
			},
			want:    "a",
			wantLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
