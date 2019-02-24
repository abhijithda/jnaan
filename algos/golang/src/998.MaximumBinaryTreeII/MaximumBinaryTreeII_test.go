package maximumbinarytreeii

import (
	"internal/tree"
	"reflect"
	"testing"
)

func Test_insertIntoMaxTree(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: tree.Create([]int{4, 1, 3, -1, -1, 2}),
				val:  5,
			},
			want: tree.Create([]int{5, 4, -1, 1, 3, -1, -1, 2}),
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{5, 2, 4, -1, 1}),
				val:  3,
			},
			want: tree.Create([]int{5, 2, 4, -1, 1, -1, 3}),
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{5, 2, 3, -1, 1}),
				val:  4,
			},
			want: tree.Create([]int{5, 2, 4, -1, 1, 3}),
		},
		{
			name: "Example 4",
			args: args{
				root: tree.Create([]int{}),
				val:  1,
			},
			want: tree.Create([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoMaxTree(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoMaxTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
