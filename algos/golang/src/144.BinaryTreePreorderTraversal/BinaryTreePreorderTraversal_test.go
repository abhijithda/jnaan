package binarytreepreordertraversal

import (
	"internal/tree"
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
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
				root: tree.Create([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{}),
			},
			want: []int{},
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{1}),
			},
			want: []int{1},
		},
		{
			name: "Example 4",
			args: args{
				root: tree.Create([]int{1, 2, 3, 4, 5, -1, 7}),
			},
			want: []int{1, 2, 4, 5, 3, 7},
		},
		{
			name: "Example 5",
			args: args{
				root: tree.Create([]int{1, -1, 3, -1, 5}),
			},
			want: []int{1, 3, 5},
		},
		{
			name: "Example 6",
			args: args{
				root: tree.Create([]int{1, 2, -1, 4, -1, 6}),
			},
			want: []int{1, 2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
