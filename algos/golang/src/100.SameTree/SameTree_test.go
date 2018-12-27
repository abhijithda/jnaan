package sametree

import (
	"internal/tree"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				p: tree.Create([]int{1, 2, 3}),
				q: tree.Create([]int{1, 2, 3}),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				p: tree.Create([]int{1, 2}),
				q: tree.Create([]int{1, -1, 2}),
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				p: tree.Create([]int{1, 2, 1}),
				q: tree.Create([]int{1, 1, 2}),
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				p: tree.Create([]int{}),
				q: tree.Create([]int{}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
