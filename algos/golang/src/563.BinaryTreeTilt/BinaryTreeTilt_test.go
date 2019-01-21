package binarytreetilt

import (
	"internal/tree"
	"testing"
)

func Test_findTilt(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: tree.Create([]int{1}),
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{1, 2}),
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{1, 2, 3}),
			},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{
				root: tree.Create([]int{1, 2, 3, -2}),
			},
			want: 5,
		},
		{
			name: "Example 5",
			args: args{
				root: tree.Create([]int{0, -1, 2, 3, -2}),
			},
			want: 8,
		},
		{
			name: "Example 6 Bug",
			args: args{
				root: tree.Create([]int{1, 2, 3, 4, -1, 5}),
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
