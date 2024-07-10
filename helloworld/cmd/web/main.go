package main

import (
	"fmt"
	"log"
	"net/http"

	"template.com/helloworld/pkg/config"
	"template.com/helloworld/pkg/handlers"
	"template.com/helloworld/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
