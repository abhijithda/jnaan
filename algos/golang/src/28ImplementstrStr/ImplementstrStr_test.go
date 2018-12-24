package implementstrstr

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				haystack: "abc",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				haystack: "aaaa",
				needle:   "aaaa",
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				haystack: "",
				needle:   "aaaa",
			},
			want: -1,
		},
		{
			name: "Example 5",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			name: "Example 6",
			args: args{
				haystack: "mississippi",
				needle:   "issi",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
