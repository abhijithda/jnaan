package symmetrictree

import (
	"internal/tree"
	"testing"
)

// import "testing"
// import "internal/tree"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root: tree.Create([]int{1,2,2,3,4,4,3}),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{1,2,2,-1,3,-1,3}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
