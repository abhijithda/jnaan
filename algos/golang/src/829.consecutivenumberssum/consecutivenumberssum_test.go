package consecutivenumberssum

import (
	"fmt"
	"testing"
)

func Test_consecutiveNumbersSumOld(t *testing.T) {
	tests := []struct {
		description   string
		num           int
		sumVariations int
	}{
		{
			description:   "5 is 2",
			num:           5,
			sumVariations: 2,
		},
		{
			description:   "9 is 3",
			num:           9,
			sumVariations: 3,
		},
		{
			description:   "15 is 4",
			num:           15,
			sumVariations: 4,
		},
		{
			description:   "933320757 is 16",
			num:           933320757,
			sumVariations: 16,
		},
		{
			description:   "855877922 is 4",
			num:           855877922,
			sumVariations: 4,
		},
	}

	for _, tc := range tests {
		fmt.Println("Testing...:", tc.description)
		res := consecutiveNumbersSum(tc.num)
		if res != tc.sumVariations {
			t.Errorf("got %v, want %v for %v", res, tc.sumVariations, tc.num)
		}
	}
}

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
