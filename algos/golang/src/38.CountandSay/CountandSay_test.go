package countandsay

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			name: "Example 2",
			args: args{
				n: 2,
			},
			want: "11",
		},
		{
			name: "Example 3",
			args: args{
				n: 3,
			},
			want: "21",
		},
		{
			name: "Example 4",
			args: args{
				n: 4,
			},
			want: "1211",
		},
		{
			name: "Example 5",
			args: args{
				n: 5,
			},
			want: "111221",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
