package oneeditdistance

import "testing"

func Test_isOneEditDistance(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s: "ab",
				t: "acb",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s: "cab",
				t: "ad",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				s: "1203",
				t: "1213",
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				s: "ab",
				t: "ab",
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				s: "ab",
				t: "abc",
			},
			want: true,
		},
		{
			name: "Example 6",
			args: args{
				s: "ab",
				t: "abcd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneEditDistance(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isOneEditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
