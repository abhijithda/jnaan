package romantointeger

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		description string
		roman       string
		num         int
	}{
		{
			description: "Simple add num",
			roman:       "III",
			num:         3,
		},
		{
			description: "Subtract 1 from 5",
			roman:       "IV",
			num:         4,
		},
		{
			description: "Subtract 1 from 10",
			roman:       "IX",
			num:         9,
		},
		{
			description: "Add different nums",
			roman:       "LVIII",
			num:         58,
		},
		{
			description: "Subtract multiple nums",
			roman:       "MCMXCIV",
			num:         1994,
		},
	}

	for _, tc := range tests {
		t.Logf("%s", tc.description)
		res := romanToInt(tc.roman)
		if res != tc.num {
			t.Errorf("got %v, want %v\n", res, tc.num)
		}
	}
}
