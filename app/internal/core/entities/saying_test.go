package entities

import (
	"quote/api/app/internal/errdefs"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sayingTestCase struct {
	name     string
	quote    string
	expected string
	err      error
}

var sayingTestCases = []sayingTestCase{
	{
		name:     "empty quote should return error",
		quote:    "",
		expected: "",
		err:      errdefs.ErrInvalidQuote,
	},
	{
		name:     "invalid characters should be trimmed",
		quote:    "*)#*@Awesome**)@#",
		expected: "Awesome",
		err:      nil,
	},
	{
		name:     "invalid characters should be trimmed from and formatted correctly",
		quote:    "*)#*@Blue**)Sky@#",
		expected: "Blue Sky",
		err:      nil,
	},
	{
		name:     "invalid characters should be removed, but numbers should be included",
		quote:    "*)#*@the 303**)night@# sky",
		expected: "The 303 night sky",
		err:      nil,
	},
}

func TestNewSaying(t *testing.T) {
	for _, tc := range sayingTestCases {
		actual, err := newSaying(tc.quote)

		if tc.err != nil && err == nil {
			t.Fatalf("newSaying(%s) = (%v, %v), expected error", tc.quote, actual, err)
		} else if tc.err == nil && err != nil {
			t.Fatalf("newSaying(%s) = (%v, %v), expected no error", tc.quote, actual, err)
		} else if err != nil && tc.err != nil {
			e := assert.ErrorIs(t, err, tc.err)

			if !e {
				t.Fatalf("newSaying(%s) = (%v, %v), expected error %v", tc.quote, actual, err, tc.err)
			}
		} else {
			if actual.String() != tc.expected {
				t.Fatalf("newSaying(%s) = (%v, %v), expected quote %s", tc.quote, actual, err, tc.expected)
			}
		}
	}
}
