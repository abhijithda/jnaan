package powerfulintegers

import (
	"reflect"
	"sort"
	"testing"
)

func Test_powerfulIntegers(t *testing.T) {
	type args struct {
		x     int
		y     int
		bound int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				x:     2,
				y:     3,
				bound: 10,
			},
			want: []int{2, 3, 4, 5, 7, 9, 10},
		},
		{
			name: "Example 2",
			args: args{
				x:     3,
				y:     5,
				bound: 15,
			},
			want: []int{2, 4, 6, 8, 10, 14},
		},
		{
			name: "Example 3",
			args: args{
				x:     100,
				y:     100,
				bound: 1000000,
			},
			want: []int{2, 101, 200, 10001, 10100, 20000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := powerfulIntegers(tt.args.x, tt.args.y, tt.args.bound)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("powerfulIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
