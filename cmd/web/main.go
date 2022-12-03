package main

import (
	"fmt"
	"github.com/kennnyz/go_pet/pckg/config"
	Handler "github.com/kennnyz/go_pet/pckg/handler"
	"github.com/kennnyz/go_pet/pckg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(" Can not create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := Handler.NewRepo(&app)
	Handler.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", Handler.Repo.Home)
	//http.HandleFunc("/about", Handler.Repo.About)

	fmt.Printf("Staring application on port %s", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
