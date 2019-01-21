package kthsmallestelementinabst

import (
	"internal/tree"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: tree.Create([]int{3, 1, 4, -1, 2}),
				k:    1,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{5, 3, 6, 2, 4, -1, -1, 1}),
				k:    3,
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{3, 1, 4, -1, 2}),
				k:    4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
