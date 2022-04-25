package quotes

import (
	"quote/api/app/internal/core/interactor"
	"quote/api/app/server/router"
)

// quoteHandler is a router for quotes handler
type quotesRouter struct {
	svc    *interactor.QuoteInteractor
	routes []router.Route
}

// NewQuoteHandler initializes a new router
func NewQuotesRouter(svc *interactor.QuoteInteractor) router.Router {
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
		router.NewGetRoute("/quotes", hdl.getAllQuotes),
	}
}
