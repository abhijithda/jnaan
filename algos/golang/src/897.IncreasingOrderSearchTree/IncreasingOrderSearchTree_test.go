package increasingordersearchtree

import (
	"reflect"
	"testing"

	it "internal/tree"
)

func Test_increasingBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: it.PrepareTree([]int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9}),
			},
			want: it.PrepareTree([]int{1, -1, 2, -1, 3, -1, 4, -1, 5, -1, 6, -1, 7, -1, 8, -1, 9}),
		},
		{
			name: "Example 2",
			args: args{
				root: it.PrepareTree([]int{}),
			},
			want: it.PrepareTree([]int{}),
		},
		{
			name: "Example 3",
			args: args{
				root: it.PrepareTree([]int{2, 1}),
			},
			want: it.PrepareTree([]int{1, -1, 2}),
		},
		{
			name: "Example 4",
			args: args{
				root: it.PrepareTree([]int{1, -1, 2}),
			},
			want: it.PrepareTree([]int{1, -1, 2}),
		},
		{
			name: "Example 5 Bug",
			args: args{
				root: it.PrepareTree([]int{54, -1, 57, 788, -1, -1, 876}),
			},
			want: it.PrepareTree([]int{54, -1, 788, -1, 876, -1, 57}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
