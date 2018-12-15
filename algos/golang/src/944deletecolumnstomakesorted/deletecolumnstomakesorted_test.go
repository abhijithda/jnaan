package deletecolumnstomakesorted

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
				A: []string{"cba", "daf", "ghi"},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				A: []string{"a", "b"},
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
				A: []string{},
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				A: []string{"c", "b"},
			},
			want: 1,
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
