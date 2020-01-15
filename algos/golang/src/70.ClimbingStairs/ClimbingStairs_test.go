package ClimbingStairs

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example - 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "Example - 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "Example - 5",
			args: args{
				n: 5,
			},
			want: 8,
		},
		{
			name: "Example - 12",
			args: args{
				n: 12,
			},
			want: 233,
		},
		}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
