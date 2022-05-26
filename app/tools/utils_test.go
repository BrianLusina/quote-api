package tools

import "testing"

func TestStandardizeSpaces(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "string with spaces", input: "This is a      Test", expected: "This is a Test"},
		{name: "string with spaces at the ends", input: "This is another      Test   ", expected: "This is another Test"},
	}

	for _, tc := range testCases {
		actual := StandardizeSpaces(tc.input)

		if actual != tc.expected {
			t.Errorf("StandardizeSpaces(%s) = %s, expected %s", tc.input, actual, tc.expected)
		}
	}
}
