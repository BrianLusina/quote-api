package quotes

type QuoteDto struct {
	author string `json:"author"`
	quote  string `json:"quote"`
}

type QuoteResponseDto []QuoteDto
