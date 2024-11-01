package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/internal/app"
	"webapp/internal/pkg"
)

func main() {
	pkg.LoadTemplates()

	fmt.Println("Running webapp on port 3000")
	router := app.Generate()
	log.Fatal(http.ListenAndServe(":3000", router))
}
