package MinCostClimbingStairs

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{ cost: []int{10, 15, 20}},
			want: 15,
		},
		{
			name: "Example 2",
			args: args{ cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			want: 6,
		},
		{
			name: "Example 3 - bug",
			args: args{ cost: []int{0,0,0,1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}