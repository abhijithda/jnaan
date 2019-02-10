package brokencalculator

import "testing"

func Test_brokenCalc(t *testing.T) {
	type args struct {
		X int
		Y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				X: 2,
				Y: 3,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				X: 5,
				Y: 8,
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				X: 3,
				Y: 10,
			},
			want: 3,
		},
		{
			name: "Example 4",
			args: args{
				X: 1024,
				Y: 1,
			},
			want: 1023,
		},
		{
			name: "Example 5",
			args: args{
				X: 1024,
				Y: 1024,
			},
			want: 0,
		},
		{
			name: "Example 6",
			args: args{
				X: 10,
				Y: 20,
			},
			want: 1,
		},
		{
			name: "Example 7",
			args: args{
				X: 20,
				Y: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := brokenCalc(tt.args.X, tt.args.Y); got != tt.want {
				t.Errorf("brokenCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
