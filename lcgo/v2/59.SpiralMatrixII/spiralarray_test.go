package spiralarray

import (
	"reflect"
	"testing"
)

func Test_spiral(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				N: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "Example 2",
			args: args{
				N: 4,
			},
			want: [][]int{
				{01, 02, 03, 04},
				{12, 13, 14, 05},
				{11, 16, 15, 06},
				{10, 9, 8, 07},
			},
		},
		{
			name: "Example 3",
			args: args{
				N: 8,
			},
			want: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{28, 29, 30, 31, 32, 33, 34, 9},
				{27, 48, 49, 50, 51, 52, 35, 10},
				{26, 47, 60, 61, 62, 53, 36, 11},
				{25, 46, 59, 64, 63, 54, 37, 12},
				{24, 45, 58, 57, 56, 55, 38, 13},
				{23, 44, 43, 42, 41, 40, 39, 14},
				{22, 21, 20, 19, 18, 17, 16, 15},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiral(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiral() = %v, want %v", got, tt.want)
			}
		})
	}
}
