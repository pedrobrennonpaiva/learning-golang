package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorAPI struct {
	Err string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func TreatError(w http.ResponseWriter, r *http.Response) {
	var errorApi ErrorAPI
	json.NewDecoder(r.Body).Decode(&errorApi)
	JSON(w, r.StatusCode, errorApi)
}
