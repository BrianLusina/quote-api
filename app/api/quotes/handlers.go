package quotes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdl *quotesRouter) getAllQuotes(ctx *gin.Context) {
	quotes, err := hdl.svc.GetAllQuotes()

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	response := make([]QuoteResponseDto, len(quotes))
	for idx, quote := range quotes {
		response[idx] = QuoteResponseDto{
			QuoteDto{
				quote:  quote.Quote,
				author: quote.Author,
			},
		}
	}

	ctx.JSON(http.StatusOK, response)
}
