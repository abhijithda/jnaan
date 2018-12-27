package leafsimilartrees

import (
	"internal/tree"
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root1: tree.Create([]int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4}),
				root2: tree.Create([]int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8}),
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				root1: tree.Create([]int{3, 5}),
				root2: tree.Create([]int{3, 5, 1, -1, -1, -1, -1}),
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				root1: tree.Create([]int{3, 1, 5}),
				root2: tree.Create([]int{3, 5, 1, -1, -1, -1, -1}),
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				root1: tree.Create([]int{3, 1, 5}),
				root2: tree.Create([]int{}),
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				root1: tree.Create([]int{}),
				root2: tree.Create([]int{}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
