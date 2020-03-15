package nimgame

import "testing"

func Test_canWinNim(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{n: 1},
			want: true,
		},
		{
			name: "Example 2",
			args: args{n: 2},
			want: true,
		},
		{
			name: "Example 3",
			args: args{n: 3},
			want: true,
		},
		{
			name: "Example 4",
			args: args{n: 4},
			want: false,
		},
		{
			name: "Example 5",
			args: args{n: 5},
			want: true,
		},
		{
			name: "Example 6",
			args: args{n: 6},
			want: true,
		},
		{
			name: "Example 7",
			args: args{n: 7},
			want: true,
		},
		{
			name: "Example 8",
			args: args{n: 8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.args.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
