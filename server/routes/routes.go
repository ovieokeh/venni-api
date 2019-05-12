package venni

import (
	"github.com/go-chi/chi"
	handlers "github.com/ovieokeh/venni-api/server/handlers"
)

// Routes - defines all the routes for the API
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", handlers.IndexHandler)
	router.HandleFunc("/*", handlers.NotFoundHandler)
	return router
}
