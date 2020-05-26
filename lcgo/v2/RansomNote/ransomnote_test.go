package ransomnote

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "Example 2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				ransomNote: "",
				magazine:   "",
			},
			want: true,
		},
		{
			name: "Example 5",
			args: args{
				ransomNote: "abaab",
				magazine:   "bbaaabz",
			},
			want: true,
		},
		{
			name: "Example 6",
			args: args{
				ransomNote: "abaabc",
				magazine:   "bbaaabz",
			},
			want: false,
		},
		{
			name: "Example 7",
			args: args{
				ransomNote: "abaab",
				magazine:   "",
			},
			want: false,
		},
		{
			name: "Example 8",
			args: args{
				ransomNote: "",
				magazine:   "abcdef",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
