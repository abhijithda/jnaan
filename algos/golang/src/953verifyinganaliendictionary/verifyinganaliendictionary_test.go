package verifyinganaliendictionary

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				words: []string{"hello", "leetcode"},
				order: "hlabcdefgijkmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				words: []string{"word", "world", "row"},
				order: "worldabcefghijkmnpqstuvxyz",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				words: []string{"apple", "app"},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				words: []string{"a", "app"},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "Example 5",
			args: args{
				words: []string{"a", "a"},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "Example 6",
			args: args{
				words: []string{"a", " "},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: false,
		},
		{
			name: "Example 7",
			args: args{
				words: []string{" ", " "},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "Example 8",
			args: args{
				words: []string{"worldo", "word", "wordo"},
				order: "worldabcefghijkmnpqstuvxyz",
			},
			want: true,
		},
		{
			name: "Example 9",
			args: args{
				words: []string{" ", "a", "b"},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "Example 10",
			args: args{
				words: []string{"", " ", "a", "b"},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "Example 10",
			args: args{
				words: []string{"", "a"},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "Example 11",
			args: args{
				words: []string{"a", ""},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: false,
		},
		{
			name: "Example 12",
			args: args{
				words: []string{"a", " "},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
