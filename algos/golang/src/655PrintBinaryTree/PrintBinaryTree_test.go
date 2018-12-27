package printbinarytree

import (
	"internal/tree"
	"reflect"
	"testing"
)

func Test_printTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Example 1",
			args: args{
				root: tree.Create([]int{1, 2}),
			},
			want: [][]string{
				{"", "1", ""},
				{"2", "", ""},
			},
		},
		{
			name: "Example 2",
			args: args{
				root: tree.Create([]int{1, 2, 3}),
			},
			want: [][]string{
				{"", "1", ""},
				{"2", "", "3"},
			},
		},
		{
			name: "Example 3",
			args: args{
				root: tree.Create([]int{1, 2, 3, 4, -1, -1, 5}),
			},
			want: [][]string{
				{"", "", "", "1", "", "", ""},
				{"", "2", "", "", "", "3", ""},
				{"4", "", "", "", "", "", "5"},
			},
		},
		{
			name: "Example 4",
			args: args{
				root: tree.Create([]int{10, 5, 15, -1, -1, 6, 20}),
			},
			want: [][]string{
				{"", "", "", "10", "", "", ""},
				{"", "5", "", "", "", "15", ""},
				{"", "", "", "", "6", "", "20"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
