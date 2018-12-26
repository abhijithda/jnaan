package maximumbinarytree

import (
	"reflect"
	"testing"

	tree "internal/tree"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{3, 2, 1, 6, 0, 5},
			},
			want: tree.PrepareTree([]int{6, 3, 5, -1, 2, 0, -1, -1, 1}),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 3},
			},
			want: tree.PrepareTree([]int{3, 2}),
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{3, 2},
			},
			want: tree.PrepareTree([]int{3, -1, 2}),
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{1},
			},
			want: tree.PrepareTree([]int{1}),
		},
		{
			name: "Example 5",
			args: args{
				nums: []int{},
			},
			want: tree.PrepareTree([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
