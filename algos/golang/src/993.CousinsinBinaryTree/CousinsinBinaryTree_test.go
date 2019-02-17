package cousinsinbinarytree

import (
	"internal/tree"
	"testing"
)

func Test_isCousins(t *testing.T) {
	type args struct {
		root *TreeNode
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root: tree.Create([]int{1, 2, 3, 4}),
				x:    4,
				y:    3,
			},
			want: false,
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{1, 2, 3, -1, 4, -1, 5}),
				x:    5,
				y:    4,
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{1, 2, 3, -1, 4}),
				x:    2,
				y:    3,
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				root: tree.Create([]int{}),
				x:    2,
				y:    3,
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				root: tree.Create([]int{1, 2, 3, 4, 5, 6, 7}),
				x:    5,
				y:    7,
			},
			want: true,
		},
		{
			name: "Example 6",
			args: args{
				root: tree.Create([]int{1, 2, 3, 4, 5, 6, 7}),
				x:    5,
				y:    4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
