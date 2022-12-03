package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kennnyz/go_pet/pckg/config"
	Handler "github.com/kennnyz/go_pet/pckg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Get("/", Handler.Repo.Home)
	mux.Get("/about", Handler.Repo.About)
	return mux
}
