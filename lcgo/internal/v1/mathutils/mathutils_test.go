package mathutils

import "testing"

func TestNCR(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n: 5,
				r: 3,
			},
			want: 10,
		},
		{
			name: "Example 2",
			args: args{
				n: 3000,
				r: 3,
			},
			want: 4495501000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NCR(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("NCR() = %v, want %v", got, tt.want)
			}
		})
	}
}
