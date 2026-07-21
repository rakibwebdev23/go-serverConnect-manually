package review

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) ReviewRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /reviews",
		manager.With(
			http.HandlerFunc(h.GetReview),
		),
	)
}
