package longestcommonprefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		description string
		input       []string
		output      string
	}{
		{
			description: "empty strings",
			input:       []string{},
			output:      "",
		},
		{
			description: "No prefix",
			input:       []string{"dog", "racecar", "car"},
			output:      "",
		},
		{
			description: "prefix 'fl'",
			input:       []string{"flower", "flow", "flight"},
			output:      "fl",
		},
		{
			description: "no prefix 'a'",
			input:       []string{"aca", "cba"},
			output:      "",
		},
	}

	for _, tc := range tests {
		t.Logf(tc.description)
		res := longestCommonPrefix(tc.input)
		if res != tc.output {
			t.Errorf("got %v, want %v\n", res, tc.output)
		}
	}
}
