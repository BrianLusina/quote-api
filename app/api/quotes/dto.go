package url

type QuoteDto struct {
	author string
	quote  string
}

type CreateQuoteRequestDto struct {
	author string
	quote  string
}

type QuoteResponseDto struct {
	quotes []QuoteDto
}
