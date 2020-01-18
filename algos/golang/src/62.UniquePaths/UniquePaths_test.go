package UniquePaths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1x1",
			args: args{
				m: 1,
				n: 1,
			},
			want: 0,
		},
		{
			name: "Example - 7x3",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "Example - 2x9",
			args: args{
				m: 2,
				n: 9,
			},
			want: 9,
		},
		{
			name: "Example - 50x9",
			args: args{
				m: 50,
				n: 9,
			},
			want: 1652411475,
		},
		{
			name: "Example - 100x100",
			args: args{
				m: 100,
				n: 100,
			},
			want: 4631081169483718960,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
