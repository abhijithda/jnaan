package minimumaddtomakeparenthesesvalid

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1:",
			args: args{
				S: "())",
			},
			want: 1,
		},
		{
			name: "Example 2:",
			args: args{
				S: "(((",
			},
			want: 3,
		},
		{
			name: "Example 3:",
			args: args{
				S: "()",
			},
			want: 0,
		},
		{
			name: "Example 4:",
			args: args{
				S: "()))((",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAddToMakeValid(tt.args.S); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
