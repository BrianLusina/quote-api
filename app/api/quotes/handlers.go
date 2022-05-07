package quotes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (hdl *quotesRouter) createQuote(ctx *gin.Context) {
	var request QuoteDto
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	quote, err := hdl.svc.CreateQuote(request.Author, request.Quote)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := QuoteResponseDto{
		Identifier: quote.ID.String(),
		QuoteDto: QuoteDto{
			Quote:  quote.Quote,
			Author: quote.Author,
		},
		CreatedAt: quote.CreatedAt.Format(time.RFC3339),
		UpdatedAt: quote.UpdatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, response)
}

func (hdl *quotesRouter) getQuote(ctx *gin.Context) {
	id := ctx.Param("id")
	quote, err := hdl.svc.GetQuote(id)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	response := QuoteResponseDto{
		Identifier: quote.ID.String(),
		QuoteDto: QuoteDto{
			Quote:  quote.Quote,
			Author: quote.Author,
		},
		CreatedAt: quote.CreatedAt.Format(time.RFC3339),
		UpdatedAt: quote.UpdatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, response)
}

func (hdl *quotesRouter) getAllQuotes(ctx *gin.Context) {
	quotes, err := hdl.svc.GetAllQuotes()

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	response := QuoteResponsesDto{}
	for _, quote := range quotes {
		response = append(response, QuoteResponseDto{
			Identifier: quote.ID.String(),
			QuoteDto: QuoteDto{
				Quote:  quote.Quote,
				Author: quote.Author,
			},
			CreatedAt: quote.CreatedAt.Format(time.RFC3339),
			UpdatedAt: quote.UpdatedAt.Format(time.RFC3339),
		})
	}

	ctx.JSON(http.StatusOK, response)
}

func (hdl *quotesRouter) getRandomQuote(ctx *gin.Context) {
	quote, err := hdl.svc.GetRandomQuote()

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	response := QuoteResponseDto{
		Identifier: quote.ID.String(),
		QuoteDto: QuoteDto{
			Quote:  quote.Quote,
			Author: quote.Author,
		},
		CreatedAt: quote.CreatedAt.Format(time.RFC3339),
		UpdatedAt: quote.UpdatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, response)
}
