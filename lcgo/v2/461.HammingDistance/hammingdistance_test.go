package hammingdistance

import "testing"

func Test_hammingDistance(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{x: 1, y: 4},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{x: 1, y: 8},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{x: 1, y: 1},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{x: 1, y: 5},
			want: 1,
		},
		{
			name: "Example 5",
			args: args{x: 1, y: 6},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
