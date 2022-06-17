package quotes

import (
	"quote/api/app/internal/core/services/quotesvc"
	"quote/api/app/server/router"
	"quote/api/app/utils/cache"
)

// quoteHandler is a router for quotes handler
type quotesRouter struct {
	svc    *quotesvc.QuoteSvc
	cache  *cache.Cache
	routes []router.Route
}

// NewQuoteHandler initializes a new router
func NewQuotesRouter(cache *cache.Cache, svc *quotesvc.QuoteSvc) router.Router {
	r := &quotesRouter{
		cache: cache,
		svc:   svc,
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
