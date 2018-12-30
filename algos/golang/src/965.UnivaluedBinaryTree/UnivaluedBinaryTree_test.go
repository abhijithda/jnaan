package univaluedbinarytree

import (
	"internal/tree"
	"testing"
)

func Test_isUnivalTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root: tree.Create([]int{1, 1, 1, 1, 1, -1, 1}),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{2, 2, 2, 5, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(tt.args.root); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
