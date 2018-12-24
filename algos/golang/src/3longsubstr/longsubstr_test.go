package longsubstr

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Null string",
		},
		{
			name: "Single char",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "3 continous char",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "One char repeated",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "3 chars in b/n",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "3 chars in b/n starting previous match",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
