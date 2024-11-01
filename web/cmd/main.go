package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/internal/router"
)

func main() {
	fmt.Println("Running webapp on port 3000")

	router := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", router))
}
