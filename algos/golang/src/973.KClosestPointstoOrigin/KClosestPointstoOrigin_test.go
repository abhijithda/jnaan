package kclosestpointstoorigin

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				points: [][]int{{1, 3}, {-2, 2}},
				K:      1,
			},
			want: [][]int{{-2, 2}},
		},
		{
			name: "Example 2",
			args: args{
				points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
				K:      2,
			},
			want: [][]int{{3, 3}, {-2, 4}},
		},
		{
			name: "Example 3",
			args: args{
				points: [][]int{{3, 3}},
				K:      2,
			},
			want: [][]int{{3, 3}},
		},
		{
			name: "Example 4",
			args: args{
				points: [][]int{},
				K:      1,
			},
			want: [][]int{},
		},
		{
			name: "Example 5",
			args: args{
				points: [][]int{{3, 3}},
				K:      0,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
