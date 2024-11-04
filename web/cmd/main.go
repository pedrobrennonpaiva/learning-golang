package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/internal/app"
	"webapp/internal/config"
	"webapp/internal/pkg"
	"webapp/internal/pkg/cookies"
)

func main() {
	config := config.Parse()
	cookies.Configure()

	pkg.LoadTemplates()
	router := app.Generate()

	fmt.Printf("Server is running on port: %s\n", config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router)
	if err != nil {
		log.Fatal(err)
	}
}
