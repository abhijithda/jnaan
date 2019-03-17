package complementofbase10integer

import "testing"

func Test_bitwiseComplement(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				N: 5,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				N: 7,
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				N: 10,
			},
			want: 5,
		},
		{
			name: "Example 4",
			args: args{
				N: 0,
			},
			want: 1,
		},
		{
			name: "Example 5",
			args: args{
				N: 1000000000,
			},
			want: 73741823,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplement(tt.args.N); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
