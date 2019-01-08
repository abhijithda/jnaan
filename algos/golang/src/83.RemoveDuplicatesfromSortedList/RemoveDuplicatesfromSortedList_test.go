package removeduplicatesfromsortedlist

import (
	"reflect"
	"testing"

	"internal/list"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: list.Prepare([]int{}),
			},
			want: list.Prepare([]int{}),
		},
		{
			name: "Example 2",
			args: args{
				head: list.Prepare([]int{1}),
			},
			want: list.Prepare([]int{1}),
		},
		{
			name: "Example 3",
			args: args{
				head: list.Prepare([]int{1, 1}),
			},
			want: list.Prepare([]int{1}),
		},
		{
			name: "Example 4",
			args: args{
				head: list.Prepare([]int{1, 1, 2, 3, 3}),
			},
			want: list.Prepare([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
