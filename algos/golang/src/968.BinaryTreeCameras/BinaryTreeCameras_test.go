package binarytreecameras

import (
	"internal/tree"
	"testing"
)

func Test_minCameraCover(t *testing.T) {
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
				root: tree.Create([]int{0, 0, -1, 0, 0}),
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{0, 0, -1, 0, -1, 0, -1, -1, 0}),
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{0, 0, 0, -1, 0, 0}),
			},
			want: 2,
		},
		{
			name: "Example 4",
			args: args{
				root: tree.Create([]int{}),
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				root: tree.Create([]int{0}),
			},
			want: 1,
		},
		{
			name: "Example 6",
			args: args{
				root: tree.Create([]int{0, 0, -1, -1, 0, 0, -1, -1, 0, 0}),
			},
			want: 2,
		},
		{
			name: "Example 7 Bug",
			args: args{
				root: tree.Create([]int{0, -1, 0, -1, 0, -1, 0}),
			},
			want: 2,
		},
		{
			name: "Example 8",
			args: args{
				root: tree.Create([]int{0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0}),
			},
			want: 2,
		},
		{
			name: "Example 9",
			args: args{
				root: tree.Create([]int{0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0}),
			},
			want: 3,
		},
		{
			name: "Example 10 Bug",
			args: args{
				root: tree.Create([]int{0, -1, 0, 0, -1, 0, 0, -1, -1, 0}),
			},
			want: 3,
		},
		{
			name: "Example 11 Bug",
			args: args{
				root: tree.Create([]int{0, -1, 0, -1, 0, 0, 0}),
			},
			want: 2,
		},
		{
			name: "Example 12 Bug",
			args: args{
				root: tree.Create([]int{0, 0, 0, -1, 0, 0, -1, -1, 0}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
