package binarytreepaths

import (
	"reflect"
	"testing"

	"internal/tree"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				root: tree.Create([]int{1, 2, 3, -1, 5}),
			},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{1, 2, 3, 4, 5, -1, 7}),
			},
			want: []string{"1->2->4", "1->2->5", "1->3->7"},
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{}),
			},
			want: []string{},
		},
		{
			name: "Example 4",
			args: args{
				root: tree.Create([]int{1}),
			},
			want: []string{"1"},
		},
		{
			name: "Example 5",
			args: args{
				root: tree.Create([]int{1, -1, 2}),
			},
			want: []string{"1->2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
