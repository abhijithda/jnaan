package sqrtx

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				x: 3,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				x: 5,
			},
			want: 2,
		},
		{
			name: "Example 4",
			args: args{
				x: 81,
			},
			want: 9,
		},
		{
			name: "Example 5",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "Example 6",
			args: args{
				x: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
