package findandreplacepattern

import (
	"log"
	"reflect"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	tests := []struct {
		description string
		words       []string
		pattern     string
		output      []string
	}{
		{
			description: "first sample",
			words:       []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
			pattern:     "abb",
			output:      []string{"mee", "aqq"},
		},
	}

	for _, tc := range tests {
		log.Println(tc.description)

		res := findAndReplacePattern(tc.words, tc.pattern)
		if reflect.DeepEqual(res, tc.output) == false {
			t.Errorf("%s:\n words: %+v pattern: %+v; got %+v, want %+v", tc.description, tc.words, tc.pattern, res, tc.output)
		}

	}

}
