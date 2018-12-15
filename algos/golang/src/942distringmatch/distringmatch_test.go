package distringmatch

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				S: "IDID",
			},
			want: []int{0, 4, 1, 3, 2},
		},
		{
			name: "Example 2",
			args: args{
				S: "III",
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "Example 3",
			args: args{
				S: "DDI",
			},
			want: []int{3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatch(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
