package venni

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"

	// postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	routes "github.com/ovieokeh/venni-api/server/routes"
)

// Server - represents the structure of the server
type Server struct {
	Router *chi.Mux
	DB     *gorm.DB
}

// Init - sets up the server
func (app *Server) Init(dbConnectionString string) {
	var err error
	app.DB, err = gorm.Open("postgres", dbConnectionString)

	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	defer app.DB.Close()

	app.Router = chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	app.Router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
		cors.Handler,
	)

	app.Router.Route("/api", func(r chi.Router) {
		r.Mount("/v1", routes.Routes())
	})
}

// Run - starts the server
func (app *Server) Run(port string) {
	log.Println("venni listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}
