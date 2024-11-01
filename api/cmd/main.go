package main

import (
	"fmt"
	"golang-api/internal/api"
	"golang-api/internal/config"
	"log"
	"net/http"
)

func main() {
	config := config.Parse()

	router := api.Generate()

	fmt.Printf("Server is running on port: %s\n", config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router)
	if err != nil {
		log.Fatal(err)
	}
}
