package plusone

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "Example 2",
			args: args{
				digits: []int{4, 3, 2, 1},
			},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "Example 3",
			args: args{
				digits: []int{4, 3, 2, 9},
			},
			want: []int{4, 3, 3, 0},
		},
		{
			name: "Example 4",
			args: args{
				digits: []int{4, 3, 9, 9},
			},
			want: []int{4, 4, 0, 0},
		},
		{
			name: "Example 5",
			args: args{
				digits: []int{9, 9, 9, 9},
			},
			want: []int{1, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
