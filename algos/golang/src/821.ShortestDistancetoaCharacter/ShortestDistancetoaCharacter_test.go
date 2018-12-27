package shortesttocharacter

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_shortestToCharOld(t *testing.T) {
	tests := []struct {
		description string
		str         string
		c           byte
		indexes     []int
	}{
		{
			description: "4 Occurrences",
			str:         "loveleetcode",
			c:           'e',
			indexes:     []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
	}

	for _, tc := range tests {
		fmt.Println(tc.description)
		indexes := shortestToCharOld(tc.str, tc.c)
		if reflect.DeepEqual(indexes, tc.indexes) {
			t.Errorf("got %v, want %v", indexes[:], tc.indexes[:])
		}
	}

}

func Test_shortestToChar(t *testing.T) {
	type args struct {
		S string
		C byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				S: "loveleetcode",
				C: 'e',
			},
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			name: "Example 1",
			args: args{
				S: "e",
				C: 'e',
			},
			want: []int{0},
		},
		{
			name: "Example 2",
			args: args{
				S: "abcde",
				C: 'e',
			},
			want: []int{4, 3, 2, 1, 0},
		},
		{
			name: "Example 3",
			args: args{
				S: "e1234",
				C: 'e',
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestToChar(tt.args.S, tt.args.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
