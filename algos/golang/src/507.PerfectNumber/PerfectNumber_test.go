package perfectnumber

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				num: 28,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				num: 6,
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				num: 5,
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				num: 10,
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				num: 100000000,
			},
			want: false,
		},
		{
			name: "Example 6 Bug",
			args: args{
				num: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
