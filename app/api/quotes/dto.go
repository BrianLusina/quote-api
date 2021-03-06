package quotes

type QuoteDto struct {
	Identifier string `json:"id"`
	Author     string `json:"author" default:"unknown"`
	Quote      string `json:"quote" validate:"required" binding:"required"`
}

type QuoteResponseDto struct {
	Identifier string `json:"id"`
	QuoteDto
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type QuoteResponsesDto []QuoteResponseDto
