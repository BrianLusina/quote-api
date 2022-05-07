package quotes

import (
	"quote/api/app/internal/core/domain"
	"quote/api/app/server/router"
)

// quoteHandler is a router for quotes handler
type quotesRouter struct {
	svc    *domain.QuotesUseCase
	routes []router.Route
}

// NewQuoteHandler initializes a new router
func NewQuotesRouter(svc *domain.QuotesUseCase) router.Router {
	r := &quotesRouter{
		svc: svc,
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the controller
func (hdl *quotesRouter) Routes() []router.Route {
	return hdl.routes
}

func (hdl *quotesRouter) initRoutes() {
	hdl.routes = []router.Route{
		router.NewPostRoute("/api/v1/quotes", hdl.createQuote),
		router.NewGetRoute("/api/v1/quotes", hdl.getAllQuotes),
		router.NewGetRoute("/api/v1/quotes/:id", hdl.getQuote),
		router.NewGetRoute("/api/v1/quotes/random", hdl.getRandomQuote),
	}
}
