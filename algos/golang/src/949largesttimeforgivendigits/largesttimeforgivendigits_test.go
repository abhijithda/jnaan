package largesttimeforgivendigits

import "testing"

func Test_largestTimeFromDigits(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				A: []int{1, 2, 3, 4},
			},
			want: "23:41",
		},
		{
			name: "Example 2",
			args: args{
				A: []int{5, 5, 5, 5},
			},
			want: "",
		},
		{
			name: "Example 3",
			args: args{
				A: []int{1, 2, 8, 9},
			},
			want: "19:28",
		},
		{
			name: "Example 4",
			args: args{
				A: []int{1, 5, 6, 6},
			},
			want: "16:56",
		},
		{
			name: "Example 5",
			args: args{
				A: []int{1, 0, 0, 0},
			},
			want: "10:00",
		},
		{
			name: "Example 6",
			args: args{
				A: []int{2, 3, 5, 9},
			},
			want: "23:59",
		},
		{
			name: "Example 7",
			args: args{
				A: []int{0, 0, 0, 0},
			},
			want: "00:00",
		},
		{
			name: "Example 8 (MLE)",
			args: args{
				A: []int{0, 0, 3, 0},
			},
			want: "03:00",
		},
		{
			name: "Example 9",
			args: args{
				A: []int{1, 9, 6, 0},
			},
			want: "19:06",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTimeFromDigits(tt.args.A); got != tt.want {
				t.Errorf("largestTimeFromDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
