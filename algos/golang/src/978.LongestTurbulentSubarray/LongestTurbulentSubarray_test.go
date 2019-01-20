package longestturbulentsubarray

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				A: []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				A: []int{4, 8, 12, 16},
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				A: []int{100},
			},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{
				A: []int{9, 4, 2, 10, 7, 8, 8, 1, 9, 7, 9, 7, 9, 7, 9},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.A); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
