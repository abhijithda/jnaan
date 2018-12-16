package prisoncellsafterndays

import (
	"reflect"
	"testing"
)

func Test_prisonAfterNDays(t *testing.T) {
	type args struct {
		cells []int
		N     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				cells: []int{0, 1, 0, 1, 1, 0, 0, 1},
				N:     7,
			},
			want: []int{0, 0, 1, 1, 0, 0, 0, 0},
		},
		{
			name: "Example 2",
			args: args{
				cells: []int{1, 0, 0, 1, 0, 0, 1, 0},
				N:     1000000000,
			},
			want: []int{0, 0, 1, 1, 1, 1, 1, 0},
		},
		{
			name: "Example 3",
			args: args{
				cells: []int{0, 0, 0, 1, 1, 0, 1, 0},
				N:     574,
			},
			want: []int{0, 0, 0, 1, 1, 0, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prisonAfterNDays(tt.args.cells, tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prisonAfterNDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
