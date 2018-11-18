package validmountainarray

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				A: []int{2, 1},
			},
			want: false,
		},
		{
			name: "Example 2",
			args: args{
				A: []int{3, 5, 5},
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				A: []int{0, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				A: []int{0, 3, 2, 1, 3},
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				A: []int{0, 3, 2, 2, 1},
			},
			want: false,
		},
		{
			name: "Example 6",
			args: args{
				A: []int{0, 3, 2},
			},
			want: true,
		},
		{
			name: "Example 7",
			args: args{
				A: []int{},
			},
			want: false,
		},
		{
			name: "Example 8",
			args: args{
				A: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "Example 9",
			args: args{
				A: []int{3, 2, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.A); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
