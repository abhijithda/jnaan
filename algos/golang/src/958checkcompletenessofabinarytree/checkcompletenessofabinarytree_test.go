package checkcompletenessofabinarytree

import (
	"internal/tree"
	"testing"
)

func Test_isCompleteTree(t *testing.T) {
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
				root: tree.Create([]int{1, 2, 3, 4, 5, 6}),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{1, 2, 3, 4, 5, -1, 7}),
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{}),
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				root: tree.Create([]int{1}),
			},
			want: true,
		},
		{
			name: "Example 5",
			args: args{
				root: tree.Create([]int{1, 2}),
			},
			want: true,
		},
		{
			name: "Example 5",
			args: args{
				root: tree.Create([]int{1, -1, 2}),
			},
			want: false,
		},
		{
			name: "Example 6",
			args: args{
				root: tree.Create([]int{1, 2, 3, -1, -1, 7, 8}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCompleteTree(tt.args.root); got != tt.want {
				t.Errorf("isCompleteTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
