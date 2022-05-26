package entities

import (
	"quote/api/app/internal/errdefs"
	"testing"
)

func TestNewQuote(t *testing.T) {
	tests := []struct {
		name           string
		author         string
		expectedAuthor string
		quote          string
		expectErr      bool
		err            error
	}{
		{
			name:           "empty quote",
			author:         "",
			expectedAuthor: "",
			quote:          "",
			err:            errdefs.ErrInvalidQuote,
			expectErr:      true,
		},
		{
			name:           "empty author should set author to Unknown",
			author:         "",
			expectedAuthor: "Unknown",
			quote:          "Some quote",
			err:            nil,
			expectErr:      false,
		},
		{
			name:           "invalid characters should be trimmed from author field",
			author:         "*)#*@Johnny**)@#",
			expectedAuthor: "Johnny",
			quote:          "Some awesome quote",
			err:            nil,
			expectErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewQuote(tt.author, tt.quote)
			if tt.expectErr {
				if err != tt.err {
					t.Errorf("NewQuote(%s, %s) = (%v, %v), wantErr %v", tt.author, tt.quote, actual, err, tt.err)
				}
			} else {
				if actual == nil {
					t.Errorf("NewQuote(%s, %s) = (%v, %v), expected no error", tt.author, tt.quote, actual, err)
				} else {
					if actual.Author != tt.expectedAuthor {
						t.Errorf("NewQuote(%s, %s) = (%v, %v), expected author %s", tt.author, tt.quote, actual, err, tt.expectedAuthor)
					}
					if actual.Quote != tt.quote {
						t.Errorf("NewQuote(%s, %s) = (%v, %v), expected quote %s", tt.author, tt.quote, actual, err, tt.quote)
					}
				}
			}
		})
	}
}
