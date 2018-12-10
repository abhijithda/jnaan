package deletecolumnstomakesortedii

import "testing"

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				A: []string{"ca", "bb", "ac"},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				A: []string{"xc", "yb", "za"},
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				A: []string{"zyx", "wvu", "tsr"},
			},
			want: 3,
		},
		{
			name: "Example 4",
			args: args{
				A: []string{"", "", ""},
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				A: []string{"xga", "xfb", "yfa"},
			},
			want: 1,
		},
		{
			name: "Example 6",
			args: args{
				A: []string{"abx", "agz", "bgc", "bfc"},
			},
			want: 1,
		},
		{
			name: "Example 7",
			args: args{
				A: []string{"doeeqiy", "yabhbqe", "twckqte"},
			},
			want: 3,
		},
		{
			name: "Example 8",
			args: args{
				A: []string{"bpubea", "dwgjgd", "dqbsih", "tzsegm"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.args.A); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
