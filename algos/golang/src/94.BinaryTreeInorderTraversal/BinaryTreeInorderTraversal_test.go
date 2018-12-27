package binarytreeinordertraversal

import (
	"reflect"
	"testing"

	it "internal/tree"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				root: it.PrepareTree([]int{1, -1, 2, 3}),
			},
			want: []int{1, 3, 2},
		},
		{
			name: "Example 2",
			args: args{
				root: it.PrepareTree([]int{}),
			},
			want: []int{},
		},
		{
			name: "Example 3",
			args: args{
				root: it.PrepareTree([]int{1, 2, 3}),
			},
			want: []int{2, 1, 3},
		},
		{
			name: "Example 4",
			args: args{
				root: it.PrepareTree([]int{1, 2, 3, 4}),
			},
			want: []int{4, 2, 1, 3},
		},
		{
			name: "Example 5",
			args: args{
				root: it.PrepareTree([]int{1, 2, -1, 3, 4}),
			},
			want: []int{3, 2, 4, 1},
		},
		{
			name: "Example 6",
			args: args{
				root: it.PrepareTree([]int{1, 2, -1, 3, -1, -1, -1, 4}),
			},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "Example 7",
			args: args{
				root: it.PrepareTree([]int{1, -1, 2, -1, -1, -1, 3, -1, -1, -1, -1, -1, -1, -1, 4}),
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
