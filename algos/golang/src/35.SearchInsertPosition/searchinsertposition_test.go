package searchinsertposition

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
		{
			name: "Example 4",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: 0,
		},
		{
			name: "Example 6 TLE",
			args: args{
				nums:   []int{1, 3},
				target: 0,
			},
			want: 0,
		},
		{
			name: "Example 7 Bug",
			args: args{
				nums:   []int{1, 3},
				target: 2,
			},
			want: 1,
		},
		{
			name: "Example 8 TLE",
			args: args{
				nums:   []int{3, 5, 7, 9, 10},
				target: 8,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
