package secondminimumnodeinabinarytree

import (
	"internal/tree"
	"testing"
)

func Test_findSecondMinimumValue(t *testing.T) {
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
				root: tree.Create([]int{2, 2, 5, -1, -1, 5, 7}),
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{2, 2, 2}),
			},
			want: -1,
		},
		{
			name: "Example 3 Bug",
			args: args{
				root: tree.Create([]int{1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
