package MinimumPathSum

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 3x3",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 7,
		},
		{
			name: "Example 1x1",
			args: args{
				grid: [][]int{
					{4},
				},
			},
			want: 4,
		},
		{
			name: "Example 3x4",
			args: args{
				grid: [][]int{
					{1, 3, 1, 1},
					{1, 5, 1, 3},
					{4, 2, 1, 6},
					{5, 4, 3, 2},
				},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
