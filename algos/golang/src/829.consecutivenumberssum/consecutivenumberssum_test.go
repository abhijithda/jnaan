package consecutivenumberssum

import (
	"testing"
)

func Test_consecutiveNumbersSum(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				N: 5,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				N: 9,
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				N: 15,
			},
			want: 4,
		},
		{
			name: "Example 4 TLE",
			args: args{
				N: 933320757,
			},
			want: 16,
		},
		{
			name: "Example 5 TLE",
			args: args{
				N: 855877922,
			},
			want: 4,
		},
		{
			name: "Example 6 TLE",
			args: args{
				N: 643267314,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consecutiveNumbersSum(tt.args.N); got != tt.want {
				t.Errorf("consecutiveNumbersSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
