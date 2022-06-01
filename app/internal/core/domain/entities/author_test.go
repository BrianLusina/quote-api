package entities

import "testing"

func TestNewAuthor(t *testing.T) {
	type testCase struct {
		name     string
		author   string
		expected string
	}

	var testCases = []testCase{
		{
			name:     "empty author should set author to Unknown",
			author:   "",
			expected: "Unknown",
		},
		{
			name:     "invalid characters should be trimmed",
			author:   "*)#*@Johnny**)@#",
			expected: "Johnny",
		},
		{
			name:     "invalid characters should be trimmed from and formatted correctly",
			author:   "*)#*@Johnny**)Five@#",
			expected: "Johnny Five",
		},
		{
			name:     "invalid characters should be trimmed and names title cased",
			author:   "*)#*@john303**)doe@#",
			expected: "John Doe",
		},
	}

	for _, tc := range testCases {
		actual, err := newAuthor(tc.author)

		if err != nil {
			t.Errorf("newAuthor(%s) = (%v, %v), expected no error", tc.author, actual, err)
		} else {
			if actual.String() != tc.expected {
				t.Errorf("newAuthor(%s) = (%v, %v), expected author %s", tc.author, actual, err, tc.expected)
			}
		}
	}
}
