package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/waldyrfelix/mockapi/app/handlers/actor"
	// "github.com/waldyrfelix/mockapi/app/models"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// setRouters sets the all required routers
func (a *App) SetRouters() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/actors", actor.GetAll).Methods("GET")

	// Routing for handling the projects
	// a.Get("/actors", a.handleRequest(actor.GetAll))
	// a.Post("/projects", a.handleRequest(handler.CreateProject))
	// a.Get("/projects/{title}", a.handleRequest(handler.GetProject))
	// a.Put("/projects/{title}", a.handleRequest(handler.UpdateProject))
	// a.Delete("/projects/{title}", a.handleRequest(handler.DeleteProject))
	// a.Put("/projects/{title}/archive", a.handleRequest(handler.ArchiveProject))
	// a.Delete("/projects/{title}/archive", a.handleRequest(handler.RestoreProject))
}

// Setup middlewares
func (a *App) SetMiddlewares() {
	a.Router.Use(JsonMiddleware)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
