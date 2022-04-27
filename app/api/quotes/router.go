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
		router.NewPostRoute("/quotes", hdl.createQuote),
		router.NewGetRoute("/quotes", hdl.getAllQuotes),
		router.NewGetRoute("/quotes/:id", hdl.getQuote),
	}
}
