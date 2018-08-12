package UncommonWordsfromTwoSentences

import (
	"log"
	"reflect"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	tests := []struct {
		description string
		A           string
		B           string
		output      []string
	}{
		{
			description: "First sample",
			A:           "this apple is sweet",
			B:           "this apple is sour",
			output:      []string{"sweet", "sour"},
		},
		{
			description: "Second sample",
			A:           "apple apple",
			B:           "banana",
			output:      []string{"banana"},
		},
	}

	for _, tc := range tests {
		log.Println(tc.description)

		res := uncommonFromSentences(tc.A, tc.B)
		if reflect.DeepEqual(res, tc.output) == false {
			t.Errorf("got %v, want %v", res, tc.output)
		}
	}
}
