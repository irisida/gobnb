package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/irisida/gobnb/pkg/config"
	"github.com/irisida/gobnb/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(sessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/americana", handlers.Repo.Americana)
	mux.Get("/swinging-london", handlers.Repo.SwingingLondon)
	mux.Get("/parisian", handlers.Repo.Parisian)

	mux.Get("/availability", handlers.Repo.Availability)
	mux.Get("/booking", handlers.Repo.Booking)

	mux.Post("/availability", handlers.Repo.PostAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
