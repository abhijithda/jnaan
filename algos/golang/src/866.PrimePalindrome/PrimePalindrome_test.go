package primepalindrome

import (
	"fmt"
	"log"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		description string
		N           int
		res         bool
	}{
		{
			description: "Palindrome integer",
			N:           101,
			res:         true,
		},
		{
			description: "Non-Palindrome integer",
			N:           10,
			res:         false,
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)
		log.Println(tc.description)

		res := isPalindrome(tc.N)
		if res != tc.res {
			t.Errorf("got %v, want %v", res, tc.res)
		}
	}

}

func Test_isPrime(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				p: 101,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.p); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_primePalindrome(t *testing.T) {
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
				N: 1,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				N: 2,
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				N: 5,
			},
			want: 5,
		},
		{
			name: "Example 4",
			args: args{
				N: 8,
			},
			want: 11,
		},
		{
			name: "Example 5",
			args: args{
				N: 90,
			},
			want: 101,
		},
		{
			name: "Example 6",
			args: args{
				N: 100,
			},
			want: 101,
		},
		{
			name: "Example 7",
			args: args{
				N: 200,
			},
			want: 313,
		},
		{
			name: "Example 8",
			args: args{
				N: 9989900,
			},
			want: 100030001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primePalindrome(tt.args.N); got != tt.want {
				t.Errorf("primePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
